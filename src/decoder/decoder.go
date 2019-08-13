package decoder
 
import (
	"strings"
	"unicode/utf8"
	BinaryConverter "../binary-converter"
	
)

/**
 This method decode the segments then mapped it into the header
 @var headers []string, messages map[string][]string
 @return decodedMessages []string
*/

func Run(headers []string, messages map[string][]string) []string{
	
	// Header Mapping
	var decodedMessages []string
	
	for _, header := range headers {
		
		var decodedMessage string
		var binaryCounter = 0

		var mapping = ""
		headerMapping := make(map[string]string)

		// Loop through header characters
		for _, letter := range string(header) {
			
			// Integer to binary conversion
			var binary = BinaryConverter.IntegerToBinary(binaryCounter, false)
			
			// Check for same length
			if(!strings.Contains(binary, "0") && len(mapping) == len(binary)) {
				
				var tmp = ""
				binary = strings.Repeat("0", utf8.RuneCountInString(string(binaryCounter)))
				
				for i := 0; i < len(mapping); i++ {
					tmp += "0"
				} 

				mapping = tmp + binary
				headerMapping[mapping] = string(letter)
				binaryCounter = 1

			} else {
				
				var tmp = ""
				
				if (len(binary) < len(mapping)) {
					
					for i := 0; i < len(mapping)-len(binary); i++ {
						tmp += "0"
					} 
				}
				
				mapping = tmp + binary
				headerMapping[mapping] = string(letter)
				binaryCounter++
			}
			
		}
		
		// Segment Decoding
		var concatenatedMessage = ""
		var counter = 0
		
		for _, message := range messages[header] {
			concatenatedMessage+=message
		}

		var tempBinaryLength = ""
		var tempSegment = ""
		
		for _, binary := range concatenatedMessage {
			
			// Get length of keys
			if counter < 3 {
				
				tempSegment = ""
				tempBinaryLength+=string(binary)
				counter++
			
				} else {
				
				if len(tempSegment) < BinaryConverter.BinaryToInteger(tempBinaryLength) {
					
					tempSegment+=string(binary)
				
				} else {
					
					decodedMessage+=headerMapping[tempSegment]
					tempSegment = ""
					tempSegment+=string(binary)
				}

				// Check End of Segment
				if len(tempSegment) == BinaryConverter.BinaryToInteger(tempBinaryLength) && !strings.Contains(tempSegment, "0") {
					
					tempBinaryLength = ""
					counter = 0
					tempSegment = ""
				}
			}
		}

		decodedMessages = append(decodedMessages,decodedMessage)
	}
	return decodedMessages
}