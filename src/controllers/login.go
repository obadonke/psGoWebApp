package controllers

import (
	"controllers/util"
	"models"
	"net/http"
	"text/template"
	"viewmodels"
)

type loginController struct {
	template *template.Template
}

func (this *loginController) login(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	if req.Method == "POST" {
		email := req.FormValue("email")
		password := req.FormValue("password")

		member, err := models.GetMember(email, password)
		if err != nil {
			responseWriter.Write([]byte(err.Error()))
			return
		}
		
		session, err := models.CreateSession(member)
		if err != nil {
			responseWriter.Write([]byte(err.Error()))
			return			
		}
		
		var cookie http.Cookie
		cookie.Name = "sessionId"
		cookie.Value = session.SessionId()
		responseWriter.Header().Add("Set-Cookie", cookie.String())
	}
	vm := viewmodels.GetLogin()
	this.template.Execute(responseWriter, vm)
}
