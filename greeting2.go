package greeting

import "fmt"

//Goodby function
func Goodby(name string) string {
	message := fmt.Sprintf("Goodby, %v !", name)
	return message

}
