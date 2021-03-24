% kubectl get secret -n ingress-nginx
NAME                                  TYPE                                  DATA   AGE
default-token-vqhp6                   kubernetes.io/service-account-token   3      5m39s
ingress-nginx-admission-token-fnl7p   kubernetes.io/service-account-token   3      5m38s
ingress-nginx-token-r48pq             kubernetes.io/service-account-token   3      5m39s


 % kubectl describe pod ingress-nginx-controller-84599cfff-k4vq5 -n ingress-nginx
Name:           ingress-nginx-controller-84599cfff-k4vq5

  ----     ------       ----                  ----               -------
  Normal   Scheduled    5m14s                 default-scheduler  Successfully assigned ingress-nginx/ingress-nginx-controller-84599cfff-k4vq5 to docker-desktop
  Warning  FailedMount  64s (x10 over 5m14s)  kubelet            MountVolume.SetUp failed for volume "webhook-cert" : secret "ingress-nginx-admission" not found
  Warning  FailedMount  57s (x2 over 3m11s)   kubelet            Unable to attach or mount volumes: unmounted volumes=[webhook-cert], unattached volumes=[webhook-cert ingress-nginx-token-r48pq]: timed out waiting for the condition


% kubectl get pod -n ingress-nginx
NAME                                       READY   STATUS              RESTARTS   AGE
ingress-nginx-admission-create-5kd47       0/1     CrashLoopBackOff    5          6m33s
ingress-nginx-admission-patch-c7s26        0/1     CrashLoopBackOff    5          6m33s
ingress-nginx-controller-84599cfff-k4vq5   0/1     ContainerCreating   0          6m34s


 % kubectl delete -f ingress-nginx.v4.yaml

 %kubectl delete ns ingress-nginx --force
 删除重启docker

  % kubectl get pod -n ingress-nginx
NAME                                       READY   STATUS      RESTARTS   AGE
ingress-nginx-admission-create-jjnn2       0/1     Completed   0          12s
ingress-nginx-admission-patch-k9n8l        0/1     Completed   1          12s
ingress-nginx-controller-84599cfff-xtx2v   0/1     Running     0          13s

% kubectl apply -f sample/apple.yaml
pod/apple-app unchanged
service/apple-service unchanged
% kubectl apply -f sample/banana.yaml
pod/banana-app created
service/banana-service created
% kubectl apply -f sample/ingress.v1.yaml
ingress.networking.k8s.io/example-ingress created

 % curl http://127.0.0.1/apple
<html>
<head><title>404 Not Found</title></head>

说明nginx 配置正确已经起来了

 % kubectl get pod
NAME         READY   STATUS             RESTARTS   AGE
apple-app    0/1     CrashLoopBackOff   4          15m
banana-app   0/1     CrashLoopBackOff   3          102s

镜像在m1 上运行不了


