package utils

func GetCache[T any](cache Cache, cacheKey string) (T, bool) {
	var blank T
	data, ok := cache.Get(cacheKey)
	if ok {
		if res, ok := data.(T); ok {
			return res, true
		} else {
			cache.Delete(cacheKey)
		}
	}
	return blank, false
}
