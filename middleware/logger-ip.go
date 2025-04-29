package middleware

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"
)

// LoggingIPMiddleware é um middleware no estilo net/http
func LoggingIPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := clientIP(r)
		log.Printf("[net/http middleware] IP do cliente: %s", ip)
		next.ServeHTTP(w, r)
	})
}

// clientIP tenta extrair IP real (leva em conta X-Forwarded-For)
func clientIP(r *http.Request) string {
	if xf := r.Header.Get("X-Forwarded-For"); xf != "" {
		parts := strings.Split(xf, ",")
		return strings.TrimSpace(parts[0])
	}
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ip
}

// GetClientIP returns the client's IP address using a simplified approach
// that handles common proxy headers in a secure manner
func GetClientIP(r *http.Request) string {
	// Try standard headers used by proxies in order of reliability
	for _, header := range []string{"X-Real-IP", "X-Forwarded-For"} {
		if ip := r.Header.Get(header); ip != "" {
			// For X-Forwarded-For, only use the first IP in the list
			if header == "X-Forwarded-For" {
				parts := strings.Split(ip, ",")
				return strings.TrimSpace(parts[0])
			}
			return ip
		}
	}

	// Fall back to RemoteAddr if no proxy headers are present
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		// If we can't split, just return RemoteAddr as is
		return r.RemoteAddr
	}
	return ip
}

func RecovererGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if rvr := recover(); rvr != nil {
				if rvr == http.ErrAbortHandler {
					// we don't recover http.ErrAbortHandler so the response
					// to the client is aborted, this should not be logged
					panic(rvr)
				}

				// Log the panic
				c.Error(fmt.Errorf("panic recovered: %v", rvr))
				debug.PrintStack()

				// Respond with internal server error if not websocket
				if c.GetHeader("Connection") != "Upgrade" {
					c.AbortWithStatus(http.StatusInternalServerError)
				}
			}
		}()

		c.Next()
	}
}
