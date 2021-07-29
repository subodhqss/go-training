package service

import (
	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
)

type loginService interface {
	PrintLogin() []*models.Employee
}

type logServ struct {
	logRepo repository.LoginReposiotry
}

func NewLoginService(logRepo repository.LoginReposiotry) loginService {
	return &logServ{logRepo: logRepo}
}

func (es *logServ) PrintLogin() []*models.Employee {
	log := es.logRepo.PrintLogin()
	return log
}
