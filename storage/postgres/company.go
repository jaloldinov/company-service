package postgres

import (
	"context"
	"errors"
	"fmt"
	cb "position_service/genproto/company_service"
	"position_service/pkg/helper"
	"position_service/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type companyRepo struct {
	db *pgxpool.Pool
}

func NewCompanyRepo(db *pgxpool.Pool) storage.CompanyI {
	return &companyRepo{
		db: db,
	}
}

func (r *companyRepo) Create(ctx context.Context, entity *cb.CreateCompany) (id string, err error) {
	query := `
		INSERT INTO company (
			id,
			name
		) 
		 VALUES ($1, $2)
	`

	id = uuid.NewString()

	_, err = r.db.Exec(
		ctx,
		query,
		id,
		entity.Name,
	)

	if err != nil {
		return "", fmt.Errorf("error while inserting company err: %w", err)
	}

	return id, nil
}

func (r *companyRepo) GetAll(ctx context.Context, req *cb.GetAllCompanyRequest) (*cb.GetAllCompanyResponse, error) {
	var (
		resp   cb.GetAllCompanyResponse
		err    error
		filter string
		params = make(map[string]interface{})
	)

	if req.Name != "" {
		filter += " AND name ILIKE '%' || :name || '%' "
		params["name"] = req.Name
	}

	countQuery := `
				SELECT count(1) FROM company WHERE true ` + filter

	q, arr := helper.ReplaceQueryParams(countQuery, params)
	err = r.db.QueryRow(ctx, q, arr...).Scan(&resp.Count)

	if err != nil {
		return nil, fmt.Errorf("error while getting company count err: %w", err)
	}

	query := `SELECT 
				id, 
				name
			FROM company WHERE true ` + filter

	query += " LIMIT :limit OFFSET :offset"
	params["limit"] = req.Limit
	params["offset"] = req.Offset

	q, arr = helper.ReplaceQueryParams(query, params)
	rows, err := r.db.Query(ctx, q, arr...)
	if err != nil {
		return nil, fmt.Errorf("error while getting company rows: %w", err)
	}

	for rows.Next() {
		var company cb.Company

		err = rows.Scan(
			&company.Id,
			&company.Name,
		)

		if err != nil {
			return nil, fmt.Errorf("error while scanning company rows: %w", err)
		}
		resp.Companies = append(resp.Companies, &company)
	}

	return &resp, nil
}

func (r *companyRepo) Get(ctx context.Context, id string) (*cb.Company, error) {
	var (
		company cb.Company
		err     error
	)

	query := `SELECT 
				id, 
				name
			FROM company WHERE id = $1`

	err = r.db.QueryRow(ctx, query, id).Scan(
		&company.Id,
		&company.Name,
	)

	if err != nil {
		return nil, fmt.Errorf("error while getting company err: %w", err)
	}

	return &company, nil
}

func (r *companyRepo) Update(ctx context.Context, entity *cb.Company) (string, error) {
	query := `
		UPDATE company SET 
			name = $1
		WHERE id = $2
	`

	result, err := r.db.Exec(
		ctx,
		query,
		entity.Name,
		entity.Id,
	)

	if err != nil {
		return "", fmt.Errorf("error while updating company err: %w", err)
	}

	if i := result.RowsAffected(); i == 0 {
		return "not found", nil
	}

	return "Updated", nil
}

func (r *companyRepo) Delete(ctx context.Context, id string) (*cb.CompanyResult, error) {

	query := `
		DELETE FROM company WHERE id = $1
	`

	result, err := r.db.Exec(ctx, query, id)

	if err != nil {
		return nil, errors.New("error while deleting company")
	}

	if i := result.RowsAffected(); i == 0 {
		return nil, errors.New("company not found")
	}

	return &cb.CompanyResult{
		Result: "Deleted",
	}, nil
}
