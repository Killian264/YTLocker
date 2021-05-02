package handlers

// func TestHandleRegistration(t *testing.T) {

// 	service := &mocks.IUser{}

// 	user := models.User{
// 		Username: "killian",
// 		Email:    "killian@ytlocker.com",
// 		Password: "coolstorybro",
// 	}

// 	body, err := json.Marshal(user)
// 	assert.Nil(t, err)

// 	request, err := http.NewRequest("GET", "/registration", bytes.NewBuffer(body))
// 	assert.Nil(t, err)

// 	fake := test.FakeRequest{
// 		Services: &services.Services{
// 			User: services.IUser(service),
// 		},
// 		Route:   "/registration",
// 		Request: request,
// 		Handler: HandleRegistration,
// 	}

// 	service.On("ValidEmail", user.Email).Return(true, nil)
// 	service.On("RegisterUser", &user).Return(nil)

// 	res := test.SendFakeRequest(fake)
// 	assert.Equal(t, res.StatusCode, 200)

// }
