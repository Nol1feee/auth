package main

import (
	"context"
	"github.com/Nol1feee/CLI-chat/auth/internal/app"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app | %s\n", err)
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app | %s\n", err)
	}
}
