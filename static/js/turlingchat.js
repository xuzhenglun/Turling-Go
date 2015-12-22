function turlingReply(m) {
  obj = JSON.parse(m.data);
  switch (obj.code) {
    case 100000:
      //	$('#chatbox').append('<blockquote>');
      $('#chatbox').append('<blockquote><p class=text-success>' + obj.text + '</p></blockquote>');
      //$('#chatbox').append('</blockquote>');
      break;
    case 200000:
      $('#chatbox').append('<blockquote><p class=text-success>' + obj.text + '</p></blockquote>');
      $('#chatbox').append('<a href="' + obj.url + '" target="_blank" style="margin-left:30px;">点击查看</a>');
      break;
      //Picture;
    case 302000:
      $('#chatbox').append('<blockquote> <p class=text-success>' + obj.text + '</p>');
      var info = ''
      for (var key in obj.list) {
        info += '<a href="' + obj.list[key].detailurl + ' "target="_blank"> <p class=text-warning>' + obj.list[key].source + ': ' + obj.list[key].article + '</p></a><br />'
      }
      $('#chatbox').append('<div style="margin-left:30px;><small><a href=""></a>' + info + '</small></div>');
      break;
      //news
    case 305000:
      $('#chatbox').append('<blockquote><p class=text-success>' + obj.text + '</p>');
      if (obj.list != undefined) {
        var info = ''
        for (var key in obj.list) {
          info += '<div><a href="' + obj.list[key].detailurl + '" target="_blank" > <p class=text-warning>' + obj.list[key].trainnum + '</p></a><br />';
          info += '<p class=text-success>' + obj.list[key].start + '发车时间：' + obj.list[key].starttime + '---> ' + obj.list[key].terminal + '结束时间： ' + obj.list[key].endtime + '</p></div><br />';
        }
        $('#chatbox').append('<div style="margin-left:30px;><blockquote><a href=""></a>' + info + '</blockquote></div>');
      }
      break;
      //trains
    case 306000:
      $('#chatbox').append('<blockquote><p class=text-success>' + obj.text + '</p>');
      if (obj.list != undefined) {
        var info = ''
        for (var key in obj.list) {
          info += '<p class=text-warning>' + obj.list[key].flight + ': ' + '发车时间：' + obj.list[key].starttime + '---> ' + '结束时间： ' + obj.list[key].endtime + '</p><br />';
        }
        $('#chatbox').append('<div style="margin-left:30px;><blockquote><a href=""></a>' + info + '</blockquote></div>');
      }
      break;
      //plains
    case 308000:
      $('#chatbox').append('<blockquote><p class=text-success>' + obj.text + '</p>');
      if (obj.list != undefined) {
        var info = ''
        for (var key in obj.list) {
          info += '<a href="' + obj.list[key].detailurl + '" target="_blank"> ' + obj.list[key].name + '</a><br />';
          info += '<p class=text-warning>' + obj.list[key].info + '</p><br />';
        }
        $('#chatbox').append('<div style="margin-left:30px;><blockquote><a href=""></a>' + info + '</blockquote></div>');
      }
      break;
      //menu
    default:
      info = '<p class=text-danger> Error:' + obj.code + '</p><br /><p class=text-danger>' + obj.text + '</p>'
      $('#chatbox').append('<div><blockquote>' + info + '</blockquote></div>');
      //error
  }
  $('#chatbox').scrollTop($('#chatbox')[0].scrollHeight);
}

function init() {
  id = uuid()
  getLocation()
}

function uuid() {
  var s = [];
  var hexDigits = "0123456789abcdef";
  for (var i = 0; i < 16; i++) {
    s[i] = hexDigits.substr(Math.floor(Math.random() * 0x10), 1);
  }
  var uuid = s.join("");
  return uuid;
}

function getLocation() {
  if (navigator.geolocation) {
    navigator.geolocation.getCurrentPosition(showLocation);
  } else {
    console.log("Geolocation is not supported by this browser.");
  }
}

function showLocation(position) {
  geo = position.coords
}
