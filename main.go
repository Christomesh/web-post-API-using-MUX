package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Users struct
type User struct {
	FullName string `json:"fullName"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Item is a struct that groups all neccessary fields into a single unit
type Post struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author User   `json:"author"`
}

var posts []Post = []Post{}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/posts", addPost).Methods("POST")
	router.HandleFunc("/posts", getAllPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", getPost).Methods("GET")
	router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", patchPost).Methods("PATCH")
	router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")

	http.ListenAndServe(":5000", router)

}

func getPost(w http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}
	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified ID"))
		return
	}
	post := posts[id]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func getAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func addPost(w http.ResponseWriter, r *http.Request) {
	var newPost Post

	json.NewDecoder(r.Body).Decode(&newPost)
	posts = append(posts, newPost)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
	}

	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specied id"))
		return
	}

	var updatedPost Post
	json.NewDecoder(r.Body).Decode(&updatedPost)

	posts[id] = updatedPost

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedPost)

}

func patchPost(w http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("ID could not be converted to integer"))
	}

	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specied id"))
		return
	}

	// get current value

	post := &posts[id]
	json.NewDecoder(r.Body).Decode(post)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)

}

func deletePost(w http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("ID could not be converted to integer"))
	}

	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specied id"))
		return
	}

	posts = append(posts[:id], posts[id+1:]...)

	w.WriteHeader(200)

}
