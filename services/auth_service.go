package service 
import (
	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"

)

type loginService interface {
	GetEmployeEmail(string) *models.Login
}
type logServ struct {
	logRepo repository.LoginReposiotry
}
func NewLoginService(logRepo repository.LoginRepository) logService {
	return &logServ{logRepo: logRepo}

}
 
func loginEmployee(logRepo repository.LoginReposiotry) loginService {
	//func (ls *logServ)GetEmployeEmail(email string) *models.Login{
     log := ls.logRepo.GetEmployeEmail(email)
	 return log
	 if log == nil {
		 return nil
	 }
	// if log.password, err == Login.password
	






