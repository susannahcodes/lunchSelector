// hello-world.go
package main

import (
	"fmt"
	"math/rand"
        "time"
)

func main() {
	rand.Seed(time.Now().Unix())
	restaurants := []string{
   	 "Locked out",
   	 "Pipes broke",
         "Food poisoning",
         "Not feeling well",
}
n := rand.Int() % len(reasons)
fmt.Print("Today's restaurant is: ", reasons[n])
}
