package subfolder1

import "fmt"

type CarSpecification interface {
	TopSpeed(speed string) string
	Cylinder(total string) string
	Detail()
}

type Car struct {
	Merk string
	Type string
	Year string
}

func (car *Car) TopSpeed(speed string) string {
	return speed
}

func (car *Car) Cylinder(total string) string {
	return total
}

func (car *Car) Detail() {
	fmt.Println("Merk : ", car.Merk)
	fmt.Println("Tahun Buatan : ", car.Year)
	fmt.Println("Tipe : ", car.Type)
}

// access modifier public
func Create(){
	fmt.Println("ini function public")
}

// access modifier private
func update(){
	fmt.Println("ini function private")
}
