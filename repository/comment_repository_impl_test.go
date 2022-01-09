package repository

import (
	"context"
	"fmt"
	go_database "go-database"
	"go-database/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRespository := NewCommentRepository(go_database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "Dimas159@gmail.com",
		Comment: "Test Repo",
	}

	result, err := commentRespository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindByID(t *testing.T) {
	commentRespository := NewCommentRepository(go_database.GetConnection())

	ctx := context.Background()

	result, err := commentRespository.FindById(ctx, 21)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindAll(t *testing.T) {
	commentRespository := NewCommentRepository(go_database.GetConnection())

	ctx := context.Background()

	result, err := commentRespository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, res := range result {
		fmt.Println(res)
	}
}
