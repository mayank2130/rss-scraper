package main

import (
  "net/http"
  "log"
  "encoding/json"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {

  dat, err := json.Marshal(payload)

  if err != nil {
    log.Printf("Failed to marshal JSON response: %v",payload)
    w.WriteHeader(500)
    return
  }

  w.Header().Add("Content-Type", "application/json")
  w.WriteHeader(code)
  w.Write(dat)
}
