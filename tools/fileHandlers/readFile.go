package fileHandlers

import (
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(filename string) []byte {
	data, err := os.ReadFile(filename)
	check(err)
	return data
}
