package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-moment-rpc/internal/model"

	"github.com/xh-polaris/meowchat-moment-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-moment-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListMomentByCommunityIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListMomentByCommunityIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListMomentByCommunityIdLogic {
	return &ListMomentByCommunityIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func transformMomentSlice(data []*model.Moment) []*pb.Moment {
	res := make([]*pb.Moment, len(data))
	for i, d := range data {
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
		res[i] = m
	}
	return res
}

func (l *ListMomentByCommunityIdLogic) ListMomentByCommunityId(in *pb.ListMomentByCommunityIdReq) (*pb.ListMomentResp, error) {
	data, total, err := l.svcCtx.MomentModel.FindManyByCommunityId(l.ctx, in.CommunityId, in.Count, in.Skip)
	if err != nil {
		return nil, err
	}
	res := transformMomentSlice(data)
	return &pb.ListMomentResp{Moments: res, Total: total}, nil
}
