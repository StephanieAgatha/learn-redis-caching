package main

import (
	"redishop/delivery"
	"redishop/util/helper"
)

func main() {
	helper.PrintAscii()
	delivery.NewServer().Run()
}
