package utils

import (
	"Frequencer/pkg/models"
	"testing"
)

func TestCacheHandler(t *testing.T) {
	// Test data
	name := "TestPreset"
	content := models.Presets{
		Name:        "Test",
		Description: "Test Description",
		Frequencies: "440",
		Type:        "Sine",
	}

	// Test CacheInformation
	CacheIt(name, content)

	// Test GetCachedInformation
	retrieved, found := GetCached(name)
	if !found {
		t.Errorf("Expected to find %s in cache", name)
	}
	if retrieved.Name != content.Name {
		t.Errorf("Expected name %s, got %s", content.Name, retrieved.Name)
	}

	// Test GetCachedInformation for non-existent item
	_, found = GetCached("NonExistent")
	if found {
		t.Error("Expected not to find NonExistent in cache")
	}

	// Test ClearCache
	ClearCache()
	_, found = GetCached(name)
	if found {
		t.Error("Expected cache to be empty after ClearCache")
	}
}
