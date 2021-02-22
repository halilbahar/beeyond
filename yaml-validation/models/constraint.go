package models

import (
	"context"
	"yaml-validation/pkg/setting"
	"yaml-validation/services"

	"go.mongodb.org/mongo-driver/bson"
)

type Constraint struct {
	Path             string
	Min              *float32 `json:"min,omitempty"`
	Max              *float32 `json:"max,omitempty"`
	Enum             []string `json:"enum,omitempty"`
	Regex            string   `json:"regex,omitempty"`
	Disabled         bool     `json:"disabled,omitempty"`
	GroupKindVersion []GroupKindVersion
}

func (constraint Constraint) IsValid() bool {
	if constraint.Enum == nil && constraint.Min == nil && constraint.Max == nil && constraint.Regex == "" {
		return false
	}

	isValidEnum := constraint.Enum != nil && constraint.Min == nil && constraint.Max == nil && constraint.Regex == ""
	isValidMinMax := constraint.Enum == nil && constraint.Min != nil && constraint.Max != nil && constraint.Regex == ""
	isValidRegex := constraint.Enum == nil && constraint.Regex != "" && constraint.Min == nil && constraint.Max == nil

	return isValidEnum || isValidMinMax || isValidRegex
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

		if err := cur.Decode(&constr); err == nil {
			constraints = append(constraints, &constr)
		}
	}

	_ = cur.Close(context.TODO())
	return constraints
}

func GetConstraint(path string, groupKindVersion *GroupKindVersion) *Constraint {
	var constraint Constraint

	err := services.GetClient().
		Database(setting.DatabaseSetting.Name).
		Collection("Constraints").
		FindOne(context.TODO(), bson.M{"path": path, "groupkindversion": bson.M{"$elemMatch": groupKindVersion.ToLower()}}).
		Decode(&constraint)

	if err != nil {
		return nil
	}

	return &constraint
}

func DeleteConstraint(path string, groupKindVersion *GroupKindVersion) {
	_, _ = services.GetClient().
		Database(setting.DatabaseSetting.Name).
		Collection("Constraints").
		DeleteMany(context.TODO(), bson.M{"path": path, "groupkindversion": bson.M{"$elemMatch": groupKindVersion.ToLower()}})
}
