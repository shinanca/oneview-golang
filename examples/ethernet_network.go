package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"os"
)

func main() {
	var (
		ClientOV *ov.OVClient
		ethernet_network = "eth1"
		ethernet_network_1 = "eth77"
		ethernet_network_2 = "eth88"
	)
	ovc := ClientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		800)
	ovVer, _ := ovc.GetAPIVersion()
	fmt.Println(ovVer)

	fmt.Println("#................... Ethernet Network by Name ...............#")
	ethernet_nw, err := ovc.GetEthernetNetworkByName(ethernet_network)
	if err != nil {
		fmt.Println(err)
	}
	else {
		fmt.Println(ethernet_nw)
	}

	sort := "name:desc"
	ethernet_nw_list, err := ovc.GetEthernetNetworks("", sort)
	if err != nil {
		fmt.Println(err)
	}
	else {
	fmt.Println("# ................... Ethernet Networks List .................#")
	for i := 0; i < len(ethernet_nw_list.Members); i++ {
		fmt.Println(ethernet_nw_list.Members[i].Name)
	}}

	ethernet_nw_id := "02bbab66-4f23-4297-88fa-5420294ec552"
	fmt.Println("#................... GetAssociatedProfiles ....................#")
	ethernet_nw_ass_pfl, err := ovc.GetAssociatedProfile(ethernet_nw_id)
	if err != nil {
		panic(err)
	}
	else {
	fmt.Println(ethernet_nw_ass_pfl)
        }

	fmt.Println("#................... GetAssociatedUplinkGroups ...............#")
	ethernet_nw_uplinkgrps, err := ovc.GetAssociatedUplinkGroup(ethernet_nw_id)
	if err != nil {
		panic(err)
	}
	else {
	fmt.Println(ethernet_nw_uplinkgrps)
        }

	bandwidth := ov.Bandwidth{MaximumBandwidth: 10000, TypicalBandwidth: 2000}

	ethernetNetwork := ov.EthernetNetwork{Name: "eth77", VlanId: 10, Purpose: "General", SmartLink: false, PrivateNetwork: false, ConnectionTemplateUri: "", EthernetNetworkType: "Tagged", Type: "ethernet-networkV4"}

	bulkEthernetNetwork := ov.BulkEthernetNetwork{VlanIdRange: "2-4", Purpose: "General", NamePrefix: "Test_eth", SmartLink: false, PrivateNetwork: false, Bandwidth: bandwidth, Type: "bulk-ethernet-networkV1"}

	er := ovc.CreateEthernetNetwork(ethernetNetwork)
	if er != nil {
		fmt.Println("............... Ethernet Network Creation Failed:", err)
	}
	else {
	fmt.Println(".... Ethernet Network Created Success")
        }

	err = ovc.CreateBulkEthernetNetwork(bulkEthernetNetwork)
	if err != nil {
		fmt.Println("............. Bulk Ethernet Network Creation Failed:", err)
	}
	else {
	fmt.Println(".... Bulk Ethernet Network Created Success")
        }

	bulk_ethernet_network_list, err := ovc.GetEthernetNetworks("", sort)
	for i := 0; i < len(bulk_ethernet_network_list.Members); i++ {
		fmt.Println(bulk_ethernet_network_list.Members[i].Name)
	}

	ethernet_nw, _ := ovc.GetEthernetNetworkByName(ethernet_network_1)
	ethernet_nw.Name = ethernet_network_2
	err = ovc.UpdateEthernetNetwork(ethernet_nw)
	if err != nil {
		panic(err)
	}
	else {
	fmt.Println("#.................... Ethernet Network after Updating ...........#")
	ethernet_nw_after_update, err := ovc.GetEthernetNetworks("", sort)
	for i := 0; i < len(ethernet_nw_after_update.Members); i++ {
		fmt.Println(ethernet_nw_after_update.Members[i].Name)
	}
        }

	ethernet_network_3 := "ppp"
	err = ovc.DeleteEthernetNetwork(ethernet_network_3)
	if err != nil {
		panic(err)
	}
	else {
	fmt.Println("#...................... Deleted Ethernet Network Successfully .....#")
        }

}
