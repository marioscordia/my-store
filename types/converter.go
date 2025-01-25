package types

import "github.com/marioscordia/my-store/pkg/gen/aggregator_v1"

func FromInvReqToDistance(req *aggregator_v1.AggregateInvoiceRequest) *Distance {
	return &Distance{
		OBUid: req.ObuID,
		Value: req.Value,
		Unix:  req.Unix,
	}
}

func FromInvToGetInvResponse(invoice *Invoice) *aggregator_v1.GetInvoiceResponse {
	return &aggregator_v1.GetInvoiceResponse{
		ObuID:    invoice.OBUid,
		Distance: invoice.Distance,
		Amount:   invoice.Amount,
	}
}
