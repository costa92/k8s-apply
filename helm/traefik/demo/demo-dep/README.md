# demo-dep 
## raefik 运行 raefik-api
1. 新建 Deployment 类型文件
[traefik-api-dep.yaml](traefik-api-dep.yaml)
```yaml
      containers:
        - name: nginx
          image: costa92/treafik-api:v0.0.1   # 镜像名称与版本
          ports:
            - containerPort: 8080
```

2. 新建 IngressRoute 类型文件
[trafike.yaml](trafike.yaml)
```yaml
apiVersion: traefik.containo.us/v1alpha1
kind: IngressRoute
metadata:
  name: traefik-service
spec:
  entryPoints:
    - web
  routes:
    - match: Host(`traefik.test.com`) && PathPrefix(`/traefik`)
      kind: Rule
      services:
        - name: treafik-service   # 与svc的name一致
          port: 8080      # 与svc的port一致
```

>测试镜像代码:  
> [raefik-api](https://github.com/costa92/traefik-api)  
> [docker image](https://hub.docker.com/r/costa92/treafik-api/tags)