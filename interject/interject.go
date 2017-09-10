/*
Copyright © 2017  <jmfgdev@outlook.com>
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

import "fmt"
import "strings"

func main() {
	var myResult string
	fmt.Println("---INTERJECT---")
	fmt.Println("What's the number one server OS in the world?")
	fmt.Printf("-->  ")
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
