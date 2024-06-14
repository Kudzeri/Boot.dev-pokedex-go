package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	if cache.Cache == nil {
		t.Errorf("cache is nil")
	}
}

func TestGetCache(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	cases := []struct {
		inputKey    string
		inputVal    []byte
		expectedVal string
	}{
		{
			inputKey:    "key1",
			inputVal:    []byte("val1"),
			expectedVal: "val1",
		},
		{
			inputKey:    "2",
			inputVal:    []byte("2"),
			expectedVal: "2",
		},
	}

	for _, cas := range cases {
		cache.Add(cas.inputKey, cas.inputVal)

		actual, ok := cache.Get(cas.inputKey)
		if !ok {
			t.Errorf("%s not found", cas.inputKey)
			continue
		}

		if string(actual) != cas.expectedVal {
			t.Errorf("%s doesn't match %s", cas.inputVal, cas.expectedVal)
			continue
		}
	}
}

func TestCache_Reap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("val1"))

	time.Sleep(interval + time.Millisecond*20)

	_, ok := cache.Get(keyOne)
	if ok {
		t.Errorf("%s should have been reaped", keyOne)
	}
}
