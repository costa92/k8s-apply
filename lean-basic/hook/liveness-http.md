# liveness http

cnych/liveness 容器的web 服务代码

```go
http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
    duration := time.Now().Sub(started)
    if duration.Seconds() > 10 {
        w.WriteHeader(500)
        w.Write([]byte(fmt.Sprintf("error: %v", duration.Seconds())))
    } else {
        w.WriteHeader(200)
        w.Write([]byte("ok"))
    }
})
```

上面的代码 最开始前 10s 返回状态码200，10s 过后就返回状态码500。所以当容器启动3秒后，kubelet 开始执行健康检查。第一次健康检查会成功，因为是在 10s 之内，但是 10 秒后，健康检查将失败，因为现在返回的是一个错误的状态码了，所以 kubelet 将会杀掉和重启容器。


查看 pod 执行的详情

```sh
kubectl describe pod liveness-http 

Events:
  Type     Reason     Age                 From               Message
  ----     ------     ----                ----               -------
  Normal   Scheduled  15m                 default-scheduler  Successfully assigned default/liveness-http to minikube
  Normal   Pulled     15m                 kubelet            Successfully pulled image "cnych/liveness" in 10.189247083s (10.189254377s including waiting)
  Normal   Pulled     15m                 kubelet            Successfully pulled image "cnych/liveness" in 3.852543295s (3.852628754s including waiting)
  Normal   Pulled     14m                 kubelet            Successfully pulled image "cnych/liveness" in 6.22557966s (6.225588599s including waiting)
  Normal   Created    14m (x3 over 15m)   kubelet            Created container livenss
  Normal   Started    14m (x3 over 15m)   kubelet            Started container livenss
  Normal   Pulling    14m (x4 over 15m)   kubelet            Pulling image "cnych/liveness"
  Warning  Unhealthy  14m (x9 over 15m)   kubelet            Liveness probe failed: HTTP probe failed with statuscode: 500
  Normal   Killing    14m (x3 over 15m)   kubelet            Container livenss failed liveness probe, will be restarted
  Warning  BackOff    35s (x54 over 14m)  kubelet            Back-off restarting failed container livenss in pod liveness-http_default(79380c26-a160-466c-8a2a-5f2c4362580e)
```

注意：

initialDelaySeconds的属性，可以来配置第一次执行探针的等待时间，对于启动非常慢的应用这个参数非常有用，比如 Jenkins、Gitlab 这类应用，但是如何设置一个合适的初始延迟时间呢？这个就和应用具体的环境有关系了，所以这个值往往不是通用的，这样的话可能就会导致一个问题，我们的资源清单在别的环境下可能就会健康检查失败了，为解决这个问题，在 Kubernetes v1.16 版本官方特地新增了一个 startupProbe（启动探针），该探针将推迟所有其他探针，直到 Pod 完成启动为止，使用方法和存活探针一样：

```yaml
startupProbe:
  httpGet:
    path: /healthz
    port: 8080
  failureThreshold: 30  # 尽量设置大点
  periodSeconds: 10
```

另外除了上面的initialDelaySeconds和periodSeconds属性外，探针还可以配置如下几个参数：

timeoutSeconds：探测超时时间，默认1秒，最小1秒。
successThreshold：探测失败后，最少连续探测成功多少次才被认定为成功。默认是 1，但是如果是liveness则必须是 1。最小值是 1。
failureThreshold：探测成功后，最少连续探测失败多少次才被认定为失败。默认是 3，最小值是 1。

