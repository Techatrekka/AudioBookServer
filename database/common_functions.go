package database

import (
	"server/configs"
	"strings"
)

func SelectById(table string, id string) []byte {
	database := configs.GetDBClient()
	data, _, err1 := database.From(table).Select("*", "", false).Eq(
		strings.ToLower(table)+"_id", id).Execute()
	if err1 != nil {
		panic(err1)
	}
	return data
}

func SelectAll(table string) []byte {
	database := configs.GetDBClient()
	data, _, err1 := database.From(table).Select("*", "", false).Execute()
	if err1 != nil {
		panic(err1)
	}
	return data
}

func UploadObjectToTable(table string, object any) {
	database := configs.GetDBClient()
	_, _, err1 := database.From(table).Upsert(object, "", "", "").Execute()
	if err1 != nil {
		panic(err1)
	}
}

func SelectByCompositeId(
	table string, column1 string,
	column2 string, id1 string, id2 string) []byte {
	database := configs.GetDBClient()
	data, _, err1 := database.From(table).Select("*", "", false).Eq(
		strings.ToLower(column1)+"_id", id1).Eq(
		strings.ToLower(column2)+"_id", id2).Execute()
	if err1 != nil {
		panic(err1)
	}
	return data
}
