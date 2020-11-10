package main

import "fmt"

// iHumanResource interface
type iHumanResource interface {
	getID() uint64
	calculateSalary() float64
}

// employee exact humanResource
type employee struct {
	id     uint64
	salary float64
}

func (e *employee) getID() uint64 {
	return e.id
}

func (e *employee) calculateSalary() float64 {
	return e.salary
}

// department exact humanResource
type department struct {
	id    uint64
	subHr []iHumanResource
}

func (d *department) getID() uint64 {
	return d.id
}

func (d *department) calculateSalary() float64 {
	var totalSalary float64
	for _, hr := range d.subHr {
		totalSalary += hr.calculateSalary()
	}
	return totalSalary
}

func (d *department) addHumanResource(hr iHumanResource) {
	d.subHr = append(d.subHr, hr)
}

// main
func main() {
	var rootHr *department

	rootHr = &department{
		id: 1001,
	}

	rootHr.addHumanResource(&employee{
		1101,
		1500,
	})
	rootHr.addHumanResource(&employee{
		1102,
		2000,
	})
	rootHr.addHumanResource(&department{
		id: 10011,
		subHr: []iHumanResource{&employee{
			1103,
			1000,
		}, &employee{
			1104,
			3000,
		}},
	})

	res := rootHr.calculateSalary()

	fmt.Println(res)

}
