Deployment同样也是Kubernetes系统的一个核心概念，主要职责和RC一样的都是保证Pod的数量和健康，二者大部分功能都是完全一致的，我们可以看成是一个升级版的RC控制器，那Deployment又具备那些新特性呢？

RC的全部功能：Deployment具备上面描述的RC的全部功能
事件和状态查看：可以查看Deployment的升级详细进度和状态
回滚：当升级Pod的时候如果出现问题，可以使用回滚操作回滚到之前的任一版本
版本记录：每一次对Deployment的操作，都能够保存下来，这也是保证可以回滚到任一版本的基础
暂停和启动：对于每一次升级都能够随时暂停和启动

% kubectl apply -f deployment/Deployment.yaml
error: unable to recognize "deployment/Deployment.yaml": no matches for kind "Deployment" in version "app/v1"

 % kubectl apply -f deployment/Deployment.yaml
error: unable to recognize "deployment/Deployment.yaml": no matches for kind "Deployment" in version "extensions/v1beta1"

应该是apps，拼写错误可以用% kubectl api-versions查看

% kubectl apply -f deployment/Deployment.yaml
error: error validating "deployment/Deployment.yaml": error validating data: ValidationError(Deployment.spec): missing required field "selector" in io.k8s.api.apps.v1.DeploymentSpec; if you choose to ignore these errors, turn validation off with --validate=false


没有在spec里面定义选择器，加上
  selector: #标签选择器
    matchLabels:
      app: apple 

 % kubectl apply -f deployment/Deployment.yaml
deployment.apps/apple-deployment created


kubectl create -f https://kubernetes.io/docs/user-guide/nginx-deployment.yaml --record 
## --record参数可以记录命令，我们可以很方便的查看每次 revision 的变化 更新的时候可以记录状态，每一步是使用什么命令进行更新的    

% kubectl get deployment                     
NAME                       READY   UP-TO-DATE   AVAILABLE   AGE
apple-deployment           3/3     3            3           87s


 % kubectl get pods                           
NAME                                        READY   STATUS      RESTARTS   AGE
apple-app                                   1/1     Running     18         102d
apple-deployment-fb8cfb965-bnjhr            1/1     Running     0          106s
apple-deployment-fb8cfb965-fftd8            1/1     Running     0          106s
apple-deployment-fb8cfb965-zznjr            1/1     Running     0          106s


% kubectl scale deployment apple-deployment --replicas 1
deployment.apps/apple-deployment scaled

% kubectl get pods                                      
NAME                                        READY   STATUS      RESTARTS   AGE
apple-app                                   1/1     Running     18         102d
apple-deployment-fb8cfb965-zznjr            1/1     Running     0          2m38s


如果集群支持horizontal pod autoscaling的话，还可以为Deployment设置自动扩展 

kubectl autoscale deployment nginx-deployment --min=10 --max=15 --cpu-percent=80 

更新容器中的镜像

kubectl set image deployment/nginx-deployment nginx=nginx:1.9.1

回滚

% kubectl rollout undo deployment/apple-deployment
error: no rollout history found for deployment "apple-deployment"

 % kubectl edit deployment/apple-deployment
Edit cancelled, no changes made.

7.使用edit命令编辑Deployment

kubectl edit deployment/nginx-deployment
　　8.查看rollout的状态

kubectl rollout status deployment/nginx-deployment
　　9.查看历史RS

kubectl get rs



#设置镜像
kubectl set image deployment/nginx-deployment nginx=nginx:1.91 
#查看当前更新状态
kubectl rollout status deployments nginx-deployment 
kubectl get pods 
#查看可回滚的历史版本
kubectl rollout history deployment/nginx-deployment 
kubectl rollout undo deployment/nginx-deployment
##可以使用--revision参数指定回退到某个历史版本 
kubectl rollout undo deployment/nginx-deployment --to-revision=2  
##暂停 deployment的更新
kubectl rollout pause deployment/nginx-deployment


% kubectl rollout history deployment/apple-deployment 
deployment.apps/apple-deployment 
REVISION  CHANGE-CAUSE
1         <none>

kubectl set image deployment/apple-deployment gorse_server:latest
error: there is no need to specify a resource type as a separate argument when passing arguments in resource/name form (e.g. 'kubectl get resource/<resource_name>' instead of 'kubectl get resource resource/<resource_name>'

kubectl set image deployment apple-deployment gorse_server:latest
Error from server (NotFound): deployments.apps "gorse_server:latest" not found

kubectl set image deployment/apple-deployment apple-app=gorse_server:latest
deployment.apps/apple-deployment image updated


% kubectl rollout history deployment/apple-deployment 
deployment.apps/apple-deployment 
REVISION  CHANGE-CAUSE
1         <none>
2         <none>

%  kubectl rollout undo deployment/apple-deployment
deployment.apps/apple-deployment rolled back

https://www.cnblogs.com/cangqinglang/p/11959300.html

https://cloud.tencent.com/developer/article/1347201