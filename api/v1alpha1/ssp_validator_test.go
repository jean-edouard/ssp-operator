/*


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

package v1alpha1

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

var _ = Describe("NodeMaintenance Validation", func() {
	var (
		client  client.Client
		objects = make([]runtime.Object, 0)
	)

	JustBeforeEach(func() {
		scheme := runtime.NewScheme()
		// add our own scheme
		SchemeBuilder.AddToScheme(scheme)
		// add more schemes
		v1.AddToScheme(scheme)

		client = fake.NewFakeClientWithScheme(scheme, objects...)
		InitValidator(client)
	})

	Context("creating NodeMaintenance", func() {
		Context("for node already in maintenance", func() {
			BeforeEach(func() {
				// add an SSP CR to fake client
				sspExisting := &SSP{
					Spec: SSPSpec{},
				}
				objects = append(objects, sspExisting)
			})

			It("should be rejected", func() {
				ssp := SSP{}
				err := ssp.ValidateCreate()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring(ErrorSSPExists))
			})
		})
	})
})
