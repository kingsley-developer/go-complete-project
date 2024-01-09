package soctetiodemo

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

type Chat struct {
	Name string
}
type Chat1 struct {
	Name string
}

func (c *Chat) Serve() {
	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})

	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})

	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})

	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		// server.Remove(s.ID())
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		// Add the Remove session id. Fixed the connection & mem leak
		server.Remove(s.ID())
		fmt.Println("closed", reason)
	})

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func (c *Chat1) Serve1(url string) {
	uri := url

	client, _ := socketio.NewClient(uri, nil)

	// Handle an incoming event
	client.OnEvent("reply", func(s socketio.Conn, msg string) {
		log.Println("Receive Message /reply: ", "reply", msg)
	})

	client.Connect()
	client.Emit("notice", "hello")
	client.Close()
}
