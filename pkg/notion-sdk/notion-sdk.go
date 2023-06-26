package notionsdk

import (
	"fmt"
	"net/http"

	"github.com/apegushin/notioncli/pkg/config"
)

type Client struct {
	cfg *config.Config
}

func NewClient(configFilePath string) *Client {
	return &Client{
		cfg: config.NewConfig(configFilePath),
	}
}

func (c *Client) Get() error {
	resp, err := http.Get("https://www.google.com/")

	if err != nil {
		fmt.Println("oh no, something terrible happened.", err)
	} else {
		if resp.Status == "200 OK" {
			fmt.Println("got 200 OK from google.com")
		} else {
			fmt.Println("got some weird status", resp.Status)
		}
	}

	return err
}
