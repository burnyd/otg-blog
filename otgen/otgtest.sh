controllerip=$(kubectl get services -n ceos service-https-keng-controller -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
export OTG_API="https://${controllerip}:8443"
otgen create device -n otg1 -p p1 -l eth1 --ip 10.10.10.2 --prefix 30 --gw 10.10.10.1 | \
otgen add device    -n otg2 -p p2 -l eth2 --ip 20.20.20.2 --prefix 30 --gw 20.20.20.1 | \
otgen add flow -n f-1-2 --tx otg1 --rx otg2 --src 10.10.10.2 --dst 20.20.20.2 --count 1000 --rate 200 --size 256 --timestamps --latency ct --proto udp | \
otgen --log info run -k -m flow | otgen transform -m flow | otgen display -m table