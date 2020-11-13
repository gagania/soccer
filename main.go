package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"test-api/config"
	"test-api/models"
	"test-api/players"
	"test-api/team"
	"test-api/utils"

	"github.com/gorilla/mux"
)

func main() {
	db, e := config.MySQL()

	if e != nil {
		log.Fatal(e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}

	fmt.Println("Success")
	//intial router
	// router := mux.NewRouter()
	router := mux.NewRouter()
	router.HandleFunc("/api/teams", getTeams).Methods("GET")
	router.HandleFunc("/api/teams", createTeam).Methods("POST")
	router.HandleFunc("/api/teams", updateTeam).Methods("PUT")
	router.HandleFunc("/api/teams/{id}", deleteTeam).Methods("DELETE")

	router.HandleFunc("/api/player", getPlayers).Methods("GET")
	router.HandleFunc("/api/player", createPlayer).Methods("POST")
	router.HandleFunc("/api/player", updatePlayer).Methods("PUT")
	router.HandleFunc("/api/player/{id}", deletePlayer).Methods("DELETE")
	fmt.Println("Running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func getTeams(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ctx := r.Context()

		teams, err := team.GetAll(ctx)

		if err != nil {
			fmt.Println(err)
		}

		utils.ResponseJSON(w, teams, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
	return
}

func createTeam(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx := r.Context()

		var t models.Team

		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := team.Insert(ctx, t); err != nil {
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Succesfully",
		}

		utils.ResponseJSON(w, res, http.StatusCreated)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}

// UpdateTeam
func updateTeam(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {

		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx := r.Context()

		var t models.Team

		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		fmt.Println(t)

		if err := team.UpdateTeam(ctx, t); err != nil {
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Succesfully",
		}

		utils.ResponseJSON(w, res, http.StatusCreated)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}

// Delete Team
func deleteTeam(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}
		ctx := r.Context()
		var t models.Team
		params := mux.Vars(r)
		id := params["id"]

		if id == "" {
			utils.ResponseJSON(w, "id tidak boleh kosong", http.StatusBadRequest)
			return
		}
		t.ID, _ = strconv.Atoi(id)

		if err := team.DeleteTeam(ctx, t); err != nil {

			kesalahan := map[string]string{
				"error": fmt.Sprintf("%v", err),
			}

			utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Succesfully",
		}

		utils.ResponseJSON(w, res, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}

func getPlayers(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ctx := r.Context()

		player, err := players.GetAllPlayer(ctx)

		if err != nil {
			fmt.Println(err)
		}

		utils.ResponseJSON(w, player, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
	return
}

func createPlayer(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx := r.Context()

		var p models.Players

		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := players.InsertPlayer(ctx, p); err != nil {
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Succesfully",
		}

		utils.ResponseJSON(w, res, http.StatusCreated)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}

// UpdatePlayer
func updatePlayer(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx := r.Context()

		var p models.Players

		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		fmt.Println(p)

		if err := players.UpdatePlayer(ctx, p); err != nil {
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Succesfully",
		}

		utils.ResponseJSON(w, res, http.StatusCreated)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}

// Delete Player
func deletePlayer(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}
		ctx := r.Context()
		var p models.Players
		params := mux.Vars(r)
		id := params["id"]

		if id == "" {
			utils.ResponseJSON(w, "id tidak boleh kosong", http.StatusBadRequest)
			return
		}
		p.ID, _ = strconv.Atoi(id)

		if err := players.DeletePlayer(ctx, p); err != nil {

			kesalahan := map[string]string{
				"error": fmt.Sprintf("%v", err),
			}

			utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Succesfully",
		}

		utils.ResponseJSON(w, res, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}
