# Deploymemt 代码分析

## 添加
1. 部署Deploymemt 文件

```sh 
kubectl create -f nginx/nginx-deploymemt.yaml
```

2. 查看 deployment 
```sh 
kubectl get deploy 
```
输出：
```sh 
NAME    READY   UP-TO-DATE   AVAILABLE   AGE
nginx   3/3     3            3           52m
```

字段含义:
  * UP-TO-DATE: 已经更新到期望状态的副本数。
  * AVAILABLE: 已经可以提供服务的副本数。
  * READY: 可以提供服务的副本数/期望副本数。

3. 查看 ReplicaSet：


```sh 
kubectl get rs --selector=app=nginx 
``` 
输出:
```sh 
NAME              DESIRED   CURRENT   READY   AGE
nginx-9456bbbf9   3         3         3       56m
```

创建了 Deploment之后，集群多了一个ReplicaSet资源，也就是Deploment 管理的其实是 ReplicaSet，而不是直接管理 Pod。



## 更新与回滚

### 修改

1. 修改镜像的信息

```sh
kubectl set image  deployment/nginx nginx=nginx:1.16.1
deployment.apps/nginx image updated
```

2. 查看 Event 信息：
```sh 
kubectl describe deploy nginx

// ....
Events:
  Type    Reason             Age    From                   Message
  ----    ------             ----   ----                   -------
  Normal  ScalingReplicaSet  2m37s  deployment-controller  Scaled up replica set nginx-ff6655784 to 1
  Normal  ScalingReplicaSet  2m32s  deployment-controller  Scaled down replica set nginx-9456bbbf9 to 2
  Normal  ScalingReplicaSet  2m32s  deployment-controller  Scaled up replica set nginx-ff6655784 to 2
  Normal  ScalingReplicaSet  2m31s  deployment-controller  Scaled down replica set nginx-9456bbbf9 to 1
  Normal  ScalingReplicaSet  2m31s  deployment-controller  Scaled up replica set nginx-ff6655784 to 3
  Normal  ScalingReplicaSet  2m30s  deployment-controller  Scaled down replica set nginx-9456bbbf9 to 0

```
从 Event 可以 deployment-controller 通过调整ReplicaSet资源nginx-9456bbbf9 与 nginx-ff6655784 的副本数完成这次滚动更新

```sh
kubectl get rs --selector=app=nginx
NAME              DESIRED   CURRENT   READY   AGE
nginx-9456bbbf9   0         0         0       73m
nginx-ff6655784   3         3         3       7m18s
```

3. 查看历史的回滚版本


```sh 
kubectl rollout history deployment/nginx
deployment.apps/nginx
REVISION  CHANGE-CAUSE
1         <none>
2         <none>
```
CHANGE-CAUSE 该字段为空，这个字段是取 kubernetes.io/change-cause 注释中提取到


添加注释
```sh 
kubectl annotate deployment/nginx kubernetes.io/change-cause="image update to nginx:1.16.1"
```

```sh 
kubectl rollout history deployment/nginx
deployment.apps/nginx
REVISION  CHANGE-CAUSE
1         <none>
2         image update to nginx:1.16.1
```


补齐第一个版本的注释

```sh 
kubectl annotate rs/nginx-9456bbbf9 kubernetes.io/change-cause="image create nginx"
```

不能这么写：


```sh 
kubectl annotate deployment/nginx-9456bbbf9 kubernetes.io/change-cause="image create nginx"
Error from server (NotFound): deployments.apps "nginx-9456bbbf9" not found
```
因为deployment 已经不存在了

### 回滚

1. 回滚到上个版本

```sh 
kubectl rollout undo deployment/nginx
```

2. 查看历史版本详情

```sh 
kubectl rollout history deployment/nginx --revision=4
```

3. 回滚到指定版本

```
kubectl rollout undo deployment/nginx --to-revison=1
```
--to-revison=number: number 是历史记录的 REVISION 


## Deploymemt 其他的特性

* ***minReadySeconds: 默认值为0，表示一个Pod 就绪之后多长时间可以提供服务, 换句话说，如果设置为 1 就是 Pod 就绪之后 1 秒才对外提供服务
* paused: 挂起。
* progressDeadlineSeconds: 默认是600，表示处理一个Deployment 任务的超时时间，比如10分钟到了还没有升级成功，则标记为 failed （失败）状态。
* replicas： 副本数。
* revisionHistoryLimit： 默认是：10，表示保留的历史版本数量。
* selector: 标签选择器。
* strategy：表示Deployment 更新Pod时的替换策略
* template: Pod 模板

注意：
  
strategy 有两个属性，分别是 type 与 rollingUpdate。 Type 的可选择值是”Recreate“ 和 ”RollingUpdate“，默认为”RollingUpdate“。
strategy.rollingUpdate 有两个属性：
  * maxSurge: 表示滚动更新的时候最多可以比期望副本数多几个，数字或者百分比配置都可以,比如 1 表示更新过程最多同时新增 1 个副本。 然后等一个旧副本删除之后才继续增加一个新副本，百分比计算的结果要向上取整。
  * maxUnavailable: 表示滚动更新的时候可以有多少副本不可用，同样是数字或者百分比配置，比如期望副本数为3， 1 表示最多删除副本到剩下2，然后要等所有新副本创建才能继续删除，百分比计算的几个是向下取整。


