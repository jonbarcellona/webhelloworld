/*
In your web browser, enter the following address:
http://localhost:8080
*/
package main

import (
	"context"
	"flag"
	"fmt"
	"go-github/github"
	"log"
	"net/http"
	"oauth2"
	"os"
)

func main() {
	portNumber := flag.Int("port", 8080, "port number to listen on")
	flag.Parse()

	http.HandleFunc("/gcp", handleGCP)
	http.HandleFunc("/github", handleGithub)

	http.HandleFunc("/", handle)
	http.HandleFunc("/health", healthCheckHandler)

	log.Printf("Go simple web server (pid=%d)", os.Getpid())
	log.Printf("Listening on port %d", *portNumber)
	portString := fmt.Sprintf(":%d", *portNumber)
	log.Fatal(http.ListenAndServe(portString, nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello world!")
}
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Environment: PID=%d", os.Getpid())
}

func handleGCP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "GCP")

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "... your access token ..."},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(nil)

	// list all organizations for user "willnorris"
	orgs, _, err := client.Organizations.List(ctx, "jonbarcellona", nil)
}

func handleGithub(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "handleGithub")
}
