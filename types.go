package genie

// rootPath : working directory for the project
// folderNames : list of folders to be created
type initPaths struct {
	rootPath    string
	folderNames []string
}

// name: name fo cookie
// lifetime: lifetime of cookie
//persistent: if true, cookie will be persistent between browser closes
// secure : if true , cookie will be encrypted
// domain: domain cookie is associated with
type cookieConfig struct {
	name     string
	lifetime string
	persist  string
	secure   string
	domain   string
}
