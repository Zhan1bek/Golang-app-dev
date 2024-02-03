package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
	"tsis/package/player"
)

var players = []player.Player{
	{ID: "1", FirstName: "Lionel", LastName: "Messi", Age: 34, Position: "Forward", Nation: "Argentina", ShirtNumber: 10, PlayPeriod: "2004-2021"},
	{ID: "2", FirstName: "Xavi", LastName: "Hernandez", Age: 40, Position: "Midfielder", Nation: "Spain", ShirtNumber: 6, PlayPeriod: "1998-2015"},
	{ID: "3", FirstName: "Andres", LastName: "Iniesta", Age: 36, Position: "Midfielder", Nation: "Spain", ShirtNumber: 8, PlayPeriod: "2002-2018"},
	{ID: "4", FirstName: "Gerard", LastName: "Pique", Age: 34, Position: "Defender", Nation: "Spain", ShirtNumber: 3, PlayPeriod: "2008-2022"},
	{ID: "5", FirstName: "Sergio", LastName: "Busquets", Age: 33, Position: "Midfielder", Nation: "Spain", ShirtNumber: 5, PlayPeriod: "2008-2021"},
	{ID: "6", FirstName: "Carles", LastName: "Puyol", Age: 44, Position: "Defender", Nation: "Spain", ShirtNumber: 5, PlayPeriod: "1999-2014"},
	{ID: "7", FirstName: "Ronaldinho", LastName: "Gaúcho", Age: 41, Position: "Forward", Nation: "Brazil", ShirtNumber: 10, PlayPeriod: "2003-2011"},
	{ID: "8", FirstName: "Samuel", LastName: "Eto'o", Age: 40, Position: "Forward", Nation: "Cameroon", ShirtNumber: 9, PlayPeriod: "1997-2019"},
	{ID: "9", FirstName: "Rivaldo", Age: 49, Position: "Forward", Nation: "Brazil", ShirtNumber: 11, PlayPeriod: "1991-2015"},
	{ID: "10", FirstName: "Marc-André", LastName: "Ter Stegen", Age: 29, Position: "Goalkeeper", Nation: "Germany", ShirtNumber: 1, PlayPeriod: "2012-2021"},
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

func GetPlayerByFirstName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range players {
		if strings.EqualFold(item.FirstName, params["name"]) {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&player.Player{})
}

func GetPlayerByPosition(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var playersOnPosition []player.Player // Создаем временный слайс для игроков на заданной позиции

	for _, item := range players {
		if strings.EqualFold(item.Position, params["position"]) { // Используем регистронезависимое сравнение
			playersOnPosition = append(playersOnPosition, item) // Добавляем игрока в слайс
		}
	}

	if len(playersOnPosition) > 0 {
		json.NewEncoder(w).Encode(playersOnPosition) // Возвращаем всех найденных игроков
	} else {
		w.WriteHeader(http.StatusNotFound) // Устанавливаем статус 404, если игроки не найдены
		json.NewEncoder(w).Encode("No players found on the given position")
	}
}

// Таже самая операция как и в поиске по позиции
func GetPlayerByNation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var PlayerByNation []player.Player
	for _, item := range players {
		if strings.EqualFold(item.Nation, params["nation"]) {
			PlayerByNation = append(PlayerByNation, item)
		}
	}
	if len(PlayerByNation) > 0 {
		json.NewEncoder(w).Encode(PlayerByNation)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("No players found on this nation")
	}
}

func GetPlayersByYear(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	year, err := strconv.Atoi(params["year"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid year format")
		return
	}

	var playersInYear []player.Player
	for _, item := range players {
		periods := strings.Split(item.PlayPeriod, "-")
		if len(periods) == 2 {
			startYear, errStart := strconv.Atoi(periods[0])
			endYear, errEnd := strconv.Atoi(periods[1])
			if errStart == nil && errEnd == nil && year >= startYear && year <= endYear {
				playersInYear = append(playersInYear, item)
			}
		}
	}

	if len(playersInYear) > 0 {
		json.NewEncoder(w).Encode(playersInYear)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("No players found for the given year")
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("API is up and running - developed by AI"))
}
