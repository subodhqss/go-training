package repository

import "fmt"

type TestRepository interface {
	PrintTest()
}

func NewTestRepo() TestRepository {
	return &testRepo{}
}

type testRepo struct{}

func (tr *testRepo) PrintTest() {
	fmt.Println("this is test repo")
}
