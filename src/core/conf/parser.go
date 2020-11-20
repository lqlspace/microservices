package conf

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func MustLoad(file string, val interface{}) {
	dataBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("file path error.")
	}

	if err := json.Unmarshal(dataBytes, val); err != nil {
		log.Fatalf("unmarshal err: %s\n", err)
	}
}
