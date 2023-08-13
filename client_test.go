package supabase_test

import (
	"fmt"
	"testing"

	"github.com/supabase-community/supabase-go"
)

const (
	API_URL = "https://jnyalqqtavnsmjgplrzt.supabase.co"
	API_KEY = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImpueWFscXF0YXZuc21qZ3Bscnp0Iiwicm9sZSI6InNlcnZpY2Vfcm9sZSIsImlhdCI6MTY2MjQyODY4NiwiZXhwIjoxOTc4MDA0Njg2fQ.R2TXvWTSy1ExDJ8qBNnjAZx3RSCCgzob2SZ3OOlf3RI"
)

func TestFrom(t *testing.T) {
	client, err := supabase.NewClient(API_URL, API_KEY, nil)
	if err != nil {
		fmt.Println("cannot initalize client", err)
	}
	data, count, err := client.From("countries").Select("*", "exact", false).Execute()
	fmt.Println(string(data), err, count)
}

func TestRpc(t *testing.T) {
	client, err := supabase.NewClient(API_URL, API_KEY, nil)
	if err != nil {
		fmt.Println("cannot initalize client", err)
	}
	result := client.Rpc("hello_world", "", nil)
	fmt.Println(result)
}
