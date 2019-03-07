package client

import (
	"fmt"

	"monita/handlers/observable"
)

// CreateObservable creates new Observable
func (mc *MonitaClient) CreateObservable(request observable.CreateObservableRequest) (*observable.CreateObservableResponse, error) {
	res, err := mc.request("POST", "v1/user/observables", request, &observable.CreateObservableResponse{}, 201)

	if err != nil {
		return nil, err
	}

	result, _ := res.(*observable.CreateObservableResponse)

	return result, nil
}

// DeleteObservable deletes Observable by id
func (mc *MonitaClient) DeleteObservable(id uint) error {
	_, err := mc.request("DELETE", fmt.Sprintf("v1/user/observables/%d", id), nil, nil, 200)

	return err
}

// LoadObservable returns Observable with new data
func (mc *MonitaClient) LoadObservable(id uint) (*observable.LoadDataResponse, error) {
	res, err := mc.request("GET", fmt.Sprintf("v1/user/observables/%d", id), nil, &observable.LoadDataResponse{}, 200)

	if err != nil {
		return nil, err
	}

	result, _ := res.(*observable.LoadDataResponse)

	return result, nil
}

// HandleObservable returns handled Observable
func (mc *MonitaClient) HandleObservable(id uint) (*observable.HandleResponse, error) {
	res, err := mc.request("POST", fmt.Sprintf("v1/user/observables/%d", id), nil, &observable.HandleResponse{}, 200)

	if err != nil {
		return nil, err
	}

	result, _ := res.(*observable.HandleResponse)

	return result, nil
}

// MuteObservable mute notifications for Observable
func (mc *MonitaClient) MuteObservable(id uint) error {
	_, err := mc.request("PATCH", fmt.Sprintf("v1/user/observables/%d/mute", id), nil, nil, 200)

	return err
}

// UnmuteObservable unmute notifications for Observable
func (mc *MonitaClient) UnmuteObservable(id uint) error {
	_, err := mc.request("PATCH", fmt.Sprintf("v1/user/observables/%d/unmute", id), nil, nil, 200)

	return err
}

// ReorderObservable change Observable order
func (mc *MonitaClient) ReorderObservable(id uint, request observable.ReorderRequest) error {
	_, err := mc.request("PATCH", fmt.Sprintf("v1/user/observables/%d/reorder", id), request, nil, 200)

	return err
}
