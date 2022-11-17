package utils

import "fmt"

func CreatePrintError(tip string) func(err error) {
	return func(err error) {
		fmt.Printf(`go-fs-extra[%v]: %v \n`, tip, err)
	}
}
