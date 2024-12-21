package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/go-multierror"
	"github.com/jmoiron/sqlx"
	"github.com/ole-larsen/plutonium/models"
)

type Contract struct {
	Created strfmt.Date `db:"created"`
	Updated strfmt.Date `db:"updated"`
	Deleted strfmt.Date `db:"deleted"`
	Name    string      `db:"name"`
	Address string      `db:"address"`
	Tx      string      `db:"tx"`
	ABI     string      `db:"abi"`
	ID      int64       `db:"id"`
}

type ContractsRepositoryInterface interface {
	InnerDB() *sqlx.DB
	Ping() error
	MigrateContext(ctx context.Context) error
	Create(ctx context.Context, contractMap map[string]interface{}) error
	GetOne(ctx context.Context, name string) (*models.Contract, error)
	GetByAddress(ctx context.Context, address common.Address) (*models.Contract, error)
	GetCollectionsContracts(ctx context.Context) ([]*models.Contract, error)
	GetAuctions(ctx context.Context) ([]*models.Contract, error)
}

// UsersRepository - repository to store users.
type ContractsRepository struct {
	DB  sqlx.DB
	TBL string
}

func NewContractsRepository(db *sqlx.DB, tbl string) *ContractsRepository {
	if db == nil {
		return nil
	}

	return &ContractsRepository{
		DB:  *db,
		TBL: tbl,
	}
}

func (r *ContractsRepository) InnerDB() *sqlx.DB {
	if r == nil {
		return nil
	}

	return &r.DB
}

func (r *ContractsRepository) Ping() error {
	if r == nil {
		return ErrDBNotInitialized
	}

	return r.DB.Ping()
}

func (r *ContractsRepository) MigrateContext(ctx context.Context) error {
	if r == nil {
		return ErrDBNotInitialized
	}

	_, err := r.DB.ExecContext(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
	id                     SERIAL PRIMARY KEY,
	name                   varchar(255) UNIQUE NOT NULL,
	address                varchar(255) UNIQUE NOT NULL,
	tx                     varchar(255) UNIQUE NOT NULL,
	abi                    text,
	created                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
	updated                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
	deleted                TIMESTAMP WITH TIME ZONE DEFAULT NULL
);`, r.TBL))

	return err
}

func (r *ContractsRepository) Create(ctx context.Context, contractMap map[string]interface{}) error {
	if r == nil {
		return ErrDBNotInitialized
	}

	_, err := r.DB.NamedExecContext(ctx, `
INSERT INTO contracts (name, address, tx, abi)
VALUES (:name, :address, :tx, :abi)
ON CONFLICT (name) 
DO 
UPDATE SET address = EXCLUDED.address, tx = EXCLUDED.tx, abi = EXCLUDED.abi;`, contractMap)

	return err
}

func (r *ContractsRepository) GetOne(ctx context.Context, name string) (*models.Contract, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var contract Contract

	sqlStatement := `
SELECT
	id,
	name,
	address,
	tx,
	abi
FROM contracts
WHERE name=$1;`

	row := r.DB.QueryRowContext(ctx, sqlStatement, name)

	err := row.Scan(
		&contract.ID, &contract.Name, &contract.Address, &contract.Tx, &contract.ABI)
	switch err {
	case sql.ErrNoRows:
		return nil, fmt.Errorf("contract not found")
	case nil:
		return &models.Contract{
			ID:      contract.ID,
			Name:    contract.Name,
			Address: contract.Address,
			Tx:      contract.Tx,
			Abi:     contract.ABI,
		}, nil
	default:
		return nil, err
	}
}

func (r *ContractsRepository) GetByAddress(ctx context.Context, address common.Address) (*models.Contract, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var contract Contract

	sqlStatement := `
SELECT
	id,
	name,
	address,
	tx,
	abi
FROM contracts
WHERE address=$1;`
	row := r.DB.QueryRowContext(ctx, sqlStatement, address.String())

	err := row.Scan(
		&contract.ID, &contract.Name, &contract.Address, &contract.Tx, &contract.ABI)
	switch err {
	case sql.ErrNoRows:
		return nil, fmt.Errorf("contract not found")
	case nil:
		return &models.Contract{
			ID:      contract.ID,
			Name:    contract.Name,
			Address: contract.Address,
			Tx:      contract.Tx,
			Abi:     contract.ABI,
		}, nil
	default:
		return nil, err
	}
}

func (r *ContractsRepository) GetCollectionsContracts(ctx context.Context) ([]*models.Contract, error) {
	return r.GetContractsByType(ctx, "collection_%")
}

func (r *ContractsRepository) GetAuctions(ctx context.Context) ([]*models.Contract, error) {
	return r.GetContractsByType(ctx, "auction_%")
}

func (r *ContractsRepository) GetContractsByType(ctx context.Context, contractType string) ([]*models.Contract, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var (
		multierr  multierror.Error
		contracts []*models.Contract
	)

	rows, err := r.DB.QueryxContext(ctx, fmt.Sprintf(`
SELECT 
    c.id,
	c.name,
   	c.address,
   	c.tx,
   	c.abi
FROM contracts c	
WHERE c.name LIKE '%s';`, contractType))

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var contract Contract

		err := rows.Scan(
			&contract.ID, &contract.Name, &contract.Address, &contract.Tx, &contract.ABI)
		if err != nil {
			return nil, err
		}

		contracts = append(contracts, &models.Contract{
			ID:      contract.ID,
			Name:    contract.Name,
			Address: contract.Address,
			Tx:      contract.Tx,
			Abi:     contract.ABI,
		})
	}

	defer rows.Close()

	return contracts, multierr.ErrorOrNil()
}
