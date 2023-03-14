package main

import (
	"net/http"
	"strconv"
)



func (app *application) quoteCreateShow(w http.ResponseWriter, r *http.Request) {
	// display the input box
	RenderTemplate(w, "poll.create.page.tmpl", nil)
}

func (app *application) quoteCreateSubmit(w http.ResponseWriter, r *http.Request) {
	// add the question to the datastore
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	question := r.PostForm.Get("new_question") // Refers to the name
	_, err = app.questions.Insert(question)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}


