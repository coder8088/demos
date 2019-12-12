package main

type File struct {
	fd int
}

func (f *File) Read(offset int64, data []byte) int {
	return 0
}

func (f *File) Close() error {
	return nil
}

// 可以把类型的方法转为函数，不依赖具体对象执行
var CloseFile = (*File).Close

var ReadFile = (*File).Read

func main() {
	f := &File{10}
	ReadFile(f, 0, []byte{})
	CloseFile(f)
}
