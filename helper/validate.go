package helper

import (
	"fmt"
	"io/ioutil"
)

func ValidateDir(director string) (result bool) {
	result = true
	_, err := ioutil.ReadDir(director)
	if err != nil {
		fmt.Printf("%v", err)
		result = false
	}

	return
}
