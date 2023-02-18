# 环境变量(env)

设置 env 的值：valueFrom，由于 Pod 的 name 和 namespace 属于元数据，是在 Pod 创建之前就已经定下来了的，所以我们可以使用 metata 就可以获取到了，但是对于 Pod 的 IP 则不一样，因为我们知道 Pod IP 是不固定的，Pod 重建了就变了，它属于状态数据，所以我们使用 status 这个属性去获取。另外除了使用 fieldRef获取 Pod 的基本信息外，还可以通过 resourceFieldRef 去获取容器的资源请求和资源限制信息。

Pod 创建
```sh
kubectl apply -f env-pod.yaml
```

Pod 创建成功后，我们可以查看日志：

```sh
 kubectl logs env-pod -n kube-system |grep POD 
```

