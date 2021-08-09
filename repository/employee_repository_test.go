package repository
import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/subodhqss/go-training/models"
)
func TestSaveEmployee1(t *testing.T){
	repo := NewEmpRepo()
	expRes := &models.Employee{
		EmployeeNumber: 109,
		LastName:      "kaur"
		FirstName:     "Neet"
		Extension:     "abcc"
		Email:         "neet@gmail.com"
		ReportsToId:   112,
		JobTitle:      "trainee"
	}
	res := repo.SaveEmployee1(models.Employee{})
	assert.Equal(t,expRes,res)
}