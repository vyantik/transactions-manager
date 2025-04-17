# Менеджер транзакций на Go

Это простое консольное приложение, написанное на Go, для управления финансовыми транзакциями. Оно позволяет пользователям добавлять транзакции и просматривать текущий список транзакций вместе с рассчитанным балансом.

## Возможности

-   Добавление нескольких транзакций за раз (через пробел).
-   Просмотр всех записанных транзакций.
-   Отображение текущего баланса на основе введенных транзакций.

## Файлы

-   `<mcfile name="main.go" path="z:\Projects\go\Learning\ARRAY_LEARNING\main.go"></mcfile>`: Содержит основную логику приложения, цикл пользовательского интерфейса и вызовы вспомогательных функций.
-   `<mcfile name="transactions_helper.go" path="z:\Projects\go\Learning\ARRAY_LEARNING\transactions_helper.go"></mcfile>`: Включает вспомогательные функции для разбора ввода транзакций (<mcsymbol name="getTransactions" filename="transactions_helper.go" path="z:\Projects\go\Learning\ARRAY_LEARNING\transactions_helper.go" startline="19" type="function"></mcsymbol>) и расчета баланса (<mcsymbol name="calculateBalance" filename="transactions_helper.go" path="z:\Projects\go\Learning\ARRAY_LEARNING\transactions_helper.go" startline="9" type="function"></mcsymbol>).

## Как запустить

1.  **Клонируйте репозиторий (если применимо) или убедитесь, что у вас есть файлы Go.**
2.  **Перейдите в каталог проекта в вашем терминале:**
    ```bash
    cd z:\Projects\go\Learning\ARRAY_LEARNING
    ```
3.  **Соберите исполняемый файл:**
    ```bash
    go build
    ```
4.  **Запустите исполняемый файл:**
    -   В Windows:
        ```bash
        .\ARRAY_LEARNING.exe
        ```
    -   В macOS/Linux:
        ```bash
        ./ARRAY_LEARNING
        ```
5.  **Следуйте инструкциям на экране:**
    -   Введите `1`, чтобы добавить транзакции (например, `100 -50 25.5`).
    -   Введите `2`, чтобы просмотреть список транзакций.
    -   Введите `3`, чтобы выйти из приложения.
