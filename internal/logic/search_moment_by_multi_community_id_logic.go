package logic

import (
	"context"

	"github.com/xh-polaris/meowchat-moment-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-moment-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchMomentByMultiCommunityIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchMomentByMultiCommunityIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchMomentByMultiCommunityIdLogic {
	return &SearchMomentByMultiCommunityIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchMomentByMultiCommunityIdLogic) SearchMomentByMultiCommunityId(in *pb.SearchMomentByMultiCommunityIdReq) (*pb.ListMomentResp, error) {
	data, total, err := l.svcCtx.MomentModel.SearchByMultiCommunityId(l.ctx, in.CommunityIds, in.Keyword, in.Count, in.Skip)
	if err != nil {
		return nil, err
	}
	res := make([]*pb.Moment, in.Count)
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
