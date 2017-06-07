package utils

import (
	"bytes"
	"log"
	"os"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Log", func() {
	var (
		b *bytes.Buffer

		initialTimeFormat string
	)

	BeforeEach(func() {
		b = bytes.NewBuffer([]byte{})
		log.SetOutput(b)
		initialTimeFormat = timeFormat
	})

	AfterEach(func() {
		log.SetOutput(os.Stdout)
		SetLogLevel(LogLevelNothing)
		timeFormat = initialTimeFormat
	})

	It("log level nothing", func() {
		SetLogLevel(LogLevelNothing)
		Debugf("debug")
		Infof("info")
		Errorf("err")
		Expect(b.Bytes()).To(Equal([]byte("")))
	})

	It("log level err", func() {
		SetLogLevel(LogLevelError)
		Debugf("debug")
		Infof("info")
		Errorf("err")
		Expect(b.Bytes()).To(ContainSubstring("err\n"))
		Expect(b.Bytes()).ToNot(ContainSubstring("info"))
		Expect(b.Bytes()).ToNot(ContainSubstring("debug"))
	})

	It("log level info", func() {
		SetLogLevel(LogLevelInfo)
		Debugf("debug")
		Infof("info")
		Errorf("err")
		Expect(b.Bytes()).To(ContainSubstring("err\n"))
		Expect(b.Bytes()).To(ContainSubstring("info\n"))
		Expect(b.Bytes()).ToNot(ContainSubstring("debug"))
	})

	It("log level debug", func() {
		SetLogLevel(LogLevelDebug)
		Debugf("debug")
		Infof("info")
		Errorf("err")
		Expect(b.Bytes()).To(ContainSubstring("err\n"))
		Expect(b.Bytes()).To(ContainSubstring("info\n"))
		Expect(b.Bytes()).To(ContainSubstring("debug\n"))
	})

	It("doesn't add a timestamp if the time format is empty", func() {
		SetLogLevel(LogLevelDebug)
		SetLogTimeFormat("")
		Debugf("debug")
		Expect(b.Bytes()).To(Equal([]byte("debug\n")))
	})

	It("adds a timestamp", func() {
		format := "Jan 2, 2006 at 3:04:05pm (MST)"
		SetLogTimeFormat(format)
		SetLogLevel(LogLevelInfo)
		Infof("info")
		t, err := time.Parse(format, string(b.Bytes()[:b.Len()-6]))
		Expect(err).ToNot(HaveOccurred())
		Expect(t).To(BeTemporally("~", time.Now(), 2*time.Second))
	})

	It("says whether debug is enabled", func() {
		Expect(Debug()).To(BeFalse())
		SetLogLevel(LogLevelDebug)
		Expect(Debug()).To(BeTrue())
	})

	It("reads log level from env", func() {
		Expect(logLevel).To(Equal(LogLevelNothing))
		os.Setenv(logEnv, "1")
		readLoggingEnv()
		Expect(logLevel).To(Equal(LogLevelDebug))
	})

	It("does not error reading invalid log levels from env", func() {
		Expect(logLevel).To(Equal(LogLevelNothing))
		os.Setenv(logEnv, "")
		readLoggingEnv()
		Expect(logLevel).To(Equal(LogLevelNothing))
		os.Setenv(logEnv, "asdf")
		readLoggingEnv()
		Expect(logLevel).To(Equal(LogLevelNothing))
	})
})
