package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "0.0.0.0:1200", "http service address")
var upgrader = websocket.Upgrader{}

type msg struct {
	Jwt  string `json:"jwt"`
	Type string `json:"type"`
	Data string `json:"data"`
}

func sendMsg(j string, t string, d string, c *websocket.Conn) {
	m := msg{j, t, d}
	if err := c.WriteJSON(m); err != nil {
		fmt.Println(err)
	}

	//mm, _ := json.Marshal(m);
	//fmt.Println(string(mm));
}

func handleAPI(w http.ResponseWriter, r *http.Request) {

	//swp := r.Header.Get("Sec-Websocket-Protocol")
	//Wupgrader.CheckOrigin = func(r *http.Request) bool { return true }
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	c, err := upgrader.Upgrade(w, r, nil) //add rh
	if err != nil {
		fmt.Print("@HandleAPI Ws Upgrade Error> ", err)
		return
	}

Loop:
	for {
		in := msg{}

		err := c.ReadJSON(&in)
		if err != nil {
			//fmt.Println("Error reading json.", err)
			c.Close()
			break Loop
		}
		switch in.Type {
		case "test":
			break
		default:
			break
		}
	}
}

func main() {
	//procon_asyncq.StartTaskDispatcher(9)
	//go procon_sse.ConfigureSystemHeartbeat()
	//go procon_sse.StartSSE()

	flag.Parse()
	log.SetFlags(0)

	r := mux.NewRouter()

	//Websocket API
	r.HandleFunc("/api", handleAPI)
	//r.HandleFunc("/pty", procon_wspty.HandleWsPty)

	//Rest API
	//r.HandleFunc("/rest/api/ui/{component}/{subcomponent}", handleUI)

	http.ListenAndServeTLS(*addr, "/etc/letsencrypt/live/server.singularis.work/fullchain.pem", "/etc/letsencrypt/live/server.singularis.work/privkey.pem", r)
}
