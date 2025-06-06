package main

import (
	"fmt"
	"context"

	"github.com/justinas/nosurf"
	"net/http"
)

func secureHeaders(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Security-Policy","default-src 'self'; style-src 'self' fonts.googleapis.com; font-srcfonts.gstatic.com")
	w.Header().Set("Referrer-Policy", "origin-when-cross-origin")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-Frame-Options", "deny")
	w.Header().Set("X-XSS-Protection", "0")
	
	next.ServeHTTP(w, r)
	})
}

func (app *application) logRequest(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
		app.infolog.Printf("%s - %s %s %s",r.RemoteAddr,r.Proto,r.Method,r.URL.RequestURI())

		next.ServeHTTP(w,r)
	})
}


func (app *application) recoverPanic(next http.Handler) http.Handler{
	fn := func(w http.ResponseWriter,r *http.Request){
		defer func(){
			if err:=recover(); err!=nil{
				w.Header().Set("Connection","Close")
				app.serverError(w,fmt.Errorf("%s",err))
			}
		}()
		next.ServeHTTP(w,r)
	}
	return http.HandlerFunc(fn)

	
}

func (app *application) requireAuthentication(next http.Handler)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !app.IsAuthenticated(r){
			http.Redirect(w,r,"/user/login",http.StatusSeeOther)
			return
		}
		w.Header().Add("Cache-Control","no-store")
		next.ServeHTTP(w,r)
	})
}

func noSrf(next http.Handler) http.Handler{
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: true,
	})
	return csrfHandler
}

func (app *application) authenticate (next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
		id := app.sessionManager.GetInt(r.Context(),"authenticatedUserID")
		if id==0{
			next.ServeHTTP(w,r)
			return
		}
		exists,err:=app.users.Exists(id)
		if err !=nil{
			app.serverError(w,err)
			return
		}
		if exists{
			ctx := context.WithValue(r.Context(),isAuthenticatedContextKey,true)
			r = r.WithContext(ctx)
		}
		next.ServeHTTP(w,r)
	})
}




