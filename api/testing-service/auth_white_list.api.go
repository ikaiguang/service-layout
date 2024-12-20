package testingapi

import (
	_ "github.com/gorilla/websocket"
	middlewareutil "github.com/ikaiguang/go-srv-kit/service/middleware"
	servicev1 "github.com/ikaiguang/service-layout/api/testing-service/v1/services"
)

// GetAuthWhiteList 验证白名单
func GetAuthWhiteList() map[string]middlewareutil.TransportServiceKind {
	// 白名单
	whiteList := make(map[string]middlewareutil.TransportServiceKind)

	// 测试
	whiteList[servicev1.OperationSrvTestdataPost] = middlewareutil.TransportServiceKindALL
	whiteList[servicev1.OperationSrvTestdataGet] = middlewareutil.TransportServiceKindALL
	whiteList[servicev1.OperationSrvTestdataPut] = middlewareutil.TransportServiceKindALL
	whiteList[servicev1.OperationSrvTestdataDelete] = middlewareutil.TransportServiceKindALL
	whiteList[servicev1.OperationSrvTestdataPatch] = middlewareutil.TransportServiceKindALL
	whiteList[servicev1.OperationSrvTestdataWebsocket] = middlewareutil.TransportServiceKindALL

	return whiteList
}
