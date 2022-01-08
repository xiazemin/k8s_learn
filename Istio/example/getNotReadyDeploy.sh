 kubectl get deploy |grep '0/1' |awk '{print $1}' |xargs -I{{}} sh -c "kubectl get deploy {{}} -o yaml > {{}}.yaml" #kubectl get deploy {{}} -o yaml > {{}}.yaml

img=`kubectl get deploy |grep '0/1' |awk '{print $1}' |xargs -I{{}} kubectl get deploy {{}} -o yaml |awk '/.*image:.*/{if($1=="-"){print $3}else{print $2}}'`
for i in $img
do 
echo $i
docker pull $i
done

dp=`kubectl get deploy |grep '0/1' |awk '{print $1}' |xargs -I{{}} echo "{{}}.yaml"`
for d in $dp 
do 
echo $d;
kubectl delete -f $d
kubectl apply -f $d
done