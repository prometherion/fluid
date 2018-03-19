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

package store

import (
	"fmt"

	"k8s.io/client-go/tools/cache"

	"github.com/NCCloud/fluid/internal/ingress"
)

// SSLCertTracker holds a store of referenced Secrets in Ingress rules
type SSLCertTracker struct {
	cache.ThreadSafeStore
}

// NewSSLCertTracker creates a new SSLCertTracker store
func NewSSLCertTracker() *SSLCertTracker {
	return &SSLCertTracker{
		cache.NewThreadSafeStore(cache.Indexers{}, cache.Indices{}),
	}
}

// ByKey searches for an ingress in the local ingress Store
func (s SSLCertTracker) ByKey(key string) (*ingress.SSLCert, error) {
	cert, exists := s.Get(key)
	if !exists {
		return nil, fmt.Errorf("local SSL certificate %v was not found", key)
	}
	return cert.(*ingress.SSLCert), nil
}
