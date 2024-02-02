package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"tsis/tsis1/package/player"
)

var players = []player.Player{
	{ID: "1", Name: "Lionel Messi", Age: 34, Position: "Forward", Nation: "Argentina", ShirtNumber: 10},
	{ID: "2", Name: "Xavi Hernandez", Age: 40, Position: "Midfielder", Nation: "Spain", ShirtNumber: 6},
	{ID: "3", Name: "Andres Iniesta", Age: 36, Position: "Midfielder", Nation: "Spain", ShirtNumber: 8},
	{ID: "4", Name: "Gerard Pique", Age: 34, Position: "Defender", Nation: "Spain", ShirtNumber: 3},
	{ID: "5", Name: "Sergio Busquets", Age: 33, Position: "Midfielder", Nation: "Spain", ShirtNumber: 5},
	{ID: "6", Name: "Carles Puyol", Age: 44, Position: "Defender", Nation: "Spain", ShirtNumber: 5},
	{ID: "7", Name: "Ronaldinho", Age: 41, Position: "Forward", Nation: "Brazil", ShirtNumber: 10},
	{ID: "8", Name: "Samuel Eto'o", Age: 40, Position: "Forward", Nation: "Cameroon", ShirtNumber: 9},
	{ID: "9", Name: "Rivaldo", Age: 49, Position: "Forward", Nation: "Brazil", ShirtNumber: 11},
	{ID: "10", Name: "Marc-André ter Stegen", Age: 29, Position: "Goalkeeper", Nation: "Germany", ShirtNumber: 1},
}

func GetPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(players)
}

func GetPlayerByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range players {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&player.Player{})
}

func GetPlayerByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range players {
		if item.Name == params["name"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&player.Player{})
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("API is up and running - developed by AI"))
}
