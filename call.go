package main

import (
  "github.com/buger/jsonparser"
  "io/ioutil"
  "net/http"
  "net/url"
  "strconv"
  "time"
)

type AddCall struct {
  CallId      int
  Balance     float64
  PhoneNumber string
  Created     time.Time
}

// Add Call by campaignId and phoneNumber
// Return AddCall struct
func (client *Client) AddCall(campaignId int, phoneNumber string) (addCall AddCall, err error) {
  data := url.Values{
    "public_key":  {client.ApiPublicKey},
    "campaign_id": {strconv.Itoa(campaignId)},
    "phone":       {phoneNumber},
  }

  response, err := http.PostForm(pathPhonesCall, data)
  if err != nil {
    return
  }

  defer response.Body.Close()

  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    return
  }

  callId, err := jsonparser.GetInt(body, "call_id")
  if err != nil {
    return
  }

  balanceString, err := jsonparser.GetString(body, "balance")
  if err != nil {
    return
  }
  balance, err := strconv.ParseFloat(balanceString, bitSize64)
  if err != nil {
    return
  }

  phoneNumber, err = jsonparser.GetString(body, "phone")
  if err != nil {
    return
  }

  createdString, err := jsonparser.GetString(body, "created")
  if err != nil {
    return
  }

  created, err := time.Parse(time.RFC3339, createdString)
  if err != nil {
    return
  }

  addCall = AddCall{
    CallId:      int(callId),
    Balance:     balance,
    PhoneNumber: phoneNumber,
    Created:     created,
  }

  return
}
