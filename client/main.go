package main

import (
	"context"
	"github.com/erikjiang/grpc_demo/helper"
	"github.com/erikjiang/grpc_demo/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	// 客户端双向认证设置:
	creds := helper.GetClientCreds()

	//conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure(), grpc.WithBlock())
	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("dial failed, err: %s\n", err.Error())
	}
	userClient := proto.NewUserServiceClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// get user info
	resp, err := userClient.GetUserInfo(ctx, &proto.UserRequest{
		UserId: 987654,
	})
	if err != nil {
		log.Fatalf("get user info failed, err: %s\n", err.Error())
	}
	log.Printf("get user info, response: %v\n", resp)

	// get user detail
	userDetail, err := userClient.GetUserDetail(ctx, &proto.UserRequest{
		UserId: 543876,
		Bank:   proto.BankType_CCB,
	})
	if err != nil {
		log.Fatalf("get user detail, err: %s\n", err.Error())
	}
	log.Printf("get user detail, data: %+v\n", userDetail)

	// get user list
	userList, err := userClient.GetUserList(ctx, &proto.UserCount{Count: 5})
	if err != nil {
		log.Fatalf("get user list failed, err: %s\n", err.Error())
	}
	log.Printf("get user list, data: %+v\n", userList.List)

	users := make([]*proto.UserRequest, 10)
	for i := 0; i < 10; i++ {
		users[i] = &proto.UserRequest{UserId: int32(200 + i)}
	}
	userScores, err := userClient.GetUserScore(ctx, &proto.UsersRequest{Users: users})
	if err != nil {
		log.Fatalf("get users score failed, err: %s\n", err.Error())
	}
	log.Printf("get users score, data: %+v\n", userScores.Users)

	/*  Server Stream  */
	sStream, err := userClient.GetUserScoreByServerStream(ctx, &proto.UsersRequest{Users: users})
	if err != nil {
		log.Fatalf("get users score by server stream failed, err: %s\n", err.Error())
	}

	for {
		usersRes, err := sStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("steam recv failed, err: %s\n", err.Error())
		}
		log.Printf("users score by server stream, data: %+v\n", usersRes.Users)
	}

	/*  Client Stream  */
	cStream, err := userClient.GetUserScoreByClientStream(ctx)
	if err != nil {
		log.Fatalf("get users score by client stream failed, err: %s\n", err.Error())
	}
	usersReq := make([]*proto.UserRequest, 0)
	count := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			usersReq = append(usersReq, &proto.UserRequest{UserId: int32(100+count)})
			count++
		}
		cStream.Send(&proto.UsersRequest{Users: usersReq})
		usersReq = usersReq[0:0]
	}

	usersRes, err := cStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("stream close and recv failed, err: %s\n", err.Error())
	}
	log.Printf("users score by client stream, data: %v\n", usersRes.Users)

	/*  Bidirectional Stream  */
	bStream, err := userClient.GetUserScoreByBidirectionalStream(ctx)
	if err != nil {
		log.Fatalf("get users score by bidirectional stream failed, err: %s\n", err.Error())
	}
	usersReq = make([]*proto.UserRequest, 0)
	count = 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			usersReq = append(usersReq, &proto.UserRequest{UserId: int32(100+count)})
			count++
		}
		err = bStream.Send(&proto.UsersRequest{Users: usersReq})
		if err != nil {
			log.Fatalf("bidirectional stream send failed, err: %s\n", err.Error())
		}
		usersReq = usersReq[0:0]

		usersRes, err = bStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("stream recv failed, err: %s\n", err.Error())
		}
		log.Printf("users score by bidirectional stream, data: %v\n", usersRes.Users)
	}

	// new team service client
	teamClient := proto.NewTeamServiceClient(conn)

	// get team info
	teamInfo, err := teamClient.GetTeamInfo(ctx, &proto.TeamRequest{
		TeamId: 321654,
	})
	if err != nil {
		log.Fatalf("get team info failed, err: %s\n", err.Error())
	}
	log.Printf("get team info, response: %v\n", teamInfo)

	// add team
	addRes, err := teamClient.AddTeam(ctx, &proto.AddRequest{
		Team: &proto.TeamDetail{
			TeamId:    432765,
			TeamName:  "Blue Star",
			MemberNum: 198,
			Business:  "devops",
			Created:   &timestamp.Timestamp{Seconds: time.Now().Unix()},
		},
	})
	if err != nil {
		log.Fatalf("add team failed, err: %s\n", err.Error())
	}
	log.Printf("add team, response: %v\n", addRes)

}
