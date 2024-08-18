package mongodb

import (
	"encoding/json"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetUpdateFieldsSelected ...
func GetUpdateFieldsSelected(request interface{}, updateFields []string) bson.M {
	var mapReq bson.M
	data, _ := bson.Marshal(request)
	_ = bson.Unmarshal(data, &mapReq)

	result := bson.M{}
	for _, v := range updateFields {
		if val, ok := mapReq[v]; ok {
			result[v] = val
		}
	}
	return result
}
func ConvertToObjectIds(ids []string) ([]primitive.ObjectID, error) {
	var objectIDs []primitive.ObjectID
	for _, id := range ids {
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		objectIDs = append(objectIDs, objectID)
	}

	return objectIDs, nil
}

// ConstructMongoPipeline Construct aggregation pipeline from raw string
func ConstructMongoPipeline(str string) (mongo.Pipeline, error) {
	var pipeline []bson.D
	str = strings.TrimSpace(str)
	if strings.Index(str, "[") != 0 {
		var doc bson.M
		err := json.Unmarshal([]byte(str), &doc)
		if err != nil {
			return nil, err
		}
		var v bson.D
		b, err := bson.Marshal(doc)
		if err != nil {
			return nil, err
		}
		err = bson.Unmarshal(b, &v)
		if err != nil {
			return nil, err
		}
		pipeline = append(pipeline, v)
	} else {
		var docs []bson.M
		err := json.Unmarshal([]byte(str), &docs)
		if err != nil {
			fmt.Printf("json.Unmarshal: %v", err.Error())
			return nil, err
		}
		for _, doc := range docs {
			var v bson.D
			b, err := bson.Marshal(doc)
			if err != nil {
				fmt.Printf("bson.Marshal: %v", err.Error())
				return nil, err
			}
			err = bson.Unmarshal(b, &v)
			if err != nil {
				fmt.Printf("bson.Unmarshal: %v", err.Error())
				return nil, err
			}
			pipeline = append(pipeline, v)
		}
	}
	return pipeline, nil
}
