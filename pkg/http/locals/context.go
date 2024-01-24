package locals

import (
	"context"
	"fmt"
	cg "warabiz/api/pkg/constants/general"
)

func GetRequestID(ctx context.Context) string {
	return fmt.Sprintf("%v", ctx.Value(cg.CtxRequestID))
}

// func GetActiveUser(ctx context.Context) *session.UserSession {
// 	activeUser, ok := ctx.Value("active_user").(*session.UserSession)
// 	if !ok || activeUser == nil {
// 		return nil
// 	} else {
// 		return activeUser
// 	}
// }

func GetTimeLoc(ctx context.Context) string {
	return fmt.Sprintf("%v", ctx.Value(cg.CtxTimeZone))
}

// func GetCreator(ctx context.Context) string {
// 	activeUser, ok := ctx.Value("active_user").(*session.UserSession)
// 	if !ok || activeUser == nil {
// 		return "system"
// 	} else {
// 		return fmt.Sprintf("%v", activeUser.UserID)
// 	}
// }
