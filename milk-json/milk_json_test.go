package milk_json

import (
	"fmt"
	"testing"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func TestBeautifulJson(t *testing.T) {
	data := []Message{
		{
			Name: "Milk",
			Body: "I love Drink milk!",
			Time: 20240623,
		},
		{
			Name: "Meat",
			Body: "Meat is delicious!",
			Time: 20240624,
		},
	}
	json, err := BeautifulJson(data)
	if err != nil {
		fmt.Println(err, json)
		t.Fatal("err")
	}
	fmt.Println(json)
}

func TestBeautifulJson2(t *testing.T) {
	data := Message{
		Name: "Milk",
		Body: "I love Drink milk!",
		Time: 20240623,
	}
	json, err := BeautifulJson(data)
	if err != nil {
		t.Fatal("err")
	}
	fmt.Println(json)

}
