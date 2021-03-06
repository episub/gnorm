// Code generated by gnorm, DO NOT EDIT!

package statistics

import (
	"database/sql"

	"github.com/episub/gnorm/database/drivers/mysql/gnorm"
)

// Row represents a row from 'STATISTICS'.
type Row struct {
	TableCatalog string         // TABLE_CATALOG
	TableSchema  string         // TABLE_SCHEMA
	TableName    string         // TABLE_NAME
	NonUnique    int64          // NON_UNIQUE
	IndexSchema  string         // INDEX_SCHEMA
	IndexName    string         // INDEX_NAME
	SeqInIndex   int64          // SEQ_IN_INDEX
	ColumnName   string         // COLUMN_NAME
	Collation    sql.NullString // COLLATION
	Cardinality  sql.NullInt64  // CARDINALITY
	SubPart      sql.NullInt64  // SUB_PART
	Packed       sql.NullString // PACKED
	Nullable     string         // NULLABLE
	IndexType    string         // INDEX_TYPE
	Comment      sql.NullString // COMMENT
	IndexComment string         // INDEX_COMMENT
}

// Field values for every column in Statistics.
var (
	TableCatalogCol gnorm.StringField        = "TABLE_CATALOG"
	TableSchemaCol  gnorm.StringField        = "TABLE_SCHEMA"
	TableNameCol    gnorm.StringField        = "TABLE_NAME"
	NonUniqueCol    gnorm.Int64Field         = "NON_UNIQUE"
	IndexSchemaCol  gnorm.StringField        = "INDEX_SCHEMA"
	IndexNameCol    gnorm.StringField        = "INDEX_NAME"
	SeqInIndexCol   gnorm.Int64Field         = "SEQ_IN_INDEX"
	ColumnNameCol   gnorm.StringField        = "COLUMN_NAME"
	CollationCol    gnorm.SqlNullStringField = "COLLATION"
	CardinalityCol  gnorm.SqlNullInt64Field  = "CARDINALITY"
	SubPartCol      gnorm.SqlNullInt64Field  = "SUB_PART"
	PackedCol       gnorm.SqlNullStringField = "PACKED"
	NullableCol     gnorm.StringField        = "NULLABLE"
	IndexTypeCol    gnorm.StringField        = "INDEX_TYPE"
	CommentCol      gnorm.SqlNullStringField = "COMMENT"
	IndexCommentCol gnorm.StringField        = "INDEX_COMMENT"
)

// Query retrieves rows from 'STATISTICS' as a slice of Row.
func Query(db gnorm.DB, where gnorm.WhereClause) ([]*Row, error) {
	const origsqlstr = `SELECT 
		TABLE_CATALOG, TABLE_SCHEMA, TABLE_NAME, NON_UNIQUE, INDEX_SCHEMA, INDEX_NAME, SEQ_IN_INDEX, COLUMN_NAME, COLLATION, CARDINALITY, SUB_PART, PACKED, NULLABLE, INDEX_TYPE, COMMENT, INDEX_COMMENT
		FROM information_schema.STATISTICS WHERE (`

	sqlstr := origsqlstr + where.String() + ") "

	var vals []*Row
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, err
	}
	for q.Next() {
		r := Row{}

		err = q.Scan(&r.TableCatalog, &r.TableSchema, &r.TableName, &r.NonUnique, &r.IndexSchema, &r.IndexName, &r.SeqInIndex, &r.ColumnName, &r.Collation, &r.Cardinality, &r.SubPart, &r.Packed, &r.Nullable, &r.IndexType, &r.Comment, &r.IndexComment)
		if err != nil {
			return nil, err
		}

		vals = append(vals, &r)
	}
	return vals, nil
}
