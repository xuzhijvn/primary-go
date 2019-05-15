// Sample program to show how to use an interface in Go.
package main

import (
	"fmt"
)

// notifier is an interface that defined notification
// type behavior.
type notifier interface {
	notify()
	notify2()
}

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements a method with a pointer receiver.
func (u *user) notify() {
	u.name = "GoneBoy"
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

func (u user) notify2() {
	u.name = "xuzhijun"
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

// main is the entry point for the application.
func main() {
	// Create a value of type User and send a notification.
	u := user{"User1", "User1@email.com"}

	u2 := &user{"User2", "User2@email.com"}

	sendNotification(&u)

	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)

	sendNotification2(u2)

	fmt.Printf("Sending user email to %s<%s>\n",
		u2.name,
		u2.email)

	sendNotification(u2)

	fmt.Printf("Sending user email to %s<%s>\n",
		u2.name,
		u2.email)
	// ./listing36.go:32: cannot use u (type user) as type
	//                     notifier in argument to sendNotification:
	//   user does not implement notifier
	//                          (notify method has pointer receiver)
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotification(n notifier) {
	n.notify()
}
func sendNotification2(n notifier) {
	n.notify2()
}
