package models

import (
	"context"
	"yaml-validation/pkg/setting"
	"yaml-validation/services"

	"go.mongodb.org/mongo-driver/bson"
)

type Constraint struct {
	Path     string   `json:"path,omitempty"`
	Kind     string   `json:"kind,omitempty"`
	Min      float32  `json:"min,omitempty"`
	Max      float32  `json:"max,omitempty"`
	Enum     []string `json:"enum,omitempty"`
	Regex    string   `json:"regex,omitempty"`
	Disabled bool     `json:"disabled,omitempty"`
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

	// TODO: error handling?
	cur, _ := services.GetClient().
		Database(setting.DatabaseSetting.Name).
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

func GetConstraint(path string, kind string) *Constraint {
	var constraint Constraint

	err := services.GetClient().
		Database(setting.DatabaseSetting.Name).
		Collection("Constraints").
		FindOne(context.TODO(), bson.M{"path": path, "kind": kind}).
		Decode(&constraint)

	if err != nil {
		return nil
	}

	return &constraint
}

func DeleteConstraint(path string, kind string) {
	services.GetClient().
		Database(setting.DatabaseSetting.Name).
		Collection("Constraints").
		DeleteMany(context.TODO(), bson.M{"path": path, "kind": kind})
}
