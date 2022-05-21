# ratel 安装
参考链接:[ratel](https://github.com/dotbalo/ratel-doc)
### 说明
 查看config 文件命令：
 ```sh
 ls -a ~/.kube/config
 ````
 再把config 内容复制到 minikube.config 中


### service.yaml 文件内容

 ```yaml
 - serverName: 'minikube'
  serverAddress: 'https://192.168.64.3:8443'
  #serverAdminUser: 'xxx'
  #serverAdminPassword: 'xxx#'
  serverAdminToken: 'null'
  serverDashboardUrl: "http://192.168.64.3:31207/#"
  production: 'false'
  kubeConfigPath: "/mnt/minikube.config"
  harborConfig: "HarborUrl, HarborUsername, HarborPassword, HarborEmail"
 ```
  serverAddress: 查询方式可以通过 kubectl cluster-info 获取：
  
  ```sh
  Kubernetes control plane is running at https://192.168.64.3:8443
  CoreDNS is running at https://192.168.64.3:8443/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy

  To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.
  ```
  serverDashboardUrl 是安装Dashboard的链接配置，因为 ratel 功能是基于Dashboard的开发的，所以需要安装Dashboard

  kubeConfigPath: 是配置 config

