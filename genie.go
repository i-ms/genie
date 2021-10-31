package genie

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/i-ms/genie/render"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const version = "1.0.0"

// Genie is the overall type for Genie package. Members that
// are exported in this type are available to any application
// that uses it
type Genie struct {
	AppName  string
	Debug    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	RootPath string
	Routes   *chi.Mux
	Render   *render.Render
	config   config
}

type config struct {
	port     string
	renderer string
}

// New reads the .env file, creates our application config, populates the Genie
// type with settings based on .env values, and creates necessary folders and
// files if they don't exist
func (g *Genie) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middleware"},
	}
	err := g.Init(pathConfig)
	if err != nil {
		return err
	}

	err = g.checkDotEnv(rootPath)
	if err != nil {
		return err
	}

	// read .env file
	err = godotenv.Load(rootPath + "/.env")
	if err != nil {
		return err
	}

	// create loggers
	infoLog, errorLog := g.startLoggers()
	g.InfoLog = infoLog
	g.ErrorLog = errorLog
	// Setting debug mode
	g.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	g.Version = version
	g.RootPath = rootPath
	// Casting http.Handler to *chi.Mux
	g.Routes = g.routes().(*chi.Mux)

	g.config = config{
		port:     os.Getenv("PORT"),
		renderer: os.Getenv("RENDERER"),
	}

	g.createRenderer()

	return nil
}

// Init creates necessary folders for our Genie application
func (g *Genie) Init(p initPaths) error {
	root := p.rootPath
	for _, path := range p.folderNames {
		// Create folder if it doesn't exist
		err := g.CreateDirIfNotExists(root + "/" + path)
		if err != nil {
			return err
		}
	}
	return nil
}

// ListenAndServe start's the web server
func (g *Genie) ListenAndServe() {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		ErrorLog:     g.ErrorLog,
		Handler:      g.Routes,
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 300 * time.Second,
	}

	g.InfoLog.Printf("Listening on port : %s", os.Getenv("PORT"))
	err := srv.ListenAndServe()
	g.ErrorLog.Fatal(err)
}

// Creates a new .env file if not present
func (g *Genie) checkDotEnv(path string) error {
	err := g.CreateFileIfNotExists(fmt.Sprintf("%s/.env", path))
	if err != nil {
		return err
	}
	return nil
}

// Creates and return's pointer to infoLogger and errorLogger
func (g *Genie) startLoggers() (*log.Logger, *log.Logger) {
	var infoLog *log.Logger
	var errorLog *log.Logger

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return infoLog, errorLog
}

func (g *Genie) createRenderer() {
	myRenderer := render.Render{
		Renderer: g.config.renderer,
		RootPath: g.RootPath,
		Port:     g.config.port,
	}
	g.Render = &myRenderer
}
