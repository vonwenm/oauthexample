package main

import (
	"log"
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
)

// The one and only access token! In real-life scenarios, a more complex authentication
// middleware than auth.Basic should be used, obviously.
// const AuthToken = "token"

// The one and only martini instance.
var m *martini.Martini

func init() {
	m = martini.New()
	// Setup middleware
	m.Use(martini.Recovery())
	m.Use(martini.Logger())
	m.Use(render.Renderer(render.Options{Directory: "views", Layout: "layout", Extensions: []string{".tmpl", ".html"}}))
	// Setup routes
	r := martini.NewRouter()
	r.Get(`/`, func(re render.Render) { re.HTML(200, "display", GetAuthCodeUrl()) })
	r.Get(`/oauth2callback`, GetUserInfo)
	// Add the router action
	m.Action(r.Handle)
}

func main() {
	// Listen on https: with the preconfigured martini instance. The certificate files
	// can be created using this command in this repository's root directory:
	//
	// go run /path/to/goroot/src/crypto/tls/generate_cert.go --host="localhost"
	//
	if err := http.ListenAndServeTLS(":8001", "cert.pem", "key.pem", m); err != nil {
		log.Fatal(err)
	}
}
