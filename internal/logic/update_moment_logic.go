package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-moment-rpc/errorx"
	"github.com/xh-polaris/meowchat-moment-rpc/internal/model"
	"github.com/xh-polaris/meowchat-moment-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-moment-rpc/pb"
	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateMomentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMomentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMomentLogic {
	return &UpdateMomentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMomentLogic) UpdateMoment(in *pb.UpdateMomentReq) (*pb.UpdateMomentResp, error) {
	m := in.Moment
	momentId, err := primitive.ObjectIDFromHex(m.Id)
	if err != nil {
		return nil, errorx.ErrInvalidObjectId
	}

	err = l.svcCtx.MomentModel.Update(l.ctx, &model.Moment{
		ID:          momentId,
		CatId:       m.CatId,
		CommunityId: m.CommunityId,
		Photos:      m.Photos,
		Title:       m.Title,
		Text:        m.Text,
		UserId:      m.UserId,
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateMomentResp{}, nil
}
