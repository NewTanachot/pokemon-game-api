package customlog

import (
	"fmt"
	"log"
)

func WriteBorderedInfoLog(message string) {
	fmt.Printf("-=-=-=-=- [ %v ] -=-=-=-=-\n\n", message)
}

func WriteBorderedErrorLog(message string) {
	fmt.Printf("-=-=- [ %v ] -=-=-\n\n", message)
}

func WriteInfoRuningServerPathLog(port string) {
	fmt.Println("")
	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")
	fmt.Println("|                                            |")
	fmt.Printf("|   Server runing on http://localhost%s   |\n", port)
	fmt.Println("|                                            |")
	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")
	fmt.Println("")
}

func WriteFatalSetGodotEnvFailLog(key string) {
	log.Fatalf("Can not get %v value from .env file\n", key)
}
