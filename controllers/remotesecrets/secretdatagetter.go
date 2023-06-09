//
// Copyright (c) 2021 Red Hat, Inc.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package remotesecrets

import (
	"context"
	"errors"
	"fmt"

	"github.com/redhat-appstudio/remote-secret/controllers/remotesecretstorage"

	api "github.com/redhat-appstudio/remote-secret/api/v1beta1"
	"github.com/redhat-appstudio/remote-secret/controllers/bindings"

	"github.com/redhat-appstudio/remote-secret/pkg/secretstorage"
)

type SecretDataGetter struct {
	Storage remotesecretstorage.RemoteSecretStorage
}

func (sb *SecretDataGetter) GetData(ctx context.Context, obj *api.RemoteSecret) (map[string][]byte, string, error) {
	data, err := sb.Storage.Get(ctx, obj)
	if err != nil {
		if errors.Is(err, secretstorage.NotFoundError) {
			return map[string][]byte{}, string(api.RemoteSecretErrorReasonTokenRetrieval), fmt.Errorf("%w: %s", bindings.SecretDataNotFoundError, err.Error())
		}
		return nil, string(api.RemoteSecretErrorReasonTokenRetrieval), fmt.Errorf("failed to get the token data from token storage: %w", err)
	}

	return *data, string(api.RemoteSecretErrorReasonNoError), nil
}

var _ bindings.SecretDataGetter[*api.RemoteSecret] = (*SecretDataGetter)(nil)
