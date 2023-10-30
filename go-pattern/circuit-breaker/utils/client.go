package utils

import (
	"errors"
	"net/http"
)

type SmsClient struct {
}

type NotificationClient interface {
	Send(baseUrl string) error // We ignore all the arguments to simplify the demo
}

func NewSmsClient() *SmsClient {
	return &SmsClient{}
}

func (s *SmsClient) Send(baseUrl string) error {
	url := baseUrl + "/"
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return errors.New("bad response")
	}

	return nil
}
