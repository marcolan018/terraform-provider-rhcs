package e2e

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/terraform-redhat/terraform-provider-rhcs/tests/ci"
	exe "github.com/terraform-redhat/terraform-provider-rhcs/tests/utils/exec"
	"github.com/terraform-redhat/terraform-provider-rhcs/tests/utils/helper"
	. "github.com/terraform-redhat/terraform-provider-rhcs/tests/utils/log"
	"github.com/terraform-redhat/terraform-provider-rhcs/tests/utils/profilehandler"
)

var _ = Describe("Break Glass Credential", ci.Day2, ci.FeatureBreakGlassCredential, func() {
	defer GinkgoRecover()
	var (
		profileHandler profilehandler.ProfileHandler
		bgcService     exe.BreakGlassCredentialService
	)

	BeforeEach(func() {
		var err error
		profileHandler, err = profilehandler.NewProfileHandlerFromYamlFile()
		Expect(err).ToNot(HaveOccurred())

		if !profileHandler.Profile().IsHCP() {
			Skip("Test can run only on Hosted cluster")
		}

		if !profileHandler.Profile().IsExternalAuthEnabled() {
			Skip("Test requires external auth enabled profile")
		}

		bgcService, err = profileHandler.Services().GetBreakGlassCredentialService()
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		bgcService.Destroy()
	})

	It("Verify break glass credential can be created - [id:break-glass-credential-01]", ci.Day2, ci.Critical, func() {
		By("Create break glass credential")
		bgcArgs := &exe.BreakGlassCredentialArgs{
			Cluster: helper.StringPointer(clusterID),
		}
		_, err := bgcService.Apply(bgcArgs)
		Expect(err).ToNot(HaveOccurred())
		bgcOutput, err := bgcService.Output()
		Expect(err).ToNot(HaveOccurred())
		Expect(bgcOutput).ToNot(BeEmpty())
		Expect(bgcOutput[0].Kubeconfig).ToNot(BeEmpty())

		Logger.Infof("Successfully verified break glass credential can be created for cluster %s", clusterID)
	})
})
