package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/devpablocristo/ponti/pkg/words-suggestors/pg_trgm-gin"
)

func main() {
	// Leer configuración desde variables de entorno
	dsn := os.Getenv("DATABASE_DSN")
	table := os.Getenv("SUGGESTER_TABLE")
	column := os.Getenv("SUGGESTER_COLUMN")

	// Inicializar el Suggester
	s, err := pkgsuggester.NewSuggester(
		pkgsuggester.WithDSN(dsn),
		pkgsuggester.WithTable(table),
		pkgsuggester.WithColumn(column),
	)
	if err != nil {
		log.Fatalf("error al crear el suggester: %v", err)
	}
	defer s.Close()

	// Consulta de ejemplo
	query := "pab" // Cambia por el prefijo que quieras probar
	fmt.Printf("Buscando sugerencias para: %s\n", query)
	suggestions, err := s.Suggest(context.Background(), query)
	if err != nil {
		log.Fatalf("error en Suggest: %v", err)
	}

	// Imprimir resultados
	fmt.Println("Sugerencias:")
	for _, sug := range suggestions {
		fmt.Printf(" - ID: %d, Texto: %s\n", sug.ID, sug.Text)
	}
}
