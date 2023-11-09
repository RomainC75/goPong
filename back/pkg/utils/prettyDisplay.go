package utils

import (
	"encoding/json"
	"fmt"
	"log"
)

func PrettyDisplay(v interface{}) {
	empJSON, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("----> \n %s \n-----------\n", string(empJSON))
}
