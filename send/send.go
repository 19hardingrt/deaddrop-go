package send

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/andey-robins/deaddrop-go/db"
)

// SendMessage takes a destination username and will
// prompt the user for a message to send to that user
func SendMessage(to string) {
	if !db.UserExists(to) {
		log.Fatalf("Destination user does not exist")
	}

	message := getUserMessage()
	f, e := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if e != nil {
		log.Fatalf("Error opening file: %v", e)
	}
	if _, err := f.WriteString("Message sent to " + to + "\n"); err != nil {
		log.Println(err)
	}
	db.SaveMessage(message, to)
}

// getUserMessage prompts the user for the message to send
// and returns it
func getUserMessage() string {

	fmt.Println("Enter your message: ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	return text
}
