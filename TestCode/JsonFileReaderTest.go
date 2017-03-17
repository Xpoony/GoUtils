package main

import (
	c "config"
	f "fmt"
)

func main() {
	var cc = c.GetObject()
	f.Println(cc.SearchValue("People", 0, "Name"))
	f.Println(cc.SearchValue("Server", "Ip"))
	f.Println(cc.SearchValue("System", "Info", "Name"))
}
