package pkg

import (
	"fmt"

	"github.com/pkg/errors"
)

func IsStrBalanced(str string) (bool, error) {
	if len(str)%2 != 0 { //Предварительная проверка. Если длина строки не делится на 2, значит какой-то скобке не хватило пары
		return false, nil
	}
	err := errors.New("")
	var round, square, curly int //счетчики скобок по парам
	for i := 0; i < len(str); i++ {
		switch str[i] {
		case '(':
			round++
		case ')':
			round--
			//проверка,не стоит ли закрывающая скобка первее открывающей
			if round < 0 {
				return false, nil
			}
		case '[':
			square++
		case ']':
			square--
			if square < 0 {
				return false, nil
			}
		case '{':
			curly++
		case '}':
			curly--
			if curly < 0 {
				return false, nil
			}
		default:
			return false, errors.Wrap(err, fmt.Sprintf("Invalid character %s. Use only ()[]{} brackets to check", string(str[i])))
		}
	}
	//проверяем, все ли скобки закрываются
	if round+square+curly == 0 {
		return true, nil
	}
	return false, nil
}
