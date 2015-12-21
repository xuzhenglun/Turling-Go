package Turling

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type TurlingRobot struct {
	Key    string
	Url    *url.URL
	values url.Values
}

type Ask struct {
	Info   string
	UserId string
	loc    string
	lat    string
}

func New(address, key string) *TurlingRobot {
	var t TurlingRobot
	var err error
	t.Url, err = url.Parse(address)
	if err != nil {
		log.Println(err)
		return nil
	}
	t.values = t.Url.Query()
	t.values.Set("key", key)
	return &t
}

func (this TurlingRobot) Reply(msg *Ask) string {
	if msg.Info == "" {
		log.Println("Empty Text")
		return ""
	}
	this.values.Set("info", msg.Info)
	if msg.UserId != "" {
		this.values.Set("userid", msg.UserId)
	}
	if msg.lat != "" && msg.loc != "" {
		this.values.Set("lat", msg.lat)
		this.values.Set("loc", msg.loc)
	}

	this.Url.RawQuery = this.values.Encode()
	log.Println(this.Url)

	resp, err := http.Get(this.Url.String())
	if err != nil {
		log.Println(err)
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return ""
	}

	return string(body)
}
