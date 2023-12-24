package utils

import (
	"be-park-ease/constants"
	"context"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type (
	QueryBuilder struct {
		idx           int
		filters       []queryFilter
		order         string
		groupBy       string
		offset, limit int
	}

	queryFilter struct {
		expression string
		args       []interface{}
	}
)

type QueryDBTX interface {
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
}

type queryWrappedDB struct {
	QueryDBTX
}

func QueryWrap(db QueryDBTX) QueryDBTX {
	return &queryWrappedDB{
		db,
	}
}

func (w queryWrappedDB) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	if b, ok := QueryBuilderFrom(ctx); ok {
		query, args = b.Build(query, args...)
	}

	return w.QueryDBTX.Exec(ctx, query, args...)
}

func (w queryWrappedDB) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	if b, ok := QueryBuilderFrom(ctx); ok {
		query, args = b.Build(query, args...)
	}

	return w.QueryDBTX.Query(ctx, query, args...)
}

func (w queryWrappedDB) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	if b, ok := QueryBuilderFrom(ctx); ok {
		query, args = b.Build(query, args...)
	}

	return w.QueryDBTX.QueryRow(ctx, query, args...)
}

type queryBuilderContextKey struct{}

func QueryWithBuilder(ctx context.Context, b *QueryBuilder) context.Context {
	return context.WithValue(ctx, queryBuilderContextKey{}, b)
}

func QueryBuilderFrom(ctx context.Context) (*QueryBuilder, bool) {
	b, ok := ctx.Value(queryBuilderContextKey{}).(*QueryBuilder)
	return b, ok
}

func QueryBuild(ctx context.Context, f func(builder *QueryBuilder)) context.Context {
	b, ok := QueryBuilderFrom(ctx)
	if !ok {
		b = &QueryBuilder{}
	} else {
		b = b.clone()
	}

	f(b)
	return QueryWithBuilder(ctx, b)
}

func (b *QueryBuilder) clone() *QueryBuilder {
	cb := QueryBuilder{}
	cb = *b
	return &cb
}

func (b *QueryBuilder) checkArgument(query *string, args ...interface{}) {
	checkArr := make(map[string]bool)
	listArr := []string{}
	replaceArr := make(map[string]string)

	re := regexp.MustCompile(`(\$[0-9]+)`)
	matches := re.FindAllString(*query, -1)
	for _, match := range matches {
		if _, ok := checkArr[match]; !ok {
			checkArr[match] = true
			listArr = append(listArr, match)
		}
	}

	if len(checkArr) != len(args) {
		panic("Query and args not match")
	}

	sort.Strings(listArr)
	for idx, key := range listArr {
		replaceArr[key] = fmt.Sprintf("$%d", b.idx+idx+1)
	}

	b.idx += len(args)
	*query = re.ReplaceAllStringFunc(*query, func(match string) string {
		return replaceArr[match]
	})

}

func (b *QueryBuilder) Where(query string, args ...interface{}) *QueryBuilder {
	b.checkArgument(&query, args...)
	b.filters = append(b.filters, queryFilter{
		expression: query,
		args:       args,
	})

	return b
}

func (b *QueryBuilder) In(column string, args ...interface{}) *QueryBuilder {
	placeholders := make([]string, len(args))
	for i := range args {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}

	query := fmt.Sprintf("%s IN (%s)", column, strings.Join(placeholders, ","))
	return b.Where(query, args...)
}

func (b *QueryBuilder) Order(cols string) *QueryBuilder {
	b.order = cols
	return b
}

func (b *QueryBuilder) Ordering(cols, ordering string) *QueryBuilder {
	b.order = fmt.Sprintf("%s %s", cols, ordering)
	return b
}

func (b *QueryBuilder) Offset(x int) *QueryBuilder {
	b.offset = x
	return b
}

func (b *QueryBuilder) Limit(x int) *QueryBuilder {
	b.limit = x
	return b
}

func (b *QueryBuilder) Pagination(page, limit int) *QueryBuilder {
	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = constants.DefaultPageLimit
	}

	offset := (page - 1) * limit
	if offset < 0 {
		offset = 0
	}

	b.limit = limit
	b.offset = offset
	return b
}

func (b *QueryBuilder) GroupBy(groupBy string) *QueryBuilder {
	b.groupBy = groupBy
	return b
}

func (b *QueryBuilder) Build(query string, args ...interface{}) (string, []interface{}) {
	var sb strings.Builder

	sb.WriteString(query)
	sb.WriteByte('\n')

	for idx, filter := range b.filters {
		if idx == 0 {
			sb.WriteString("WHERE ")
		} else {
			sb.WriteString("AND ")
		}

		sb.WriteByte('(')
		sb.WriteString(filter.expression)
		sb.WriteByte(')')
		sb.WriteByte('\n')

		args = append(args, filter.args...)
	}

	if b.groupBy != "" {
		sb.WriteString("GROUP BY ")
		sb.WriteString(b.groupBy)
		sb.WriteByte('\n')
	}

	if b.order != "" {
		sb.WriteString("ORDER BY ")
		sb.WriteString(b.order)
		sb.WriteByte('\n')
	}

	if b.limit > 0 {
		sb.WriteString("LIMIT ")
		sb.WriteString(strconv.Itoa(b.limit))
		sb.WriteByte('\n')
	}

	if b.offset > 0 {
		sb.WriteString("OFFSET ")
		sb.WriteString(strconv.Itoa(b.offset))
		sb.WriteByte('\n')
	}

	return sb.String(), args
}
