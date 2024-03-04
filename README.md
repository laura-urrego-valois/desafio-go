# Taller de código: Desafío Go

## Objetivo

### A continuación se plantea un desafío integrador que nos permitirá evaluar los temas que hemos visto hasta el momento.

## Planteo

### Una aerolínea pequeña tiene un sistema de reservas de pasajes a diferentes países. Este retorna un archivo con la información de los pasajes sacados en las últimas 24 horas. La aerolínea necesita un programa para extraer información de las ventas del día y, así, analizar las tendencias de compra.

### El archivo en cuestión es del tipo valores separados por coma (CSV), donde los campos están compuestos por: id, nombre, email, país de destino, hora del vuelo y precio.

## Requerimiento 1

### Una función que calcule cuántas personas viajan a un país determinado.

## Requerimiento 2

### Una o varias funciones que calculen cuántas personas viajan en madrugada (0 → 6), mañana (7 → 12), tarde (13 → 19), y noche (20 → 23).

## Requerimiento 3

### Calcular el porcentaje de personas que viajan a un país determinado en un día.

## Requerimiento 4

### Ejecutar al menos una vez cada requerimiento en la función main. Las ejecuciones deben realizarse de manera concurrente (utilizando diferentes goroutines).