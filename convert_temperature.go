package main

import "fmt"

const boilingPointKelvin float32 = 373

func main() {
	boilingPointCelsius := boilingPointKelvin - 273

	fmt.Printf("O ponto de ebulição da água em kelvin é aproximadamente %g°K e Celsius é aproximadamente %g°C", boilingPointKelvin, boilingPointCelsius)
}
