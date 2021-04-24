package user

import (
	"net/http"
	"testing"

	"github.com/Killian264/YTLocker/golocker/data"
	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/stretchr/testify/assert"
)

var validUser = models.User{
	Username: "Killian",
	Email:    "killiandebacker@gmail.com",
	Password: "superpassword1234567",
}

////////////////////// Register //////////////////////

func Test_Register_User(t *testing.T) {

	service := createMockServices()

	err := service.RegisterUser(&validUser)
	assert.Nil(t, err)

	user, err := service.GetUserByID(validUser.ID)
	assert.Nil(t, err)

	assert.Equal(t, user.Email, validUser.Email)
}

func Test_Registration_Validates_Information(t *testing.T) {

	service := createMockServices()

	badUsername := validUser
	badUsername.Username = ""

	badEmail := validUser
	badEmail.Email = ""

	badPassword := validUser
	badPassword.Password = ""

	err := service.RegisterUser(&badUsername)
	assert.NotNil(t, err)

	err = service.RegisterUser(&badEmail)
	assert.NotNil(t, err)

	err = service.RegisterUser(&badPassword)
	assert.NotNil(t, err)

}

func Test_Error_On_Duplicate_Email(t *testing.T) {
	service := createMockServices()

	err := service.RegisterUser(&validUser)
	assert.Nil(t, err)

	err = service.RegisterUser(&validUser)
	assert.NotNil(t, err)
}

func Test_Password_Is_Hashed(t *testing.T) {
	service := createMockServices()

	original := validUser

	err := service.RegisterUser(&validUser)
	assert.Nil(t, err)

	assert.NotEqual(t, original.Password, validUser.Password)
}

////////////////////// GetUserFromRequest //////////////////////
func Test_Valid_Bearer(t *testing.T) {

	service := createMockServices()
	service.RegisterUser(&validUser)

	req, _ := http.NewRequest("GET", "/adsfasdf/asdf", nil)
	req.Header["Authorization"] = []string{"TEMP_API_BEARER"}

	actual, err := service.GetUserFromRequest(req)
	assert.Nil(t, err)
	assert.Equal(t, validUser.ID, actual.ID)
}

func Test_InValid_Bearer(t *testing.T) {

	service := createMockServices()

	req, _ := http.NewRequest("GET", "/adsfasdf/asdf", nil)
	req.Header["Authorization"] = []string{"INVALID_BEARER"}

	_, err := service.GetUserFromRequest(req)
	assert.NotNil(t, err)
}

func Test_Unset_Bearer(t *testing.T) {

	service := createMockServices()
	req, _ := http.NewRequest("GET", "/adsfasdf/asdf", nil)

	_, err := service.GetUserFromRequest(req)
	assert.NotNil(t, err)
}

////////////////////// ValidEmail //////////////////////

func TestValidEmail(t *testing.T) {

	service := createMockServices()

	valid, err := service.ValidEmail(validUser.Email)
	assert.True(t, valid)
	assert.Nil(t, err)

	service.RegisterUser(&validUser)

	valid, err = service.ValidEmail(validUser.Email)
	assert.False(t, valid)
	assert.Nil(t, err)

}

func createMockServices() *User {

	db := data.InMemorySQLiteConnect()

	service := NewUser(
		IUserData(db),
	)

	return service
}
