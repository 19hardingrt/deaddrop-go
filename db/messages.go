package db

import (
	"fmt"
	"log"
)

type Messages struct {
	Sender  string
	Message string
	Hash    string
}

// GetMessagesForUser assumes that a user has already been
// authenticated through a call to session.Authenticate(user)
// and then returns all the messages stored for that user
func GetMessagesForUser(user string) []string {
	database := Connect().Db

	rows, err := database.Query(`
			SELECT (data) FROM Messages
			WHERE recipient = (
				SELECT id FROM Users WHERE user = ?
			)
	`, user)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	defer rows.Close()

	// marshall rows into an array
	messages := make([]string, 0)
	for rows.Next() {
		var message string
		err := rows.Scan(&message)
		if err != nil {
			log.Fatalf("unable to scan row")
		}
		messages = append(messages, message)
	}
	return messages
}

// saveMessage will process the transaction to place a message
// into the database
func SaveMessage(message, recipient, sender, hash string) {
	database := Connect().Db

	fmt.Println(hash)
	database.Exec(`
		INSERT INTO Messages (sender, recipient, data, hash)
		VALUES (
			(SELECT id FROM Users WHERE user = ?), 
			(SELECT id FROM Users WHERE user = ?), 
			?,
			?
		);
	`, sender, recipient, message, hash)

	log.Println(recipient + " received a message from " + sender)

}

func GetMessagesForUser2(user string) []Messages {
	database := Connect().Db

	rows, err := database.Query(`
		SELECT Messages.data, Users.user, Messages.hash
		FROM Messages
		INNER JOIN Users ON Messages.sender=Users.id
		AND Messages.recipient=(SELECT id FROM Users WHERE user = ?)
	`, user)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	defer rows.Close()

	messages := []Messages{}
	for rows.Next() {
		res := Messages{}

		if err := rows.Scan(&res.Message, &res.Sender, &res.Hash); err != nil {
			log.Fatalf("could not scan row: %v", err)
		}
		messages = append(messages, res)
	}
	return messages
}
