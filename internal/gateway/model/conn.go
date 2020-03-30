package model

import "github.com/gorilla/websocket"

type Conn struct {
	Uid    uint64
	Wsconn *websocket.Conn
}
