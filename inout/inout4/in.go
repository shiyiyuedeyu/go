package main
import "fmt"
import "bufio"
import "os"

var str string

func main() {
	for {
		fmt.Println("Enter : ")
		inputReader := bufio.NewReader(os.Stdin)
		input, err := inputReader.ReadString('\n')
		if err == nil {
			fmt.Printf("The input is : %s\n", input)
		} else {
			fmt.Println("shibai!!!!!!!!!")
		}
	}
}
