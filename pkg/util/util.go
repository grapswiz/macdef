package util

import (
	"os"
)

func ExistsFile(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
