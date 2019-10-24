package main

import "os/exec"
import "fmt"

func main() {
	exec.Command("javac", "Hello.java").Run()
	out, _ := exec.Command("java", "Hello").Output()
	fmt.Println(string(out))
}
