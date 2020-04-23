package code

import "strings"

func simplifyPath(path string) string {
	paths := strings.Split(path, "/")
	var result []string
	for _, p := range paths {
		switch p {
		case "", ".":
		case "..":
			if len(result) != 0 {
				result = result[:len(result)-1]
			} else {
				result = nil
			}
		default:
			result = append(result, p)
		}
	}
	return "/" + strings.Join(result, "/")
}
