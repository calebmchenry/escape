package server

import (
	"encoding/json"
	"fmt"
	"net"
	"webl-fun/pkg/engine/game"
)

type SocketServer struct {
	game *game.Game
}

func NewSocketServer() *SocketServer {
	s := &SocketServer{}
	s.game = game.New()
	s.game.Start()
	return s
}

func (s *SocketServer) Serve() {
	port := ":8080"

	fmt.Printf("Listening on :8080\n")
	ln, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Accept incoming connections and handle them
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("Connection received")
		// Handle the connection in a new goroutine
		go s.HandleConnection(conn)
	}
}

type AuthMessage struct {
	Token string `json:"token"`
}

type SocketMessage struct {
	ActionID string `json:"actionId"`
}

func (s *SocketServer) HandleConnection(conn net.Conn) {
	enc := json.NewEncoder(conn)
	dec := json.NewDecoder(conn)

	var auth AuthMessage
	err := dec.Decode(&auth)
	if err != nil {
		fmt.Printf("Error while decoding auth: %s\n", err)
	}

	fmt.Printf("Adding a new character with pid: %s\n", auth.Token)

	src := s.game.Subscribe(auth.Token)
	go func() {
		for {
			d, ok := <-src
			if !ok {
				return
			}
			err = enc.Encode(d)
			if err != nil {
				fmt.Printf("Error while encoding tick: %s\n", err)
			}
		}
	}()
	defer conn.Close()
	defer s.game.Unsubscribe(auth.Token)
	for {
		var command ClientCommand
		err = dec.Decode(&command)
		if err != nil {
			fmt.Printf("Error while decoding message: %s\n", err)
			return
		}
		s.game.Act(auth.Token, command.Action, command.TargetID, command.SrcID)
	}
}
