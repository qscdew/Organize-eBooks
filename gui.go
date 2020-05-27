package main

import "fyne.io/fyne/widget"

type entry_bookname_author struct {
	bookname widget.Entry
	author   widget.Entry
}

var entrys [10]entry_bookname_author
var editname int

var min int = 0
var max int = 9

func SwitchBooksPages(min, max int) *widget.Box {
	listBooks := widget.NewVBox()
	for i, b := range books[min : max+1] {
		index := i // capture
		list2 := widget.NewHBox()

		list2.Append(widget.NewButton("全部", func() {

			if editname == 0 {
				entrys[index].bookname.SetText(books[min+index].file_name)

			} else {
				entrys[index].author.SetText(books[min+index].file_name)
			}
		}))

		list2.Append(widget.NewButton("清空", func() {
			if editname == 0 {
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

		for i, t := range b.text {
			indext := i
			button := widget.NewButton(t, func() {

				if editname == 0 {
					entrys[index].bookname.SetText(entrys[index].bookname.Text + books[min+index].text[indext])
				} else {
					entrys[index].author.SetText(entrys[index].author.Text + books[min+index].text[indext])
				}

			})
			list2.Append(button)
		}

		listBooks.Append(list2)

	}
	return listBooks
}
