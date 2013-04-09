package main

import (
	"arq/sr"
	"fmt"
)

func main() {
	queue := sr.NewQueue(8)

	fmt.Printf("%v\n\n", queue)

	queue.MarkAcknowledged(0)

	fmt.Printf("%v\n", queue)
}
