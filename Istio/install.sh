# Knative v0.9 在 Istio v1.1.7 验证过
export ISTIO_VERSION=1.1.7
curl -L https://git.io/getLatestIstio | sh -
cd istio-${ISTIO_VERSION}
# 安装 Istio CRD
for i in install/kubernetes/helm/istio-init/files/crd*yaml; do kubectl apply -f $i; done
# 创建 istio-system namespace
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Namespace
metadata:
  name: istio-system
  labels:
    istio-injection: disabled
EOF
# 启用 sidecar 注入的模板
helm template --namespace=istio-system \
  --set sidecarInjectorWebhook.enabled=true \
  --set sidecarInjectorWebhook.enableNamespacesByDefault=true \
  --set global.proxy.autoInject=disabled \
  --set global.disablePolicyChecks=true \
  --set prometheus.enabled=false \
  `# 禁用 mixer prometheus adapter，删除 istio 默认的 metrics` \
  --set mixer.adapters.prometheus.enabled=false \
  `# 禁用 mixer policy check，我们的模板里不使用 policy` \
  --set global.disablePolicyChecks=true \
  `# 将 gateway pod 设置为 1 以规避最终一致性/readiness 问题` \
  --set gateways.istio-ingressgateway.autoscaleMin=1 \
  --set gateways.istio-ingressgateway.autoscaleMax=1 \
  --set gateways.istio-ingressgateway.resources.requests.cpu=500m \
  --set gateways.istio-ingressgateway.resources.requests.memory=256Mi \
  `# 多个 pilot replica 便于伸缩` \
  --set pilot.autoscaleMin=2 \
  `# 将 pilot 追踪采样设置为 100%` \
  --set pilot.traceSampling=100 \
  install/kubernetes/helm/istio \
  > ./istio.yaml
# 部署 istio
kubectl apply -f istio.yaml