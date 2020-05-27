package main

import (
	"bufio"
	"fmt"
	"github.com/yanyiwu/gojieba"
	"io"
	"os"
	"path"
	"path/filepath"
	"time"
)

var books []Book

type Book struct {
	book_name   string
	author_name string
	book_path   string
	file_name   string
	book_type   string
	text        []string
}

//复制文件
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

//创建新文件夹
func MakeDir(path string) {
	//_dir := "./gzFiles2"
	_dir := path
	exist, err := PathExists(_dir)
	if err != nil {
		fmt.Printf("get dir error![%v]\n", err)
		return
	}

	if exist {
		fmt.Printf("has dir![%v]\n", _dir)
	} else {
		fmt.Printf("no dir![%v]\n", _dir)

		// 创建文件夹
		err := os.Mkdir(_dir, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
		} else {
			fmt.Printf("mkdir success!\n")
		}
	}
}

//另存所有书籍
func SaveBooks() {
	MakeDir("./newbooks")

	for _, b := range books {
		//fmt.Println(b.book_path+b.file_name+b.book_type)
		//fmt.Println(b.book_path+b.book_name+b.author_name+b.book_type)

		if b.author_name != "" {

			MakeDir("./newbooks/" + b.author_name)
			old_path := b.book_path + b.file_name + b.book_type
			var new_path string
			if b.book_name != "" {
				new_path = "newbooks/" + b.author_name + "/" + b.book_name + b.book_type
			} else {
				new_path = "newbooks/" + b.author_name + "/" + b.file_name + b.book_type
			}

			CopyFile(new_path, old_path)
		} else {
			MakeDir("./newbooks/未知")
			old_path := b.book_path + b.file_name + b.book_type
			var new_path string
			if b.book_name != "" {
				new_path = "newbooks/未知/" + b.book_name + b.book_type
			} else {
				new_path = "newbooks/未知/" + b.file_name + b.book_type
			}

			CopyFile(new_path, old_path)
		}

	}
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//加载指定目录中的书
func GetFilelist(root_path string) {
	t1 := time.Now()
	err := filepath.Walk(root_path, func(root_path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		paths, fileName := filepath.Split(root_path)

		new_book := Book{}
		new_book.book_path = paths
		new_book.book_type = path.Ext(root_path)

		new_book.file_name = fileName[0 : len(fileName)-len(new_book.book_type)]
		books = append(books, new_book)

		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	elapsed := time.Since(t1)
	fmt.Println("加载全部书籍用时: ", elapsed)
}

//处理书籍信息
func ShowBooks(book []Book) {
	t1 := time.Now()
	var s string
	var words []string

	x := gojieba.NewJieba()
	defer x.Free()

	for _, a := range LoadDictionary() {
		x.AddWord(a)
	}

	for i, b := range book {
		s = b.file_name
		words = x.Cut(s, true)
		book[i].text = words

		//fmt.Println(s+" :")
		//fmt.Println( "    ",strings.Join(words, "/"))
	}
	elapsed := time.Since(t1)
	fmt.Println("处理书籍信息用时: ", elapsed)
}

//加载作家列表,并添加进分词字典
func LoadDictionary() []string {
	t1 := time.Now()
	var list []string
	fi, err := os.Open("dictionary.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		list = append(list, string(a))
	}
	elapsed := time.Since(t1)
	fmt.Println("加载作家列表用时: ", elapsed)
	return list
}
