package helloworld

import "fmt"

func GetHello(name string) string {
	return fmt.Sprint("Hello ", name)
}
