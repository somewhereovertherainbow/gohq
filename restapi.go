package gohq

import (
	"net/http"
	"bytes"
	"encoding/json"
	"log"
	"io/ioutil"
	"errors"
	"strconv"
)

// Request makes (GET/POST/PUT/PATCH/etc..) requests to the HQ API
func (a *Account) Request(method string, urlStr string, data interface{}, auth bool) (response []byte, err error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return
	}

	req, err := http.NewRequest(method, urlStr, bytes.NewBuffer(dataBytes))
	if err != nil {
		return
	}

	if auth {
		req.Header.Set("Authorization", "Bearer "+a.AccessToken)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "hq-viewer/1.2.18 (iPhone; iOS 11.2.2; Scale/3.00)")
	req.Header.Set("Content-Length", strconv.Itoa(len(dataBytes)))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	defer func() {
		if resp.Body.Close() != nil {
			log.Println("error closing resp body")
		}
	}()

	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var hqerr HQError
	if err = json.Unmarshal(response, &hqerr); err == nil && hqerr.Error != "" {
		err = errors.New(hqerr.Error)
		if err.Error() == "not authorized" && urlStr != EndpointTokens {
			if a.LoginToken != "" {
				tokens, err_two := a.Tokens()
				if err_two == nil {
					a.AccessToken = tokens.AccessToken
					a.LoginToken = tokens.LoginToken
					a.AuthToken = tokens.AuthToken

					response, err = a.Request(method, urlStr, data, auth)
				}
			}
		}
	}

	// TODO: Add a check to see if HQ ever goes down

	return
}

// Tokens refreshes new tokens based on the login token
func (a *Account) Tokens() (t *Tokens, err error) {
	type Data struct {
		Token string `json:"token"`
	}

	resp, err := a.Request("POST", EndpointTokens, Data{Token: a.LoginToken}, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &t)

	return
}

// Me gets updated profile information
func (a *Account) Me() (t *Me, err error) {
	type Data struct {
	}

	resp, err := a.Request("GET", EndpointMe, Data{}, true)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &t)

	return
}

// Cashout sends a cashout request to HQ
func (a *Account) Cashout(email string) (cd *CashoutData, err error) {
	type Data struct {
		Email string `json:"email"`
	}

	resp, err := a.Request("POST", EndpointMe, Data{Email: email}, true)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &cd)

	return
}

// Payouts gets all of the past payout data
func (a *Account) Payouts() (pd *PayoutData, err error) {
	type Data struct {
	}

	resp, err := a.Request("GET", EndpointMe, Data{}, true)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &pd)

	return
}

// Weekly runs the makeItRain easter egg
func (a *Account) Weekly() (err error) {
	type Data struct {
	}

	if _, err = a.Request("GET", EndpointMe, Data{}, true); err != nil {
		return
	}

	return
}

// SearchUser searches for a user
func (a *Account) SearchUser(username string) (sd *SearchData, err error) {
	type Data struct {
	}

	resp, err := a.Request("GET", EndpointSearchUser(username), Data{}, true)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &sd)

	return
}
