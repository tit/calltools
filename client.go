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
  pathPhonesCall   = baseUrl + "/phones/call"
)

// CallTools is main struct for the API
type Client struct {
  // get API public key on https://zvonok.com/manager/users/profile/
  // @see https://zvonok.com/manager/users/profile/
  ApiPublicKey string
}

type Call struct {
  CallId      int
  Balance     float64
  PhoneNumber string
  Created     time.Time
}

func (client *Client) Balance() (balance float64, err error) {
  requestUrl := fmt.Sprintf("%s?public_key=%s", pathUsersBalance, client.ApiPublicKey)

  response, err := http.Get(requestUrl)
  if err != nil {
    return
  }

  body := response.Body
  defer response.Body.Close()

  json, err := ioutil.ReadAll(body)
  if err != nil {
    return
  }

  balanceString, err := jsonparser.GetString(json, "balance")
  if err != nil {
    return
  }

  balance, err = strconv.ParseFloat(balanceString, bitSize64)
  if err != nil {
    return
  }

  return
}

func (client *Client) AddCall(campaignId int, phoneNumber string) (call Call, err error) {
  data := url.Values{
    "public_key":  {client.ApiPublicKey},
    "campaign_id": {strconv.Itoa(campaignId)},
    "phone":       {phoneNumber},
  }

  response, err := http.PostForm(pathPhonesCall, data)
  if err != nil {
    return
  }

  body := response.Body
  defer response.Body.Close()

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
