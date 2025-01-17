/*
Yoink: https://github.com/pilcrowOnPaper/manaba-assignments/blob/main/http.go
*/

package utils

import "strings"

func ParseSetCookieHeaders(headers []string) map[string]string {
	values := map[string]string{}
	for _, header := range headers {
		pair := strings.Split(strings.Split(header, ";")[0], "=")
		if len(pair) == 1 {
			values[pair[0]] = ""
		} else {
			values[pair[0]] = pair[1]
		}
	}
	return values
}
