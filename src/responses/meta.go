package responses

type Meta struct {
	Page      int   `json:"page" form:"page"`
	PerPage   int   `json:"perPage" form:"perPage"`
	MaxPage   int   `json:"maxPage" form:"maxPage"`
	TotalData int64 `json:"totalData" form:"totalData"`
}

func (p *Meta) GetOffset() int {
	return (p.GetPage() - 1) * p.GetPerPage()
}

func (m *Meta) GetPerPage() int {
	if m.PerPage == 0 {
		m.PerPage = 10
	}
	return m.PerPage
}

func (m *Meta) GetPage() int {
	if m.Page == 0 {
		m.Page = 1
	}
	return m.Page
}

func (m *Meta) GetMaxPage() int {
	return int(m.TotalData)
}
