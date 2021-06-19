// // package services

// import (
// 	"github.com/subodhqss/go-training/models"
// 	"github.com/subodhqss/go-training/repository"
// )

// type TestService interface {
// 	PrintTest() *models.Test
// }

// type testSrv struct {
// 	testRepo repository.TestRepository
// }

// func NewTestService(testRepo repository.TestRepository) TestService {
// 	return &testSrv{testRepo: testRepo}
// }

// func (ts *testSrv) PrintTest() *models.Test {
// 	test := ts.testRepo.PrintTest()
// 	// test := &models.Test{ID: 1, Message: "Hi SUbodh"}
// 	return test
// }
