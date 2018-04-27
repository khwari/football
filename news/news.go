package news

import "net/http"

func GetNews(w http.ResponseWriter, r *http.Request)  {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("News Page"))
}
