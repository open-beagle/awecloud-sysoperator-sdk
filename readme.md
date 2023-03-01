# 替换包路径
github.com/open-beagle/awecloud-sysoperator-sdk/pkg/apis/bcc/v1 => github.com/open-beagle/awecloud-sysoperator-sdk/pkg/apis/bcc/v1
github.com/open-beagle/awecloud-sysoperator-sdk/pkg/traefik/v1alpha1 => github.com/open-beagle/awecloud-sysoperator-sdk/pkg/traefik/v1alpha1


# operator-sdk
operator-sdk-v0.17 generate k8s


## dev

```bash
# 新建一个Tag
git tag v0.6.6-beagle.5

# 推送一个Tag ，-f 强制更新
git push -f origin v0.6.6-beagle.5

# 删除本地Tag
git tag -d v0.6.6-beagle.5
```

## realse

```bash
# 新建一个Tag
git tag v0.6.6-beagle

# 推送一个Tag ，-f 强制更新
git push -f origin v0.6.6-beagle

# 删除本地Tag
git tag -d v0.6.6-beagle
```

# tag
git tag v6.0.1 -f
git push origin v6.0.1 -f

# protos
go-to-protobuf \
  --proto-import="./protos" \
  --packages="github.com/open-beagle/awecloud-sysoperator-sdk/pkg/apis/bcc/v1" \
  --go-header-file "./hack/boilerplate/boilerplate.generatego.txt"