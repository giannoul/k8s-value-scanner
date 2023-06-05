PWD           = $(shell pwd)
home		  = $(shell )
cluster_name  = $(or $(shell printenv KIND_CLUSTER), local-sandbox)
dev_toolbox   = k8s-value-finder-dev

.SILENT: .use-context-kind

.PHONY: .use-context-kind
.use-context-kind:
	kubectl config use-context kind-$(cluster_name)


.PHONY: kind-create
kind-create:
	kind \
	  create \
	  cluster \
	  --name $(cluster_name) \
	  --config kind-k8s/kind.conf

.PHONY: kind-delete
kind-delete:
	kind \
	  delete \
	  cluster \
	  --name $(cluster_name)
 


.PHONY: .build-dev-toolbox
.build-dev-toolbox:
	docker build -t $(dev_toolbox) -f toolbox/Dockerfile.dev ./toolbox

.PHONY: dev-toolbox
dev-toolbox: .build-dev-toolbox
	docker run --rm -d -v $(PWD)/:/code --network=host --name $(dev_toolbox) $(dev_toolbox) tail -f /dev/null
	docker exec $(dev_toolbox) mkdir -p /home/gopher/.kube
	docker cp $(HOME)/.kube/config $(dev_toolbox):/home/gopher/.kube/config
	docker exec -ti $(dev_toolbox) bash

.PHONY: apply-test-manifests
apply-test-manifests:
	kubectl apply -f test-manifests
