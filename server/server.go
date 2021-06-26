package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
)

// "github.com/gorilla/websocket"

// C:\Program Files\Go\src\github.com\gorilla\mux

// C:\Users\shror\go\bin

// var upgrader = websocket.Upgrader{}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/getMarketInfo", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		url := "https://api.coingecko.com/api/v3/coins/markets?vs_currency=inr&order=market_cap_desc"
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(w, string(body))
	})

	http.HandleFunc("/getMarketInfows", func(w http.ResponseWriter, r *http.Request) {

		// var conn, _ = upgrader.Upgrade(w, r, nil)
		// go func(conn *websocket.Conn) {
		// 	for {
		// 		mType, msg, _ := conn.ReadMessage()

		// 		conn.WriteMessage(mType, msg)
		// 	}
		// }(conn)

	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
