package hwd

import (
	"debug/pe"
	"fmt"
	"gorm.io/gorm"
	"nokowebapi/apis"
	"nokowebapi/globals"
	"nokowebapi/nokocore"
	"os"
)

func NewService() {
	var err error
	var stat os.FileInfo
	var file *pe.File
	var DB *gorm.DB
	nokocore.KeepVoid(err, stat, file, DB)

	config := globals.GetConfigGlobals[Config]()

	// check config.Exec is not empty, the file exists
	if config.Exec == "" {
		panic(fmt.Errorf("hwd.exec is empty"))
	}

	if stat, err = os.Stat(config.Exec); os.IsNotExist(err) {
		panic(fmt.Errorf("hwd.exec file not found, %w", err))
	}

	if stat.IsDir() {
		panic(fmt.Errorf("hwd.exec is a directory"))
	}

	fmt.Printf("Exec: %s\n", config.Exec)
	if file, err = pe.Open(config.Exec); err != nil {
		panic(fmt.Errorf("failed to open file, %w", err))
	}

	defer nokocore.NoErr(file.Close())

	if header, ok := file.OptionalHeader.(*pe.OptionalHeader32); ok {
		fmt.Printf("File: PE32\nMagic (Hex): %#04x\n", header.Magic)
	}

	if header, ok := file.OptionalHeader.(*pe.OptionalHeader64); ok {
		fmt.Printf("File: PE32+\nMagic (Hex): %#04x\n", header.Magic)
	}

	if DB, err = apis.SqliteOpenFile(config.DB, &gorm.Config{}); err != nil {
		panic(fmt.Errorf("failed to open database, %w", err))
	}
}
