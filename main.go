package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
)

func main() {
	// Define and parse command line flags
	kubeconfigPath := flag.String("kubeconfig", "/Users/aswina/.kube/config", "kubeconfig file path")
	flag.Parse()

	// Set up the context
	ctx := context.Background()

	// Build Kubernetes config from the specified kubeconfig file
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfigPath)
	handleError(err)

	// Create Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	handleError(err)

	// List pods in the "testclientgo" namespace
	pods, err := clientset.CoreV1().Pods("testclientgo").List(ctx, metav1.ListOptions{})
	handleError(err)

	// Print Pod names
	fmt.Println("========Pods========")
	for _, pod := range pods.Items {
		// Print the name of each Pod
		fmt.Println(pod.Name)
	}

	fmt.Println()

	// List deployments in the "testclientgo" namespace
	fmt.Println("========Deployments========")
	deployments, err := clientset.AppsV1().Deployments("testclientgo").List(ctx, metav1.ListOptions{})
	handleError(err)

	// Print Deployment names
	for _, deployment := range deployments.Items {
		// Print the name of each Deployment
		fmt.Println(deployment.Name)
	}

	fmt.Println()

	// List services in the "testclientgo" namespace
	fmt.Println("========Services========")
	services, err := clientset.CoreV1().Services("testclientgo").List(ctx, metav1.ListOptions{})
	handleError(err)

	// Print Service names
	for _, service := range services.Items {
		// Print the name of each Service
		fmt.Println(service.Name)
	}

	// Get complete GVR for Kubernetes resource - RESTMapper interface
	configFlags := genericclioptions.NewConfigFlags(true).WithDeprecatedPasswordFlag()
	matchVersionFlags := cmdutil.NewMatchVersionFlags(configFlags)
	matchVersion, err := cmdutil.NewFactory(matchVersionFlags).ToRESTMapper()
	handleError(err)

	// Take input from the user about which Kube resource they want to interact with
	fmt.Print("Enter the Kube resource you want to interact with: ")
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	resource := reader.Text()
	handleError(err)

	// Get Group Version Kind for the specified Kube resource
	gvr, err := matchVersion.ResourcesFor(schema.GroupVersionResource{
		Resource: resource,
	})
	handleError(err)

	fmt.Println("Group Version Kind is", gvr)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
