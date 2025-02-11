package fixturez_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
	"go.uber.org/mock/gomock"

	"github.com/ibrt/golang-utils/errorz"
	"github.com/ibrt/golang-utils/fixturez"
)

type contextKey int

const (
	beforeSuiteContextKey contextKey = iota
	beforeTestContextKey  contextKey = iota
)

// Helper implements a test helper.
type Helper struct {
	beforeSuite, beforeTest, afterTest, afterSuite int
}

// BeforeSuite implements the fixturez.BeforeSuite interface.
func (m *Helper) BeforeSuite(ctx context.Context, g *WithT) context.Context {
	g.Expect(m).ToNot(BeNil())
	g.Expect(ctx).ToNot(BeNil())
	g.Expect(ctx.Value(beforeSuiteContextKey)).To(BeNil())
	g.Expect(ctx.Value(beforeTestContextKey)).To(BeNil())

	m.beforeSuite++
	return context.WithValue(ctx, beforeSuiteContextKey, true)
}

// BeforeTest implements the fixturez.BeforeTest interface.
func (m *Helper) BeforeTest(ctx context.Context, g *WithT, _ *gomock.Controller) context.Context {
	g.Expect(m).ToNot(BeNil())
	g.Expect(ctx).ToNot(BeNil())
	g.Expect(ctx.Value(beforeSuiteContextKey)).To(BeTrue())
	g.Expect(ctx.Value(beforeTestContextKey)).To(BeNil())

	m.beforeTest++
	return context.WithValue(ctx, beforeTestContextKey, true)
}

// AfterTest implements the fixturez.AfterTest interface.
func (m *Helper) AfterTest(ctx context.Context, g *WithT) {
	g.Expect(m).ToNot(BeNil())
	g.Expect(ctx).ToNot(BeNil())
	g.Expect(ctx.Value(beforeSuiteContextKey)).To(BeTrue())
	g.Expect(ctx.Value(beforeTestContextKey)).To(BeTrue())

	m.afterTest++
}

func (m *Helper) AfterSuite(ctx context.Context, g *WithT) {
	g.Expect(m).ToNot(BeNil())
	g.Expect(ctx).ToNot(BeNil())
	g.Expect(ctx.Value(beforeSuiteContextKey)).To(BeTrue())
	g.Expect(ctx.Value(beforeTestContextKey)).To(BeNil())

	m.afterSuite++
}

// SuiteCorrect implements a test suite.
type SuiteCorrect struct {
	Helper *Helper
}

func (*SuiteCorrect) TestFirst(ctx context.Context, g *WithT) {
	g.Expect(ctx).ToNot(BeNil())
}

func (*SuiteCorrect) TestSecond(ctx context.Context, g *WithT, ctrl *gomock.Controller) {
	g.Expect(ctx).ToNot(BeNil())
	g.Expect(ctrl).ToNot(BeNil())
}

func (*SuiteCorrect) TestErrorFormatter(g *WithT) {
	g.Expect(format.Object(errorz.Errorf("test error"), 0)).
		To(ContainSubstring("\"message\": \"test error\""))

	g.Expect(format.Object("test string", 0)).
		To(Equal("<string>: \"test string\""))
}

func TestSuite_Correct(t *testing.T) {
	s := &SuiteCorrect{}
	fixturez.RunSuite(t, s)

	g := NewWithT(t)
	g.Expect(s.Helper.beforeSuite).To(Equal(1))
	g.Expect(s.Helper.beforeTest).To(Equal(3))
	g.Expect(s.Helper.afterTest).To(Equal(3))
	g.Expect(s.Helper.afterSuite).To(Equal(1))
}

func TestSuite_CorrectButNotPointer(t *testing.T) {
	g := NewWithT(t)

	tt := &testing.T{}
	fixturez.RunSuite(tt, SuiteCorrect{})
	g.Expect(tt.Failed()).To(BeTrue())
}

type SuiteIncorrectFieldString struct {
	Field string
}

func TestSuite_IncorrectFieldString(t *testing.T) {
	g := NewWithT(t)

	tt := &testing.T{}
	fixturez.RunSuite(tt, &SuiteIncorrectFieldString{})
	g.Expect(tt.Failed()).To(BeTrue())
}

type SuiteIncorrectFieldStruct struct {
	Field *struct{}
}

func TestSuite_IncorrectFieldStruct(t *testing.T) {
	g := NewWithT(t)

	tt := &testing.T{}
	fixturez.RunSuite(tt, &SuiteIncorrectFieldStruct{})
	g.Expect(tt.Failed()).To(BeTrue())
}

type SuiteIncorrectMethodName struct {
	// intentionally empty
}

func (*SuiteIncorrectMethodName) Method(_ context.Context, _ *WithT) {
	// intentionally empty
}

func TestSuite_IncorrectMethodName(t *testing.T) {
	g := NewWithT(t)

	tt := &testing.T{}
	fixturez.RunSuite(tt, &SuiteIncorrectMethodName{})
	g.Expect(tt.Failed()).To(BeTrue())
}

type SuiteIncorrectMethodSignature1 struct {
	// intentionally empty
}

func (*SuiteIncorrectMethodSignature1) TestMethod(_ string) {
	// intentionally empty
}

func TestSuite_IncorrectMethodSignature1(t *testing.T) {
	g := NewWithT(t)

	tt := &testing.T{}
	fixturez.RunSuite(tt, &SuiteIncorrectMethodSignature1{})
	g.Expect(tt.Failed()).To(BeTrue())
}

type SuiteIncorrectMethodSignature2 struct {
	// intentionally empty
}

func (*SuiteIncorrectMethodSignature2) TestMethod() {
	// intentionally empty
}

func TestSuite_IncorrectMethodSignature2(t *testing.T) {
	g := NewWithT(t)

	tt := &testing.T{}
	fixturez.RunSuite(tt, &SuiteIncorrectMethodSignature2{})
	g.Expect(tt.Failed()).To(BeTrue())
}
