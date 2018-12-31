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

const (
  bitSize64 = 64

  baseUrl          = "https://calltools.ru/lk/cabapi_external/api/v1"
  pathUsersBalance = baseUrl + "/users/balance"
  pathPhonesCall   = baseUrl + "/phones/call"
  pathCallsByPhone = baseUrl + "/phones/calls_by_phone"
)

// CallTools is main struct for the API
type Client struct {
  // get API public key on https://zvonok.com/manager/users/profile/
  // @see https://zvonok.com/manager/users/profile/
  ApiPublicKey string
}

type CallByPhoneNumber struct {
  campaignId      int
  phoneNumber     string
  fromDateCreated time.Time
  toDateCreated   time.Time
  fromDateUpdated time.Time
  toDateUpdated   time.Time
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

type CallByPhoneNumberResult struct {
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
  sourceJson        string
}

func (client *Client) CallResultByPhoneNumber(callByPhoneNumber CallByPhoneNumber) (callByPhoneNumberResults []CallByPhoneNumberResult) {
  // "from_created_date": {callByPhoneNumber.fromDateCreated.Format("2006-01-02 15:04:05")},
  // "to_created_date":   {callByPhoneNumber.toDateCreated.Format("2006-01-02 15:04:05")},
  // "from_updated_date": {callByPhoneNumber.fromDateUpdated.Format("2006-01-02 15:04:05")},
  // "to_updated_date":   {callByPhoneNumber.toDateUpdated.Format("2006-01-02 15:04:05")},

  apiPublicKey := client.ApiPublicKey
  campaignId := strconv.Itoa(callByPhoneNumber.campaignId)
  phoneNumber := callByPhoneNumber.phoneNumber

  dateTimeFormat := "2006-01-02 15:04:05"
  fromDateCreated := callByPhoneNumber.fromDateCreated.Format(dateTimeFormat)
  toCreatedDate := callByPhoneNumber.toDateCreated.Format(dateTimeFormat)
  fromUpdatedDate := callByPhoneNumber.fromDateUpdated.Format(dateTimeFormat)
  toUpdatedDate := callByPhoneNumber.toDateUpdated.Format(dateTimeFormat)

  response, err := http.Get(fmt.Sprintf("%s/?public_key=%s&campaign_id=%s&phone=%s&from_created_date=%s&to_created_date=%s&from_updated_date=%s&to_updated_date%s", pathCallsByPhone, apiPublicKey, campaignId, phoneNumber, fromDateCreated, toCreatedDate, fromUpdatedDate, toUpdatedDate))
  if err != nil {
    return
  }

  body := response.Body
  defer response.Body.Close()

  json, err := ioutil.ReadAll(body)
  if err != nil {
    return
  }

  var callByPhoneNumberResult CallByPhoneNumberResult

  _, err = jsonparser.ArrayEach(json, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
    phoneNumber, err := jsonparser.GetString(value, "phone")
    callByPhoneNumberResult.phoneNumber = phoneNumber
    // callByPhoneNumberResult.status = status
    // callByPhoneNumberResult.callId = int(callId)
    //
    // status, err := jsonparser.GetString(value, "status")
    // callId, err := jsonparser.GetInt(value, "call_id")
  })
  if err != nil {
    return
  }

  // status, err := jsonparser.GetString(json, "status")
  // if err != nil {
  //   return
  // }
  //
  // callId, err := jsonparser.GetInt(json, "call_id")
  // if err != nil {
  //   return
  // }

  // callByPhoneNumberResult.created = created
  // callByPhoneNumberResult.updated = updated
  // callByPhoneNumberResult.duration = duration
  // callByPhoneNumberResult.ivrData = ivrData
  // callByPhoneNumberResult.completed = completed
  // callByPhoneNumberResult.buttonNum = buttonNum
  // callByPhoneNumberResult.actionType = actionType
  // callByPhoneNumberResult.dialStatus = dialStatus
  // callByPhoneNumberResult.userChoice = userChoice
  // callByPhoneNumberResult.audioclipId = audioclipId
  // callByPhoneNumberResult.recordedAudio = recorded_audio
  // callByPhoneNumberResult.statusDisplay = statusDisplay
  // callByPhoneNumberResult.userChoiceDisplay = userChoiceDisplay
  // callByPhoneNumberResult.sourceJson = string(json)

  callByPhoneNumberResults = append(callByPhoneNumberResults, callByPhoneNumberResult)
  return
}
