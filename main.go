package main // определение пакета для текущего файла
import (
	"fmt"
	"strings"
)

// ********* placeholder *********
//

func WordCount(s string) map[string]int {

	var result = make(map[string]int)

	words := strings.Fields(s)
	for _, v := range words {
		result[v] = result[v] + 1
	}

	return result
}

func main() {
	fmt.Println(WordCount("one two three two four two one"))
}
