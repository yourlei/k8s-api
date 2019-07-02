package main

import (
	"flag"
	"fmt"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 配置 k8s 集群外 kubeconfig 配置文件
	var kubeconfig *string
	kubeconfig = flag.String("kubeconfig", "/etc/kubernetes/admin.conf", "absolute path to the kubeconfig file")
	flag.Parse()

	//在 kubeconfig 中使用当前上下文环境，config 获取支持 url 和 path 方式
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	// 根据指定的 config 创建一个新的 clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 NamespacesGetter 接口方法 Namespaces 返回 NamespaceInterface
	// NamespaceInterface 接口拥有操作 Namespace 资源的方法，例如 Create、Update、Get、List 等方法
	name := "client-go-test"
	namespacesClient := clientset.CoreV1().Namespaces()
	namespace := &apiv1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Status: apiv1.NamespaceStatus{
			Phase: apiv1.NamespaceActive,
		},
	}

	// 创建一个新的 Namespaces
	fmt.Println("Creating Namespaces...")
	result, err := namespacesClient.Create(namespace)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created Namespaces %s on %s\n", result.ObjectMeta.Name, result.ObjectMeta.CreationTimestamp)

	namespceList, err := namespacesClient.List()
	fmt.Printf("%s", namespceList)
	// 获取指定名称的 Namespaces 信息
	// fmt.Println("Getting Namespaces...")
	// result, err = namespacesClient.Get(name, metav1.GetOptions{})
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Name: %s, Status: %s, selfLink: %s, uid: %s\n",
	// 	result.ObjectMeta.Name, result.Status.Phase, result.ObjectMeta.SelfLink, result.ObjectMeta.UID)

	// // 删除指定名称的 Namespaces 信息
	// fmt.Println("Deleting Namespaces...")
	// deletePolicy := metav1.DeletePropagationForeground
	// if err := namespacesClient.Delete(name, &metav1.DeleteOptions{
	// 	PropagationPolicy: &deletePolicy,
	// }); err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Deleted Namespaces %s\n", name)
}
