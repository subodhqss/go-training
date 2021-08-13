// Code generated by mockery v2.7.5. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/subodhqss/go-training/models"
)

// LoginReposiotry is an autogenerated mock type for the LoginReposiotry type
type LoginReposiotry struct {
	mock.Mock
}

// GetEmployeEmail provides a mock function with given fields: email
func (_m *LoginReposiotry) GetEmployeEmail(email string) *models.Employee {
	ret := _m.Called(email)

	var r0 *models.Employee
	if rf, ok := ret.Get(0).(func(string) *models.Employee); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Employee)
		}
	}

	return r0
}
