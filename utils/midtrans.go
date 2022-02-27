package utils

import (
	"errors"

	"github.com/midtrans/midtrans-go/coreapi"
)

func CreateTransaction(core coreapi.Client, req *coreapi.ChargeReq) (*coreapi.ChargeResponse, error) {

	apiRes, err := core.ChargeTransaction(req)
	// log.Info(apiRes)

	if err != nil {
		return nil, errors.New("Failed")
	}
	// fmt.Println(apiRes)
	return apiRes, nil
}
