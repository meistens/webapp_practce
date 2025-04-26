package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
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
	// new cmd flag for mysql dsn string
	dsn := flag.String("dsn", "root:password012@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// db here
	db, err := openDB(*dsn)
	if err != nil {
		errLog.Fatal(err)
	}

	// defer call to db.Close() so pool closes before the main() exits
	defer db.Close()

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
	// to get rid of no new vars left side of :=, remove :
	err = srv.ListenAndServe()
	errLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
