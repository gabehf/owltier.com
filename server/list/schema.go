package list

type List struct {
	Id        string   `json:"id"`
	CreatedAt int64    `json:"created_at"`
	CreatedBy string   `json:"created_by"`
	Format    string   `json:"format"`
	Breaks    []bool   `json:"breaks"`
	NA        []string `json:"na"`
	APAC      []string `json:"apac"`
	Combined  []string `json:"combined"`
}
