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
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

const (
	ErrorSSPExists = "an SSP already exists"
)

// log is for logging in this package.
var ssplog = logf.Log.WithName("ssp-resource")

func (r *SSP) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// introduce an SSPValidator, which gets a k8s client injected
// +k8s:deepcopy-gen=false
type SSPValidator struct {
	client client.Client
}

var validator *SSPValidator

var _ webhook.Validator = &SSP{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *SSP) ValidateCreate() error {
	ssplog.Info("validate create", "name", r.Name)

	if validator == nil {
		return fmt.Errorf("ssp validator isn't initialized yet")
	}

	return validator.ValidateCreate(r)
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *SSP) ValidateUpdate(old runtime.Object) error {
	ssplog.Info("validate update", "name", r.Name)

	if validator == nil {
		return fmt.Errorf("ssp validator isn't initialized yet")
	}

	return validator.ValidateUpdate(r, old.(*SSP))
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *SSP) ValidateDelete() error {
	ssplog.Info("validate delete", "name", r.Name)

	if validator == nil {
		return fmt.Errorf("ssp validator isn't initialized yet")
	}

	return nil
}

// Initialize the SSPValidator
func InitValidator(client client.Client) {
	validator = &SSPValidator{
		client: client,
	}
}

func (v *SSPValidator) ValidateCreate(nm *SSP) error {
	// Validate that no SSP for given node exists yet
	if err := v.validateNoSSPExists(); err != nil {
		ssplog.Info("validation failed", "error", err)
		return err
	}

	return nil
}

func (v *SSPValidator) ValidateUpdate(new, old *SSP) error {
	return nil
}

func (v *SSPValidator) validateNoSSPExists() error {
	var ssps SSPList
	if err := v.client.List(context.TODO(), &ssps, &client.ListOptions{}); err != nil {
		return fmt.Errorf("could not list SSPs, please try again: %v", err)
	}

	if len(ssps.Items) > 0 {
		return fmt.Errorf(ErrorSSPExists)
	}
	return nil
}
