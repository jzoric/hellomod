package hellomod

import "fmt"

func Hello(name string, department string) string {
  return fmt.Sprintf("Hello %s, from %s!", name, department)
}
