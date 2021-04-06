package transformer

import (
	"github.com/pkg/errors"
	"github.com/qtumproject/janus/pkg/eth"
	"github.com/qtumproject/janus/pkg/qtum"
	"github.com/qtumproject/janus/pkg/utils"
	"github.com/shopspring/decimal"
)

type ProxyQTUMGetUTXOs struct {
	*qtum.Qtum
}

var _ ETHProxy = (*ProxyQTUMGetUTXOs)(nil)

func (p *ProxyQTUMGetUTXOs) Method() string {
	return "qtum_getUTXOs"
}

func (p *ProxyQTUMGetUTXOs) Request(req *eth.JSONRPCRequest) (interface{}, error) {
	var params eth.GetUTXOsRequest
	if err := unmarshalRequest(req.Params, &params); err != nil {
		return nil, errors.WithMessage(err, "couldn't unmarshal request parameters")
	}

	err := params.CheckHasValidValues()
	if err != nil {
		return nil, errors.WithMessage(err, "couldn't validate parameters value")
	}

	return p.request(params)
}

func (p *ProxyQTUMGetUTXOs) request(params eth.GetUTXOsRequest) (*qtum.ListUnspentResponse, error) {
	address, err := convertETHAddress(utils.RemoveHexPrefix(params.Address), p.Chain())
	if err != nil {
		return nil, errors.WithMessage(err, "couldn't convert Ethereum address to Qtum address")
	}

	req := qtum.NewListUnspentRequest(qtum.ListUnspentQueryOptions{}, address)
	resp, err := p.Qtum.ListUnspent(req)
	if err != nil {
		return nil, err
	}

	var minUTXOsSum decimal.Decimal
	for _, utxo := range *resp {
		minUTXOsSum = minUTXOsSum.Add(utxo.Amount)

		if minUTXOsSum.GreaterThanOrEqual(params.MinSumAmount) {
			return resp, nil
		}
	}

	return nil, errors.Errorf("required minimum amount is greater than total amount of UTXOs")
}