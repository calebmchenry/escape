package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"webl-fun/pkg/engine/game"

	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Server struct {
	router *mux.Router
	game   *game.Game
}

type ResponsePayload struct {
	Data   *Success `json:"data"`
	Errors []string `json:"errors"`
}

type Success struct {
	CharacterID string `json:"characterId"`
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	p := &ResponsePayload{
		Data: &Success{
			CharacterID: uuid.New().String(),
		},
		Errors: []string{},
	}
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	_, err = w.Write(b)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}

func New() *Server {
	s := &Server{}
	s.router = mux.NewRouter()
	s.router.HandleFunc("/login", HandleLogin)
	s.router.HandleFunc("/ws/{id}", s.HandleWS)
	s.game = game.New()
	s.game.Start()
	return s
}

func (s *Server) Serve() {
	port := ":8080"
	fmt.Printf("Serving on %s\n", port)
	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"GET"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:5173"})
	r := handlers.CORS(credentials, methods, origins)(s.router)
	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatalln("There's an error with the server,", err)
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin:     func(r *http.Request) bool { return true },
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Message struct {
	CharacterID string `json:"characterId"`
	ActionID    string `json:"actionId"`
}

func (s *Server) HandleWS(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cID := vars["id"]
	if cID == "" {
		fmt.Printf("Expected a character ID but received an empty string")
		return
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	src := s.game.Subscribe(cID)
	go func() {
		for {
			messageType, b, err := conn.ReadMessage()
			if messageType < 0 {
				// TODO(calebmenry): remove character
				s.game.Unsubscribe(cID)
				break
			}
			if err != nil {
				fmt.Printf("%s\n", err)
			}
			if err != nil {
				fmt.Printf("%s\n", err)
			}
			var msg Message
			err = json.Unmarshal(b, &msg)
			if err != nil {
				fmt.Printf("Failed to parse message: %s \n with error: %s\n", string(b), err)
			}
			s.game.Act(msg.CharacterID, msg.ActionID)
		}
	}()
	go func() {
		for {
			d, ok := <-src
			if !ok {
				return
			}
			b, err := json.Marshal(d)
			if err != nil {
				fmt.Printf("%s\n", err)
			}
			conn.WriteMessage(1, b)
		}
	}()
}
