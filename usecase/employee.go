package usecase

import (
	"encoding/json"
	"go-mongo-project/model"
	"go-mongo-project/repository"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeService struct {
	MongoCollection *mongo.Collection
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func (svc *EmployeeService) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	var emp model.Employee
	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid body", err)
		res.Error = err.Error()
		return
	}

	// assign new employee id
	emp.EmployeeID = uuid.NewString()

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	// insert employee
	insertID, err := repo.InsertEmployee(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error inserting employee", err)
		res.Error = err.Error()
		return
	}

	res.Data = emp.EmployeeID
	w.WriteHeader(http.StatusOK)

	log.Println("Employee inserted with ID", insertID, emp)
}

func (svc *EmployeeService) GetEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	//get employee id
	empID := mux.Vars(r)["id"]
	log.Println("Employee ID: ", empID)

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	emp, err := repo.FindEmployee(empID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error getting employee", err)
		res.Error = err.Error()
		return

	}

	res.Data = emp
	w.WriteHeader(http.StatusOK)
}

func (svc *EmployeeService) GetAllEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	emp, err := repo.FindAllEmployee()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error getting employee", err)
		res.Error = err.Error()
		return

	}

	res.Data = emp
	w.WriteHeader(http.StatusOK)
}
func (svc *EmployeeService) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	empID := mux.Vars(r)["id"]
	log.Println("Employee Id: ", empID)

	if empID == "" {
		w.WriteHeader(http.StatusBadRequest)
		res.Error = "Invalid employee id"
		return
	}

	var emp model.Employee

	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid body", err)
		res.Error = err.Error()
		return
	}

	emp.EmployeeID = empID

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}
	count, err := repo.UpdateEmployeeID(empID, &emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error updating employee", err)
		res.Error = err.Error()
		return
	}

	res.Data = count
	w.WriteHeader(http.StatusOK)

}
func (svc *EmployeeService) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	empID := mux.Vars(r)["id"]
	log.Println("Employee Id: ", empID)

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	count, err := repo.DeleteEmployee(empID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error deleting employee", err)
		res.Error = err.Error()
		return
	}

	res.Data = count
	w.WriteHeader(http.StatusOK)

}
