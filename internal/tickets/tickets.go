package tickets

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	ID                                    int
	Nombre, Email, PaisDestino, HoraVuelo string
	Precio                                int
}

// Creé una función llamada CargaTiquetes que abre un archivo CSV especificado -para este caso "tickets.csv"- lee sus
// registros y los convierte en una lista de estructuras Ticket, que incluye detalles como ID, nombre, email, país
// de destino, hora del vuelo y precio del mismo.
// Retorna la lista de tickets y un error si ocurre algún problema durante la lectura del archivo.

func CargaTiquetes(nombreArchivo string) ([]Ticket, error) {
	archivo, err := os.Open(nombreArchivo)
	if err != nil {
		return nil, err
	}
	defer archivo.Close()

	var tiquetes []Ticket
	lector := csv.NewReader(archivo)
	for {
		registro, err := lector.Read()
		if err != nil {
			break
		}
		id, _ := strconv.Atoi(registro[0])
		precio, _ := strconv.Atoi(registro[5])

		tiquete := Ticket{
			ID:          id,
			Nombre:      registro[1],
			Email:       registro[2],
			PaisDestino: registro[3],
			HoraVuelo:   registro[4],
			Precio:      precio,
		}
		tiquetes = append(tiquetes, tiquete)
	}
	return tiquetes, nil
}

/* ---------------------------- Requerimiento # 1 --------------------------- */
// Una función que calcule cuántas personas viajan a un país determinado.
// Esta función recorre una lista de tiquetes y cuenta cuántos de ellos tienen un destino específico.
// Retorna el total de tiquetes que coinciden con el destino proporcionado y un error si ocurre algún problema.

func TotalTiquetesPorPais(tiquetes []Ticket, destino string) (int, error) {
	total := 0
	encontrado := false
	for _, tiquete := range tiquetes {
		if tiquete.PaisDestino == destino {
			total++
			encontrado = true
		}
	}
	if !encontrado {
		return 0, fmt.Errorf("No se encontró el país %s", destino)
	}
	return total, nil
}

/* ---------------------------- Requerimiento # 2 --------------------------- */
// Una o varias funciones que calculen cuántas personas viajan en madrugada (0 → 6), mañana (7 → 12),
// tarde (13 → 19), y noche (20 → 23).

const (
	Madrugada = "madrugada"
	Manana    = "mañana"
	Tarde     = "tarde"
	Noche     = "noche"
)

func TotalTiquetesPorPeriodos(tiquetes []Ticket, periodo string) (int, error) {
	total := 0
	for _, tiquete := range tiquetes {
		horaString := strings.Split(tiquete.HoraVuelo, ":")[0]
		hora, err := strconv.Atoi(horaString)
		if err != nil {
			return 0, fmt.Errorf("Error al convertir la hora del vuelo a entero: %v", err)
		}
		switch periodo {
		case Madrugada:
			if hora >= 0 && hora < 6 {
				total++
			}
		case Manana:
			if hora >= 6 && hora < 12 {
				total++
			}
		case Tarde:
			if hora >= 12 && hora < 19 {
				total++
			}
		case Noche:
			if hora >= 19 && hora <= 23 {
				total++
			}

		default:
			return 0, fmt.Errorf("Período no válido: %s", periodo)
		}
	}
	return total, nil
}

/* ---------------------------- Requerimiento #3 ---------------------------- */
// Calcular el porcentaje de personas que viajan a un país determinado en un día.

func PorcentajeTiquetesPorDestinoDia(tiquetes []Ticket, destino string) (float64, error) {
	total, err := TotalTiquetesPorPais(tiquetes, destino)
	if err != nil {
		return 0, err
	}
	return float64(total) / float64(len(tiquetes)) * 100, nil
}
