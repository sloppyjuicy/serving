/*
Copyright 2022 The Knative Authors

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

// Code generated by injection-gen. DO NOT EDIT.

package certificate

import (
	context "context"

	v1 "github.com/cert-manager/cert-manager/pkg/client/informers/externalversions/certmanager/v1"
	controller "knative.dev/pkg/controller"
	logging "knative.dev/pkg/logging"
	factory "knative.dev/serving/pkg/client/certmanager/injection/informers/factory"
)

// Key is used for associating the Informer inside the context.Context.
type Key struct{}

func WithInformer(ctx context.Context) (context.Context, controller.Informer) {
	f := factory.Get(ctx)
	inf := f.Certmanager().V1().Certificates()
	return context.WithValue(ctx, Key{}, inf), inf.Informer()
}

// Get extracts the typed informer from the context.
func Get(ctx context.Context) v1.CertificateInformer {
	untyped := ctx.Value(Key{})
	if untyped == nil {
		logging.FromContext(ctx).Panic(
			"Unable to fetch github.com/cert-manager/cert-manager/pkg/client/informers/externalversions/certmanager/v1.CertificateInformer from context.")
	}
	return untyped.(v1.CertificateInformer)
}
