package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"yaml-validation/services"
)

type Constraint struct {
	Path     string `json:"path"`
	Kind     string `json:"kind"`
	Regex    string `json:"regex"`
	Disabled bool   `json:"disabled"`
}

func SaveConstraint(constraint Constraint) error {
	collection := services.GetClient().
		Database("beeyond_validation_db").
		Collection("Constraints")

	_, err := collection.InsertOne(context.TODO(), constraint)
	return err
}

func GetConstraints() []*Constraint {
	var constraints []*Constraint

	// TODO: error handling?
	cur, _ := services.GetClient().
		Database("beeyond_validation_db").
		Collection("Constraints").
		Find(context.TODO(), bson.D{})

	for cur.Next(context.TODO()) {
		var constr Constraint
		_ = cur.Decode(&constr)
		// TODO: append only when there is no error?
		constraints = append(constraints, &constr)
	}

	_ = cur.Close(context.TODO())
	return constraints
}
