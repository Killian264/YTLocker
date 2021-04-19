package user

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Killian264/YTLocker/golocker/interfaces"
	"github.com/Killian264/YTLocker/golocker/mocks"
	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

////////////////////// GetUserFromRequest //////////////////////
func TestValidBearer(t *testing.T) {

	service, data := createMockServices()

	req, err := http.NewRequest("GET", "/adsfasdf/asdf", nil)
	assert.Nil(t, err)

	req.Header["Authorization"] = []string{"TEMP_API_BEARER"}

	expected := &models.User{
		Username: "Cool",
	}

	data.On("GetFirstUser").Return(expected, nil)

	actual, err := service.GetUserFromRequest(req)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestInvalidBearer(t *testing.T) {

	service, _ := createMockServices()

	req, err := http.NewRequest("GET", "/adsfasdf/asdf", nil)
	assert.Nil(t, err)

	req.Header["Authorization"] = []string{"INVALID_BEARER"}

	_, err = service.GetUserFromRequest(req)
	assert.NotNil(t, err)
}

func TestUnsetBearer(t *testing.T) {

	service, _ := createMockServices()

	req, err := http.NewRequest("GET", "/adsfasdf/asdf", nil)
	assert.Nil(t, err)

	_, err = service.GetUserFromRequest(req)
	assert.NotNil(t, err)
}

////////////////////// Register //////////////////////

type SimpleTest struct {
	input    models.User
	expected error
}

func TestRegisterUser(t *testing.T) {

	service, data := createMockServices()

	data.On("NewUser", mock.Anything).Return(nil)

	tests := []SimpleTest{
		{
			input:    models.User{Username: ""},
			expected: fmt.Errorf(""),
		},
		{
			input:    models.User{Username: "qqq", Email: ""},
			expected: fmt.Errorf(""),
		},
		{
			input:    models.User{Username: "qqq", Email: "email", Password: ""},
			expected: fmt.Errorf(""),
		},
		{
			input:    models.User{Username: "qqq", Email: "email", Password: "qqqqqq"},
			expected: fmt.Errorf(""),
		},
		{
			input:    models.User{Username: "qqq", Email: "email", Password: "qqqqqqqqqq66666666"},
			expected: nil,
		},
	}

	for _, test := range tests {
		err := service.RegisterUser(&test.input)
		if test.expected == nil {
			assert.Nil(t, err)
		} else {
			assert.NotNil(t, err)
		}
	}
}

////////////////////// ValidEmail //////////////////////

func TestValidEmail(t *testing.T) {

	service, data := createMockServices()

	data.On("GetUserByEmail", mock.Anything).Return(nil, fmt.Errorf("helsdfjasd")).Once()
	valid, err := service.ValidEmail("sadjfka")
	assert.False(t, valid)
	assert.NotNil(t, err)

	data.On("GetUserByEmail", mock.Anything).Return(&models.User{}, nil).Once()
	valid, err = service.ValidEmail("sadjfka")
	assert.False(t, valid)
	assert.Nil(t, err)

	data.On("GetUserByEmail", mock.Anything).Return(nil, nil).Once()
	valid, err = service.ValidEmail("sadjfka")
	assert.True(t, valid)
	assert.Nil(t, err)
}

func createMockServices() (*User, *mocks.IUserData) {

	dataMock := &mocks.IUserData{}

	service := NewUser(
		interfaces.IUserData(dataMock),
	)

	return service, dataMock
}
