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

	"github.com/gorilla/mux"
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

type Coin struct {
	Ath                              float64
	Ath_change_percentage            float64
	Ath_date                         string
	Atl                              float64
	Atl_change_percentage            float64
	Atl_date                         string
	Circulating_supply               float64
	Current_price                    float64
	Fully_diluted_valuation          float64
	High_24h                         float64
	Id                               string
	Image                            string
	Last_updated                     string
	Low_24h                          float64
	Market_cap                       float64
	Market_cap_change_24h            float64
	Market_cap_change_percentage_24h float64
	Market_cap_rank                  float64
	Max_supply                       float64
	Name                             string
	Price_change_24h                 float64
	Price_change_percentage_24h      float64
	Symbol                           string
	Total_supply                     float64
	Total_volume                     float64
}

// DB is a global variable to hold db connection
var DB *sql.DB

// router
var router *mux.Router

func getAllPurchases(w http.ResponseWriter, r *http.Request) {
	fmt.Println("=============reaches here =================")

	w.Header().Set("Content-type", "application/json")

	name := mux.Vars(r)["name"]
	var purchases []CryptoInfo

	fmt.Println("============= name  =================" + name)

	query := `select user_name, coin_name, quantity, purchase_price, total_amount from Purchases where user_name = $1`

	rows, err := DB.Query(query, name)
	checkErr(err)

	for rows.Next() {
		var userName, coinName, quantity, purchasePrice, totalAmount string

		err := rows.Scan(&userName, &coinName, &quantity, &purchasePrice, &totalAmount)

		fmt.Println(userName)
		checkErr(err)

		purchase := CryptoInfo{
			UserName:      userName,
			CoinName:      coinName,
			Quantity:      quantity,
			PurchasePrice: purchasePrice,
			TotalAmount:   totalAmount,
		}

		purchases = append(purchases, purchase)
	}

	defer rows.Close()

	json.NewEncoder(w).Encode(purchases)

}

func saveCoinList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	fmt.Println("Yo .. we reached the save coin section")

	decoder := json.NewDecoder(r.Body)
	var coinData []Coin

	err := decoder.Decode(&coinData)
	checkErr(err)

	fmt.Println("-=============")
	fmt.Println(coinData[0].Name)
	for key, coin := range coinData {
		fmt.Println(key, " -- "+coin.Name)
		// insert query
		query := "insert into coins (id, name, symbol, total_supply, isFav) values ($1, $2, $3, $4, $5)"

		DB.Exec(query, coin.Id, coin.Name, coin.Symbol, coin.Max_supply, false)

	}
}

func handleRequests() {
	router.HandleFunc("/getPurchaseInfo/{name}", getAllPurchases).Methods("GET")
	router.HandleFunc("/saveCoin", saveCoinList).Methods("POST")
}

func main() {

	router = mux.NewRouter()

	// DB related connection
	connection := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		hostname, host_port, username, password, database_name)

	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err)
	}
	DB = db
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	// DB connection here closes

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	router.HandleFunc("/getMarketInfo", func(w http.ResponseWriter, r *http.Request) {
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

	router.HandleFunc("/getMarketInfows", func(w http.ResponseWriter, r *http.Request) {

		// var conn, _ = upgrader.Upgrade(w, r, nil)
		// go func(conn *websocket.Conn) {
		// 	for {
		// 		mType, msg, _ := conn.ReadMessage()

		// 		conn.WriteMessage(mType, msg)
		// 	}
		// }(conn)

	})

	router.HandleFunc("/sendmail", func(w http.ResponseWriter, r *http.Request) {
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

	router.HandleFunc("/getUsers", func(w http.ResponseWriter, r *http.Request) {
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

	router.HandleFunc("/savePurchaseInfo", func(w http.ResponseWriter, r *http.Request) {
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

	handleRequests()

	log.Fatal(http.ListenAndServe(":8081", router))

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
