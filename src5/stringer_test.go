package interfaces

import (
	"fmt"
	"testing"
	"time"
)

func Test_types_first(t *testing.T) {
	c := Customer{
		Name: "Bob",
		Address: Address{
			Line1: "The Street",
			Geo: Position{
				Lat:  52,
				Long: 0,
			},
		},
		Registered: time.Now(),
	}

	fmt.Printf("%v", c)
}

type (
	Customer struct {
		Name       string
		Address    Address
		Registered time.Time
	}
	Address struct {
		Line1 string
		Geo   Position
	}
	Position struct {
		Lat  Degree
		Long Degree
	}
	Degree float64
)

// func (d Degree) String() string {
// 	return fmt.Sprintf("%2.3f°", d)
// }

// func (p Position) String() string {
// 	return fmt.Sprintf("(%v, %v)", p.Lat, p.Long)
// }
