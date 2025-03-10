package addresses

import "strings"

func TypeOfAddress(address string) bool {
	validTypes := []string{"Street", "Address"}

	firstWord := strings.Split(address, " ")[0]
	addressHasValidType := false
	for _, validType := range validTypes {
		if firstWord == validType {
			addressHasValidType = true
			break
		}
	}

	return addressHasValidType
}
