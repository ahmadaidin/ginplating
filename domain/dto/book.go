package dto

type FindAllBookOptions struct {
	Title      string `form:"title"`
	Author     string `form:"author"`
	Pagination Pagination
	Search     string      `form:"search"`
	Sort       []SortQuery `form:"sort"`
}

func MergeFindAllBookOptions(opts ...FindAllBookOptions) FindAllBookOptions {
	var opt FindAllBookOptions
	for _, o := range opts {
		if o.Pagination.Size != 0 {
			opt.Pagination.Size = o.Pagination.Size
		}
		if o.Pagination.Page != 0 {
			opt.Pagination.Page = o.Pagination.Page
		}
		if o.Search != "" {
			opt.Search = o.Search
		}
		if len(o.Sort) != 0 {
			opt.Sort = o.Sort
		}
	}
	return opt
}
