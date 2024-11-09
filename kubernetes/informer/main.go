package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	podInformer()
}

func podInformer() {
	// 加载 KubeConfig 文件
	home := os.Getenv("HOME")
	kubeConfigPath := filepath.Join(home, ".kube", "config")
	// log.Println(kubeConfigPath)

	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		panic(err.Error())
	}

	// 创建 kubernetes 客户端
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// 创建 Informer 工厂
	factory := informers.NewSharedInformerFactory(clientset, 30*time.Second)

	// 创建 Pod Informer
	podInformer := factory.Core().V1().Pods().Informer()

	// list
	lister := factory.Core().V1().Pods().Lister()
	go func() {
		requirement, _ := labels.NewRequirement("app", selection.Equals, []string{"nginx"})
		selector := labels.NewSelector().Add(*requirement)
		for {
			time.Sleep(time.Second * 3)
			pods, err := lister.List(selector)
			if err != nil {
				log.Println("[lister] list error:", err)
				continue
			}
			for _, pod := range pods {
				fmt.Printf("[lister] Pod: %+v\n", *pod)
			}
		}
	}()

	// 添加事件处理程序
	podInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			pod := obj.(*v1.Pod)
			fmt.Printf("Pod Added: %s\n", pod.Name)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oldPod := oldObj.(*v1.Pod)
			newPod := newObj.(*v1.Pod)
			fmt.Printf("Pod Updated: %s -> %s\n", oldPod.Name, newPod.Name)
		},
		DeleteFunc: func(obj interface{}) {
			pod := obj.(*v1.Pod)
			fmt.Printf("Pod Deleted: %s\n", pod.Name)
		},
	})

	// 启动 Informer
	stopCh := make(chan struct{})
	defer close(stopCh)
	factory.Start(stopCh)

	// 等待 Informer 同步
	if !cache.WaitForCacheSync(stopCh, podInformer.HasSynced) {
		fmt.Println("Timeout while waiting for informer to sync")
		return
	}

	// 阻塞主 goroutine
	<-stopCh
}
