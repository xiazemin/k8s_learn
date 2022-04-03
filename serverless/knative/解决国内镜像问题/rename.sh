for img in $(docker images --format "{{.Repository}}:{{.Tag}}"| grep "anjia0532"); do
  n=$(echo ${img}| awk -F'[/.:]' '{printf "gcr.io/%s",$2}')
  image=$(echo ${img}| awk -F'[/.:]' '{printf "/knative.%s/%s/%s/%s",$4,$5,$6,$7}')
  tag=$(echo ${img}| awk -F'[:]' '{printf ":%s",$2}')
  echo "${n}${image}${tag}"
  docker tag $img "${n}${image}${tag}"
   [[ ${n} == "gcr.io/google-containers" ]] && docker tag $img "k8s.gcr.io${image}${tag}"
  docker rmi $img
done

#gcr.io/knative-releases/knative.dev/serving/cmd/queue:latest 
#docker.io/anjia0532/knative-releases.knative.dev.serving.cmd.webhook:latest