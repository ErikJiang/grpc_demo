package svc

import (
	"context"
	"github.com/erikjiang/grpc_demo/proto"
	"log"
)

type TeamService struct {
	proto.UnimplementedTeamServiceServer
}

func (ts *TeamService) GetTeamInfo(ctx context.Context, in *proto.TeamRequest) (*proto.TeamResponse, error) {
	log.Printf("team id: %d\n", in.TeamId)

	return &proto.TeamResponse{
		TeamId:   in.TeamId,
		TeamName: "courage team",
		TeamDesc: "Courage is a great team",
	}, nil
}

func (ts *TeamService) AddTeam(ctx context.Context, in *proto.AddRequest) (*proto.AddResponse, error) {
	log.Printf("add team, body: %+v\n", in.Team)

	return &proto.AddResponse{
		Status:  200,
		Message: "add team is successd",
	}, nil
}
