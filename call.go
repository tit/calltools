package main

import (
  "fmt"
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

type Call struct {
  phoneNumber       string
  status            string
  callId            int
  created           time.Time
  updated           time.Time
  duration          int
  ivrData           []IvrData
  completed         time.Time
  buttonNum         int
  actionType        string
  dialStatus        int
  userChoice        string
  audioclipId       int
  recordedAudio     url.URL
  statusDisplay     string
  userChoiceDisplay string
}

type IvrData struct {
  ivrNum       int
  webhook      string
  smsName      string
  smsText      string
  toPhone      string
  buttonNum    int
  toSipname    string
  actionType   int
  statusName   string
  recognizeNum string
  followIvrNum string
}

func (client *Client) CallByPhoneNumber(campaignId int, phoneNumber string, fromDateCreated time.Time, toDateCreated time.Time, fromDateUpdated time.Time, toDateUpdated time.Time) (calls []Call, err error) {
  dateTimeFormat := "2006-01-02 15:04:05"

  queryUrl := fmt.Sprintf("%s/?public_key=%s&campaign_id=%d&phone=%s&from_created_date=%s&to_created_date=%s&from_updated_date=%s&to_updated_date%s", pathCallsByPhone, client.ApiPublicKey, campaignId, phoneNumber,
    fromDateCreated.Format(dateTimeFormat),
    toDateCreated.Format(dateTimeFormat),
    fromDateUpdated.Format(dateTimeFormat),
    toDateUpdated.Format(dateTimeFormat))

  response, err := http.Get(queryUrl)

  if err != nil {
    return
  }

  defer response.Body.Close()

  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    return
  }

  _, err = jsonparser.ArrayEach(body, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
    phoneNumber, err := jsonparser.GetString(value, "phone")
    if err != nil {
      return
    }

    status, err := jsonparser.GetString(value, "status")
    if err != nil {
      return
    }

    callIdInt64, err := jsonparser.GetInt(value, "call_id")
    if err != nil {
      return
    }
    callId := int(callIdInt64)

    createdString, err := jsonparser.GetString(value, "created")
    if err != nil {
      return
    }

    created, err := time.Parse(time.RFC3339, createdString)
    if err != nil {
      return
    }

    updatedString, err := jsonparser.GetString(value, "updated")
    if err != nil {
      return
    }

    updated, err := time.Parse(time.RFC3339, updatedString)
    if err != nil {
      return
    }

    durationInt64, err := jsonparser.GetInt(value, "duration")
    if err != nil {
      return
    }
    duration := int(durationInt64)

    var ivrDatas []IvrData

    _, err = jsonparser.ArrayEach(body, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
      ivrNumInt64, err := jsonparser.GetInt(value, "ivr_num")
      if err != nil {
        return
      }

      ivrNum := int(ivrNumInt64)

      ivrData := IvrData{
        ivrNum: ivrNum,
      }

      ivrDatas = append(ivrDatas, ivrData)

    }, "ivr_data")

    call := Call{
      phoneNumber: phoneNumber,
      status:      status,
      callId:      callId,
      created:     created,
      updated:     updated,
      duration:    duration,
    }

    calls = append(calls, call)
  })
  if err != nil {
    return
  }

  return
}

func (client *Client) CallByCallId(callId int) (calls []Call, err error) {
  return
}
