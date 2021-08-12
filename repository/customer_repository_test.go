package repository
import (
	"log"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/subodhqss/go-training/models"
)
func TestSaveCustomer(t *testing.T){
	repo := NewCustRepo();
	customerDummy := &models.Customer{
		CustomerNumber: 109,
		CustomerName:  "kaur",
		State:  "Neet",
		City:   "abcc",
		PostalCode:"251201",
		Country:  "India",
	}
	res := repo.SaveCustomer(*customerDummy)
	customerDummy.CustomerNumber = res.CustomerNumber
	assert.Equal(t,customerDummy,res)

	customerDummy1:=models.Customer{
		CustomerNumber: 999,
	}

	errCheck := repo.SaveCustomer(customerDummy1)
	assert.Nil(t,errCheck)

}

func TestGetCustomer(t *testing.T) {
	repo := NewCustRepo()

	var dummyCustomer []*models.Customer
	result := gormDB.Preload("PaymentDetails").Preload("CustomerNumber").Limit(10).Find(&dummyCustomer)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	res := repo.PrintCustomer()
	assert.Equal(t, dummyCustomer, res)

}

func TestGetCustomerId(t *testing.T) {
	repo := NewCustRepo()

	var dummyCustomer *models.Customer
	result := gormDB.Preload("PaymentDetails").Preload("CustomerNumber").Where("customerNumber", 99).Find(&dummyCustomer)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}

	res := repo.PrintCustomerId("99")
	assert.Equal(t, dummyCustomer, res)

	errCheck := repo.PrintCustomerId("jsdasj") //its not working
	assert.NotNil(t, errCheck)
}

func TestUpdateCustomer(t *testing.T) {
	repo := NewCustRepo()
	customerDummy := &models.Customer{
		CustomerNumber: 109,
		CustomerName:  "kaur",
		State:  "UP",
		City:   "abcc",
		PostalCode:"251201",
		Country:  "India",
		
	}

	res := repo.UpdateCustomer(*customerDummy)
	assert.Equal(t, customerDummy, res)
	errModel := models.Customer{
		CustomerNumber: 109,
		CustomerName:  "kaur",
		State:  "UP",
		City:   "abcc",
		PostalCode:"251201",
		Country:  "India",
		
	}

	errCheck := repo.UpdateCustomer(errModel)
	assert.NotNil(t,errCheck)//not working

}

/*func TestDeleteC(t *testing.T){
	repo := NewCustRepo()

	dummyData:=models.Employee{}

	res:=repo.DeleteCustomer(dummyData,"836674")
	// expResult := gormDB.Where("employeeNumber",836674).Delete(&dummyData)
	assert.Equal(t,res,&dummyData)
}*/

