package main

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/yanyiwu/gojieba"
	"os"
	"path"
	"path/filepath"
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
var entrys[20] entry_bookname_author

func main(){

	getFilelist(root_path)
	show_books(books)

	fmt.Println(books)
	app := app.New()
	app.Settings().SetTheme(theme.LightTheme())

	w := app.NewWindow("Hello")
	w.Resize(fyne.NewSize(1200, 400))
	list := widget.NewVBox()//总体纵向列表

	var editname int


	list3 := widget.NewHBox()//顶层功能按钮横向列表
	list3.Append(widget.NewButton("修改作者", func() {
		fmt.Println("修改作者")
		editname =1
	}))
	list3.Append(widget.NewButton("修改书名", func() {
		fmt.Println("修改书名")
		editname =0
	}))
	list3.Append(widget.NewButton("上一页", func() {
		fmt.Println("上一页")
	}))
	list3.Append(widget.NewButton("下一页", func() {
		fmt.Println("下一页")
	}))
	list.Append(list3)
	for i,b := range books{
		index := i // capture
		list2 := widget.NewHBox()

		entrys[index].author.SetPlaceHolder("作者")
		list2.Append(&entrys[index].author)

		entrys[index].bookname.SetPlaceHolder("书名")
		list2.Append(&entrys[index].bookname)

		for i,t :=range b.text{
			indext:=i
			list2.Append(widget.NewButton(t, func() {
				fmt.Println("",books[index].text[indext])
				if editname==0{
					entrys[index].bookname.SetText(entrys[index].bookname.Text+books[index].text[indext])

				} else{
					entrys[index].author.SetText(entrys[index].author.Text+books[index].text[indext])


				}
			}))
		}




		list.Append(list2)



	}
	horiz := widget.NewHScrollContainer(list)

	w.SetContent(horiz)

	w.ShowAndRun()
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

func show_books(book[] Book){
	var s string
	var words []string
	use_hmm := true
	x := gojieba.NewJieba()
	defer x.Free()
	x.AddWord("厄普代克")
	x.AddWord("辻村深月")

	for i,b :=range book{
		s=b.file_name
		words = x.Cut(s,use_hmm)
		book[i].text=words

		fmt.Println(s+" :")

		fmt.Println( "    ",strings.Join(words, "/"))

	}
}