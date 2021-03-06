/*
 * Copyright © 2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
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
 *
 * @author		Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @copyright 	2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 */

package cli

import (
	"github.com/ory/hydra/config"
)

type Handler struct {
	Clients   *ClientHandler
	Policies  *PolicyHandler
	Keys      *JWKHandler
	Warden    *IntrospectionHandler
	Token     *TokenHandler
	Groups    *GroupHandler
	Migration *MigrateHandler
}

func NewHandler(c *config.Config) *Handler {
	return &Handler{
		Clients:   newClientHandler(c),
		Policies:  newPolicyHandler(c),
		Keys:      newJWKHandler(c),
		Warden:    newIntrospectionHandler(c),
		Token:     newTokenHandler(c),
		Groups:    newGroupHandler(c),
		Migration: newMigrateHandler(c),
	}
}
