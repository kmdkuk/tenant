/*
MIT License

Copyright (c) 2021 kouki

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var tenantlog = logf.Log.WithName("tenant-resource")

func (r *Tenant) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-multitenancy-kmdkuk-com-v1-tenant,mutating=true,failurePolicy=fail,sideEffects=None,groups=multitenancy.kmdkuk.com,resources=tenants,verbs=create;update,versions=v1,name=mtenant.kb.io,admissionReviewVersions={v1,v1beta1}

var _ webhook.Defaulter = &Tenant{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Tenant) Default() {
	tenantlog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-multitenancy-kmdkuk-com-v1-tenant,mutating=false,failurePolicy=fail,sideEffects=None,groups=multitenancy.kmdkuk.com,resources=tenants,verbs=create;update,versions=v1,name=vtenant.kb.io,admissionReviewVersions={v1,v1beta1}

var _ webhook.Validator = &Tenant{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Tenant) ValidateCreate() error {
	tenantlog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Tenant) ValidateUpdate(old runtime.Object) error {
	tenantlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Tenant) ValidateDelete() error {
	tenantlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
