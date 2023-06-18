package main
import (
  "github.com/bogdan-lytvynov/go-shell/internal/engine"
)


func main() {
  e := engine.New()
  e.Exec("echo Hello\n")
}
