package main

import "fmt"

// Define a magazine type and embed the publication interface.
type magazine struct {
	publication
}

// Define a Stringer interface that gives a string representation of the type
func (n magazine) String() string {
	return fmt.Sprintf("This is a magazine named %s", n.name)
}

// the createMagazine function returns a new Newspaper object
func createMagazine(name string, pages int, publisher string) iPublication {
	return &magazine{
		publication{
			name:      name,
			pages:     pages,
			publisher: publisher,
		},
	}
}
