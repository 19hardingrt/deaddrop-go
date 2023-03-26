package send

import (
	"bufio"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/andey-robins/deaddrop-go/db"
	"github.com/andey-robins/deaddrop-go/session"
)

// SendMessage takes a destination username and will
// prompt the user for a message to send to that user
func SendMessage(to, from string) {
	if !db.NoUsers() && !db.UserExists(from) {
		log.Fatalf("Sender not recognized")
	}
	err := session.Authenticate(from)
	if err != nil {
		log.Println("Recipient: " + to + "is not a valid user\n")
		log.Fatalf("Unable to authenticate user")
	}

	if !db.UserExists(to) {
		log.Println("User: " + from + " attempted to send a message with an incorrect password")
		log.Fatalf("Destination user does not exist")
	}

	message := getUserMessage()

	h := hmac.New(sha256.New, []byte(os.Getenv("KEY")))
	_, err = h.Write([]byte(message))
	if err != nil {
		log.Println("MAC could not be written")
		log.Fatalf("MAC could not be written")
	}

	hash := hex.EncodeToString(h.Sum(nil))

	db.SaveMessage(message, to, from, hash)
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
