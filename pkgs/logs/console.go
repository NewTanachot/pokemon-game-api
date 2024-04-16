package logs

import "fmt"

func WriteBorderedInfoLog(message string) {
	fmt.Printf("-=-=-=-=- [ %v ] -=-=-=-=-\n\n", message)
}
