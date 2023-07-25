package repository

import (
	"context"

	"github.com/Akshayvij07/ecommerce/pkg/helper/respondse"
)

func (c *AdminDataBase) ViewSalesReport(ctx context.Context) ([]respondse.SalesReport, error) {
	var report []respondse.SalesReport
	FetchReports := `SELECT u.id, u.name, pm.payment_method AS payment_method, o.order_date, o.order_total,u.mobile, a.house_number,a.pincode           
	FROM orders o
	JOIN users u ON u.id = o.users_id
	JOIN payment_methods pm ON o.payment_method_id = pm.id
	JOIN user_addresses ua ON o.shipping_address_id = ua.id 
	JOIN addresses a ON ua.address_id = a.id 
	WHERE o.order_status_id = 1;`
	err := c.DB.Raw(FetchReports).Scan(&report).Error
	return report, err
}
