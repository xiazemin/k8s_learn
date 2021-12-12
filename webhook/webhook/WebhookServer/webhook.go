package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/golang/glog"
	"k8s.io/api/admission/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 定义WebhookServer
type WebhookServer struct {
	server *http.Server
}

// Webhook Server parameters
type WhSvrParameters struct {
	certFile       string // path to the x509 certificate for https
	keyFile        string // path to the x509 private key matching `CertFile`
	sidecarCfgFile string // path to sidecar injector configuration file
}

// 核心业务逻辑实现
func (whsvr *WebhookServer) serve(w http.ResponseWriter, r *http.Request) {
	//拿到Apiserver传进来的body
	var body []byte
	if r.Body != nil {
		if data, err := ioutil.ReadAll(r.Body); err == nil {
			body = data
		}
	}

	//根据传进来的path判断是mutate还是validate
	var admissionResponse *v1beta1.AdmissionResponse = &v1beta1.AdmissionResponse{}
	ar := v1beta1.AdmissionReview{}
	if err := json.Unmarshal(body, &ar); err != nil {
		glog.Errorf("cant unmarshal data %v", err)
		resp, _ := json.Marshal(
			&v1beta1.AdmissionReview{
				Response: admissionResponse,
				TypeMeta: metav1.TypeMeta{
					APIVersion: "admission.k8s.io/v1",
				},
			})
		if _, err := w.Write(resp); err != nil {
			glog.Errorf("Can't write response: %v", err)
		}
		return
	}
	if r.URL.Path == "/mutate" {
		admissionResponse = whsvr.mutate(&ar)
	} else if r.URL.Path == "/validate" {
		admissionResponse = whsvr.validate(&ar)
	}

	ar.Response = admissionResponse
	resp, err := json.Marshal(ar)
	if err != nil {
		glog.Errorf("Can't marshal response: %v", err)
	}
	//回写response
	glog.Infof("Ready to write reponse ...")
	if _, err := w.Write(resp); err != nil {
		glog.Errorf("Can't write response: %v", err)
	}
}

//mutate处理
func (whsvr *WebhookServer) mutate(ar *v1beta1.AdmissionReview) *v1beta1.AdmissionResponse {
	//根据不同的资源类型做处理
	switch ar.Request.Kind.Kind {
	case "Deployment":
	case "Service":
	}
	return &v1beta1.AdmissionResponse{}
}

/**
{
  "apiVersion": "admission.k8s.io/v1",
  "kind": "AdmissionReview",
  "request": {
    # 唯一标识此准入回调的随机 uid
    "uid": "705ab4f5-6393-11e8-b7cc-42010a800002",

    # 传入完全正确的 group/version/kind 对象
    "kind": {"group":"autoscaling","version":"v1","kind":"Scale"},
    # 修改 resource 的完全正确的的 group/version/kind
    "resource": {"group":"apps","version":"v1","resource":"deployments"},
    # subResource（如果请求是针对 subResource 的）
    "subResource": "scale",

    # 在对 API 服务器的原始请求中，传入对象的标准 group/version/kind
    # 仅当 webhook 指定 `matchPolicy: Equivalent` 且将对 API 服务器的原始请求转换为 webhook 注册的版本时，这才与 `kind` 不同。
    "requestKind": {"group":"autoscaling","version":"v1","kind":"Scale"},
    # 在对 API 服务器的原始请求中正在修改的资源的标准 group/version/kind
    # 仅当 webhook 指定了 `matchPolicy：Equivalent` 并且将对 API 服务器的原始请求转换为 webhook 注册的版本时，这才与 `resource` 不同。
    "requestResource": {"group":"apps","version":"v1","resource":"deployments"},
    # subResource（如果请求是针对 subResource 的）
    # 仅当 webhook 指定了 `matchPolicy：Equivalent` 并且将对 API 服务器的原始请求转换为该 webhook 注册的版本时，这才与 `subResource` 不同。
    "requestSubResource": "scale",

    # 被修改资源的名称
    "name": "my-deployment",
    # 如果资源是属于名字空间（或者是名字空间对象），则这是被修改的资源的名字空间
    "namespace": "my-namespace",

    # 操作可以是 CREATE、UPDATE、DELETE 或 CONNECT
    "operation": "UPDATE",

    "userInfo": {
      # 向 API 服务器发出请求的经过身份验证的用户的用户名
      "username": "admin",
      # 向 API 服务器发出请求的经过身份验证的用户的 UID
      "uid": "014fbff9a07c",
      # 向 API 服务器发出请求的经过身份验证的用户的组成员身份
      "groups": ["system:authenticated","my-admin-group"],
      # 向 API 服务器发出请求的用户相关的任意附加信息
      # 该字段由 API 服务器身份验证层填充，并且如果 webhook 执行了任何 SubjectAccessReview 检查，则应将其包括在内。
      "extra": {
        "some-key":["some-value1", "some-value2"]
      }
    },

    # object 是被接纳的新对象。
    # 对于 DELETE 操作，它为 null。
    "object": {"apiVersion":"autoscaling/v1","kind":"Scale",...},
    # oldObject 是现有对象。
    # 对于 CREATE 和 CONNECT 操作，它为 null。
    "oldObject": {"apiVersion":"autoscaling/v1","kind":"Scale",...},
    # options 包含要接受的操作的选项，例如 meta.k8s.io/v CreateOptions、UpdateOptions 或 DeleteOptions。
    # 对于 CONNECT 操作，它为 null。
    "options": {"apiVersion":"meta.k8s.io/v1","kind":"UpdateOptions",...},

    # dryRun 表示 API 请求正在以 `dryrun` 模式运行，并且将不会保留。
    # 带有副作用的 Webhook 应该避免在 dryRun 为 true 时激活这些副作用。
    # 有关更多详细信息，请参见 http://k8s.io/docs/reference/using-api/api-concepts/#make-a-dry-run-request
    "dryRun": false
  }
}
//https://kubernetes.io/zh/docs/reference/access-authn-authz/extensible-admission-controllers/
*/
//validate处理
func (whsvr *WebhookServer) validate(ar *v1beta1.AdmissionReview) *v1beta1.AdmissionResponse {
	//根据不同的资源类型做处理
	switch ar.Request.Kind.Kind {
	case "Deployment":
	case "Service":
	}

	/**
				{
			  "apiVersion": "admission.k8s.io/v1",
			  "kind": "AdmissionReview",
			  "response": {
			    "uid": "<value from request.uid>",
			    "allowed": true
			  }
			}

			{
		  "apiVersion": "admission.k8s.io/v1",
		  "kind": "AdmissionReview",
		  "response": {
		    "uid": "<value from request.uid>",
		    "allowed": false
		  }
		}

		{
	  "apiVersion": "admission.k8s.io/v1",
	  "kind": "AdmissionReview",
	  "response": {
	    "uid": "<value from request.uid>",
	    "allowed": false,
	    "status": {
	      "code": 403,
	      "message": "You cannot do this because it is Tuesday and your name starts with A"
	    }
	  }
	}
	*/
	glog.Infof(ar.Request.Name, ar.Request.UID, ar.APIVersion, ar.Request.Name, ar.Request.Operation, ar.Request.UserInfo.Username, ar.Request.UserInfo.Extra, ar.Request.UserInfo.Groups)
	return &v1beta1.AdmissionResponse{
		Allowed: true,
		UID:     ar.Request.UID,
	}
}
