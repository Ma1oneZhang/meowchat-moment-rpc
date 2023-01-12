package logic

import (
	"context"

	"github.com/xh-polaris/meowchat-moment-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-moment-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListMomentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListMomentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListMomentLogic {
	return &ListMomentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListMomentLogic) ListMoment(in *pb.ListMomentReq) (*pb.ListMomentResp, error) {
	data, err := l.svcCtx.MomentModel.FindMany(l.ctx, in.CommunityId, in.Count, in.Skip)
	if err != nil {
		return nil, err
	}
	res := make([]*pb.Moment, 0, 20)
	for _, d := range data {
		m := &pb.Moment{
			Id:          d.ID.Hex(),
			CreateAt:    d.CreateAt.Unix(),
			Photos:      d.Photos,
			Title:       d.Title,
			Text:        d.Text,
			UserId:      d.UserId,
			CommunityId: d.CommunityId,
			CatId:       d.CatId,
		}
		res = append(res, m)
	}
	return &pb.ListMomentResp{Moments: res}, nil
}
