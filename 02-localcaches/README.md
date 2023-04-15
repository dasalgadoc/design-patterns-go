# 🧑🏻‍💻 Local Caches

When working with expensive functions, it's possible to save in memory results already calculated to saving time for repeated calculations.

In Go, caches can be implemented as object in memory.

Concurrent implementations can also be implemented with GoRoutines and locks to avoid race conditions. There's some libraries who simplify the source code implementations.

## 📕 Library

- [PatrickMN Go cache](https://github.com/patrickmn/go-cache)
- [Eko Go Cache](https://github.com/eko/gocache)
- [List of caches](https://awesome-go.com/caches/) Find a specific cache