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

	// Testing rendering of pages using go template
	testRenderer.Renderer = "go"
	testRenderer.RootPath = "./testdata"

	err = testRenderer.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error rendering page : ", err)
	}

	// Testing for non-existent go template
	err = testRenderer.Page(w, r, "no-file", nil, nil)
	if err == nil {
		t.Error("Error rendering non-existent go template: ", err)
	}

	// Testing rendering of pages using jet template
	testRenderer.Renderer = "jet"
	err = testRenderer.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error rendering page : ", err)
	}

	// Testing for non-existent jet template
	err = testRenderer.Page(w, r, "no-file", nil, nil)
	if err == nil {
		t.Error("Error rendering non-existent jet template: ", err)
	}

}
