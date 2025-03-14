package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
)

var db *sql.DB
var config Config = Config{BasePath: "/", Port: 8000}

type Config struct {
	Port          int    `json:"port"`
	BasePath      string `json:"basePath"`
	SecureCookies bool   `json:"secureCookies"`
	DatabaseURL   string `json:"databaseUrl"`
}

// TODO: implement e-mail verification option
// TODO: add forgot password and reset password endpoint
func main() {
	log.SetOutput(os.Stderr)
	configFile, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatalln("Failed to read config file!", err)
	}
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalln("Failed to parse config file!", err)
	}

	dsn, err := mysql.ParseDSN(config.DatabaseURL)
	if err != nil {
		log.Fatalln("Failed to parse MariaDB DSN!", err)
	}
	dsn.MultiStatements = true
	dsn.ParseTime = true
	dsn.Params = map[string]string{"time_zone": "'+00:00'"} // dsn.Loc is already UTC
	config.DatabaseURL = dsn.FormatDSN()

	db, err = sql.Open("mysql", config.DatabaseURL)
	if err != nil {
		log.Fatalln("Failed to open connection to database!", err)
	}
	db.SetMaxOpenConns(10)
	CreateSqlTables()
	PrepareSqlStatements()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" || r.Method != "GET" {
			http.NotFound(w, r)
		} else {
			StatusEndpoint(w, r)
		}
	})
	http.HandleFunc("POST /api/login", LoginEndpoint)
	http.HandleFunc("POST /api/logout", LogoutEndpoint)

	port := strconv.Itoa(config.Port)
	log.SetOutput(os.Stdout)
	log.Println("Listening to port " + port)
	log.SetOutput(os.Stderr)
	log.Fatalln(http.ListenAndServe(":"+port, handlers.CORS(
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PATCH", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowedOrigins([]string{"*"}), // Breaks credentialed auth
		handlers.AllowCredentials(),
	)(http.DefaultServeMux)))
}
