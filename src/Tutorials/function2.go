// Function call multiple times....

package main
import ("fmt")

func myMessage() {
  fmt.Println("Hei Pascal, I just got executed multiple times!!!")
}

func main() {
  myMessage()
  myMessage()
  myMessage()
}
