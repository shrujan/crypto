package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"

	_ "github.com/lib/pq"
)

const (
	hostname      = "localhost"
	host_port     = 5432
	username      = "postgres"
	password      = "Shrujan@123"
	database_name = "Crypto"
)

type User struct {
	UserName string
	Email    string
}

type JsonResponse struct {
	Type    string `json:"type"`
	Data    []User `json:"data"`
	Message string `json:"message"`
}

type CryptoInfo struct {
	UserName      string `json:"userName"`
	CoinName      string `json:"coinName"`
	Quantity      string `json:"quantity"`
	PurchasePrice string `json:"purchasePrice"`
	PurchaseDate  string `json:"purchaseDate"`
	TotalAmount   string `json:"totalAmount"`
}

func main() {
	// DB related connection
	connection := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		hostname, host_port, username, password, database_name)

	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	// DB connection here closes

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

	http.HandleFunc("/getUsers", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		sql := `select * from "Users"`
		usersList, err := db.Query(sql)
		if err != nil {
			log.Fatal(err)
		}

		if usersList == nil {
			fmt.Fprintf(w, string("No Users found"))
		}

		var users []*User      // declare a slice of courses that will hold all of the Course instances scanned from the rows object
		for usersList.Next() { // this stops when there are no more usersList
			userObj := new(User)                                     // initialize a new instance
			err := usersList.Scan(&userObj.UserName, &userObj.Email) // scan contents of the current row into the instance
			if err != nil {
				log.Fatal(err)
			}

			users = append(users, userObj)
		}

		jsonResp, err := json.Marshal(users)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResp)
	})

	http.HandleFunc("/savePurchaseInfo", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method != http.MethodPost {
			fmt.Fprintf(w, "Post method only")
			return
		}

		var response = JsonResponse{}

		decoder := json.NewDecoder(r.Body)

		var cryptInfo CryptoInfo
		err := decoder.Decode(&cryptInfo)

		checkErr(err)

		fmt.Println("==============================")
		userName := cryptInfo.UserName
		coinName := cryptInfo.CoinName
		quantity := cryptInfo.Quantity
		purchasePrice := cryptInfo.PurchasePrice
		purchaseDate := cryptInfo.PurchaseDate
		totalAmount := cryptInfo.TotalAmount
		fmt.Println("==============================")

		if userName == "" || coinName == "" || quantity == "" || purchasePrice == "" || purchaseDate == "" || totalAmount == "" {
			response = JsonResponse{Type: "error", Message: "You are missing important parameter."}
		} else {
			fmt.Println("user: " + userName + " purchased : " + quantity + coinName + " for " + purchasePrice + " on " + purchaseDate)

			var lastInsertID int
			err := db.QueryRow("INSERT INTO purchases(user_name, coin_name, quantity, purchase_price, date, total_amount) VALUES($1, $2, $3, $4, $5, $6) returning id;", userName, coinName, quantity, purchasePrice, purchaseDate, totalAmount).Scan(&lastInsertID)
			checkErr(err)

			if err != nil {
				panic(err)
			}
			response = JsonResponse{Type: "success", Message: "Crypto info inserted!"}
		}

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		// fmt.Fprintf(w, string("Success"))
		json.NewEncoder(w).Encode(response)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
