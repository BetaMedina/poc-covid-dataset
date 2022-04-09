package main

import (
	"dataset-covid/pkg/config"
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const CSV_HEADER int = 0

func ConvertInt32(initialValue *string) int32 {
	convertedValue, err := strconv.Atoi(*initialValue)
	if err != nil {
		return 0
	}
	return int32(convertedValue)
}
func ConvertInt64(initialValue *string) int64 {
	convertedValue, err := strconv.Atoi(*initialValue)
	if err != nil {
		return 0
	}
	return int64(convertedValue)
}

func ConvertFloat64(initialValue *string) float64 {
	convertedValue, err := strconv.ParseFloat(*initialValue, 64)
	if err != nil {
		return 0.0
	}
	return convertedValue
}

func insertCollectionRows(collectionRows *[]interface{}, collection *mongo.Collection, channel <-chan interface{}, hasFinished <-chan bool) {
	start := time.Now()
	for {
		select {
		case row := <-channel:
			if len(*collectionRows) == 100 {
				collection.InsertMany(config.Ctx, *collectionRows)
				*collectionRows = nil
			}
			*collectionRows = append(*collectionRows, row)
		case <-hasFinished:
			collection.InsertMany(config.Ctx, *collectionRows)
			elapsed := time.Since(start)
			log.Println(elapsed)
			return
		}
	}
}

func main() {
	channel := make(chan interface{}, 2)
	hasFinished := make(chan bool)
	file, _ := os.Open("./public/caso.csv")
	defer file.Close()
	reader := csv.NewReader(file)
	connection := config.ConfigDb()
	collection := config.GetColleciton(connection, "covid_dataset", "covid_data")
	go func() {
		i := 0
		for record, err := reader.Read(); err == nil; record, err = reader.Read() {
			if i != CSV_HEADER {
				formattedRow := bson.M{
					"date":                           record[0],
					"state":                          record[1],
					"city":                           record[2],
					"place_type":                     record[3],
					"confirmed":                      ConvertInt32(&record[4]),
					"deaths":                         ConvertInt32(&record[5]),
					"order_for_place":                ConvertInt32(&record[6]),
					"estimated_population_2019":      ConvertInt32(&record[8]),
					"estimated_population":           ConvertInt32(&record[9]),
					"city_ibge_code":                 ConvertInt32(&record[10]),
					"confirmed_per_100k_inhabitants": ConvertFloat64(&record[11]),
					"death_rate":                     ConvertFloat64(&record[12]),
				}
				channel <- formattedRow
			}
			i++
		}
		hasFinished <- true
	}()
	var collectionRows []interface{}
	defer insertCollectionRows(&collectionRows, collection, channel, hasFinished)
}
