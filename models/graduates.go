package models

import (
	"context"
	"errors"
	"gts/config"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Graduates struct {
	FullName     string `json:"fullname" bson:"fullname"`
	Email        string `json:"email" bson:"email"`
	PhoneNumber  string `json:"phoneNumber" bson:"phoneNumber"`
	Telegram     string `json:"telegram" bson:"telegram"`
	Gender       string `json:"gender" bson:"gender"`
	PlaceOfBirth string `json:"placeOfBirth" bson:"placeOfBirth"`
	BirthDate    string `json:"birthDate" bson:"birthDate"`
	Education    string `json:"education" bson:"education"`
	Major        string `json:"major" bson:"major"`
	Domicile     string `json:"domicile" bson:"domicile"`
	Occupation   string `json:"occupation" bson:"occupation"`
	Company      string `json:"company" bson:"company"`
	LinkedinURL  string `json:"linkedinURL" bson:"linkedinURL"`
}

func (grad *Graduates) Add() error {
	graduatesCollection := config.DB().Database("gts").Collection("graduates")
	_, err := graduatesCollection.InsertOne(context.TODO(), grad)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (grad *Graduates) GetByName(name string) (Graduates, error) {
	var data Graduates
	filter := bson.M{
		"fullname": name,
	}

	graduatesCollection := config.DB().Database("gts").Collection("graduates")

	if err := graduatesCollection.FindOne(context.TODO(), filter).Decode(&data); err != nil {
		if err == mongo.ErrNoDocuments {
			return Graduates{}, errors.New("Data was not found")
		}
		return Graduates{}, err
	}
	return data, nil
}
