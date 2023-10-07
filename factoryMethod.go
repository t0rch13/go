package main

import (
	"fmt"
)


type Vehicle interface { // Vehicle is the interface for different types of vehicles
	Drive() string
}


type Car struct{} // Car is a concrete implementation of the Vehicle interface for cars

func (c *Car) Drive() string {
	return "Driving a car"
}


type Bike struct{} // Bike is a concrete implementation of the Vehicle interface for bikes

func (b *Bike) Drive() string {
	return "Riding a bike"
}


type VehicleFactory interface { // VehicleFactory is the interface for creating vehicles
	CreateVehicle() Vehicle
}


type CarFactory struct{} // CarFactory is a concrete implementation of the VehicleFactory interface for creating cars

func (cf *CarFactory) CreateVehicle() Vehicle {
	return &Car{}
}


type BikeFactory struct{} // BikeFactory is a concrete implementation of the VehicleFactory interface for creating bikes

func (bf *BikeFactory) CreateVehicle() Vehicle {
	return &Bike{}
}

func main() {
	
	carFactory := &CarFactory{}
	bikeFactory := &BikeFactory{}

	// Use the factory methods to create vehicles
	car := carFactory.CreateVehicle()
	bike := bikeFactory.CreateVehicle()

	// Drive the vehicles
	fmt.Println(car.Drive()) // Output: Driving a car
	fmt.Println(bike.Drive()) // Output: Riding a bike
}
