package server

import (
	"context"

	"github.com/mratri10/go-rich/db"
	"github.com/mratri10/go-rich/pb/github.com/atri/go-grpc-purchase/pb"
)

type PurchaseServer struct {
	pb.UnimplementedPurchaseServiceServer
	DB *db.DB
}

func (s *PurchaseServer) AddPurchase(ctx context.Context, req *pb.PurchaseRequest) (*pb.PurchaseResponse, error) {
	err := s.DB.AddPurchase(db.TestData{
		CustomerId: req.CustomerId,
		ProductId:  req.ProductId,
		Amount:     int(req.Amount),
	})

	if err != nil {
		return &pb.PurchaseResponse{
			Status:  "FAILED",
			Message: err.Error(),
		}, err
	}
	return &pb.PurchaseResponse{
		Status:  "Success",
		Message: "Purchase added successfully",
	}, nil
}

func (s *PurchaseServer) GetPurchases(ctx context.Context, req *pb.CustomerRequest) (*pb.PurchaseList, error) {
	purchases, err := s.DB.GetPurchases(req.CustomerId)

	if err != nil {
		return nil, err
	}

	var list []*pb.Purchase
	for _, p := range purchases {
		list = append(list, &pb.Purchase{
			ProductId: p.ProductId,
			Amount:    int32(p.Amount),
			Date:      p.Date.Format("2025-01-02 15:45:23"),
		})
	}

	return &pb.PurchaseList{
		Purchases: list,
	}, nil
}
