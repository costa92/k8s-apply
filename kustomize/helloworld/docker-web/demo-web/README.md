# demo-web 

web的 golang代码
```go
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	risky := os.Getenv("ENABLE_RISKY")  // 获取环境变量
	var alt string
	if risky == "true" {
		alt = os.Getenv("ALT_GREEITMG") 获取环境变量
	}
	fmt.Fprintf(w, fmt.Sprintf("Hello astaxie: %s", alt)) // 这个写入到 w 的是输出到客户端的
}

func main() {
	http.HandleFunc("/", sayhelloName)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

```

Dockerfile 文件
```Dockerfile
FROM golang:1.19-buster as builder

WORKDIR /app
COPY . ./

RUN go build -v -o server

FROM debian:buster-slim

RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/server /app/server
EXPOSE 8000
CMD ["/app/server"]
```
## docker 生成
### 生成 docker 镜像

```shell
docker build -t demo-web:v2 .
```

### 运行镜像
```shell
docker run --name demo-web -d  demo-web:v2
```

## k8s 导入docker文件
### minikube
```shell
minikube image load demo-web:v2
```

### kind
```shell
kind image load demo-web:v2
```

## 参考文档
[minikube](https://www.zhaowenyu.com/minikube-doc/cmd/minikube-image.html#minikube-image)