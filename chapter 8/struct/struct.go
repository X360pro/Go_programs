package mystruct

type Subscriber struct {
	Name string
	Rate float64
	Active bool
	Home Address
}

type Employee struct {
	Name string
	Salary float64
	Home Address
}
type Address struct {
	Street string
	Pincode int
}