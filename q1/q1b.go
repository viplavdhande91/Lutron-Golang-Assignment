

package main
import "fmt"
import "strings" // Needed to use Split

func main() {
    str := "go bat cat gopher lang go toy go" 
    split_array := strings.Split(str, " ")
   
	frequency_count := make(map[string]int)

	//POPULATING MAP
	for i := 0; i < len(split_array); i++ {
		value := frequency_count[split_array[i]]

		frequency_count[split_array[i]] = value + 1

	}

	
	//ITERATNG OVER MAP
	fmt.Println("Word Frequency Count:", frequency_count)
	

}