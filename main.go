package main

import (
	"fmt"
	"sync"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {

	tiquetes, err := tickets.CargaTiquetes("tickets.csv")
	if err != nil {
		fmt.Println("¡Error al cargar los tiquetes!", err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		total, err := tickets.TotalTiquetesPorPais(tiquetes, "China")
		if err != nil {
			fmt.Println("¡Error al calcular el total de tiquetes!", err)
			return
		}
		fmt.Printf("Total de tiquetes para el país de destino: %d\n", total)
	}()

	go func() {
		defer wg.Done()
		total, err := tickets.TotalTiquetesPorPeriodos(tiquetes, "madrugada")
		if err != nil {
			fmt.Println("¡Error al calcular el total de tiquetes!", err)
			return
		}
		fmt.Printf("Total de tiquetes para el período establecido: %d\n", total)
	}()

	go func() {
		defer wg.Done()
		porcentaje, err := tickets.PorcentajeTiquetesPorDestinoDia(tiquetes, "Brazil")
		if err != nil {
			fmt.Println("¡Error al calcular el total de tiquetes!", err)
			return
		}
		fmt.Printf("Porcentaje de tiquetes para el país de destino: %.2f%%\n", porcentaje)
	}()

	wg.Wait()
}
