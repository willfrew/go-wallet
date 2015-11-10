package wallet_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/willfrew/wallet-test"
	"testing"
)

func TestWallet(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Wallet Suite")
}

var _ = Describe("Wallet", func() {

	Describe("Withdraw", func() {

		var w wallet.Wallet

		BeforeEach(func() {
			w = *wallet.NewWallet(100)
		})

		It("should allow valid withdrawals", func() {
			err, cash := w.Withdraw(50)
			Expect(err).NotTo(HaveOccurred())
			Expect(cash).To(Equal(wallet.Cash(50)))
		})

		It("should take the withdrawn amount from the balance", func() {
			w.Withdraw(30)
			Expect(w.Balance()).To(Equal(wallet.Cash(70)))

			w.Withdraw(69)
			Expect(w.Balance()).To(Equal(wallet.Cash(1)))

			w.Withdraw(1)
			Expect(w.Balance()).To(Equal(wallet.Cash(0)))
		})

		It("should refuse unfulfillable transactions", func() {
			err, cash := w.Withdraw(150)
			Expect(err).To(HaveOccurred())
			Expect(cash).To(Equal(wallet.Cash(0)))
		})

		It("should refuse negative withdrawal amounts", func() {
			err, _ := w.Withdraw(-10)
			Expect(err).To(HaveOccurred())
		})

		It("should refuse zero withdrawals", func() {
			err, _ := w.Withdraw(0)
			Expect(err).To(HaveOccurred())
		})

	})

	Describe("Deposit", func() {

		var w wallet.Wallet

		BeforeEach(func() {
			w = *wallet.NewWallet(0)
		})

		It("should refuse negative deposits", func() {
			err := w.Deposit(-1)
			Expect(err).To(HaveOccurred())
		})

		It("should refuse zero deposits", func() {
			err := w.Deposit(0)
			Expect(err).To(HaveOccurred())
		})

		It("should add deposits to the balance", func() {
			w.Deposit(500)
			Expect(w.Balance()).To(Equal(wallet.Cash(500)))
		})

		It("should support multiple deposits", func() {
			w.Deposit(10)
			w.Deposit(40)
			w.Deposit(50)
			Expect(w.Balance()).To(Equal(wallet.Cash(100)))
		})

	})

	Describe("Transfer", func() {

		var (
			w1 wallet.Wallet
			w2 wallet.Wallet
		)

		BeforeEach(func() {
			w1 = *wallet.NewWallet(100)
			w2 = *wallet.NewWallet(1000)
		})

		It("should facilitate balance transfers", func() {
			err := w2.Transfer(300, &w1)
			Expect(err).NotTo(HaveOccurred())
			Expect(w1.Balance()).To(Equal(wallet.Cash(400)))
			Expect(w2.Balance()).To(Equal(wallet.Cash(700)))
		})

	})

})
