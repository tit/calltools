package calltools

import (
	"fmt"
	"github.com/buger/jsonparser"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	bitSize64 = 64

	baseUrl          = "https://calltools.ru/lk/cabapi_external/api/v1"
	pathUsersBalance = baseUrl + "/users/balance"
	pathPhonesCall   = baseUrl + "/phones/Call"
)

// CallTools is main struct for the API
type Client struct {
	// get API public key on https://zvonok.com/manager/users/profile/
	// @see https://zvonok.com/manager/users/profile/
	apiPublicKey string
}

func NewClient(apiPublicKey string) *Client {
	return &Client{apiPublicKey: apiPublicKey}
}

type Call struct {
	CallId      int
	Balance     float64
	PhoneNumber string
	Created     time.Time
}

func (client *Client) Balance() (balance float64, err error) {
	requestUrl := fmt.Sprintf("%s?public_key=%s", pathUsersBalance, client.apiPublicKey)

	response, err := http.Get(requestUrl)
	if err != nil {
		return balance, err
	}

	body := response.Body

	json, err := ioutil.ReadAll(body)
	if err != nil {
		return balance, err
	}

	balanceString, err := jsonparser.GetString(json, "balance")
	if err != nil {
		return balance, err
	}

	balance, err = strconv.ParseFloat(balanceString, bitSize64)
	if err != nil {
		return balance, err
	}

	return balance, err
}

func (client *Client) AddCall(campaignId int, phoneNumber string) (call Call, err error) {
	data := url.Values{
		"public_key":  {client.apiPublicKey},
		"campaign_id": {strconv.Itoa(campaignId)},
		"phone":       {phoneNumber},
	}

	response, err := http.PostForm(pathPhonesCall, data)
	if err != nil {
		return
	}

	body := response.Body

	json, err := ioutil.ReadAll(body)
	if err != nil {
		return
	}

	callId, err := jsonparser.GetInt(json, "call_id")
	if err != nil {
		return call, fmt.Errorf(string(json))
	}

	balanceString, err := jsonparser.GetString(json, "balance")
	if err != nil {
		return
	}
	balance, err := strconv.ParseFloat(balanceString, bitSize64)
	if err != nil {
		return
	}

	phoneNumber, err = jsonparser.GetString(json, "phone")
	if err != nil {
		return
	}

	createdString, err := jsonparser.GetString(json, "created")
	if err != nil {
		return
	}

	created, err := time.Parse(time.RFC3339, createdString)
	if err != nil {
		return
	}

	call.CallId = int(callId)
	call.Balance = balance
	call.PhoneNumber = phoneNumber
	call.Created = created

	return
}
