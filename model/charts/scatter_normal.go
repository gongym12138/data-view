package charts

import (
	"data-view/schema"
	"encoding/json"

	"github.com/jmoiron/sqlx"
)

const ScatterNormal = "ScatterNormal"

// ScatterNormalGetDataHandle 散点图
var ScatterNormalGetDataHandle = &ChartDataHandler{RunGetDataFromDB: func(db *sqlx.DB, params schema.ChartDataParams) (map[string]interface{}, error) {
	return DatasetGetDataHandle.GetDataFromDB(db, params)
}, RunGetDataFromCsv: func(params schema.ChartDataParams) (map[string]interface{}, error) {
	return nil, nil
}, RunGetDataFromRest: func(params schema.ChartDataParams) (map[string]interface{}, error) {
	example := `[
		["dimension"],
		[10.0, 8.04],
		[8.07, 6.95],
		[13.0, 7.58],
		[9.05, 8.81],
		[11.0, 8.33],
		[14.0, 7.66],
		[13.4, 6.81],
		[10.0, 6.33],
		[14.0, 8.96],
		[12.5, 6.82],
		[9.15, 7.20],
		[11.5, 7.20],
		[3.03, 4.23],
		[12.2, 7.83],
		[2.02, 4.47],
		[1.05, 3.33],
		[4.05, 4.96],
		[6.03, 7.24],
		[12.0, 6.26],
		[12.0, 8.84],
		[7.08, 5.82],
		[5.02, 5.68]
	]`

	dataset := make([]interface{}, 0)
	if err := json.Unmarshal([]byte(example), &dataset); err != nil {
		return nil, err
	}
	return map[string]interface{}{Source: dataset}, nil
}}
