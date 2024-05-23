package main

import (
	"IAM/config"
)

func main() {
	config.InitConnection()
	defer config.DB.Close() // close when application was closed
}
