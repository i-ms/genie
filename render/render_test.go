package render

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestData for Table Testing
// name-> name of test
// renderer-> renderer used
// template-> name of template file used
// errorExpected-> true if error is expected
// errorMessage-> error message if error is expected
var pageData = []struct {
	name          string
	renderer      string
	template      string
	errorExpected bool
	errorMessage  string
}{
	{"go_page", "go", "home", false, "error rendering go template"},
	{"go_page_no_template", "go", "no-file", true, "no error rendering non-existent go template, when one is expected"},
	{"jet_page", "jet", "home", false, "error rendering jet template"},
	{"jet_page_no_template", "jet", "no-file", true, "no error rendering non-existent jet template, when one is expected"},
	{"invalid_render_engine", "foo", "home", true, "no error rendering with non-existent template engine"},
}

// Testing rendering of templates created using go template and jet templates
func TestRender_Page(t *testing.T) {

	// Sample request used for test
	r, err := http.NewRequest("GET", "/test-url", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()

	// Testing for all specified conditions in pageData
	for _, e := range pageData {
		testRenderer.Renderer = e.renderer
		testRenderer.RootPath = "./testdata"

		err = testRenderer.Page(w, r, e.template, nil, nil)
		if e.errorExpected {
			if err == nil {
				t.Errorf("%s: %s", e.name, e.errorMessage)
			}
		} else {
			if err != nil {
				t.Errorf("%s: %s: %s", e.name, e.errorMessage, err.Error())
			}
		}
	}
}
