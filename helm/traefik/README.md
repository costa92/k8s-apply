#  Traefik Ingress

## 安装 Traefik

1. 执行以下命令，添加 Traefik 的 Helm chart repo 源。示例如下：
```shell
 helm repo add traefik https://helm.traefik.io/traefik
```

2. 准备安装配置文件 my-traefik.yaml。示例如下：
```yaml
image:
  registry: docker.io
  repository: traefik
  # defaults to appVersion
  tag: "v2.9.10"
  pullPolicy: IfNotPresent
```
具体参考 my-traefik.yaml 文件

> 完整的默认配置可执行 helm show values traefik/traefik 命令查看。  

3. 执行以下命令将 Traefik 安装到 TKE 集群。示例如下：

```shell
kubectl create ns traefik
helm upgrade --install traefik -f  my-traefik.yaml traefik/traefik
```

4. 执行以下命令，获取流量入口的 IP 地址（如下为 EXTERNAL-IP 字段）。示例如下：

```shell
$ kubectl get service -n traefik
NAME      TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)                                     AGE
traefik   NodePort    10.98.249.108   <none>        9000:32437/TCP,80:31279/TCP,443:30084/TCP   13d
```


> 文档说明：
> [在 TKE 上使用 Traefik Ingress](https://cloud.tencent.com/document/product/457/51235)  

> Traefik 更多用法请参见 [Traefik 官方文档](https://doc.traefik.io/traefik/routing/providers/kubernetes-crd/)。  