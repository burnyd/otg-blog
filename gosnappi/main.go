package main

import (
	"fmt"
	"log"
	"time"

	"github.com/open-traffic-generator/snappi/gosnappi"
)

func main() {
	// Create a new API handle to make API calls against a traffic generator
	api := gosnappi.NewApi()

	// Set the transport protocol to either HTTP or GRPC
	api.NewHttpTransport().SetLocation("https://172.18.100.2:8443")

	// Create a new traffic configuration that will be set on traffic generator
	config := api.NewConfig()
	config.Ports().Add().SetLocation("eth1").SetName("p1")
	config.Ports().Add().SetLocation("eth2").SetName("p2")

	intf1 := config.Devices().Add().SetName("otg1")
	eth1 := intf1.Ethernets().Add().
		SetName("otg1.eth[0]").
		SetMac("02:00:01:01:01:01")
	eth1.Connection().SetPortName("p1")
	eth1.Ipv4Addresses().Add().
		SetName("otg1.eth[0].ipv4[0]").
		SetAddress("10.10.10.2").
		SetPrefix(24).
		SetGateway("10.10.10.1")

	intf2 := config.Devices().Add().SetName("otg2")
	eth2 := intf2.Ethernets().Add().
		SetName("otg2.eth[0]").
		SetMac("02:00:02:01:01:01")
	eth2.Connection().SetPortName("p2")
	eth2.Ipv4Addresses().Add().
		SetName("otg2.eth[0].ipv4[0]").
		SetAddress("20.20.20.2").
		SetPrefix(24).
		SetGateway("20.20.20.1")

	// Configure the flow and set the endpoints
	flow := config.Flows().Add().SetName("flows1")

	flow.TxRx().Device().
		SetRxNames([]string{"otg2.eth[0].ipv4[0]"}).
		SetTxNames([]string{"otg1.eth[0].ipv4[0]"})

	flow.Size().SetFixed(128)
	flow.Duration().FixedPackets().SetPackets(1000)

	pkt := flow.Packet()
	eth := pkt.Add().Ethernet()
	ipv4 := pkt.Add().Ipv4()
	tcp := pkt.Add().Tcp()

	eth.Dst().SetValue("00:11:22:33:44:55")
	eth.Src().SetValue("00:11:22:33:44:66")

	ipv4.Src().SetValue("10.10.10.2")
	ipv4.Dst().SetValue("20.20.20.2")

	tcp.SrcPort().SetValue(5000)
	tcp.DstPort().SetValue(6000)

	// Push traffic configuration constructed so far to traffic generator
	log.Println(config.ToYaml())
	_, err := api.SetConfig(config)
	if err != nil {
		fmt.Print(err)
	}

	ts := gosnappi.NewStateTrafficFlowTransmit()
	ts.SetState(gosnappi.StateTrafficFlowTransmitState.START)

	req := api.NewMetricsRequest()
	req.Port().SetPortNames([]string{"p1"})
	time.Sleep(25 * time.Second)
	metrics, err := api.GetMetrics(req)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(metrics.ToYaml())
	for _, m := range metrics.Msg().PortMetrics {
		fmt.Println("Port ", *m.Name)
		fmt.Println("Received ", *m.BytesRx)
		fmt.Println("Transmitted ", *m.BytesTx)
	}
}
