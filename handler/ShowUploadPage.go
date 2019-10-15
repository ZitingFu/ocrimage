package handler

import "net/http"

type ShowUploadPage struct{}

func (s *ShowUploadPage) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	http.Redirect(res, req, "/html", http.StatusFound)
	//res.Write([]byte("this is upload page"))
}
