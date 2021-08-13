// Code generated by mockery v2.7.5. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/subodhqss/go-training/models"
)

// OrderService is an autogenerated mock type for the OrderService type
type OrderService struct {
	mock.Mock
}

// PrintOrder provides a mock function with given fields:
func (_m *OrderService) PrintOrder() []*models.Order {
	ret := _m.Called()

	var r0 []*models.Order
	if rf, ok := ret.Get(0).(func() []*models.Order); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Order)
		}
	}

	return r0
}

// PrintOrderId provides a mock function with given fields: _a0
func (_m *OrderService) PrintOrderId(_a0 string) []*models.Order {
	ret := _m.Called(_a0)

	var r0 []*models.Order
	if rf, ok := ret.Get(0).(func(string) []*models.Order); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Order)
		}
	}

	return r0
}

// SaveOrder provides a mock function with given fields: _a0
func (_m *OrderService) SaveOrder(_a0 models.Order) *models.Order {
	ret := _m.Called(_a0)

	var r0 *models.Order
	if rf, ok := ret.Get(0).(func(models.Order) *models.Order); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Order)
		}
	}

	return r0
}

// UpdateOrder provides a mock function with given fields: _a0
func (_m *OrderService) UpdateOrder(_a0 models.Order) *models.Order {
	ret := _m.Called(_a0)

	var r0 *models.Order
	if rf, ok := ret.Get(0).(func(models.Order) *models.Order); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Order)
		}
	}

	return r0
}
