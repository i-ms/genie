package render

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Testing rendering of templates created using go template and jet templates
func TestRender_Page(t *testing.T) {
	r, err := http.NewRequest("GET", "/test-url", nil)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()

	testRenderer.Renderer = "go"
	testRenderer.RootPath = "./testdata"

	// Testing rendering of pages using go template
	err = testRenderer.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error rendering page : ", err)
	}
}
