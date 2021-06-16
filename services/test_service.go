package services

import (
	"fmt"

	"github.com/subodhqss/go-training/repository"
)

type TestService interface {
	PrintTest()
}

type testSrv struct {
	testRepo repository.TestRepository
}

func NewTestService() TestService {
	return &testSrv{}
}

func (ts *testSrv) PrintTest() {
	ts.testRepo.PrintTest()
	fmt.Println("this is test service")
}
