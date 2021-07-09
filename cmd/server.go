package main

import (
	"context"
	"errors"
	"net/http"
	"vtp-apis/cmd/config"
	"vtp-apis/cmd/handler"
	"vtp-apis/external/dao"
	"vtp-apis/external/delivery/rest"
	"vtp-apis/external/repository"
	"vtp-apis/external/usecase"
	commonConf "vtp-apis/packages/config"
	"vtp-apis/packages/ginh"
	"vtp-apis/packages/logger/zapl"
	"vtp-apis/packages/shutdown"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run() {
	shutdown.SigtermHandler().RegisterErrorFuncContext(context.Background(), s.httpServer.Shutdown)
	if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		zapl.Error(context.TODO(), "http Server listen and serve error", err)
	}
}

func initServer() *Server {
	conf := config.Config()

	es, err := commonConf.InitElasticClient(conf.Elastic)
	if err != nil {
		panic(err)
	}
	esDAO := dao.NewElasticsearchDAO(es)
	esRepo := repository.NewESRepository(esDAO)
	//or, err := commonConf.InitOracleClient(conf.Oracle)
	//orDAO := dao.NewOracleDAO()
	uCase := usecase.NewVNSaleUserCase(esRepo, nil)

	ginHandler := handler.NewGin(&handler.GinDependencies{
		RunMode:         conf.RunMode,
		HealthCheck:     &ginh.BaseSerializer{},
		VNSaleSerialize: rest.NewVNSaleSerializer(uCase),
	})

	return &Server{
		httpServer: &http.Server{
			Addr:    conf.HTTP.Addr,
			Handler: ginHandler,
		},
	}
}
