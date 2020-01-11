package disk_cache

import (
	"fmt"
	"os"
	"time"
)

type ItemCache struct {
	cacheID string
	day     string
	root    string
}

func NewItemCache(cacaheRoot string, cacheID string) *ItemCache {
	c := time.Now()
	day := c.Day()
	path := fmt.Sprint("%s/%02d/%s", cacaheRoot, day, cacheID)
	ic := &ItemCache{
		cacheID: cacheID,
		day:     fmt.Sprintf("%02d", day),
		root:    path,
	}
	return ic
}

func (ic *ItemCache) Root() string {
	return ic.root
}

func (ic *ItemCache) SubDir(name string) (string, error) {
	path := fmt.Sprintf("%s/%s", ic.root, name)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		return path, err
	}
	return path, nil
}

func (ic *ItemCache) Clean() error {
	err := os.RemoveAll(ic.root)
	return err
}

type DiskCache struct {
	root        string
	initialized bool
}

func NewDiskCache(root string) *DiskCache {

	fc := &DiskCache{
		root:        root,
		initialized: false,
	}
	return fc
}

func fileExist(fileName string) bool {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return false
	}
	return true
}

func (fc *DiskCache) InitDirectory() error {
	for i := 1; i <= 31; i++ {
		path := fmt.Sprintf("%s/%02d", fc.root, i)
		if !fileExist(path) {
			err := os.MkdirAll(path, 0755)
			if err != nil {
				return err
			}
		}
	}
	fc.initialized = true
	return nil
}

func (fc *DiskCache) Initialized() bool {
	return fc.initialized
}

func (fc *DiskCache) CreateItemCache(cacheID string) (*ItemCache, error) {
	itemCache := NewItemCache(fc.root, cacheID)
	err := os.MkdirAll(itemCache.root, 0755)
	if err != nil {
		return nil, err
	}
	return itemCache, nil
}
