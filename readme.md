# 替换包路径
github.com/open-beagle/awecloud-sysoperator-sdk/pkg/apis/bcc/v1 => github.com/open-beagle/awecloud-sysoperator-sdk/pkg/apis/bcc/v1
github.com/open-beagle/awecloud-sysoperator-sdk/pkg/traefik/v1alpha1 => github.com/open-beagle/awecloud-sysoperator-sdk/pkg/traefik/v1alpha1


# operator-sdk
operator-sdk-v0.17 generate k8s

# tag
git tag v0.5.3 -f
git push origin v0.5.3 -f

# protos
go-to-protobuf \
  --proto-import="./protos" \
  --packages="github.com/open-beagle/awecloud-sysoperator-sdk/pkg/apis/bcc/v1" \
  --go-header-file "./hack/boilerplate/boilerplate.generatego.txt"