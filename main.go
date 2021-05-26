package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

type Candi struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Slug   string `json:"slug"`
	Trivia string `json:"trivia"`
}

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {

	db, err = gorm.Open("mysql", "pakpak:123456@tcp(34.101.119.51)/relic?charset=utf8mb4&parseTime=True")

	if err != nil {
		log.Println("Koneksi gagal !", err)
	} else {
		log.Println("Koneksi Berhasil !")
	}

	db.AutoMigrate(&Candi{})

	handleRequest()

}

func handleRequest() {

	port := ":3000"

	log.Println("Start Development server on port", port)

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", welcome)
	myRouter.HandleFunc("/api/relics", getRelics).Methods("GET", "OPTIONS")
	myRouter.HandleFunc("/api/relic/create", createRelic).Methods("POST", "OPTIONS")
	myRouter.HandleFunc("/api/relic/{slug}", getRelic).Methods("GET", "OPTIONS")
	myRouter.HandleFunc("/api/relic/{id}/update", updateRelic).Methods("PUT", "OPTIONS")
	myRouter.HandleFunc("/api/relic/{id}/delete", deleteRelic).Methods("DELETE", "OPTIONS")

	log.Fatal(http.ListenAndServe(port, myRouter))

}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Hexcap REST API")
}

func createRelic(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var candi Candi
	json.Unmarshal(payloads, &candi)

	db.Create(&candi)

	res := Result{Code: 200, Data: candi, Message: "Data Created"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func getRelics(w http.ResponseWriter, r *http.Request) {
	candis := []Candi{}

	// db.Where("slug = ?", relicSlug).First(&candi)

	// if err := db.Find(&candis).Error; err != nil {
	// 	http.Error(w, err.Error(), http.StatusNotFound)
	// 	return
	// }

	db.Find(&candis)

	res := Result{Code: 200, Data: candis, Message: "Data Received"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func getRelic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	relicSlug := vars["slug"]

	log.Println(" isi slug ", relicSlug)

	var candi Candi
	// db.Where("slug = ?", relicSlug).First(&candi)

	if err := db.Where("slug = ?", relicSlug).First(&candi).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	res := Result{Code: 200, Data: candi, Message: "Data Received"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func updateRelic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	relicId := vars["id"]

	log.Println(" ID ", relicId)

	payloads, _ := ioutil.ReadAll(r.Body)

	var candiUpdates Candi
	json.Unmarshal(payloads, &candiUpdates)
	// db.Where("slug = ?", relicSlug).First(&candi)

	var candi Candi
	if err := db.First(&candi, relicId).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	db.Model(&candi).Update(candiUpdates)

	res := Result{Code: 200, Data: candi, Message: "Data Updated"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func deleteRelic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	relicId := vars["id"]

	log.Println(" ID ", relicId)

	var candi Candi
	if err := db.First(&candi, relicId).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	db.Delete(&candi)

	res := Result{Code: 200, Message: "Data Deleted"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
