package bowling_test

import (
	. "bowling"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bowling", func() {
	var game *Game

	BeforeEach(func() {
		game = NewGame()
	})

	It("should have a starting score of 0", func() {
		Expect(game.Score()).To(Equal(0))
	})

	Context("If in 2 tries, the bowler fails to knock down all the pins", func() {
		BeforeEach(func() {
			game.Roll(4)
			game.Roll(4)
		})

		It("their score is the sum of the number of pins they've knocked down in the 2 attempts", func() {
			Expect(game.Score()).To(Equal(8))
		})
	})

	Context("If in 2 tries, the bowler knocks down all the pins, it is a spare", func() {
		BeforeEach(func() {
			game.Roll(4)
			game.Roll(6)
			game.Roll(5)
		})

		It("The scoring of a spare is the sum of the number of pins knocked down plus the number of pins knocked down in the next bowl.", func() {
			Expect(game.Score()).To(Equal(20))
		})
	})

	Context("If in 2 tries, the bowler knocks down all the pins, only one spare", func() {
		BeforeEach(func() {
			game.Roll(4)
			game.Roll(6)
			game.Roll(5)
			game.Roll(5)
			game.Roll(2)
		})

		It("only the next try gets double points.", func() {
			Expect(game.Score()).To(Equal(29))
		})
	})

	Context("If in one try, the bowler knocks down all the pins, it is a strike", func() {
		BeforeEach(func() {
			game.Roll(10)
			game.Roll(5)
			game.Roll(4)
		})

		It("The scoring of a strike is the sum of the number of pins knocked down plus the number of pins knocked down in the next two bowls..", func() {
			Expect(game.Score()).To(Equal(28))
		})
	})

	Context("Number of pins hit per roll", func() {

		It("should not be 0", func() {
			Expect(game.Roll(-4)).To(HaveOccurred())
		})

		It("should not be > 10", func() {
			Expect(game.Roll(11)).To(HaveOccurred())
		})
	})

	Context("There are 10 pins in a try", func() {
		BeforeEach(func() {
			game.Roll(5)
		})

		It("should not be greater than 10 within the same try", func() {
			Expect(game.Roll(6)).To(HaveOccurred())
			game.Roll(5)
			game.Roll(5)
			Expect(game.Roll(6)).To(HaveOccurred())
		})
	})

	Context("There are 10 frames in a match", func() {
		BeforeEach(func() {
			for i := 0; i < 18; i++ {
				game.Roll(3)
			}
		})

		It("should stop scoring after the 10th frame", func() {
			Expect(game.Roll(3)).ToNot(HaveOccurred())
			Expect(game.Roll(3)).ToNot(HaveOccurred())
			Expect(game.Roll(3)).To(HaveOccurred())
		})
	})

	Context("If one has a strike for every roll", func() {
		BeforeEach(func() {
			for i := 0; i < 12; i++ {
				game.Roll(10)
			}
		})

		It("should score them 300", func() {
			Expect(game.Score()).To(Equal(300))
		})
	})

	Context("If one has a strike for every roll except the last", func() {
		BeforeEach(func() {
			for i := 0; i < 10; i++ {
				game.Roll(10)
			}
			game.Roll(5)
			game.Roll(3)
		})

		It("should calculate the score correctly", func() {
			Expect(game.Score()).To(Equal(283))
		})
	})

	Context("In the last frame, if the bowler bowls a spare, they get another bowl", func() {
		BeforeEach(func() {
			for i := 0; i < 18; i++ {
				game.Roll(0)
			}
			game.Roll(7)
			game.Roll(3)
		})

		It("The score of this frame is the sum of the three bowls.", func() {
			Expect(game.Roll(3)).ToNot(HaveOccurred())
			Expect(game.Score()).To(Equal(13))
		})
	})

	Context("In the last frame, if the bowler bowls a strike, they get another 2 bowls", func() {
		BeforeEach(func() {
			for i := 0; i < 18; i++ {
				game.Roll(0)
			}
			game.Roll(10)
		})

		It("The score of this frame is the sum of the three bowls", func() {
			Expect(game.Roll(3)).ToNot(HaveOccurred())
			Expect(game.Roll(3)).ToNot(HaveOccurred())
			Expect(game.Score()).To(Equal(16))
		})

		It("The score of this frame is the sum of the three bowls", func() {
			Expect(game.Roll(10)).ToNot(HaveOccurred())
			Expect(game.Roll(10)).ToNot(HaveOccurred())
			Expect(game.Score()).To(Equal(30))
		})
	})

	Context("If all spares", func() {
		BeforeEach(func() {
			for i := 0; i < 10; i++ {
				game.Roll(3)
				game.Roll(7)
			}
			game.Roll(10)
		})

		It("should have the correct score.", func() {
			Expect(game.Score()).To(Equal(137))
		})
	})

	Context("If no spares or strikes", func() {
		BeforeEach(func() {
			game.Roll(3)
			game.Roll(1)
			game.Roll(2)
			game.Roll(0)
			game.Roll(3)
			game.Roll(5)
			game.Roll(3)
			game.Roll(6)
			game.Roll(4)
			game.Roll(0)
			game.Roll(0)
			game.Roll(5)
			game.Roll(7)
			game.Roll(0)
			game.Roll(8)
			game.Roll(1)
			game.Roll(2)
			game.Roll(0)
		})

		It("have the correct score", func() {
			Expect(game.Score()).To(Equal(50))
		})
	})

})
