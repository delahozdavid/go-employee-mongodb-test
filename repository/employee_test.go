package repository

import (
	"context"
	"go-mongo-project/model"
	"log"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newMongoClient() *mongo.Client {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatalf("MONGODB_URI not set in .env file")
	}

	mongoTestClient, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	return mongoTestClient
}

func TestMongoOperations(t *testing.T) {
	mongoTestClient := newMongoClient()
	defer mongoTestClient.Disconnect(context.Background())

	// dummy data
	employee_1 := uuid.New().String()
	employee_2 := uuid.New().String()

	// connect to collection
	collection := mongoTestClient.Database("companydb").Collection("employee_test")

	employeeRepo := EmployeeRepo{MongoCollection: collection}

	// Insert Employee 1 data
	t.Run("Insert Employee 1", func(t *testing.T) {
		emp := model.Employee{
			Name:       "David Castro",
			Department: "IT",
			EmployeeID: employee_1,
			Role:       "Software Engineer",
		}

		_, err := employeeRepo.InsertEmployee(&emp)
		if err != nil {
			t.Errorf("error while inserting employee: %v", err)
		}
		t.Log("Employee 1 inserted")
	})

	// Insert Employee 2 data
	t.Run("Insert Employee 2", func(t *testing.T) {
		emp := model.Employee{
			Name:       "José Raúl Del Bosque",
			Department: "Project Manager",
			EmployeeID: employee_2,
			Role:       "Project Manager",
		}

		_, err := employeeRepo.InsertEmployee(&emp)
		if err != nil {
			t.Errorf("error while inserting employee: %v", err)
		}
		t.Log("Employee 2 inserted")
	})

	// Get Employee 1 data
	t.Run("Get Employee 1", func(t *testing.T) {
		result, err := employeeRepo.FindEmployee(employee_1)

		if err != nil {
			t.Errorf("error while getting employee: %v", err)
		} else {
			t.Log("Employee 1 found: ", result.Name)
		}
	})

	// Get all employees data
	t.Run("Get All Employees", func(t *testing.T) {
		results, err := employeeRepo.FindAllEmployee()

		if err != nil {
			t.Errorf("error while getting employees: %v", err)
		}

		t.Log("Employees found: ", results)
	})

	// Update Employee 1 data
	t.Run("Update Employee 1", func(t *testing.T) {
		emp := model.Employee{
			Name:       "David Castro",
			Department: "IT",
			EmployeeID: employee_1,
		}

		emp.Name = "David Castro Sr."

		_, err := employeeRepo.UpdateEmployeeID(employee_1, &emp)
		if err != nil {
			t.Errorf("error while updating employee: %v", err)
		}

		t.Log("Employee 1 updated")
	})

	// Delete Employee 1 data
	t.Run("Delete Employee 1", func(t *testing.T) {
		_, err := employeeRepo.DeleteEmployee(employee_1)
		if err != nil {
			t.Fatalf("Failed to delete employee: %v", err)
		}
		t.Log("Employee 1 deleted")
	})
}
