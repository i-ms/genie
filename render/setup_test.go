package render

import (
	"github.com/CloudyKit/jet/v6"
	"os"
	"testing"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./testdata/views"),
	jet.InDevelopmentMode(),
)

var testRenderer = Render{
	Renderer: "",
	RootPath: "",
	JetViews: views,
}

// Entry point for test execution.
// Runs all the functions written in it and then exit's
func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
