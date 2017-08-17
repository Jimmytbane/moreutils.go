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

// import "os/exec"

func main() {
	var username string
	var ipOrDomain string

	fmt.Println("---SSHME---")
	fmt.Println("What username@server do you want to login to?")
	fmt.Printf("Put a server ip or domain here -->  ")
	fmt.Scanf("%s", &ipOrDomain)
	fmt.Printf("Put a username here -->  ")
	fmt.Scanf("%s", &username)
	fmt.Printf("About to run ssh %s@%s", username, ipOrDomain)
}
