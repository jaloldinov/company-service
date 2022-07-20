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

type positionRepo struct {
	db *pgxpool.Pool
}

func NewPositionRepo(db *pgxpool.Pool) storage.PositionI {

	return &positionRepo{
		db: db,
	}
}

func (r *positionRepo) Create(ctx context.Context, entity *pb.CreatePositionRequest) (position_id string, err error) {

	position_id = uuid.NewString()

	query := `INSERT INTO position (
		id,
		name,
		profession_id,
		company_id
	) VALUES ($1, $2, $3, $4)`

	_, err = r.db.Exec(ctx, query, position_id, entity.Name, entity.ProfessionId, entity.CompanyId)
	if err != nil {
		return "", fmt.Errorf("error while inserting position err: %w", err)
	}

	query = `INSERT INTO position_attributes (
		id,
		attribute_id,
		position_id,
		value
	) 
	VALUES ($1, $2, $3, $4) `

	positionAttributes := entity.PositionAttributes

	for key := range positionAttributes {
		attribute_id, err := uuid.NewRandom()
		if err != nil {
			return "", err
		}
		_, err = r.db.Exec(ctx, query, attribute_id, positionAttributes[key].AttributeId, position_id, positionAttributes[key].Value)

		if err != nil {
			return "", fmt.Errorf("error while inserting position_attributes err: %w", err)
		}
	}

	return position_id, nil
}

func (r *positionRepo) GetAll(ctx context.Context, req *pb.GetAllPositionRequest) (*pb.GetAllPositionResponse, error) {
	var (
		filter    string
		err       error
		params    = make(map[string]interface{})
		count     int32
		positions []*pb.GetPositionResponse
	)

	fmt.Println(len(req.Name))

	// if len(req.Name) != 0 {
	if req.Name != "" {
		filter += " AND position.name ILIKE '%' || :name || '%' "
		params["name"] = req.Name
	}

	if req.ProfessionId != "" {
		filter += " AND position.profession_id = :profession_id"
		params["profession_id"] = req.ProfessionId
	}

	if req.CompanyId != "" {
		filter += " AND position.company_id = :company_id"
		params["company_id"] = req.CompanyId
	}

	countQuery := `SELECT count(1) FROM position where true ` + filter

	q, arr := helper.ReplaceQueryParams(countQuery, params)
	err = r.db.QueryRow(ctx, q, arr...).Scan(&count)

	if err != nil {
		return nil, fmt.Errorf("error while positoin count %w", err)
	}

	query := `SELECT
			position.id,
			position.name,
			position.profession_id,
			position.company_id,
			position_attributes.id,
			position_attributes.attribute_id,
			position_attributes.value,
			attribute.id,
			attribute.name,
			attribute.type
		FROM position
		LEFT JOIN
		position_attributes
		ON
		position.id = position_attributes.position_id
		LEFT JOIN
		attribute
		ON
		position_attributes.attribute_id = attribute.id
		WHERE true ` + filter

	query += " LIMIT :limit OFFSET :offset"
	params["limit"] = req.Limit
	params["offset"] = req.Offset

	q, arr = helper.ReplaceQueryParams(query, params)
	rows, err := r.db.Query(ctx, q, arr...)
	if err != nil {
		return nil, fmt.Errorf("error while getting  positions rows %w", err)
	}

	for rows.Next() {
		var (
			position          pb.GetPositionResponse
			positionAttribute pb.GetPositionAttribute
			attribute         pb.Attribute
		)

		err = rows.Scan(
			&position.Id,
			&position.Name,
			&position.ProfessionId,
			&position.CompanyId,
			&positionAttribute.Id,
			&positionAttribute.AttributeId,
			&positionAttribute.Value,
			&attribute.Id,
			&attribute.Name,
			&attribute.AttributeType,
		)

		if err != nil {
			return nil, fmt.Errorf("error while scanning all: %w", err)
		}

		position.PositionAttributes = append(position.PositionAttributes,
			&pb.GetPositionAttribute{
				Id:          positionAttribute.Id,
				AttributeId: positionAttribute.AttributeId,
				Value:       positionAttribute.Value,
				Attribute:   &attribute,
			})
		positions = append(positions, &position)
	}

	return &pb.GetAllPositionResponse{
		Positions: positions,
		Count:     count,
	}, nil

}

func (r *positionRepo) Get(ctx context.Context, id string) (*pb.GetPositionResponse, error) {
	var position pb.GetPositionResponse

	queryPosition := `
				SELECT 
					id,
					name,
					profession_id,
					company_id
				FROM
					position
				WHERE
					id = $1`
	row := r.db.QueryRow(ctx, queryPosition, id)
	err := row.Scan(
		&position.Id,
		&position.Name,
		&position.ProfessionId,
		&position.CompanyId,
	)

	if err != nil {
		return nil, err
	}

	queryPositionAttributes := `
				SELECT
					pa.id,
					pa.attribute_id,
					pa.value,
					a.id,
					a.name,
					a.type
				FROM
					position_attributes pa
				INNER JOIN
					attribute a
				ON
					pa.attribute_id =	a.id
	
				WHERE 
					pa.position_id = $1`

	rows, _ := r.db.Query(ctx, queryPositionAttributes, position.Id)
	for rows.Next() {
		var PositionAttributes pb.GetPositionAttribute
		var attribute pb.Attribute
		err = rows.Scan(
			&PositionAttributes.Id,
			&PositionAttributes.AttributeId,
			&PositionAttributes.Value,
			&attribute.Id,
			&attribute.Name,
			&attribute.AttributeType,
		)

		if err != nil {
			return nil, err
		}

		position.PositionAttributes = append(position.PositionAttributes, &pb.GetPositionAttribute{
			Id:          PositionAttributes.Id,
			AttributeId: PositionAttributes.AttributeId,
			Value:       PositionAttributes.Value,
			Attribute:   &attribute,
		})
	}

	return &position, nil
}

func (r *positionRepo) Update(ctx context.Context, req *pb.GetPositionResponse) (string, error) {

	delQuery := `DELETE FROM position_attributes WHERE position_id = $1`
	DelResult, err := r.db.Exec(ctx, delQuery, req.Id)

	if err != nil {
		return "", fmt.Errorf("error while deleting position_attributes err: %w", err)
	}

	if i := DelResult.RowsAffected(); i == 0 {
		return "", errors.New("position_attributes not found")
	}

	query := `UPDATE 
					position 
				SET 
					name = $1, 
					profession_id = $2, 
					company_id = $3 
				WHERE id = $4`

	result, err := r.db.Exec(
		ctx,
		query,
		req.Name,
		req.ProfessionId,
		req.CompanyId,
		req.Id,
	)

	if err != nil {
		return "", fmt.Errorf("error while updating position err: %w", err)
	}

	if i := result.RowsAffected(); i == 0 {
		return "not found", nil
	}

	query = `INSERT INTO position_attributes (
		id,
		attribute_id,
		position_id,
		value
		) 
		VALUES ($1, $2, $3, $4) `

	positionAttributes := req.PositionAttributes

	for key := range positionAttributes {
		attribute_id, err := uuid.NewRandom()
		if err != nil {
			return "", err
		}
		result, err = r.db.Exec(ctx,
			query,
			attribute_id,
			positionAttributes[key].AttributeId,
			req.Id,
			positionAttributes[key].Value)

		if err != nil {
			return "", fmt.Errorf("error while updating position_attributes err: %w", err)
		}

		if i := result.RowsAffected(); i == 0 {
			return "not updated", nil
		}
	}

	return "Updated", nil

}

func (r *positionRepo) Delete(ctx context.Context, id string) (*pb.PositionStatus, error) {
	query := `DELETE FROM position_attributes WHERE position_id = $1`
	result, err := r.db.Exec(ctx, query, id)

	if err != nil {
		return nil, fmt.Errorf("error while deleting position_attributes err: %w", err)
	}

	if i := result.RowsAffected(); i == 0 {
		return nil, errors.New("position_attributes not found")
	}

	query = `DELETE FROM position WHERE id = $1`
	result, err = r.db.Exec(ctx, query, id)

	if err != nil {
		return nil, fmt.Errorf("error while deleting position err: %w", err)
	}

	if i := result.RowsAffected(); i == 0 {
		return nil, errors.New("position not found")
	}

	return &pb.PositionStatus{
		Status: "Deleted",
	}, nil

}
