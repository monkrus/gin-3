package main

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
)

func main() {

  url := "http://localhost:8080/videos"
  method := "POST"
  // Straignten the payload line !
  payload := strings.NewReader("{\n  \"title\": \"title 7\",\n  \"description\" : \"desc 7\",\n  \"url\" : \"url 7\" \n}")

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
  }
  req.Header.Add("Authorization", "Basic cHJhZ21hdGljOnJldmlld3M=")
  req.Header.Add("Content-Type", "application/json")

  res, err := client.Do(req)
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)

  fmt.Println(string(body))
}