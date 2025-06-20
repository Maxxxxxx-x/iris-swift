// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlc

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createUserStmt, err = db.PrepareContext(ctx, createUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.deleteApiKeyByCreatorIdStmt, err = db.PrepareContext(ctx, deleteApiKeyByCreatorId); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteApiKeyByCreatorId: %w", err)
	}
	if q.deleteApiKeyByIdStmt, err = db.PrepareContext(ctx, deleteApiKeyById); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteApiKeyById: %w", err)
	}
	if q.deleteUserByEmailStmt, err = db.PrepareContext(ctx, deleteUserByEmail); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteUserByEmail: %w", err)
	}
	if q.deleteUserByIdStmt, err = db.PrepareContext(ctx, deleteUserById); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteUserById: %w", err)
	}
	if q.forceUserChangePasswordByEmailStmt, err = db.PrepareContext(ctx, forceUserChangePasswordByEmail); err != nil {
		return nil, fmt.Errorf("error preparing query ForceUserChangePasswordByEmail: %w", err)
	}
	if q.forceUserChangePasswordByIdStmt, err = db.PrepareContext(ctx, forceUserChangePasswordById); err != nil {
		return nil, fmt.Errorf("error preparing query ForceUserChangePasswordById: %w", err)
	}
	if q.getAllUsersStmt, err = db.PrepareContext(ctx, getAllUsers); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllUsers: %w", err)
	}
	if q.getApiKeyByCreatorIdStmt, err = db.PrepareContext(ctx, getApiKeyByCreatorId); err != nil {
		return nil, fmt.Errorf("error preparing query GetApiKeyByCreatorId: %w", err)
	}
	if q.getApiKeyByIdAndCreatorIdStmt, err = db.PrepareContext(ctx, getApiKeyByIdAndCreatorId); err != nil {
		return nil, fmt.Errorf("error preparing query GetApiKeyByIdAndCreatorId: %w", err)
	}
	if q.getApiKeyByKeyHashStmt, err = db.PrepareContext(ctx, getApiKeyByKeyHash); err != nil {
		return nil, fmt.Errorf("error preparing query GetApiKeyByKeyHash: %w", err)
	}
	if q.getApiKeyByNameAndCreatorIdStmt, err = db.PrepareContext(ctx, getApiKeyByNameAndCreatorId); err != nil {
		return nil, fmt.Errorf("error preparing query GetApiKeyByNameAndCreatorId: %w", err)
	}
	if q.getApiKeyUsageByKeyIdAndUserIdStmt, err = db.PrepareContext(ctx, getApiKeyUsageByKeyIdAndUserId); err != nil {
		return nil, fmt.Errorf("error preparing query GetApiKeyUsageByKeyIdAndUserId: %w", err)
	}
	if q.getApiKeysStmt, err = db.PrepareContext(ctx, getApiKeys); err != nil {
		return nil, fmt.Errorf("error preparing query GetApiKeys: %w", err)
	}
	if q.getUserByEmailStmt, err = db.PrepareContext(ctx, getUserByEmail); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserByEmail: %w", err)
	}
	if q.getUserByIdStmt, err = db.PrepareContext(ctx, getUserById); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserById: %w", err)
	}
	if q.refreshApiKeyStmt, err = db.PrepareContext(ctx, refreshApiKey); err != nil {
		return nil, fmt.Errorf("error preparing query RefreshApiKey: %w", err)
	}
	if q.saveApiKeyStmt, err = db.PrepareContext(ctx, saveApiKey); err != nil {
		return nil, fmt.Errorf("error preparing query SaveApiKey: %w", err)
	}
	if q.upatePasswordByEmailStmt, err = db.PrepareContext(ctx, upatePasswordByEmail); err != nil {
		return nil, fmt.Errorf("error preparing query UpatePasswordByEmail: %w", err)
	}
	if q.updatePasswordByIdStmt, err = db.PrepareContext(ctx, updatePasswordById); err != nil {
		return nil, fmt.Errorf("error preparing query UpdatePasswordById: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.deleteApiKeyByCreatorIdStmt != nil {
		if cerr := q.deleteApiKeyByCreatorIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteApiKeyByCreatorIdStmt: %w", cerr)
		}
	}
	if q.deleteApiKeyByIdStmt != nil {
		if cerr := q.deleteApiKeyByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteApiKeyByIdStmt: %w", cerr)
		}
	}
	if q.deleteUserByEmailStmt != nil {
		if cerr := q.deleteUserByEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUserByEmailStmt: %w", cerr)
		}
	}
	if q.deleteUserByIdStmt != nil {
		if cerr := q.deleteUserByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUserByIdStmt: %w", cerr)
		}
	}
	if q.forceUserChangePasswordByEmailStmt != nil {
		if cerr := q.forceUserChangePasswordByEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing forceUserChangePasswordByEmailStmt: %w", cerr)
		}
	}
	if q.forceUserChangePasswordByIdStmt != nil {
		if cerr := q.forceUserChangePasswordByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing forceUserChangePasswordByIdStmt: %w", cerr)
		}
	}
	if q.getAllUsersStmt != nil {
		if cerr := q.getAllUsersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllUsersStmt: %w", cerr)
		}
	}
	if q.getApiKeyByCreatorIdStmt != nil {
		if cerr := q.getApiKeyByCreatorIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getApiKeyByCreatorIdStmt: %w", cerr)
		}
	}
	if q.getApiKeyByIdAndCreatorIdStmt != nil {
		if cerr := q.getApiKeyByIdAndCreatorIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getApiKeyByIdAndCreatorIdStmt: %w", cerr)
		}
	}
	if q.getApiKeyByKeyHashStmt != nil {
		if cerr := q.getApiKeyByKeyHashStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getApiKeyByKeyHashStmt: %w", cerr)
		}
	}
	if q.getApiKeyByNameAndCreatorIdStmt != nil {
		if cerr := q.getApiKeyByNameAndCreatorIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getApiKeyByNameAndCreatorIdStmt: %w", cerr)
		}
	}
	if q.getApiKeyUsageByKeyIdAndUserIdStmt != nil {
		if cerr := q.getApiKeyUsageByKeyIdAndUserIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getApiKeyUsageByKeyIdAndUserIdStmt: %w", cerr)
		}
	}
	if q.getApiKeysStmt != nil {
		if cerr := q.getApiKeysStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getApiKeysStmt: %w", cerr)
		}
	}
	if q.getUserByEmailStmt != nil {
		if cerr := q.getUserByEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByEmailStmt: %w", cerr)
		}
	}
	if q.getUserByIdStmt != nil {
		if cerr := q.getUserByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByIdStmt: %w", cerr)
		}
	}
	if q.refreshApiKeyStmt != nil {
		if cerr := q.refreshApiKeyStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing refreshApiKeyStmt: %w", cerr)
		}
	}
	if q.saveApiKeyStmt != nil {
		if cerr := q.saveApiKeyStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing saveApiKeyStmt: %w", cerr)
		}
	}
	if q.upatePasswordByEmailStmt != nil {
		if cerr := q.upatePasswordByEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing upatePasswordByEmailStmt: %w", cerr)
		}
	}
	if q.updatePasswordByIdStmt != nil {
		if cerr := q.updatePasswordByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updatePasswordByIdStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                                 DBTX
	tx                                 *sql.Tx
	createUserStmt                     *sql.Stmt
	deleteApiKeyByCreatorIdStmt        *sql.Stmt
	deleteApiKeyByIdStmt               *sql.Stmt
	deleteUserByEmailStmt              *sql.Stmt
	deleteUserByIdStmt                 *sql.Stmt
	forceUserChangePasswordByEmailStmt *sql.Stmt
	forceUserChangePasswordByIdStmt    *sql.Stmt
	getAllUsersStmt                    *sql.Stmt
	getApiKeyByCreatorIdStmt           *sql.Stmt
	getApiKeyByIdAndCreatorIdStmt      *sql.Stmt
	getApiKeyByKeyHashStmt             *sql.Stmt
	getApiKeyByNameAndCreatorIdStmt    *sql.Stmt
	getApiKeyUsageByKeyIdAndUserIdStmt *sql.Stmt
	getApiKeysStmt                     *sql.Stmt
	getUserByEmailStmt                 *sql.Stmt
	getUserByIdStmt                    *sql.Stmt
	refreshApiKeyStmt                  *sql.Stmt
	saveApiKeyStmt                     *sql.Stmt
	upatePasswordByEmailStmt           *sql.Stmt
	updatePasswordByIdStmt             *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                                 tx,
		tx:                                 tx,
		createUserStmt:                     q.createUserStmt,
		deleteApiKeyByCreatorIdStmt:        q.deleteApiKeyByCreatorIdStmt,
		deleteApiKeyByIdStmt:               q.deleteApiKeyByIdStmt,
		deleteUserByEmailStmt:              q.deleteUserByEmailStmt,
		deleteUserByIdStmt:                 q.deleteUserByIdStmt,
		forceUserChangePasswordByEmailStmt: q.forceUserChangePasswordByEmailStmt,
		forceUserChangePasswordByIdStmt:    q.forceUserChangePasswordByIdStmt,
		getAllUsersStmt:                    q.getAllUsersStmt,
		getApiKeyByCreatorIdStmt:           q.getApiKeyByCreatorIdStmt,
		getApiKeyByIdAndCreatorIdStmt:      q.getApiKeyByIdAndCreatorIdStmt,
		getApiKeyByKeyHashStmt:             q.getApiKeyByKeyHashStmt,
		getApiKeyByNameAndCreatorIdStmt:    q.getApiKeyByNameAndCreatorIdStmt,
		getApiKeyUsageByKeyIdAndUserIdStmt: q.getApiKeyUsageByKeyIdAndUserIdStmt,
		getApiKeysStmt:                     q.getApiKeysStmt,
		getUserByEmailStmt:                 q.getUserByEmailStmt,
		getUserByIdStmt:                    q.getUserByIdStmt,
		refreshApiKeyStmt:                  q.refreshApiKeyStmt,
		saveApiKeyStmt:                     q.saveApiKeyStmt,
		upatePasswordByEmailStmt:           q.upatePasswordByEmailStmt,
		updatePasswordByIdStmt:             q.updatePasswordByIdStmt,
	}
}
