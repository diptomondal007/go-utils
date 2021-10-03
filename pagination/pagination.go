package pagination

import (
	"errors"
	"math"
	"reflect"

	"github.com/diptomondal007/go-utils/variables"
)

type Paginator interface {
	previous() *int64
	next() *int64
	validate() error
}

type paginator struct {
	data       interface{}
	count      int64
	page       int64
	pageSize   int64
	totalPages int64
}

type PaginatedResponse struct {
	Meta PaginatorMeta `json:"meta"`
	Data interface{}   `json:"data"`
}

type PaginatorMeta struct {
	TotalCount int64  `json:"total_count"`
	Count      int64  `json:"count"`
	PageSize   int64  `json:"page_size"`
	Next       *int64 `json:"next"`
	Previous   *int64 `json:"previous"`
	TotalPages int64  `json:"total_pages"`
}

func GeneratePagination(data interface{}, count, page, pageSize int64) (PaginatedResponse, error) {
	p := newPagination(count, page, pageSize)

	if data == nil {
		data = make([]interface{}, 0)
	}

	err := p.validate()
	if err != nil {
		return PaginatedResponse{}, err
	}

	return PaginatedResponse{
		Meta: PaginatorMeta{
			TotalCount: count,
			Count:      int64(reflect.ValueOf(data).Len()),
			PageSize:   pageSize,
			Next:       p.next(),
			Previous:   p.previous(),
		},
		Data: data,
	}, nil
}

func (p *paginator) previous() *int64 {
	if int(p.totalPages)-int(p.page) >= 0 && (p.page-1 > 0) {
		return variables.Int64P(p.page - 1)
	}
	return nil
}

func (p *paginator) next() *int64 {
	if p.totalPages-p.page > 0 {
		return variables.Int64P(p.page + 1)
	}
	return nil
}

func (p *paginator) validate() error {
	if p.page > p.totalPages {
		return errors.New("page number out of bound")
	}
	return nil
}

func newPagination(count, page, pageSize int64) Paginator {
	return &paginator{
		data:       nil,
		count:      count,
		page:       page,
		pageSize:   pageSize,
		totalPages: totalPages(count, pageSize),
	}
}

func totalPages(count, pageSize int64) int64 {
	if pageSize == 0 {
		return 0
	}

	return int64(math.Ceil(float64(count) / float64(pageSize)))
}
