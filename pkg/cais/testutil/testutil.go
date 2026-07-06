package testutil

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/puppe1990/cais-inertia/pkg/cais"
)

// ProjectRoot walks up from cwd to find go.mod.
func ProjectRoot(t *testing.T) string {
	t.Helper()
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	for {
		if _, err := os.Stat(filepath.Join(wd, "go.mod")); err == nil {
			return wd
		}
		parent := filepath.Dir(wd)
		if parent == wd {
			t.Fatal("go.mod not found")
		}
		wd = parent
	}
}

// TemplatesDir returns web/templates in scaffolded apps, or pkg/cais/testdata/templates
// in the framework-only repo.
func TemplatesDir(t *testing.T) string {
	t.Helper()
	root := ProjectRoot(t)
	web := filepath.Join(root, "web", "templates")
	if info, err := os.Stat(web); err == nil && info.IsDir() {
		return web
	}
	return filepath.Join(root, "pkg", "cais", "testdata", "templates")
}

// NewRenderer loads templates from TemplatesDir.
func NewRenderer(t *testing.T) *cais.Renderer {
	t.Helper()
	r, err := cais.NewRendererFromDir(TemplatesDir(t), nil)
	if err != nil {
		t.Fatal(err)
	}
	return r
}

type RequestOption func(*http.Request)

// PathValue sets a Go 1.22+ path parameter on the request.
func PathValue(key, value string) RequestOption {
	return func(r *http.Request) {
		r.SetPathValue(key, value)
	}
}

// NewRequest builds an httptest request with optional path values.
func NewRequest(method, target string, opts ...RequestOption) *http.Request {
	req := httptest.NewRequest(method, target, nil)
	for _, opt := range opts {
		opt(req)
	}
	return req
}
