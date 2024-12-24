package filez_test

import (
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-utils/errorz"
	"github.com/ibrt/golang-utils/filez"
	"github.com/ibrt/golang-utils/filez/internal/itfilez"
	"github.com/ibrt/golang-utils/fixturez"
)

type Suite struct {
	// intentionally empty
}

func TestSuite(t *testing.T) {
	fixturez.RunSuite(t, &Suite{})
}

func (*Suite) TestMustAbs(g *WithT) {
	g.Expect(func() {
		g.Expect(filez.MustAbs("path")).ToNot(Equal("path"))
	}).ToNot(Panic())
}

func (*Suite) TestMustRel(g *WithT) {
	g.Expect(func() {
		g.Expect(filez.MustRel(filepath.Join("a", "b", "c"), filepath.Join("a", "b"))).To(Equal(".."))
		g.Expect(filez.MustRel(filepath.Join("a"), filepath.Join("a", "b"))).To(Equal("b"))
		g.Expect(filez.MustRel("", filepath.Join("a", "b"))).To(Equal(filepath.Join("a", "b")))
		g.Expect(filez.MustRel(filepath.Join("a", "b"), "")).To(Equal(filepath.Join("..", "..")))
	}).ToNot(Panic())
}

func (*Suite) TestGetwdAndChdir(g *WithT) {
	g.Expect(func() {
		wd1 := filez.MustGetwd()
		defer filez.MustChdir(wd1)

		wd2 := filepath.Dir(wd1)
		filez.MustChdir(wd2)
		g.Expect(filez.MustGetwd()).To(Equal(wd2))
	}).ToNot(Panic())
}

func (*Suite) TestMustUserHomeDir(g *WithT) {
	g.Expect(func() {
		g.Expect(filez.MustUserHomeDir()).ToNot(BeEmpty())
	}).ToNot(Panic())
}

func (*Suite) TestMustRemoveAll(g *WithT) {
	g.Expect(
		func() {
			filePath := filez.MustCreateTempFileString("content")
			filez.MustRemoveAll(filePath)
			g.Expect(filez.MustCheckFileExists(filePath)).To(BeFalse())
			filez.MustRemoveAll(filePath)
		}).
		ToNot(Panic())

	g.Expect(func() { filez.MustRemoveAll(string([]byte{0})) }).To(Panic())
}

func (*Suite) TestFilesAndDirs(g *WithT) {
	{
		filePath := filez.MustCreateTempFile([]byte("content"))
		defer func() { errorz.MaybeMustWrap(os.Remove(filePath)) }()

		g.Expect(filez.MustReadFile(filePath)).To(Equal([]byte("content")))
		g.Expect(filez.MustReadFileString(filePath)).To(Equal("content"))
	}

	{
		filePath := filez.MustCreateTempFileString("content")
		defer func() { errorz.MaybeMustWrap(os.Remove(filePath)) }()

		g.Expect(filez.MustReadFile(filePath)).To(Equal([]byte("content")))
		g.Expect(filez.MustReadFileString(filePath)).To(Equal("content"))
	}

	{
		dirPath := filez.MustCreateTempDir()
		defer filez.MustRemoveAll(dirPath)

		filePath := filez.MustWriteFile(filepath.Join(dirPath, "first"), 0777, 0666, []byte("content"))
		g.Expect(filez.MustReadFile(filePath)).To(Equal([]byte("content")))
		g.Expect(filez.MustReadFileString(filePath)).To(Equal("content"))

		filePath = filez.MustWriteFileString(filepath.Join(dirPath, "second"), 0777, 0666, "content")
		g.Expect(filez.MustReadFile(filePath)).To(Equal([]byte("content")))
		g.Expect(filez.MustReadFileString(filePath)).To(Equal("content"))

		g.Expect(filez.MustCheckPathExists(dirPath)).To(BeTrue())
		g.Expect(filez.MustCheckPathExists(filepath.Join(dirPath, "third"))).To(BeFalse())
		g.Expect(func() { filez.MustCheckPathExists(string([]byte{0})) }).To(Panic())

		g.Expect(filez.MustCheckFileExists(filePath)).To(BeTrue())
		g.Expect(filez.MustCheckFileExists(filepath.Join(dirPath, "third"))).To(BeFalse())
		g.Expect(filez.MustCheckFileExists(dirPath)).To(BeFalse())
		g.Expect(func() { filez.MustCheckFileExists(string([]byte{0})) }).To(Panic())

		filez.MustPrepareDir(dirPath, 0777)
		g.Expect(filez.MustCheckFileExists(filePath)).To(BeFalse())
	}
}

func (*Suite) TestMustIsChild(g *WithT) {
	g.Expect(filez.MustIsChild("a", "a")).To(BeTrue())
	g.Expect(filez.MustIsChild("", "a")).To(BeTrue())
	g.Expect(filez.MustIsChild("a", filepath.Join("a", "b"))).To(BeTrue())
	g.Expect(filez.MustIsChild("a", "")).To(BeFalse())
	g.Expect(filez.MustIsChild(filepath.Join("a", "b"), "a")).To(BeFalse())
	g.Expect(filez.MustIsChild("a", "b")).To(BeFalse())
}

func (*Suite) TestMustRelForDisplay(g *WithT) {
	g.Expect(filez.MustRelForDisplay("a")).To(Equal("a"))
	g.Expect(filez.MustRelForDisplay(filez.MustAbs("a"))).To(Equal("a"))
	g.Expect(filez.MustRelForDisplay(filez.MustAbs(filepath.Join("..", "a")))).To(Equal(filez.MustAbs(filepath.Join("..", "a"))))
}

func (*Suite) TestMustExport_Level0(g *WithT) {
	dirPath := filez.MustCreateTempDir()
	defer filez.MustRemoveAll(dirPath)

	filez.MustExport(itfilez.AssetsEmbed, ".", dirPath)

	g.Expect(filez.MustListRegularFilePaths(dirPath)).To(Equal([]string{
		"assets/child/second.txt",
		"assets/first.txt",
	}))
}

func (*Suite) TestMustExport_Level1(g *WithT) {
	dirPath := filez.MustCreateTempDir()
	defer filez.MustRemoveAll(dirPath)

	filez.MustExport(itfilez.AssetsEmbed, "assets", dirPath)

	g.Expect(filez.MustListRegularFilePaths(dirPath)).To(Equal([]string{
		"child/second.txt",
		"first.txt",
	}))
}

func (*Suite) TestMustExport_Level2(g *WithT) {
	dirPath := filez.MustCreateTempDir()
	defer filez.MustRemoveAll(dirPath)

	filez.MustExport(itfilez.AssetsEmbed, "assets/child", dirPath)
	g.Expect(filez.MustListRegularFilePaths(dirPath)).To(Equal([]string{
		"second.txt",
	}))
}
