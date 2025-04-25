package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// dependency injection, might be different if using imported packages
// basically, struct holds app-wide dependencies to be used
type application struct {
	errLog  *log.Logger
	infoLog *log.Logger
}

// copy-paste, check VC for how it came to this
func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app := &application{
		errLog:  errLog,
		infoLog: infoLog,
	}
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errLog,
		// Call the new app.routes() method to get the servemux containing our routes.
		Handler: app.routes(),
	}
	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errLog.Fatal(err)
}
