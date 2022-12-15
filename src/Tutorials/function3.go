// function with parameter(which acts as variables inside a function)..

package main
import ("fmt")

func familyName(fname string) {
  fmt.Println("Hello", fname, "Attama")
}

func main() {
  familyName("Pascal")
  familyName("Pedro")
  familyName("Emeka")
}
