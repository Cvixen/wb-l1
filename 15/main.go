package main

/**
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}
*/

//В Go самый простой способ создать утечку памяти - определить глобальную переменную, массив и добавить данные в этот массив. Также сборщик  мусора
//не сможет очистить переменную v, потому что на нее ссылается justString (поведение слайсов)

/**
func someFunc() string {
	v := createHugeString(1 << 10)
	justString := make([]byte, 100)
	copy(justString, v)

	return string(justString)
}

func main() {
	a := someFunc()

}
*/
