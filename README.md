# k8s apply desc

## 安装 k8s 服务

### 安装 ingress 服务

```sh
kubectl apply -f ingress-nginx.yaml
```

参考使用: [ingress](https://kubernetes.io/zh/docs/concepts/services-networking/ingress/)

## 安装服务

### dtm 安装
```sh
kubectl apply -f dtm/deployment.yml -f dtm/configMap.yml -f dtm/service.yml
```
参考项目地址：[GitHub](https://github.com/dtm-labs/dtm%20)
参考文档： [Docs](https://dtm.pub/)

###  