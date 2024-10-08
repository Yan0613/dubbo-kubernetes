package k8s_test

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"

	bootstrap_k8s "github.com/apache/dubbo-kubernetes/pkg/plugins/bootstrap/k8s"
	"github.com/apache/dubbo-kubernetes/pkg/test"
)

var (
	k8sClient       client.Client
	testEnv         *envtest.Environment
	k8sClientScheme *runtime.Scheme
)

func TestKubernetes(t *testing.T) {
	test.RunSpecs(t, "Kubernetes Secrets Suite")
}

var _ = BeforeSuite(test.Within(time.Minute, func() {
	By("bootstrapping test environment")
	testEnv = &envtest.Environment{}

	cfg, err := testEnv.Start()
	Expect(err).ToNot(HaveOccurred())
	Expect(cfg).ToNot(BeNil())

	k8sClientScheme, err = bootstrap_k8s.NewScheme()
	Expect(err).ToNot(HaveOccurred())

	// +kubebuilder:scaffold:scheme

	k8sClient, err = client.New(cfg, client.Options{Scheme: k8sClientScheme})
	Expect(err).ToNot(HaveOccurred())
	Expect(k8sClient).ToNot(BeNil())
}))

var _ = AfterSuite(func() {
	By("tearing down the test environment")
	err := testEnv.Stop()
	Expect(err).ToNot(HaveOccurred())
})
