package pokecache

import (
    "fmt"
    "testing"
    "time"
)

func TestAddGet(t *testing.T) {
    c := NewCache(5 * time.Second)

    cases := []struct {
        key string
        val []byte
    }{
        {"https://example.com", []byte("a")},
        {"https://example.com/x", []byte("b")},
    }

    for i, tc := range cases {
        t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
            c.Add(tc.key, tc.val)
            got, ok := c.Get(tc.key)
            if !ok {
                t.Fatalf("expected to find key")
            }
            if string(got) != string(tc.val) {
                t.Fatalf("got %q, want %q", string(got), string(tc.val))
            }
        })
    }
}

func TestReapLoop(t *testing.T) {
    interval := 5 * time.Millisecond
    c := NewCache(interval)

    c.Add("k", []byte("v"))
    if _, ok := c.Get("k"); !ok {
        t.Fatalf("expected key present before reap")
    }

    time.Sleep(interval + 10*time.Millisecond)

    if _, ok := c.Get("k"); ok {
        t.Fatalf("expected key to be reaped")
    }
}