package main

const (
  bitSize64 = 64

  baseUrl          = "https://calltools.ru/lk/cabapi_external/api/v1"
  pathUsersBalance = baseUrl + "/users/balance"
  pathPhonesCall   = baseUrl + "/phones/call"
  pathCallsByPhone = baseUrl + "/phones/calls_by_phone"
)

// CallTools is main struct for the API
type Client struct {
  // get API public key on https://zvonok.com/manager/users/profile/
  // @see https://zvonok.com/manager/users/profile/
  ApiPublicKey string
}