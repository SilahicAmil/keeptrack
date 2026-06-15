package helpers

import "strings"

func ContainsAll(main string, subs ...string) bool {
	for _, sub := range subs {
		if !strings.Contains(main, sub) {
			return false
		}
	}
	return true
}
