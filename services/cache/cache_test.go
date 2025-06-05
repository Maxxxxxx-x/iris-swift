package cache

import (
	"github.com/Maxxxxxx-x/iris-swift/config"
	"testing"
)

func TestCache(t *testing.T) {
	cfg := config.BigCache{
		Enabled:            true,
		Shards:             1024,
		LifeWindow:         "10m",
		CleanWindow:        "1m",
		MaxEntriesInWindow: 600000,
		MaxEntrySize:       500,
		Verbose:            true,
		HardMaxCacheSize:   8192,
	}
	cache, err := New(cfg)
	if err != nil {
		t.Errorf("NewCache err: %v", err)
		return
	}
	cache.Set("foo", "bar")
	value, err := cache.Get("foo")
	if err != nil {
		t.Errorf("Get err: %v", err)
		return
	}
	if value != "bar" {
		t.Errorf("Get value: %v, want %v", value, "bar")
		return
	}
	t.Logf("Cache test passed, value: %s", value)
}
