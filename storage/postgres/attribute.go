package postgres

import (
	"context"
	"errors"
	"fmt"
	pb "position_service/genproto/position_service"
	"position_service/pkg/helper"
	"position_service/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type attributeRepo struct {
	db *pgxpool.Pool
}

func NewAttributeRepo(db *pgxpool.Pool) storage.AttributeI {
	return &attributeRepo{
		db: db,
	}
}

func (r *attributeRepo) Create(ctx context.Context, entity *pb.CreateAttributeRequest) (id string, err error) {
	query := `
			INSERT INTO attribute (
				id,
				name,
				type
			) VALUES ($1, $2, $3)`
	id = uuid.NewString()

	_, err = r.db.Exec(context.Background(), query, id, entity.Name, entity.AttributeType)
	if err != nil {
		return "", fmt.Errorf("error while inserting profession err: %w", err)
	}

	return id, nil
}

func (r *attributeRepo) GetAll(ctx context.Context, req *pb.GetAllAttributeRequest) (*pb.GetAllAttributeResponse, error) {
	var (
		resp   pb.GetAllAttributeResponse
		err    error
		filter string
		params = make(map[string]interface{})
	)

	if req.Name != "" {
		filter += " AND name ILIKE '%' || :name || '%' "
		params["name"] = req.Name
	}

	countQuery := `SELECT count(1) FROM attribute WHERE true ` + filter

	q, arr := helper.ReplaceQueryParams(countQuery, params)
	err = r.db.QueryRow(ctx, q, arr...).Scan(&resp.Count)

	if err != nil {
		return nil, fmt.Errorf("error while scanning count %w", err)
	}

	query := `SELECT
			id,
			name,
			type
		FROM attribute WHERE true` + filter

	query += " LIMIT :limit OFFSET :offset"
	params["limit"] = req.Limit
	params["offset"] = req.Offset

	q, arr = helper.ReplaceQueryParams(query, params)
	rows, err := r.db.Query(ctx, q, arr...)
	if err != nil {
		return nil, fmt.Errorf("error while getting rows %w", err)
	}

	for rows.Next() {
		var attribute pb.Attribute

		err = rows.Scan(
			&attribute.Id,
			&attribute.Name,
			&attribute.AttributeType,
		)

		if err != nil {
			return nil, fmt.Errorf("error while scanning attribute err: %w", err)
		}

		resp.Attributes = append(resp.Attributes, &attribute)
	}

	return &resp, nil
}

func (r *attributeRepo) Get(id string) (*pb.Attribute, error) {
	var attribute pb.Attribute

	query := `
		SELECT	
			id,
			name, 
			type
		FROM attribute 
		WHERE id = $1
	`

	row := r.db.QueryRow(context.Background(), query, id)
	err := row.Scan(
		&attribute.Id,
		&attribute.Name,
		&attribute.AttributeType,
	)

	if err != nil {
		return nil, fmt.Errorf("error while Getting attribute err: %w", err)
	}

	return &attribute, nil
}

func (r *attributeRepo) Update(req *pb.Attribute) (string, error) {
	query := `
		UPDATE attribute SET
			name = $1,
			type = $2
		WHERE id = $3
	`

	result, err := r.db.Exec(context.Background(),
		query,
		req.Name,
		req.AttributeType,
		req.Id,
	)
	if err != nil {
		return "", fmt.Errorf("error while updating attribute err: %w", err)
	}

	if i := result.RowsAffected(); i == 0 {
		return "not found", nil
	}

	return "Updated", nil
}

func (r *attributeRepo) Delete(id string) (*pb.AttributeResult, error) {

	query := `DELETE from attribute
					WHERE id = $1`

	result, err := r.db.Exec(context.Background(), query, id)

	if err != nil {
		return nil, fmt.Errorf("error while deleting attribute err: %w", err)
	}

	if i := result.RowsAffected(); i == 0 {
		return nil, errors.New("not found")
	}

	return &pb.AttributeResult{
		Status: result.String(),
	}, nil
}
