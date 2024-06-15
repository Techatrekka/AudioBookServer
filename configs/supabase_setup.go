package configs

import (
	"encoding/json"
	"fmt"

	"github.com/supabase-community/supabase-go"
)

func ConnectDB() (*supabase.Client, error) {
	client, err := supabase.NewClient("https://vgrqqwqnnhlbknrjzavn.supabase.co",
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6InZncnFxd3FubmhsYmtucmp6YXZuIiwicm9sZSI6InNlcnZpY2Vfcm9sZSIsImlhdCI6MTcxODMxNTIxMywiZXhwIjoyMDMzODkxMjEzfQ.Frm5-0qtNTjATmeafaDX1RXuozyP09sZ7oWzl81p2O4",
		&supabase.ClientOptions{})
	if err != nil {
		fmt.Println("cannot initalize client", err)
	}
	fmt.Println("Connected to Supabase")
	return client, nil
}

type Genre struct {
	Genre_id    int    `json:"genre_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func testfunction(database supabase.Client) {
	var value []Genre
	// value = Genre{id: 1, title: "test", description: ""}
	//database.From("Genre").Insert(value, false, "*", "exact", "1").Execute()
	data, _, err1 := database.From("Genre").Select("*", "1", false).Execute()
	if err1 != nil {
		panic(err1)
	}

	// print(data)
	// print(count)
	err := json.Unmarshal(data, &value)
	if err != nil {
		print(err)
	}
	print(value[0].Title)
}
