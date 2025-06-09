package main

import (
	"church_consolidation/config"
	"church_consolidation/domain"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/xuri/excelize/v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.InitDb()
	log.Println("db start")
	log.Println(db)
	log.Println("db end")

	f, err := excelize.OpenFile("consolidation.xlsx")
	if err != nil {
		log.Fatal("Error abriendo el archivo:", err)
	}

	// Suponemos que los datos están en la primera hoja
	sheet := f.GetSheetName(0)
	rows, err := f.GetRows(sheet)
	if err != nil {
		log.Fatal("Error leyendo las filas:", err)
	}

	for i, row := range rows {
		if i == 0 {
			continue // saltamos el encabezado
		}

		// Cambia los índices según el orden de las columnas en tu archivo
		user := domain.Consolidation{
			FullName:         strings.TrimSpace(row[1]),
			Phone:            strings.TrimSpace(row[2]),
			Address:          strings.TrimSpace(row[3]),
			Age:              parseInt(strings.TrimSpace(row[4])),
			AttendsCellGroup: parseBool(strings.TrimSpace(row[5])),
			CallDay:          strings.TrimSpace(row[6]),
			CallTime:         strings.TrimSpace(row[7]),
			VisitDay:         strings.TrimSpace(row[8]),
			VisitTime:        strings.TrimSpace(row[9]),
			InvitedBy:        strings.TrimSpace(row[10]),
			Consolidator:     strings.TrimSpace(row[11]),
			DocumentType:     strings.TrimSpace(row[12]),
			DocumentNumber:   strings.TrimSpace(row[13]),
			MaritalStatus:    strings.TrimSpace(row[14]),
			Petition:         strings.TrimSpace(row[15]),

			// y así sucesivamente...

			// Ejemplo de fecha

		}

		// Assuming config.DB is undefined because it's not exported or a getter function is preferred.
		// The previous code attempted to use config.GetDB(), but the lint error indicates it's undefined.
		// Assuming config.DB is the exported GORM DB instance initialized by config.InitDb().
		err := db.Create(&user).Error
		if err != nil {
			fmt.Printf("Error guardando fila %d: %v\n", i+1, err)
		}
	}
	fmt.Println("Importación completada.")
}

func parseDate(s string) time.Time {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return time.Time{}
	}
	return t
}
func parseInt(s string) int {
	num, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return 0
	}
	return int(num)
}

func parseBool(s string) bool {
	if s == "Si" {
		return true
	}
	return false
}
