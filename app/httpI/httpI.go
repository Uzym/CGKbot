package httpI

import (
	"2bot/api"
	"2bot/app/bot"
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	_ "google.golang.org/grpc"
)

type GRPCServer struct {
	api.UnimplementedBotApiServer
}

func (s *GRPCServer) GetLeaderboard(context.Context, *api.LeaderboardWrite) (*api.LeaderboardRead, error) {
	leaders, err := bot.Leaderboard()
	if err != nil {
		return nil, err
	}
	var leadersGRPC []*api.Leader

	for i := 0; i < len(leaders); i++ {
		l := leaders[i]
		leadersGRPC = append(leadersGRPC, &api.Leader{Username: l.GetUsername(), Cnt: int32(l.GetCnt())})
	}

	return &api.LeaderboardRead{Successfully: true, Leaders: leadersGRPC}, nil
}

func (s *GRPCServer) SetNewUser(cnt context.Context, user *api.NewUserWrite) (*api.NewUserRead, error) {
	err := bot.SetNewUser(user.GetId(), user.GetUsername())
	if err != nil {
		return nil, err
	}
	return &api.NewUserRead{Successfully: true}, nil
}

func (s *GRPCServer) GetQuestion(cnt context.Context, q *api.QuestionWrite) (*api.QuestionRead, error) {
	question, err := bot.Question(q.GetChatId())
	if err != nil {
		return nil, err
	}
	return &api.QuestionRead{Successfully: true, ChatId: q.GetChatId(), Question: question}, nil
}

func (s *GRPCServer) PostAnswer(cnt context.Context, ans *api.AnswerWrite) (*api.AnswerRead, error) {
	ok, err := bot.Answer(ans.GetChatId(), ans.GetAnswer(), time.Now())
	if err != nil {
		return nil, err
	}
	return &api.AnswerRead{Successfully: true, ChatId: ans.GetChatId(), Correct: ok}, nil
}

var Port string

func StartHTTP() {
	s := grpc.NewServer()
	srv := &GRPCServer{}
	api.RegisterBotApiServer(s, srv)

	l, err := net.Listen("tcp", Port)
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
