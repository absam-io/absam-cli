package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/absam-io/absam-cli/config"
)

const (
	SUCCESS = 0
	FAIL    = 1
)

func Get(reqUrl string) ([]byte, error) {
	conf, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	parsedUrl, _ := url.Parse(reqUrl)

	req := &http.Request{
		Method: "GET",
		URL:    parsedUrl,
		Header: map[string][]string{
			"Content-Type":        {"application/json"},
			"access-token":        {conf.Token},
			"secret-access-token": {conf.Secret},
		},
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	res.Body.Close()

	return data, nil
}

func Post(reqUrl string, form url.Values) ([]byte, error) {
	conf, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, reqUrl, strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("access-token", conf.Token)
	req.Header.Add("secret-access-token", conf.Secret)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	res.Body.Close()

	return data, nil
}

func Die(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(FAIL)
}
