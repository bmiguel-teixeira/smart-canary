.PHONY: bootstrap setup all

bootstrap:
	kind create cluster --image kindest/node:v1.28.0 --config bootstrap/kind.yaml

setup:
	kubectl -n argocd-system apply -f bootstrap/manifests/argocd.yaml
	kubectl -n argo-rollouts apply -f bootstrap/manifests/argo-rollouts.yaml
	kubectl apply -f bootstrap/manifests/prom-svc.yaml
	kubectl -n argocd-system apply -f bootstrap/manifests/argo-app.yaml
	helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
	helm repo update
	helm upgrade -i prom prometheus-community/kube-prometheus-stack -f bootstrap/prom-values.yaml

build:
	cd server && make build
	kind load docker-image server:1.0 --name smart-canary

install:
	helm template server-charts/ | kubectl apply -f -

all: bootstrap setup build install
	echo done

destroy:
	kind delete cluster --name smart-canary
