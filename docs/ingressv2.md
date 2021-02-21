
```
$ kubectl apply -k 'github.com/kubernetes-sigs/service-apis/config/crd?ref=v0.2.0'
```

```
$ go test -v   -tags=e2e -count=1  ./test/conformance/ingressv2/... -run "TestIngressConformance/"
$ go test -v  -tags=e2e -count=1  ./test/conformance/ingressv2/... -run "TestIngressConformance/headers/tags" -enable-beta
```
