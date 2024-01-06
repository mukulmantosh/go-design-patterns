package main

import "fmt"

func main() {
	mag1, _ := newPublication("magazine", "Times", 50, "The Times")
	mag2, _ := newPublication("magazine", "Times", 50, "The Times")
	news1, _ := newPublication("newspaper", "Hindu", 50, "The Hindu")
	news2, _ := newPublication("newspaper", "Indian Express", 50, "The Indian Express")

	pubDetails(mag1)
	pubDetails(mag2)
	pubDetails(news1)
	pubDetails(news2)

}

func pubDetails(pub iPublication) {
	fmt.Printf("=============\n")
	fmt.Printf("%s\n", pub)
	fmt.Printf("Type: %T\n", pub)
	fmt.Printf("Name: %s\n", pub.getName())
	fmt.Printf("Pages : %d\n", pub.getPages())
	fmt.Printf("Publisher : %s\n", pub.getPublisher())
	fmt.Printf("=============\n")

}
