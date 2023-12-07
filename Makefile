.PHONY: install-generators clean

install-generators:
	GOOS=darwin; \
	/usr/local/bin/go install k8s.io/code-generator/cmd/client-gen@latest; \
	/usr/local/bin/go install k8s.io/code-generator/cmd/deepcopy-gen@latest

zz_generated.deepcopy.go: pkg/apis/mygroup.example.com/v1alpha1/types.go
	deepcopy-gen --input-dirs github.com/gwuud/my-first-crd/pkg/apis/mygroup.example.com/v1alpha1 \
		-O zz_generated.deepcopy \
		--output-base ../../.. \
		--go-header-file ./hack/boilerplate.go.txt

clientset: pkg/apis/mygroup.example.com/v1alpha1/types.go
	client-gen --clientset-name clientset \
		--input-base "" \
		--input github.com/gwuud/my-first-crd/pkg/apis/mygroup.example.com/v1alpha1 \
		--output-package github.com/gwuud/my-first-crd/pkg/clientset \
		--output-base ../../.. \
		--go-header-file hack/boilerplate.go.txt

clean:
	rm -rvf \
		pkg/apis/mygroup.example.com/v1alpha1/zz_generated.deepcopy.go \
		pkg/clientset
