package model

type Employee struct {
	ID         string `bson:"_id,omitempty"`
	Name       string `bson:"name"`
	Department string `bson:"department"`
	EmployeeID string `bson:"employee_id"`
	Role       string `bson:"role"`
}
