package utils

import (
	en "terraform-provider-vitalqip/vitalqip/entities"
)

func (objMgr *ObjectManager) CreateZone(zone *en.Zone) (*en.Zone, error) {
	zoneAdded := &en.Zone{}
	err := objMgr.connector.CreateObjectWithResponse(zone, &zoneAdded, "qipaddzone")
	if err != nil {
		return nil, err
	}
	return zoneAdded, err
}

func (objMgr *ObjectManager) GetZone(query map[string]string) (*en.Zone, error) {

	zoneResponse := &en.Zone{}
	queryParams := en.NewQueryParams(query)
	err := objMgr.connector.GetObject(nil, "qipgetzone", &zoneResponse, queryParams)
	return zoneResponse, err
}

func (objMgr *ObjectManager) DeleteZone(query map[string]string) error {

	queryParams := en.NewQueryParams(query)
	_, err := objMgr.connector.DeleteObject(nil, "qipdeletezone", queryParams)
	return err
}

func (objMgr *ObjectManager) UpdateZone(zone *en.ZoneUpdate) error {

	_, err := objMgr.connector.UpdateObject(zone, "qipmodifyzone")
	if err != nil {
		return err
	}
	return nil
}
