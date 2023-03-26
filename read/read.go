package read

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/andey-robins/deaddrop-go/db"
	"github.com/andey-robins/deaddrop-go/session"
)

func ReadMessages(user string) {
	if !db.UserExists(user) {
		log.Println("Cannot read messages from non-existent user: " + user + "\n")
		log.Fatalf("User not recognized")
	}

	err := session.Authenticate(user)
	if err != nil {
		log.Println("Cannot read messages with the wrong password for: " + user + "\n")
		log.Fatalf("Unable to authenticate user")
	}

	messages := db.GetMessagesForUser2(user)
	for _, message := range messages {
		if verify([]byte(message.Message), []byte(os.Getenv("KEY")), message.Hash) {
			fmt.Println(message.Sender + " sent: " + message.Message)
		} else {
			log.Println("MAC failure: '" + message.Message + "' from " + message.Sender + " cannot be verified!")
			fmt.Println("WARNING! Message cannot be authenticated!! " + message.Sender + " sent: " + message.Message)
		}
	}
}

func verify(msg, key []byte, hash string) bool {
	sig, err := hex.DecodeString(hash)
	if err != nil {
		return false
	}

	mac := hmac.New(sha256.New, key)
	mac.Write(msg)

	return hmac.Equal(sig, mac.Sum(nil))
}
