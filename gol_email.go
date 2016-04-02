package main

import "encoding/hex"
import "fmt"

func main() {
	fmt.Println(len(hex.EncodeToString([]byte("suyash@gmail.com"))))
}
