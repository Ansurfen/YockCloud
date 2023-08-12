package main

// #include "../cloud.h"
import "C"
import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"unsafe"
	"yockcloud/cgo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoOption struct {
	URL                string `json:"url"`
	options.Credential `json:"cred"`
}

//export Dial
func Dial(dataSourceName *C.char) *C.component {
	opt := MongoOption{}
	err := json.Unmarshal([]byte(C.GoString(dataSourceName)), &opt)
	if err != nil {
		panic(err)
	}
	cli, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(opt.URL).SetAuth(opt.Credential))
	if err != nil {
		panic(err)
	}
	ret := cgo.Malloc(C.component{})
	ret.ptr = unsafe.Pointer(cli)
	return ret
}

//export Ping
func Ping(comp *C.component) *C.char {
	conn := cgo.CastPtr[mongo.Client](comp.ptr)
	err := conn.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		return C.CString(err.Error())
	}
	return nil
}

//export InsertOne
func InsertOne(comp *C.component, db, col, stu *C.char) *C.char {
	conn := cgo.CastPtr[mongo.Client](comp.ptr)
	var filter any
	err := bson.UnmarshalExtJSON([]byte(C.GoString(stu)), true, &filter)
	if err != nil {
		return (*C.char)(cgo.JSON(cgo.MAP{
			"data": nil,
			"err":  err.Error(),
		}))
	}
	collection := conn.Database(C.GoString(db)).Collection(C.GoString(col))
	res, err := collection.InsertOne(context.TODO(), filter)
	if err != nil {
		return (*C.char)(cgo.JSON(cgo.MAP{
			"data": nil,
			"err":  err.Error(),
		}))
	}
	return (*C.char)(cgo.JSON(cgo.MAP{
		"data": res.InsertedID,
		"err":  nil,
	}))
}

//export InsertMany
func InsertMany(comp *C.component, db, col, stu *C.char) *C.char {
	conn := cgo.CastPtr[mongo.Client](comp.ptr)
	var filter []any
	err := bson.UnmarshalExtJSON([]byte(C.GoString(stu)), true, &filter)
	if err != nil {
		return (*C.char)(cgo.JSON(cgo.MAP{
			"data": nil,
			"err":  err.Error(),
		}))
	}
	collection := conn.Database(C.GoString(db)).Collection(C.GoString(col))
	res, err := collection.InsertMany(context.TODO(), filter)
	if err != nil {
		return (*C.char)(cgo.JSON(cgo.MAP{
			"data": nil,
			"err":  err.Error(),
		}))
	}
	return (*C.char)(cgo.JSON(cgo.MAP{
		"data": res.InsertedIDs,
		"err":  nil,
	}))
}

//export FindMany
func FindMany(comp *C.component, db, col, stu *C.char) *C.char {
	conn := cgo.CastPtr[mongo.Client](comp.ptr)
	var filter any
	err := bson.UnmarshalExtJSON([]byte(C.GoString(stu)), true, &filter)
	if err != nil {
		return (*C.char)(cgo.JSON(cgo.MAP{
			"data": nil,
			"err":  err.Error(),
		}))
	}
	collection := conn.Database(C.GoString(db)).Collection(C.GoString(col))
	ctx := context.TODO()
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return (*C.char)(cgo.JSON(cgo.MAP{
			"data": nil,
			"err":  err.Error(),
		}))
	}
	var result bson.A
	if err = cursor.All(ctx, &result); err != nil {
		return (*C.char)(cgo.JSON(cgo.MAP{
			"data": nil,
			"err":  err.Error(),
		}))
	}
	arr := []string{}
	for _, res := range result {
		data, err := bson.MarshalExtJSON(res, false, false)
		if err != nil {
			return (*C.char)(cgo.JSON(cgo.MAP{
				"data": nil,
				"err":  err.Error(),
			}))
		}
		arr = append(arr, string(data))
	}
	return C.CString(fmt.Sprintf(`{"data": [%s], "err": null}`, strings.Join(arr, ",")))
}

//export FindOne
func FindOne(comp *C.component, db, col, stu *C.char) *C.char {
	conn := cgo.CastPtr[mongo.Client](comp.ptr)
	var filter any
	err := bson.UnmarshalExtJSON([]byte(C.GoString(stu)), true, &filter)
	if err != nil {
		return (*C.char)(cgo.JSON(cgo.MAP{
			"data": nil,
			"err":  err.Error(),
		}))
	}
	collection := conn.Database(C.GoString(db)).Collection(C.GoString(col))
	var result bson.M
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return (*C.char)(cgo.JSON(cgo.MAP{
			"data": nil,
			"err":  err.Error(),
		}))
	}
	data, err := bson.MarshalExtJSON(result, false, false)
	if err != nil {
		return (*C.char)(cgo.JSON(cgo.MAP{
			"data": nil,
			"err":  err.Error(),
		}))
	}
	return C.CString(fmt.Sprintf(`{"data": [%s], "err": null}`, data))
}

//export DeleteOne
func DeleteOne(comp *C.component, db, col, stu *C.char) *C.char {
	conn := cgo.CastPtr[mongo.Client](comp.ptr)
	var filter any
	err := bson.UnmarshalExtJSON([]byte(C.GoString(stu)), true, &filter)
	if err != nil {
		return (*C.char)(cgo.JSON(cgo.MAP{
			"count": nil,
			"err":   err.Error(),
		}))
	}
	collection := conn.Database(C.GoString(db)).Collection(C.GoString(col))
	res, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return (*C.char)(cgo.JSON(cgo.MAP{
			"count": nil,
			"err":   err.Error(),
		}))
	}
	return (*C.char)(cgo.JSON(cgo.MAP{
		"count": res.DeletedCount,
		"err":   nil,
	}))
}

//export DeleteMany
func DeleteMany(comp *C.component, db, col, stu *C.char) *C.char {
	conn := cgo.CastPtr[mongo.Client](comp.ptr)
	var filter any
	err := bson.UnmarshalExtJSON([]byte(C.GoString(stu)), true, &filter)
	if err != nil {
		return (*C.char)(cgo.JSON(cgo.MAP{
			"count": nil,
			"err":   err.Error(),
		}))
	}
	collection := conn.Database(C.GoString(db)).Collection(C.GoString(col))
	res, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		return (*C.char)(cgo.JSON(cgo.MAP{
			"count": nil,
			"err":   err.Error(),
		}))
	}
	return (*C.char)(cgo.JSON(cgo.MAP{
		"count": res.DeletedCount,
		"err":   nil,
	}))
}

//export UpdateByID
func UpdateByID(comp *C.component, db, col, id, stu *C.char) *C.char {
	conn := cgo.CastPtr[mongo.Client](comp.ptr)
	var filter any
	err := bson.UnmarshalExtJSON([]byte(C.GoString(stu)), true, &filter)
	if err != nil {
		return (*C.char)(cgo.JSON(cgo.MAP{
			"data": nil,
			"err":  err.Error(),
		}))
	}
	collection := conn.Database(C.GoString(db)).Collection(C.GoString(col))
	res, err := collection.UpdateByID(context.TODO(), id, filter)
	if err != nil {
		return (*C.char)(cgo.JSON(cgo.MAP{
			"data": nil,
			"err":  err.Error(),
		}))
	}
	return (*C.char)(cgo.JSON(cgo.MAP{
		"data": res,
		"err":  nil,
	}))
}

//export UpdateOne
func UpdateOne(comp *C.component, db, col, stu1, stu2 *C.char) *C.char {
	conn := cgo.CastPtr[mongo.Client](comp.ptr)
	var (
		filter any
		update any
	)
	err := bson.UnmarshalExtJSON([]byte(C.GoString(stu1)), true, &filter)
	if err != nil {
		return (*C.char)(cgo.JSON(cgo.MAP{
			"data": nil,
			"err":  err.Error(),
		}))
	}
	err = bson.UnmarshalExtJSON([]byte(C.GoString(stu2)), true, &update)
	if err != nil {
		return (*C.char)(cgo.JSON(cgo.MAP{
			"data": nil,
			"err":  err.Error(),
		}))
	}
	collection := conn.Database(C.GoString(db)).Collection(C.GoString(col))
	res, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return (*C.char)(cgo.JSON(cgo.MAP{
			"data": nil,
			"err":  err.Error(),
		}))
	}
	return (*C.char)(cgo.JSON(cgo.MAP{
		"data": res,
		"err":  nil,
	}))
}

//export UpdateMany
func UpdateMany(comp *C.component, db, col, stu1, stu2 *C.char) *C.char {
	conn := cgo.CastPtr[mongo.Client](comp.ptr)
	var (
		filter any
		update any
	)
	err := bson.UnmarshalExtJSON([]byte(C.GoString(stu1)), true, &filter)
	if err != nil {
		return (*C.char)(cgo.JSON(cgo.MAP{
			"data": nil,
			"err":  err.Error(),
		}))
	}
	err = bson.UnmarshalExtJSON([]byte(C.GoString(stu2)), true, &update)
	if err != nil {
		return (*C.char)(cgo.JSON(cgo.MAP{
			"data": nil,
			"err":  err.Error(),
		}))
	}
	collection := conn.Database(C.GoString(db)).Collection(C.GoString(col))
	res, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return (*C.char)(cgo.JSON(cgo.MAP{
			"data": nil,
			"err":  err.Error(),
		}))
	}
	return (*C.char)(cgo.JSON(cgo.MAP{
		"data": res,
		"err":  nil,
	}))
}

//export ReplaceOne
func ReplaceOne(comp *C.component, db, col, stu1, stu2 *C.char) *C.char {
	conn := cgo.CastPtr[mongo.Client](comp.ptr)
	var (
		filter  any
		replace any
	)
	err := bson.UnmarshalExtJSON([]byte(C.GoString(stu1)), true, &filter)
	if err != nil {
		return (*C.char)(cgo.JSON(cgo.MAP{
			"data": nil,
			"err":  err.Error(),
		}))
	}
	err = bson.UnmarshalExtJSON([]byte(C.GoString(stu2)), true, &replace)
	if err != nil {
		return (*C.char)(cgo.JSON(cgo.MAP{
			"data": nil,
			"err":  err.Error(),
		}))
	}
	collection := conn.Database(C.GoString(db)).Collection(C.GoString(col))
	res, err := collection.ReplaceOne(context.TODO(), filter, replace)
	if err != nil {
		return (*C.char)(cgo.JSON(cgo.MAP{
			"data": nil,
			"err":  err.Error(),
		}))
	}
	return (*C.char)(cgo.JSON(cgo.MAP{
		"data": res,
		"err":  nil,
	}))
}

//export FindOneAndUpdate
func FindOneAndUpdate(comp *C.component, db, col, stu1, stu2 *C.char) *C.char {
	conn := cgo.CastPtr[mongo.Client](comp.ptr)
	var (
		filter  any
		replace any
	)
	err := bson.UnmarshalExtJSON([]byte(C.GoString(stu1)), true, &filter)
	if err != nil {
		return (*C.char)(cgo.JSON(cgo.MAP{
			"data": nil,
			"err":  err.Error(),
		}))
	}
	err = bson.UnmarshalExtJSON([]byte(C.GoString(stu2)), true, &replace)
	if err != nil {
		return (*C.char)(cgo.JSON(cgo.MAP{
			"data": nil,
			"err":  err.Error(),
		}))
	}
	collection := conn.Database(C.GoString(db)).Collection(C.GoString(col))
	res := collection.FindOneAndUpdate(context.TODO(), filter, replace)
	data, err := bson.MarshalExtJSON(res, false, false)
	if err != nil {
		return (*C.char)(cgo.JSON(cgo.MAP{
			"data": nil,
			"err":  err.Error(),
		}))
	}
	return C.CString(fmt.Sprintf(`{"data": [%s], "err": null}`, data))
}

//export FindOneAndReplace
func FindOneAndReplace(comp *C.component, db, col, stu1, stu2 *C.char) *C.char {
	conn := cgo.CastPtr[mongo.Client](comp.ptr)
	var (
		filter  any
		replace any
	)
	err := bson.UnmarshalExtJSON([]byte(C.GoString(stu1)), true, &filter)
	if err != nil {
		return (*C.char)(cgo.JSON(cgo.MAP{
			"data": nil,
			"err":  err.Error(),
		}))
	}
	err = bson.UnmarshalExtJSON([]byte(C.GoString(stu2)), true, &replace)
	if err != nil {
		return (*C.char)(cgo.JSON(cgo.MAP{
			"data": nil,
			"err":  err.Error(),
		}))
	}
	collection := conn.Database(C.GoString(db)).Collection(C.GoString(col))
	res := collection.FindOneAndReplace(context.TODO(), filter, replace)
	data, err := bson.MarshalExtJSON(res, false, false)
	if err != nil {
		return (*C.char)(cgo.JSON(cgo.MAP{
			"data": nil,
			"err":  err.Error(),
		}))
	}
	return C.CString(fmt.Sprintf(`{"data": [%s], "err": null}`, data))
}

func main() {}
