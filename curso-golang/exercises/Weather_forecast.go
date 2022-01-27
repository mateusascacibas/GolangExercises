//Package weather fala sobre o tempo.
package main

//CurrentCondition fala sobre a condicao de tempo atual.
var CurrentCondition string

//CurrentLocation fala sobre a localizao atual.
var CurrentLocation string

//Forecast fala sobre a localizacao atual e a condicao atual neste lugar.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
