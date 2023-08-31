package play

import "fmt"

func Hello(name string) string {
	return fmt.Sprintf("Hello %s", name)
} 

func Bye() string {
	return "Bye Bye"
}

func Ask() string {
	return "Can you introduce yourself ?"
}