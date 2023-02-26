package logic

import (
	"context"

	"github.com/xh-polaris/meowchat-moment-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-moment-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchMomentByCommunityIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchMomentByCommunityIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchMomentByCommunityIdLogic {
	return &SearchMomentByCommunityIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchMomentByCommunityIdLogic) SearchMomentByCommunityId(in *pb.SearchMomentByCommunityIdReq) (*pb.ListMomentResp, error) {
	data, total, err := l.svcCtx.MomentModel.SearchByCommunityId(l.ctx, in.CommunityId, in.Keyword, in.Count, in.Skip)
	if err != nil {
		return nil, err
	}
	res := make([]*pb.Moment, 0, in.Count)
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
	return &pb.ListMomentResp{Moments: res, Total: total}, nil
}
