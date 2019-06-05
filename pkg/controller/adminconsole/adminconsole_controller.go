package adminconsole

import (
	"admin-console-operator/pkg/service"
	"context"
	"time"

	edpv1alpha1 "admin-console-operator/pkg/apis/edp/v1alpha1"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	logPrint "log"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_adminconsole")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new AdminConsole Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	scheme := mgr.GetScheme()
	client := mgr.GetClient()
	platformService, _ := service.NewPlatformService(scheme)
	adminConsoleService := service.NewAdminConsoleService(platformService, client)

	return &ReconcileAdminConsole{
		client:  client,
		scheme:  scheme,
		service: adminConsoleService,
	}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("adminconsole-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource AdminConsole
	err = c.Watch(&source.Kind{Type: &edpv1alpha1.AdminConsole{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner AdminConsole
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &edpv1alpha1.AdminConsole{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileAdminConsole implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileAdminConsole{}

// ReconcileAdminConsole reconciles a AdminConsole object
type ReconcileAdminConsole struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client  client.Client
	scheme  *runtime.Scheme
	service service.AdminConsoleService
}

// Reconcile reads that state of the cluster for a AdminConsole object and makes changes based on the state read
// and what is in the AdminConsole.Spec
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileAdminConsole) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling AdminConsole")

	// Fetch the AdminConsole instance
	instance := &edpv1alpha1.AdminConsole{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	err = r.service.Install()
	if err != nil {
		logPrint.Printf("[ERROR] Cannot install Admin Console %s. The reason: %s", instance.Name, err)
		return reconcile.Result{RequeueAfter: 10 * time.Second}, nil
	}

	err = r.service.Configure()
	if err != nil {
		logPrint.Printf("[ERROR] Cannot configure Admin Console %s. The reason: %s", instance.Name, err)
		return reconcile.Result{RequeueAfter: 10 * time.Second}, nil
	}

	err = r.service.ExposeConfiguration()
	if err != nil {
		logPrint.Printf("[ERROR] Cannot expose configuration for Admin Console %s. The reason: %s", instance.Name, err)
		return reconcile.Result{RequeueAfter: 10 * time.Second}, nil
	}

	err = r.service.Integration()
	if err != nil {
		logPrint.Printf("[ERROR] Cannot integrate Admin Console %s. The reason: %s", instance.Name, err)
		return reconcile.Result{RequeueAfter: 10 * time.Second}, nil
	}

	return reconcile.Result{}, nil
}
