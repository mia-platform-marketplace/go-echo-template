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
	"net/http"

	"github.com/labstack/echo/v4"
)

// StatusResponse type.
type StatusResponse struct {
	Status  string `json:"status"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

func handleStatusRoutes(c echo.Context, serviceName, serviceVersion string) error {
	status := StatusResponse{
		Status:  "OK",
		Name:    serviceName,
		Version: serviceVersion,
	}
	return c.JSON(http.StatusOK, status)
}

// StatusRoutes add status routes to router.
func StatusRoutes(r *echo.Echo, serviceName, serviceVersion string) {
	r.GET("/-/healthz", func(c echo.Context) error {
		return handleStatusRoutes(c, serviceName, serviceVersion)
	})

	r.GET("/-/ready", func(c echo.Context) error {
		return handleStatusRoutes(c, serviceName, serviceVersion)
	})

	r.GET("/-/check-up", func(c echo.Context) error {
		return handleStatusRoutes(c, serviceName, serviceVersion)
	})
}
