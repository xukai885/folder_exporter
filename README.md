#### 检测指定文件夹大小
默认端口：9101
#### 启动参数
```shell
./folder_exporter Path1 Path2 .... -p 9102
```

```shell
curl localhost:9101/metrics

# 文件大小字节 bytes 
folder_size_bytes{folder="/data/redis"} 16260945
folder_size_bytes{folder="/data/monitor"} 455118851
folder_size_bytes{folder="/data/localai"} 4634935
```