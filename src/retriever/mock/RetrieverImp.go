package mock

type RetrieverImp struct {
	Content string
}

func (r RetrieverImp) Get(url string) string {
	return r.Content
}
