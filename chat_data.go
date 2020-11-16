package main

import "net"

type ChatData struct {
	UserName   string   `json:"user_name"`
	Connection net.Conn `json:"-"`
}

const (
	MESSAGE   = "MESSAGE"
	FILE      = "FILE"
	SEND      = "SEND"
	REGISTER  = "REGISTER"
	DISCONECT = "DISCONECT"
)
