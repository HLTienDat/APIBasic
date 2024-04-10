package main

import (
	transport "test3/module/item/transport"
)

func main() {
	transport := transport.NewTransport()
	transport.SetEmployeeRoutes()
	transport.SetAccountRoutes()
	transport.Run()
}
