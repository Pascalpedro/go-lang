// A function with multiple parameters...
// When you are working with multiple parameters, the function call must have the same number of arguments as there are parameters, and the arguments must be passed in the same order.

package main
import ("fmt")

func familyName(fname string, age int) {
  fmt.Println("Hello", age, "year old", fname, "Attama")
}

func main() {
  familyName("Pascal", 28)
  familyName("Juliet", 26)
  familyName("Bianca", 22)
}
