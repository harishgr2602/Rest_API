package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Article struct
type Article struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Genre         string `json:"genre"`
	Author        string `json:"author"`
	PublishedYear string `json:"publishedyear"`
	Publisher     string `json:"publisher"`
}

var articles []Article

func main() {
	r := mux.NewRouter()

	articles = append(articles, Article{ID: "1", Title: "War and Peace", Genre: "Historical Fiction", Author: "Leo Tolstoy", PublishedYear: "1897", Publisher: "The Russian Messanger"})
	articles = append(articles, Article{ID: "2", Title: "The Wind In The Willows", Genre: "Fanatsy", Author: "Kenneth Grahame", PublishedYear: "1908", Publisher: "Methuen"})
	articles = append(articles, Article{ID: "3", Title: "The Dark World", Genre: "Fantasy", Author: "Henry Kuttner", PublishedYear: "1965", Publisher: "Ace Books"})
	articles = append(articles, Article{ID: "4", Title: "The Time Machine", Genre: "Science Fiction", Author: "H.G.Wells", PublishedYear: "1895", Publisher: "William Heinemann"})

	r.HandleFunc("/article", getArticles).Methods("GET")
	r.HandleFunc("/article/{id}", getarticle).Methods("GET")
	r.HandleFunc("/article/{id}", postArticle).Methods("POST")
	r.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	r.HandleFunc("/article/{id}", updateArticle).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}
func getarticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, v := range articles {
		if v.ID == params["id"] {
			json.NewEncoder(w).Encode(v)
			return
		}
	}

	json.NewEncoder(w).Encode(&Article{})
}
func postArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := mux.Vars(r)
	var arr Article
	_ = json.NewDecoder(r.Body).Decode(&arr)
	arr.ID = p["id"]
	articles = append(articles, arr)
	json.NewEncoder(w).Encode(articles)
}
func deleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content/Type", "application/json")
	p := mux.Vars(r)
	for k, v := range articles {
		if v.ID == p["id"] {
			articles = append(articles[:k], articles[(k+1):]...)
			break
		}
	}
	json.NewEncoder(w).Encode(articles)
}
func updateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content/Type", "application/json")
	params := mux.Vars(r)
	for k, v := range articles {
		if v.ID == params["id"] {
			articles = append(articles[:k], articles[k+1:]...)
			var art Article
			_ = json.NewDecoder(r.Body).Decode(&art)
			art.ID = params["id"]
			articles = append(articles, art)
			json.NewEncoder(w).Encode(articles)
			return
		}
	}
	json.NewEncoder(w).Encode(articles)
}
