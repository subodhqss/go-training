// Code generated by mockery v2.7.5. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/subodhqss/go-training/models"

	multipart "mime/multipart"
)

// employeService is an autogenerated mock type for the employeService type
type employeService struct {
	mock.Mock
}

// DeleteEmployee provides a mock function with given fields: _a0, _a1
func (_m *employeService) DeleteEmployee(_a0 models.Employee, _a1 string) *models.Employee {
	ret := _m.Called(_a0, _a1)

	var r0 *models.Employee
	if rf, ok := ret.Get(0).(func(models.Employee, string) *models.Employee); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Employee)
		}
	}

	return r0
}

// PrintEmploye provides a mock function with given fields:
func (_m *employeService) PrintEmploye() []*models.Employee {
	ret := _m.Called()

	var r0 []*models.Employee
	if rf, ok := ret.Get(0).(func() []*models.Employee); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Employee)
		}
	}

	return r0
}

// PrintEmployeId provides a mock function with given fields: _a0
func (_m *employeService) PrintEmployeId(_a0 string) *models.Employee {
	ret := _m.Called(_a0)

	var r0 *models.Employee
	if rf, ok := ret.Get(0).(func(string) *models.Employee); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Employee)
		}
	}

	return r0
}

// SaveEmployee provides a mock function with given fields: _a0
func (_m *employeService) SaveEmployee(_a0 models.Employee) *models.Employee {
	ret := _m.Called(_a0)

	var r0 *models.Employee
	if rf, ok := ret.Get(0).(func(models.Employee) *models.Employee); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Employee)
		}
	}

	return r0
}

// SaveEmployeeBulk provides a mock function with given fields: _a0
func (_m *employeService) SaveEmployeeBulk(_a0 []models.Employee) {
	_m.Called(_a0)
}

// SaveEmployeeCSV provides a mock function with given fields: _a0
func (_m *employeService) SaveEmployeeCSV(_a0 models.Employee) *models.Employee {
	ret := _m.Called(_a0)

	var r0 *models.Employee
	if rf, ok := ret.Get(0).(func(models.Employee) *models.Employee); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Employee)
		}
	}

	return r0
}

// Update provides a mock function with given fields: _a0
func (_m *employeService) Update(_a0 models.Employee) *models.Employee {
	ret := _m.Called(_a0)

	var r0 *models.Employee
	if rf, ok := ret.Get(0).(func(models.Employee) *models.Employee); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Employee)
		}
	}

	return r0
}

// UpdateEmployee provides a mock function with given fields: _a0
func (_m *employeService) UpdateEmployee(_a0 models.Employee) *models.Employee {
	ret := _m.Called(_a0)

	var r0 *models.Employee
	if rf, ok := ret.Get(0).(func(models.Employee) *models.Employee); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Employee)
		}
	}

	return r0
}

// UpdateImage provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *employeService) UpdateImage(_a0 string, _a1 multipart.File, _a2 *multipart.FileHeader, _a3 string) {
	_m.Called(_a0, _a1, _a2, _a3)
}
