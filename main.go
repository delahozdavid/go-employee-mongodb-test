package main

import (
	"context"
	"go-mongo-project/usecase"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Print("Connecting to MongoDB")

	// create mongo client
	mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))

	if err != nil {
		log.Fatal("Error connecting to MongoDB", err)
	}

	err = mongoClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Error pinging MongoDB", err)
	}

	log.Print("Connection to MongoDB successful")

}

func main() {
	// close the mongo connection
	defer mongoClient.Disconnect(context.Background())

	collection := mongoClient.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))
	empService := usecase.EmployeeService{MongoCollection: collection}
	// create employee service

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/health", healthHandler).Methods("GET")

	r.HandleFunc("/api/v1/employee", empService.CreateEmployee).Methods("POST")
	r.HandleFunc("/api/v1/employee/{id}", empService.GetEmployee).Methods("GET")
	r.HandleFunc("/api/v1/employee", empService.GetAllEmployee).Methods("GET")
	r.HandleFunc("/api/v1/employee/{id}", empService.UpdateEmployee).Methods("PUT")
	r.HandleFunc("/api/v1/employee/{id}", empService.DeleteEmployee).Methods("DELETE")

	log.Print("Starting server on port 8080")
	http.ListenAndServe(":8080", r)

}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("mongodb-go-api is up and running"))
}
