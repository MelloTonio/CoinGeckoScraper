package Helpers

import (
	"encoding/json"
	"io/ioutil"
	"log"

	Model "github.com/mellotonio/coinfinder/src/Models"
)

func IsLetter(str string) bool {
	for x := 0; x < len(str); x++ {
		ch := str[x]
		if !((ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == ' ' || ch == '.' || ch == '0' || ch == '-') {
			return false
		}
	}
	return true
}

func WriteJSON(data []Model.Coin, name string) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable to create json file")
		return
	}

	_ = ioutil.WriteFile(name, file, 0644)
}
