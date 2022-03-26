images=`echo "gcr.io/knative-releases/knative.dev/serving/cmd/default-domain:latest"`
eval $(echo ${images}|
        sed 's/k8s\.gcr\.io/anjia0532\/google-containers/g;s/gcr\.io/anjia0532/g;s/\//\./g;s/ /\n/g;s/anjia0532\./anjia0532\//g' |
        uniq |
        awk '{print "docker pull "$1";"}'
       )

# docker tag docker.io/anjia0532/knative-releases.knative.dev.serving.cmd.default-domain:latest gcr.io/knative-releases/knative.dev/serving/cmd/default-domain:latest
# docker rmi docker.io/anjia0532/knative-releases.knative.dev.serving.cmd.default-domain:latest