# OTG traffic generator examples for KNE / Kind.


## Create the KIND cluster with the cluster.yaml file
kind create cluster --config cluster.yaml
## Create meshnet daemonset.
kubectl apply -k /home/burnyd/projects/meshnet-cni/manifests/base
## Apply ceos operator
kubectl apply -f arista-ceoslab-operator/config/kustomized/manifest.yaml
## Apply metalblb
kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.13.7/config/manifests/metallb-native.yaml
## IPaddress pool for kind..
kubectl apply -f metallb-config.yaml
## Install the operators for
kubectl apply -f manifests/.
## Move ceos images
ceoslab                                  4.31.0F       ecc52bdfdf25   7 weeks ago     2.42GB
kind load docker-image  ceoslab:4.31.0F
## Apply KNE Binary
git pull kne && cd kne_cli && go build.. then move file here
./kne_cli create ceos.pb.txt
kubectl get pods -n ceos
NAME   READY   STATUS    RESTARTS   AGE
r1     1/1     Running   0          32s
r2     1/1     Running   0          32s

## Install the otgen binary on your system if you would like.
bash -c "$(curl -sL https://get.otgcdn.net/otgen)"

which otgen
/usr/local/bin/otgen

## Run the otgen script
```
cd otgen
➜  otgen source otgtest.sh
INFO[0000] Applying OTG config...
INFO[0000] ready.
INFO[0000] Starting protocols...
INFO[0000] waiting for protocols to come up...
INFO[0000] Starting traffic...
INFO[0000] started...
INFO[0000] Total packets to transmit: 1000, ETA is: 5s
+-------+-----------+-----------+
| NAME  | FRAMES TX | FRAMES RX |
+-------+-----------+-----------+
| f-1-2 |      1000 |      1000 |
+-------+-----------+-----------+

INFO[0005] Stopping traffic...
INFO[0005] stopped.
INFO[0005] Stopping protocols...
INFO[0005] stopped.
```

## Run it in go with gosnappi
```
cd gosnappi
go run main.go
..truncated output
2024/03/01 11:10:42 choice: port_metrics
port_metrics:
- bytes_rx: "243"
  bytes_rx_rate: 0
  bytes_tx: "243"
  bytes_tx_rate: 0
  capture: stopped
  frames_rx: "2"
  frames_rx_rate: 0
  frames_tx: "0"
  frames_tx_rate: 0
  link: up
  location: service-keng-port-eth1.ceos.svc.cluster.local:5555;1
  name: p1
  transmit: stopped
 <nil>
Port  p1
Received  243
Transmitted  243
```