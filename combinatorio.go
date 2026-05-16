package main

import (
	"fmt"
)

// Producto representa un artículo del catálogo de la cafetería.
type Producto struct {
	Nombre      string
	Precio      float64
	Categoria   string
	Descripcion string
}

// catalogo contiene todos los productos disponibles.
var catalogo = []Producto{
	{Nombre: "Agua mineral", Precio: 20.00, Categoria: "Bebida", Descripcion: "Botella 600ml sin gas"},
	{Nombre: "Té verde", Precio: 30.00, Categoria: "Bebida", Descripcion: "Infusión caliente"},
	{Nombre: "Café Americano", Precio: 35.00, Categoria: "Bebida", Descripcion: "Café negro doble shot"},
	{Nombre: "Jugo de naranja", Precio: 45.00, Categoria: "Bebida", Descripcion: "Natural exprimido 355ml"},
	{Nombre: "Pastel de chocolate", Precio: 55.00, Categoria: "Postre", Descripcion: "Rebanada individual 120g"},
	{Nombre: "Burrito de res", Precio: 75.00, Categoria: "Comida", Descripcion: "Tortilla, carne, frijoles"},
	{Nombre: "Sandwich de pollo", Precio: 85.00, Categoria: "Comida", Descripcion: "Pan integral, pollo, verduras"},
	{Nombre: "Ensalada César", Precio: 95.00, Categoria: "Comida", Descripcion: "Lechuga romana, crutones"},
	{Nombre: "Pizza personal", Precio: 110.00, Categoria: "Comida", Descripcion: "4 rebanadas, queso y jitomate"},
	{Nombre: "Combo del día", Precio: 130.00, Categoria: "Combo", Descripcion: "Plato fuerte + bebida + postre"},
}

// encontrarCombinaciones usa backtracking para generar todas las combinaciones
func encontrarCombinaciones(productos []Producto, presupuesto float64) [][]Producto {
	var resultado [][]Producto

	// función recursiva interna
	var backtrack func(inicio int, actual []Producto, total float64)
	backtrack = func(inicio int, actual []Producto, total float64) {
		// Guardamos la combinación actual si no excede el presupuesto
		if total <= presupuesto && len(actual) > 0 {
			// copiamos la combinación para evitar referencias
			combi := make([]Producto, len(actual))
			copy(combi, actual)
			resultado = append(resultado, combi)
		}

		// iteramos sobre los productos restantes
		for i := inicio; i < len(productos); i++ {
			nuevoTotal := total + productos[i].Precio
			if nuevoTotal <= presupuesto {
				backtrack(i+1, append(actual, productos[i]), nuevoTotal)
			}
		}
	}

	backtrack(0, []Producto{}, 0)
	return resultado
}

// imprimirResultados muestra el resumen y detalle de las combinaciones
func imprimirResultados(combis [][]Producto, presupuesto float64) {
	fmt.Printf("\nPresupuesto: $%.2f\n", presupuesto)
	fmt.Println("Total de combinaciones:", len(combis))

	// Agrupación por cantidad de productos
	agrupacion := make(map[int]int)
	var mejor []Producto
	var mejorTotal float64

	for _, c := range combis {
		agrupacion[len(c)]++
		var suma float64
		for _, p := range c {
			suma += p.Precio
		}
		if suma > mejorTotal {
			mejorTotal = suma
			mejor = c
		}
	}

	fmt.Println("\nPor cantidad de productos:")
	for k, v := range agrupacion {
		fmt.Printf("  %d producto(s): %d combinación(es)\n", k, v)
	}

	// Detalle de cada combinación
	for i, c := range combis {
		var suma float64
		for _, p := range c {
			suma += p.Precio
		}
		fmt.Printf("\n[%d] %d producto(s) — Total: $%.2f — Cambio: $%.2f\n", i+1, len(c), suma, presupuesto-suma)
		for _, p := range c {
			fmt.Printf("     • %s  $%.2f\n", p.Nombre, p.Precio)
		}
	}

	// Mejor combinación
	fmt.Println("\nMejor combinación (mayor gasto):")
	var suma float64
	for _, p := range mejor {
		fmt.Printf("     • %s  $%.2f\n", p.Nombre, p.Precio)
		suma += p.Precio
	}
	fmt.Printf("     Total: $%.2f  Cambio: $%.2f\n", suma, presupuesto-suma)
}

func main() {
	var presupuesto float64
	fmt.Print("Ingresa tu presupuesto: $")
	fmt.Scan(&presupuesto)

	combis := encontrarCombinaciones(catalogo, presupuesto)
	imprimirResultados(combis, presupuesto)
}
