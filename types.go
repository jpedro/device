package main

import (
    "time"
)

type Device struct {
    Serial  string `json:"serial"`
    Model   string `json:"model"`
    Version string `json:"version"`
    UUID    string `json:"uuid"`
}

type Message struct {
    Content string    `json:"content"`
    Created time.Time `json:"created"`
}
