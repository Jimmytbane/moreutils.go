/*
BSD 3-Clause License

Copyright (c) 2017, JimmyBot
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

* Neither the name of the copyright holder nor the names of its
  contributors may be used to endorse or promote products derived from
  this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package main

import(
  "strings"
  "fmt"
  "runtime"
  "os"
  "os/exec"
)

func whatami() {
		fmt.Printf("%s on %s\n", runtime.Version(), runtime.GOOS)
}

func ntkernelbestkernel() {
  var input string
	var truth string
	fmt.Println("What is the best kernel of all time?")
	fmt.Printf("-->  ")
	fmt.Scanf("%s\n", &input)
	input = strings.ToLower(input)

	truth = "NT KERNEL BEST KERNEL\n"

	if (input == "nt" || input == "nt kernel") {
		fmt.Println(truth)
	} else {
		fmt.Printf("What are you talking about? %s", truth)
	}
}

func downcasethis() {
	var input string
	var ignore string
	fmt.Printf("Input a string you would like to downcase -->  ")
	fmt.Scanf("%s", &input)
	fmt.Println("Your downcased string is --> " + strings.ToLower(input))
	fmt.Scanf("%s", &ignore)
}

func interject() {
	var myResult string
	fmt.Printf("What's the number one server OS in the world?\n-->  ")
	fmt.Scanln(&myResult)
	myResult = strings.ToLower(myResult)

	if myResult == "linux" {
		fmt.Printf("I'd like to interject for a moment--\nwhat you've referred to as a %s is in fact, GNU/%s \nor, as I've recently taken to calling it, GNU + %s\n", myResult, myResult, myResult)
	} else if myResult == "gnu/linux" || myResult == "gnu plus linux" || myResult == "gnu + linux" || myResult == "lignux" {
		fmt.Println("That's right!")
	} else {
		fmt.Printf("%s is not the number on server OS in the world...\n", myResult)
	}
}

func linuxentropy() {
	stdout, err := exec.Command("cat", "/proc/sys/kernel/random/entropy_avail").Output()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Your Linux entropy is -->  %s", string(stdout))
}


func main() {
  var another string
  var command string
  fmt.Printf("What command do you want to run?\n-->  ")
  fmt.Scanf("%s\n", &command)
  if (command == "whatami") {
    whatami()
  } else if (command == "ntkernelbestkernel") {
    ntkernelbestkernel()
  } else if (command == "interject") {
    interject()
  } else if (command == "downcasethis") {
    downcasethis()
  } else if (command == "linuxentropy") {
    linuxentropy()
  }
  fmt.Printf("Would you like to run another command?(command/n)\n-->  ")
  fmt.Scanf("%s\n", &another)
  if (another == "n" || another == "no") {
    os.Exit(0)
  } else if (another == "whatami") {
    whatami()
  } else if (another == "ntkernelbestkernel") {
    ntkernelbestkernel()
  } else if (another == "interject") {
    interject()
  } else if (another == "downcasethis") {
    downcasethis()
  } else if (another == "linuxentropy") {
    linuxentropy()
  }
}
