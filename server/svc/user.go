package svc

import (
	"context"
	"fmt"
	"github.com/erikjiang/grpc_demo/helper"
	"github.com/erikjiang/grpc_demo/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
	"io"
	"log"
	"time"
)

type UserService struct {
	proto.UnimplementedUserServiceServer
}

func (us *UserService) GetUserInfo(ctx context.Context, in *proto.UserRequest) (*proto.UserResponse, error) {
	log.Printf("user id: %d\n", in.UserId)
	userInfo := &proto.UserResponse{
		UserId:   in.UserId,
		UserName: "Jason",
		UserDesc: "Jason is a engineer",
	}
	return userInfo, nil
}

func (us *UserService) GetUserDetail(ctx context.Context, in *proto.UserRequest) (*proto.UserDetail, error) {
	log.Printf("user id: %d\n", in.UserId)
	userInfo := &proto.UserDetail{
		UserId:   in.UserId,
		UserName: "Jason",
		Address:  "Changjiang road No.290, Chengdu city",
		Company:  "XCloud China Co. Ltd.",
		Score:    95,
		Updated:  &timestamp.Timestamp{Seconds: time.Now().Unix()},
	}
	switch in.Bank {
	case proto.BankType_ABC:
		userInfo.Bank = "Agricultural Bank of China"
		userInfo.Balance = 1200
	case proto.BankType_BCM:
		userInfo.Bank = "Bank of Communications"
		userInfo.Balance = 5000
	case proto.BankType_BOC:
		userInfo.Bank = "Bank Of China"
		userInfo.Balance = 3500
	case proto.BankType_CCB:
		userInfo.Bank = "China Constuction Bank"
		userInfo.Balance = 6800
	case proto.BankType_CEB:
		userInfo.Bank = "China Everbright Bank"
		userInfo.Balance = 2600
	}
	return userInfo, nil
}

func (us *UserService) GetUserList(ctx context.Context, in *proto.UserCount) (*proto.UserList, error) {
	userList := make([]*proto.UserResponse, in.Count)
	for i := 0; i < int(in.Count); i++ {
		name := helper.RandStringRunes(5)
		userList[i] = &proto.UserResponse{
			UserId:   100 + int32(i),
			UserName: name,
			UserDesc: fmt.Sprintf("%s is great person", name),
		}
	}
	return &proto.UserList{List: userList}, nil
}

func (us *UserService) GetUserScore(ctx context.Context, in *proto.UsersRequest) (*proto.UsersResponse, error) {
	userScores := make([]*proto.UserDetail, len(in.Users))

	for idx, user := range in.Users {
		userScores[idx] = &proto.UserDetail{
			UserId: user.UserId,
			Score:  int32(100 + idx),
		}
	}

	return &proto.UsersResponse{Users: userScores}, nil
}

// Server stream
func (us *UserService) GetUserScoreByServerStream(in *proto.UsersRequest, stream proto.UserService_GetUserScoreByServerStreamServer) error {
	userScores := make([]*proto.UserDetail, 2)

	for idx, user := range in.Users {
		userScores = append(userScores, &proto.UserDetail{
			UserId: user.UserId,
			Score:  int32(100 + idx),
		})
		if (idx+1)%2 == 0 {
			err := stream.Send(&proto.UsersResponse{Users: userScores})
			if err != nil {
				log.Fatalf("stream send failed, err: %s\n", err.Error())
			}
			userScores = userScores[0:0]
		}
	}
	if len(userScores) > 0 {
		err := stream.Send(&proto.UsersResponse{Users: userScores})
		if err != nil {
			log.Fatalf("stream send failed, err: %s\n", err.Error())
		}
	}
	return nil
}

// Client stream
func (us *UserService) GetUserScoreByClientStream(stream proto.UserService_GetUserScoreByClientStreamServer) error {
	userScores := []*proto.UserDetail{}
	count := 0
	for {
		usersReq, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		for _, user := range usersReq.Users {
			userScores = append(userScores, &proto.UserDetail{
				UserId: user.UserId,
				Score:  int32(100 + count),
			})
			count++
		}
		log.Printf("get user scrore by client stream: data: %v\n", usersReq.Users)
	}
	err := stream.SendAndClose(&proto.UsersResponse{Users: userScores})
	if err != nil {
		return err
	}
	return nil
}

// BidirectionalStream
func (us *UserService) GetUserScoreByBidirectionalStream(stream proto.UserService_GetUserScoreByBidirectionalStreamServer) error {
	userScores := []*proto.UserDetail{}
	count := 0
	for {
		usersReq, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		for _, user := range usersReq.Users {
			userScores = append(userScores, &proto.UserDetail{
				UserId: user.UserId,
				Score:  int32(100 + count),
			})
			count++
		}
		log.Printf("get user scrore by bidirectional stream: data: %v\n", usersReq.Users)
		err = stream.Send(&proto.UsersResponse{Users: userScores})
		if err != nil {
			return err
		}
		userScores = userScores[0:0]
	}
	return nil
}
