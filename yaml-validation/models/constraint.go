package models

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"yaml-validation/conf"
	"yaml-validation/services"

	"go.mongodb.org/mongo-driver/bson"
)

type Constraint struct {
	Path             string           `json:"-"`
	Min              *float32         `json:"min,omitempty"`
	Max              *float32         `json:"max,omitempty"`
	Enum             []string         `json:"enum,omitempty"`
	Regex            *string          `json:"regex,omitempty"`
	Disabled         bool             `json:"disabled,omitempty"`
	GroupKindVersion GroupKindVersion `json:"-"`
}

// Extension method to Constraint
// checks if exactly one of MinMax, Enum and Regex is filled out, but not multiple
// and if MinMax is present if the valueType is integer, since it can only be used on integer fields
// Parameters:
// 		- constraint (Constraint): represents the constraint we want to validate
// 		- valueType (string): represents the type of the property (e.g. integer, string, ...)
// Returns: bool: true if valid, otherwise false
func (constraint Constraint) IsValid(valueType string) bool {
	if constraint.Enum == nil && constraint.Min == nil && constraint.Max == nil && constraint.Regex == nil {
		return false
	}

	isValidEnum := constraint.Enum != nil && constraint.Min == nil && constraint.Max == nil && constraint.Regex == nil
	isValidMinMax := constraint.Enum == nil && constraint.Min != nil && constraint.Max != nil && constraint.Regex == nil && valueType == "integer"
	isValidRegex := constraint.Enum == nil && constraint.Regex != nil && constraint.Min == nil && constraint.Max == nil

	return isValidEnum || isValidMinMax || isValidRegex
}

// Saves a constraint in the database
// Parameter: constraint (Constraint): represents the constraint we want to store
// Returns: error if anny occur when inserting
func SaveConstraint(constraint Constraint) error {
	collection := services.GetClient().
		Database(conf.Configuration.Database.Name).
		Collection("Constraints")

	_, err := collection.InsertOne(context.TODO(), constraint)
	return err
}

// Gets the constraint that correspond to the given groupKindVersion and path from the database
// Parameters:
// 		- path (string): represents the path of the we want to get
// 		- groupKindVersion (*GroupKindVersion): represents the Group, Kind and Version of the constraint we want to get
// Returns: *Constraints: Represents the constraint that matches the path and group kind version
func GetConstraint(path string, groupKindVersion GroupKindVersion) *Constraint {
	var constraint Constraint

	err := services.GetClient().
		Database(conf.Configuration.Database.Name).
		Collection("Constraints").
		FindOne(context.TODO(), bson.M{"path": path, "groupkindversion": groupKindVersion}).
		Decode(&constraint)

	if err != nil {
		return nil
	}

	return &constraint
}

// Gets all constraint that correspond to the given groupKindVersion from the database
// Parameter: groupKindVersion (*GroupKindVersion): represents the Group, Kind and Version of the constraints we want to get
// Returns: []*Constraints: An array of all constraints found constraints that match to given groupKindVersion
func GetConstraintsByGKV(groupKindVersion *GroupKindVersion) []*Constraint {
	var constraints []*Constraint

	cur, err := services.GetClient().
		Database(conf.Configuration.Database.Name).
		Collection("Constraints").
		Find(context.TODO(), bson.M{"disabled": false, "groupkindversion": groupKindVersion})

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

// Deletes all Constraints from the database that match the given path and groupKindVersion
// Parameters:
//		- path (string): Represents the path of the constraint we want to delete
//		- groupKindVersion (GroupKindVersion): Represents the Group, Kind and Version of the constraint we want to delete
// Returns: the deleteResult  (contains the number of deleted documents)
func DeleteConstraint(path string, groupKindVersion GroupKindVersion) *mongo.DeleteResult {
	deleteResult, _ := services.GetClient().
		Database(conf.Configuration.Database.Name).
		Collection("Constraints").
		DeleteMany(context.TODO(), bson.M{"path": path, "groupkindversion": groupKindVersion})
	return deleteResult
}

// Deletes all Constraints from the database
// Returns: the deleteResult  (contains the number of deleted documents)
func DeleteAll() *mongo.DeleteResult {
	deleteResult, _ := services.GetClient().
		Database(conf.Configuration.Database.Name).
		Collection("Constraints").
		DeleteMany(context.TODO(), bson.M{})
	return deleteResult
}
