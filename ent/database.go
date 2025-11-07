// Copyright (C) moneta. 2025-present.
//
// Created at 2025-07-29, by liasica

package ent

import (
	"context"
	"database/sql"
	"fmt"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/zerolog/log"

	"github.com/liasica/orbit/ent/configure"
	"github.com/liasica/orbit/ent/migrate"
)

var Database *Client

func Setup(dsn string, debug bool) {
	OpenDatabase(dsn, debug)
	autoMigrate()
	dataInitialization()
}

func OpenDatabase(dsn string, debug bool) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal().Err(err)
	}

	drv := entsql.OpenDB(dialect.Postgres, db)
	Database = NewClient(Driver(drv))
	if debug {
		Database = Database.Debug()
	}
}

func autoMigrate() {
	err := Database.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithGlobalUniqueID(true),
		migrate.WithForeignKeys(false),
	)
	if err != nil {
		log.Fatal().Err(err)
	}
}

type TxFunc func(tx *Tx) error

func WithTx(ctx context.Context, fn TxFunc) error {
	tx, err := Database.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			_ = tx.Rollback()
			panic(v)
		}
	}()
	if err = fn(tx); err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			err = fmt.Errorf("rolling back transaction: %w", txErr)
		}
		return err
	}
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}

func dataInitialization() {
	ctx := context.Background()

	if exists, _ := Database.Configure.Query().Where(configure.KeyEQ(configure.KeyGitlabMergeTargets)).Exist(ctx); !exists {
		_ = Database.Configure.Create().
			SetKey(configure.KeyGitlabMergeTargets).
			SetData([]byte(`["development","main","master","next"]`)).
			Exec(ctx)
	}
}
