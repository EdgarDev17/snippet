package main

import "net/http"

func secureHandlers(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Security-Policy",
			"default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com")
		w.Header().Set("Referrer-Policy", "origin-when-cross-origin")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "deny")
		w.Header().Set("X-XSS-Protection", "0")

		next.ServeHTTP(w, r)
	})
}

// este middleware tiene que estar vinculado a al struck Application
func (app *Application) logginRequest(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// de esta manera se declaran variables en grupo
		var (
			ip       = r.RemoteAddr
			protocol = r.Proto
			method   = r.Method
			uri      = r.URL.RequestURI()
		)

		app.logger.Info("received request", "ip", ip, "protocol", protocol, "method", method, "uri", uri)
		next.ServeHTTP(w, r)
	})
}
