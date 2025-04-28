package main

import "fmt"

const (
	Red = iota // Red is 0
	Blue
	Green
)

type BitSize int

const (
	_          = iota
	KG BitSize = 1 << (10 * iota)
	MB BitSize = 1 << (10 * iota)
	GB BitSize = 1 << (10 * iota)
)

type M1 struct {
	seconds int
	ok      bool
}

func main() {
	// var highUint8 uint8 = 255
	// var highInt8 int8 = 127

	// fmt.Println("highInt8", highInt8)
	// fmt.Println("highUint8", highUint8)
	// fmt.Printf("%b \t \t \b %b \n", KG, KG)
	// fmt.Printf("%b \t \t \b %b", MB, MB)

	timeZone := make(map[string]M1, 2)

	keys := []string{"one", "two"}

	timeZone["one"] = M1{
		seconds: 10,
		ok:      true,
	}

	timeZone["two"] = M1{
		seconds: 100,
		ok:      false,
	}

	for key := range timeZone {
		seconds := timeZone[key]

		if seconds.ok == true {
			fmt.Println("[seconds]", seconds)
		}

	}

	for _, value := range keys {
		// The ok variable in the "comma ok" idiom when used with map lookups in Go indicates only whether the key exists in the map. It does not directly tell you anything about the value associated with that key (other than that a value exists for that key)

		if seconds, ok := timeZone[value]; ok && seconds.ok {
			fmt.Println("[seconds] xxxx", seconds)
		}

	}

}
