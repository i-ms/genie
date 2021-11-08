package genie

import "net/http"

// SessionLoad : provides middleware to load and save
// session data from current request and communicate
// session token to and from client in a cookie
func (g *Genie) SessionLoad(next http.Handler) http.Handler {
	g.InfoLog.Println("SessionLoad called")
	return g.Session.LoadAndSave(next)
}
