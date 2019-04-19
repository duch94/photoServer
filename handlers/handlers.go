package handlers

import (
	"fmt"
	"github.com/duch94/awesomeProject/clients/sqlite"
	"github.com/duch94/awesomeProject/photo_controller"
	"log"
	"net/http"
)

func UploadPhotoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("upload photo handler")
	err := r.ParseMultipartForm(100000)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	multipartForm := r.MultipartForm
	photos := multipartForm.File["photos"]
	for i := range photos {
		photo, err := photos[i].Open()
		err = photo_controller.SavePhoto(photo, photos[i].Filename)
		if err != nil {
			_ = photo.Close()
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = photo.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func PhotoListHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("[func PhotoListHandler]: ploto list haha")
	sqlite.GetPhotoList()
	_, err := w.Write([]byte("Photo list haha"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DeletePhotoHandler(w http.ResponseWriter, r *http.Request) {
	photo_controller.DeletePhoto()
	sqlite.DeletePhoto()
}

func GetPhotoHandler(w http.ResponseWriter, r *http.Request) {
	photo_controller.LoadPhoto()
}

func GetPreviewHandler(w http.ResponseWriter, r *http.Request) {
	photo_controller.LoadPhoto()
}