% kubectl apply -f deployment/job.yaml
job.batch/alpine-job created

% kubectl get job
NAME                             COMPLETIONS   DURATION   AGE
alpine-job                       0/1           8s         8s

% kubectl describe job alpine-job 
Name:           alpine-job
Namespace:      default
Selector:       controller-uid=30d935b7-2c64-4968-a366-bea8a7e3407e
Labels:         controller-uid=30d935b7-2c64-4968-a366-bea8a7e3407e
                job-name=alpine-job
Annotations:    <none>
Parallelism:    1
Completions:    1
Start Time:     Sun, 05 Dec 2021 20:38:42 +0800
Pods Statuses:  1 Running / 0 Succeeded / 0 Failed
Pod Template:
  Labels:  controller-uid=30d935b7-2c64-4968-a366-bea8a7e3407e
           job-name=alpine-job
  Containers:
   counter:
    Image:      alpine:alpine
    Port:       <none>
    Host Port:  <none>
    Command:
      bin/sh
      -c
      for i in 9 8 7 6 5 4 3 2 1; do echo $i; done
    Environment:  <none>
    Mounts:       <none>
  Volumes:        <none>
Events:
  Type    Reason            Age   From            Message
  ----    ------            ----  ----            -------
  Normal  SuccessfulCreate  77s   job-controller  Created pod: alpine-job-9pvq2


  % kubectl get pod
NAME                                        READY   STATUS             RESTARTS   AGE
alpine-job-9pvq2                            0/1     ImagePullBackOff   0          108s


% kubectl delete -f deployment/job.yaml
job.batch "alpine-job" deleted


  image: alpine:3.13

% kubectl apply -f deployment/job.yaml 
job.batch/alpine-job created

 % kubectl get pod
NAME                                        READY   STATUS      RESTARTS   AGE
alpine-job-xhlqk                            0/1     Completed   0          28s

https://www.cnblogs.com/lvcisco/p/9670100.html

注意Job的RestartPolicy仅支持Never和OnFailure两种，不支持Always，我们知道Job就相当于来执行一个批处理任务，执行完就结束了，如果支持Always的话是不是就陷入了死循环了？注意Job的RestartPolicy仅支持Never和OnFailure两种，不支持Always，我们知道Job就相当于来执行一个批处理任务，执行完就结束了，如果支持Always的话是不是就陷入了死循环了？