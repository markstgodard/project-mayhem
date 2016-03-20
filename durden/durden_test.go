package durden_test

import (
	. "github.com/markstgodard/project-mayhem/durden"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Durden", func() {

	Context("NewProjectMayhem", func() {

		It("creates a new project mayhem", func() {
			pm := NewProjectMayhem()
			Expect(pm).ToNot(BeNil())
			Expect(pm.Status).To(Equal(StatusStopped))
		})
	})

})
