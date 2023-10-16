# Developing in-memory cache as a separate library on Github

A separate package that implements an in-memory cache with the following methods:

- `Set(key string, value interface{})` - writing the `value` to the cache by the `key`
- `Get(key string)`
- `Delete(key)`

## Usage

You can install the package into your project using:
`go get -u github.com/rovezuka/cache`

### Example
```
func main() {
        // Instance creation
	cache := cache.New()

        // Writing a value by key
	cache.Set("userId", 42)

        // Getting value by key
	userId := cache.Get("userId")

	fmt.Println(userId)

        // Deleting a value by key
	cache.Delete("userId")
	userId := cache.Get("userId")

	fmt.Println(userId)
}
```

## License
This project is distributed under the MIT license. See the `LICENSE` file for details.
