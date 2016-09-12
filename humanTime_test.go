package gohumantime_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/urjitbhatia/gohumantime"
)

var _ = Describe("HumanTime", func() {
	Describe("Valid humanTime conversions", func() {
		Context("seconds only", func() {
			It("numeric seconds", func() {
				Expect(ToMilliseconds("1 second")).To(Equal(Second))
			})

			It("word seconds", func() {
				Expect(ToMilliseconds("three second")).To(Equal(3 * Second))
			})

			It("word *** minutes", func() {
				Expect(ToMilliseconds("10 minutes")).To(Equal(10 * Minute))
			})

			It("missing seconds, return 0", func() {
				Expect(ToMilliseconds("second")).To(Equal(0))
			})

			It("fifty seconds", func() {
				Expect(ToMilliseconds("fifty seconds")).To(Equal(50 * Second))
			})
		})

		Context("minutes", func() {
			It("seventyfour minutes", func() {
				Expect(ToMilliseconds("74 minutes")).To(Equal(74 * Minute))
			})
			It("60 minutes", func() {
				Expect(ToMilliseconds("60 minutes")).To(Equal(60 * Minute))
			})
		})

		Context("mixed units", func() {
			It("numeric units", func() {
				Expect(ToMilliseconds("10 minutes")).To(Equal(Minute * 10))
			})
		})

		Context("now", func() {
			It("now", func() {
				Expect(ToMilliseconds("now")).To(Equal(0))
			})
		})

		Context("seconds and hours", func() {
			It("numeric units", func() {
				Expect(ToMilliseconds("1 seconds and 3 hour")).To(Equal(Second + 3*Hour))
			})

			It("mixed units", func() {
				Expect(ToMilliseconds("1 hour 14 minutes 0 seconds")).To(Equal(1*Hour + 14*Minute))
			})

			It("numeric fractional units", func() {
				Expect(ToMilliseconds("1.3 seconds and 3 hour")).To(Equal(int(1.3*Second) + 3*Hour))
			})

			It("word units", func() {
				Expect(ToMilliseconds("three second, two hours")).To(Equal(3*Second + 2*Hour))
			})

			It("word units", func() {
				Expect(ToMilliseconds("ninety second, twenty hours")).To(Equal(90*Second + 20*Hour))
			})

			It("missing seconds, with hour, ignore missing seconds", func() {
				Expect(ToMilliseconds("second 2 hour")).To(Equal(2 * Hour))
			})
		})

		Context("days and weeks", func() {
			It("numeric units", func() {
				Expect(ToMilliseconds("1 days and 3 weeks")).To(Equal(Day + 3*Week))
			})

			It("word units", func() {
				Expect(ToMilliseconds("three days and one week")).To(Equal(3*Day + Week))
			})

			It("mixed numeric and word units", func() {
				Expect(ToMilliseconds("4 days and one week")).To(Equal(4*Day + Week))
			})
		})

		Context("invalid input, return 0", func() {
			It("empty input", func() {
				Expect(ToMilliseconds("")).To(Equal(0))
			})

			It("empty input with spaces", func() {
				Expect(ToMilliseconds("   ")).To(Equal(0))
			})

			It("empty input with spaces and useless characters", func() {
				Expect(ToMilliseconds("  foobar[] ")).To(Equal(0))
			})
		})

		Context("benchMark", func() {
			Measure("convert humanreadable strings to time", func(b Benchmarker) {
				runtime := b.Time("runtime", func() {
					output, _ := ToMilliseconds("4 seconds, 2.4 minutes 1 hour and 3.33 days and 10weeks, 1 month and 1 year")
					Expect(output).To(Equal(34419460000))
				})

				Ω(runtime.Seconds()).Should(BeNumerically("<", 0.006), "ToMilliseconds() shouldn't take more than 600 ns.")
			}, 3000)
		})
	})
})
