package repository
import (
	"log"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/subodhqss/go-training/models"
)
func TestSaveOffice(t *testing.T){
	repo := NewOfficeRepo();
	officeDummy := &models.Office{
		OfficeCode: "E99",
		City:       "Noida",
		Phone:      "123456",
		AddressLine1: "sector 63",
		Country:       "India",
		PostalCode:     "189",
		State:       "UP",
		
	}
	res := repo.SaveOffice(*officeDummy)
	officeDummy.OfficeCode = res.OfficeCode

	officeDummy1:=models.Office{
		OfficeCode: "E99",
	}

	errCheck := repo.SaveOffice(officeDummy1)
	assert.Nil(t,errCheck)

}

func TestGetOffice(t *testing.T) {
	repo := NewOfficeRepo()

	var dummyOffice []*models.Office
	result := gormDB.Preload("OfficeDetail").Preload("OfficeCode").Limit(10).Find(&dummyOffice)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	res := repo.PrintOffice()
	assert.Equal(t, dummyOffice, res)

}

func TestGetOfficeId(t *testing.T) {
	repo := NewOfficeRepo()

	var dummyOffice *models.Office
	result := gormDB.Preload("OfficeDetail").Preload("OfficeCode").Where("officeCode", "E99").Find(&dummyOffice)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}

	res := repo.PrintOfficeId("99")
	assert.Equal(t, dummyOffice, res)

	errCheck := repo.PrintOfficeId("jsdasj") //its not working
	assert.NotNil(t, errCheck)
}

func TestUpdateOffice(t *testing.T) {
	repo := NewOfficeRepo()
	officeDummy := &models.Office{
		OfficeCode: "E99",
		City:       "Noida",
		Phone:      "123456",
		AddressLine1: "sector 63",
		Country:       "India",
		PostalCode:     "189",
		State:       "UP",
	}

	res := repo.EditOffice(*officeDummy)
	assert.Equal(t, officeDummy, res)
	errModel := models.Office{
		OfficeCode: "E99",
		City:       "Noida",
		Phone:      "123456",
		AddressLine1: "sector 63",
		Country:       "India",
		PostalCode:     "189",
		State:       "UP",
	}
	errCheck1 := repo.EditOffice(errModel)
	assert.NotNil(t,errCheck1)//not working

}

/*func TestDeleteO(t *testing.T){
	repo := NewOfficeRepo()

	dummyData:=models.Employee{}

	res:=repo.DeleteOffice(dummyData,"836674")
	// expResult := gormDB.Where("employeeNumber",836674).Delete(&dummyData)
	assert.Equal(t,res,&dummyData)
}*/

