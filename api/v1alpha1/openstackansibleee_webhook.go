/*
Copyright 2022.

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

//
// Generated by:
//
// operator-sdk create webhook --group ansibleee --version v1alpha1 --kind OpenStackAnsibleEE --programmatic-validation --defaulting
//

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// OpenStackAnsibleEEDefaults -
type OpenStackAnsibleEEDefaults struct {
	ContainerImageURL string
}

var openstackAnsibleEEDefaults OpenStackAnsibleEEDefaults

// log is for logging in this package.
var openstackansibleeelog = logf.Log.WithName("openstackansibleee-resource")

// SetupOpenStackAnsibleEEDefaults - initialize OpenStackAnsibleEE spec defaults for use with either internal or external webhooks
func SetupOpenStackAnsibleEEDefaults(defaults OpenStackAnsibleEEDefaults) {
	openstackAnsibleEEDefaults = defaults
	openstackansibleeelog.Info("OpenStackAnsibleEE defaults initialized", "defaults", defaults)
}

// SetupWebhookWithManager sets up the webhook with the Manager
func (r *OpenStackAnsibleEE) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-ansibleee-openstack-org-v1alpha1-openstackansibleee,mutating=true,failurePolicy=fail,sideEffects=None,groups=ansibleee.openstack.org,resources=openstackansibleees,verbs=create;update,versions=v1alpha1,name=mopenstackansibleee.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &OpenStackAnsibleEE{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *OpenStackAnsibleEE) Default() {
	openstackansibleeelog.Info("default", "name", r.Name)

	r.Spec.Default()
}

// Default - set defaults for this OpenStackAnsibleEE spec
func (spec *OpenStackAnsibleEESpec) Default() {
	if spec.Image == "" {
		spec.Image = openstackAnsibleEEDefaults.ContainerImageURL
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-ansibleee-openstack-org-v1alpha1-openstackansibleee,mutating=false,failurePolicy=fail,sideEffects=None,groups=ansibleee.openstack.org,resources=openstackansibleees,verbs=create;update,versions=v1alpha1,name=vopenstackansibleee.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &OpenStackAnsibleEE{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *OpenStackAnsibleEE) ValidateCreate() error {
	openstackansibleeelog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *OpenStackAnsibleEE) ValidateUpdate(old runtime.Object) error {
	openstackansibleeelog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *OpenStackAnsibleEE) ValidateDelete() error {
	openstackansibleeelog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
