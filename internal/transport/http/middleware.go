package http

import (
	"context"
	"net/http"
	"time"

	logger "ArtAPI/util/logging"

	"github.com/rs/xid"
	"github.com/rs/zerolog"
)

type reqIDKey struct{}
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

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

		reqID := xid.New().String()
		ctx := context.WithValue(r.Context(),reqIDKey{}, reqID)
		r = r.WithContext(ctx)
		l.UpdateContext(func(c zerolog.Context) zerolog.Context {
			return c.Str("request_id", reqID)
		})
		w.Header().Add("X-Request-ID", reqID)

		r = r.WithContext(l.WithContext(r.Context())) //add logger to contxt 
		lrw := newLoggingResponseWriter(w)
		//defer log in order to still get entry when panic occurs down stream
		defer func () { 
			panicVal := recover()
			if panicVal != nil {
				lrw.statusCode = http.StatusInternalServerError
				panic(panicVal)
			}

			l.
				Info().
				Str("method", r.Method).
				Str("url", r.URL.RequestURI()).
				Str("user_agent", r.UserAgent()). 
				Dur("elapsed_ms", time.Since(start)).
				Int("status_code", lrw.statusCode).
				Msg("incoming request")
		}()

		next.ServeHTTP(lrw, r)
	})
}

func TimeoutMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
		defer cancel()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}