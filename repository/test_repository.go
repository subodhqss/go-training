package repository

import (
	"github.com/subodhqss/go-training/models"
)

type TestRepository interface {
	PrintTest() *models.Test
}

func NewTestRepo() TestRepository {
	return &testRepo{}
}

type testRepo struct{}

func (tr *testRepo) PrintTest() *models.Test {
	test := &models.Test{ID: 1, Message: "Hello Sarthak"}

	return test
}
