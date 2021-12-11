package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/golang/glog"
	"k8s.io/api/admission/v1beta1"
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
	var admissionResponse *v1beta1.AdmissionResponse
	ar := v1beta1.AdmissionReview{}
	if err := json.Unmarshal(body, &ar); err != nil {
		glog.Errorf("cant unmarshal data %v", err)
		return
	}
	if r.URL.Path == "/mutate" {
		admissionResponse = whsvr.mutate(&ar)
	} else if r.URL.Path == "/validate" {
		admissionResponse = whsvr.validate(&ar)
	}

	resp, err := json.Marshal(admissionResponse)
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

//validate处理
func (whsvr *WebhookServer) validate(ar *v1beta1.AdmissionReview) *v1beta1.AdmissionResponse {
	//根据不同的资源类型做处理
	switch ar.Request.Kind.Kind {
	case "Deployment":
	case "Service":
	}
	return &v1beta1.AdmissionResponse{}
}
