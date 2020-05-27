package main

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"strconv"
)

var root_path string = "/home/hanyuluo/文档/books"

func main() {

	GetFilelist(root_path)
	ShowBooks(books)

	app := app.New()
	app.Settings().SetTheme(theme.LightTheme())

	w := app.NewWindow("御书")
	w.Resize(fyne.NewSize(1200, 400))
	list := widget.NewVBox() //总体纵向列表

	listBooks := widget.NewVBox()

	list3 := widget.NewHBox() //顶层功能按钮横向列表

	list3.Append(widget.NewButton("修改作者", func() {
		fmt.Println("修改作者")
		editname = 1
	}))
	list3.Append(widget.NewButton("修改书名", func() {
		fmt.Println("修改书名")
		editname = 0
	}))

	page_label := widget.NewLabel(strconv.Itoa(min) + " -- " + strconv.Itoa(max) + " 共" + strconv.Itoa(len(books)+1))
	list3.Append(page_label)

	list3.Append(widget.NewButton("保存信息", func() {
		fmt.Println("保存信息")
		for i, _ := range books[min:max] {
			index := i // capture
			books[min+index].book_name = entrys[index].bookname.Text
			books[min+index].author_name = entrys[index].author.Text

		}
	}))

	list3.Append(widget.NewButton("全部信息", func() {
		list := widget.NewVBox() //总体纵向列表
		for i, b := range books {
			index := i // capture
			page_label := widget.NewLabel(string(index) + b.book_name)
			list.Append(page_label)

		}
		dialog.ShowCustom("全部信息", "Done", list, w)
	}))

	list3.Append(widget.NewButton("另存到文件夹", func() {
		fmt.Println("另存到文件夹")

		SaveBooks()

	}))

	list4 := widget.NewHBox() //底层功能按钮横向列表
	list4.Append(widget.NewButton("上一页", func() {
		for i, _ := range books[min : max+1] {
			index := i // capture
			books[min+index].book_name = entrys[index].bookname.Text
			books[min+index].author_name = entrys[index].author.Text

		}
		if min < 10 {

		} else {
			min -= 10
			max -= 10
			if max < 10 {
				max = 9
			}
		}
		page_label.SetText(strconv.Itoa(min) + " -- " + strconv.Itoa(max))

		list.Children[1] = SwitchBooksPages(min, max)

	}))
	list4.Append(widget.NewButton("下一页", func() {
		for i, _ := range books[min : max+1] {
			index := i // capture
			books[min+index].book_name = entrys[index].bookname.Text
			books[min+index].author_name = entrys[index].author.Text

		}
		min += 10
		if min > len(books) {
			min -= 10
			max = len(books) - 1
		} else {
			max += 10
			if max >= len(books) {
				max = len(books) - 1
			}
		}

		page_label.SetText(strconv.Itoa(min) + " -- " + strconv.Itoa(max))

		for i, _ := range books[min : max+1] {
			index := i
			entrys[index].author.SetText(books[min+index].author_name)
			entrys[index].bookname.SetText(books[min+index].book_name)
		}
		list.Children[1] = SwitchBooksPages(min, max)

	}))

	list.Append(list3)
	list.Append(listBooks)
	list.Append(list4)

	horiz := widget.NewHScrollContainer(list)

	w.SetContent(horiz)

	w.ShowAndRun()
}
