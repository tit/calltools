package main

import "fmt"

func main() {
  client := Client{ApiPublicKey: "621a00762d174a32e159ec52781f35af"}
  balance, err := client.Balance()
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(balance)
  // var callByPhoneNumber CallByPhoneNumber
  // callByPhoneNumber.campaignId = 148186874
  // callByPhoneNumber.phoneNumber = "+79137812231"
  // _ = client.CallResultByPhoneNumber(callByPhoneNumber)
  // fmt.Println(callResultByPhoneNumberResults[0].phoneNumber)
}
