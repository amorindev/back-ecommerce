package middlewares

import (
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Creamos un ResponseWriter personalizado para capturar el status code
		rw := &responseWriter{w, http.StatusOK}

		// Ejecutamos el handler
		next.ServeHTTP(rw, r)

		// Configuramos logrus para salida con colores
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		})
		logrus.SetOutput(os.Stdout)

		// Construimos el mensaje de log
		latency := time.Since(start)
		statusCode := rw.status
		method := r.Method
		path := r.URL.Path

		// Mapeamos colores según el status code
		var statusColor string
		switch {
		case statusCode >= 200 && statusCode < 300:
			statusColor = "\033[32m" // Verde
		case statusCode >= 300 && statusCode < 400:
			statusColor = "\033[36m" // Cian
		case statusCode >= 400 && statusCode < 500:
			statusColor = "\033[33m" // Amarillo
		default:
			statusColor = "\033[31m" // Rojo
		}

		// Reset de color después del mensaje
		resetColor := "\033[0m"

		logrus.Infof(
			"%s | %s %3d %s | %13v | %-7s %s",
			time.Now().Format("2006/01/02 - 15:04:05"),
			statusColor, statusCode, resetColor,
			latency,
			method,
			path,
		)
	})
}

// responseWriter personalizado para capturar el status code
type responseWriter struct {
    http.ResponseWriter
    status int
}

func (rw *responseWriter) WriteHeader(code int) {
    rw.status = code
    rw.ResponseWriter.WriteHeader(code)
}