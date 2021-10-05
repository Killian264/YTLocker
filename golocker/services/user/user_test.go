package user

import (
	"testing"

	"github.com/Killian264/YTLocker/golocker/data"
	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/stretchr/testify/assert"
)

var validUser = models.User{
	Username: "killian",
	Email:    "killiandebacker@gmail.com",
	Picture:  "https://lh3.googleusercontent.com/a/default-user=s96-c",
}

func Test_Login(t *testing.T) {
	service := createMockServices()

	bearer, err := service.GenerateTemporarySessionBearer()
	assert.Nil(t, err)

	user, err := service.Login(validUser, bearer)
	assert.Nil(t, err)
	assert.NotEmpty(t, user)
	assert.NotEmpty(t, user.Session.Bearer)
}

func Test_RefreshSession(t *testing.T) {
	service := createMockServices()

	bearer, _ := service.GenerateTemporarySessionBearer()
	user, _ := service.Login(validUser, bearer)

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

	bearer, _ := service.GenerateTemporarySessionBearer()
	expected, _ := service.Login(validUser, bearer)

	actual, err := service.GetUserFromBearer(expected.Session.Bearer)
	assert.Nil(t, err)
	assert.Equal(t, expected.ID, actual.ID)

	service.RefreshSession(expected)

	actual, err = service.GetUserFromBearer(expected.Session.Bearer)
	assert.NotNil(t, err)
}

func Test_GenerateTemporarySessionBearer(t *testing.T) {
	service := createMockServices()

	one, err := service.GenerateTemporarySessionBearer()
	assert.Nil(t, err)
	assert.NotEmpty(t, one, "")

	two, err := service.GenerateTemporarySessionBearer()
	assert.Nil(t, err)
	assert.NotEmpty(t, two, "")

	assert.NotEqual(t, one, two)
}

func createMockServices() *User {
	db := data.InMemorySQLiteConnect()

	service := NewUser(
		IUserData(db),
	)

	return service
}
