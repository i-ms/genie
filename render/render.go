package render

import (
	"errors"
	"fmt"
	"github.com/CloudyKit/jet/v6"
	"github.com/alexedwards/scs/v2"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Render struct {
	Renderer   string
	RootPath   string
	Secure     bool
	Port       string
	ServerName string
	JetViews   *jet.Set
	Session    *scs.SessionManager
}

type TemplateData struct {
	IsAuthenticated bool
	IntMap          map[string]int
	StringMap       map[string]string
	FloatMap        map[string]float64
	Data            map[string]interface{}
	CSRFToken       string
	Port            string
	ServerName      string
	Secure          bool
}

// Adding default template data to each request
func (ren *Render) defaultData(td *TemplateData, r *http.Request) *TemplateData {
	td.Secure = ren.Secure
	td.Port = ren.Port
	td.ServerName = ren.ServerName
	if ren.Session.Exists(r.Context(), "userID") {
		td.IsAuthenticated = true
	}
	return td
}

func (ren *Render) Page(w http.ResponseWriter, r *http.Request, view string, variables, data interface{}) error {

	// Choosing the render function according to template specified in env file
	switch strings.ToLower(ren.Renderer) {
	case "go":
		return ren.GoPage(w, r, view, data)

	case "jet":
		return ren.JetPage(w, r, view, variables, data)

	default:

	}

	return errors.New("no rendering engine specified")
}

// GoPage renders a standard Go template
func (ren *Render) GoPage(w http.ResponseWriter, r *http.Request, view string, data interface{}) error {
	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/views/%s.page.tmpl", ren.RootPath, view))
	if err != nil {
		return err
	}

	// If data is present then casting it to TemplateData
	td := &TemplateData{}
	if data != nil {
		td = data.(*TemplateData)
	}

	err = tmpl.Execute(w, &td)
	if err != nil {
		return err
	}

	return nil
}

// JetPage renders a template using the Jet Templating engine
func (ren *Render) JetPage(w http.ResponseWriter, r *http.Request, templateName string, variables, data interface{}) error {

	// DataStructure used by Jet to pass data into template
	var vars jet.VarMap

	//
	if variables == nil {
		vars = make(jet.VarMap)
	} else {
		vars = variables.(jet.VarMap)
	}

	td := &TemplateData{}
	if data != nil {
		td = data.(*TemplateData)
	}

	// Adding default data to the template data
	td = ren.defaultData(td, r)

	// Loading jet template
	t, err := ren.JetViews.GetTemplate(fmt.Sprintf("%s.jet", templateName))
	if err != nil {
		log.Println(err)
		return err
	}

	// Execute template in w
	if err = t.Execute(w, vars, td); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
