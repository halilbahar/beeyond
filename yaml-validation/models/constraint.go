package models

import (
	"../services"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

type Constraint struct {
	Path     string `json:"path"`
	Kind     string `json:"kind"`
	Regex    string `json:"regex"`
	Disabled bool   `json:"disabled"`
}

func SaveConstraintToDb(constraint Constraint) error {
	collection := services.GetClient().Database("beeyond_validation_db").Collection("Constraints")
	_, err := collection.InsertOne(context.TODO(), constraint)
	fmt.Print(err)
	return err
}

func GetConstraints() []*Constraint {
	var constraints []*Constraint

	cur, _ := services.GetClient().Database("beeyond_validation_db").Collection("Constraints").Find(context.TODO(), bson.D{})
	for cur.Next(context.TODO()) {
		var constr Constraint
		cur.Decode(&constr)
		constraints = append(constraints, &constr)
	}

	cur.Close(context.TODO())
	return constraints
}
