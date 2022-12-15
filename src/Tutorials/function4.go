// A function with multiple parameters...

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
