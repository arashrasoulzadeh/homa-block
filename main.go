package main

import "fmt"

func main() {

	blockchain := CreateBlockchain(2)

	blockchain.addBlock("Alice", "Bob", 5)
	blockchain.addBlock("John", "Bob", 2)

	fmt.Println(blockchain.chain)

	fmt.Println(blockchain.isValid())
}
