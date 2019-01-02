package main

import (
  "fmt"
  "time"
)

func main() {
  client := Client{ApiPublicKey: "621a00762d174a32e159ec52781f35af"}
  // fmt.Println(client.Balance())

  result, _ := client.CallByPhoneNumber(148186874, "+79137812231",
    time.Date(2017, 1, 1, 1, 1, 1, 1, time.Local),
    time.Date(2018, 12, 1, 1, 1, 1, 1, time.Local),
    time.Date(2017, 1, 1, 1, 1, 1, 1, time.Local),
    time.Date(2018, 12, 1, 1, 1, 1, 1, time.Local))

  fmt.Println(result)
}
