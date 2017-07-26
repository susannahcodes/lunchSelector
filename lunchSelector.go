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
   	"Monsoon",
        "Jack Brown's",
        "Tai Fresh",
        "Market Street Market",
        "Himalayan Fusion",
        "The Pie Chest",
        "The Whiskey Jar",
        "Omni Buffet",
        "Revolutionary Soup",
        "Christian's Pizza",
        "The Spot",
        "Pearl Island Catering",
}
n := rand.Int() % len(restaurants)
fmt.Print("Today's restaurant is: ", restaurants[n])
}
