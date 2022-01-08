docker images |grep 3.7.7-stretch
docker pull docker.io/library/python:3.7.7-stretch
docker pull docker.io/library/python:3.7.7-stretch
export HTTPS_PROXY=""
docker pull docker.io/library/python:3.7.7-stretch
docker pull ruby:2.7.1-slim-buster
docker pull gradle:4.8.1
docker pull odidev/websphere-liberty:21.0.0.3-full-java11-openj9
 docker pull node:12.18.1-slim
docker pull mariadb:10.7.1-focal
docker pull mongo:4.0.19-xenial

cd istio/samples/bookinfo
src/build-services.sh 1.16.2 docker.io/istio

cd ../../..

cd example 
sh restart.sh 

cd istio
% kubectl apply -f <(istioctl kube-inject -f samples/bookinfo/platform/kube/bookinfo.yaml)
Error: failed to run injection template: requested template "sidecar" not found; have 
service/details unchanged
serviceaccount/bookinfo-details unchanged


% kubectl apply -f samples/bookinfo/platform/kube/bookinfo.yaml
service/details unchanged
serviceaccount/bookinfo-details unchanged
deployment.apps/details-v1 created
service/ratings unchanged
serviceaccount/bookinfo-ratings unchanged
deployment.apps/ratings-v1 created
service/reviews unchanged
serviceaccount/bookinfo-reviews unchanged
deployment.apps/reviews-v1 created
deployment.apps/reviews-v2 created
deployment.apps/reviews-v3 created
service/productpage unchanged
serviceaccount/bookinfo-productpage unchanged
deployment.apps/productpage-v1 created
