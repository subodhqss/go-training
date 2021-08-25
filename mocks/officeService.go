// Code generated by mockery v2.7.5. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/subodhqss/go-training/models"
)

// officeService is an autogenerated mock type for the officeService type
type officeService struct {
	mock.Mock
}

// EditOffice provides a mock function with given fields: _a0
func (_m *officeService) EditOffice(_a0 models.Office) *models.Office {
	ret := _m.Called(_a0)

	var r0 *models.Office
	if rf, ok := ret.Get(0).(func(models.Office) *models.Office); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Office)
		}
	}

	return r0
}

// PrintOffice provides a mock function with given fields:
func (_m *officeService) PrintOffice() []*models.Office {
	ret := _m.Called()

	var r0 []*models.Office
	if rf, ok := ret.Get(0).(func() []*models.Office); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Office)
		}
	}

	return r0
}

// PrintOfficeId provides a mock function with given fields: _a0
func (_m *officeService) PrintOfficeId(_a0 string) *models.Office {
	ret := _m.Called(_a0)

	var r0 *models.Office
	if rf, ok := ret.Get(0).(func(string) *models.Office); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Office)
		}
	}

	return r0
}

// SaveOffice provides a mock function with given fields: _a0
func (_m *officeService) SaveOffice(_a0 models.Office) *models.Office {
	ret := _m.Called(_a0)

	var r0 *models.Office
	if rf, ok := ret.Get(0).(func(models.Office) *models.Office); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Office)
		}
	}

	return r0
}
