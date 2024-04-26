package customlog

import (
	"fmt"
	"log"
	stringutils "pokemon-game-api/pkgs/utils/string"
)

func WriteBorderedInfoLog(message string) {
	fmt.Printf("\n-=-=-=-=- [ %v ] -=-=-=-=-\n\n", message)
}

func WriteBorderedErrorLog(message string) {
	fmt.Printf("\n-=-=- [ %v ] -=-=-\n\n", message)
}

func WriteInfoRuningServerPathLog(port string) {
	fmt.Println(stringutils.Empty)
	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")
	fmt.Println("|                                            |")
	fmt.Printf("|   Server runing on http://localhost%s   |\n", port)
	fmt.Println("|                                            |")
	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")
	fmt.Println(stringutils.Empty)
}

func WriteGodotEnvFailPanicLog(key string) {
	log.Panicf("Can not get %v value from .env file\n", key)
}

func WriteMongoClientPanicLog(err string) {
	log.Panicln("Can not connect to mongoDb and initialize client. -", err)
}
