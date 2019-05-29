package main

import "fmt"
import "bufio"

func getInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(string)
	text, _ := reader.ReadString('\n')
	return text
}

func main() {
	fmt.Println("Web crawler.")
	i := getInput("Gimme a URL:")
}
