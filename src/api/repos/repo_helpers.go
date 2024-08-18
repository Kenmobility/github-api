package repos

import "github.com/kenmobility/github-api/src/api/models"

const (
	DEFAULTPAGE                  = 1
	DEFAULTLIMIT                 = 10
	PageDefaultSortBy            = "created_at"
	PageDefaultSortDirectionDesc = "desc"
)

func getPaginationInfo(query models.APIPagingDto) (models.APIPagingDto, int) {
	var offset int
	// load defaults
	if query.Page == 0 {
		query.Page = DEFAULTPAGE
	}
	if query.Limit == 0 {
		query.Limit = DEFAULTLIMIT
	}

	if query.Sort == "" {
		query.Sort = PageDefaultSortBy
	}

	if query.Direction == "" {
		query.Direction = PageDefaultSortDirectionDesc
	}

	if query.Page > 1 {
		offset = query.Limit * (query.Page - 1)
	}
	return query, offset
}

func getPagingInfo(query models.APIPagingDto, count int) models.PagingInfo {
	var hasNextPage bool

	next := int64((query.Page * query.Limit) - count)
	if next < 0 {
		hasNextPage = true
	}

	pagingInfo := models.PagingInfo{
		TotalCount:  int64(count),
		HasNextPage: hasNextPage,
		Page:        int(query.Page),
	}

	return pagingInfo
}
