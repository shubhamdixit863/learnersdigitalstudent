package services

import (
	"html/template"
	"net/http"
)

const (
	htmlPath            = "../internal/static/index.html"
	get                 = "GET"
	post                = "POST"
	put                 = "PUT"
	deletes             = "DELETE"
	idString            = "id"
	equalString         = "="
	userRegistered      = "User Registered:"
	userDeleted         = "User deleted:"
	userNotFound        = "User not found"
	userUpdated         = "User Updated:"
	bodyNotSupported    = "Request body not supported"
	methodNotAllowed    = 405
	InternalServerError = 500
)

func HtmlRead(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(htmlPath)
	if err != nil {
		w.WriteHeader(InternalServerError)
	}
	tmpl.Execute(w, nil)
	return
}
