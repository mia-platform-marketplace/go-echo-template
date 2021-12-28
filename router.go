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

	swagger "github.com/davidebianchi/gswagger"
	"github.com/labstack/echo/v4"
)

func setupRouter(e *echo.Echo, router *swagger.Router) error {

	e.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	// Setup your routes here.
	_, err := router.AddRoute(http.MethodGet, "/", func(w http.ResponseWriter, req *http.Request) {
		ctx := e.NewContext(req, w)

		ctx.Error(ctx.String(http.StatusOK, ""))
	}, swagger.Definitions{})
	if err != nil {
		return err
	}
	return nil
}
