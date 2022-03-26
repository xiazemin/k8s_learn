#serverless/knative/setup/serving-core.yaml
#serverless/knative/setup/serving-crds.yaml

## 替换 @sha256: :latest\nimagePullPolicy: IfNotPresent\n#


grep 'image:' serverless/knative/setup/serving-core.yaml |grep 'gcr.io'
#grep 'image:' serverless/knative/setup/serving-crds.yaml |grep 'gcr.io'

#https://github.com/anjia0532/gcr.io_mirror

# #原镜像
# gcr.io/knative-releases/knative.dev/eventing/cmd/controller:latest
# #转换后镜像
# anjia0532/knative-releases.knative.dev.eventing.cmd.controller:latest

# #下载并重命名镜像
# docker pull anjia0532/knative-releases.knative.dev.eventing.cmd.controller:latest
# docker tag anjia0532/knative-releases.knative.dev.eventing.cmd.controller:latest gcr.io/knative-releases/knative.dev/eventing/cmd/controller:latest
# docker images | grep $(echo gcr.io/knative-releases/knative.dev/eventing/cmd/controller:latest|awk -F':' '{print $1}')

images=`grep 'image:' serverless/knative/setup/serving-core.yaml |grep 'gcr.io'`
eval $(echo ${images}|
        sed 's/k8s\.gcr\.io/anjia0532\/google-containers/g;s/gcr\.io/anjia0532/g;s/\//\./g;s/ /\n/g;s/anjia0532\./anjia0532\//g' |
        uniq |
        awk '{print "docker pull "$1";"}'
       )

for img in $(docker images --format "{{.Repository}}:{{.Tag}}"| grep "anjia0532"); do
  n=$(echo ${img}| awk -F'[/.:]' '{printf "gcr.io/%s",$2}')
  image=$(echo ${img}| awk -F'[/.:]' '{printf "/%s",$3}')
  tag=$(echo ${img}| awk -F'[:]' '{printf ":%s",$2}')
  echo "${n}${image}${tag}"
#   docker tag $img "${n}${image}${tag}"
#   [[ ${n} == "gcr.io/google-containers" ]] && docker tag $img "k8s.gcr.io${image}${tag}"
done

#https://zhuanlan.zhihu.com/p/369106824