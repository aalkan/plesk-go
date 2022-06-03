package client

import (
	"encoding/json"
	"fmt"
	"github.com/aalkan/plesk-go/pkg/models"
	"io/ioutil"
	"net/http"
	"reflect"
	"time"
)

//const baseURL = "https://api.central.plesk.com/30/keys"
const baseURL = "https://ka.demo.plesk.com:7050/jsonrest/business-partner/30/"

//const baseURL = "http://195.214.233.82:4000/api"

type client struct {
	baseURL    string
	config     *models.Config
	httpClient *http.Client
}

func NewClient(config *models.Config) *client {
	return &client{
		baseURL: baseURL,
		config:  config,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}
func (c *client) request(method, endpoint string, params, response interface{}) error {
	request, err := http.NewRequest(method, baseURL+endpoint, nil)
	if err != nil {
		return err
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	//request.Header.Set("HTTP_AUTH_LOGIN", c.config.Username)
	//request.Header.Set("HTTP_AUTH_PASSWD", c.config.Password)
	request.SetBasicAuth(c.config.Username, c.config.Password)
	query := request.URL.Query()
	reflectVal := reflect.ValueOf(params).Elem()
	for i := 0; i < reflectVal.NumField(); i++ {
		value := reflectVal.Field(i).Interface().(string)
		if value != "" {
			query.Add(reflectVal.Type().Field(i).Tag.Get("json"), value)
		}
	}
	request.URL.RawQuery = query.Encode()

	resp, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println("err" + fmt.Sprintf("%v", err))
	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return err
	}

	return nil
}
