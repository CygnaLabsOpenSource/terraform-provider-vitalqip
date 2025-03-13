package utils

import (
	en "terraform-provider-vitalqip/vitalqip/entities"
)

func (objMgr *ObjectManager) CreateIPv4Address(ipv4Address *en.IPv4Address) (*en.IPv4Address, error) {

	_, err := objMgr.connector.CreateObject(ipv4Address, "qipaddaddress")
	if err != nil {
		return nil, err
	}
	return ipv4Address, err
}

func (objMgr *ObjectManager) GetIPv4Address(query map[string]string) (*en.IPv4Address, error) {

	ipv4Address := &en.IPv4Address{}
	queryParams := en.NewQueryParams(query)
	err := objMgr.connector.GetObject(nil, "qipgetaddress", &ipv4Address, queryParams)
	return ipv4Address, err
}

func (objMgr *ObjectManager) DeleteIPv4Address(query map[string]string) error {

	queryParams := en.NewQueryParams(query)
	_, err := objMgr.connector.DeleteObject(nil, "qipdeleteaddress", queryParams)
	return err
}

func (objMgr *ObjectManager) UpdateIPv4Address(ipv4Address *en.IPv4Address) (*en.IPv4Address, error) {

	_, err := objMgr.connector.UpdateObject(ipv4Address, "qipmodifyaddress")
	if err != nil {
		return nil, err
	}
	return ipv4Address, nil
}
