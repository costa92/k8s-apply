# livenessProbe  

1. periodSeconds：表示让 kubelet 每隔5秒执行一次存活探针，也就是每5秒执行一次上面的cat /tmp/healthy命令，如果命令执行成功了，将返回0，那么 kubelet 就会认为当前这个容器是存活的，如果返i回的是非0值，那么 kubelet 就会把该容器杀掉然后重启它。默认是10秒，最小1秒

2. initialDelaySeconds：表示在第一次执行探针的时候要等待5秒，这样能够确保我们的容器能够有足够的时间启动起来。大家可以想象下，如果你的第一次执行探针等候的时间太短，是不是很有可能容器还没正常启动起来，所以存活探针很可能始终都是失败的，这样就会无休止的重启下去了，对吧？