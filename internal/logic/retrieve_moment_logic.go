package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-moment-rpc/errorx"
	"github.com/xh-polaris/meowchat-moment-rpc/internal/model"

	"github.com/xh-polaris/meowchat-moment-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-moment-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RetrieveMomentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRetrieveMomentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RetrieveMomentLogic {
	return &RetrieveMomentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RetrieveMomentLogic) RetrieveMoment(in *pb.RetrieveMomentReq) (*pb.RetrieveMomentResp, error) {
	data, err := l.svcCtx.MomentModel.FindOne(l.ctx, in.MomentId)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errorx.ErrNoSuchMoment
	default:
		return nil, err
	}
	m := &pb.Moment{
		Id:          data.ID.Hex(),
		CreateAt:    data.CreateAt.Unix(),
		Photos:      data.Photos,
		Title:       data.Title,
		Text:        data.Text,
		UserId:      data.UserId,
		CommunityId: data.CommunityId,
		CatId:       data.CatId,
	}
	return &pb.RetrieveMomentResp{Moment: m}, nil
}
