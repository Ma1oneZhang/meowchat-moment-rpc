package logic

import (
	"context"

	"github.com/xh-polaris/meowchat-moment-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-moment-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMomentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMomentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMomentLogic {
	return &DeleteMomentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteMomentLogic) DeleteMoment(in *pb.DeleteMomentReq) (*pb.DeleteMomentResp, error) {
	err := l.svcCtx.MomentModel.Delete(l.ctx, in.MomentId)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteMomentResp{}, nil
}
