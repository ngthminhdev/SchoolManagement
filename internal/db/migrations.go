package db

import (
	"GolangBackend/internal/global"
	"GolangBackend/helper"
	"context"
	"log"
	"time"
)

type MigrationQuery struct {
	Name  string
	Query []string
}

var userTable MigrationQuery = MigrationQuery{
	Name: `User table`,
	Query: []string{
		`CREATE TABLE IF NOT EXISTS public.users (
      id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
      name TEXT,
      email TEXT UNIQUE NOT NULL,
      phone TEXT UNIQUE NOT NULL, 
      password TEXT UNIQUE NOT NULL, 
      gender SMALLINT DEFAULT 1,
      deleted BOOLEAN DEFAULT false,
      created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
      modified_at TIMESTAMPTZ NOT NULL DEFAULT now()
    );`,
		// "CREATE INDEX IF NOT EXISTS idx_user_phone ON public.users (email)",
		// "CREATE INDEX IF NOT EXISTS idx_user_email ON public.users (phone)",
	},
}

var migrations = []MigrationQuery{
	userTable,
}

func RunMigrations() {
	ctx, cancelConnect := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelConnect()

	for _, migration := range migrations {
		for _, query := range migration.Query {
			_, err := global.DB.Exec(ctx, query)
			if err != nil {
				log.Fatalf("Run migration at %s error: %v", migration.Name, err)
			}
		}
		helper.LogInfo("Successful run: %s", migration.Name)
	}
}
