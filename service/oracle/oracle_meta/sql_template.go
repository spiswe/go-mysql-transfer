package meta

const DOT string = "."

var SQLTemplateControl = &SQLTemplate{}

type SQLTemplate struct {
}

func (s *SQLTemplate) GetColumnNameByMeta(meta ColumnMeta) string {
	return meta.Name()
}

func (s *SQLTemplate) MakeColumn(columns []ColumnMeta) string {
	columnSize := len(columns)
	var str string
	for idx, col := range columns {
		str += s.GetColumnNameByMeta(col)
		if idx < columnSize {
			str += ","
		}
	}
	return str
}
