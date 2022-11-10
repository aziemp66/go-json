package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {
	bytes, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestMarshal(t *testing.T) {
	logJson("Eko")
	logJson(12)
	logJson(true)
	logJson([]string{"Azie", "Melza", "Pratama"})
	logJson(map[string]string{"Azie": "Backend", "Fadhil": "Frontend"})
}
