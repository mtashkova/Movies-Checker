package env_test

import (
	"errors"
	"os"
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.tools.sap/distribution-store/moovies/pkg/env"
)

var _ = Describe("Env", func() {
	const envName = "ENV_NAME"
	var reader EnvReader

	Describe("ReadString", func() {
		const value = "abcd"
		var env string

		BeforeEach(func() {
			reader = ReadString(envName, &env)
		})

		It("reads env variable when is set", func() {
			Expect(os.Setenv(envName, value)).To(Succeed())
			Expect(reader()).To(Succeed())
			Expect(env).To(Equal(value))
		})

		It("returns an error when env variable is not set", func() {
			Expect(os.Unsetenv(envName)).To(Succeed())
			Expect(reader()).ToNot(Succeed())
		})
	})

	Describe("ReadInt", func() {
		const value = 8
		var env int

		BeforeEach(func() {
			reader = ReadInt(envName, &env)
		})

		It("reads env variable when is set", func() {
			Expect(os.Setenv(envName, strconv.Itoa(value))).To(Succeed())
			Expect(reader()).To(Succeed())
			Expect(env).To(Equal(value))
		})

		It("returns an error when env variable is not set", func() {
			Expect(os.Unsetenv(envName)).To(Succeed())
			Expect(reader()).ToNot(Succeed())
		})

		It("returns an error when env variable is not an integer", func() {
			Expect(os.Setenv(envName, "value")).To(Succeed())
			Expect(reader()).ToNot(Succeed())
		})
	})

	Describe("ReadEnvExec", func() {
		var readerOne, readerTwo EnvReader

		It("executes successfully when all readers succeed", func() {
			const valueOne = 8
			const valueTwo = 9
			var envOne, envTwo int

			readerOne = func() error {
				envOne = valueOne
				return nil
			}
			readerTwo = func() error {
				envTwo = valueTwo
				return nil
			}
			Expect(ReadEnvExec(readerOne, readerTwo)).To(Succeed())
			Expect(envOne).To(Equal(valueOne))
			Expect(envTwo).To(Equal(valueTwo))
		})

		It("returns an error when any reader fails", func() {
			readerOne = func() error {
				return nil
			}

			readerError := errors.New("reading env variable failed")

			readerTwo = func() error {
				return readerError
			}
			err := ReadEnvExec(readerOne, readerTwo)
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(readerError))
		})
	})

})
