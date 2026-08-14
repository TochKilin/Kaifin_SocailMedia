package sql

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"

	"kaifin_clone_api/pkg/logs"
	"kaifin_clone_api/pkg/share"
)

func BuildFilterClause(filters []share.Filter, startIndex int) (string, []interface{}) {
	// check filtler
	if len(filters) == 0 {
		return "", nil
	}
	// create for store sql
	sqlFilters := make([]string, 0, len(filters))
	// create for store param
	params := make([]interface{}, 0, len(filters))
	// work to follow filter
	for _, filter := range filters {
		// no field name continue
		if strings.TrimSpace(filter.Property) == "" {
			continue
		}
		// create  PostgreSQL parameter
		paramPlaceholder := fmt.Sprintf("$%d", startIndex+len(params))

		switch v := filter.Value.(type) {
		case int, int16, int32, int64, float32, float64, bool:
			sqlFilters = append(sqlFilters, fmt.Sprintf("%s = %s", filter.Property, paramPlaceholder))
			params = append(params, v)
		case string:
			if strings.Contains(v, "%") {
				sqlFilters = append(sqlFilters, fmt.Sprintf("%s LIKE %s", filter.Property, paramPlaceholder))
			} else {
				sqlFilters = append(sqlFilters, fmt.Sprintf("%s = %s", filter.Property, paramPlaceholder))
			}
			params = append(params, v)
		case time.Time:
			sqlFilters = append(sqlFilters, fmt.Sprintf("%s::DATE = %s", filter.Property, paramPlaceholder))
			params = append(params, v)
		default:
			continue
		}
	}

	if len(sqlFilters) == 0 {
		return "", nil
	}

	return strings.Join(sqlFilters, " AND "), params
}

// create order by for sort data
func BuildSortClause(sorts []share.Sort) string {
	if len(sorts) == 0 {
		return ""
	}
	// store decs & asc
	sqlSorts := make([]string, 0, len(sorts))
	/// loop sort
	for _, sort := range sorts {
		// clean space
		property := strings.TrimSpace(sort.Property)
		direction := strings.ToLower(strings.TrimSpace(sort.Direction))
		// skip emty properry
		if property == "" {
			continue
		}
		// validation
		if direction != "asc" && direction != "desc" {
			direction = "asc"
		}

		sqlSorts = append(sqlSorts, fmt.Sprintf("%s %s", property, direction))
	}

	if len(sqlSorts) == 0 {
		return ""
	}
	// ad order by == ORDER BY created_at desc, user_name asc
	return "ORDER BY " + strings.Join(sqlSorts, ", ")
}

// create pagination
func BuildPagingClause(paging share.Paging, startIndex int) (string, []interface{}) {
	page := paging.Page
	perPage := paging.Perpage

	if page < 1 {
		page = 1
	}
	if perPage < 1 {
		perPage = 10
	}

	offset := (page - 1) * perPage
	return fmt.Sprintf("LIMIT $%d OFFSET $%d", startIndex, startIndex+1), []interface{}{perPage, offset}
}

// func BuildSQLFilter(filters []share.Filter) (string, []interface{}) {
// 	return BuildFilterClause(filters, 1)
// }

// func BuildSQLSort(sorts []share.Sort) string {
// 	return BuildSortClause(sorts)
// }

func BuildSQLSearch(columns []string, search string, startIndex int) (string, []interface{}) {
	search = strings.TrimSpace(search)
	if search == "" || len(columns) == 0 {
		return "", nil
	}

	placeholder := fmt.Sprintf("$%d", startIndex)
	parts := make([]string, 0, len(columns))
	for _, column := range columns {
		column = strings.TrimSpace(column)
		if column == "" {
			continue
		}
		parts = append(parts, fmt.Sprintf("%s ILIKE %s", column, placeholder))
	}

	if len(parts) == 0 {
		return "", nil
	}

	return "(" + strings.Join(parts, " OR ") + ")", []interface{}{"%" + search + "%"}
}

type SeqResult struct {
	ID int `db:"id"`
}

func GetUserIdByField(tableName, fieldName string, value interface{}, db *sqlx.DB) (*int, error) {
	// Ensure table and field names are sanitized
	query := fmt.Sprintf(`SELECT id FROM %s WHERE %s = $1 AND deleted_at IS NULL LIMIT 1`, tableName, fieldName)

	var userID *int
	err := db.Get(&userID, query, value)
	if err != nil {
		return nil, fmt.Errorf("failed to get user id: %w", err)
	}
	return userID, nil
}

func IsExits(tbl_name string, field_name string, value interface{}, db *sqlx.DB) (bool, error) {
	var exists int

	query := fmt.Sprintf(`SELECT 1 as id FROM %s WHERE %s=$1 AND deleted_at IS NULL`, tbl_name, field_name)

	err := db.Get(&exists, query, value)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func GetSeqNextVal(seqName string, db *sqlx.DB) (*int, error) {
	var result SeqResult
	sql := `SELECT nextval($1) AS id`

	err := db.Get(&result, sql, seqName)
	if err != nil {
		logs.NewCustomLog("failed_to_get_sequence", err.Error(), "error")
		return nil, fmt.Errorf("failed to get sequence value: %w", err)
	}
	return &result.ID, nil
}

func SetSeqNextVal(seqName string, value int, db *sqlx.DB) (*int, error) {

	var result SeqResult

	// Define the SQL query
	sql := `SELECT setval($1, $2) AS id`
	err := db.Get(&result, sql, seqName, value)
	if err != nil {
		logs.NewCustomLog("failed_to_get_sequence_value", err.Error(), "error")
		return nil, fmt.Errorf("failed to set sequence value: %w", err)
	}

	return &result.ID, nil
}

func BuildSQLSort(sorts []share.Sort) string {
	if len(sorts) == 0 {
		return " ORDER BY id"
	}
	var orderClauses []string
	for _, sort := range sorts {
		orderClauses = append(orderClauses, fmt.Sprintf("%s %s", sort.Property, sort.Direction))
	}
	return " ORDER BY " + strings.Join(orderClauses, ", ")
}

func BuildSQLFilter(req []share.Filter) (string, []interface{}) {
	var sqlFilters []string
	var param []interface{}

	for i, filter := range req {
		paramPlaceholder := fmt.Sprintf("$%d", i+1)

		// Convert the filter value to the appropriate type
		switch v := filter.Value.(type) {
		case string:
			if intValue, err := strconv.Atoi(v); err == nil {
				filter.Value = intValue
			} else if boolValue, err := strconv.ParseBool(v); err == nil {
				filter.Value = boolValue
			} else if dateValue, err := time.Parse("2006-01-02", v); err == nil {
				filter.Value = dateValue
			} else {
				filter.Value = v
			}
		}

		// Handle the converted value
		switch v := filter.Value.(type) {
		case int:
			sqlFilters = append(sqlFilters, fmt.Sprintf("%s = %s", filter.Property, paramPlaceholder))
			param = append(param, v)
		case bool:
			sqlFilters = append(sqlFilters, fmt.Sprintf("%s = %s", filter.Property, paramPlaceholder))
			param = append(param, v)
		case string:
			if strings.Contains(v, "%") {
				// Handle cases with LIKE for wildcard searches
				sqlFilters = append(sqlFilters, fmt.Sprintf("%s LIKE %s", filter.Property, paramPlaceholder))
			} else {
				sqlFilters = append(sqlFilters, fmt.Sprintf("%s = %s", filter.Property, paramPlaceholder))
			}
			param = append(param, v)
		case time.Time:
			// Handle date comparison
			sqlFilters = append(sqlFilters, fmt.Sprintf("%s::DATE = %s", filter.Property, paramPlaceholder))
			param = append(param, v)
		default:
			return "", nil
		}
	}
	filterClause := strings.Join(sqlFilters, " AND ")
	return filterClause, param
}
