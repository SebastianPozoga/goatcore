package memfs_test

import (
	"testing"

	"github.com/goatcms/goat-core/filesystem/filespace/memfs"
	"github.com/goatcms/goat-core/testbase"
)

func TestMkdir(t *testing.T) {
	// init
	fs, err := memfs.NewFilespace()
	if err != nil {
		t.Error(err)
	}
	// create directories
	path := "/mydir1/mydir2/mydir3"
	if err := fs.MkdirAll(path, 0777); err != nil {
		t.Error("Fail when create directories", err)
	}
	// test node type
	if !fs.IsDir("/mydir1/mydir2") {
		t.Error("node is not a directory or not exists")
	}
	if !fs.IsDir(path) {
		t.Error("node is not a directory or not exists")
	}
	if fs.IsDir("/noExistPath") {
		t.Error("node is not a directory or not exists")
	}
}

func TestRemove(t *testing.T) {
	// init
	fs, err := memfs.NewFilespace()
	if err != nil {
		t.Error(err)
	}
	// create directories
	path := "/mydir1/mydir2/mydir3/mydir4"
	if err := fs.MkdirAll(path, 0777); err != nil {
		t.Error("Fail when create directories", err)
		return
	}
	// test node type
	if fs.Remove("/mydir1/mydir2") == nil {
		t.Error("Remove remove no empty directory is not allowed")
	}
	if err := fs.Remove("/mydir1/mydir2/mydir3/mydir4"); err != nil {
		t.Errorf("Remove should remove empty directory (Error: %v)", err)
	}
	if err := fs.RemoveAll("/mydir1"); err != nil {
		t.Errorf("RemoveAll should remove empty directory (Error: %v )", err)
	}
}

func TestWriteAndRead(t *testing.T) {
	// init
	fs, err := memfs.NewFilespace()
	if err != nil {
		t.Error(err)
	}
	//Create data
	testData := []byte("There is test data")

	// create directories
	path := "/mydir1/mydir2/mydir3/myfile.ex"
	fs.WriteFile(path, testData, 0777)
	readData, err := fs.ReadFile(path)
	if err != nil {
		t.Error("can not read file after write data ", err)
	}
	if !testbase.ByteArrayEq(readData, testData) {
		t.Error("read data are diffrent ", readData, testData)
	}
}

func TestCopy(t *testing.T) {
	const (
		srcPath   = "src"
		destPath  = "dest"
		file1Path = "/d1/d2/f1.ex"
		file2Path = "/d1/z1/f2.exx"
	)

	// init
	fs, err := memfs.NewFilespace()
	if err != nil {
		t.Error(err)
	}
	//Create data
	testData1 := []byte("Content of file 1")
	testData2 := []byte("Content of file 2")

	// create test model
	fs.WriteFile(srcPath+file1Path, testData1, 0777)
	fs.WriteFile(srcPath+file2Path, testData2, 0777)

	// copy
	fs.Copy(srcPath, destPath)

	// test
	readData1, err := fs.ReadFile(destPath + file1Path)
	if err != nil {
		t.Error("can not read file1 after write data ", err)
	} else {
		if !testbase.ByteArrayEq(testData1, readData1) {
			t.Error("read1 and test1 data are diffrent ", testData1, readData1)
		}
	}
	readData2, err := fs.ReadFile(destPath + file2Path)
	if err != nil {
		t.Error("can not read file2 after write data ", err)
	} else {
		if !testbase.ByteArrayEq(testData2, readData2) {
			t.Error("read2 and test2 data are diffrent ", testData2, readData2)
		}
	}
}

func TestCopySingleFile(t *testing.T) {
	const (
		srcPath   = "src"
		destPath  = "dest"
		file1Path = "/d1/d2/f1.ex"
	)
	// init
	fs, err := memfs.NewFilespace()
	if err != nil {
		t.Error(err)
	}
	//Create data
	testData1 := []byte("Content of file 1")
	// create test model
	fs.WriteFile(srcPath+file1Path, testData1, 0777)
	// copy
	fs.Copy(srcPath+file1Path, destPath+file1Path)
	// test
	readData1, err := fs.ReadFile(destPath + file1Path)
	if err != nil {
		t.Error("can not read file1 after write data ", err)
	} else {
		if !testbase.ByteArrayEq(testData1, readData1) {
			t.Error("read1 and test1 data are diffrent ", testData1, readData1)
		}
	}
}

func TestWriteStreamAndRead(t *testing.T) {
	// init
	fs, err := memfs.NewFilespace()
	if err != nil {
		t.Error(err)
	}
	//Create data
	testData := []byte("There is test data")
	// create directories
	path := "/mydir1/mydir2/mydir3/myfile.ex"
	writer, err := fs.Writer(path)
	if err != nil {
		t.Error(err)
		return
	}
	n, err := writer.Write(testData)
	if err != nil {
		t.Error(err)
		return
	}
	if n != len(testData) {
		t.Errorf("return length should be equal to data size %v %v", n, len(testData))
		return
	}
	err = writer.Close()
	if err != nil {
		t.Error(err)
		return
	}
	readData, err := fs.ReadFile(path)
	if err != nil {
		t.Error(err)
	}
	if !testbase.ByteArrayEq(readData, testData) {
		t.Error("read data are diffrent ", readData, testData)
	}
}

func TestWriteAndReader(t *testing.T) {
	// init
	fs, err := memfs.NewFilespace()
	if err != nil {
		t.Error(err)
	}
	//Create data
	testData := []byte("There is test data")

	// create directories
	path := "/mydir1/mydir2/mydir3/myfile.ex"
	err = fs.WriteFile(path, testData, 0777)
	if err != nil {
		t.Error(err)
		return
	}
	reader, err := fs.Reader(path)
	if err != nil {
		t.Error(err)
		return
	}
	buf := make([]byte, 222)
	n, err := reader.Read(buf)
	if err != nil {
		t.Error(err)
		return
	}
	err = reader.Close()
	if err != nil {
		t.Error(err)
		return
	}
	if n != len(testData) {
		t.Errorf("return length should be equal to data size %v %v", n, len(testData))
		return
	}
	if !testbase.ByteArrayEq(buf[:n], testData) {
		t.Error("read data are diffrent ", buf, testData)
	}
}

func TestReadDir(t *testing.T) {
	var dir1 bool
	var dir2 bool
	// init
	fs, err := memfs.NewFilespace()
	if err != nil {
		t.Error(err)
	}
	// prepare data
	if err := fs.MkdirAll("dir1", 0777); err != nil {
		t.Error(err)
		return
	}
	if err := fs.MkdirAll("dir2", 0777); err != nil {
		t.Error(err)
		return
	}
	//testing
	list, err := fs.ReadDir("./")
	if err != nil {
		t.Error(err)
		return
	}
	for _, file := range list {
		switch file.Name() {
		case "dir1":
			dir1 = true
		case "dir2":
			dir2 = true
		default:
			t.Errorf("unknown file %s", file.Name())
			return
		}
	}
	if !dir1 {
		t.Errorf("don't read dir1")
	}
	if !dir2 {
		t.Errorf("don't read dir2")
	}
}
