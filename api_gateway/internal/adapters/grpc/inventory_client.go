package grpc

import (
	pb "api_gateway/proto/inventory"
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

// InventoryServiceClient provides a wrapper for interacting with the inventory service
type InventoryServiceClient struct {
	productClient  pb.ProductServiceClient
	categoryClient pb.CategoryServiceClient
	conn           *grpc.ClientConn
}

// NewInventoryServiceClient creates a new gRPC client for the inventory service
func NewInventoryServiceClient(inventoryServiceURL string) (*InventoryServiceClient, error) {
	conn, err := grpc.Dial(inventoryServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &InventoryServiceClient{
		productClient:  pb.NewProductServiceClient(conn),
		categoryClient: pb.NewCategoryServiceClient(conn),
		conn:           conn,
	}, nil
}

// Close closes the client connection
func (c *InventoryServiceClient) Close() error {
	return c.conn.Close()
}

// Product service methods
func (c *InventoryServiceClient) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return c.productClient.CreateProduct(ctx, req)
}

func (c *InventoryServiceClient) GetProduct(ctx context.Context, id string) (*pb.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return c.productClient.GetProduct(ctx, &pb.GetProductRequest{Id: id})
}

func (c *InventoryServiceClient) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return c.productClient.UpdateProduct(ctx, req)
}

func (c *InventoryServiceClient) DeleteProduct(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	_, err := c.productClient.DeleteProduct(ctx, &pb.DeleteProductRequest{Id: id})
	return err
}

func (c *InventoryServiceClient) ListProducts(ctx context.Context, filter *pb.ProductFilter, page, limit int32) (*pb.ListProductsResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return c.productClient.ListProducts(ctx, &pb.ListProductsRequest{
		Filter: filter,
		Page:   page,
		Limit:  limit,
	})
}

// Category service methods
func (c *InventoryServiceClient) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return c.categoryClient.CreateCategory(ctx, req)
}

func (c *InventoryServiceClient) GetCategory(ctx context.Context, id string) (*pb.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return c.categoryClient.GetCategory(ctx, &pb.GetCategoryRequest{Id: id})
}

func (c *InventoryServiceClient) UpdateCategory(ctx context.Context, req *pb.UpdateCategoryRequest) (*pb.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return c.categoryClient.UpdateCategory(ctx, req)
}

func (c *InventoryServiceClient) DeleteCategory(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	_, err := c.categoryClient.DeleteCategory(ctx, &pb.DeleteCategoryRequest{Id: id})
	return err
}

func (c *InventoryServiceClient) ListCategories(ctx context.Context) (*pb.ListCategoriesResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return c.categoryClient.ListCategories(ctx, &emptypb.Empty{})
}
