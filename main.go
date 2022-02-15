package main

import "app/tool"

func main() {
	tool.LoadEnv()
	tool.FirebaseInit()
}
