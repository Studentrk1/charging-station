package main

import (
	"encoding/json"
	"time"

	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type chargingstation struct {
	StationID    int
	EnergyOutput string
	Type         string
}

type OccuChargingStation struct {
	StationID        string
	EnergyOutput     string
	Type             string
	AvailabilityTime string
}

var chargestation []chargingstation
var OccupiedChargeStation []OccuChargingStation

func Requesthandler() {
	router := mux.NewRouter()
	router.HandleFunc("/chst", ChargingStation).Methods("POST")
	router.HandleFunc("/allchst", ChargingStation)
	//router.HandleFunc("/AvailableChargingSt", AvChargingst)
	router.HandleFunc("/OccupiedCharge", OccCharge)
	router.HandleFunc("/stch", startCharging).Methods("POST")
	http.ListenAndServe("127.0.0.1:8000", router)

}
func chargingStation(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var chst chargingstation
	json.Unmarshal(reqBody, &chst)
	chargestation = append(chargestation, chst)
	json.NewEncoder(w).Encode(chst)
}
func ChargingStation(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chargestation)
}

func startCharging(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))

}

//func AvChargingst(w http.ResponsWriter, r *http.Request) {
//json.NewEncorder(w).Encode(chargestation)
//}
func OccCharge(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(OccupiedChargeStation)
}

func main() {
	chargestation = []chargingstation{}

	OccupiedChargeStation = []OccuChargingStation{
		{StationID: "1", EnergyOutput: "50kwh", Type: "DC"},
		{"2", "10kwh", "DC", time.Now().UTC().String()},
		{"3", "100kwh", "DC", time.Now().UTC().String()},
	}
	Requesthandler()
}
