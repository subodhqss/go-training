package controllers
import (
	"encoding/json"
	"net/http"
	"github.com/subodhqss/go-training/repository"
	"github.com/subodhqss/go-training/services"
	
)

var officeService = service.NewOfficeService(repository.NewOfficeRepo())

func GetOffice(rw http.ResponseWriter, r *http.Request){
	data := officeService.PrintOffice()
	jsonData, _ := json.Marshal(data)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(jsonData)

}