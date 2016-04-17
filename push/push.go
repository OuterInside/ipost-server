package push

import (
  "fmt"
  apns "github.com/anachronistic/apns"
)

func PushToIos(message string) {
  payload := apns.NewPayload()
  payload.Alert = message
  payload.Badge = 42
  payload.Sound = "bingbong.aiff"

  pn := apns.NewPushNotification()
  pn.DeviceToken = "5e01160f190f4252a26c0192b6014e0df18c8368713901fad025bfd57270aa0a"
  pn.AddPayload(payload)

  client := apns.NewClient("gateway.sandbox.push.apple.com:2195", "./keys/cert.pem", "./keys/key.pem")

  resp := client.Send(pn)

  alert, _ := pn.PayloadString()
  fmt.Println("Alert\t", alert)
  fmt.Println("Succe\t", resp.Success)
  fmt.Println("Error\t", resp.Error)
}

func PushToAndroid(message string) {
}

