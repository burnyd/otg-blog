controllerip=$(kubectl get services -n ceos service-https-keng-controller -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
export OTG_API="https://${controllerip}:8443"

curl -k $OTG_API/config | jq .devices