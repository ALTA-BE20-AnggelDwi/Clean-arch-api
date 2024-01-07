package handler

import (
	"clean-arch/features/product"
	"clean-arch/utils/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productService product.ProductServiceInterface
}

func NewProductHandler(productService product.ProductServiceInterface) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (handler *ProductHandler) GetAllProducts(c echo.Context) error {
	results, errSelect := handler.productService.GetAll()
	if errSelect != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error read data. "+errSelect.Error(), nil))
	}

	productsResult := CoreToResponseList(results)

	return c.JSON(http.StatusOK, responses.WebResponse("success read data.", productsResult))
}

func (handler *ProductHandler) CreateProduct(c echo.Context) error {
	newProduct := ProductRequest{}
	errBind := c.Bind(&newProduct)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	productCore := RequestToCore(newProduct)
	errInsert := handler.productService.Create(productCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert data"+errInsert.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success insert data", nil))
}

func (handler *ProductHandler) UpdateProduct(c echo.Context) error {
	id := c.Param("product_id")
	idParam, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error convert id param", nil))
	}

	updatedProduct := ProductRequest{}
	errBind := c.Bind(&updatedProduct)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	productCore := RequestToCore(updatedProduct)
	err := handler.productService.Update(idParam, productCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error update data"+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success update data", nil))
}

func (handler *ProductHandler) DeleteProduct(c echo.Context) error {
	id := c.Param("product_id")
	idParam, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error convert id param", nil))
	}

	err := handler.productService.Delete(idParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error delete data"+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success delete data", nil))
}

func (handler *ProductHandler) GetProductByID(c echo.Context) error {
	// Mendapatkan ID dari parameter URL
	idParam := c.Param("product_id")

	// Konversi string ID ke int
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("Invalid ID", nil))
	}

	// Memanggil fungsi logic untuk mendapatkan produk berdasarkan ID
	product, err := handler.productService.SelectByProductID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("Failed to get product", nil))
	}

	// Mengembalikan produk dalam respons JSON
	return c.JSON(http.StatusOK, responses.WebResponse("Success read data.", product))
}

func (handler *ProductHandler) GetProductsByUserID(c echo.Context) error {
	// Mendapatkan ID pengguna dari parameter URL
	userIDParam := c.Param("user_id")

	// Konversi string ID pengguna ke int
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("Invalid User ID", nil))
	}

	// Memanggil fungsi logic untuk mendapatkan produk berdasarkan ID pengguna
	products, err := handler.productService.SelectByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("Failed to get products by user ID", nil))
	}

	// Mengembalikan produk dalam respons JSON
	return c.JSON(http.StatusOK, responses.WebResponse("Success read data.", products))
}
