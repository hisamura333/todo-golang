package main

import (
	"github.com/hisamura333/todo-golang/infra"
)

func main() {
	infra.Route()
	infra.DBInit()
}

