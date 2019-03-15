"/C=FR/ST=Paris/L=Paris/O=Global Security/OU=IT Department/CN=bobba-best-drink.com"

See this example: https://www.shellhacks.com/create-csr-openssl-without-prompt-non-interactive/

see this raticle

https://medium.com/google-cloud/install-secure-helm-in-gke-254d520061f7

helm install --tls --debug --dry-run -f bobba-helm-chart/values-api.yaml ./bobba-helm-chart

use ingress.enabled = true | false for activating ingress