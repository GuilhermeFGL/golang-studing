package addresses

import "strings"

func TypeOfAddress(address string) bool {
	validTypes := []string{"Street", "Avenue"}

	firstWord := strings.Split(address, " ")[0]
	for _, validType := range validTypes {
		if firstWord == validType {
			return true
		}
	}

	return false
}
