package gohumantime_test

import (
	. "github.com/urjitbhatia/gohumantime"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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

			It("missing seconds, return 0", func() {
				Expect(ToMilliseconds("second")).To(Equal(0))
			})
		})

		Context("seconds and hours", func() {
			It("numeric units", func() {
				Expect(ToMilliseconds("1 seconds and 3 hour")).To(Equal(Second + 3*Hour))
			})

			It("word units", func() {
				Expect(ToMilliseconds("three second, two hours")).To(Equal(3*Second + 2*Hour))
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
					output, _ := ToMilliseconds("4 seconds, 2 minutes 1 hour and 3 days and 10weeks, 1 month and 1 year")
					Expect(output).To(Equal(34390924000))
				})
				Ω(runtime.Seconds()).Should(BeNumerically("<", 0.0005), "ToMilliseconds() shouldn't take more than 500 ns.")
			}, 2000)
		})
	})
})
