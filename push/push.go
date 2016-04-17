package push

import (
  "fmt"
  "net/http"
  "net/url"
  "strings"
  "io/ioutil"
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
  client := &http.Client{}

  //data := url.Values{"registration_ids": "APA91bE8FHm5E7yy4qVAvmb6A8woQTUu7iMVR-zSKqNpJHCsDEy96LI-3fTXR5CNAdh5d0XX7BpSSIwMVMrlbCDal1S1i3Ah-0WzIxOxMWO-e4XAWsxtBdkkMGOsfF7fTH3SZxtG2hJAweGExTxQCTPhAxnojRtT8w", "data": {"message": "bar"}}  //data := url.Values{"data": {"message": message} }

  req, _ := http.NewRequest(
    "POST",
    "https://android.googleapis.com/gcm/send",
    strings.NewReader(data.Encode()),
  )
  req.Header.Set("Content-Type", "application/json")
  req.Header.Set("Authorization", "AIzaSyBgf-rHJQ8g2hC2dhpwkugiI6Judx68V24")

  resp, _ := client.Do(req)
  body, _ := ioutil.ReadAll(resp.Body)
  defer resp.Body.Close()

  println(string(body))
}

