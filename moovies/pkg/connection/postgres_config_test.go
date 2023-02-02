package connection_test

import (
	"os"
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.tools.sap/distribution-store/moovies/pkg/connection"
)

var _ = Describe("PostgresConfig", func() {
	Describe("LoadConfig", func() {
		const (
			host     = "localhost"
			port     = 5423
			password = "secret"
			user     = "db"
			dbname   = "db"

			hostEnv     = "HOST"
			portEnv     = "PORT"
			passwordEnv = "PASSWORD"
			userEnv     = "USER"
			dbnameEnv   = "DB_NAME"
		)

		When("all env variables exist", func() {
			BeforeEach(func() {
				Expect(os.Setenv(hostEnv, host)).To(Succeed())
				Expect(os.Setenv(portEnv, strconv.Itoa(port))).To(Succeed())
				Expect(os.Setenv(passwordEnv, password)).To(Succeed())
				Expect(os.Setenv(userEnv, user)).To(Succeed())
				Expect(os.Setenv(dbnameEnv, dbname)).To(Succeed())
			})

			AfterEach(func() {
				Expect(os.Unsetenv(hostEnv)).To(Succeed())
				Expect(os.Unsetenv(portEnv)).To(Succeed())
				Expect(os.Unsetenv(passwordEnv)).To(Succeed())
				Expect(os.Unsetenv(userEnv)).To(Succeed())
				Expect(os.Unsetenv(dbnameEnv)).To(Succeed())
			})

			It("reads successfully", func() {
				config, err := LoadConfig()
				Expect(err).ToNot(HaveOccurred())
				Expect(config.Host).To(Equal(host))
				Expect(config.Port).To(Equal(port))
				Expect(config.Password).To(Equal(password))
				Expect(config.User).To(Equal(user))
				Expect(config.DBName).To(Equal(dbname))
			})

			When("env variable is missing", func() {
				BeforeEach(func() {
					Expect(os.Unsetenv(hostEnv)).To(Succeed())
				})

				It("returns an error", func() {
					_, err := LoadConfig()
					Expect(err).To(HaveOccurred())
				})
			})
		})
	})
})
