package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"github.com/xederro/portfolio/shortener-service/cmd/auth"
	"github.com/xederro/portfolio/shortener-service/cmd/shortener"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

var url = fmt.Sprintf("%s?authToken=%s", os.Getenv("TursoDataBaseLink"), os.Getenv("TursoDataBaseToken"))

type App struct {
	shortener.UnimplementedShortenerServiceServer
}

func (d App) GetOne(ctx context.Context, s *shortener.ShortenerGetOneRequest) (*shortener.ShortenerGetOneResponse, error) {
	db := d.StartDB()
	defer db.Close()

	var result shortener.ShortenerGetOneResponse

	rows, err := db.Query("SELECT long FROM shortener WHERE short = ? LIMIT 1;", s.Short)
	if err != nil {
		fmt.Println("failed to execute query: %v\n", err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	if rows.Next() {
		err := rows.Scan(&result.Long)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
	} else {
		return nil, errors.New("No link with provided short")
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error during rows iteration:", err)
		return nil, err
	}

	return &result, nil
}

func (d App) GetAll(ctx context.Context, s *shortener.ShortenerGetAllRequest) (*shortener.ShortenerGetAllResponse, error) {
	db := d.StartDB()
	defer db.Close()
	var results shortener.ShortenerGetAllResponse

	uuid, err := d.GetUUID(s.Token)
	if err != nil {
		fmt.Println("Error while getting UUID")
		return nil, err
	}

	rows, err := db.Query("SELECT short, long FROM shortener WHERE user = ?;", uuid)
	if err != nil {
		log.Fatalf("failed to execute query: %v\n", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		var row shortener.ShortenerRow

		if err := rows.Scan(&row.Short, &row.Long); err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}

		results.Rows = append(results.Rows, &row)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error during rows iteration:", err)
		return nil, err
	}

	return &results, nil
}

func (d App) Insert(ctx context.Context, s *shortener.ShortenerInsertRequest) (*shortener.ShortenerInsertResponse, error) {
	db := d.StartDB()
	defer db.Close()
	var short shortener.ShortenerInsertResponse

	uuid, err := d.GetUUID(s.Token)
	if err != nil {
		fmt.Println("Error while getting UUID")
		return nil, err
	}

	rows, err := db.Query("INSERT INTO shortener (long, user) values (?,?);", s.Long, uuid)
	if err != nil {
		log.Fatalf("failed to execute query: %v\n", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	rows, err = db.Query("SELECT max(short) FROM shortener;")
	if err != nil {
		log.Fatalf("failed to execute query: %v\n", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	if rows.Next() {
		err := rows.Scan(&short.Short)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error during rows iteration:", err)
		return nil, err
	}

	return &short, nil
}

func (d App) Delete(ctx context.Context, s *shortener.ShortenerDeleteRequest) (*shortener.ShortenerDeleteResponse, error) {
	db := d.StartDB()
	defer db.Close()
	uuid, err := d.GetUUID(s.Token)
	if err != nil {
		fmt.Println("Error while getting UUID")
		return nil, err
	}

	rows, err := db.Query("delete from shortener where short = ? and user = ?;", s.Short, uuid)
	if err != nil {
		log.Fatalf("failed to execute query: %v\n", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	return &shortener.ShortenerDeleteResponse{}, nil
}

func (d App) GetUUID(token string) (string, error) {
	q := auth.AuthRequest{
		Token: token,
	}
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial("auth:8000", opts...)
	if err != nil {
		log.Println("There was an error while establishing connection")
		return "", err
	}
	defer conn.Close()

	client := auth.NewAuthServiceClient(conn)

	response, err := client.CheckAuth(context.TODO(), &q)
	if err != nil {
		log.Println("There was an error")
		return "", err
	}

	if !response.GetIsAuth() {
		err := errors.New("user cannot be authenticated")
		log.Println(err)
		return "", err
	}

	return response.GetUser(), nil
}

func (d App) StartDB() *sql.DB {
	db, err := sql.Open("libsql", url)
	if err != nil {
		log.Fatalf("failed to open db %s: %s", url, err)
	}
	return db
}
