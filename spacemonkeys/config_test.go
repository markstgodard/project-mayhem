package spacemonkeys_test

import (
	"os"

	. "github.com/markstgodard/project-mayhem/spacemonkeys"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {

	var (
		cfg *Config
	)

	Context("Load", func() {

		BeforeEach(func() {
			os.Setenv("BOSH_DIRECTOR_HOST", "192.168.50.4:25555")
			os.Setenv("BOSH_DIRECTOR_USERNAME", "tyler")
			os.Setenv("BOSH_DIRECTOR_PASSWORD", "durden")
		})

		AfterEach(func() {
			os.Unsetenv("BOSH_DIRECTOR_HOST")
			os.Unsetenv("BOSH_DIRECTOR_USERNAME")
			os.Unsetenv("BOSH_DIRECTOR_PASSWORD")
		})

		It("loads config from env", func() {
			cfg = LoadConfig()
			Expect(cfg.Deployment.Host).To(Equal("192.168.50.4:25555"))
			Expect(cfg.Deployment.Username).To(Equal("tyler"))
			Expect(cfg.Deployment.Password).To(Equal("durden"))
		})
	})

})
