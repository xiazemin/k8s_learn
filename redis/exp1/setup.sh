
kubectl apply -f redis-config.yaml  --force -n devops
kubectl apply -f redis-deploy.yaml -n devops  --force

# 查看pod的ip
kubectl get pods -n devops  -o wide
# 运行一个自删除的redis 容器 测试pod的连通性
docker run -it  --rm  redis:3.0.7 /bin/bash
# 使用reids-cli命令连接测试