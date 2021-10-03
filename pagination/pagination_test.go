package pagination

import (
	"reflect"
	"testing"

	"github.com/diptomondal007/go-utils/variables"
)

func TestGeneratePagination(t *testing.T) {
	type args struct {
		data     interface{}
		count    int64
		page     int64
		pageSize int64
	}
	tests := []struct {
		name    string
		args    args
		want    PaginatedResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test-01",
			args: args{
				data:     nil,
				count:    0,
				page:     0,
				pageSize: 0,
			},
			want: PaginatedResponse{
				Meta: PaginatorMeta{
					TotalCount: 0,
					Count:      0,
					PageSize:   0,
					Next:       nil,
					Previous:   nil,
				},
				Data: make([]interface{}, 0),
			},
		},
		{
			name: "test-02",
			args: args{
				data:     nil,
				count:    20,
				page:     1,
				pageSize: 10,
			},
			want: PaginatedResponse{
				Meta: PaginatorMeta{
					TotalCount: 20,
					Count:      0,
					PageSize:   10,
					Next:       variables.Int64P(2),
					Previous:   nil,
				},
				Data: make([]interface{}, 0),
			},
		},
		{
			name: "test-03",
			args: args{
				data: []struct {
					a int
					b string
				}{{
					a: 10, b: "a",
				},
					{
						a: 20, b: "ab",
					},
				},
				count:    25,
				page:     2,
				pageSize: 10,
			},
			want: PaginatedResponse{
				Meta: PaginatorMeta{
					TotalCount: 25,
					Count:      2,
					PageSize:   10,
					Next:       variables.Int64P(3),
					Previous:   variables.Int64P(1),
				},
				Data: []struct {
					a int
					b string
				}{
					{
						a: 10, b: "a",
					},
					{
						a: 20, b: "ab",
					},
				},
			},
		},
		{
			name: "test-04",
			args: args{
				data: []struct {
					a int
					b string
				}{{
					a: 10, b: "a",
				},
					{
						a: 20, b: "ab",
					},
				},
				count:    25,
				page:     3,
				pageSize: 10,
			},
			want: PaginatedResponse{
				Meta: PaginatorMeta{
					TotalCount: 25,
					Count:      2,
					PageSize:   10,
					Next:       nil,
					Previous:   variables.Int64P(2),
				},
				Data: []struct {
					a int
					b string
				}{
					{
						a: 10, b: "a",
					},
					{
						a: 20, b: "ab",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GeneratePagination(tt.args.data, tt.args.count, tt.args.page, tt.args.pageSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("GeneratePagination() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GeneratePagination() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newPagination(t *testing.T) {
	type args struct {
		count    int64
		page     int64
		pageSize int64
	}
	tests := []struct {
		name string
		args args
		want Paginator
	}{
		// TODO: Add test cases.
		{
			name: "test-01",
			args: args{
				count:    40,
				page:     3,
				pageSize: 10,
			},
			want: &paginator{
				data:       nil,
				count:      40,
				page:       3,
				pageSize:   10,
				totalPages: 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newPagination(tt.args.count, tt.args.page, tt.args.pageSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newPagination() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_paginator_next(t *testing.T) {
	type fields struct {
		data       interface{}
		count      int64
		page       int64
		pageSize   int64
		totalPages int64
	}
	tests := []struct {
		name   string
		fields fields
		want   *int64
	}{
		// TODO: Add test cases.
		{
			name: "test-01",
			fields: fields{
				data:       nil,
				count:      40,
				page:       3,
				pageSize:   10,
				totalPages: 4,
			},
			want: variables.Int64P(4),
		},
		{
			name: "test-01",
			fields: fields{
				data:       nil,
				count:      20,
				page:       2,
				pageSize:   10,
				totalPages: 2,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &paginator{
				data:       tt.fields.data,
				count:      tt.fields.count,
				page:       tt.fields.page,
				pageSize:   tt.fields.pageSize,
				totalPages: tt.fields.totalPages,
			}
			if got := p.next(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_paginator_previous(t *testing.T) {
	type fields struct {
		data       interface{}
		count      int64
		page       int64
		pageSize   int64
		totalPages int64
	}
	tests := []struct {
		name   string
		fields fields
		want   *int64
	}{
		// TODO: Add test cases.
		{
			name: "test-01",
			fields: fields{
				data:       nil,
				count:      40,
				page:       3,
				pageSize:   10,
				totalPages: 4,
			},
			want: variables.Int64P(2),
		},
		{
			name: "test-01",
			fields: fields{
				data:       nil,
				count:      20,
				page:       1,
				pageSize:   10,
				totalPages: 2,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &paginator{
				data:       tt.fields.data,
				count:      tt.fields.count,
				page:       tt.fields.page,
				pageSize:   tt.fields.pageSize,
				totalPages: tt.fields.totalPages,
			}
			if got := p.previous(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("previous() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_paginator_validate(t *testing.T) {
	type fields struct {
		data       interface{}
		count      int64
		page       int64
		pageSize   int64
		totalPages int64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test-01",
			fields: fields{
				data:       nil,
				count:      20,
				page:       3,
				pageSize:   10,
				totalPages: 2,
			},
			wantErr: true,
		},
		{
			name: "test-01",
			fields: fields{
				data:       nil,
				count:      20,
				page:       2,
				pageSize:   10,
				totalPages: 2,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &paginator{
				data:       tt.fields.data,
				count:      tt.fields.count,
				page:       tt.fields.page,
				pageSize:   tt.fields.pageSize,
				totalPages: tt.fields.totalPages,
			}
			if err := p.validate(); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_totalPages(t *testing.T) {
	type args struct {
		count    int64
		pageSize int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{
			name: "test-01",
			args: args{
				count:    25,
				pageSize: 10,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalPages(tt.args.count, tt.args.pageSize); got != tt.want {
				t.Errorf("totalPages() = %v, want %v", got, tt.want)
			}
		})
	}
}
