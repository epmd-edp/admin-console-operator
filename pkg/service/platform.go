package service

import (
	"github.com/epmd-edp/admin-console-operator/v2/pkg/apis/edp/v1alpha1"
	appsV1Api "github.com/openshift/api/apps/v1"
	coreV1Api "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

type PlatformService interface {
	AddServiceAccToSecurityContext(scc string, ac v1alpha1.AdminConsole) error
	CreateDeployConf(ac v1alpha1.AdminConsole) error
	CreateSecret(ac v1alpha1.AdminConsole, name string, data map[string][]byte) error
	CreateExternalEndpoint(ac v1alpha1.AdminConsole) error
	CreateService(ac v1alpha1.AdminConsole) error
	CreateServiceAccount(ac v1alpha1.AdminConsole) (*coreV1Api.ServiceAccount, error)
	CreateSecurityContext(ac v1alpha1.AdminConsole, sa *coreV1Api.ServiceAccount) error
	CreateUserRole(ac v1alpha1.AdminConsole) error
	CreateUserRoleBinding(ac v1alpha1.AdminConsole, name string, binding string, kind string) error
	GetConfigmap(namespace string, name string) (map[string]string, error)
	GetDisplayName(ac v1alpha1.AdminConsole) (string, error)
	GetSecret(namespace string, name string) (map[string][]byte, error)
	GetAdminConsole(ac v1alpha1.AdminConsole) (*v1alpha1.AdminConsole, error)
	GetDeployConf(ac v1alpha1.AdminConsole) (*appsV1Api.DeploymentConfig, error)
	GenerateDbSettings(ac v1alpha1.AdminConsole) ([]coreV1Api.EnvVar, error)
	GenerateKeycloakSettings(ac v1alpha1.AdminConsole) ([]coreV1Api.EnvVar, error)
	PatchDeployConfEnv(ac v1alpha1.AdminConsole, dc *appsV1Api.DeploymentConfig, env []coreV1Api.EnvVar) error
	UpdateAdminConsole(ac v1alpha1.AdminConsole) (*v1alpha1.AdminConsole, error)
}

func NewPlatformService(scheme *runtime.Scheme) (PlatformService, error) {
	config := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{},
	)
	restConfig, err := config.ClientConfig()
	if err != nil {
		return nil, logErrorAndReturn(err)
	}

	platform := OpenshiftService{}

	err = platform.Init(restConfig, scheme)
	if err != nil {
		return nil, logErrorAndReturn(err)
	}
	return platform, nil
}

func logErrorAndReturn(err error) error {
	log.Printf("[ERROR] %v", err)
	return err
}

func generateLabels(name string) map[string]string {
	return map[string]string{
		"app": name,
	}
}