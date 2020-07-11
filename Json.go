package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Name    string   `json:"name"` // if provide the space between  like `json: "name"` then it will not consider tag value while marshal
	Age     int      `json:"age"`
	Weight  float32  `json:"weight"`
	Subject Subjects `json:"subject"`
}

type Subjects struct {
	Physics   bool `json:"physics"`
	Maths     bool `json:"maths"`
	Chemistry bool `json:"chemistry"`
}

/* Go requires all exported fields to start with a capitalized letter.
It’s not common to use that style in JSON however. We use the tag to let the parser know where to actually look for the value.
source : https://eager.io/blog/go-and-json/
*/

func main() {
	fmt.Println("Implementing json")
	//-----------------------------DECODING JSON
	var data1 Data
	jdata := []byte(
		`{
		"name":   "Cena",
		"age":    42,
		"weight": 95.8,
		"subject": {"physics":true,"maths":false,"chemistry":true}
		}`)
	json.Unmarshal(jdata, &data1) //decode the json and store
	fmt.Println(data1)

	//------------------------------------ ENCODING JSON
	data := Data{"John", 25, 78.5, Subjects{true, true, true}}
	jd, err := json.Marshal(data) // encode in json format
	fmt.Println(string(jd))
	fmt.Println(err)
	jd1, _ := json.MarshalIndent(data, "", " ") // encode in json format
	fmt.Println(string(jd1))
}
