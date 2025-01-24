package main

import (
    "context"
    "errors"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "github.com/golang-migrate/migrate/v4"
    _ "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
    "github.com/jackc/pgx/v5/pgxpool"
    "net/http"
    "sego6/entity"
)

var db *pgxpool.Pool

func main() {
    var err error

    r := gin.Default()

    r.Use(cors.New(cors.Config{
        AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
        AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
        AllowOrigins: []string{"http://localhost:5173"},
    }))

    api := r.Group("/api")

    api.POST("/account", createAccount)
    api.GET("/accounts", getAccounts)
    api.POST("/transaction", createTransaction)
    api.GET("transactions", getTransactions)
    api.GET("/categories", getCategories)

    db, err = pgxpool.New(context.Background(), "postgres://postgres:postgres@localhost:5432/sego?sslmode=disable")
    if err != nil {
        logger.Fatal(err)
    }
    err = db.Ping(context.Background())
    if err != nil {
        logger.Fatal(err)
    }
    m, err := migrate.New("file://migrations", "postgres://postgres:postgres@localhost:5432/sego?sslmode=disable")
    if err != nil {
        logger.Fatal(err)
    }
    err = m.Up()
    if err != nil && !errors.Is(err, migrate.ErrNoChange) {
        logger.Fatal(err)
    }
    defer func() {
        srcErr, dbErr := m.Close()
        if srcErr != nil {
            logger.Fatal(srcErr)
        }
        if dbErr != nil {
            logger.Fatal(dbErr)
        }
    }()

    err = r.Run(":8080")
    if err != nil {
        logger.Fatal(err)
    }
}

func createAccount(c *gin.Context) {
    var account entity.Account
    if err := c.ShouldBindJSON(&account); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    _, err := db.Exec(context.Background(),
        "INSERT INTO account(name, type, balance, created_at, currency_code) VALUES ($1, $2, $3, $4, $5)", account.Name, account.Type, account.Balance, account.CreatedAt, account.Currency)
    if err != nil {
        logger.Error(err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"account": account})
}

func getAccounts(c *gin.Context) {
    var accounts []entity.Account
    rows, err := db.Query(context.Background(), "SELECT * FROM account")
    if err != nil {
        logger.Error(err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()
    for rows.Next() {
        var account entity.Account
        if err := rows.Scan(&account.Name, &account.Type, &account.Balance, &account.CreatedAt, &account.Currency); err != nil {
            logger.Error(err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        accounts = append(accounts, account)
    }
    c.JSON(http.StatusOK, gin.H{"accounts": accounts})
}

func createTransaction(c *gin.Context) {
    var transaction entity.Transaction
    if err := c.ShouldBindJSON(&transaction); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    _, err := db.Exec(context.Background(),
        "INSERT INTO \"transaction\"(amount, type, created_at, account, category) VALUES ($1, $2, $3, $4, $5)",
        transaction.Amount, transaction.Type, transaction.CreatedAt, transaction.Account, transaction.Category,
    )
    if err != nil {
        logger.Error(err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    switch transaction.Type {
    case "Deposit":
        _, err := db.Exec(context.Background(),
            "UPDATE account SET balance = balance + $1 WHERE name = $2",
            transaction.Amount, transaction.Account,
        )
        if err != nil {
            logger.Error(err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
    case "Withdrawal":
        _, err := db.Exec(context.Background(),
            "UPDATE account SET balance = balance - $1 WHERE name = $2",
            transaction.Amount, transaction.Account,
        )
        if err != nil {
            logger.Error(err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
    }

    c.JSON(http.StatusCreated, gin.H{"transaction": transaction})
}

func getCategories(c *gin.Context) {
    var categories []entity.Category
    rows, err := db.Query(context.Background(), "SELECT * FROM category")
    if err != nil {
        logger.Error(err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()
    for rows.Next() {
        var category entity.Category
        if err := rows.Scan(&category.Title); err != nil {
            logger.Error(err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        categories = append(categories, category)
    }

    c.JSON(http.StatusOK, gin.H{"categories": categories})
}

func getTransactions(c *gin.Context) {
    var transactions []entity.Transaction
    rows, err := db.Query(context.Background(), "SELECT transaction_id, amount, type, created_at, account, category FROM transaction")
    if err != nil {
        logger.Error(err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()
    for rows.Next() {
        var transaction entity.Transaction
        if err := rows.Scan(&transaction.ID, &transaction.Amount, &transaction.Type, &transaction.CreatedAt, &transaction.Account, &transaction.Category); err != nil {
            logger.Error(err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        transactions = append(transactions, transaction)
    }

    c.JSON(http.StatusOK, gin.H{"transactions": transactions})
}
