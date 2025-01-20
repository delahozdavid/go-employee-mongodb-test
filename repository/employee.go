package repository

import (
	"context"
	"fmt"
	"go-mongo-project/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeRepo struct {
	MongoCollection *mongo.Collection
}

func (r *EmployeeRepo) InsertEmployee(emp *model.Employee) (*mongo.InsertOneResult, error) {
	return r.MongoCollection.InsertOne(context.Background(), emp)
}

func (r *EmployeeRepo) FindEmployee(empID string) (*model.Employee, error) {
	var emp model.Employee
	err := r.MongoCollection.FindOne(context.Background(), bson.M{"employee_id": empID}).Decode(&emp)
	if err != nil {
		return nil, err
	}
	return &emp, nil
}

func (r *EmployeeRepo) FindAllEmployee() ([]model.Employee, error) {
	results, err := r.MongoCollection.Find(context.Background(), bson.D{})

	if err != nil {
		return nil, err
	}

	var emps []model.Employee
	err = results.All(context.Background(), &emps)
	if err != nil {
		return nil, fmt.Errorf("error while decoding employees: %v", err)
	}

	return emps, nil
}

func (r *EmployeeRepo) UpdateEmployeeID(empID string, updateEmp *model.Employee) (int64, error) {
	result, err := r.MongoCollection.UpdateOne(context.Background(), bson.D{{Key: "employee_id", Value: empID}}, bson.D{bson.E{Key: "$set", Value: updateEmp}})
	if err != nil {
		return 0, err
	}

	return result.ModifiedCount, nil
}

func (r *EmployeeRepo) DeleteEmployee(empID string) (int64, error) {
	result, err := r.MongoCollection.DeleteOne(context.Background(), bson.D{{Key: "employee_id", Value: empID}})
	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}

func (r *EmployeeRepo) DeleteAllEmployee() (int64, error) {
	result, err := r.MongoCollection.DeleteMany(context.Background(), bson.D{})
	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}
