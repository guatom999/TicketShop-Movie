package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

func ConvertStringToObjectId(id string) primitive.ObjectID {
	objectId, _ := primitive.ObjectIDFromHex(id)
	return objectId
}

func ConvertObjectIdToString(id primitive.ObjectID) string {
	objectString := id.Hex()

	return objectString
}
