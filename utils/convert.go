package utils

import (
	"encoding/json"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConvertStringToObjectId(id string) primitive.ObjectID {
	objectId, _ := primitive.ObjectIDFromHex(id)
	return objectId
}

func ConvertObjectIdToString(id primitive.ObjectID) string {
	objectString := id.Hex()

	return objectString
}

func EncodeMessage(obj any) []byte {

	raw, err := json.Marshal(obj)
	if err != nil {
		log.Printf("Error: Marshal Failed %s", err.Error())
		return nil
	}

	return raw
}
