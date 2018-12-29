package main

import "fmt"

func main() {
  client := Client{ApiPublicKey: ""}
  balance, _ := client.Balance()
  fmt.Println(balance)
  // var callByPhoneNumber CallByPhoneNumber
  // callByPhoneNumber.campaignId = 148186874
  // callByPhoneNumber.phoneNumber = "+79137812231"
  // _ = client.CallResultByPhoneNumber(callByPhoneNumber)
  // fmt.Println(callResultByPhoneNumberResults[0].phoneNumber)
}
