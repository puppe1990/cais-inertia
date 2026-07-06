package handlers

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/puppe1990/cais-inertia/pkg/cais"
	"github.com/puppe1990/cais-inertia/pkg/cais/i18n"
	"github.com/puppe1990/cais-inertia/pkg/cais/meta"
)

func testSite() meta.Site {
	return meta.Site{AppName: "Cais", AppURL: "https://cais.example.com"}
}

func projectRoot(t *testing.T) string {
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

func setupTestRenderer(t *testing.T) *cais.Renderer {
	t.Helper()
	templatesDir := filepath.Join(projectRoot(t), "web", "templates")
	r, err := cais.NewRendererFromDir(templatesDir, i18n.DefaultCatalog())
	if err != nil {
		t.Fatal(err)
	}
	return r
}