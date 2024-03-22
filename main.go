package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

var folders []string
var port string

func main() {
	flag.StringVar(&port, "p", "9101", "监听端口")
	flag.Parse()

	folders = flag.Args()
	if len(folders) == 0 {
		fmt.Println("请输入需要监控的文件夹")
		os.Exit(1)
	}

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		folderSizes := calculateTotalSize(folders)
		fmt.Fprintf(w, "# 文件大小字节 bytes \n")
		for folder, size := range folderSizes {
			fmt.Fprintf(w, "folder_size_bytes{folder=\"%s\"} %d\n", folder, size)
		}
	})

	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func calculateTotalSize(folders []string) map[string]int64 {
	folderSizes := make(map[string]int64)
	var wg sync.WaitGroup
	var lock sync.Mutex

	for _, folder := range folders {
		wg.Add(1)
		go func(folder string) {
			defer wg.Done()
			var folderSize int64

			filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					fmt.Printf("error accessing path %s: %v\n", path, err)
					return nil
				}

				if !info.IsDir() {
					folderSize += info.Size()
				}

				return nil
			})

			lock.Lock()
			folderSizes[folder] = folderSize
			lock.Unlock()
		}(folder)
	}

	wg.Wait()
	return folderSizes
}
