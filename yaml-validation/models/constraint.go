package models

import (
	"context"
	"strings"
	"yaml-validation/pkg/setting"
	"yaml-validation/services"

	"go.mongodb.org/mongo-driver/bson"
)

type Constraint struct {
	Path             string             `json:"path,omitempty"`
	Min              *float32           `json:"min,omitempty"`
	Max              *float32           `json:"max,omitempty"`
	Enum             []string           `json:"enum,omitempty"`
	Regex            string             `json:"regex,omitempty"`
	Disabled         bool               `json:"disabled,omitempty"`
	GroupKindVersion []GroupKindVersion `json:"x-kubernetes-group-version-kind,omitempty"`
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

func GetConstraint(path string, gkv *GroupKindVersion) *Constraint {
	var constraint Constraint
	gkvLower := GkvToLower(gkv)

	err := services.GetClient().
		Database(setting.DatabaseSetting.Name).
		Collection("Constraints").
		FindOne(context.TODO(), bson.M{"path": path, "groupkindversion": bson.M{"$elemMatch": gkvLower}}).
		Decode(&constraint)

	if err != nil {
		return nil
	}

	return &constraint
}

func DeleteConstraint(path string, gkv *GroupKindVersion) {
	gkvLower := GkvToLower(gkv)

	services.GetClient().
		Database(setting.DatabaseSetting.Name).
		Collection("Constraints").
		DeleteMany(context.TODO(), bson.M{"path": path, "groupkindversion": bson.M{"$elemMatch": gkvLower}})
}

func GkvToLower(gkv *GroupKindVersion) *GroupKindVersion {
	var gkvLower GroupKindVersion
	gkvLower.Group = strings.ToLower(gkv.Group)
	gkvLower.Kind = strings.ToLower(gkv.Kind)
	gkvLower.Version = strings.ToLower(gkv.Version)

	return &gkvLower
}
