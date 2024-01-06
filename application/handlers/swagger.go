package handlers

import (
    "net/http"
)

func SwaggerMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Serve Swagger documentation if the request matches the pattern
        if r.URL.Path == "/swagger/index.html" {
            http.ServeFile(w, r, "docs/index.html")
            return
        }

        // Se não for a rota Swagger, continue para o próximo handler
        next.ServeHTTP(w, r)
    })
}

