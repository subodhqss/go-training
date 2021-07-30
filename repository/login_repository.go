// package repository

// import (
// 	"fmt"
// 	"log"

// 	"github.com/subodhqss/go-training/models"
// )

// type LoginReposiotry interface {
// 	//PrintLogin() []*models.Login
// 	//GetEmployeEmail(models.Login) *models.Login
// }

// func NewLogRepo() LoginReposiotry {
// 	return &logRepo{}

//  }

// type logRepo struct{}

// // func (er *logRepo) PrintLogin() []*models.Employee {

// // 	var login []*models.Employee
// // 	result := gormDB.Limit(10).Find(&login)
// // 	if err := result.Error; err != nil {
// // 		log.Print("Error in getting all records")
// // 	}

// // 	return login
// // }

// // func (tr *logRepo) SaveLogin(login models.Employee) *models.Employee {
// // 	b := gormDB.First(&login)
// // 	result := gormDB.Save(&b)
// // 	if err := result.Error; err != nil {
// // 		log.Print("Error in getting all records")
// // 	}
// // 	fmt.Println("created", login)
// // 	return &login

// // }


// // func (tr *logRepo) GetEmployeEmail(login models.Login) *models.Login {
// // 	b := gormDB.First(&login)
// // 	result := gormDB.Save(&b)
// // 	if err := result.Error; err != nil {
// // 		log.Print("Error in getting all records")
// // 	}
// // 	fmt.Println("created", login)
// // 	return &login

// // }



// func (lr *logRepo) GetEmployeByEmail(email string) *models.Employee {
// 	var employee *models.Employee

// 	fmt.Println(email) 

// 	result := gormDB.Where("email", email).Find(&employee)
// 	if err := result.Error; err != nil {
// 		log.Print("Error in getting all records")
// 	}
// 	fmt.Println("there is no error in get method")
//     return employee
// }
package repository

import (
	"fmt"
	"log"

	"github.com/subodhqss/go-training/models"
)

type LoginReposiotry interface {
	//PrintLogin() []*models.Employee
	//SaveLogin(models.Employee) *models.Employee
	GetEmployeEmail(email string) *models.Employee
}

func NewLogRepo() LoginReposiotry {
	return &logRepo{}

}

type logRepo struct{}

func (lr *logRepo) GetEmployeEmail(email string) *models.Employee {
	var employee *models.Employee

	fmt.Println(email)

	result := gormDB.Preload("OfficeDetail").Preload("ReportsTo").Where("email", email).Find(&employee)
	if err := result.Error; err != nil {
		log.Print("Error in getting all records")
	}
	fmt.Println("there is no error in get method")
	return employee
}