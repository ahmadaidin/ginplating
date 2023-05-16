package dto

var (
	DefaultLimit      = 10
	DefaultPage       = 1
	DefaultPagination = Pagination{
		Page: DefaultPage,
		Size: DefaultLimit,
	}
)

type Pagination struct {
	Page int `form:"page"`
	Size int `form:"size" validate:"min=0"`
}

func (p Pagination) Offset() int {
	return (p.Page - 1) * p.Size
}

type SortQuery string // e.g. "name" for ascending or "-name" for descending

func (sortQ SortQuery) String() string {
	return string(sortQ)
}

type SortOrder int

const (
	SortAcending  SortOrder = 1
	SortDecending SortOrder = -1
)

func (sortQ SortQuery) BreakDown() (order SortOrder, field string) {
	if len(sortQ) == 0 {
		return SortAcending, ""
	}
	if sortQ[0] == '-' {
		return SortDecending, string(sortQ[1:])
	}
	return SortAcending, string(sortQ)
}
