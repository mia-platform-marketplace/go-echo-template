/*
 * Copyright 2019 Mia srl
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path"
	"syscall"

	"mia_template_service_name_placeholder/helpers"

	"github.com/caarlos0/env/v6"
	swagger "github.com/davidebianchi/gswagger"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mia-platform/glogger"
)

func main() {
	entrypoint(make(chan os.Signal, 1))
	os.Exit(0)
}

func entrypoint(shutdown chan os.Signal) {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		panic(err.Error())
	}

	// Init logger instance.
	log, err := glogger.InitHelper(glogger.InitOptions{Level: cfg.LogLevel})
	if err != nil {
		panic(err.Error())
	}

	// Routing
	router := echo.New()
	router.Pre(middleware.RemoveTrailingSlash())
	router.Use(echo.WrapMiddleware(glogger.RequestMiddlewareLogger(log, []string{"/-/"})))
	StatusRoutes(router, "mia_template_service_name_placeholder", cfg.ServiceVersion)

	swaggerRouter, err := swagger.NewRouter(newSwaggerEcho(router), swagger.Options{
		Context: context.Background(),
		Openapi: &openapi3.T{
			Info: &openapi3.Info{
				Title:   "mia_template_service_name_placeholder",
				Version: cfg.ServiceVersion,
			},
		},
	})
	if err != nil {
		log.Panicf("fails to create swagger router: %s", err.Error())
		return
	}

	if cfg.ServicePrefix != "" && cfg.ServicePrefix != "/" {
		swaggerRouter, err = swaggerRouter.SubRouter(newSwaggerEcho(router), swagger.SubRouterOptions{
			PathPrefix: fmt.Sprintf("%s/", path.Clean(cfg.ServicePrefix)),
		})
		if err != nil {
			log.Panicf("fails to create swagger router: %s", err.Error())
			return
		}
	}

	if err := setupRouter(router, swaggerRouter); err != nil {
		log.Panicf("fails to setup router: %s", err.Error())
		return
	}

	if err = swaggerRouter.GenerateAndExposeSwagger(); err != nil {
		log.Errorf("fails to generate and expose swagger: %s", err.Error())
		return
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", cfg.HTTPPort),
		Handler: router,
	}

	go func() {
		log.WithField("port", cfg.HTTPPort).Info("Starting server")
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// sigterm signal sent from kubernetes
	signal.Notify(shutdown, syscall.SIGTERM)
	// We'll accept graceful shutdowns when quit via  and SIGTERM (Ctrl+/)
	// SIGINT (Ctrl+C), SIGKILL or SIGQUIT will not be caught.
	helpers.GracefulShutdown(srv, shutdown, log, cfg.DelayShutdownSeconds)
}
