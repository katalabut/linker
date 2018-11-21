package linker

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (l *Linker) urlCreate(w http.ResponseWriter, r *http.Request) {
	reqStruct := new(Url)
	// создаем новый декодер, использующий ридер r.Body
	// и читаем данные в структуру
	if err := json.NewDecoder(r.Body).Decode(&reqStruct); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// используем слой базы данных для добавления URL
	l.db.AddUrls()
	//Ls.myconnection.AddUrls(reqBodyStruct.LongUrl, reqBodyStruct.ShortUrl)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reqStruct)
}

func (l *Linker) urlShow(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UrlShow")
	return
}

func (l *Linker) urlRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UrlRoot")
	return
}
