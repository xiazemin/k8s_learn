dp=`kubectl get deploy |grep '0/1' |awk '{print $1}' |xargs -I{{}} echo "{{}}.yaml"`
for d in $dp 
do 
echo $d;
kubectl delete -f $d
kubectl apply -f $d
done