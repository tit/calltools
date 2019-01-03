package main

import "fmt"

func main() {
  client := Client{ApiPublicKey: "621a00762d174a32e159ec52781f35af"}
  balance, _ := client.Balance()
  fmt.Println(balance)

  // calls, _ := client.CallByPhoneNumber(148186874, "+79137812231")
  // calls, _ := client.CallByCallId(181228342977211)

  // fmt.Println(calls)
}
