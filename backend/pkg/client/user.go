package client

import (
	"fmt"
	"monita/handlers/user"
)

// CreateUser returns new User
func (mc *MonitaClient) CreateUser(request user.CreateUserRequest) (*user.CreateUserResponse, error) {
	res, err := mc.request("POST", "v1/users", request, &user.CreateUserResponse{}, 201)

	if err != nil {
		return nil, err
	}

	result, _ := res.(*user.CreateUserResponse)

	return result, nil
}

// DeleteUser deletes User by name
func (mc *MonitaClient) DeleteUser(name string) error {
	_, err := mc.request("DELETE", fmt.Sprintf("v1/users/%s", name), nil, nil, 200)

	return err
}

// Login returns token
func (mc *MonitaClient) Login(request user.LoginRequest) (*user.LoginResponse, error) {
	res, err := mc.request("POST", "v1/login", request, &user.LoginResponse{}, 200)

	if err != nil {
		return nil, err
	}

	result, _ := res.(*user.LoginResponse)

	mc.SetAuthToken(result.Token)

	return result, nil
}

// CurrentUser returns current User data
func (mc *MonitaClient) CurrentUser() (*user.CurrentUserResponse, error) {
	res, err := mc.request("GET", "v1/user", nil, &user.CurrentUserResponse{}, 200)

	if err != nil {
		return nil, err
	}

	result, _ := res.(*user.CurrentUserResponse)

	return result, nil
}

// ChangePassword changes authenticated User password
func (mc *MonitaClient) ChangePassword(request user.ChangePasswordRequest) error {
	_, err := mc.request("PATCH", "v1/user/password", request, nil, 200)

	return err
}

// ChangeEmail changes authenticated User email
func (mc *MonitaClient) ChangeEmail(request user.ChangeEmailRequest) error {
	_, err := mc.request("PATCH", "v1/user/email", request, nil, 200)

	return err
}

// MuteNotifications mute all notifications
func (mc *MonitaClient) MuteNotifications() error {
	_, err := mc.request("PATCH", "v1/user/mute", nil, nil, 200)

	return err
}

// UnmuteNotifications unmute all notifications
func (mc *MonitaClient) UnmuteNotifications() error {
	_, err := mc.request("PATCH", "v1/user/unmute", nil, nil, 200)

	return err
}
