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
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Lat         float32 `json:"lat"`
	Long        float32 `json:"long"`
	Rating      float32 `json:"rating"`
	Address     string  `json:"address"`
	City        string  `json:"city"`
	Tagline     string  `json:"tagline"`
	Img         string  `json:"img"`
}

type Trivia struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	Candi_id int    `json:"candi_id"`
	Trivia   string `json:"trivia"`
	Img      string `json:"img"`
}

type Article struct {
	ID          int         `json:"id"`
	Category    string      `json:"category"`
	Writer      string      `json:"writer"`
	Description string      `json:"description"`
	Img         string      `json:"img"`
	Date_post   interface{} `json:"date_post"`
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

	// db.AutoMigrate(&Candi{})

	handleRequest()

}

func handleRequest() {

	port := ":5000"

	log.Println("Start Development server on port", port)

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", welcome)

	// Trivia Routes

	myRouter.HandleFunc("/api/trivias", getTrivias).Methods("GET", "OPTIONS")
	myRouter.HandleFunc("/api/trivia/create", createTrivia).Methods("POST", "OPTIONS")
	myRouter.HandleFunc("/api/trivia/{slug}", getTrivia).Methods("GET", "OPTIONS")
	myRouter.HandleFunc("/api/trivia/{id}/update", updateTrivia).Methods("PUT", "OPTIONS")
	myRouter.HandleFunc("/api/trivia/{id}/delete", deleteTrivia).Methods("DELETE", "OPTIONS")

	// Candi Routes
	myRouter.HandleFunc("/api/candis", getCandis).Methods("GET", "OPTIONS")
	myRouter.HandleFunc("/api/candi/create", createCandi).Methods("POST", "OPTIONS")
	myRouter.HandleFunc("/api/candi/{id}", getCandi).Methods("GET", "OPTIONS")
	myRouter.HandleFunc("/api/candi/{id}/update", updateCandi).Methods("PUT", "OPTIONS")
	myRouter.HandleFunc("/api/candi/{id}/delete", deleteCandi).Methods("DELETE", "OPTIONS")

	// Article Routes
	myRouter.HandleFunc("/api/articles", getArticles).Methods("GET", "OPTIONS")
	myRouter.HandleFunc("/api/article/create", createArticle).Methods("POST", "OPTIONS")
	myRouter.HandleFunc("/api/article/{category}", getArticleByCategory).Methods("GET", "OPTIONS")
	myRouter.HandleFunc("/api/article/{id}/update", updateArticle).Methods("PUT", "OPTIONS")
	myRouter.HandleFunc("/api/article/{id}/delete", deleteArticle).Methods("DELETE", "OPTIONS")

	log.Fatal(http.ListenAndServe(port, myRouter))

}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Hexcap REST API")
}

func createTrivia(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var trivia Trivia
	json.Unmarshal(payloads, &trivia)

	db.Create(&trivia)

	res := Result{Code: 200, Data: trivia, Message: "Data Created"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// semua method diperbolehkan masuk
	w.Header().Set("Access-Control-Allow-Methods", "*")

	// semua header diperbolehkan untuk disisipkan
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func getTrivias(w http.ResponseWriter, r *http.Request) {
	trivias := []Trivia{}

	// db.Where("slug = ?", relicSlug).First(&candi)

	// if err := db.Find(&candis).Error; err != nil {
	// 	http.Error(w, err.Error(), http.StatusNotFound)
	// 	return
	// }

	db.Find(&trivias)

	res := Result{Code: 200, Data: trivias, Message: "Data Received"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// semua method diperbolehkan masuk
	w.Header().Set("Access-Control-Allow-Methods", "*")

	// semua header diperbolehkan untuk disisipkan
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func getTrivia(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	relicSlug := vars["slug"]

	log.Println(" isi slug ", relicSlug)

	var trivia Trivia

	// db.Where("slug = ?", relicSlug).First(&candi)

	if err := db.Where("slug = ?", relicSlug).First(&trivia).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	res := Result{Code: 200, Data: trivia, Message: "Data Received"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// semua method diperbolehkan masuk
	w.Header().Set("Access-Control-Allow-Methods", "*")

	// semua header diperbolehkan untuk disisipkan
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func updateTrivia(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	triviaId := vars["id"]

	log.Println(" ID ", triviaId)

	payloads, _ := ioutil.ReadAll(r.Body)

	var candiUpdates Trivia
	json.Unmarshal(payloads, &candiUpdates)
	// db.Where("slug = ?", relicSlug).First(&candi)

	var trivia Trivia
	if err := db.First(&trivia, triviaId).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	db.Model(&trivia).Update(candiUpdates)

	res := Result{Code: 200, Data: trivia, Message: "Data Updated"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// semua method diperbolehkan masuk
	w.Header().Set("Access-Control-Allow-Methods", "*")

	// semua header diperbolehkan untuk disisipkan
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func deleteTrivia(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	triviaId := vars["id"]

	log.Println(" ID ", triviaId)

	var trivia Trivia
	if err := db.First(&trivia, triviaId).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	db.Delete(&trivia)

	res := Result{Code: 200, Message: "Data Deleted"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// semua method diperbolehkan masuk
	w.Header().Set("Access-Control-Allow-Methods", "*")

	// semua header diperbolehkan untuk disisipkan
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

// CANDI FUNCTIONS

func createCandi(w http.ResponseWriter, r *http.Request) {
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

	// semua method diperbolehkan masuk
	w.Header().Set("Access-Control-Allow-Methods", "*")

	// semua header diperbolehkan untuk disisipkan
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func getCandis(w http.ResponseWriter, r *http.Request) {
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

	// semua method diperbolehkan masuk
	w.Header().Set("Access-Control-Allow-Methods", "*")

	// semua header diperbolehkan untuk disisipkan
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func updateCandi(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	candiId := vars["id"]

	log.Println(" ID ", candiId)

	payloads, _ := ioutil.ReadAll(r.Body)

	var candiUpdates Trivia
	json.Unmarshal(payloads, &candiUpdates)
	// db.Where("slug = ?", relicSlug).First(&candi)

	var candi Candi
	if err := db.First(&candi, candiId).Error; err != nil {
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

	// semua method diperbolehkan masuk
	w.Header().Set("Access-Control-Allow-Methods", "*")

	// semua header diperbolehkan untuk disisipkan
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func deleteCandi(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	candiId := vars["id"]

	log.Println(" ID ", candiId)

	var candi Candi
	if err := db.First(&candi, candiId).Error; err != nil {
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

	// semua method diperbolehkan masuk
	w.Header().Set("Access-Control-Allow-Methods", "*")

	// semua header diperbolehkan untuk disisipkan
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func getCandi(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	candiID := vars["id"]

	log.Println(" isi id ", candiID)

	var candi Candi

	// db.Where("slug = ?", relicSlug).First(&candi)

	if err := db.First(&candi, candiID).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	res := Result{Code: 200, Data: candi, Message: "Data Received"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// semua method diperbolehkan masuk
	w.Header().Set("Access-Control-Allow-Methods", "*")

	// semua header diperbolehkan untuk disisipkan
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

// fungsi untuk article
func createArticle(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var article Article
	json.Unmarshal(payloads, &article)

	db.Create(&article)

	res := Result{Code: 200, Data: article, Message: "Article Data Created"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// semua method diperbolehkan masuk
	w.Header().Set("Access-Control-Allow-Methods", "*")

	// semua header diperbolehkan untuk disisipkan
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func getArticles(w http.ResponseWriter, r *http.Request) {
	articles := []Article{}

	// db.Where("slug = ?", relicSlug).First(&candi)

	// if err := db.Find(&candis).Error; err != nil {
	// 	http.Error(w, err.Error(), http.StatusNotFound)
	// 	return
	// }

	db.Find(&articles)

	res := Result{Code: 200, Data: articles, Message: "Articles Data Received"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// semua method diperbolehkan masuk
	w.Header().Set("Access-Control-Allow-Methods", "*")

	// semua header diperbolehkan untuk disisipkan
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleId := vars["id"]

	log.Println(" ID ", articleId)

	payloads, _ := ioutil.ReadAll(r.Body)

	var articleUpdates Article
	json.Unmarshal(payloads, &articleUpdates)
	// db.Where("slug = ?", relicSlug).First(&candi)

	var article Article
	if err := db.First(&article, articleId).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	db.Model(&article).Update(articleUpdates)

	res := Result{Code: 200, Data: article, Message: "Article Data Updated"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// semua method diperbolehkan masuk
	w.Header().Set("Access-Control-Allow-Methods", "*")

	// semua header diperbolehkan untuk disisipkan
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleId := vars["id"]

	log.Println(" ID ", articleId)

	var article Article
	if err := db.First(&article, articleId).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	db.Delete(&article)

	res := Result{Code: 200, Message: "Article Data Deleted"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// semua method diperbolehkan masuk
	w.Header().Set("Access-Control-Allow-Methods", "*")

	// semua header diperbolehkan untuk disisipkan
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func getArticleByCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleCategory := vars["category"]

	// log.Println(" isi slug ", relicSlug)

	var article []Article

	// db.Where("slug = ?", relicSlug).First(&candi)

	if err := db.Where("category = ?", articleCategory).Find(&article).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	res := Result{Code: 200, Data: article, Message: "Article Category Received"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// semua method diperbolehkan masuk
	w.Header().Set("Access-Control-Allow-Methods", "*")

	// semua header diperbolehkan untuk disisipkan
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
