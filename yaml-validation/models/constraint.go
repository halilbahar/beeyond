package models

import (
	"context"
	"yaml-validation/pkg/setting"
	"yaml-validation/services"

	"go.mongodb.org/mongo-driver/bson"
)

type Constraint struct {
	Path     string `json:"path"`
	Kind     string `json:"kind"`
	Regex    string `json:"regex"`
	Disabled bool   `json:"disabled"`
}

func SaveConstraint(constraint Constraint) error {
	collection := services.GetClient().
		Database(setting.DatabaseSetting.Name).
		Collection("Constraints")

	_, err := collection.InsertOne(context.TODO(), constraint)
	return err
}

func GetConstraints() []*Constraint {
	var constraints []*Constraint

	cur, err := services.GetClient().
		Database(setting.DatabaseSetting.Name).
		Collection("Constraints").
		Find(context.TODO(), bson.D{})

	if err != nil {
		return nil
	}

	for cur.Next(context.TODO()) {
		var constr Constraint

		if cur.Decode(&constr) != nil {
			constraints = append(constraints, &constr)
		}
	}

	_ = cur.Close(context.TODO())
	return constraints
}
