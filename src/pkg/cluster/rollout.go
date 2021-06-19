package cluster

import (
	"context"
	"flag"
	"os"
	"path/filepath"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var (
	config    *rest.Config
	clientset *kubernetes.Clientset
	initErr   error
)

func Init() {
	if os.Getenv("INCLUSTER") == "TRUE" {

		config, initErr = rest.InClusterConfig()
		if initErr != nil {
			panic(initErr.Error())
		}

		clientset, initErr = kubernetes.NewForConfig(config)
		if initErr != nil {
			panic(initErr.Error())
		}
	} else {
		// This won't work for everyone
		// TODO: implement `--kubeconfig` flag parameter
		kubeconfig := flag.String("kubeconfig", filepath.Join(homedir.HomeDir(), ".kube", "config"), "(optional) absolute path to the kubeconfig file")

		config, initErr = clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if initErr != nil {
			panic(initErr)
		}
		clientset, initErr = kubernetes.NewForConfig(config)
		if initErr != nil {
			panic(initErr)
		}
	}
}

// RolloutRestart scale deployment replicas to 0 and re-scale
// it back to its initial replicas
func RolloutRestart(deployment, namespace string) error {
	deploymentsClient := clientset.AppsV1().Deployments(namespace)

	_deployment, err := deploymentsClient.Get(context.TODO(), deployment, v1.GetOptions{})
	if err != nil {
		return err
	}

	temp_replicas := _deployment.Spec.Replicas // save initial replicas
	_deployment.Spec.Replicas = int32p(0)
	if _, err := deploymentsClient.Update(context.TODO(), _deployment, v1.UpdateOptions{}); err != nil {
		return err
	}

	_deployment, err = deploymentsClient.Get(context.TODO(), deployment, v1.GetOptions{})
	if err != nil {
		return err
	}

	_deployment.Spec.Replicas = temp_replicas // re-scale to `temp_replicas`
	if _, err := deploymentsClient.Update(context.TODO(), _deployment, v1.UpdateOptions{}); err != nil {
		return err
	}

	return nil
}

func int32p(i int32) *int32 {
	return &i
}
