package cal

import (
	"fmt"
	"time"
	_ "github.com/ibmdb/go_ibm_db"
)

//Div is divide 2 float values
func Div(x, y float64) float64 {
	fmt.Println("This is Add")
	fmt.Println(time.Now())
	return x / y
}
