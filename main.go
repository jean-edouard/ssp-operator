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

package main

import (
	"flag"
	"os"
	"path/filepath"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	sspv1alpha1 "kubevirt.io/ssp-operator/api/v1alpha1"
	"kubevirt.io/ssp-operator/controllers"
	"kubevirt.io/ssp-operator/internal/operands/metrics"
	// +kubebuilder:scaffold:imports
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

const (
	// Must match port in config/webhook/ssp_webhook.yaml
	WebhookPort = 8443
	// This is the cert location as configured by OLM
	WebhookCertDir  = "/tmp/k8s-webhook-server/serving-certs"
	WebhookCertName = "tls.crt"
	WebhookKeyName  = "tls.key"
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	utilruntime.Must(sspv1alpha1.AddToScheme(scheme))
	utilruntime.Must(metrics.AddWatchTypesToScheme(scheme))
	// +kubebuilder:scaffold:scheme
}

func main() {
	var metricsAddr string
	var enableLeaderElection bool
	flag.StringVar(&metricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "enable-leader-election", false,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")
	flag.Parse()

	ctrl.SetLogger(zap.New(zap.UseDevMode(true)))

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:             scheme,
		MetricsBindAddress: metricsAddr,
		Port:               9443,
		LeaderElection:     enableLeaderElection,
		LeaderElectionID:   "734f7229.kubevirt.io",
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	if err = (&controllers.SSPReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("controllers").WithName("SSP"),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SSP")
		os.Exit(1)
	}
	if err = (&sspv1alpha1.SSP{}).SetupWebhookWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create webhook", "webhook", "SSP")
		os.Exit(1)
	}
	// +kubebuilder:scaffold:builder

	setupLog.Info("setting up the webhook server")
	if err := setupWebhookServer(mgr); err != nil {
		setupLog.Error(err, "Failed to setup webhook server")
		os.Exit(1)
	}

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}

func setupWebhookServer(mgr manager.Manager) error {
	// Make sure the certificates are mounted, this should be handled by the OLM
	certs := []string{filepath.Join(WebhookCertDir, WebhookCertName), filepath.Join(WebhookCertDir, WebhookKeyName)}
	for _, fname := range certs {
		if _, err := os.Stat(fname); err != nil {
			setupLog.Error(err, "Failed to prepare webhook server, certificates not found")
			return err
		}
	}

	server := mgr.GetWebhookServer()
	server.Port = WebhookPort
	server.CertDir = WebhookCertDir
	server.CertName = WebhookCertName
	server.KeyName = WebhookKeyName

	server.Register("/validate-ssp-kubevirt-io-v1alpha1-ssp", admission.ValidatingWebhookFor(&sspv1alpha1.SSP{}))

	sspv1alpha1.InitValidator(mgr.GetClient())

	return nil
}
