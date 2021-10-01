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

	user, err := service.Login(validUser)
	assert.Nil(t, err)
	assert.NotEmpty(t, user)
	assert.NotEmpty(t, user.Session.Bearer)
}

func Test_RefreshSession(t *testing.T) {
	service := createMockServices()

	user, _ := service.Login(validUser)

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

	expected, _ := service.Login(validUser)

	actual, err := service.GetUserFromBearer(expected.Session.Bearer)
	assert.Nil(t, err)
	assert.Equal(t, expected.ID, actual.ID)

	service.RefreshSession(expected)

	actual, err = service.GetUserFromBearer(expected.Session.Bearer)
	assert.NotNil(t, err)
}

func createMockServices() *User {
	db := data.InMemorySQLiteConnect()

	service := NewUser(
		IUserData(db),
	)

	return service
}
