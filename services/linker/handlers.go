package linker

import (
	"encoding/json"
	"fmt"
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
		json.NewEncoder(w).Encode("BadRequest 2")
		return
	}

	urlStruct.Short = utils.MakeUrl(urlModel.ShortUrl)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(urlStruct)
}

func (l *Linker) urlShow(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UrlShow")
	return
}

func (l *Linker) urlRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UrlRoot")
	return
}
