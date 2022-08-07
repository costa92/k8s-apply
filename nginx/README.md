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

### 产生不同的环境
```sh
$ kustomize build base  # build 出来的 YAML 太长就不贴处理了
$ kustomize build base | kubectl apply -f -  # 这种方式直接部署在集群中
$ kubectl apply -k # 1.14 版本可以直接使用该命令部署应用于集群中
```

