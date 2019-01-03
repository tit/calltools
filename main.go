package main

import "fmt"

func main() {
  client := Client{ApiPublicKey: "621a00762d174a32e159ec52781f35af"}
  // fmt.Println(client.Balance())

  // result, _ := client.CallByPhoneNumber(148186874, "+79137812231")
  result, _ := client.CallByCallId(181228342977211)

  fmt.Println(result[0].recordedAudio)

}
