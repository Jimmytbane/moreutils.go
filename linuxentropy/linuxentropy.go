package main

import "os/exec"
import "fmt"

func main() {
	stdout, err := exec.Command("cat", "/proc/sys/kernel/random/entropy_avail").Output()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(stdout))
}