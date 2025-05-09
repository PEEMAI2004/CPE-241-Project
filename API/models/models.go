package models

// HarvestLog represents a log entry for a honey harvest
type HarvestLog struct {
	// HarvestID      int     `json:"harvest_id"`
	BeeHiveID      int     `json:"beehive_id"`
	HarvestDate    string  `json:"harvestdate"`
	Production     float64 `json:"production"`
	Unit           string  `json:"unit"`
	ProductionNote string  `json:"production_note"`
}

// HoneyStock represents the honey stock table
type HoneyStock struct {
	// StockID    int `json:"stock_id"`
	BeeHiveID  int     `json:"beehive_id"`
	Quantity   float64 `json:"quantity"`
	Unit       string  `json:"unit"`
	HarvestID  int     `json:"harvest_id"`
	StockDate  string  `json:"stock_date"`
	IsSold     bool    `json:"is_sold"`
}

// OrderRequest represents an incoming order request
type OrderRequest struct {
	UserID     int         `json:"user_id"`
	CustomerID int         `json:"customer_id"`
	Items      []OrderItem `json:"items"`
	OrderDate  string      `json:"order_date"`
	OrderStatus string     `json:"status"`
}

// OrderItem represents an item in an order request
type OrderItem struct {
	StockID int     `json:"stock_id"`
	Price   float64 `json:"price"`
}

// OrderList represents an entry in the orderlist table
type OrderList struct {
	// OrderID    int `json:"order_id"` // Will be set by database
	CustomerID int    `json:"customer_id"`
	UserID     int    `json:"user_id"`
	OrderDate  string `json:"order_date"`
	OrderStatus string     `json:"status"`
}

// OrderItemDB represents an entry in the orderitem table
type OrderItemDB struct {
	// OrderItemID int     `json:"orderitem_id"` // Will be set by database
	OrderID  int     `json:"order_id"`
	StockID  int     `json:"stock_id"`
	Price    float64 `json:"price"`
}