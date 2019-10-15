package main

import (
	"drink_check/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &handler.ShowUploadPage{})

	doUploadObj := handler.DoUpload{
		DrinkCheckAPI: "https://aip.baidubce.com/rpc/2.0/easydl/v1/retail/drink",
		TokenAPI:      "https://aip.baidubce.com/oauth/2.0/token",
	}
	mux.Handle("/doupload", &doUploadObj)
	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("files/css"))))
	mux.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir("files"))))
	mux.Handle("/upload/", http.StripPrefix("/upload/", http.FileServer(http.Dir("upload"))))

	server := &http.Server{
		Addr:    ":1111",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
