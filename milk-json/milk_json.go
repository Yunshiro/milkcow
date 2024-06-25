package milk_json

import (
	"encoding/json"
	"fmt"
)

//MilkJson makes the data to json style

func BeautifulJson(d interface{}) (string, error) {
	// Marshal the data to a json
	b, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", fmt.Errorf("json marshal fail\n")
	}
	return string(b), nil
}
