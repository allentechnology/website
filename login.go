package clientLogin

import (
	"errors"
	"golang.org/x/net/publicsuffix"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

func Login(loginPage, username, password string, timeout time.Duration) (*http.Client, error) {
	options := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	jar, err := cookiejar.New(&options)
	if err != nil {
		return nil, err
	}
	client := http.Client{Jar: jar, Timeout: timeout}
	resp, err := client.PostForm(loginPage, url.Values{
		"password": {password},
		"username": {username},
	})
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if strings.Contains(string(data), "Invalid") {
		return nil, errors.New("Login failed")
	}
	return &client, nil
}
