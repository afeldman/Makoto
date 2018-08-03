package spec_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/afeldman/Makoto/kpc"
)

var _ = Describe("Authors", func() {
	var (
		aut1 Author
		aut2 Author
	)

	BeforeEach(func() {
		aut1 = *Author_Init("Anton Feldmann")
		aut2 = *Author_Init("Dr. Feldmann")
	})

	Context("Emailaddres empty", func() {
		It("Author 1 has empty email", func(){
			Expect(*(aut1.GetEmail())).Should(Equal(""))
		})
		It("Author 2 has empty email", func(){
			Expect(*(aut2.GetEmail())).Should(Equal(""))
		})
		It("Author 1 Name is Anton Feldmann", func(){
			Expect(*(aut1.GetName())).Should(Equal("Anton Feldmann"))
		})
		It("Author 2 Name is Dr. Feldmann", func(){
			Expect(*(aut2.GetName())).Should(Equal("Dr. Feldmann"))
		})
	})

	Context("change of email", func(){
		It("Author 1 has an non empty email", func(){
			aut1.SetEmail("anton.feldmann@gmail.com")
			Expect(*(aut1.GetEmail())).Should(Equal("anton.feldmann@gmail.com"))
		})
		It("Author 2 has an non empty email", func(){
			aut2.SetEmail("anton.feldmann@gmail.com")
			Expect(*(aut2.GetEmail())).Should(Equal("anton.feldmann@gmail.com"))
		})
		It("Author 1 change name", func(){
			aut1.SetName("foo bar")
			Expect(*(aut1.GetName())).Should(Equal("foo bar"))
		})
		It("Author 2 change name", func(){
			aut2.SetName("bar foo")
			Expect(*(aut2.GetName())).Should(Equal("bar foo"))
		})
	})
})
