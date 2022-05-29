package cache

import (
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestCache(t *testing.T) {

	cache := NewCache()
	cache.Put("1", "one")
	cache.Put("2", "two")
	cache.PutTill("3", "tree", time.Now().Add(time.Second*3))
	cache.Put("4", "four")

	time.Sleep(time.Second * 1)

	keysBefore := cache.Keys()
	sort.Strings(keysBefore)
	expkeysBefore := []string{"1", "2", "3", "4"}
	if !reflect.DeepEqual(keysBefore, expkeysBefore) {
		t.Errorf("Unexpected result:\n\tExpected: %v\n\tGot: %v", expkeysBefore, keysBefore)
	}

	get1b, ok1b := cache.Get("1")
	expget1b := "one"
	expok1b := true
	if get1b != expget1b {
		t.Errorf("Unexpected result:\n\tExpected: %v\n\tGot: %v", expget1b, get1b)
	}
	if ok1b != expok1b {
		t.Errorf("Unexpected result:\n\tExpected: %v\n\tGot: %v", expok1b, ok1b)
	}

	get3b, ok3b := cache.Get("3")
	expget3b := "tree"
	expok3b := true
	if get3b != expget3b {
		t.Errorf("Unexpected result:\n\tExpected: %v\n\tGot: %v", expget3b, get3b)
	}
	if ok3b != expok3b {
		t.Errorf("Unexpected result:\n\tExpected: %v\n\tGot: %v", expok3b, ok3b)
	}

	time.Sleep(time.Second * 3)

	keysAfter := cache.Keys()
	sort.Strings(keysAfter)
	expkeysAfter := []string{"1", "2", "4"}
	if !reflect.DeepEqual(keysAfter, expkeysAfter) {
		t.Errorf("Unexpected result:\n\tExpected: %v\n\tGot: %v", expkeysAfter, keysAfter)
	}

	get1a, ok1a := cache.Get("1")
	expget1a := "one"
	expok1a := true
	if get1a != expget1a {
		t.Errorf("Unexpected result:\n\tExpected: %v\n\tGot: %v", expget1a, get1a)
	}
	if ok1a != expok1a {
		t.Errorf("Unexpected result:\n\tExpected: %v\n\tGot: %v", expok1a, ok1a)
	}

	get3a, ok3a := cache.Get("3")
	expget3a := ""
	expok3a := false
	if get3a != expget3a {
		t.Errorf("Unexpected result:\n\tExpected: %v\n\tGot: %v", expget3a, get3a)
	}
	if ok3a != expok3a {
		t.Errorf("Unexpected result:\n\tExpected: %v\n\tGot: %v", expok3a, ok3a)
	}

	get5, ok5 := cache.Get("5")
	expget5 := ""
	expok5 := false
	if get5 != expget5 {
		t.Errorf("Unexpected result:\n\tExpected: %v\n\tGot: %v", expget5, get5)
	}
	if ok5 != expok5 {
		t.Errorf("Unexpected result:\n\tExpected: %v\n\tGot: %v", expok5, ok5)
	}
	
}
