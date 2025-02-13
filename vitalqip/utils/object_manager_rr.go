package utils

import (
	en "terraform-provider-vitalqip/vitalqip/entities"
)

func (objMgr *ObjectManager) CreateRR(rr *en.RR) (*en.RR, error) {
	rrAdded := &en.RR{}
	err := objMgr.connector.CreateObjectWithResponse(rr, &rrAdded, "qipaddrr")
	if err != nil {
		return nil, err
	}
	return rrAdded, err
}

func (objMgr *ObjectManager) GetRR(query map[string]string) (*en.RRResponse, error) {

	rrResponse := &en.RRResponse{}
	queryParams := en.NewQueryParams(query)
	err := objMgr.connector.GetObject(nil, "qipgetrr", &rrResponse, queryParams)
	return rrResponse, err
}

func (objMgr *ObjectManager) DeleteRR(query map[string]string) error {

	queryParams := en.NewQueryParams(query)
	_, err := objMgr.connector.DeleteObject(nil, "qipdeleterr", queryParams)
	return err
}

func (objMgr *ObjectManager) UpdateRR(rr *en.RRUpdate) (*en.RRUpdate, error) {

	_, err := objMgr.connector.UpdateObject(rr, "qipmodifyrr")
	if err != nil {
		return nil, err
	}
	return rr, nil
}
