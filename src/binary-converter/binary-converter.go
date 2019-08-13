package binaryConverter
 
import (
	"strconv"
)

/**
 This method converts integer to binary
 @var integerInput int, prefixInput bool
 @return string binary
*/

func IntegerToBinary(integerInput int, prefixInput bool) string {
	
	i64 := int64(integerInput)

	if prefixInput {
		return "0b" + strconv.FormatInt(i64, 2)
	} else {
		return strconv.FormatInt(i64, 2)
	}
}

/**
 This method converts binary to integer
 @var binaryInput string
 @return int
*/

func BinaryToInteger(binaryInput string) int {

	result, _ := strconv.ParseInt(binaryInput, 2, 64)
	return int(result)
}