package server

import (
	"github.com/rs/cors"
)

func (s *HttpServer) configureCors() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost", "http://localhost:9000"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	handler := c.Handler(s.svr.Handler)
	s.svr.Handler = handler
}
