package main

import (
	"crypto/tls"
	// "fmt"
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	// "strconv"
	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"
	"muskiteer.net/internal/models"

	"github.com/go-playground/form/v4"

	_ "github.com/go-sql-driver/mysql"
)

type application struct{
	errorlog *log.Logger
	infolog *log.Logger
	snippets *models.SnippetModel
	templateCache map[string]*template.Template
	formDecoder *form.Decoder
	sessionManager *scs.SessionManager
	users *models.UserModel
}



func main(){	
	errorlog := log.New(os.Stdout,"ERROR\t", log.Ldate| log.Ltime|log.Lshortfile)
	infolog := log.New(os.Stdout,"INFO\t", log.Ldate| log.Ltime)
	
	addr := flag.String("addr",":4000", "network port ")
	dsn:= flag.String("dsn","web:123456@/snippetbox?parseTime=true","MySQL data source name")	
	flag.Parse()
	db ,err:= openDB(*dsn)
	if err != nil{
		errorlog.Fatal(err)
	}
		
	defer db.Close()

	templateCache,err := newTemplateCache()
	if err!=nil{
		errorlog.Fatal(err)
	}
	formDecoder := form.NewDecoder()

	sessionManager:=scs.New()
	sessionManager.Store = mysqlstore.New(db)
	sessionManager.Lifetime = 12*time.Hour

	sessionManager.Cookie.Secure = true

	app := &application{
		errorlog: errorlog,
		infolog: infolog,
		snippets: &models.SnippetModel{DB: db},
		templateCache: templateCache,
		formDecoder: formDecoder,
		sessionManager: sessionManager,
		users: &models.UserModel{DB:db},
		}

		tlsConfig := &tls.Config{
			CurvePreferences : []tls.CurveID{tls.X25519,tls.CurveP256},
		}
		
	
	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorlog,
		Handler: app.routes(),
		TLSConfig: tlsConfig,
		IdleTimeout: time.Minute,
		ReadTimeout: 5*time.Second,
		WriteTimeout: 10*time.Second,
		}

	infolog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServeTLS("./tls/cert.pem","./tls/key.pem")
	errorlog.Fatal(err)
}

	func openDB(dsn string) (*sql.DB,error){
		db,err:=sql.Open("mysql",dsn)
		if err!=nil{
			return nil,err
		}
		if err=db.Ping();err!=nil{
			return nil,err
		}
		return db,nil
	}

