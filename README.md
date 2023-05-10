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

### 测试
参考 sample文件中的 apple.yaml、banana.yaml、ingress.yaml

**注意：** 在 mac docker 中容器 hostPort 参数不能生效
```yaml
    ports:
    - name: web
        containerPort: 80
        hostPort: 80  # 不能生效 mac docker 中
        protocol: TCP
```
原因参考:
[Docker 的(Linux/Mac OS)网络配置问题] (https://yuanmomo.net/2019/06/13/docker-network/)



宿主机访问容器，使用 -p 参数映射端口。容器访问宿主机，可以在宿主机使用下面的命令获取 宿主机的 ip 地址：
```sh
ps -ef | grep -i docker | grep -i  "\-\-host\-ip" |awk -F "host-ip" '{print $2}' | awk -F '--lowest-ip' '{print $1}'
```