package code

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1, v2 := strings.Split(version1, "."), strings.Split(version2, ".")
	len1, len2 := len(v1), len(v2)
	for i := 0; i < max(len1, len2); i++ {
		var i1, i2 int
		if i < len1 {
			i1, _ = strconv.Atoi(v1[i])
		}
		if i < len2 {
			i2, _ = strconv.Atoi(v2[i])
		}
		if i1 < i2 {
			return -1
		} else if i1 > i2 {
			return 1
		}
	}
	return 0
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
