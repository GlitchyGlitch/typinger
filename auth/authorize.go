package auth

import "github.com/GlitchyGlitch/typinger/models"

func Authorize(user *models.User) bool { // TODO: Add groups later.
	if user == nil {
		return false
	}
	return true
}
