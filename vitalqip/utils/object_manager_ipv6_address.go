package utils

import (
	en "terraform-provider-vitalqip/vitalqip/entities"
)

func (objMgr *ObjectManager) CreateIPv6Address(IPv6Address *en.IPv6Address) (*en.IPv6Address, error) {

	_, err := objMgr.connector.CreateObject(IPv6Address, "qipaddaddress")
	if err != nil {
		return nil, err
	}

	return IPv6Address, err
}

func (objMgr *ObjectManager) GetIPv6Address(query map[string]string) (*en.IPv6Address, error) {

	IPv6Address := &en.IPv6Address{}
	queryParams := en.NewQueryParams(query)
	err := objMgr.connector.GetObject(nil, "qipgetaddress", &IPv6Address, queryParams)
	return IPv6Address, err
}

func (objMgr *ObjectManager) DeleteIPv6Address(query map[string]string) error {

	queryParams := en.NewQueryParams(query)
	_, err := objMgr.connector.DeleteObject(nil, "qipdeleteaddress", queryParams)
	return err
}

func (objMgr *ObjectManager) UpdateIPv6Address(IPv6Address *en.IPv6Address) (*en.IPv6Address, error) {

	_, err := objMgr.connector.UpdateObject(IPv6Address, "qipmodifyaddress")
	if err != nil {
		return nil, err
	}
	return IPv6Address, nil
}
