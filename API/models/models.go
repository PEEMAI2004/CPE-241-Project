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
	BeeHiveID  int `json:"beehive_id"`
	Quantity   float64    `json:"quantity"`
	Unit       string `json:"unit"`
	HarvestID  int `json:"harvest_id"`
	StockDate  string `json:"stock_date"`
}
