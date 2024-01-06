package main

import "fmt"

func main() {
	// TODO : Create a NotificationBuilder and use it to set properties.
	var buildData = newNotificationBuilder()

	// TODO : Use the builder to set some properties.
	buildData.SetTitle("New Notification")
	buildData.SetSubTitle("Inbox Received")
	buildData.SetIcon("icon.jpg")
	buildData.SetImage("image.jpg")
	buildData.SetPriority(5)
	buildData.SetMessage("This is a basic notification")
	buildData.SetType("alert")

	// TODO : Use the Build function to create a finished object.
	notify, err := buildData.Build()
	if err != nil {
		fmt.Printf("NOTIFICATION ERROR: %s", err)
	} else {
		fmt.Printf("! Notification: %+v", notify)
	}
}
