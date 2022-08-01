package config

import "net/http"

var SessionName = "board-session"

func SameSiteMode() http.SameSite {
	if IsProduction() {
		return http.SameSiteNoneMode
	}
	return http.SameSiteLaxMode
}