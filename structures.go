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
  IvrNum       int
  Webhook      string
  SmsName      string
  SmsText      string
  ToPhone      string
  ButtonNum    int
  ToSipname    string
  ActionType   int
  StatusName   string
  RecognizeNum string
  FollowIvrNum string
}

type Call struct {
  PhoneNumber       string
  Status            string
  CallId            int
  Created           time.Time
  Updated           time.Time
  Duration          int
  IvrData           []IvrData
  Completed         time.Time
  ButtonNum         int
  ActionType        string
  DialStatus        int
  UserChoice        string
  AudioclipId       int
  RecordedAudio     *url.URL
  StatusDisplay     string
  UserChoiceDisplay string
  DialStatusDisplay string
}
