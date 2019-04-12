package server

import (
	"context"

	"v5u.win/add"
)

// 业务逻辑部分
// AddServer struct for AddService
type AddServer struct {
}

// Add method ——> AddServiceClient interface 的实现
func (a AddServer) Add(ctx context.Context, r *add.AddRequest) (*add.AddResponse, error) {
	rs := r.A + r.B
	s := add.AddResponse{
		Result: uint64(rs),
	}
	return &s, nil
}
