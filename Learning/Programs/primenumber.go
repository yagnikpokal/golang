// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math/big"
)

func main() {

	const n = 1212121
	if big.NewInt(n).ProbablyPrime(0) {
		fmt.Println(n, "is prime")
	} else {
		fmt.Println(n, "is not prime")
	}

}
