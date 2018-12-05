package linker

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"html/template"
	"linker/utils"
	"net/http"
)

func (l *Linker) urlCreate(w http.ResponseWriter, r *http.Request) {
	urlStruct := &Url{}

	if err := json.NewDecoder(r.Body).Decode(&urlStruct); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("BadRequest")
		return
	}

	urlModel, err := l.Db.AddUrls(urlStruct.Long)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("BadRequest 2" + err.Error())
		return
	}

	urlStruct.Short = utils.MakeUrl(urlModel.ShortUrl)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(urlStruct)
}

func (l *Linker) urlShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sUrl := vars["shorturl"]

	if !(len(sUrl) > 0) {
		json.NewEncoder(w).Encode("BadRequest shorturl")
		return
	}

	lUrl, err := l.Db.FindlongUrl(sUrl)
	if err != nil {
		tmpl := template.Must(template.ParseFiles("static/index.html"))
		tmpl.Execute(w, nil)
		return
	}

	http.Redirect(w, r, lUrl, http.StatusFound)
}

func (l *Linker) urlRoot(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/index.html"))
	tmpl.Execute(w, nil)
}
