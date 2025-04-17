# Go Transaction Manager

This is a simple command-line application written in Go to manage financial transactions. It allows users to add transactions and view the current list of transactions along with the calculated balance.

## Features

-   Add multiple transactions at once (separated by spaces).
-   View all recorded transactions.
-   Display the current balance based on the entered transactions.

## Files

-   `<mcfile name="main.go" path="z:\Projects\go\Learning\ARRAY_LEARNING\main.go"></mcfile>`: Contains the main application logic, user interface loop, and calls to helper functions.
-   `<mcfile name="transactions_helper.go" path="z:\Projects\go\Learning\ARRAY_LEARNING\transactions_helper.go"></mcfile>`: Includes helper functions for parsing transaction input (<mcsymbol name="getTransactions" filename="transactions_helper.go" path="z:\Projects\go\Learning\ARRAY_LEARNING\transactions_helper.go" startline="19" type="function"></mcsymbol>) and calculating the balance (<mcsymbol name="calculateBalance" filename="transactions_helper.go" path="z:\Projects\go\Learning\ARRAY_LEARNING\transactions_helper.go" startline="9" type="function"></mcsymbol>).

## How to Run

1.  **Clone the repository (if applicable) or ensure you have the Go files.**
2.  **Navigate to the project directory in your terminal:**
    ```bash
    cd z:\Projects\go\Learning\ARRAY_LEARNING
    ```
3.  **Build the executable:**
    ```bash
    go build
    ```
4.  **Run the executable:**
    -   On Windows:
        ```bash
        .\ARRAY_LEARNING.exe
        ```
    -   On macOS/Linux:
        ```bash
        ./ARRAY_LEARNING
        ```
5.  **Follow the on-screen prompts:**
    -   Enter `1` to add transactions (e.g., `100 -50 25.5`).
    -   Enter `2` to view the list of transactions.
    -   Enter `3` to exit the application.
