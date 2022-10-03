package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type Cars struct {
	Id   int    `json:"Id"`
	Name string `json:"name"`
}

func main() {

	str := "abcd"

	for i := 0; i <= len(str); i++ {
		for j := i + 1; j <= len(str); j++ {
			fmt.Println(str[i:j])
		}
	}

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

func GetRandom() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int()
}
