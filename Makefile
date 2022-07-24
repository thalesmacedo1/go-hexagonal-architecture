cmock:
	mockgen -destination=application/mocks/application.go -source=application/product.go
	
.PHONY: cmock