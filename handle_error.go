package main

import (
	"database/sql"
	"fmt"
	xerrors "github.com/pkg/errors"
)

// 问题：是否已经wrap dao 层的 sql.ErrNoRows
// 应该wrap 这个error,因为sql.ErrNoRows 是标准库sql 的一个sentinel error,需要handle他，并增加一些额外的信息

func SelectPeopleById(id int) error {
	errorString := fmt.Sprintf("main.SelectPeopleById select people no rows id: %d ", id)
	return xerrors.Wrap(sql.ErrNoRows, errorString)
}

func main() {
	err := SelectPeopleById(123)
	fmt.Printf("%+v", err)
	if xerrors.Cause(err) == sql.ErrNoRows {
		fmt.Printf("\nyes")
	}
}
