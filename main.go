package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Item struct {
	price int
	stock int
}

// Declaration of the vending machine object type
type vendingMachine struct {
	// 10, 50, 100, 500
	change_drawer map[int]int

	// map containing the machines Items
	items map[string]Item

	// queue to keep track of the currently inserted coins
	// for the instance in which we need to return inserted
	inserted_coins []int

	// bool to keep track of if the machine has in standby or not
	standby_mode bool

	// current balance of the user (amount of money they have inserted)
	current_balance int
}

func (vm *vendingMachine) add_coin(coin int) {

	if _, ok := vm.change_drawer[coin]; !ok {
		fmt.Println("Denomination not Supported")
	}
	vm.change_drawer[coin] += 1
	vm.inserted_coins = append(vm.inserted_coins, coin)
	vm.current_balance += coin

	fmt.Println(vm.current_balance, 0, "-")
}

func (vm *vendingMachine) change_possible(coin int) (bool, map[int]int) {

	change_amount := vm.current_balance - coin

	var copy_drawer map[int]int = make(map[int]int)
	for k, v := range vm.change_drawer {
		copy_drawer[k] = v
	}

	// make an empty change drawer
	change := map[int]int{10: 0, 50: 0, 100: 0, 500: 0}

	// we should use as many 500 yen coins as possible
	for change_amount >= 500 && copy_drawer[500] > 0 {
		change_amount -= 500
		copy_drawer[500] -= 1
		change[500] += 1
	}

	// next we do 100 yen coins
	for change_amount >= 100 && copy_drawer[100] > 0 {
		change_amount -= 100
		copy_drawer[100] -= 1
		change[100] += 1
	}

	if change_amount >= 100 {
		return false, change
	}

	for change_amount >= 50 && copy_drawer[50] > 0 {
		change_amount -= 50
		copy_drawer[50] -= 1
		change[50] += 1
	}

	for change_amount >= 10 && copy_drawer[10] > 0 {
		change_amount -= 10
		copy_drawer[10] -= 1
		change[10] += 1
	}

	return change_amount == 0, change
}

func (vm *vendingMachine) change_lever() {

	if vm.standby_mode {
		return
	}

	for coin := range vm.inserted_coins {
		vm.change_drawer[coin] -= 1
	}

	fmt.Println(0, vm.current_balance, "-")
	vm.current_balance = 0
	vm.inserted_coins = vm.inserted_coins[:0]
}

func (vm *vendingMachine) dispense_item(item_name string, change_amount map[int]int) {

	for k, v := range change_amount {
		vm.change_drawer[k] -= v
	}

	reduce := vm.items[item_name].stock
	reduce -= 1
	vm.current_balance -= vm.items[item_name].price
	vm.inserted_coins = vm.inserted_coins[:0]
	fmt.Println(vm.current_balance, 0, item_name)
}

func main() {
	var vm = vendingMachine{}

	filename := "test.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()

	// Line One is Events, Items
	lineOne := scanner.Text()
	parts := strings.Split(lineOne, " ")

	if len(parts) != 2 {
		fmt.Println("Incorrect number of parts")
		return
	}

	numEvents, err := strconv.Atoi(parts[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	numItems, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	//Read Line Two
	scanner.Scan()
	lineTwo := scanner.Text()
	change_amounts := strings.Split(lineTwo, " ")
	// Next Line is the starting change draw amounts
	change_drawer := make(map[int]int)
	denominations := []int{10, 50, 100, 500}

	for i, amountStr := range change_amounts {
		amount, err := strconv.Atoi(amountStr)
		if err != nil {
			fmt.Println(err)
			return
		}
		change_drawer[denominations[i]] = amount
	}

	//Read the Items within Machine
	items := make(map[string]Item)
	for range numItems {
		scanner.Scan()
		curr_line := scanner.Text()
		item_amounts := strings.Split(curr_line, " ")
		item_name := item_amounts[0]
		price, err := strconv.Atoi(item_amounts[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		stock, err := strconv.Atoi(item_amounts[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		curr_Item := Item{price, stock}
		items[item_name] = curr_Item
	}

	var inserted_coins []int
	vm.inserted_coins = inserted_coins

	var standby_mode = false
	vm.standby_mode = standby_mode

	var current_balance = 0
	vm.current_balance = current_balance

	vm.items = items
	vm.change_drawer = change_drawer

	for range numEvents {

		scanner.Scan()
		curr_line := scanner.Text()
		inputs := strings.Split(curr_line, " ")

		command := inputs[0]

		if command == "+" {
			coin_amount, err := strconv.Atoi(inputs[1])

			if err != nil {
				fmt.Println(err)
				return
			}

			vm.add_coin(coin_amount)

		} else if command == "*" {

			item_name := inputs[1]
			// check if item exists within machine
			if item, ok := vm.items[item_name]; !ok {
				fmt.Printf("%s not valid item\n", item)
				return
			}

			curr_item := vm.items[item_name]
			possible, change := vm.change_possible(curr_item.price)

			if possible {
				vm.dispense_item(item_name, change)
			} else {
				fmt.Println(vm.current_balance, 0, "-")
			}

		} else if command == "#" {
			vm.change_lever()
		}

	}
}
