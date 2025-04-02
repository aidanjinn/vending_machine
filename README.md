# Vending Machine Simulator

This program simulates a vending machine that processes transactions from a text file input. It's designed as a practice exercise for system design, focusing on handling coin insertion, item purchases, and change calculation.

## Purpose

This project was created as a redemption for a failed OA and practice exercise for system design, specifically to improve understanding of handling input streams, managing state, and implementing business logic.

## Features

-   **Coin Insertion:** Processes coin insertions and updates the current balance.
-   **Item Purchase:** Handles item purchases, calculates change, and updates the vending machine's inventory.
-   **Change Return:** Returns inserted coins when requested.
-   **Change Calculation:** Determines if change can be provided and returns the correct denominations.
-   **Input from Text File:** Reads transaction data from a text file, allowing for batch processing of events.
-   **Error Handling:** Provides basic error handling for invalid input and insufficient funds.

## Input Format

The program reads input from a text file with the following format:

1.  **Line 1:** Number of events and number of item options (e.g., `5 3`).
2.  **Line 2:** Initial amounts of coins in the change drawer (10, 50, 100, 500 yen) separated by spaces (e.g., `10 5 2 1`).
3.  **Lines 3 to (2 + number of items):** Item details: item name, price, and stock (e.g., `Cola 120 5`).
4.  **Remaining lines:** Event commands:
    -   `+ <coin>`: Insert a coin (e.g., `+ 100`).
    -   `* <item_name>`: Purchase an item (e.g., `* Cola`).
    -   `#`: Return inserted coins.

## Example Input File (test.txt)

```text
5 3
10 5 2 1
Cola 120 5
Water 100 10
Juice 150 3
+ 100
+ 50
* Cola
#
+ 500
