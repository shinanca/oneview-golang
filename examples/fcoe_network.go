package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"os"
	"strconv"
)

func main() {
	var (
		ClientOV      *ov.OVClient
		testName      = "TestFCoENetworkGOsdk"
		new_fcoe_name = "RenamedFCoENetwork"
	)
	apiversion, _ := strconv.Atoi(os.Getenv("ONEVIEW_APIVERSION"))
	ovc := ClientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		apiversion,
		"*")
	initialScopeUris := &[]utils.Nstring{utils.NewNstring("/rest/scopes/7e4f76b0-bb2c-49d2-a641-d785475df423")}
	fcoeNetwork := ov.FCoENetwork{
		Name:                  testName,
		Type:                  "fcoe-networkV4", //The Type value is for API>500.
		VlanId:                201,
		ConnectionTemplateUri: "",
		ManagedSanUri:         "",
		InitialScopeUris:      *initialScopeUris, //added for API>500
	}
	fmt.Println(fcoeNetwork)
	err := ovc.CreateFCoENetwork(fcoeNetwork)
	if err != nil {
		fmt.Println("FCoE Network Creation Failed: ", err)
	} else {
		fmt.Println("FCoE Network created successfully...")
	}

	sort := "name:desc"
	fcoeNetworks, err := ovc.GetFCoENetworks("", sort, "", "")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("#---Get FCoE Networks sorted by name in descending order----#")
		for i := range fcoeNetworks.Members {
			fmt.Println(fcoeNetworks.Members[i].Name)
		}
	}
	fcoeNetwork2, err := ovc.GetFCoENetworkByName(testName)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#-------------Get FCoENetworks by name----------------#")
		fmt.Println(fcoeNetwork2)
	}
	fcoeNetwork2.Name = new_fcoe_name
	err = ovc.UpdateFCoENetwork(fcoeNetwork2)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("FCoENetwork has been updated with name: " + fcoeNetwork2.Name)
	}
	err = ovc.DeleteFCoENetwork(new_fcoe_name)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Deleted FCoENetworks successfully...")
	}

}
