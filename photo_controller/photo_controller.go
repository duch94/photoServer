package photo_controller

import (
	"bytes"
	"errors"
	"github.com/duch94/awesomeProject/clients/sqlite"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"io"
	"mime/multipart"
	"os"
)

func SavePhoto(photo multipart.File, photoFilename string) error {
	dirname := "./photos/"
	destination, err := os.Create(dirname + photoFilename)
	if err != nil {
		return errors.New("could not save file to directory " + dirname + " because it does not exists")
	}
	defer destination.Close()
	if _, err := io.Copy(destination, photo); err != nil {
		return err
	}
	err = CreatePreview(photo, dirname + "preview_" + photoFilename)
	if err != nil {
		return err
	}
	err = sqlite.AddPhotoPath()
	if err != nil {
		return err
	}
	return nil
}

func CreatePreview(photo multipart.File, photoFilepath string) error {
	var binaryPhoto []byte
	_, err := photo.Read(binaryPhoto)
	if err != nil {
		return err
	}
	decodedPhoto, _, err := image.Decode(bytes.NewReader(binaryPhoto))
	if err != nil {
		return err
	}
	newPhoto := resize.Resize(100, 0, decodedPhoto, resize.Lanczos3)
	newBinaryPhotoWriter, err := os.Create(photoFilepath)
	if err != nil {
		return err
	}
	defer newBinaryPhotoWriter.Close()
	err = jpeg.Encode(newBinaryPhotoWriter, newPhoto, nil)
	return nil
}

func LoadPhoto() {

}

func DeletePhoto() {

}