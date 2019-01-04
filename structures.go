package calltools

import (
  "net/url"
  "time"
)

const (
  bitSize64 int = 64

  baseUrl          string = "https://calltools.ru/lk/cabapi_external/api/v1"
  pathUsersBalance        = baseUrl + "/users/balance"
  pathPhonesCall          = baseUrl + "/phones/call"
  pathCallsByPhone        = baseUrl + "/phones/calls_by_phone"
  pathCallById            = baseUrl + "/phones/call_by_id"
  pathRemoveCall          = baseUrl + "/phones/remove_call"
)

type Client struct {
  ApiPublicKey string
}

type AddCall struct {
  CallId      int
  Balance     float64
  PhoneNumber string
  Created     time.Time
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
  recordedAudio     *url.URL
  statusDisplay     string
  userChoiceDisplay string
  dialStatusDisplay string
}
