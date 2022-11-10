package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJson(t *testing.T) {
	jsonRequest := `{
		"Firstname" : "Azie",
		"Middlename":"Melza",
		"Lastname":"Pratama",
		"Age":12,
		"IsMarried":true
	}`
	jsonBytes := []byte(jsonRequest)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Age)
}
