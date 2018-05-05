package main

import (
	"github.com/go-martini/martini"
	"github.com/gorilla/websocket"
	"html/template"
	"net/http"
	"./game"
)

const addr = "127.0.0.1:8080"

func httpHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Host string
		MaxWidth int
		MaxHeight int
	}{addr, game.MAX_WIDTH, game.MAX_HEIGHT}

	mainTemplate, err := template.ParseFiles("templates/main.html")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	err = mainTemplate.Execute(w, data)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wsHandler(writer http.ResponseWriter, request *http.Request) {
	conn, err := upgrader.Upgrade(writer, request, nil)

	if err != nil {
		return
	}

	readChannel := make(chan *MessageData, 1024)
	go ReadMessages(readChannel, conn, writer)

	writeChannel := make(chan *GlobalState, 1024)
	go WriteMessages(writeChannel, conn)

	AddController(readChannel, writeChannel)
}

func ReadMessages(c chan *MessageData, conn *websocket.Conn, writer http.ResponseWriter) {

	/* defer func() {
		recover()
		log.Println(123)
		conn.Close()
		c <- nil
	}() */

	for {
		var data map[string]interface{}
		err := conn.ReadJSON(&data)
		if err != nil {
			conn.Close()
			c <- nil
			return
		}

		c <- NewMessageData(data)
	}
}

func WriteMessages(c chan *GlobalState, conn *websocket.Conn) {
	for {
		msg := <-c
		err := conn.WriteJSON(msg.ToJsonMap())
		if err != nil {
			conn.Close()
			return
		}
	}
}

func startServer() {
	m := martini.Classic()
	m.Get("/", httpHandler)
	m.Get("/ws", wsHandler)
	m.Use(martini.Static(""))
	m.RunOnAddr(addr)
}
