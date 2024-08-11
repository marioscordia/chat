package utils

import "github.com/marioscordia/chat/internal/constants"

// ValidChannelType is a function that validates channel type
func ValidChannelType(t string) bool {
	types := []string{
		constants.ChatTypeDirect,
		constants.ChatTypeGroup,
	}

	for _, r := range types {
		if r == t {
			return true
		}
	}

	return false
}
