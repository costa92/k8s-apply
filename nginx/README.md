# nginx install
## pod 安装

```
kubectl apply -f nginx-deployment.yaml
```
## service 安装

```
kubectl apply -f service.yml
```

## 查看安装
```
kubectl get pod -A
kubectl get service -A
````

## 利用 Kustomize 管理安装

```
kubectl apply -k .
```

