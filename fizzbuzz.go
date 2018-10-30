package fizzbuzz

import (
	"bytes"
	"strconv"
)

func fizzbuzz() (result string){
	var buffer bytes.Buffer
	for i := 0; i <= 15; i++ {
		var partialResult string = strconv.Itoa(i)
			if i%15 == 0 {
      		partialResult = "fizzbuzz"
			} else if i%3 == 0 {
      		partialResult = "fizz"
			} else if i%5 == 0 {
      		partialResult = "buzz"
			}
			buffer.WriteString(partialResult)
			buffer.WriteString("\n")
	}
	sss
  result = buffer.String()
	return
}
