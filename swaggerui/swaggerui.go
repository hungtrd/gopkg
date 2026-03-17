package swaggerui

import (
	"embed"
	"io/fs"
	"net/http"
	"strings"
	"text/template"
)

//go:embed assets
var assetsFS embed.FS

// Options configures the Swagger UI handler.
type Options struct {
	Title           string // HTML page title; default "Swagger UI"
	BasePath        string // Mount prefix, e.g. "/swagger-ui"; default "/swagger-ui"
	DocExpansion    string // "none" | "list" | "full"; default "none"
	TryItOutEnabled bool   // enable try-it-out buttons; default false
}

// Option is a functional option for configuring the Swagger UI handler.
type Option func(*Options)

// WithTitle sets the HTML page title.
func WithTitle(title string) Option {
	return func(o *Options) { o.Title = title }
}

// WithBasePath sets the mount prefix for the Swagger UI.
func WithBasePath(path string) Option {
	return func(o *Options) { o.BasePath = path }
}

// WithDocExpansion sets the default expansion mode ("none", "list", "full").
func WithDocExpansion(mode string) Option {
	return func(o *Options) { o.DocExpansion = mode }
}

// WithTryItOut enables the try-it-out buttons.
func WithTryItOut() Option {
	return func(o *Options) { o.TryItOutEnabled = true }
}

func defaultOptions() *Options {
	return &Options{
		Title:        "Swagger UI",
		BasePath:     "/swagger-ui",
		DocExpansion: "none",
	}
}

var indexTmpl = template.Must(template.ParseFS(assetsFS, "assets/index.html"))

var initializerTmpl = template.Must(template.New("initializer").Parse(`window.onload = function () {
    window.ui = SwaggerUIBundle({
        url: "{{.SpecURL}}",
        dom_id: '#swagger-ui',
        deepLinking: true,
        validatorUrl: null,
        supportedSubmitMethods: {{if .TryItOutEnabled}}["get", "put", "post", "delete", "options", "head", "patch", "trace"]{{else}}[]{{end}},
        docExpansion: "{{.DocExpansion}}"
    });
};
`))

type indexData struct {
	Title string
}

type initializerData struct {
	SpecURL         string
	DocExpansion    string
	TryItOutEnabled bool
}

// Handler returns an http.Handler serving Swagger UI.
// specURL is the URL to the OpenAPI spec (relative or absolute).
func Handler(specURL string, opts ...Option) http.Handler {
	cfg := defaultOptions()
	for _, o := range opts {
		o(cfg)
	}

	sub, err := fs.Sub(assetsFS, "assets")
	if err != nil {
		panic("swaggerui: failed to sub assets FS: " + err.Error())
	}
	fileServer := http.FileServer(http.FS(sub))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/")

		switch path {
		case "swagger-initializer.js":
			w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
			_ = initializerTmpl.Execute(w, initializerData{
				SpecURL:         specURL,
				DocExpansion:    cfg.DocExpansion,
				TryItOutEnabled: cfg.TryItOutEnabled,
			})
		case "", "index.html":
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			_ = indexTmpl.Execute(w, indexData{Title: cfg.Title})
		default:
			fileServer.ServeHTTP(w, r)
		}
	})
}

// Register mounts Swagger UI onto mux.
// Registers: GET {basePath}  → 301 → {basePath}/
//
//	GET {basePath}/ → serve UI
func Register(mux *http.ServeMux, specURL string, opts ...Option) {
	cfg := defaultOptions()
	for _, o := range opts {
		o(cfg)
	}

	h := Handler(specURL, opts...)

	mux.HandleFunc("GET "+cfg.BasePath, func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, cfg.BasePath+"/", http.StatusMovedPermanently)
	})
	mux.Handle(cfg.BasePath+"/", http.StripPrefix(cfg.BasePath+"/", h))
}
