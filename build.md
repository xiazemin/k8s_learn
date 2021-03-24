docker buildx build \
  --platform linux/arm64 -t banana:v1 .


   4 | >>> ADD  ./main.go /go/src/app/
   5 |     ADD  ./go.mod /go/src/app/
   6 |     WORKDIR /go/src/app/
--------------------
error: failed to solve: rpc error: code = Unknown desc = failed to compute cache key: failed commit on ref "layer-sha256:a34041dfaf60ca49c213d83e89205fc2b7e222817c9728a4aa4f7d3f09d579c9": rename /var/lib/buildkit/runc-overlayfs/content/ingest/29bada6c000594ad3bbf38e2f655f7bea0e29d351a14b1508024d26afbfce58d/data /var/lib/buildkit/runc-overlayfs/content/blobs/sha256/a34041dfaf60ca49c213d83e89205fc2b7e222817c9728a4aa4f7d3f09d579c9: no such file or directory


docker  build  -t banana:v1 .

 => exporting to image                                                                                                                                                0.8s
 => => exporting layers                                                                                                                                               0.7s
 => => writing image sha256:537e690e93423f670676be52a7b3361b436f1ba58996ff4930fb37965a2020fb                                                                          0.0s
 => => naming to docker.io/library/banana:v1

 docker build -t apple:v1 .
  => => exporting layers                                                                                                                                               0.8s
 => => writing image sha256:cf9db1182183ebd3752e31387099d5d18d2a064fabfe513c20b238193fd934ec                                                                          0.0s
 => => naming to docker.io/library/apple:v1



 % kubectl apply -f sample/apple.yaml
pod/apple-app created
service/apple-service created

% kubectl apply -f sample/banana.yaml
pod/banana-app created
service/banana-service created


% kubectl get pods
NAME         READY   STATUS    RESTARTS   AGE
apple-app    1/1     Running   0          11s
banana-app   1/1     Running   0          6s


% curl http://127.0.0.1/apple
apple%
% curl http://127.0.0.1/banana
banana%