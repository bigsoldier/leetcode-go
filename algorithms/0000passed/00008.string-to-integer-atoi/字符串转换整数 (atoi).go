package code

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	regx := regexp.MustCompile(`^[+-]?\d*`)
	ret := regx.FindString(str)
	number, _ := strconv.Atoi(ret)
	if number > math.MaxInt32 {
		return math.MaxInt32
	} else if number < math.MinInt32 {
		return math.MinInt32
	} else {
		return number
	}
}
