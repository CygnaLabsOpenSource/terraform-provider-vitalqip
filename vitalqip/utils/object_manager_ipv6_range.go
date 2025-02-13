package utils

import (
	en "terraform-provider-vitalqip/vitalqip/entities"
)

func (objMgr *ObjectManager) CreateIPv6Range(ipv6Range *en.IPv6Range) (*en.IPv6Range, error) {

	_, err := objMgr.connector.CreateObject(ipv6Range, "qipaddv6range")
	if err != nil {
		return nil, err
	}
	return ipv6Range, err
}

func (objMgr *ObjectManager) GetIPv6Range(query map[string]string) (*en.IPv6RangeResponse, error) {

	ipv6RangeResponse := &en.IPv6RangeResponse{}
	queryParams := en.NewQueryParams(query)
	err := objMgr.connector.GetObject(nil, "qipgetv6range", &ipv6RangeResponse, queryParams)
	return ipv6RangeResponse, err
}

func (objMgr *ObjectManager) DeleteIPv6Range(query map[string]string) error {

	queryParams := en.NewQueryParams(query)
	_, err := objMgr.connector.DeleteObject(nil, "qipdeletev6range", queryParams)
	return err
}

func (objMgr *ObjectManager) UpdateIPv6Range(ipv6Range *en.IPv6Range) (*en.IPv6Range, error) {

	_, err := objMgr.connector.UpdateObject(ipv6Range, "qipmodifyv6range")
	if err != nil {
		return nil, err
	}
	return ipv6Range, nil
}
