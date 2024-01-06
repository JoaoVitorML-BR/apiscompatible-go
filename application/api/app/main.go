package main

import (
	"login/infra/server"
)

// USED for gen a key to used on .env

// func init(){
// 	key := make([]byte, 64)

// 	if _, err := rand.Read(key); err != nil{
// 		log.Fatal(err)
// 	}

// 	keyBased64 := base64.StdEncoding.EncodeToString(key)
// 	fmt.Print(keyBased64, "\n")
// }

func main() {

	// init server
	server.Server();
}