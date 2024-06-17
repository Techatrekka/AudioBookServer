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
