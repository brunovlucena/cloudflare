package util

import "fmt"

func ParseErr(err error) {
	if err != nil {
		fmt.Println("Error!")
	}
}
