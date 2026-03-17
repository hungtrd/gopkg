// Package main demonstrates serving Swagger UI alongside an API using the
// standard net/http package.
//
// Run:
//
//	go run ./pkg/swaggerui/examples/stdhttp
//
// Then open http://localhost:8080/swagger-ui/ in your browser.
package main

import (
	"log"
	"net/http"

	"github.com/hungtrd/gopkg/swaggerui"
)

func main() {
	mux := http.NewServeMux()

	// Serve the OpenAPI spec bundled alongside this example.
	mux.HandleFunc("GET /openapi.yaml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./openapi.yaml")
	})

	// Mount Swagger UI at /swagger-ui/.
	// Visiting /swagger-ui redirects to /swagger-ui/ automatically.
	swaggerui.Register(mux, "/openapi.yaml",
		swaggerui.WithTitle("My API – Swagger UI"),
		swaggerui.WithBasePath("/swagger-ui"),
		swaggerui.WithDocExpansion("list"),
		swaggerui.WithTryItOut(),
	)

	// Your API routes go here.
	mux.HandleFunc("GET /api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message":"hello"}`))
	})

	log.Println("listening on http://localhost:8080")
	log.Println("swagger UI → http://localhost:8080/swagger-ui/")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
