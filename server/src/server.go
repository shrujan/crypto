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
	CoinName         string `json:"coinName"`
	PurchaseDate     string `json:"purchaseDate"`
	TransactionPrice string `json:"transactionPrice"`
	Quantity         string `json:"quantity"`
	TotalAmount      string `json:"totalAmount"`
	UserName         string `json:"userName"`
	BuySell          string `json:"buySell"`
}

type FavCoin struct {
	IsFav  bool   `json:"isFav"`
	Symbol string `json:"symbol"`
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

type UserCoin struct {
	IsFav  bool
	Name   string
	Symbol string
}

// DB is a global variable to hold db connection
var DB *sql.DB

// router
var router *mux.Router

func getAllPurchases(w http.ResponseWriter, r *http.Request) {
	fmt.Println("=============reaches here =================")

	w.Header().Set("Content-type", "application/json")
	// avoid cors error
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	name := mux.Vars(r)["name"]
	var purchases []CryptoInfo

	fmt.Println("============= name  =================" + name)

	query := `select user_name, coin_name, quantity, purchase_price, total_amount from Purchases where user_name = $1`

	rows, err := DB.Query(query, name)
	checkErr(err)

	for rows.Next() {
		var userName, coinName, quantity, transactionPrice, totalAmount string

		err := rows.Scan(&userName, &coinName, &quantity, &transactionPrice, &totalAmount)

		fmt.Println(userName)
		checkErr(err)

		purchase := CryptoInfo{
			UserName:         userName,
			CoinName:         coinName,
			Quantity:         quantity,
			TransactionPrice: transactionPrice,
			TotalAmount:      totalAmount,
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

func getCoinList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	sql := `select name, symbol, isfav from Coins`
	coinList, err := DB.Query(sql)
	if err != nil {
		log.Fatal(err)
	}

	if coinList == nil {
		fmt.Fprintf(w, string("No Coins found"))
	}

	var coins []*UserCoin // declare a slice of courses that will hold all of the Course instances scanned from the rows object
	for coinList.Next() { // this stops when there are no more coin
		coinObj := new(UserCoin)
		fmt.Println(coinObj)                                                 // initialize a new instance
		err := coinList.Scan(&coinObj.Name, &coinObj.Symbol, &coinObj.IsFav) // scan contents of the current row into the instance
		if err != nil {
			log.Fatal(err)
		}

		coins = append(coins, coinObj)
	}

	jsonResp, err := json.Marshal(coins)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

func getWazirxData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	url := "https://api.wazirx.com/api/v2/tickers"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, string(body))
}

func updateFavoriteCoin(w http.ResponseWriter, r *http.Request) {

	// w.Header().Set("Content-type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// fmt.Fprintf(w, "Success")
	// json.NewEncoder(w).Encode(response)

	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// fmt.Println("favorite coins ")
	// var response = JsonResponse{}

	// decoder := json.NewDecoder(r.Body)

	// var coin UserCoin
	// err := decoder.Decode(&coin)
	// fmt.Println("symbol ", coin.Symbol)
	// fmt.Println("fav ", coin.IsFav)

	// checkErr(err)

	// if coin.Symbol == "" {
	// 	response = JsonResponse{Type: "Fail", Message: "Please pass all details"}
	// 	w.WriteHeader(http.StatusPreconditionFailed)
	// } else {

	// 	query := "update coins set isFav=$1 where symbol=$2"
	// 	_, err := DB.Exec(query, coin.IsFav, coin.Symbol)
	// 	checkErr(err)

	// 	response = JsonResponse{Type: "success", Message: "Updated Favorite"}
	// }

	// w.Header().Set("Content-type", "application/json")
	// json.NewEncoder(w).Encode(response)
}

func handleRequests() {
	router.HandleFunc("/getPurchaseInfo/{name}", getAllPurchases).Methods("GET")
	router.HandleFunc("/saveCoin", saveCoinList).Methods("POST")
	router.HandleFunc("/getCoin", getCoinList).Methods("GET")
	router.HandleFunc("/getMarketInfoWX", getWazirxData).Methods("GET")
	// router.HandleFunc("/favoriteCoin", updateFavoriteCoin).Methods("POST")

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

	// CORS Handle
	// corsObj := handlers.AllowedOrigins([]string{"*"})
	//

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
		transactionPrice := cryptInfo.TransactionPrice
		purchaseDate := cryptInfo.PurchaseDate
		totalAmount := cryptInfo.TotalAmount
		buySell := cryptInfo.BuySell
		fmt.Println("==============================")

		if userName == "" || coinName == "" || quantity == "" || transactionPrice == "" || purchaseDate == "" || totalAmount == "" || buySell == "" {
			response = JsonResponse{Type: "error", Message: "You are missing important parameter."}
		} else {
			fmt.Println("user: " + userName + " purchased : " + quantity + coinName + " for " + transactionPrice + " on " + purchaseDate)

			var lastInsertID int
			err := db.QueryRow("INSERT INTO purchases(user_name, coin_name, quantity, purchase_price, date, total_amount, buy_sell) VALUES($1, $2, $3, $4, $5, $6, $7) returning id;", userName, coinName, quantity, transactionPrice, purchaseDate, totalAmount, buySell).Scan(&lastInsertID)
			checkErr(err)

			if err != nil {
				panic(err)
			}
			response = JsonResponse{Type: "success", Message: "Crypto info inserted!"}
		}

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	})

	router.HandleFunc("/favoriteCoin", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		var response = JsonResponse{}
		decoder := json.NewDecoder(r.Body)

		var coin FavCoin
		err := decoder.Decode(&coin)
		fmt.Println("symbol ", coin.Symbol)
		fmt.Println("fav ", coin.IsFav)
		fmt.Println(err)

		query := "update coins set isFav=$1 where symbol=$2"
		_, dbErr := DB.Exec(query, coin.IsFav, coin.Symbol)
		checkErr(dbErr)

		response = JsonResponse{Type: "success", Message: "Updated Favorite"}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Success")

		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	handleRequests()

	log.Fatal(http.ListenAndServe(":8081", router))
	// log.Fatal(http.ListenAndServe(":3000", handlers.CORS(corsObj)(router)))

}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
