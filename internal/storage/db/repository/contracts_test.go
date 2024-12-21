package repository_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jmoiron/sqlx"
	repo "github.com/ole-larsen/plutonium/internal/storage/db/repository"
	"github.com/ole-larsen/plutonium/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewContractsRepository(t *testing.T) {
	// Valid case
	db, _, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	tbl := "contracts"
	repository := repo.NewContractsRepository(sqlxDB, tbl)

	assert.NotNil(t, repository)
	assert.Equal(t, tbl, repository.TBL)

	// Nil database case
	repositoryNil := repo.NewContractsRepository(nil, "contracts")
	assert.Nil(t, repositoryNil)
}

func TestContractsRepository_InnerDB(t *testing.T) {
	// Valid case
	db, _, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repository := repo.NewContractsRepository(sqlxDB, "contracts")

	assert.Equal(t, sqlxDB, repository.InnerDB())

	// Nil receiver case
	var nilRepository *repo.ContractsRepository

	assert.Nil(t, nilRepository.InnerDB())
}

func TestContractsRepository_Ping(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true)) // Enable ping monitoring
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repository := repo.NewContractsRepository(sqlxDB, "contracts")

	mock.ExpectPing().WillReturnError(nil)

	err = repository.Ping()
	assert.NoError(t, err)

	mock.ExpectPing().WillReturnError(errors.New("ping error"))

	err = repository.Ping()
	assert.Error(t, err)

	// Nil receiver case
	var nilRepository *repo.ContractsRepository
	err = nilRepository.Ping()
	assert.Equal(t, repo.ErrDBNotInitialized, err)
}

func TestContractsRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repository := repo.NewContractsRepository(sqlxDB, "contracts")

	ctx := context.Background()
	contractMap := map[string]interface{}{
		"name":    "contract1",
		"address": "0x123",
		"tx":      "0x456",
		"abi":     "abi_data",
	}

	mock.ExpectExec(`INSERT INTO contracts \(name, address, tx, abi\).*`).
		WithArgs(contractMap["name"], contractMap["address"], contractMap["tx"], contractMap["abi"]).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repository.Create(ctx, contractMap)
	assert.NoError(t, err)

	mock.ExpectExec(`INSERT INTO contracts \(name, address, tx, abi\).*`).
		WithArgs(contractMap["name"], contractMap["address"], contractMap["tx"], contractMap["abi"]).
		WillReturnError(errors.New("insert error"))

	err = repository.Create(ctx, contractMap)
	assert.Error(t, err)

	// Nil receiver case
	var nilRepository *repo.ContractsRepository
	err = nilRepository.Create(ctx, contractMap)
	assert.Equal(t, repo.ErrDBNotInitialized, err)
}

func TestContractsRepository_GetOne(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repository := repo.NewContractsRepository(sqlxDB, "contracts")

	ctx := context.Background()
	name := "contract1"

	mock.ExpectQuery(`SELECT .* FROM contracts WHERE name=\$1`).
		WithArgs(name).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "address", "tx", "abi"}).
			AddRow(1, name, "0x123", "0x456", "abi_data"))

	contract, err := repository.GetOne(ctx, name)
	assert.NoError(t, err)
	assert.Equal(t, &models.Contract{
		ID:      1,
		Name:    name,
		Address: "0x123",
		Tx:      "0x456",
		Abi:     "abi_data",
	}, contract)

	mock.ExpectQuery(`SELECT .* FROM contracts WHERE name=\$1`).
		WithArgs(name).
		WillReturnError(sql.ErrNoRows)

	contract, err = repository.GetOne(ctx, name)
	assert.Error(t, err)
	assert.Nil(t, contract)

	// Nil receiver case
	var nilRepository *repo.ContractsRepository
	contract, err = nilRepository.GetOne(ctx, name)
	assert.Equal(t, repo.ErrDBNotInitialized, err)
	assert.Nil(t, contract)
}

func TestContractsRepository_GetByAddress(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repository := repo.NewContractsRepository(sqlxDB, "contracts")

	ctx := context.Background()
	address := "0x0000000000000000000000000000000000000123"

	mock.ExpectQuery(`SELECT .* FROM contracts WHERE address=\$1`).
		WithArgs(address).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "address", "tx", "abi"}).
			AddRow(1, "contract1", address, "0x456", "abi_data"))

	commonAddress := common.HexToAddress(address)
	contract, err := repository.GetByAddress(ctx, commonAddress)
	assert.NoError(t, err)
	assert.Equal(t, &models.Contract{
		ID:      1,
		Name:    "contract1",
		Address: address,
		Tx:      "0x456",
		Abi:     "abi_data",
	}, contract)

	mock.ExpectQuery(`SELECT .* FROM contracts WHERE address=\$1`).
		WithArgs(address).
		WillReturnError(sql.ErrNoRows)

	contract, err = repository.GetByAddress(ctx, commonAddress)
	assert.Error(t, err)
	assert.Nil(t, contract)

	var nilRepository *repo.ContractsRepository
	contract, err = nilRepository.GetByAddress(ctx, commonAddress)
	assert.Equal(t, repo.ErrDBNotInitialized, err)
	assert.Nil(t, contract)
}

func TestContractsRepository_GetCollectionsContracts(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repository := repo.NewContractsRepository(sqlxDB, "contracts")

	ctx := context.Background()

	mock.ExpectQuery(`SELECT c.id, c.name, c.address, c.tx, c.abi FROM contracts c WHERE c.name LIKE 'collection_%'`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "address", "tx", "abi"}).
			AddRow(1, "collection_1", "0x123", "0x456", "abi_data").
			AddRow(2, "collection_2", "0x789", "0xABC", "abi_data_2"))

	contracts, err := repository.GetCollectionsContracts(ctx)
	assert.NoError(t, err)
	assert.Len(t, contracts, 2)

	var nilRepository *repo.ContractsRepository
	contracts, err = nilRepository.GetCollectionsContracts(ctx)
	assert.Equal(t, repo.ErrDBNotInitialized, err)
	assert.Nil(t, contracts)
}

func TestContractsRepository_GetAuctions(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repository := repo.NewContractsRepository(sqlxDB, "contracts")

	ctx := context.Background()

	mock.ExpectQuery(`SELECT c.id, c.name, c.address, c.tx, c.abi FROM contracts c WHERE c.name LIKE 'auction_%'`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "address", "tx", "abi"}).
			AddRow(1, "auction_1", "0x123", "0x456", "abi_data").
			AddRow(2, "auction_2", "0x789", "0xABC", "abi_data_2"))

	contracts, err := repository.GetAuctions(ctx)
	assert.NoError(t, err)
	assert.Len(t, contracts, 2)

	var nilRepository *repo.ContractsRepository
	contracts, err = nilRepository.GetAuctions(ctx)
	assert.Equal(t, repo.ErrDBNotInitialized, err)
	assert.Nil(t, contracts)
}

func TestContractsRepository_EmptyGetByAddress(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repository := repo.NewContractsRepository(sqlxDB, "contracts")

	ctx := context.Background()
	address := common.HexToAddress("0x0000000000000000000000000000000000000123")

	mock.ExpectQuery(`SELECT .* FROM contracts WHERE address=\$1`).
		WithArgs(address.Hex()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "address", "tx", "abi"}))

	contract, err := repository.GetByAddress(ctx, address)
	assert.Error(t, err)
	assert.Nil(t, contract)

	var nilRepository *repo.ContractsRepository
	contract, err = nilRepository.GetByAddress(ctx, address)
	assert.Equal(t, repo.ErrDBNotInitialized, err)
	assert.Nil(t, contract)
}

func TestContractsRepository_CreateEdgeCases(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repository := repo.NewContractsRepository(sqlxDB, "contracts")

	ctx := context.Background()
	contractMap := map[string]interface{}{
		"name":    "", // Empty name (edge case)
		"address": nil,
		"tx":      nil,
		"abi":     nil,
	}

	mock.ExpectExec(`INSERT INTO contracts .*`).
		WithArgs(contractMap["name"], contractMap["address"], contractMap["tx"], contractMap["abi"]).
		WillReturnResult(sqlmock.NewResult(1, 0)) // Simulate no rows affected

	err = repository.Create(ctx, contractMap)
	assert.NoError(t, err)
}

func TestContractsRepository_ErrorHandling(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repository := repo.NewContractsRepository(sqlxDB, "contracts")

	ctx := context.Background()
	contractMap := map[string]interface{}{
		"name":    "contract1",
		"address": "0x123",
		"tx":      "0x456",
		"abi":     "abi_data",
	}

	// Simulate database error
	mock.ExpectExec(`INSERT INTO contracts .*`).
		WithArgs(contractMap["name"], contractMap["address"], contractMap["tx"], contractMap["abi"]).
		WillReturnError(errors.New("database error"))

	err = repository.Create(ctx, contractMap)
	assert.Error(t, err)
	assert.Equal(t, "database error", err.Error())
}

func TestContractsRepository_PingEdgeCase(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true)) // Enable ping monitoring
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repository := repo.NewContractsRepository(sqlxDB, "contracts")

	// Simulate ping timeout
	mock.ExpectPing().WillReturnError(errors.New("timeout error"))

	err = repository.Ping()
	assert.Error(t, err)
	assert.Equal(t, "timeout error", err.Error())
}

func TestContractsRepository_GetCollectionsContracts_NoResults(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repository := repo.NewContractsRepository(sqlxDB, "contracts")

	ctx := context.Background()

	mock.ExpectQuery(`SELECT c.id, c.name, c.address, c.tx, c.abi FROM contracts c WHERE c.name LIKE 'collection_%'`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "address", "tx", "abi"}))

	contracts, err := repository.GetCollectionsContracts(ctx)
	assert.NoError(t, err)
	assert.Len(t, contracts, 0)
}

func TestContractsRepository_GetAuctions_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repository := repo.NewContractsRepository(sqlxDB, "contracts")

	ctx := context.Background()

	mock.ExpectQuery(`SELECT .* FROM contracts c WHERE c.name like 'auction_%'`).
		WillReturnError(errors.New("query error"))

	contracts, err := repository.GetAuctions(ctx)
	assert.Error(t, err)
	assert.Nil(t, contracts)
}

func TestContractsRepository_GetContractsByType_ScanError(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repository := repo.NewContractsRepository(sqlxDB, "contracts")

	ctx := context.Background()
	contractType := "type1"

	// Mock the query to return invalid data that causes Scan to fail
	mock.ExpectQuery(`SELECT .* FROM contracts c WHERE c.name LIKE 'type1'`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "address", "tx", "abi"}).
			AddRow("invalid_id", "contract1", "0x123", "0x456", "abi_data")) // "invalid_id" will cause Scan to fail

	contracts, err := repository.GetContractsByType(ctx, contractType)

	// Assert that an error occurred and no contracts were returned
	assert.Error(t, err)
	assert.Nil(t, contracts)

	// Ensure all expectations on the mock were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestContractsRepository_GetByAddress_DefaultError(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repository := repo.NewContractsRepository(sqlxDB, "contracts")

	ctx := context.Background()
	address := common.HexToAddress("0x0000000000000000000000000000000000000123")

	// Simulate a generic error for the QueryRowContext Scan
	mock.ExpectQuery(`SELECT.*FROM contracts WHERE address=\$1`).
		WithArgs(address.String()).
		WillReturnError(errors.New("unexpected error"))

	contract, err := repository.GetByAddress(ctx, address)

	// Assert that the error is returned and contract is nil
	assert.Error(t, err)
	assert.EqualError(t, err, "unexpected error")
	assert.Nil(t, contract)

	// Ensure all expectations on the mock were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestContractsRepository_GetOne_DefaultError(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repository := repo.NewContractsRepository(sqlxDB, "contracts")

	ctx := context.Background()
	name := "contract1"

	// Simulate a generic error for the QueryRowContext Scan
	mock.ExpectQuery(`SELECT.*FROM contracts WHERE name=\$1`).
		WithArgs(name).
		WillReturnError(errors.New("unexpected error"))

	contract, err := repository.GetOne(ctx, name)

	// Assert that the error is returned and contract is nil
	assert.Error(t, err)
	assert.EqualError(t, err, "unexpected error")
	assert.Nil(t, contract)

	// Ensure all expectations on the mock were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

// TestContractsRepository_MigrateContext tests the MigrateContext method of UsersRepository.
func TestContractsRepository_MigrateContext(t *testing.T) {
	db, mock, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer db.Close()

	repository := &repo.ContractsRepository{
		DB:  *sqlxDB,
		TBL: "contracts",
	}

	ctx := context.Background()

	// Test successful migration
	mock.ExpectExec(`CREATE TABLE IF NOT EXISTS contracts`).WillReturnResult(sqlmock.NewResult(1, 1))

	err := repository.MigrateContext(ctx)
	assert.NoError(t, err, "MigrateContext() should not return an error")

	// Test migration with an error
	mock.ExpectExec(`CREATE TABLE IF NOT EXISTS contracts`).WillReturnError(errors.New("exec error"))

	err = repository.MigrateContext(ctx)
	assert.Error(t, err, "MigrateContext() should return an error")
	assert.Equal(t, "exec error", err.Error(), "MigrateContext() should return the correct error message")

	// Test case: repo is nil
	var nilRepo *repo.ContractsRepository
	err = nilRepo.MigrateContext(ctx)
	assert.Error(t, err, "MigrateContext() on nil repository should return error")
}
