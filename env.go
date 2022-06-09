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

// config struct with the mapping of desired environment variables.
type config struct {
	LogLevel             string `env:"LOG_LEVEL" envDefault:"info"`
	HTTPPort             string `env:"HTTP_PORT" envDefault:"8080"`
	ServicePrefix        string `env:"SERVICE_PREFIX"`
	ServiceVersion       string `env:"SERVICE_VERSION" envDefault:"latest"`
	DelayShutdownSeconds int    `env:"DELAY_SHUTDOWN_SECONDS" envDefault:"10"`
}
