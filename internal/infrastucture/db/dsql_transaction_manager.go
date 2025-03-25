package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

func NewDSQLTransactionManager(ctx context.Context) (*pgxpool.Pool, *TransactionManager, error) {
	region := os.Getenv("DSQL_REGION")
	clusterEndpoint := os.Getenv("DSQL_CLUSTER_ENDPOINT")
	if region == "" || clusterEndpoint == "" {
		panic("AWS_REGION or AWS_CLUSTER_ENDPOINT is not set.")
	}

	pool, err := GetPool(ctx, region, clusterEndpoint)
	if err != nil {
		return pool, nil, err
	}
	var manager = NewTransactionManager(pool)
	return pool, &manager, nil
}
