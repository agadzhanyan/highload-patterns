# highload-patterns

List of patterns to handle high load easily

### Refresh-ahead caching
Caching data in app and update in background

![Refresh-ahead caching](pics/refresh-ahead.png "Refresh-ahead caching")
Figure out [example](./refresh-ahead/main.go)

```
curl --location --request GET 'localhost:8890/getPopularMovies'
```

> More details about caching on [system-design-primer](https://github.com/donnemartin/system-design-primer#refresh-ahead)

### Singleflight

Making work once by calculating and syncing by work hash

![Singleflight](pics/singleflight.png "Refresh-ahead caching")
Figure out [example](./singleflight/main.go)

> curl --location --request GET 'localhost:8890/getBook/100'

### Worker pool
