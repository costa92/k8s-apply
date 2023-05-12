# 灰度发布 (GrayRelease)

1. 创建 TraefikService
```yaml
apiVersion: traefik.containo.us/v1alpha1
kind: TraefikService
metadata:
  name: app-wrr
spec:
  weighted:
    services:
      - name:  treafik-service  # 注意这个 demo-dep 中的 service 中名字
        weight: 1
        port: 8080
        kind: Service
      - name: treafik-service-2 # 注意这个 demo-dep-2 中的 service 中名字
        weight: 1
        port: 8080
        kind: Service
```


2. 创建 IngressRoute 
```yaml
apiVersion: traefik.containo.us/v1alpha1
kind: IngressRoute
metadata:
  name: traefik-api-ingress-route
spec:
  entryPoints:
    - web
  routes:
    - match: Host(`traefik.test.com`) && PathPrefix(`/traefik`)  # 访问的路由规则
      kind: Rule
      services:
        - name: app-wrr
          kind: TraefikService
```

