package util

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
)

func UploadFile(file multipart.File, header *multipart.FileHeader, filePath string) string {
	imagePath := "./image/" + filePath + ".jpg"

	fileBytes, _ := ioutil.ReadAll(file)
	err := ioutil.WriteFile(imagePath, fileBytes, 0755)

	if err != nil {
		fmt.Println(err)
	}
	
	return imagePath 
}
