package main

import (
  "log"
  "net/http"
)

func main() {
  // ディレクトリを指定する
  fs := http.FileServer(http.Dir("static"))

  // ルーティングを設定する
  http.Handle("/", fs)
  log.Println("Listening...")

  // 8080ポートでサーバを起動する
  log.Fatal(http.ListenAndServe(":8080", nil))
}