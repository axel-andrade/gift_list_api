package entities

import "math"

type PaginateResult struct {
	Docs        interface{} `json:"docs"`
	TotalDocs   int64       `json:"total_docs,omitempty"`
	Limit       int64       `json:"limit,omitempty"`
	Page        int64       `json:"page,omitempty"`
	TotalPages  int         `json:"total_pages,omitempty"`
	HasPrevPage bool        `json:"has_prev_page"`
	HasNextPage bool        `json:"has_next_page"`
	PrevPage    int64       `json:"prev_page,omitempty"`
	NextPage    int64       `json:"next_page,omitempty"`
}

func BuildPaginateResult(docs interface{}, pagination PaginationOptions, totalDocs int64) PaginateResult {
	var result PaginateResult

	totalPages := int(math.Ceil(float64(totalDocs) / float64(pagination.Limit)))

	result.Docs = docs
	result.TotalPages = totalPages
	result.Limit = int64(pagination.Limit)
	result.Page = int64(pagination.Page)
	result.HasPrevPage = int64(pagination.Page) > 1
	result.HasNextPage = int64(pagination.Page) < int64(totalPages)

	if result.HasPrevPage {
		result.PrevPage = int64(pagination.Page) - 1
	}

	if result.HasNextPage {
		result.NextPage = int64(pagination.Page) + 1
	}

	return result
}
