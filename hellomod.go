package hellomod

import "fmt"

func Hello(name string) string {
  return fmt.Sprintf("Hello %s. Nice to see you!", name)
}
