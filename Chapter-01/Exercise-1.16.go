// Constants
package main

import "fmt"

const GlobalLimit = 100
const MaxCacheSize = 10 * GlobalLimit

// Cache prefixes
const (
	CacheKeyBook = "book_"
	CacheKeyCd   = "cd_"
)

// Declare a map that has a string for a key and a string for its values as our cache
var cache map[string]string

// Create a function to get items from the cache (Search and return the item)
func cacheGet(key string) string {
	return cache[key]
}

// Create a function that sets items in the cache
func cacheSet(key, val string) {
	if len(cache)+1 >= MaxCacheSize {
		return
	}
	cache[key] = val
}

// GetBook Create a function to get a book from the cache
func GetBook(isbn string) string {
	return cacheGet(CacheKeyBook + isbn)
}

// SetBook Use the book cache prefix to create a unique key
func SetBook(isbn string, name string) {
	cacheSet(CacheKeyBook+isbn, name)
}

func GetCd(sku string) string {
	return cacheGet(CacheKeyCd + sku)
}
func SetCd(sku string, title string) {
	cacheSet(CacheKeyCd+sku, title)
}
func main() {
	cache = make(map[string]string) // make a slice of maps
	SetBook("1234-5678", "Get Ready To Go")
	SetCd("1234-5678", "Get Ready To Go Audio Book")
	fmt.Println("Book :", GetBook("1234-5678"))
	fmt.Println("CD :", GetCd("1234-5678"))

	for key, value := range cache {
		fmt.Println("Key:", key, "Value:", value)
	}
}
