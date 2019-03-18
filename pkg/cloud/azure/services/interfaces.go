/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package services

import (
	"context"
)

// Service is a generic interface used by components offering a type of service.
// example: Network service would offer get/createorupdate/delete.
type Service interface {
	Get(ctx context.Context) (interface{}, error)
	CreateOrUpdate(ctx context.Context) error
	Delete(ctx context.Context) error
}
