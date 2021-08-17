package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"strconv"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("myFile")
	if err != nil {       
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)
	tempFile, _ := ioutil.TempFile("temp-images", "upload-*.jpg")

	t := time.Now().Unix()
	s := strconv.FormatInt(t, 10)
	imagePath:= "./image/" + s + ".jpg"

	fileBytes, err := ioutil.ReadAll(file)
	err1 := ioutil.WriteFile(imagePath, fileBytes, 0755)

	if err1 != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	tempFile.Write(fileBytes)

	fmt.Fprintf(w, "Successfully Uploaded File\n")
	

}
	
