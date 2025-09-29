package utils

import "github.com/lucsky/cuid"

func GenerateUserAgent() string {
	randomID := cuid.Slug()

	return "GoMaluum/" + "1.0.0" + " (" + randomID + ")"
}
