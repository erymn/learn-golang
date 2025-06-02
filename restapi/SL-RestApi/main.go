package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Create Slice/array of struct
var profiles []Profile = []Profile{}

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type Profile struct {
	Department  string `json:"department"`
	Designation string `json:"designation"`
	Employee    User   `json:"employee"`
}

func addItem(w http.ResponseWriter, r *http.Request) {
	var profile Profile

	//ambil profile dari request body dengan mekanisme decode
	//json.NewDecoder(r.Body) -> membaca data dari request body
	//Decode(&profile) -> mengubah data dari request body menjadi struct Profile
	_ = json.NewDecoder(r.Body).Decode(&profile)

	//tambahkan header content type application/json untuk memberitahu client bahwa response berupa json
	w.Header().Set("Content-Type", "application/json")

	//tambahkan profile baru ke slice profiles
	profiles = append(profiles, profile)

	//kirim response ke client dengan status code 201 dan data profile yang baru ditambahkan
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}

func getAllProfiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profiles)
}

func getProfile(w http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	if id >= len(profiles) {
		w.WriteHeader(404)
		w.Write([]byte("No profile found with specified ID"))
		return
	}

	profile := profiles[id]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/profiles", addItem).Methods("POST")
	router.HandleFunc("/profiles", getAllProfiles).Methods("GET")
	router.HandleFunc("/profiles/{id}", getProfile).Methods("GET")

	http.ListenAndServe(":8080", router)
}
