package main

import (
	"encoding/json"
	"github.com/xuzhenglun/Turling-Go/Turling"
	"golang.org/x/net/websocket"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	listenAddr = "localhost:4000" // server address
	key        = "1172c3986ecaeb20ec066284eb35b041"
	address    = `Http://www.tuling123.com/openapi/api`
	publicAddr string
)

type config struct {
	ListenAddr string
	Key        string
	Address    string
	PublicAddr string
}

var (
	pwd, _   = os.Getwd()
	RootTemp = template.Must(template.ParseFiles(pwd + "/chat.html"))
	//JSON     = websocket.JSON // codec for JSON
	Message       = websocket.Message        // codec for string, []byte
	ActiveClients = make(map[ClientConn]int) // map containing clients
)

// Initialize handlers and websocket handlers
func init() {
	http.HandleFunc("/", RootHandler)
	http.Handle("/sock", websocket.Handler(SockServer))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}

// Client connection consists of the websocket and the client ip
type ClientConn struct {
	websocket *websocket.Conn
	clientIP  string
}

// WebSocket server to handle chat between clients
func SockServer(ws *websocket.Conn) {
	var err error
	var clientMessage string
	// use []byte if websocket binary type is blob or arraybuffer
	// var clientMessage []byte

	// cleanup on server side
	defer func() {
		if err = ws.Close(); err != nil {
			log.Println("Websocket could not be closed", err.Error())
		}
	}()

	client := ws.Request().RemoteAddr
	log.Println("Client connected:", client)
	sockCli := ClientConn{ws, client}
	ActiveClients[sockCli] = 0
	log.Println("Number of clients connected ...", len(ActiveClients))

	// for loop so the websocket stays open otherwise
	// it'll close after one Receieve and Send
	Robot := Turling.New(address, key)
	for {
		if err = Message.Receive(ws, &clientMessage); err != nil {
			// If we cannot Read then the connection is closed
			log.Println("Websocket Disconnected waiting", err.Error())
			// remove the ws client conn from our active clients
			delete(ActiveClients, sockCli)
			log.Println("Number of clients still connected ...", len(ActiveClients))
			return
		}

		Reply := Robot.Reply(clientMessage)

		for cs, _ := range ActiveClients {
			if err = Message.Send(cs.websocket, Reply); err != nil {
				// we could not send the message to a peer
				log.Println("Could not send message to ", cs.clientIP, err.Error())
			}
		}
	}
}

// RootHandler renders the template for the root page
func RootHandler(w http.ResponseWriter, req *http.Request) {
	if publicAddr == "" {
		publicAddr = listenAddr
	}
	err := RootTemp.Execute(w, publicAddr)
	log.Println("Public Address is "+publicAddr, listenAddr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	cfg := getconfig()
	if cfg != nil {
		if cfg.ListenAddr != "" {
			listenAddr = cfg.ListenAddr
		}
		if cfg.Key != "" {
			key = cfg.Key
		}
		if cfg.Address != "" {
			address = cfg.Address
		}
		if cfg.PublicAddr != "" {
			publicAddr = cfg.PublicAddr
		}
	}
	log.Println("Listern at: ", listenAddr)
	err := http.ListenAndServe(listenAddr, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func getconfig() *config {
	file, err := os.Open("config.json")
	if err != nil {
		log.Println("Config file is not exist")
		log.Println("Default Setting will be actived")
		return nil
	}
	defer file.Close()

	buf := make([]byte, 1024)
	conf := make([]byte, 0)

	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)

		}
		if 0 == n {
			break

		}
		conf = append(conf, buf[:n]...)

	}

	var cfg config
	err = json.Unmarshal(conf, &cfg)
	if err != nil {
		log.Println(err)
		return nil

	}
	return &cfg
}
