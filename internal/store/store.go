// internal/store/store.go
package store

import (
	"github.com//Ashikurrahaman287//url-shortener/internal/model"
	"sync"
)

type URLStore struct {
	sync.RWMutex
	urls map[string]*model.URL
}

func NewURLStore() *URLStore {
	return &URLStore{
		urls: make(map[string]*model.URL),
	}
}

func (s *URLStore) Save(url *model.URL) {
	s.Lock()
	defer s.Unlock()
	s.urls[url.Short] = url
}

func (s *URLStore) Find(short string) (*model.URL, bool) {
	s.RLock()
	defer s.RUnlock()
	url, exists := s.urls[short]
	return url, exists
}
