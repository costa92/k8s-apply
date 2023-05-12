# 流量复制

## 创建复制 service
1. 创建 TraefikService
```yaml
apiVersion: traefik.containo.us/v1alpha1
kind: TraefikService
metadata:
  name: mirror-from-service
spec:
  mirroring:
    name: treafik-service  # 发送 100% 的请求到 K8S 的 Service "treafik-service
    port: 8080
    mirrors:
      - name: treafik-service-2  # 然后复制 20% 的请求到 treafik-service-2
        port: 8080
        percent: 20  # 百分比
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
    - match: Host(`traefik.test.com`) && PathPrefix(`/traefik`)
      kind: Rule
      services:
        - name: mirror-from-service
          kind: TraefikService

```


## 创建复制 TraefikService
1. 创建TraefikService 复制 
```yaml
apiVersion: traefik.containo.us/v1alpha1
kind: TraefikService
metadata:
  name: mirror-from-traefikservice
  namespace: default
spec:
  mirroring:
    name: mirror-from-service     #流量入口从TraefikService 来   
    kind: TraefikService
    mirrors:
      - name: treafik-service-2
        port: 8080
        percent: 20
```
注意： mirror-from-service 是复制 service 中的已经创建了 TraefikService，所以在这里直接使用的，直接访问的是: treafik-service 

2. 创建IngressRoute 复制 
```yaml
apiVersion: traefik.containo.us/v1alpha1
kind: IngressRoute
metadata:
  name: ingressroute-mirror-traefikservice
  namespace: default
spec:
  entryPoints:
    - web
  routes:
    - match: Host(`traefik.test.com`) && PathPrefix(`/traefik`)
      kind: Rule
      services:
        - name: mirror-from-traefikservice
          kind: TraefikService
```