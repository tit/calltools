package main

import "fmt"

func main() {
  client := Client{ApiPublicKey: "621a00762d174a32e159ec52781f35af"}
  fmt.Println(client.Balance())
  // var callByPhoneNumber CallByPhoneNumber = CallByPhoneNumber{
  //   campaignId:  148186874,
  //   phoneNumber: "+79137812231",
  // }
}
