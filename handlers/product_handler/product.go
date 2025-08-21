package product_handler

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/contract_resp"
	"github.com/umarkotak/ytkidd-api/repos/product_repo"
	"github.com/umarkotak/ytkidd-api/utils/render"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	products, err := product_repo.GetAll(ctx)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		render.Error(w, r, err, "")
		return
	}

	publicProducts := make([]contract_resp.PublicProduct, 0, len(products))
	for _, product := range products {
		publicProducts = append(publicProducts, contract_resp.PublicProduct{
			Code:        product.Code,
			BenefitType: product.BenefitType,
			Name:        product.Name,
			ImageUrl:    product.ImageUrl,
			BasePrice:   product.BasePrice,
			Price:       product.Price,
			Metadata:    product.Metadata,
		})
	}

	render.Response(w, r, http.StatusOK, publicProducts)
}
