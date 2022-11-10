package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	IsMarried  bool
}

func TestJSONObject(t *testing.T) {
	bytes, _ := json.Marshal(Customer{
		FirstName:  "Azie",
		MiddleName: "Melza",
		LastName:   "Pratama",
		Age:        19,
		IsMarried:  false,
	})

	fmt.Println(string(bytes))
}
