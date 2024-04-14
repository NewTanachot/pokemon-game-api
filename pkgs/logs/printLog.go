package logs

import "fmt"

func WriteInfoLog(message string) {
	fmt.Printf("-=-=-=-=- [ %v ] -=-=-=-=-\n\n", message)
}
