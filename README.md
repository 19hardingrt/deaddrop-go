# deaddrop-go

A deaddrop utility written in Go. Put files in a database behind a password to be retrieved at a later date.

This is a part of the University of Wyoming's Secure Software Design Course (Spring 2023). This is the base repository to be forked and updated for various assignments. Alternative language versions are available in:
- [Javascript](https://github.com/andey-robins/deaddrop-js)
- [Rust](https://github.com/andey-robins/deaddrop-rs)

## Versioning

`deaddrop-go` is built with:
- go version go1.19.4 linux/amd64

## Usage

`go run main.go --help` for instructions

Then run `go run main.go -new -user <username here>` and you will be prompted to create the initial password.

## Database

Data gets stored into the local database file dd.db. This file will not by synched to git repos. Delete this file if you don't set up a user properly on the first go

## Logging Strategy

The idea behind my logging code was to keep track new users that are created, who messages are sent to, and when someone reads the messages they were sent
(not referring to a specific timestamp). I chose to omit any information regarding who sent the message to the following user in order to ensure privacy of the sender as well as the specific contents of the message.

## MAC Strategy

The idea behind my MAC strategy was to ensure that the sender authenticates that they are logged in to their proper user account before sending a message.
A MAC is then added to the database with the message and is checked when a user reads their messages. If the MAC does not match, it will alert the user. I have also added an additional flag for the -send flag where you need to specify both the sender and the recipient. The new sending usage is: 

go run main.go -send -to <recipient> -from <sender>