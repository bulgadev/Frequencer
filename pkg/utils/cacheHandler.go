package utils

import "Frequencer/pkg/models"

var presetCache = make(map[string]models.Presets)

// CacheInformation stores a preset in the cache by name.
func CacheIt(name string, content models.Presets) {
	presetCache[name] = content
}

// GetCachedInformation retrieves a preset from the cache by name.
// Returns the preset and a boolean indicating if it was found.
func GetCached(name string) (models.Presets, bool) {
	content, found := presetCache[name]
	return content, found
}

// ClearCache removes all items from the cache.
func ClearCache() {
	presetCache = make(map[string]models.Presets)
}
