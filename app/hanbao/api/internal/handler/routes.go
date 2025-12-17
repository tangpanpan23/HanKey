package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest"
	"hanbao-engine/app/hanbao/api/internal/logic"
	"hanbao-engine/app/hanbao/api/internal/svc"
	"hanbao-engine/app/hanbao/api/internal/types"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	// 词根解锁仪式
	server.AddRoute(rest.Route{
		Method:  http.MethodPost,
		Path:    "/api/v1/hanbao/unlock",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			var req types.UnlockRequest
			// 简化的请求处理，实际应该使用go-zero的请求解析
			_, err := logic.NewHanbaoUnlockLogic(serverCtx).HanbaoUnlock(&req)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"message": "unlock successful"}`)) // 简化响应
		},
	})

	// 会话开始
	server.AddRoute(rest.Route{
		Method:  http.MethodPost,
		Path:    "/api/v1/hanbao/session/start",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			var req types.StartSessionRequest
			resp, err := logic.NewHanbaoStartSessionLogic(serverCtx).HanbaoStartSession(&req)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"session_id": "` + resp.SessionID + `", "message": "` + resp.Message + `"}`))
		},
	})

	// 获取关卡
	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/api/v1/hanbao/level/:levelId",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			req := &types.LevelRequest{LevelId: "pron_1_1"} // 简化
			resp, err := logic.NewHanbaoGetLevelLogic(serverCtx).HanbaoGetLevel(req)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"level": "` + resp.Title + `"}`))
		},
	})

	// 提交答案
	server.AddRoute(rest.Route{
		Method:  http.MethodPost,
		Path:    "/api/v1/hanbao/level/:levelId/answer",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			var req types.AnswerRequest
			resp, err := logic.NewHanbaoAnswerLevelLogic(serverCtx).HanbaoAnswerLevel(&req)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			correct := "false"
			if resp.Correct {
				correct = "true"
			}
			w.Write([]byte(`{"correct": ` + correct + `, "explanation": "` + resp.Explanation + `"}`))
		},
	})

	// 获取藏宝图
	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/api/v1/hanbao/session/:sessionId/treasure-map",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			req := &types.TreasureMapRequest{SessionID: "demo"}
			resp, err := logic.NewHanbaoGetTreasureMapLogic(serverCtx).HanbaoGetTreasureMap(req)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"user_id": "` + resp.UserID + `", "total_roots": ` + string(rune(resp.Stats.TotalRoots+'0')) + `}`))
		},
	})

	// 获取推荐
	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/api/v1/hanbao/recommendations/:sessionId",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			req := &types.RecommendationsRequest{SessionID: "demo"}
			resp, err := logic.NewHanbaoGetRecommendationsLogic(serverCtx).HanbaoGetRecommendations(req)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"recommended_count": ` + string(rune(len(resp.RecommendedRoots)+'0')) + `}`))
		},
	})
}
