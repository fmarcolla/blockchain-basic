package main

import "fmt";

func main(){

	blockChain := CreateBlockchain(3);

	blockChain.addBlock("user-a", "user-b", 100);
	blockChain.addBlock("user-b", "user-c", 200);
	
	fmt.Println(blockChain.isValid());
}