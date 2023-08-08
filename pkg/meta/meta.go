package meta

type Meta struct {
	TotalCount int `xml:"total_count"`
}

func New(total int) (*Meta, error) {
	return &Meta{
		TotalCount: total,
	}, nil
}
