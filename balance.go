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
    return
  }
  defer response.Body.Close()

  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    return
  }

  rawBalance, err := jsonparser.GetString(body, "balance")
  if err != nil {
    return
  }

  balanceFloat, err := strconv.ParseFloat(rawBalance, bitSize64)
  if err != nil {
    return
  }

  balance = Balance{approximate: int(balanceFloat), exact: balanceFloat}

  return
}
