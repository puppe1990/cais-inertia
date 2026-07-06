package session

import (
	"testing"

	"github.com/puppe1990/cais-inertia/pkg/cais"
)

func TestCookieOptionsFromConfig_productionSecure(t *testing.T) {
	opts := CookieOptionsFromConfig(cais.Config{Env: "production"})
	if !opts.Secure {
		t.Error("Secure = false, want true")
	}
}
