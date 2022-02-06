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

安装 dtm 的 ingress

```sh
kubectl apply -f ingress/ingress-nginx.yaml
```
修改 /etc/hosts 添加 

```sh
hello-jane.test  192.168.11.10
```

**注意**: hello-jane.test 是自己定义的域名， 192.168.11.10 是节点的 ip


参考项目地址：[GitHub](https://github.com/dtm-labs/dtm)

参考文档： [Docs](https://dtm.pub/)


### mysql-endpoint

mysql 一般不会安装到 k8s 集群中，因此需要通过 endpoint 连接外部服务，

```sh
// 安装 endpoint
kubectl apply -f mysql-endpoint.yaml
// 安装服务
kubectl apply -f mysql-service.yaml
```

参考自定义服务文档： [定义 Service](https://kubernetes.io/zh/docs/concepts/services-networking/service/#%E5%AE%9A%E4%B9%89-service)

### nginx 服务

```sh
kubectl apply -f nginx-deployment.yaml
kubectl apply -f service.yml
```


