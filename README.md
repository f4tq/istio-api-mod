# Istio API Adaption for all kubernetes versions

The repository simplifies the use of Istio APIs by keeping the base simple.  Just the protos.

Any assumption beyond the proto directly ties into a particular kubernetes version.

The purpose of this project has been to work with kubebuilder to build custom controllers that work with Istio.  Istio is very difficult to reuse at the moment.

This repository uses go modules to make it possible to address the istio.io/api go naming convention.  
This repository use git subtrees to extract directories from the official istio.io/api project.
This repository works best with the companion project, [f4tq/kubebuilder-dev](https://github.com/f4tq/kubebuilder-dev), that provides all the proto tools, kubebuild, go 1.11 `GOMODULE=on` enablement.

## Building
```
git clone https://github.com/f4tq/istio-api-mod.git
cd istio-api-mod
```
### Using docker image `f4tq/kubebuilder-dev`
- as published
```istio-api-mod $ touch ~/.bash_history-istio-api-mod ;docker run -it --rm -v ~/.bash_history-istio-api-mod:/root/.bash_history  -e HISTFILESIZE=10000 -e HISTSIZE=10000 -v `pwd`:/work -w /work  -e GOPATH=/go f4tq/kubebuilder:1.0.8 
 root # go build ./...


``` 

