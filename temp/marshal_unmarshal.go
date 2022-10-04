package main

import (
	"encoding/json"
	"fmt"
)

type Cars struct {
	Id   int    `json:"Id"`
	Name string `json:"name"`
}

func marshalUnmarshal() {
	tempCar2 := Cars{Id: 34, Name: "suzuki"}

	b, err := json.Marshal(tempCar2)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b))

	tempJson := `{"Id":12 , "name":"toyota"}`
	var tempCar Cars
	err = json.Unmarshal([]byte(tempJson), &tempCar)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(tempCar)
}
