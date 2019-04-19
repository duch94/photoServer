package main

import (
	"github.com/duch94/awesomeProject/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/uploadPhoto", handlers.UploadPhotoHandler).Methods("POST")
	router.HandleFunc("/photoList", handlers.PhotoListHandler).Methods("GET")
	router.HandleFunc("/deletePhoto", handlers.DeletePhotoHandler).Methods("DELETE")
	router.HandleFunc("/getPhoto", handlers.GetPhotoHandler).Methods("GET")
	router.HandleFunc("/getPreview", handlers.GetPreviewHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}