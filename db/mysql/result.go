package mysql

type QueryResult struct {
	Cols []string
	Rows []Row

	Error error

	colToIndex map[string]int
}

func (q *QueryResult) ColID(colName string) int {
	if q.colToIndex == nil {
		colToIndex := make(map[string]int, len(q.Cols))
		for i, n := range q.Cols {
			colToIndex[n] = i
		}
		q.colToIndex = colToIndex
	}
	return q.colToIndex[colName]
}
