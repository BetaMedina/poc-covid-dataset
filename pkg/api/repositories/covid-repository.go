package repositories

import (
	"dataset-covid/pkg/config"
	"dataset-covid/pkg/models"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DATASET = "covid_dataset"
const SCHEMA = "covid_data"

func GetCovidPerRegion(state string, limit int64, page int64) ([]models.CovidModel, int64, error) {
	connection := config.ConfigDb()
	collection := config.GetColleciton(connection, DATASET, SCHEMA)
	filter := bson.M{"state": strings.ToUpper(state)}

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * limit)
	findOptions.SetLimit(limit)

	documents, _ := collection.Find(config.Ctx, filter, findOptions)
	if documents == nil {
		return nil, 0, nil
	}
	var rows []models.CovidModel
	for documents.Next(config.Ctx) {
		var row models.CovidModel

		documents.Decode(&row)
		rows = append(rows, row)
	}

	total, _ := collection.CountDocuments(config.Ctx, filter)
	return rows, total, nil
}

func GetCovidPerCity(city string, limit int64, page int64) ([]models.CovidModel, int64, error) {
	connection := config.ConfigDb()
	collection := config.GetColleciton(connection, DATASET, SCHEMA)
	filter := bson.M{"city": city}
	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * limit)
	findOptions.SetLimit(limit)

	documents, _ := collection.Find(config.Ctx, filter, findOptions)
	if documents == nil {
		return nil, 0, nil
	}
	var rows []models.CovidModel
	for documents.Next(config.Ctx) {
		var row models.CovidModel

		documents.Decode(&row)
		rows = append(rows, row)
	}

	total, _ := collection.CountDocuments(config.Ctx, filter)
	return rows, total, nil
}
