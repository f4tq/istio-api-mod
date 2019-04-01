SHELL:=/bin/bash
makefile_dir := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
current_dir := $(shell pwd)

# istio-api source
src_pkg:='https://github.com/f4tq/istio-api.git'
src_branch:=istio-8772


proto_dir:= /tmp/foo.io/$(src_branch)

pkg:=github.com/f4tq/istio-api-mod

proto_path=.
repo_mount=/go/src/$(pkg)
pwd=$(current_dir)

export all_img='f4tq/kubebuilder:1.0.8'

ifeq ("","$(wildcard /.dockerenv)")
export docker_run='docker run --rm -e LOCAL_BUILD=1 -v $(pwd):$(repo_mount) -w $(repo_mount) $(all_img)'
else
export LOCAL_BUILD=1
export docker_run='set -x;'
endif

export protoc="set -x ;protoc  -I. -I./vendor -I./vendor/github.com/gogo/googleapis -I./vendor/github.com/gogo/protobuf/protobuf -I../.."
export protolock='echo DEATH '
export protoc_gen_docs_plugin=' --docs_out=mode=html_fragment_with_front_matter:./'

comma := ,
empty:=
space := $(empty) $(empty)

kube_istio_source_packages = $(subst $(space),$(empty), \
	github.com/f4tq/istio-api-mod/authentication/..., \
	github.com/f4tq/istio-api-mod/mixer/..., \
	github.com/f4tq/istio-api-mod/networking/..., \
	github.com/f4tq/istio-api-mod/policy/..., \
	github.com/f4tq/istio-api-mod/rbac/... \
	)
kube_base_output_package = github.com/f4tq/istio-api-mod/pkg/kube

kube_proto_api_packages = $(subst $(space),$(empty), \
	-github.com/f4tq/istio-api-mod/authentication/v1alpha1=istio.authentication.v1alpha1, \
	-github.com/f4tq/istio-api-mod/mixer/v1/config/client=istio.mixer.v1.config.client, \
	-github.com/f4tq/istio-api-mod/networking/v1alpha3=istio.networking.v1alpha3, \
	-github.com/f4tq/istio-api-mod/policy/v1beta1=istio.policy.v1beta1, \
	-github.com/f4tq/istio-api-mod/rbac/v1alpha1=istio.rbac.v1alpha1, \
	-k8s.io/apimachinery/pkg/apis/meta/v1, \
	-k8s.io/apimachinery/pkg/runtime/schema \
	)
kube_go_header_text=boilerplate.go.txt

#make generate-routing-go all_img='f4tq/kubebuilder:1.0.8' docker_run='docker run --rm -e LOCAL_BUILD=1 -v $(pwd):$(repo_mount) -w $(repo_mount) $(all_img)' protoc="echo protoc -I. -I./vendor -I./vendor/github.com/gogo/googleapis -I./vendor/github.com/gogo/protobuf/protobuf -I../.." protolock='echo ' protoc_gen_docs_plugin=' --docs_out=mode=html_fragment_with_front_matter:./'

generate: vendor mutate_src
generate: generate-mcp-go \
	generate-mesh-go \
	generate-mixer-go \
	generate-routing-go \
	generate-rbac-go \
	generate-authn-go \
	generate-envoy-go
generate: generate-kube-apis generate-kube-apis-go generate-kube-apis-proto

generate-mcp-go generate-mesh-go generate-mixer-go \
	generate-routing-go \
	generate-rbac-go \
	generate-authn-go \
	generate-envoy-go:

	make -f Makefile.istio out_path=$$GOPATH/src $@
#	make -f Makefile.istio protolock=$(protolock) protoc=$(protoc) protoc_gen_docs_plugin=$(proto_gen_docs_plugin) out_path=$$GOPATH/src $@

generate-kube-%: vendor mutate_src
	make -f Makefile.istio kube_base_output_package=$(kube_base_output_package) kube_proto_api_packages='$(kube_proto_api_packages)' kube_istio_source_packages='$(kube_istio_source_packages)' kube_go_header_text=./$(kube_go_header_text)  $@

generate-%: vendor mutate_src
	make -f Makefile.istio out_path=$$GOPATH/src $@

lint: vendor mutate_src
	make -f Makefile.istio protolock=$(protolock) protoc=$(protoc) protoc_gen_docs_plugin=$(proto_gen_docs_plugin) $@

# change protos to our package name (not the k8 crd name though)
mutate_src:  prototool.yaml
	@echo "Source mutated"

prototool.yaml: $(proto_dir)/prototool.yaml
	@cp $^ $@

Makefile.istio: $(proto_dir)/Makefile
	@set -ex; ls -ld $^; mkdir -p $(proto_path) $(proto_dir) util; \
	cd $(proto_dir)  ;\
	for i in $(shell cd $(proto_dir) ; find . -iname vendor -prune -o -iname pkg -prune -o -iname \*.proto  | grep -v -e pkg -e vendor ); do \
		cp --parents  $$i $(makefile_dir)/$(proto_path) && test -f $(makefile_dir)/$(proto_path)/$$i && \
		grep -l istio $(makefile_dir)/$(proto_path)/$$i && sed -i 's@option\s\+go_package\s*=\s*"istio.io/api@option go_package="$(pkg)@g' $(makefile_dir)/$(proto_path)/$$i ; \
	done ; \
	for i in $(shell cd $(proto_dir) ; find . -iname vendor -prune -o -iname pkg -prune -o -iname zz_custom.deepcopy.go  | grep -v -e pkg -e vendor ); do \
		cp --parents  $$i $(makefile_dir)/$(proto_path) && test -f $(makefile_dir)/$(proto_path)/$$i && \
		sed -i 's@istio.io/api/util@$(pkg)/util@g' $(makefile_dir)/$(proto_path)/$$i ; \
	done ;\
	cp util/deepcopy.go $(makefile_dir)/util ; \
	sed -i '/\n/!N;/\n.*.*error Source/{$d;N;d};P;D' Makefile && \
	sed -i 's/install-k8s-code-generators:/install-k8s-code-generators.old:/g' Makefile && \
	sed -i 's/install-gogo-protoc-gen:/install-gogo-protoc-gen.old:/g' Makefile && \
	sed -i 's@/usr/bin/protoc@/usr/local/bin/protoc@g' Makefile && \
	sed -i 's@warnings=true@warnings=false@g' Makefile && \
	cp Makefile $(makefile_dir)/$@

# 	sed -i 's@:../..@:\$$GOPATH/src@g' Makefile 

# checkout the upstream source
get_src: $(proto_dir)

$(proto_dir)/prototool.yaml $(proto_dir)/Makefile: $(proto_dir)

$(proto_dir):
	@set -x; mkdir -p $@; cd $@ ; git clone $(src_pkg) .  && git checkout $(src_branch) 


vendor:
	dep ensure

prereqs: vendor/vend.info

vendor/vend.info: mutate_src
	go build ./...
	go mod vendor
	vend

help:
	@echo
	@echo "go pkg: $(PKG)"
	@echo "current directory: $(current_dir)"
	@echo "makefile_dir: $(makefile_dir)"
	@echo
	@echo "kubetype_dirs: $(kubetype_dirs)"
	@echo
	@echo "deepcopy_dirs: $(deepcopy_dirs)"
	@echo
	@echo "protoc: $(protoc)"
	@echo
	@echo "gogoslick_plugin: $(gogoslick_plugin)"
	@echo
	@echo "protoc_gen_jsonshim_plugin: $(protoc_gen_jsonshim_plugin)"
	@echo
	@echo "protoc_gen_docs_plugin: $(protoc_gen_docs_plugin) "
	@echo
	@echo "gogo_mapping: $(gogo_mapping)"

