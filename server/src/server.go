package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"

	// "github.com/gorilla/websocket"
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

	http.HandleFunc("/sendmail", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			fmt.Fprintf(w, "Post method only")
			return
		}
		// Sender data.
		from := "shrujantestmail@gmail.com"
		password := "yqqpxjuwbjmstrah"
	  
		// Receiver email address.
		to := []string{
		  "shrork@gmail.com",
		}
	  
		// smtp server configuration.
		smtpHost := "smtp.gmail.com"
		smtpPort := "587"
	  
		// Message.
		message := []byte("This is a test email message.")
		
		// Authentication.
		auth := smtp.PlainAuth("", from, password, smtpHost)
		
		// Sending email.
		err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
		if err != nil {
		  fmt.Println(err)
		  return
		}
		fmt.Fprintf(w, "Success")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
