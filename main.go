package main

import (
	"fmt"
	"milkcow/milk-http"
	milk_json "milkcow/milk-json"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	milk_http.GET("/get/", ":9090", SelectData())
}

func SelectData() milk_http.MilkFunc {
	data := []Message{
		{
			Name: "Milk",
			Body: "I love Drink milk!!!",
			Time: 20240623,
		},
		{
			Name: "Meat",
			Body: "Meat is delicious!",
			Time: 20240624,
		},
	}
	dst, err := milk_json.BeautifulJson(data)
	if err != nil {
		fmt.Printf("BeautifulJson Err: %v\n", err)
	}
	return milk_http.MilkTransfer(dst)
}
