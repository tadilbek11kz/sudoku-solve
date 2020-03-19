package main

import (
	"fmt"
	"os"
)

//Создаем пустой массив двумерных массивов для хранения там решенных судоку
var boards [][9][9]rune

//Backtracking Algo
func Sudoku(board [9][9]rune) bool {
	//Если в переменной решенных судоку есть два и больше решение то возвращаем правда(true), чтобы остановить рекурсию
	if len(boards) > 1 {
		return true
	}
	/*
		Запуск алгоритма проверки доски на наличие пустых клеток.
		Если пустых клеток не найдено то доска решена.
		И мы добавляем решение в переменную решенных судоку.
		Возвращаем ложь(false) чтобы продолжить поиск решений для судоку(Чтобы проверить наше судоку на уникальное решение)
	*/
	if checkDone(board) {
		boards = appendBoard(boards, board)
		return false
	}
	//Создаем две переменные(row, col) для последующего поиска пустых клеток
	row, col := -1, -1
	//Запускаем алгоритм для поиска пустых клеток(FindEmpty)
	find := FindEmpty(board)
	//Проверяеми есть ли пустые клетки в поле
	if len(find) == 0 {
		//Если пустых клеток нету то возвращаем true(правда), что означает то что судоку уже заполнен
		return true
	} else {
		//Если пустая клетка найдена то присваеваем ее положения в переменные row(строка) и col(столбец)
		row, col = find[0], find[1]
	}
	/*
		Цикл для подбора числа в пустую клетку
		Ходит от 1 по ascii до 9 по ascii
	*/
	for i := '1'; i <= '9'; i++ {
		//Проверяем можем ли мы поставить число (i) в клетку [row][col]
		if canPlaceValue(board, row, col, i) {
			//Если число можно поставить в клетку [row][col], то заменяем келтку [row][col] на число (i)
			board[row][col] = i
			//Запускаем рекурсию
			if Sudoku(board) {
				return true
			}
			//Если клетка подобрана не верна сбрасываем ее
			board[row][col] = '.'
		}
	}
	return false
}

//Проверка решена ли доска
func checkDone(board [9][9]rune) bool {
	//Перебираем всю доску по row и col
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			//Если найдена пустая клетка '.'(0) то возвращаем ложь(false), что означает что доска не решена
			if board[i][j] == '.' {
				return false
			}
		}
	}
	//Если пустых клеток нету то возвращаем правда(true), что означает что судоку решен
	return true
}

//Самописный append для решенных досок
func appendBoard(boards [][9][9]rune, board [9][9]rune) [][9][9]rune {
	//Проверяем решенную доску на уникальность
	if !uniqueCheck(boards, board) {
		//Если такого решения еще нету в переменной с решениями, то добавляем его
		//Создаем временный массив с размером на 1 больше, чем массив с решениями
		newBoards := make([][9][9]rune, len(boards)+1)
		//Перебираем массив с решениями и присваеваем решения новому массиву
		for i := 0; i < len(boards); i++ {
			newBoards[i] = boards[i]
		}
		//Добавляем наше новое решение в массив с решениями
		newBoards[len(newBoards)-1] = board
		//Возвращаем наш новый массив с решениями
		return newBoards
	}
	//Если решение уже есть в переменной с решения то возвращаем переменную с рещениями без изменения
	return boards
}

//Проверка на уникальность
func uniqueCheck(boards [][9][9]rune, board [9][9]rune) bool {
	//Проверка есть решение в массиве с решениями
	//Перебираем массив с решения и сверяем его с новым решением
	for _, b := range boards {
		if b == board {
			//Если такое решение уже есть то возвращаем правда(true)
			return true
		}
	}
	//Если такого решения еще нету, то возвращаем ложь(false)
	return false
}

//Валдитор для судоку
func canPlaceValue(board [9][9]rune, row int, col int, charToPlace rune) bool {
	//Проверка на наличие charToPlace(число которе мы хотим вставить) по вертикали(rows check)
	for _, i := range board {
		/*
			Если charToPlace(число которе мы хотим вставить)
			уже лежит на вертикали то возвращаем false
			(данный символ в клетку [row][col] поставить нельзя)
		*/
		if charToPlace == i[col] {
			return false
		}
	}
	//Проверка на наличие charToPlace(число которе мы хотим вставить) по горизонтали(col check)
	for i := 0; i < len(board[row]); i++ {
		/*
			Если charToPlace(число которе мы хотим вставить)
			уже лежит на горизонтали то возвращаем false
			(данный символ в клетку [row][col] поставить нельзя)
		*/
		if charToPlace == board[row][i] {
			return false
		}
	}
	//Размер полей внутри судоку(subGrid)
	subGridSize := 3

	//Индекс по вертикали внутреннего поля(subGrid) в котором находиться элемент для которого мы ищем значение ([row][col])
	vertBoxIndex := row / subGridSize
	//Индекс по горизонтали внутреннего поля(subGrid) в котором находиться элемент для которого мы ищем значение([row][col])
	horizBoxIndex := col / subGridSize

	//Индекс левого верхнего угла (по вертикали) внутреннего поля(subGrid)
	topLeftSubBoxRow := subGridSize * vertBoxIndex
	//Индекс левого верхнего угла (по горизонтали) внутреннего поля(subGrid)
	topLeftSubBoxCol := subGridSize * horizBoxIndex

	//Проверка внутреннго поля (subGrid)(размером subGridSize x subGridSize) на наличие charToPlace(число которе мы хотим вставить)
	for i := 0; i < subGridSize; i++ {
		for j := 0; j < subGridSize; j++ {
			/*
				Если charToPlace(число которе мы хотим вставить)
				уже лежит во внутреннем поле(subGrid) то возвращаем false
				(данный символ в клетку [row][col] поставить нельзя)
			*/
			if charToPlace == board[topLeftSubBoxRow+i][topLeftSubBoxCol+j] {
				return false
			}
		}
	}
	//Если все проверки пройдены то данное число можно установить в клетку ([row][col])
	return true

}

//Поиск пустого элемента в судоку
func FindEmpty(board [9][9]rune) []int {
	/*
		Перебираем двумерный массив(поле судоку (массив board))
		переменная i отвечает за вертикаль(за строки)
		переменная j отвечает за горизонталь(за столбцы)
	*/
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				//Если значение поля в клетке [i][j] равно "."(пустая клетка), то возвращаем положения этой пустой клетки [row][col]([i][j])
				return []int{i, j}
			}
		}
	}
	//Если пустая клетка не найдено, то возвращаем nil(пусто)
	return nil
}

//Вывод доски
func printBoardTable(board [9][9]rune) {
	for i := 0; i < len(board); i++ {
		if i%3 == 0 && i != 0 {
			fmt.Println("- - - - - - - - - -")
		}
		for j := 0; j < len(board[0]); j++ {
			if j%3 == 0 && j != 0 {
				fmt.Printf("%v", "|")
			}
			if j == 8 {
				fmt.Printf("%v\n", board[i][j])
			} else {
				fmt.Printf("%v ", board[i][j])
			}
		}
	}
}

//Вывод доски
func printBoard(board [9][9]rune) {
	//Перебираем всю доску и выводим ее
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(board[i][j] - '0')
			if j != 8 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

//Проверка аргументов
func checkArgs() bool {
	//Проверка кол-ва аргументов(Для корректного судоку должно быть 9 аргументов)
	//len(os.Args) - 1 потому-что первый элемента массива os.Args это названия файла
	if len(os.Args)-1 == 9 {
		for _, i := range os.Args[1:] {
			if len(i) != 9 {
				return false
			}
			for _, j := range i {
				//Перебираем все строки и колонки и проверяем каждый элемент. Элемент должен быть в диапозоне от 1 до 9 или '.'(0)
				if !(j >= '1' && j <= '9' || j == '.') {
					//Если элемент не является чисилом от 1-9 или не является '.'(0) то возвращаем ложь(false)
					return false
				}
			}
		}
		//Если проверка пройдена возвращаем правда(true)
		return true
	}
	//Если кол-во аргументов массиве меньше или больше 9 то возвращаем ложь(false)
	return false
}

//Создание двумерного массива
func createBoard(args []string) [9][9]rune {
	//Создаем пустой двумерный массив
	board := [9][9]rune{}
	//Перебираем все аргументы и все элементы каждого аргумента
	for row, i := range args {
		for col, j := range i {
			//Если элемент(j) входит в диапозон от 1-9 или равен '.'(0)
			if j >= '1' && j <= '9' || j == '.' {
				board[row][col] = j
			}
		}
	}
	//Возвращаем двумерный массив хранящий наше судоку
	return board
}

//Проверка доски для судоку
func checkBoard(board [9][9]rune) bool {
	//Перебираем судоку по row и col
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			//Проверяем если на позиции [i][j] не пусто, то запускаем алгоритм проверки данного числа
			if board[i][j] != '.' {
				//Создаем временную переменную чтобы сохранить в ней число на позиции [i][j]
				tmp := board[i][j]
				//Очищаем положение [i][j] в судоку
				board[i][j] = '.'
				//Запускаем проверку для определения может ли стоять это число(tmp) на клетке [i][j]
				if !canPlaceValue(board, i, j, tmp) {
					//Если не может возвращаем ложь(false)
					return false
				}
				//Если проверка прошла возвращаем значение клетке [i][j]
				board[i][j] = tmp
			}
		}
	}
	//Вся проверка пройдена, возвращаем правду(true)
	return true
}

func main() {
	//Проверка правильности ввода аргументов
	//Иначе выдаем ошибку
	if checkArgs() {
		//С помощью функции createBoard() создаем доску
		board := createBoard(os.Args[1:])
		//Проверяем доску на коректность
		//Иначе выдаем ошибку
		if !checkBoard(board) {
			fmt.Println("Error")
			return
		}
		//Запускаем наш solver
		Sudoku(board)
		//Если найденно всего 1 решение то выводим его
		//Иначе выдаем ошибку
		if len(boards) == 1 {
			//Алгоритм для отрисовывания доски судоку
			printBoard(boards[0])
		} else {
			fmt.Println("Error")
		}
	} else {
		fmt.Println("Error")
	}
}
