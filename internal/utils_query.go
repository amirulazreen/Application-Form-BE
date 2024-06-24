package internal

import (
	"context"
	"database/sql"
	"log"
)

func (form Form) Check(ctx context.Context, db *sql.DB, query Queries) (bool, error) {
	var nameExists bool
	err := db.QueryRowContext(ctx ,query.Check, form.Name).Scan(&nameExists)
	if err != nil {
		log.Println(err)
		return false, err
	}
	return nameExists, nil
}

func (form Form) Insert(ctx context.Context, db *sql.DB, query Queries) error {
	_, err := db.ExecContext(ctx,query.Insert, form.Name, form.Type, form.Bank, form.OpsYears, form.SSM, form.PrevGateway, form.ProdType, form.StoreType, form.Inventory, form.Reference, form.SocMedia, form.Litigation, form.Score)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}