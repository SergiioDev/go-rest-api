package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/SergiioDev/learning-go/config"
	"github.com/julienschmidt/httprouter"
)

type Driver struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	LastName  string `json:"lastname"`
	Team      string `json:"team"`
	DateBirth string `json:"date_birth"`
}

func main() {
	conf := config.New()

	router := httprouter.New()
	router.GET("/drivers", Drivers)
	router.POST("/drivers", DriversCreate)
	router.GET("/drivers/:name", DriverByName)
	serverPort := ":%s"
	log.Fatal(http.ListenAndServe(fmt.Sprintf(serverPort, conf.Server.Port), router))

}

var drivers = []Driver{
	{ID: "1", Name: "Fernando", LastName: "Alonso", Team: "Aston Martin", DateBirth: "07/29/1981"},
}

func Drivers(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Println("Handling GET /drivers request")
	w.Header().Set("Content-Type", "application/json")
	res, err := json.Marshal(drivers)
	if err != nil {
		fmt.Println("Error parsing json")
	}
	w.Write(res)
}

func DriversCreate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var driver Driver
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&driver)
	drivers = append(drivers, driver)
	w.WriteHeader(http.StatusCreated)
	res, err := json.Marshal(driver)
	if err != nil {
		fmt.Println("Error parsing driver into a json")
	}
	w.Write(res)
}

func DriverByName(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	driverName := params.ByName("name")
	diver, err := findDriverByName(driverName)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
	res, err := json.Marshal(diver)
	if err != nil {
		http.Error(w, "Can Not parse driver into json", http.StatusNotFound)
	}
	w.Write(res)
}

func findDriverByName(name string) (Driver, error) {
	for i := 0; i < len(drivers); i++ {
		if drivers[i].Name == name {

			fmt.Println(drivers[i].Name)
			return drivers[i], nil
		}
	}
	return Driver{}, errors.New("Can Not find driver with name " + name)
}
