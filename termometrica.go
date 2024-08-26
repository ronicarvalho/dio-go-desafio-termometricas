package main

import "fmt"

func main() {

	const boiling_kelvin = 373.0
	const temperature_kelvin = 273.0

	boiling_celsius := boiling_kelvin - temperature_kelvin

	fmt.Println("O ponto de ebulição da água em kelvin é", boiling_kelvin, "K e em celsius é", boiling_celsius, "Cº")
}
