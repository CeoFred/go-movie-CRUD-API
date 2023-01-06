package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	Director *Director `json:"director"`
	ID   string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
}

type Director struct {
	FirstName string `json:"firstname"`
  LastName  string `json:"lastname"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{
		ID: "1",
		Isbn: "431",
		Title: "Movie 1",
		Director: &Director{
			FirstName: "John",
      LastName:  "Doe",
		},
	})
	r.HandleFunc("/movies", GetMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", GetMovie).Methods("GET")
	r.HandleFunc("/movies", CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", DeleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8080 \n")
	log.Fatal(http.ListenAndServe(":8080", r))

}

func DeleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, movie := range movies {
		if movie.ID == params["id"] {
      movies = append(movies[:i], movies[i+1:]...)
      break
    }
	}

	json.NewEncoder(w).Encode(movies)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

  for i, movie := range movies {
		if movie.ID == params["id"] {
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies[i] = movie
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}


func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
	var movie Movie

	 for i, item := range movies {
		if item.ID == params["id"] {
      movie =  movies[i]
      break
    }
	}

	json.NewEncoder(w).Encode(movie)

}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type","application/json")
		var movie Movie
		_ = json.NewDecoder(r.Body).Decode(&movie)
		movie.ID = strconv.Itoa(rand.Intn(10000000))
		movies = append(movies, movie)
		json.NewEncoder(w).Encode(movie)
}