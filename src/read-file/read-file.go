package readFile

import (
	"bufio"
	"log"
	"os"
	"strings"
)

/**
 This method reads text file then mapped the headers and messages accordingly
 @var textFileLocation string
 @return messages map[string][]string, headers []string
*/

func TextFile(textFileLocation string) (map[string][]string, []string) {

	file, err := os.Open(textFileLocation)
 
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
 
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var messages []string
	var headers []string

	messageMapping := make(map[string][]string)
 
	for scanner.Scan() {
		
		// Get messages
		if strings.Contains(scanner.Text(), "0") || strings.Contains(scanner.Text(), "1") {
			
			messages = append(messages, scanner.Text())	
			
		} else {
			
			// Get Headers
			if len(headers) > 0 {
				messageMapping[headers[len(headers)-1]] = messages
				messages = nil
			}

			headers = append(headers, scanner.Text())
			
		}
	}

	messageMapping[headers[len(headers)-1]] = messages
 
	file.Close()

	return messageMapping, headers
}