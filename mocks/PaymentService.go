// Code generated by mockery v2.7.5. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/subodhqss/go-training/models"
)

// PaymentService is an autogenerated mock type for the PaymentService type
type PaymentService struct {
	mock.Mock
}

// PrintPayment provides a mock function with given fields:
func (_m *PaymentService) PrintPayment() []*models.Payment {
	ret := _m.Called()

	var r0 []*models.Payment
	if rf, ok := ret.Get(0).(func() []*models.Payment); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Payment)
		}
	}

	return r0
}

// SavePayment provides a mock function with given fields: _a0
func (_m *PaymentService) SavePayment(_a0 models.Payment) *models.Payment {
	ret := _m.Called(_a0)

	var r0 *models.Payment
	if rf, ok := ret.Get(0).(func(models.Payment) *models.Payment); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Payment)
		}
	}

	return r0
}
