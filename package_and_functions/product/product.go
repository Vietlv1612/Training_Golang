package product

import "fmt"

type Product struct {
	Name  string
	Price float64
}

// AddProduct thêm một sản phẩm mới vào hệ thống
func AddProduct(p Product) {
	// Sử dụng defer để ghi nhật ký sau khi sản phẩm được thêm thành công
	defer func() {
		fmt.Printf("Product %s added successfully\n", p.Name)
	}()

	// Kiểm tra tính hợp lệ của sản phẩm
	if p.Price < 0 {
		// Nếu giá sản phẩm là âm, gây ra panic
		panic(fmt.Sprintf("Invalid price %f for product %s", p.Price, p.Name))
	}

	// Nếu sản phẩm hợp lệ, xử lý thêm sản phẩm vào hệ thống (giả lập)
	fmt.Printf("Adding product: %s with price: %.2f\n", p.Name, p.Price)
}

// HandleAddProduct là hàm xử lý việc thêm sản phẩm với việc bắt lỗi panic
func HandleAddProduct(p Product) {
	// Sử dụng recover để bắt và xử lý lỗi panic nếu có
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Error adding product: %v\n", r)
		}
	}()

	// Gọi hàm AddProduct
	AddProduct(p)
}
