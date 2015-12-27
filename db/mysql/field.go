package mysql

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

const (
	selectTableInfo = "SELECT COLUMN_NAME ,DATA_TYPE " +
		"FROM information_schema.`COLUMNS` " +
		"WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ? ORDER BY ORDINAL_POSITION ASC"
)

type FieldType uint8

const (
	UNKNOWN FieldType = iota
	BYTE
	STRING
	INT
	UINT
	INT64
	UINT64
	DATE
	TIME
	BOOL
	FLOAT64
)

var FieldTypeMap map[string]FieldType = map[string]FieldType{
	"tinyint":  INT,
	"smallint": INT,
	"int":      INT,

	"bigint": INT64,
	"number": INT64,

	"bool": BOOL,

	"float":   FLOAT64,
	"double":  FLOAT64,
	"real":    FLOAT64,
	"decimal": FLOAT64,

	"char":     STRING,
	"varchar":  STRING,
	"text":     STRING,
	"longtext": STRING,

	"datetime":  TIME,
	"timestamp": TIME,

	"date": DATE,
}

var FieldTypeNameMap map[FieldType]string

func init() {
	FieldTypeNameMap = make(map[FieldType]string, len(FieldTypeMap))
	for n, k := range FieldTypeMap {
		FieldTypeNameMap[k] = n
	}
}

type Field struct {
	Name string
	Type FieldType
}

// Get the fields of one table, and the order is exactly the same as they are in the table.
func GetFields(db *sql.DB, dbName string, tableName string) []Field {
	dbWrp := DBWrapper{db, dbName}
	res := dbWrp.Query(selectTableInfo, dbName, tableName)
	if res.Error != nil {
		return nil
	}
	fields := make([]Field, len(res.Rows))
	for i, row := range res.Rows {
		fName := row.Str(0)
		tName := row.Str(1)
		fType, ok := FieldTypeMap[tName]
		if !ok {
			// Try with lower case.
			fType, ok = FieldTypeMap[strings.ToLower(tName)]
		}
		if !ok {
			panic(fmt.Errorf("unsupported filed type(%s), consider configuring the FieldTypeMap to support it.", tName))
		}
		fields[i] = Field{Name: fName, Type: fType}
	}
	return fields
}

// Get the types of all fields of this table.
func GetFieldTypes(db *sql.DB, dbName string, tableName string) []FieldType {
	fields := GetFields(db, dbName, tableName)
	if fields == nil {
		return nil
	}
	types := make([]FieldType, len(fields))
	for i, f := range fields {
		types[i] = f.Type
	}
	return types
}

// Get the names of all fields of this table.
func GetFieldNames(db *sql.DB, dbName string, tableName string) []string {
	fields := GetFields(db, dbName, tableName)
	if fields == nil {
		return nil
	}
	names := make([]string, len(fields))
	for i, f := range fields {
		names[i] = f.Name
	}
	return names
}

// Get the specific field types of one table.
// Returns the types in the order of 'fieldNames'.
// If a field cannot be found in the db, a panic rise.
func GetFieldTypesByOrder(db *sql.DB, dbName string, tableName string, fieldNames []string) []FieldType {
	fields := GetFields(db, dbName, tableName)
	types := make([]FieldType, len(fieldNames))
	for i, fn := range fieldNames {
		t := UNKNOWN
		for _, f := range fields {
			if f.Name == fn {
				t = f.Type
				break
			}
		}
		if t == UNKNOWN {
			panic(fmt.Errorf("cannot find field(%s) in table(%s).", fn, tableName))
		}
		types[i] = t
	}
	return types
}

// Cast the values in the 'row' into go type specific by 'types'.
func GetValuesByType(row Row, types []FieldType) []interface{} {
	if len(row) != len(types) {
		panic(fmt.Sprintf("column and type's length are not the same! row:%d, types:%d", len(row), len(types)))
	}
	values := make([]interface{}, len(types))
	for i, t := range types {
		values[i] = GetValueByTypeAtIndex(row, t, i)
	}
	return values
}

// Cast the values in the 'row' into go type specific by 'fType' at index.
func GetValueByTypeAtIndex(row Row, fType FieldType, index int) interface{} {
	switch fType {
	case BYTE:
		return row.Bin(index)
	case STRING:
		return row.Str(index)
	case INT:
		return row.ForceInt(index)
	case UINT:
		return row.ForceUint(index)
	case INT64:
		return row.ForceInt64(index)
	case UINT64:
		return row.ForceUint64(index)
	case DATE:
		return row.ForceDate(index)
	case TIME:
		return row.ForceTime(index, time.Local)
	case BOOL:
		return row.ForceBool(index)
	case FLOAT64:
		return row.ForceFloat(index)
	default:
		panic(fmt.Sprintf("unsupported type: %d on index %d, configure error?", fType, index))
	}
}

// Get the corresponding empty values of types.
func GetEmptyValuesByType(types []FieldType) []interface{} {
	values := make([]interface{}, len(types))
	for i, t := range types {
		switch t {
		case BYTE:
			values[i] = byte(0)
		case STRING:
			values[i] = ""
		case INT:
			values[i] = 0
		case UINT:
			values[i] = uint(0)
		case INT64:
			values[i] = int64(0)
		case UINT64:
			values[i] = uint64(0)
		case DATE:
			values[i] = Date{}
		case TIME:
			values[i] = time.Time{}
		case BOOL:
			values[i] = false
		case FLOAT64:
			values[i] = 0.0
		default:
			panic(fmt.Sprintf("unsupported type: %d on index %d, configure error?", t, i))
		}
	}
	return values
}
