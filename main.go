package main

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/yanyiwu/gojieba"
	"io"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

var root_path string="books"
var books[]Book


type Book struct {
	book_name string
	author_name string
	book_path string
	file_name string
	book_type string
	text[] string
}
type entry_bookname_author struct {
	bookname widget.Entry
	author widget.Entry
}
var entrys[10] entry_bookname_author
var editname int

var min int=0
var max int=9
func main(){

	getFilelist(root_path)
	show_books(books)

	fmt.Println(books)
	app := app.New()
	app.Settings().SetTheme(theme.LightTheme())

	w := app.NewWindow("御书")
	w.Resize(fyne.NewSize(1200, 400))
	list := widget.NewVBox()//总体纵向列表

	 listBooks :=widget.NewVBox()


	list3 := widget.NewHBox()//顶层功能按钮横向列表

	list3.Append(widget.NewButton("修改作者", func() {
		fmt.Println("修改作者")
		editname =1
	}))
	list3.Append(widget.NewButton("修改书名", func() {
		fmt.Println("修改书名")
		editname =0
	}))


	page_label:=widget.NewLabel( strconv.Itoa(min)+" -- "+strconv.Itoa(max)+" 共"+strconv.Itoa(len(books)+1))
	list3.Append(page_label)

	list3.Append(widget.NewButton("保存信息", func() {
		fmt.Println("保存信息")
		for i,_:= range books[min:max]{
			index := i // capture
			books[min+index].book_name=entrys[index].bookname.Text
			books[min+index].author_name=entrys[index].author.Text

			 }
	}))

	list3.Append(widget.NewButton("全部信息", func() {
		list := widget.NewVBox()//总体纵向列表
		for i,b:= range books{
			index := i // capture
			page_label:=widget.NewLabel(string(index)+b.book_name)
			list.Append(page_label)

		}
		dialog.ShowCustom("全部信息", "Done", list, w)
	}))

	list3.Append(widget.NewButton("另存到文件夹", func() {
		fmt.Println("另存到文件夹")
		MakeDir( "./newbooks")

		for _,b:=range books{
			//fmt.Println(b.book_path+b.file_name+b.book_type)
			//fmt.Println(b.book_path+b.book_name+b.author_name+b.book_type)

			if b.author_name!="" {

				MakeDir( "./newbooks/"+b.author_name)
				old_path:=b.book_path+b.file_name+b.book_type
				var new_path string
				if b.book_name!="" {
					new_path="newbooks/"+b.author_name+"/"+b.book_name+b.book_type
				}else{
					new_path="newbooks/"+b.author_name+"/"+b.file_name+b.book_type
				}

				CopyFile(new_path,old_path)
			}else{
				MakeDir( "./newbooks/未知")
				old_path:=b.book_path+b.file_name+b.book_type
				var new_path string
				if b.book_name!="" {
					new_path="newbooks/未知/"+b.book_name+b.book_type
				}else{
					new_path="newbooks/未知/"+b.file_name+b.book_type
				}

				CopyFile(new_path,old_path)
			}

		}

	}))




	list4 := widget.NewHBox()//底层功能按钮横向列表
	list4.Append(widget.NewButton("上一页", func() {
		for i,_:= range books[min:max+1]{
			index := i // capture
			books[min+index].book_name=entrys[index].bookname.Text
			books[min+index].author_name=entrys[index].author.Text

		}
		if min<10 {

		}else{
			min-=10
			max-=10
			if max <10{
				max=9
			}
		}
		page_label.SetText( strconv.Itoa(min)+" -- "+strconv.Itoa(max))

		list.Children[1]=switchBooksPages(min,max)

	}))
	list4.Append(widget.NewButton("下一页", func() {
		for i,_:= range books[min:max+1]{
			index := i // capture
			books[min+index].book_name=entrys[index].bookname.Text
			books[min+index].author_name=entrys[index].author.Text

		}
		min+=10
		if  min > len(books){
			min-=10
			max=len(books)-1
		}else {
			max+=10
			if max >=len(books){
				max=len(books)-1
			}
		}

		page_label.SetText( strconv.Itoa(min)+" -- "+strconv.Itoa(max))

		for i,_:= range books[min:max+1]{
			index:=i
			entrys[index].author.SetText(books[min+index].author_name)
			entrys[index].bookname.SetText(books[min+index].book_name)
		}
		list.Children[1]=switchBooksPages(min,max)

	}))



	list.Append(list3)
	list.Append(listBooks)
	list.Append(list4)

	horiz := widget.NewHScrollContainer(list)

	w.SetContent(horiz)

	w.ShowAndRun()
}

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
func MakeDir(path string){
	//_dir := "./gzFiles2"
	_dir :=path
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

func getFilelist(root_path string) {
	err := filepath.Walk(root_path, func(root_path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		paths, fileName := filepath.Split(root_path)

		new_book:=Book{}
		new_book.book_path=paths
		new_book.book_type=path.Ext(root_path)

		new_book.file_name=fileName[0:len(fileName)-len(new_book.book_type)]
		books = append(books, new_book)

		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

func switchBooksPages(min,max int)  *widget.Box{
	listBooks :=  widget.NewVBox()
	for i,b := range books[min:max+1]{
		index := i // capture
		list2 := widget.NewHBox()



		list2.Append(widget.NewButton("全部", func() {

			if editname==0{
				entrys[index].bookname.SetText(books[min+index].file_name)

			} else {
				entrys[index].author.SetText(books[min+index].file_name)
			}
		}))

		list2.Append(widget.NewButton("清空", func() {
			if editname==0{
				entrys[index].bookname.SetText("")

			} else {
				entrys[index].author.SetText("")
			}
		}))

		entrys[index].author.SetPlaceHolder("作者")
		entrys[index].author.SetText(books[min+index].author_name)
		list2.Append(&entrys[index].author)

		entrys[index].bookname.SetPlaceHolder("书名")
		entrys[index].bookname.SetText(books[min+index].book_name)
		list2.Append(&entrys[index].bookname)

		for i,t :=range b.text{
			indext:=i
			button:=widget.NewButton(t, func() {

				if editname==0{
					entrys[index].bookname.SetText(entrys[index].bookname.Text+books[min+index].text[indext])
				} else{
					entrys[index].author.SetText(entrys[index].author.Text+books[min+index].text[indext])
				}

			})
			list2.Append(button)
		}

		listBooks.Append(list2)

	}
	return listBooks
}

func show_books(book[] Book){
	var s string
	var words []string

	x := gojieba.NewJieba()
	defer x.Free()
	x.AddWord("厄普代克")
	x.AddWord("辻村深月")

	for i,b :=range book{
		s=b.file_name
		words = x.Cut(s, true)
		book[i].text=words

		fmt.Println(s+" :")

		fmt.Println( "    ",strings.Join(words, "/"))

	}
}