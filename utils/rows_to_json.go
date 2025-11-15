package utils

import "database/sql"

func RowToJson(rows *sql.Rows) ([]map[string]any, error) {
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	results := make([]map[string]any, 0)

	for rows.Next() {
		values := make([]any, len(columns))

		// collect the address of each element
		args := make([]any, len(values))
		for i := range values {
			args[i] = &values[i]
		}

		if err := rows.Scan(args...); err != nil {
			return nil, err
		}

		object := make(map[string]any)
		for i, col := range columns {
			object[col] = values[i]
		}

		results = append(results, object)
	}

	return results, nil
}
