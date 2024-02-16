package main

import (
	"fmt"
	"programming-patterns/server"
	"programming-patterns/vehicle"
)

func main() {
	chevroletSpark := vehicle.NewVehicle()
	chevroletLacetti := vehicle.NewVehicle().
		WithModel("Lacetti").
		WithWarranty("5 years")

	fmt.Println(chevroletLacetti, chevroletSpark)

	s := server.NewServer(
		server.WithSSL,
		server.WithListenAddr(":4000"),
		server.WithLogLevel(server.LogLevelWarning),
	)
	s.HandleSomeStuff()
}
