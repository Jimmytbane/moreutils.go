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

import "fmt"
import "strings"

func main() {
	var input string
	var truth string
	fmt.Println("---NTKERNELBESTKERNEL---")
	fmt.Println("What is the best kernel of all time?")
	fmt.Printf("Input your response -->  ")
	fmt.Scanf("%s", &input)
	input = strings.ToLower(input)

	truth = "NT KERNEL BEST KERNEL\n,"

	if input == "nt" || input == "nt kernel" {
		fmt.Println(truth)
	} else {
		fmt.Printf("What are you talking about? %s", truth)
	}
}
