package logic

import (
	"context"

	"github.com/xh-polaris/meowchat-moment-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-moment-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchMomentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchMomentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchMomentLogic {
	return &SearchMomentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchMomentLogic) SearchMoment(in *pb.SearchMomentReq) (*pb.SearchMomentResp, error) {
	data, total, err := l.svcCtx.MomentModel.Search(l.ctx, in.CommunityId, in.Keyword, in.Count, in.Skip)
	if err != nil {
		return nil, err
	}
	res := make([]*pb.Moment, 0, 10)
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
	return &pb.SearchMomentResp{Moments: res, Total: total}, nil
}
