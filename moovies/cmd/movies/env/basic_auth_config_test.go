package env_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.tools.sap/distribution-store/moovies/cmd/movies/env"
)

var _ = Describe("BasicAuthConfig", func() {
	Describe("BasicAuthConfiguration", func() {
		const (
			username = "user1"
			password = "pass1"

			usernameEnv = "APP_USERNAME"
			passwordEnv = "APP_PASSWORD"
		)

		When("all env variables are set", func() {
			BeforeEach(func() {
				Expect(os.Setenv(usernameEnv, username)).To(Succeed())
				Expect(os.Setenv(passwordEnv, password)).To(Succeed())
			})

			AfterEach(func() {
				Expect(os.Unsetenv(usernameEnv)).To(Succeed())
				Expect(os.Unsetenv(passwordEnv)).To(Succeed())
			})

			It("reads successfully", func() {
				basicAuthConfig, err := env.BasicAuthConfiguration()
				Expect(err).ToNot(HaveOccurred())
				Expect(basicAuthConfig.Username).To(Equal(username))
				Expect(basicAuthConfig.Password).To(Equal(password))
			})

			When("env variable is missing", func() {
				BeforeEach(func() {
					Expect(os.Unsetenv(usernameEnv)).To(Succeed())
				})

				It("returns an error", func() {
					_, err := env.BasicAuthConfiguration()
					Expect(err).To(HaveOccurred())
				})
			})
		})
	})
})
