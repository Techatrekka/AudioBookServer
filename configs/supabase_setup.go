package configs

import (
	"encoding/json"
	"fmt"

	"github.com/supabase-community/supabase-go"
)

var database *supabase.Client

func ConnectDB() (*supabase.Client, error) {
	client, err := supabase.NewClient("https://vgrqqwqnnhlbknrjzavn.supabase.co",
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6InZncnFxd3FubmhsYmtucmp6YXZuIiwicm9sZSI6InNlcnZpY2Vfcm9sZSIsImlhdCI6MTcxODMxNTIxMywiZXhwIjoyMDMzODkxMjEzfQ.Frm5-0qtNTjATmeafaDX1RXuozyP09sZ7oWzl81p2O4",
		&supabase.ClientOptions{})
	if err != nil {
		fmt.Println("cannot initalize client", err)
	}
	fmt.Println("Connected to Supabase")
	database = client
	// testfunction(client)
	return client, nil
}

func GetDBClient() *supabase.Client {
	return database
}

type Genre struct {
	Genre_id    int    `json:"tag_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func testfunction(database *supabase.Client) {
	var value []Genre
	// value = Genre{id: 1, title: "test", description: ""}
	//database.From("Genre").Insert(value, false, "*", "exact", "1").Execute()
	data, count, err1 := database.From("Tags").Select("*", "", false).Execute()
	if err1 != nil {
		panic(err1)
	}

	// print(data)
	print(count)
	err := json.Unmarshal(data, &value)
	if err != nil {
		print(err)
	}
	print(value[0].Title)
	print(len(value))
}
