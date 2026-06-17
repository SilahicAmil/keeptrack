package helpers

import "strings"

func ContainsAny(main string, subs ...string) bool {
	for _, sub := range subs {
		if strings.Contains(main, sub) {
			return true
		}
	}
	return false
}
