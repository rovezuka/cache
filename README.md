# Developing in-memory cache as a separate library on Github

A separate package that implements an in-memory cache with the following methods:

- `Set(key string, val any, ttl time.Duration)` - writing the `value` to the cache by the `key` with Time-To-Live

- `Get(key string) (any, error)`

- `Delete(key string)`

## Usage

You can install the package into your project using:
`go get -u github.com/rovezuka/cache`

### Example
```
func main() {
	// Instance creation
	cache := cache.New()

	// Writing a value by key
	cache.Set("userId", 42, time.Second*5)

	// Getting value by key
	userID, _ := cache.Get("userId")

	fmt.Println(userID)

	// Waiting for the time of life to run out
	time.Sleep(time.Second * 6)

	userID, _ = cache.Get("userId")

	fmt.Println(userID)
}
```

## License
This project is distributed under the MIT license. See the `LICENSE` file for details.
