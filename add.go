package cal

import (
	"fmt"
	"time"
)

//Add is add 2 float values
func Add(x, y float64) float64 {
	fmt.Println("This is Add")
	fmt.Println(time.Now())
	return x + y
}
