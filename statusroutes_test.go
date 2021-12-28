/*
 * Copyright 2021 Mia srl
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
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestStatusRoutes(testCase *testing.T) {
	testRouter := echo.New()
	serviceName := "my-service-name"
	serviceVersion := "0.0.0"
	StatusRoutes(testRouter, serviceName, serviceVersion)

	testCase.Run("/-/healthz - ok", func(t *testing.T) {
		expectedResponse := fmt.Sprintf("{\"status\":\"OK\",\"name\":\"%s\",\"version\":\"%s\"}", serviceName, serviceVersion)
		responseRecorder := httptest.NewRecorder()
		request, requestError := http.NewRequest(http.MethodGet, "/-/healthz", nil)
		require.NoError(t, requestError, "Error creating the /-/healthz request")

		testRouter.ServeHTTP(responseRecorder, request)
		statusCode := responseRecorder.Result().StatusCode
		require.Equal(t, http.StatusOK, statusCode, "The response statusCode should be 200")

		rawBody := responseRecorder.Result().Body
		body, readBodyError := ioutil.ReadAll(rawBody)
		require.NoError(t, readBodyError)
		require.Equal(t, expectedResponse, strings.TrimSpace(string(body)), "The response body should be the expected one")
	})

	testCase.Run("/-/ready - ok", func(t *testing.T) {
		expectedResponse := fmt.Sprintf("{\"status\":\"OK\",\"name\":\"%s\",\"version\":\"%s\"}", serviceName, serviceVersion)
		responseRecorder := httptest.NewRecorder()
		request, requestError := http.NewRequest(http.MethodGet, "/-/ready", nil)
		require.NoError(t, requestError, "Error creating the /-/ready request")

		testRouter.ServeHTTP(responseRecorder, request)
		statusCode := responseRecorder.Result().StatusCode
		require.Equal(t, http.StatusOK, statusCode, "The response statusCode should be 200")

		rawBody := responseRecorder.Result().Body
		body, readBodyError := ioutil.ReadAll(rawBody)
		require.NoError(t, readBodyError)
		require.Equal(t, expectedResponse, strings.TrimSpace(string(body)), "The response body should be the expected one")
	})

	testCase.Run("/-/check-up - ok", func(t *testing.T) {
		expectedResponse := fmt.Sprintf("{\"status\":\"OK\",\"name\":\"%s\",\"version\":\"%s\"}", serviceName, serviceVersion)
		responseRecorder := httptest.NewRecorder()
		request, requestError := http.NewRequest(http.MethodGet, "/-/check-up", nil)
		require.NoError(t, requestError, "Error creating the /-/check-up request")

		testRouter.ServeHTTP(responseRecorder, request)
		statusCode := responseRecorder.Result().StatusCode
		require.Equal(t, http.StatusOK, statusCode, "The response statusCode should be 200")

		rawBody := responseRecorder.Result().Body
		body, readBodyError := ioutil.ReadAll(rawBody)
		require.NoError(t, readBodyError)
		require.Equal(t, expectedResponse, strings.TrimSpace(string(body)), "The response body should be the expected one")
	})
}
