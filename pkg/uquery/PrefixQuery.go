package uquery

import (
	"github.com/blugelabs/bluge"
	v1 "github.com/prabhatsharma/zinc/pkg/meta/v1"
)

func PrefixQuery(iQuery v1.ZincQuery) (bluge.SearchRequest, error) {
	dateQuery := bluge.NewDateRangeQuery(iQuery.Query.StartTime, iQuery.Query.EndTime).SetField("@timestamp")

	var field string
	if iQuery.Query.Field != "" {
		field = iQuery.Query.Field
	} else {
		field = "_all"
	}

	prefixQuery := bluge.NewPrefixQuery(iQuery.Query.Term).SetField(field)
	query := bluge.NewBooleanQuery().AddMust(dateQuery).AddMust(prefixQuery)

	searchRequest := buildRequest(iQuery, query)

	return searchRequest, nil
}
