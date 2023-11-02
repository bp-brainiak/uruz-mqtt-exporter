package main

import "fmt"

func PrintSalutation() {
	fmt.Println("uruz-mqtt-exporter v1.0")
	fmt.Println("a simple mqtt prometheus exporter")
	fmt.Println("Copyright (c) 2023 Ulises Francisco Ruz Puga\n")

	licence := `This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.`
	fmt.Println(licence)
	fmt.Println()
}
