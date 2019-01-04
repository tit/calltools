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

  addCall, err = addCallsByJson(body)
  if err != nil {
    return
  }

  return
}

func (client *Client) AddCallWithTTS(campaignId int, phoneNumber string, text string, speaker string) (addCall AddCall, err error) {
  data := url.Values{
    "public_key":  {client.ApiPublicKey},
    "campaign_id": {strconv.Itoa(campaignId)},
    "phone":       {phoneNumber},
    "text":        {text},
    "speaker":     {speaker},
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

  addCall, err = addCallsByJson(body)
  if err != nil {
    return
  }

  return

}

func (client *Client) CallByPhoneNumber(campaignId int, phoneNumber string) (calls []Call, err error) {
  queryUrl := fmt.Sprintf("%s/?public_key=%s&campaign_id=%d&phone=%s", pathCallsByPhone, client.ApiPublicKey, campaignId, phoneNumber)

  response, err := http.Get(queryUrl)

  if err != nil {
    return
  }

  defer response.Body.Close()

  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    return
  }

  calls, err = callsByJson(body)
  if err != nil {
    return
  }

  return
}

func (client *Client) CallByCallId(callId int) (calls []Call, err error) {
  queryUrl := fmt.Sprintf("%s/?public_key=%s&call_id=%d", pathCallById, client.ApiPublicKey, callId)

  response, err := http.Get(queryUrl)

  if err != nil {
    return
  }

  defer response.Body.Close()

  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    return
  }

  calls, err = callsByJson(body)
  if err != nil {
    return
  }

  return
}

func (client *Client) RemoveCallByPhoneNumber(campaignId int, phoneNumber string) (err error) {
  data := url.Values{
    "public_key":  {client.ApiPublicKey},
    "campaign_id": {strconv.Itoa(campaignId)},
    "phone":       {phoneNumber},
  }

  _, err = http.PostForm(pathRemoveCall, data)
  if err != nil {
    return
  }

  return
}

func (client *Client) RemoveCallByCallID(campaignId int, callId int) (err error) {
  data := url.Values{
    "public_key":  {client.ApiPublicKey},
    "campaign_id": {strconv.Itoa(campaignId)},
    "call_id":     {strconv.Itoa(callId)},
  }

  _, err = http.PostForm(pathRemoveCall, data)
  if err != nil {
    return
  }

  return
}

func addCallsByJson(body []byte) (addCall AddCall, err error) {
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

  phoneNumber, err := jsonparser.GetString(body, "phone")
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

func callsByJson(body []byte) (calls []Call, err error) {
  _, err = jsonparser.ArrayEach(body, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
    phoneNumber, err := jsonparser.GetString(value, "phone")
    if err != nil {
      return
    }

    status, err := jsonparser.GetString(value, "status")
    if err != nil {
      return
    }

    callId, err := jsonparser.GetInt(value, "call_id")
    if err != nil {
      return
    }

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

    duration, err := jsonparser.GetInt(value, "duration")
    if err != nil {
      return
    }
    var ivrDatas []IvrData

    _, err = jsonparser.ArrayEach(value, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
      ivrNum, err := jsonparser.GetInt(value, "ivr_num")
      if err != nil {
        return
      }

      webhook, err := jsonparser.GetString(value, "webhook")
      if err != nil {
        return
      }

      smsName, err := jsonparser.GetString(value, "sms_name")
      if err != nil {
        return
      }

      smsText, err := jsonparser.GetString(value, "sms_text")
      if err != nil {
        return
      }

      toPhone, _, _, err := jsonparser.Get(value, "to_phone")
      if err != nil {
        return
      }

      buttonNum, err := jsonparser.GetInt(value, "button_num")
      if err != nil {
        return
      }

      toSipname, _, _, err := jsonparser.Get(value, "to_sipname")
      if err != nil {
        return
      }

      actionType, err := jsonparser.GetInt(value, "action_type")
      if err != nil {
        return
      }

      statusName, err := jsonparser.GetString(value, "status_name")
      if err != nil {
        return
      }

      recognizeNum, _, _, err := jsonparser.Get(value, "recognize_num")
      if err != nil {
        return
      }

      followIvrNum, _, _, err := jsonparser.Get(value, "follow_ivr_num")
      if err != nil {
        return
      }

      ivrData := IvrData{
        ivrNum:       int(ivrNum),
        webhook:      webhook,
        smsName:      smsName,
        smsText:      smsText,
        toPhone:      string(toPhone),
        buttonNum:    int(buttonNum),
        toSipname:    string(toSipname),
        actionType:   int(actionType),
        statusName:   statusName,
        recognizeNum: string(recognizeNum),
        followIvrNum: string(followIvrNum),
      }

      ivrDatas = append(ivrDatas, ivrData)

    }, "ivr_data")

    completedString, err := jsonparser.GetString(value, "completed")
    if err != nil {
      return
    }

    completed, err := time.Parse(time.RFC3339, completedString)
    if err != nil {
      return
    }

    buttonNum, err := jsonparser.GetInt(value, "button_num")
    if err != nil {
      return
    }

    actionType, err := jsonparser.GetString(value, "action_type")
    if err != nil {
      return
    }

    dialStatus, err := jsonparser.GetInt(value, "dial_status")
    if err != nil {
      return
    }

    userChoice, err := jsonparser.GetString(value, "user_choice")
    if err != nil {
      return
    }

    audioclipId, err := jsonparser.GetInt(value, "audioclip_id")
    if err != nil {
      return
    }

    recordedAudioString, err := jsonparser.GetString(value, "recorded_audio")
    if err != nil {
      return
    }

    recordedAudio, err := url.Parse(recordedAudioString)
    if err != nil {
      return
    }

    statusDisplay, err := jsonparser.GetString(value, "status_display")
    if err != nil {
      return
    }

    dialStatusDisplay, err := jsonparser.GetString(value, "dial_status_display")
    if err != nil {
      return
    }

    userChoiceDisplay, err := jsonparser.GetString(value, "user_choice_display")
    if err != nil {
      return
    }

    call := Call{
      phoneNumber:       phoneNumber,
      status:            status,
      callId:            int(callId),
      created:           created,
      updated:           updated,
      duration:          int(duration),
      ivrData:           ivrDatas,
      completed:         completed,
      buttonNum:         int(buttonNum),
      actionType:        actionType,
      dialStatus:        int(dialStatus),
      userChoice:        userChoice,
      audioclipId:       int(audioclipId),
      recordedAudio:     recordedAudio,
      statusDisplay:     statusDisplay,
      dialStatusDisplay: dialStatusDisplay,
      userChoiceDisplay: userChoiceDisplay,
    }

    calls = append(calls, call)
  })
  if err != nil {
    return
  }

  return
}
