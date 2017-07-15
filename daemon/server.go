package main

import (
	"net"
	"github.com/Sirupsen/logrus"
)

func startServer(host string, port string) {

	ln, err := net.Listen("tcp", host + ":" + port)

	if err != nil {

		panic("Can`t start daemon: " + err.Error())

	} else {

		logrus.Info("Started daemon on: " + host + ":" + port)

	}

	for {

		conn, err := ln.Accept()

		logrus.Info("New command from: " + conn.LocalAddr().String())

		if err != nil {
			logrus.Error("Connection error: " + err.Error())
		}

		go handleCmd(conn)
	}
}

func handleCmd(conn net.Conn) {

}
