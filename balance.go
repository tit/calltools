package main

import (
  "fmt"
  "github.com/buger/jsonparser"
  "io/ioutil"
  "net/http"
  "strconv"
)

type Balance struct {
  approximate int
  exact       float64
}

func (client *Client) Balance() (balance Balance, err error) {
  requestUrl := fmt.Sprintf("%s?public_key=%s", pathUsersBalance, client.ApiPublicKey)

  response, err := http.Get(requestUrl)
  if err != nil {
    return balance, fmt.Errorf(requestUrl)
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

  balanceFloat, err := strconv.ParseFloat(balanceString, bitSize64)
  if err != nil {
    return
  }

  balance = Balance{approximate: int(balanceFloat), exact: balanceFloat}

  return
}
