package user

import (
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

func Test_Register_User(t *testing.T) {

	service := createMockServices()

	user, err := service.Register(validUser)
	assert.Nil(t, err)

	saved, err := service.GetUser(user.ID)
	assert.Nil(t, err)

	assert.Equal(t, user.Email, saved.Email)
}

func Test_Error_On_Duplicate_Email(t *testing.T) {
	service := createMockServices()

	_, err := service.Register(validUser)
	assert.Nil(t, err)

	_, err = service.Register(validUser)
	assert.NotNil(t, err)
}

func Test_Password_Is_Hashed(t *testing.T) {
	service := createMockServices()

	original := validUser

	user, err := service.Register(validUser)
	assert.Nil(t, err)

	assert.NotEqual(t, original.Password, user.Password)
}

func TestValidEmail(t *testing.T) {

	service := createMockServices()

	valid, err := service.ValidEmail(validUser.Email)
	assert.True(t, valid)
	assert.Nil(t, err)

	service.Register(validUser)

	valid, err = service.ValidEmail(validUser.Email)
	assert.False(t, valid)
	assert.Nil(t, err)

}

func Test_Login(t *testing.T) {

	service := createMockServices()

	service.Register(validUser)

	bearer, err := service.Login(validUser.Email, validUser.Password)
	assert.Nil(t, err)
	assert.NotEmpty(t, bearer)

}

func Test_RefreshSession(t *testing.T) {

	service := createMockServices()

	user, _ := service.Register(validUser)

	session, err := service.RefreshSession(user)
	assert.Nil(t, err)
	assert.NotEmpty(t, session.Bearer)

	session2, err := service.RefreshSession(user)
	assert.Nil(t, err)
	assert.NotEmpty(t, session2.Bearer)

	assert.NotEqual(t, session.Bearer, session2.Bearer)

}

func Test_GetUserFromBearer(t *testing.T) {

	service := createMockServices()

	expected, _ := service.Register(validUser)

	actual, err := service.GetUserFromBearer(expected.Session.Bearer)
	assert.Nil(t, err)
	assert.Equal(t, expected.ID, actual.ID)

	service.RefreshSession(expected)

	actual, err = service.GetUserFromBearer(expected.Session.Bearer)
	assert.NotNil(t, err)
	assert.Nil(t, actual)

}

func createMockServices() *User {

	db := data.InMemorySQLiteConnect()

	service := NewUser(
		IUserData(db),
	)

	return service
}
