package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var articles = []Article{
	Article{
		ID:          "1",
		Title:       "Learning golang",
		Description: "an introduction to the go ecosystem",
		Content:     "...",
		Author: &User{
			ID:       "1",
			Username: "salekin",
			Email:    "ss@gmail.com",
		},
	},
	Article{
		ID:          "2",
		Title:       "Learning GIN",
		Description: "intro to GIN framework",
		Content:     "...",
		Author: &User{
			ID:       "1",
			Username: "salekin",
			Email:    "ss@gmail.com",
		},
	},
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"user"`
	Email    string `json:"email"`
}

type Article struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Author      *User  `json:"author"`
}

func getArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, article := range articles {
		if article.ID == params["id"] {
			json.NewEncoder(w).Encode(article)
			return
		}
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	article := Article{}
	_ = json.NewDecoder(r.Body).Decode(&article)
	article.ID = strconv.Itoa(len(articles) + 1)
	articles = append(articles, article)

	json.NewEncoder(w).Encode(article)
}

func update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	article := Article{}
	_ = json.NewDecoder(r.Body).Decode(&article)
	article.ID = params["id"]

	for i, article := range articles {
		if article.ID == params["id"] {
			articles[i] = article
			break
		}
	}

	json.NewEncoder(w).Encode(article)
}

func delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for i, article := range articles {
		if article.ID == params["id"] {
			articles = append(articles[:i], articles[:i+1]...)
			break
		}
	}

	json.NewEncoder(w).Encode(articles)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/articles", getArticles).Methods("GET")
	router.HandleFunc("/api/articles/{id}", getArticle).Methods("GET")
	router.HandleFunc("/api/articles/", create).Methods("POST")
	router.HandleFunc("/api/articles/{id}", update).Methods("PUT")
	router.HandleFunc("/api/articles/{id}", delete).Methods("DELETE")

	http.ListenAndServe(":3000", router)
}