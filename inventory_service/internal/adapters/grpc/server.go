package grpc

import (
	"context"
	"inventory_service/internal/application"
	"inventory_service/internal/domain"
	pb "inventory_service/proto/inventory"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Server implements the gRPC server for inventory service
type Server struct {
	pb.UnimplementedProductServiceServer
	pb.UnimplementedCategoryServiceServer
	productService  *application.ProductService
	categoryService *application.CategoryService
}

// NewServer creates a new gRPC server instance
func NewServer(productService *application.ProductService, categoryService *application.CategoryService) *Server {
	return &Server{
		productService:  productService,
		categoryService: categoryService,
	}
}

// Domain to protobuf conversion helpers
func domainProductToProto(p *domain.Product) *pb.Product {
	return &pb.Product{
		Id:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		CategoryId:  p.CategoryID,
		Price:       p.Price,
		Stock:       int32(p.Stock),
		CreatedAt:   timestamppb.New(p.CreatedAt),
		UpdatedAt:   timestamppb.New(p.UpdatedAt),
	}
}

func domainCategoryToProto(c *domain.Category) *pb.Category {
	return &pb.Category{
		Id:          c.ID,
		Name:        c.Name,
		Description: c.Description,
	}
}

func protoFilterToDomain(filter *pb.ProductFilter) domain.ProductFilter {
	return domain.ProductFilter{
		Name:       filter.Name,
		CategoryID: filter.CategoryId,
		MinPrice:   filter.MinPrice,
		MaxPrice:   filter.MaxPrice,
	}
}

// Product Service implementations
func (s *Server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.Product, error) {
	// Check if the category exists
	_, err := s.categoryService.GetCategory(req.CategoryId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Category with ID %s not found: %v", req.CategoryId, err)
	}

	product := &domain.Product{
		Name:        req.Name,
		Description: req.Description,
		CategoryID:  req.CategoryId,
		Price:       req.Price,
		Stock:       int(req.Stock),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.productService.CreateProduct(product); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create product: %v", err)
	}

	return domainProductToProto(product), nil
}

func (s *Server) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.Product, error) {
	product, err := s.productService.GetProduct(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Product not found: %v", err)
	}

	return domainProductToProto(product), nil
}

func (s *Server) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.Product, error) {
	product, err := s.productService.GetProduct(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Product not found: %v", err)
	}

	// Check if the category exists
	_, err = s.categoryService.GetCategory(req.CategoryId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Category with ID %s not found: %v", req.CategoryId, err)
	}

	product.Name = req.Name
	product.Description = req.Description
	product.CategoryID = req.CategoryId
	product.Price = req.Price
	product.Stock = int(req.Stock)
	product.UpdatedAt = time.Now()

	if err := s.productService.UpdateProduct(product); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update product: %v", err)
	}

	return domainProductToProto(product), nil
}

func (s *Server) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*emptypb.Empty, error) {
	if err := s.productService.DeleteProduct(req.Id); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to delete product: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	filter := protoFilterToDomain(req.Filter)
	products, err := s.productService.ListProducts(filter, int(req.Page), int(req.Limit))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to list products: %v", err)
	}

	response := &pb.ListProductsResponse{}
	for _, product := range products {
		response.Products = append(response.Products, domainProductToProto(product))
	}

	return response, nil
}

// Category Service implementations
func (s *Server) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.Category, error) {
	category := &domain.Category{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := s.categoryService.CreateCategory(category); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create category: %v", err)
	}

	return domainCategoryToProto(category), nil
}

func (s *Server) GetCategory(ctx context.Context, req *pb.GetCategoryRequest) (*pb.Category, error) {
	category, err := s.categoryService.GetCategory(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Category not found: %v", err)
	}

	return domainCategoryToProto(category), nil
}

func (s *Server) UpdateCategory(ctx context.Context, req *pb.UpdateCategoryRequest) (*pb.Category, error) {
	category, err := s.categoryService.GetCategory(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Category not found: %v", err)
	}

	category.Name = req.Name
	category.Description = req.Description

	if err := s.categoryService.UpdateCategory(category); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update category: %v", err)
	}

	return domainCategoryToProto(category), nil
}

func (s *Server) DeleteCategory(ctx context.Context, req *pb.DeleteCategoryRequest) (*emptypb.Empty, error) {
	if err := s.categoryService.DeleteCategory(req.Id); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to delete category: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) ListCategories(ctx context.Context, _ *emptypb.Empty) (*pb.ListCategoriesResponse, error) {
	categories, err := s.categoryService.ListCategories()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to list categories: %v", err)
	}

	response := &pb.ListCategoriesResponse{}
	for _, category := range categories {
		response.Categories = append(response.Categories, domainCategoryToProto(category))
	}

	return response, nil
}
