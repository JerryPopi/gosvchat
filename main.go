package main

import(
	// "fmt"
	// "os"
	"flag"
)

func main(){
	var server = flag.Bool("s", false, "Selects if program will run in server mode")
	var port = flag.String("p", "3000", "Sets server port")
	var host = flag.String("h", "localhost:3000", "<host:port>")
	var username = flag.String("n", "", "Username to be used in chat")
	flag.Parse()

	if *server {
		startServer(*port)
	} else {
		startClient(*username, *host)
		// createUI()
	}
}