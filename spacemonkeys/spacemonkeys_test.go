package spacemonkeys_test

import (
	. "github.com/markstgodard/project-mayhem/spacemonkeys"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SpaceMonkeys", func() {

	var (
		cfg Config
	)

	Context("NewProjectMayhem", func() {

		BeforeEach(func() {
			cfg = Config{
				Deployment: BoshDeployment{
					Host:     "https://192.168.50.4:25555",
					Username: "admin",
					Password: "admin",
				},
			}
		})

		It("creates a new project mayhem from config", func() {
			pm := NewProjectMayhem()
			Expect(pm).ToNot(BeNil())
			Expect(pm.Status).To(Equal(StatusStopped))
		})
	})

})
