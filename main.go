package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/abdullah-alaadine/http-client/internal/commands"
)

func main() {
	if len(os.Args) < 2 {
		printHelpMessage()
		return
	}

	httpCommands := commands.GetHttpCommands()
	command := os.Args[1]

	if httpCommand, ok := httpCommands[command]; ok {
		args := os.Args[2:]
		httpResponse, err := httpCommand.Run(args)
		if err != nil {
			log.Fatal(err)
		}

		// Handle the HTTP response based on the command
		switch command {
		case "httphead":
			// For HEAD request, only print the response status
			fmt.Println(formatHeaders(httpResponse.Header))
		default:
			// For other requests, print the response body
			defer httpResponse.Body.Close()
			body, err := io.ReadAll(httpResponse.Body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(body))
		}
	} else {
		log.Fatalf("Invalid command: %s", command)
	}
}

func printHelpMessage() {
	httpCommands := commands.GetHttpCommands()

	helpMessage := `
Usage: http-client [command] [arguments]

Commands:`

	for _, httpCommand := range httpCommands {
		helpMessage += fmt.Sprintf("\n  %s\t%s", httpCommand.Name, httpCommand.Description)
	}

	helpMessage += `
	
Example:
  http-client httpget http://example.com
  http-client httppost http://example.com "body={key1: value, key2: value}"
`

	fmt.Println(helpMessage)
}

func formatHeaders(headers http.Header) string {
	var formattedHeaders strings.Builder

	for key, values := range headers {
		formattedHeaders.WriteString(fmt.Sprintf("%s: %s\n", key, strings.Join(values, ", ")))
	}

	return formattedHeaders.String()
}
