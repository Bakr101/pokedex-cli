package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T){
	cache := NewCache(30 * time.Second)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}


func TestCacheImplementation(t *testing.T){
	cache := NewCache(30 * time.Second)
	
	cases := []struct{
		inputKey	string
		inputVal	[]byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("val2"),
		},
		{
			inputKey: "",
			inputVal: []byte("val3"),
		},
	}

	for _, cas := range cases {
		cache.Add(cas.inputKey, cas.inputVal)
		actual, ok := cache.Get(cas.inputKey)
		if !ok {
			t.Errorf("%v not found", cas.inputKey)
			continue
		}
		
		if string(actual) != string(cas.inputVal) {
			t.Errorf("actual not equal expected.  Actual: %v Expected: %v", string(actual), string(cas.inputVal))
		}
	}

}

func TestReap(t *testing.T){
	interval := 10 * time.Millisecond
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("val1"))

	time.Sleep(interval + time.Millisecond)
	_, ok := cache.Get(keyOne)
	if ok {
		t.Errorf("%s should have been reaped", keyOne)
	}
}


func TestReapFail(t *testing.T){
	interval := 10 * time.Millisecond
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("val1"))

	time.Sleep(time.Millisecond)
	_, ok := cache.Get(keyOne)
	if !ok {
		t.Errorf("%s should have been reaped", keyOne)
	}
}