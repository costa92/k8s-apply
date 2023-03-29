# kustomize 使用

## 文件

```sh 
helloworld
├── base
│   ├── configMap.yaml
│   ├── deployment.yaml
│   ├── kustomization.yaml
|   ├── ingress.yaml
│   └── service.yaml   
└── overlays
    ├── dev
    |   ├── ingress.yaml
    │   ├── map.yaml
    │   └── kustomization.yaml
    └── prod
        ├── kustomization.yaml
        └── map.yaml
```
* **base**
base 指的是一个 kustomization , 任何的 kustomization 包括 overlay (后面提到)，都可以作为另一个 kustomization 的 base (简单理解为基础目录)。base 中描述了共享的内容，如资源和常见的资源配置

* **overlay**
overlay 是一个 kustomization, 它修改(并因此依赖于)另外一个 kustomization. overlay 中的 kustomization指的是一些其它的 kustomization, 称为其 base. 没有 base, overlay 无法使用，并且一个 overlay 可以用作 另一个 overlay 的 base(基础)。简而言之，overlay 声明了与 base 之间的差异。通过 overlay 来维护基于 base 的不同 variants(变体)，例如开发、QA 和生产环境的不同 variants

* **variant**
variant 是在集群中将 overlay 应用于 base 的结果。例如开发和生产环境都修改了一些共同 base 以创建不同的 variant。这些 variant 使用相同的总体资源，并与简单的方式变化，例如 deployment 的副本数、ConfigMap使用的数据源等。简而言之，variant 是含有同一组 base 的不同 kustomization

* **resource**
在 kustomize 的上下文中，resource 是描述 k8s API 对象的 YAML 或 JSON 文件的相对路径。即是指向一个声明了 kubernetes API 对象的 YAML 文件

* **patch**
修改文件的一般说明。文件路径，指向一个声明了 kubernetes API patch 的 YAML 文件


## kustomize 相关命令

查看完整的配置
```sh
$ kustomize build helloworld/base  # build 出来的 YAML 太长就不贴处理了
$ kustomize build helloworld/base | kubectl apply -f -  # 这种方式直接部署在集群中
$ kubectl apply -k # 1.14 版本可以直接使用该命令部署应用于集群中
```

部署不同环境

运行开发环境
```sh 
$ kustomize build helloworld/overlays/dev | kubectl apply -f -   # 或者 kubectl apply -k
or
$ kubectl apply -k helloworld/overlays/dev 
```
运行生产环境
```sh 
$ kustomize build helloworld/overlays/prod | kubectl apply -f -     # 或者 kubectl apply -k
or 
$ kubectl apply -k helloworld/overlays/prod
```

删除不同环境
```sh 
$ kubectl delete -k helloworld/overlays/dev
```


## 修改镜像信息
### 修改镜像的命令
```shell
$ cd helloworld/overlays/prod && kustomize edit set image demo-web=demo-web:v2
```
### 查看镜像的命令之后区别
```shell
kubectl diff -k ./
```
结果:
```shell
➜  prod git:(main) ✗ kubectl diff -k ./
diff -u -N /tmp/LIVE-3767992257/apps.v1.Deployment.prod.the-deployment-prod /tmp/MERGED-2896871326/apps.v1.Deployment.prod.the-deployment-prod
--- /tmp/LIVE-3767992257/apps.v1.Deployment.prod.the-deployment-prod    2023-03-29 19:24:33.919332931 +0800
+++ /tmp/MERGED-2896871326/apps.v1.Deployment.prod.the-deployment-prod  2023-03-29 19:24:33.919332931 +0800
@@ -6,7 +6,7 @@
     kubectl.kubernetes.io/last-applied-configuration: |
       {"apiVersion":"apps/v1","kind":"Deployment","metadata":{"annotations":{},"labels":{"app":"hello"},"name":"the-deployment-prod","namespace":"prod"},"spec":{"replicas":3,"selector":{"matchLabels":{"app":"hello"}},"strategy":{"rollingUpdate":{"maxSurge":"25%","maxUnavailable":"25%"},"type":"RollingUpdate"},"template":{"metadata":{"labels":{"app":"hello"}},"spec":{"containers":[{"command":["/app/server"],"env":[{"name":"ALT_GREEITMG","valueFrom":{"configMapKeyRef":{"key":"altGreeting","name":"the-map-prod"}}},{"name":"ENABLE_RISKY","valueFrom":{"configMapKeyRef":{"key":"enableRisky","name":"the-map-prod"}}}],"image":"demo-web:v1","name":"the-containers","ports":[{"containerPort":8080}]}]}}}}
   creationTimestamp: "2023-03-29T11:18:27Z"
-  generation: 1
+  generation: 2
   labels:
     app: hello
   managedFields:
@@ -145,7 +145,7 @@
             configMapKeyRef:
               key: enableRisky
               name: the-map-prod
-        image: demo-web:v1
+        image: demo-web:v2
         imagePullPolicy: IfNotPresent
         name: the-containers
         ports:
```

### 修改空间
```shell
kustomize edit set namespace prod
```

## 参考文档
1. [kustomize 官网文档](https://kubernetes.io/zh-cn/docs/tasks/manage-kubernetes-objects)  
2. [kustomize examples](https://github.com/kubernetes-sigs/kustomize/tree/master/examples)