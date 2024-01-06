package main

import "fmt"

type NotificationBuilder struct {
	Title    string
	SubTitle string
	Message  string
	Image    string
	Icon     string
	Priority int
	NotType  string
}

func newNotificationBuilder() *NotificationBuilder {
	return &NotificationBuilder{}
}

func (nb *NotificationBuilder) SetTitle(title string) {
	nb.Title = title
}

func (nb *NotificationBuilder) SetSubTitle(subtitle string) {
	nb.SubTitle = subtitle
}

func (nb *NotificationBuilder) SetMessage(message string) {
	nb.Message = message
}

func (nb *NotificationBuilder) SetImage(image string) {
	nb.Image = image
}

func (nb *NotificationBuilder) SetIcon(icon string) {
	nb.Icon = icon
}

func (nb *NotificationBuilder) SetPriority(priority int) {
	nb.Priority = priority
}

func (nb *NotificationBuilder) SetType(notType string) {
	nb.NotType = notType
}

// The build method returns a fully finished Notification object

func (nb *NotificationBuilder) Build() (*Notification, error) {
	// error checking at build stage
	if nb.Icon != "" && nb.SubTitle == "" {
		return nil, fmt.Errorf("you need to specify a subtitle when using an icon")
	}

	if nb.Priority > 5 {
		return nil, fmt.Errorf("priority must be 0 to 5")
	}

	// return a newly created Notification object.
	return &Notification{
		title:    nb.Title,
		subtitle: nb.SubTitle,
		message:  nb.Message,
		icon:     nb.Icon,
		priority: nb.Priority,
		notType:  nb.NotType,
	}, nil

}
