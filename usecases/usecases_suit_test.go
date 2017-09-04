package usecases_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/golang/mock/gomock"
	"github.com/onsi/ginkgo/reporters"
	"testing"
)

var mockCtrl *gomock.Controller

var _ = BeforeEach(func() {
	defer GinkgoRecover()
	mockCtrl = gomock.NewController(GinkgoT())
})

var _ = AfterEach(func() {
	mockCtrl.Finish()
})

func TestManagers(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("junit.xml")
	RunSpecsWithDefaultAndCustomReporters(t,
		"Usecases Suite",
		[]Reporter{junitReporter})
}
