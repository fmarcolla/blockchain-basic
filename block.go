package main

import (
	"time"
	"encoding/json"
	"strconv"
	"crypto/sha256"
	"fmt"
	"strings"
)

type Block struct {
	data map[string]interface{}
	hash string
	previousHash string
	timestamp time.Time
	pow int
}

func(b Block) calculateHash() string {
	data, _ := json.Marshal(b.data);
	blockData := b.previousHash + string(data) + b.timestamp.String() + strconv.Itoa(b.pow);
	blockHash := sha256.Sum256([]byte(blockData));

	return fmt.Sprintf("%x", blockHash);
}

func(b *Block) mine(difficult int){
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficult)) {
		b.pow++;
		b.hash = b.calculateHash();
	}
}