package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	jsonReq := `{"id": "12sc3", "name": "Iponk 69 pro max success zucc buck jack supreme", "price":3000000 }`
	jsonBytes := []byte(jsonReq)

	var result map[string]interface{}

	_ = json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}
