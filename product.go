package shopify

import (
	"context"
	"fmt"
	"strings"

	"github.com/r0busta/go-shopify-graphql-model/v4/graph/model"
)

//go:generate mockgen -destination=./mock/product_service.go -package=mock . ProductService
type ProductService interface {
	Query(ctx context.Context, q string, vars map[string]any) (*model.Product, error)
	GetAllProducts(ctx context.Context, fields string, filter string) ([]model.Product, error)
	StreamProducts(ctx context.Context, fields string, filter string, limit int) (<-chan model.Product, <-chan error)
	BulkQuery(ctx context.Context, q string) ([]model.Product, error)
	List(ctx context.Context, query string) ([]model.Product, error)
	ListAll(ctx context.Context) ([]model.Product, error)

	Get(ctx context.Context, id string) (*model.Product, error)

	Create(ctx context.Context, product model.ProductCreateInput, media []model.CreateMediaInput) (*string, error)

	Update(ctx context.Context, product model.ProductUpdateInput, media []model.CreateMediaInput) error

	Delete(ctx context.Context, product model.ProductDeleteInput) error

	VariantsBulkCreate(ctx context.Context, id string, input []model.ProductVariantsBulkInput, strategy model.ProductVariantsBulkCreateStrategy) error
	VariantsBulkUpdate(ctx context.Context, id string, input []model.ProductVariantsBulkInput) error
	VariantsBulkReorder(ctx context.Context, id string, input []model.ProductVariantPositionInput) error

	MediaCreate(ctx context.Context, id string, input []model.CreateMediaInput) error
}

type ProductServiceOp struct {
	client *Client
}

var _ ProductService = &ProductServiceOp{}

type mutationProductCreate struct {
	ProductCreateResult struct {
		Product *struct {
			ID string `json:"id,omitempty"`
		} `json:"product,omitempty"`

		UserErrors []model.UserError `json:"userErrors,omitempty"`
	} `graphql:"productCreate(product: $product, media: $media)" json:"productCreate"`
}

type mutationProductUpdate struct {
	ProductUpdateResult struct {
		UserErrors []model.UserError `json:"userErrors,omitempty"`
	} `graphql:"productUpdate(product: $product, media: $media)" json:"productUpdate"`
}

type mutationProductDelete struct {
	ProductDeleteResult struct {
		UserErrors []model.UserError `json:"userErrors,omitempty"`
	} `graphql:"productDelete(input: $input)" json:"productDelete"`
}

type mutationProductVariantsBulkCreate struct {
	ProductVariantsBulkCreateResult struct {
		UserErrors []model.UserError `json:"userErrors,omitempty"`
	} `graphql:"productVariantsBulkCreate(productId: $productId, variants: $variants, strategy: $strategy)" json:"productVariantsBulkCreate"`
}

type mutationProductVariantsBulkUpdate struct {
	ProductVariantsBulkUpdateResult struct {
		UserErrors []model.UserError `json:"userErrors,omitempty"`
	} `graphql:"productVariantsBulkUpdate(productId: $productId, variants: $variants)" json:"productVariantsBulkUpdate"`
}

type mutationProductVariantsBulkReorder struct {
	ProductVariantsBulkReorderResult struct {
		UserErrors []model.UserError `json:"userErrors,omitempty"`
	} `graphql:"productVariantsBulkReorder(positions: $positions, productId: $productId)" json:"productVariantsBulkReorder"`
}

type mutationProductCreateMedia struct {
	ProductCreateMediaResult struct {
		MediaUserErrors []model.UserError `json:"mediaUserErrors,omitempty"`
	} `graphql:"productCreateMedia(productId: $productId, media: $media)" json:"productCreateMedia"`
}

const productBaseQuery = `
	id
	legacyResourceId
	handle
	options{
		id
		name
		values
		position
	}
	tags
	title
	description
	descriptionPlainSummary
	priceRangeV2{
		minVariantPrice{
			amount
			currencyCode
		}
		maxVariantPrice{
			amount
			currencyCode
		}
	}
	productType
	vendor
	totalInventory
	onlineStoreUrl	
	descriptionHtml
	seo{
		description
		title
	}
	templateSuffix
	customProductType
	featuredImage{
		id
		altText
		height
		width
		url
	}
`

var productQuery = fmt.Sprintf(`
	%s
	variants(first:100, after: $cursor){
		edges{
			node{
				id
				legacyResourceId
				title
				displayName
				sku
				selectedOptions{
					name
					value
					optionValue{
						id
						name
					}
				}
				position
				image {
					id
					altText
					height
					width
					url
				}
				compareAtPrice
				price
				inventoryQuantity
				inventoryItem{
					id
					legacyResourceId	
					sku						
				}
				availableForSale
			}
		}
		pageInfo{
			hasNextPage
		}
	}
`, productBaseQuery)

var productBulkQuery = fmt.Sprintf(`
	%s
	metafields{
		edges{
			node{
				id
				legacyResourceId
				namespace
				key
				value
				type
			}
		}
	}
	variants{
		edges{
			node{
				id
				legacyResourceId
				title
				displayName
				sku
				selectedOptions{
					name
					value
					optionValue{
						id
						name
					}
				}
				position
				image {
					id
					altText
					height
					width
					url
				}
				compareAtPrice
				price
				inventoryQuantity
				inventoryItem{
					id
					legacyResourceId
					sku							
				}
				availableForSale
			}
		}
	}
`, productBaseQuery)

func (s *ProductServiceOp) BulkQuery(ctx context.Context, query string) ([]model.Product, error) {
	res := []model.Product{}

	if err := s.client.BulkOperation.BulkQuery(ctx, query, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ProductServiceOp) Query(ctx context.Context, query string, vars map[string]any) (*model.Product, error) {
	out := struct {
		Product *model.Product `json:"product"`
	}{}

	err := s.client.gql.QueryString(ctx, query, vars, &out)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	return out.Product, nil
}

func (s *ProductServiceOp) ListAll(ctx context.Context) ([]model.Product, error) {
	q := fmt.Sprintf(`
		{
			products{
				edges{
					node{
						%s
					}
				}
			}
		}
	`, productBulkQuery)

	return s.BulkQuery(ctx, q)
}

func (s *ProductServiceOp) List(ctx context.Context, query string) ([]model.Product, error) {
	q := fmt.Sprintf(`
		{
			products(query: "$query"){
				edges{
					node{
						%s
					}
				}
			}
		}
	`, productBulkQuery)

	q = strings.ReplaceAll(q, "$query", query)

	return s.BulkQuery(ctx, q)
}

func (s *ProductServiceOp) Get(ctx context.Context, id string) (*model.Product, error) {
	out, err := s.getPage(ctx, id, "")
	if err != nil {
		return nil, err
	}

	nextPageData := out
	hasNextPage := out.Variants.PageInfo.HasNextPage
	for hasNextPage && len(nextPageData.Variants.Edges) > 0 {
		cursor := nextPageData.Variants.Edges[len(nextPageData.Variants.Edges)-1].Cursor
		nextPageData, err := s.getPage(ctx, id, cursor)
		if err != nil {
			return nil, fmt.Errorf("get page: %w", err)
		}
		out.Variants.Edges = append(out.Variants.Edges, nextPageData.Variants.Edges...)
		hasNextPage = nextPageData.Variants.PageInfo.HasNextPage
	}

	return out, nil
}

func (s *ProductServiceOp) getPage(ctx context.Context, id string, cursor string) (*model.Product, error) {
	q := fmt.Sprintf(`
		query product($id: ID!, $cursor: String) {
			product(id: $id){
				%s
			}
		}
	`, productQuery)

	vars := map[string]interface{}{
		"id": id,
	}
	if cursor != "" {
		vars["cursor"] = cursor
	}

	out := struct {
		Product *model.Product `json:"product"`
	}{}
	err := s.client.gql.QueryString(ctx, q, vars, &out)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	return out.Product, nil
}

func (s *ProductServiceOp) Create(ctx context.Context, product model.ProductCreateInput, media []model.CreateMediaInput) (*string, error) {
	m := mutationProductCreate{}

	vars := map[string]interface{}{
		"product": product,
		"media":   media,
	}

	err := s.client.gql.Mutate(ctx, &m, vars)
	if err != nil {
		return nil, fmt.Errorf("mutation: %w", err)
	}

	if len(m.ProductCreateResult.UserErrors) > 0 {
		return nil, fmt.Errorf("%+v", m.ProductCreateResult.UserErrors)
	}

	return &m.ProductCreateResult.Product.ID, nil
}

func (s *ProductServiceOp) Update(ctx context.Context, product model.ProductUpdateInput, media []model.CreateMediaInput) error {
	m := mutationProductUpdate{}

	vars := map[string]interface{}{
		"product": product,
		"media":   media,
	}
	err := s.client.gql.Mutate(ctx, &m, vars)
	if err != nil {
		return fmt.Errorf("mutation: %w", err)
	}

	if len(m.ProductUpdateResult.UserErrors) > 0 {
		return fmt.Errorf("%+v", m.ProductUpdateResult.UserErrors)
	}

	return nil
}

func (s *ProductServiceOp) Delete(ctx context.Context, product model.ProductDeleteInput) error {
	m := mutationProductDelete{}

	vars := map[string]interface{}{
		"input": product,
	}
	err := s.client.gql.Mutate(ctx, &m, vars)
	if err != nil {
		return fmt.Errorf("mutation: %w", err)
	}

	if len(m.ProductDeleteResult.UserErrors) > 0 {
		return fmt.Errorf("%+v", m.ProductDeleteResult.UserErrors)
	}

	return nil
}

func (s *ProductServiceOp) VariantsBulkCreate(ctx context.Context, id string, input []model.ProductVariantsBulkInput, strategy model.ProductVariantsBulkCreateStrategy) error {
	m := mutationProductVariantsBulkCreate{}

	vars := map[string]interface{}{
		"productId": id,
		"variants":  input,
		"strategy":  strategy,
	}
	err := s.client.gql.Mutate(ctx, &m, vars)
	if err != nil {
		return fmt.Errorf("mutation: %w", err)
	}

	if len(m.ProductVariantsBulkCreateResult.UserErrors) > 0 {
		return fmt.Errorf("%+v", m.ProductVariantsBulkCreateResult.UserErrors)
	}

	return nil
}

func (s *ProductServiceOp) VariantsBulkUpdate(ctx context.Context, id string, input []model.ProductVariantsBulkInput) error {
	m := mutationProductVariantsBulkUpdate{}

	vars := map[string]interface{}{
		"productId": id,
		"variants":  input,
	}
	err := s.client.gql.Mutate(ctx, &m, vars)
	if err != nil {
		return fmt.Errorf("mutation: %w", err)
	}

	if len(m.ProductVariantsBulkUpdateResult.UserErrors) > 0 {
		return fmt.Errorf("%+v", m.ProductVariantsBulkUpdateResult.UserErrors)
	}

	return nil
}

func (s *ProductServiceOp) VariantsBulkReorder(ctx context.Context, id string, input []model.ProductVariantPositionInput) error {
	m := mutationProductVariantsBulkReorder{}

	vars := map[string]interface{}{
		"productId": id,
		"positions": input,
	}
	err := s.client.gql.Mutate(ctx, &m, vars)
	if err != nil {
		return fmt.Errorf("mutation: %w", err)
	}

	if len(m.ProductVariantsBulkReorderResult.UserErrors) > 0 {
		return fmt.Errorf("%+v", m.ProductVariantsBulkReorderResult.UserErrors)
	}

	return nil
}

func (s *ProductServiceOp) MediaCreate(ctx context.Context, id string, input []model.CreateMediaInput) error {
	m := mutationProductCreateMedia{}

	vars := map[string]interface{}{
		"productId": id,
		"media":     input,
	}

	err := s.client.gql.Mutate(ctx, &m, vars)
	if err != nil {
		return fmt.Errorf("mutation: %w", err)
	}

	if len(m.ProductCreateMediaResult.MediaUserErrors) > 0 {
		return fmt.Errorf("%+v", m.ProductCreateMediaResult.MediaUserErrors)
	}

	return nil
}

// getProductsQuery returns a reusable GraphQL query for fetching products
func getProductsQuery(fields string) string {
	return fmt.Sprintf(`
		query GetProducts($cursor: String, $pageSize: Int!, $filter: String!) {
			products(first: $pageSize, after: $cursor, query: $filter) {
				edges {
					node {
						%s
					}
					cursor
				}
				pageInfo {
					hasNextPage
				}
			}
		}
	`, fields)
}

func (s *ProductServiceOp) GetAllProducts(ctx context.Context, fields string, filter string) ([]model.Product, error) {
	// Initialize result slice
	var allProducts []model.Product

	// Use the shared query
	query := getProductsQuery(fields)

	// Default page size for batch fetching
	pageSize := 250

	// All parameters are passed as variables
	vars := map[string]any{
		"cursor":   nil,
		"pageSize": pageSize,
		"filter":   filter,
	}

	// Keep track of whether there are more pages
	hasNextPage := true

	// Loop until we've fetched all pages
	for hasNextPage {
		out := struct {
			Products model.ProductConnection `json:"products"`
		}{}

		// Execute the query with current variables
		err := s.client.gql.QueryString(ctx, query, vars, &out)
		if err != nil {
			return nil, fmt.Errorf("query: %w", err)
		}

		// Extract products from the current page
		for _, edge := range out.Products.Edges {
			allProducts = append(allProducts, *edge.Node)
		}

		// Check if there are more pages
		hasNextPage = out.Products.PageInfo.HasNextPage

		// If there are more pages, update the cursor for the next query
		if hasNextPage && len(out.Products.Edges) > 0 {
			vars["cursor"] = out.Products.Edges[len(out.Products.Edges)-1].Cursor
		} else if hasNextPage {
			// If we can't get a next cursor but hasNextPage is true,
			// we should break to avoid an infinite loop
			return nil, fmt.Errorf("pagination error: hasNextPage is true but no cursor found")
		}
	}

	return allProducts, nil
}

func (s *ProductServiceOp) StreamProducts(ctx context.Context, fields string, filter string, limit int) (<-chan model.Product, <-chan error) {
	productChan := make(chan model.Product)
	errorChan := make(chan error)

	go func() {
		defer close(productChan)
		defer close(errorChan)

		// If limit is 0 or negative, use default page size of 250
		pageSize := 250
		if limit > 0 && limit < pageSize {
			pageSize = limit
		}

		// Use the shared query with variables
		query := getProductsQuery(fields)

		// All parameters are passed as variables
		vars := map[string]any{
			"cursor":   nil,
			"pageSize": pageSize,
			"filter":   filter,
		}

		// Keep track of whether there are more pages and how many products we've processed
		hasNextPage := true
		productsProcessed := 0

		// Loop until we've fetched all pages or reached the limit
		for hasNextPage && (limit <= 0 || productsProcessed < limit) {
			out := struct {
				Products model.ProductConnection `json:"products"`
			}{}

			// Execute the query with current variables
			err := s.client.gql.QueryString(ctx, query, vars, &out)
			if err != nil {
				errorChan <- fmt.Errorf("query: %w", err)
				return
			}

			// Stream products from the current page
			for _, edge := range out.Products.Edges {
				select {
				case productChan <- *edge.Node:
					productsProcessed++
					// If we've reached the limit, stop processing
					if limit > 0 && productsProcessed >= limit {
						return
					}
				case <-ctx.Done():
					return
				}
			}

			// Check if there are more pages
			hasNextPage = out.Products.PageInfo.HasNextPage

			// If there are more pages, update the cursor for the next query
			if hasNextPage && len(out.Products.Edges) > 0 {
				vars["cursor"] = out.Products.Edges[len(out.Products.Edges)-1].Cursor
			} else if hasNextPage {
				// If we can't get a next cursor but hasNextPage is true,
				// we should break to avoid an infinite loop
				errorChan <- fmt.Errorf("pagination error: hasNextPage is true but no cursor found")
				return
			}
		}
	}()

	return productChan, errorChan
}
