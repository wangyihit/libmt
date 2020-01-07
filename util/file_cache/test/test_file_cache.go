package main

import "../../file_cache"

func main() {
	fc := file_cache.NewFileCache("file_cache_root")
	fc.InitDirectory()
	fc.CreateCacheDir("test001")
}
