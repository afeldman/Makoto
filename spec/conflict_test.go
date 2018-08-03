package spec_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/afeldman/Makoto/kpc"
)

var _ = Describe("Conflicts", func() {
	var (
		con Conflict
	)

	BeforeEach(func() {
		con = *Conflict_Init("project name","0.1.1")
	})

	Context("add version numbers to conflict ", func() {
		It("add one version to conflict", func(){
			con.AddElement("0.12.1");
			Expect(con.GetLength()).Should(Equal(2))
		})
		It("change version number", func(){
			con.AddElement("0.12.1");
			change := con.ChangeVersion("0.12.1","0.1.3")
			Expect(change).Should(Equal(true))

			Expect(con.ContainsVersion("0.1.3")).Should(Equal(true))
			Expect(con.ContainsVersion("0.12.1")).Should(Equal(false))
		})
		It("check name", func(){
			Expect(*(con.GetName())).Should(Equal("project name"))
		})
		It("change name", func(){
			con.SetName("abcd")
			Expect(*(con.GetName())).Should(Equal("abcd"))
		})
		It("delete version", func(){
			con.AddElement("0.12.1");
			Expect(con.GetLength()).Should(Equal(2))
			con.RejectVersion("0.12.1");
			Expect(con.ContainsVersion("0.12.1")).Should(Equal(false))
		})
	})
})
