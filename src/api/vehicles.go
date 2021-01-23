package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

type Vehicle struct {
	Make     string
	Model    string
	Type     string
	Price    string
	Currency string
	Image    string
}

func RegisterVehicleHandlers() {
	InitVehicles()
	http.HandleFunc("/api/vehicles/list", ListVehicles)
	http.HandleFunc("/api/vehicles/add", CreateVehicle)
}

func InitVehicles() {
	database, err := sql.Open("sqlite3", "vehicles.db")
	if err != nil {
		log.Println(err)
	}
	database.SetMaxOpenConns(1)
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS vehicles (id INTEGER PRIMARY KEY, make TEXT, model TEXT, bodyType TEXT, price TEXT, currency TEXT, image TEXT)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO vehicles (id, make, model, bodyType, price, currency, image) VALUES (?, ?, ?, ?, ?, ?, ?)")
	statement.Exec(1, "Tesla", "Model S", "Sedan", "100600", "CAD", "")
	statement.Exec(2, "Tesla", "Model 3", "Sedan", "51600", "CAD", "")
	statement.Exec(3, "Tesla", "Model X", "SUV", "111600", "CAD", "")
	statement.Exec(4, "Tesla", "Model Y", "SUV", "54900", "CAD", "")
	database.Close()
}

func ListVehicles(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var jsonData []byte
		jsonData, err := json.Marshal(GetVehiclesFromDb())
		if err != nil {
			log.Println(err)
		}
		fmt.Fprint(w, string(jsonData))
	default:
		fmt.Fprintf(w, "Sorry, only GET methods are supported.")
	}
}

func GetVehiclesFromDb() []Vehicle {
	database, err := sql.Open("sqlite3", "vehicles.db")
	if err != nil {
		log.Println(err)
	}

	var vehicles []Vehicle

	rows, _ := database.Query("SELECT id, make, model, bodyType, price, currency, image FROM vehicles")
	var id int
	var make string
	var model string
	var bodyType string
	var price string
	var currency string
	var image string
	for rows.Next() {
		rows.Scan(&id, &make, &model, &bodyType, &price, &currency, &image)
		v:= Vehicle {
			make,
			model,
			bodyType,
			price,
			currency,
			image,
		}
		vehicles = append(vehicles, v)
	}

	database.Close()
	return vehicles
}

// todo: secure endpoint
func CreateVehicle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var vehicle Vehicle
		err := json.NewDecoder(r.Body).Decode(&vehicle)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		database, err := sql.Open("sqlite3", "vehicles.db")
		if err != nil {
			log.Println(err)
		}
		statement, _ := database.Prepare("INSERT INTO vehicles (id, make, model, bodyType, price, currency, image) VALUES (?, ?, ?, ?, ?, ?, ?)")
		statement.Exec(5, vehicle.Make, vehicle.Model, vehicle.Type, vehicle.Price, vehicle.Currency, vehicle.Image)
		database.Close()
		jsonData, err := json.Marshal(vehicle)
		fmt.Fprintf(w, string(jsonData))
	default:
		fmt.Fprintf(w, "Sorry, only POST methods are supported.")
	}
}
