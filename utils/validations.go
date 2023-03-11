package utils

import (
	"regexp"
	"strings"

	"enigmacamp.com/fine_dms/model"
	"enigmacamp.com/fine_dms/repo"
)

type user struct {
	model.User
	userRepo repo.UserRepo
}

func IsValidTag(tagName string) bool {
	if !regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString(tagName) {
		return false
	}

	words := strings.Fields(tagName)
	if len(words) > 16 {
		return false
	}

	return true
}
