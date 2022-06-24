package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Movie struct {
	MovieID   string `json:"movieid"`
	MovieName string `json:"moviename"`
}

type JsonResponse struct {
	Type    string  `json:"type"`
	Data    []Movie `json:"data"`
	Message string  `json:"message"`
}

// DB set up
func setupDB() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	fmt.Println("Setting up DB")
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable\n", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("PORT"))
	db, err := sql.Open("postgres", dbinfo)

	checkErr(err)

	return db
}

// Function for handling errors
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

func getMovies(writer http.ResponseWriter, reader *http.Request) {
	log.Println("Endpoint hit: /movies")
	db := setupDB()

	printMessage("Getting movies...")

	rows, err := db.Query("SELECT * FROM movies")

	checkErr(err)

	var movies []Movie

	// Iterate through the query results
	for rows.Next() {
		var id int
		var movieID string
		var movieName string

		err = rows.Scan(&id, &movieID, &movieName)

		checkErr(err)

		movies = append(movies, Movie{MovieID: movieID, MovieName: movieName})
	}

	var response = JsonResponse{Type: "success", Data: movies, Message: "Successfully got all movies from DB"}
	printMessage("Successfully got all movies from DB")
	json.NewEncoder(writer).Encode(response)
}

func getMovie(writer http.ResponseWriter, reader *http.Request) {
	log.Println("Endpoint hit: /getmovie/{movieid}")
	// Get the map of route variables from the reader
	params := mux.Vars(reader)

	movieID := params["movieid"]

	var response = JsonResponse{}

	// movieID must be provided
	if movieID == "" {
		response = JsonResponse{Type: "error", Message: "You are missing the movieid parameter."}
	} else {
		db := setupDB()
		printMessage("Getting movie from DB")

		// Execute the query but don't return any rows
		row := db.QueryRow("SELECT * FROM movies WHERE movieid = $1", movieID)

		var movies []Movie

		var id int
		var movieID string
		var movieName string

		err := row.Scan(&id, &movieID, &movieName)

		if err != nil {
			response = JsonResponse{Type: "failure", Message: "A movie with that movieid does not exist."}
		} else {
			movies = append(movies, Movie{MovieID: movieID, MovieName: movieName})

			printMessage("Successfully got movie from DB")
			response = JsonResponse{Type: "success", Data: movies, Message: "Successfully got movie from DB"}
		}
	}

	json.NewEncoder(writer).Encode(response)
}

func createMovie(writer http.ResponseWriter, reader *http.Request) {
	log.Println("Endpoint hit: /addmovie")

	err := reader.ParseForm()
	checkErr(err)

	movieID := reader.FormValue("movieid")
	movieName := reader.FormValue("moviename")

	var response = JsonResponse{}
	// movieID and movieName must both be provided
	if movieID == "" || movieName == "" {
		log.Println("Bad")
		response = JsonResponse{Type: "error", Message: "You are missing movieID or movieName"}
	} else {
		// Setup the DB and insert a new record
		db := setupDB()
		printMessage("Inserting movie into DB")
		fmt.Printf("Inserting new movie with ID %s and name %s\n", movieID, movieName)
		var lastInsertID int
		// Execute the query and get the first (and only) row
		err := db.QueryRow("INSERT INTO movies(movieid, moviename) VALUES($1, $2)", movieID, movieName).Scan(&lastInsertID)
		checkErr(err)
		response = JsonResponse{Type: "success", Message: "The movie has been inserted successfully!"}
	}
	json.NewEncoder(writer).Encode(response)
}

func deleteMovie(writer http.ResponseWriter, reader *http.Request) {
	log.Println("Endpoint hit: /deletemovie/{movieid}")
	// Get the map of route variables from the reader
	params := mux.Vars(reader)

	movieID := params["movieid"]

	var response = JsonResponse{}

	// movieID must be provided
	if movieID == "" {
		response = JsonResponse{Type: "error", Message: "You are missing the movieid parameter."}
	} else {
		db := setupDB()
		printMessage("Deleting movie from DB")

		// Execute the query but don't return any rows
		_, err := db.Exec("DELETE FROM movies WHERE movieid = $1", movieID)

		if err != nil {
			response = JsonResponse{Type: "failure", Message: "Failed to delete the specified movie."}
		} else {
			response = JsonResponse{Type: "success", Message: "The movie has been deleted successfully."}
		}
	}

	json.NewEncoder(writer).Encode(response)
}

func deleteAllMovies(writer http.ResponseWriter, reader *http.Request) {
	log.Println("Endpoint hit: /deletemovies")
	db := setupDB()

	printMessage("Deleting all movies...")

	_, err := db.Exec("DELETE FROM movies")

	checkErr(err)

	printMessage("All movies have been deleted successfully!")

	var response = JsonResponse{Type: "success", Message: "All movies have been deleted successfully!"}

	json.NewEncoder(writer).Encode(response)
}

func main() {

	// Init the mux router
	router := mux.NewRouter().StrictSlash(true)

	// Route handles & endpoints

	// Get all movies
	router.HandleFunc("/movies", getMovies).Methods("GET")

	// Get a specific movie by the movieID
	router.HandleFunc("/getmovie/{movieid}", getMovie).Methods("GET")

	// Create a movie
	router.HandleFunc("/addmovie", createMovie).Methods("POST")

	// Delete a specific movie by the movieID
	router.HandleFunc("/deletemovie/{movieid}", deleteMovie).Methods("DELETE")

	// Delete all movies
	router.HandleFunc("/deletemovies", deleteAllMovies).Methods("DELETE")

	// Serve the app
	fmt.Println("Server at 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
