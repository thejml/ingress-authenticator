# ingress-authenticator
A Simple Golang based authenticator that can be used with nginx-ingress and other Kubernetes ingresses

# cross-platform support
This image is currently supported on both amd64 and arm64 systems. To build for a system other than yours, buildx can be used as follows:

```
docker buildx build --platform linux/arm64 . -t thejml/ingress-authenticator:6-arm64
```
