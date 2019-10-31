package manager

import "../bean"

var myWifes map[string]bean.WifeBean = make(map[string]bean.WifeBean)

func AddWife(name string,wife bean.WifeBean)  {
	_,ok := myWifes[name]
	if !ok {
		myWifes[name] = wife
	}
}

func RemoveWife(name string)  {
	_,ok := myWifes[name]
	if ok {
		delete(myWifes, name)
	}
}

func GetMyWifes() map[string]bean.WifeBean  {
	return myWifes
}

