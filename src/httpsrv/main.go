/*
* WARNING: this code has a lot of comments because it is a test app with no
* goal that justifies having a "good code" written. Actually, it is expected
* to find clunky code. I promise I will write a more streamline code in the
* future, but don't expect that top happen on this app.
*
* :-)
*
* Note: I'm so sorry you came to this repo after good practices and a
* well written code. But I had no time for "clean code" and the few miutes
* I had left I spent writing this long comment to justify my laziness.
 */
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

const fallbackport string = "8080"

var tpl *template.Template

// main type for test handler interface
type hotdog int

// HTTP handler
func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()

	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

// On App initialization parses template
func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

// main function - entry point
// 1. retrieves env vars
// 2. starts http listener
func main() {
	port := getPort()

	var d hotdog
	http.ListenAndServe(port, d)
}

// Returns TCP port
func getPort() string {
	value := getenv("PORT", fallbackport)
	return fmt.Sprintf(":%s", value)
}

// Reads env var and returns fallback value if it is not set
func getenv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
