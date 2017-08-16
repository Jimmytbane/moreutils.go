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
	var my_result string
	fmt.Println("---INTERJECT---")
	fmt.Println("What's the number one server OS in the world?")
	fmt.Printf("-->  ")
	fmt.Scanln(&my_result)
	my_result = strings.ToLower(my_result)

	if (my_result == "linux") {
		fmt.Printf("I'd like to interject for a moment--\nwhat you've referred to as a %s is in fact, GNU/%s \nor, as I've recently taken to calling it, GNU + %s", my_result, my_result, my_result)
	} else if (my_result == "gnu/linux" || my_result == "gnu plus linux" || my_result == "gnu + linux" || my_result == "lignux") {
		fmt.Println("That's right!")
	} else {
		fmt.Printf("%s is not the number on server OS in the world...", my_result)
	}
}