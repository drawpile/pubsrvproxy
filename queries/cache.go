package queries

import (
	"time"
	"sync"
)

type cachedResponse struct {
	body []byte
	err error
	timestamp int64
}

type QueryCache struct {
	ExpirationTime int64
	queries map[string]cachedResponse
	mutex sync.Mutex
}

func (q *QueryCache) get(url string) (cachedResponse, bool) {
	if q.queries == nil {
		return cachedResponse{}, false
	}
	val, ok := q.queries[url]
	if !ok || val.timestamp + q.ExpirationTime < time.Now().Unix() {
		return val, false
	}
	return val, true
}

func (q *QueryCache) put(url string, body []byte, err error) {
	if q.ExpirationTime > 0 {
		if q.queries == nil {
			q.queries = make(map[string]cachedResponse)
		}

		q.queries[url] = cachedResponse{
			body,
			err,
			time.Now().Unix(),
		}
	}
}

