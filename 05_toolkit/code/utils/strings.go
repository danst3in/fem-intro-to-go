package utils

import (
	"fmt"
	"strings"
)

func MakeExcited(phrase string) string {
	upperString := strings.ToUpper(phrase) + "!"
	fmt.Println(upperString)
	return upperString
}
