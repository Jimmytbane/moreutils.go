/*
Copyright © 2017  jmfgdev
This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published
by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.
You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
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
    }
}
