package templates

func CliTemplate() []byte {
	return []byte(`// Code generated by Gormseed (gorms).

package main

import (
	"log"
	"os"

	"github.com/raspiantoro/gormseeder"
	"{{ .SeedModuleName }}"
	"github.com/raspiantoro/gormseeder/gormseed"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	var command *string

	if len(os.Args) > 1 {
		command = &os.Args[1]
	}

	dsn := "host={{ .Host }} user={{ .Username }} password={{ .Password }} dbname={{ .DbName }} port= {{ .Port }}"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln(err)
	}

	seeds := gormseed.Load(&seeds.Seeds{})

	seeder := gormseeder.New(db, seeds)

	if command != nil && *command == "rollback" {
		if err = seeder.Rollback(); err != nil {
			log.Fatalln(err)
		}
	} else {
		if err = seeder.Seed(); err != nil {
			log.Fatalln(err)
		}
	}
}	
`)
}