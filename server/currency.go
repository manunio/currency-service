package server

import (
	"context"
	"github.com/hashicorp/go-hclog"
)

// Currency is a gRPC server it implements the methods defined by CurrencyServer interface
type Currency struct {
	log hclog.Logger
}

func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{log: l}
}

func GetRate(ctx context.Context, rr *protos.RateRee	)  {
	
}


