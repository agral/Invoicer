package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

// A custom handler func. Just to try it out.
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})
}

// Adds CSRF protection to all POST requests.
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.IsProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// Loads and saves the session on every request.
func SessionLoad(next http.Handler) http.Handler {
	return sessionManager.LoadAndSave(next)
}
