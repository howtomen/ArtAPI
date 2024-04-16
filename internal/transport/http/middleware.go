package http

import (
	"context"
	"net/http"
	"time"

	logger "ArtAPI/util/logging"
)

func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		next.ServeHTTP(w, r)
	})
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		start := time.Now()
		l := logger.GetLogger()
		//defer log in order to still get entry when panic occurs down stream
		defer func () { 
			l.
				Info().
				Str("method", r.Method).
				Str("url", r.URL.RequestURI()).
				Str("user_agent", r.UserAgent()). 
				Dur("elapsed_ms", time.Since(start)).
				Msg("incoming request")
		}()

		next.ServeHTTP(w, r)
	})
}

func TimeoutMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
		defer cancel()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}