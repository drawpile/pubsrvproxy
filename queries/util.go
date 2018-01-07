package queries

import (
	"net/http"
	"encoding/json"
	"log"
	"io/ioutil"
)

func getJson(url string, v interface{}, cache *QueryCache) error {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	cachedResponse, cached := cache.get(url)

	var body []byte

	if cached {
		if cachedResponse.err != nil {
			return cachedResponse.err
		}
		body = cachedResponse.body

	} else {
		log.Println("Querying:", url)
		resp, err := http.Get(url)
		if err != nil {
			log.Println(url, "Server request error:", err)
			cache.put(url, nil, err)
			return err
		}

		defer resp.Body.Close()

		body, err = ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Println(url, "Server request error:", err)
			cache.put(url, nil, err)
			return err
		}
		cache.put(url, body, nil)
	}

	return json.Unmarshal(body, v)
}

