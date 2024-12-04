package fixturez

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
	"go.uber.org/mock/gomock"

	"github.com/ibrt/golang-utils/errorz"
)

// BeforeSuite describes a method invoked before starting a test suite.
type BeforeSuite interface {
	BeforeSuite(context.Context, *gomega.WithT) context.Context
}

// AfterSuite represents a method invoked after completing a test suite.
type AfterSuite interface {
	AfterSuite(context.Context, *gomega.WithT)
}

// BeforeTest represents a method invoked before each test method in a suite.
type BeforeTest interface {
	BeforeTest(context.Context, *gomega.WithT, *gomock.Controller) context.Context
}

// AfterTest represents a method invoked after each test method in a suite.
type AfterTest interface {
	AfterTest(context.Context, *gomega.WithT)
}

// RunSuite runs the test suite.
func RunSuite(t *testing.T, suite any) {
	t.Helper()

	rs, err := newRunnableSuite(t, suite)
	if err != nil {
		t.Logf("invalid suite: %v", err.Error())
		t.Fail()
		return
	}

	rs.run()
}

type runnableSuite struct {
	t        *testing.T
	g        *gomega.WithT
	gFmtKeys []format.CustomFormatterKey
	helpers  []reflect.Value
	tests    []int
	sV, sVI  reflect.Value
	sT, sTI  reflect.Type
	ctx      context.Context
}

func newRunnableSuite(t *testing.T, s any) (*runnableSuite, error) {
	t.Helper()

	rs := &runnableSuite{
		t:        t,
		g:        gomega.NewWithT(t),
		gFmtKeys: make([]format.CustomFormatterKey, 0),
		helpers:  make([]reflect.Value, 0),
		tests:    make([]int, 0),
		sV:       reflect.ValueOf(s),
		sVI:      reflect.Indirect(reflect.ValueOf(s)),
		sT:       reflect.TypeOf(s),
		sTI:      reflect.Indirect(reflect.ValueOf(s)).Type(),
		ctx:      context.Background(),
	}

	if rs.sT.Kind() != reflect.Ptr || rs.sT.Elem().Kind() != reflect.Struct {
		return nil, errorz.Errorf("suite must be a struct pointer")
	}

	if err := rs.inspectFields(); err != nil {
		return nil, errorz.Wrap(err)
	}

	if err := rs.inspectMethods(); err != nil {
		return nil, errorz.Wrap(err)
	}

	return rs, nil
}

func (rs *runnableSuite) inspectFields() error {
	rs.t.Helper()

	for i := 0; i < rs.sVI.NumField(); i++ {
		fV := rs.sVI.Field(i)
		f := rs.sTI.Field(i)

		if f.Type.Kind() != reflect.Ptr || f.Type.Elem().Kind() != reflect.Struct {
			return errorz.Errorf("suite field is not helper: %v", f.Name)
		}

		switch fV.Interface().(type) {
		case BeforeSuite, AfterSuite, BeforeTest, AfterTest:
			if fV.IsNil() {
				fV.Set(reflect.New(f.Type.Elem()))
			}
			rs.helpers = append(rs.helpers, fV)
		default:
			return errorz.Errorf("suite field is not helper: %v", f.Name)
		}
	}

	return nil
}

func (rs *runnableSuite) inspectMethods() error {
	rs.t.Helper()

	for i := 0; i < rs.sV.NumMethod(); i++ {
		m := rs.sT.Method(i)

		if !rs.isTestMethod(rs.sV.Method(i), m) {
			return errorz.Errorf("suite method is not test: %v", m.Name)
		}

		rs.tests = append(rs.tests, i)
	}

	return nil
}

func (rs *runnableSuite) isTestMethod(mV reflect.Value, m reflect.Method) bool {
	rs.t.Helper()

	if !strings.HasPrefix(m.Name, "Test") {
		return false
	}

	ctxT := reflect.TypeOf((*context.Context)(nil)).Elem()
	gmgT := reflect.TypeOf((*gomega.WithT)(nil))
	ctrT := reflect.TypeOf((*gomock.Controller)(nil))

	if mV.Type().NumIn() < 1 || mV.Type().NumIn() > 3 {
		return false
	}

	for i := 0; i < mV.Type().NumIn(); i++ {
		switch mV.Type().In(i) {
		case ctxT, gmgT, ctrT:
			// ok
		default:
			return false
		}
	}

	return true
}

func (rs *runnableSuite) beforeSuite() {
	rs.t.Helper()

	fmt.Printf("          %v [BeforeSuite] START\n", rs.t.Name())
	defer fmt.Printf("          %v [BeforeSuite] END\n", rs.t.Name())

	rs.registerCustomFormatters()

	for _, helper := range rs.helpers {
		if beforeSuite, ok := helper.Interface().(BeforeSuite); ok {
			rs.ctx = beforeSuite.BeforeSuite(rs.ctx, rs.g)
		}
	}
}

func (rs *runnableSuite) registerCustomFormatters() {
	rs.t.Helper()

	format.MaxLength = 1000000
	format.MaxDepth = 25
	format.TruncatedDiff = false

	rs.gFmtKeys = []format.CustomFormatterKey{
		format.RegisterCustomFormatter(func(v any) (string, bool) {
			switch err := v.(type) {
			case error:
				return format.Indent + errorz.SDump(err), true
			default:
				return "", false
			}
		}),
	}
}

func (rs *runnableSuite) afterSuite() {
	rs.t.Helper()

	fmt.Printf("          %v [AfterSuite] START\n", rs.t.Name())
	defer fmt.Printf("          %v [AfterSuite] END\n", rs.t.Name())

	for i := len(rs.helpers) - 1; i >= 0; i-- {
		if afterSuite, ok := rs.helpers[i].Interface().(AfterSuite); ok {
			afterSuite.AfterSuite(rs.ctx, rs.g)
		}
	}

	for _, gFmtKey := range rs.gFmtKeys {
		format.UnregisterCustomFormatter(gFmtKey)
	}

	format.MaxLength = 4000
	format.MaxDepth = 10
	format.TruncatedDiff = true
}

func (rs *runnableSuite) beforeTest(tst *testing.T, g *gomega.WithT, ctrl *gomock.Controller) context.Context {
	g.THelper()
	ctx := rs.ctx

	fmt.Printf("          %v [BeforeTest] START\n", tst.Name())
	defer fmt.Printf("          %v [BeforeTest] END\n", tst.Name())

	for _, helper := range rs.helpers {
		if beforeTest, ok := helper.Interface().(BeforeTest); ok {
			ctx = beforeTest.BeforeTest(ctx, g, ctrl)
		}
	}

	return ctx
}

func (rs *runnableSuite) afterTest(ctx context.Context, tst *testing.T, g *gomega.WithT) {
	g.THelper()

	fmt.Printf("          %v [AfterTest] START\n", tst.Name())
	defer fmt.Printf("          %v [AfterTest] END\n", tst.Name())

	for _, helper := range rs.helpers {
		if afterTest, ok := helper.Interface().(AfterTest); ok {
			afterTest.AfterTest(ctx, g)
		}
	}
}

func (rs *runnableSuite) run() {
	rs.t.Helper()

	defer func() {
		rs.t.Helper()
		rs.g.Expect(errorz.MaybeWrapRecover(recover())).To(gomega.Succeed())
	}()

	rs.beforeSuite()
	defer rs.afterSuite()

	for _, i := range rs.tests {
		mT := rs.sT.Method(i)
		mV := rs.sV.Method(i)

		rs.t.Run(mT.Name, func(tst *testing.T) {
			tst.Helper()

			gmg := gomega.NewWithT(tst)
			ctr := gomock.NewController(tst)

			defer func() {
				tst.Helper()
				gmg.Expect(errorz.MaybeWrapRecover(recover())).To(gomega.Succeed())
			}()

			ctx := rs.beforeTest(tst, gmg, ctr)
			defer rs.afterTest(ctx, tst, gmg)

			fmt.Printf("          %v [TestMethod] START\n", tst.Name())
			defer fmt.Printf("          %v [TestMethod] END\n", tst.Name())

			rs.invokeTestMethod(ctx, gmg, ctr, mV)
		})
	}
}

func (rs *runnableSuite) invokeTestMethod(
	ctx context.Context,
	gmg *gomega.WithT,
	ctr *gomock.Controller,
	mV reflect.Value) {

	rs.t.Helper()

	args := make([]reflect.Value, 0)
	ctxT := reflect.TypeOf((*context.Context)(nil)).Elem()
	gmgT := reflect.TypeOf((*gomega.WithT)(nil))
	ctrT := reflect.TypeOf((*gomock.Controller)(nil))

	for i := 0; i < mV.Type().NumIn(); i++ {
		switch mV.Type().In(i) {
		case ctxT:
			args = append(args, reflect.ValueOf(ctx))
		case gmgT:
			args = append(args, reflect.ValueOf(gmg))
		case ctrT:
			args = append(args, reflect.ValueOf(ctr))
		}
	}

	mV.Call(args)
}
