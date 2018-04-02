package real

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

type Retriever struct {
	User    string
	TImeOut int
}

func (r *Retriever) String() string {
	return fmt.Sprintf("%s %d", r.User, r.TImeOut)
}

func (r *Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, e := httputil.DumpResponse(resp, true)

	if e != nil {
		panic(e)
	}
	return string(bytes)
}
