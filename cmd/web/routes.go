package main

import (
	"net/http"
	"github.com/justinas/alice"
	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler{
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
		app.notFound(w)
	})

	FileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handler(http.MethodGet,"/static/*filepath",http.StripPrefix("/static",FileServer))

	dynamic := alice.New(app.sessionManager.LoadAndSave)

	router.Handler(http.MethodGet,"/",dynamic.ThenFunc(app.home))
	
	router.Handler(http.MethodGet,"/snippet/view/:id",dynamic.ThenFunc(app.view))
	router.Handler(http.MethodGet,"/snippet/create",dynamic.ThenFunc(app.create))
	router.Handler(http.MethodPost,"/snippet/create",dynamic.ThenFunc(app.createPost))


	router.Handler(http.MethodGet,"/user/login",dynamic.ThenFunc(app.userlogin))
	router.Handler(http.MethodPost,"/user/login",dynamic.ThenFunc(app.userloginPost))
	router.Handler(http.MethodGet,"/user/signup",dynamic.ThenFunc(app.usersignup))
	router.Handler(http.MethodPost,"/user/signup",dynamic.ThenFunc(app.usersignupPost))
	router.Handler(http.MethodPost,"/user/logout",dynamic.ThenFunc(app.userlogout))
	


	standard:= alice.New(app.recoverPanic,app.logRequest,secureHeaders)
	return standard.Then(router)
}