package iris_middleware_auth

import "strings"

/**
* String
 */

func ContainsPrefix(s string, values []string) bool {
	for _, value := range values {
		if strings.HasPrefix(s, value) {
			return true
		}
	}
	return false
}

/**
* json
 */
func StringArrayContains(array []string, value string) bool {
	for _, str := range array {
		if str == value {
			return true
		}
	}
	return false
}
