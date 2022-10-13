cmock:
	mockgen -destination=application/mocks/application.go -source=application/product.go application
	
.PHONY: cmock