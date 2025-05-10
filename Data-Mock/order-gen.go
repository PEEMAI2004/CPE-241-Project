package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

// StockItem represents a single stock item from the JSON file
type StockItem struct {
	StockID   int       `json:"stock_id"`
	StockDate string    `json:"stock_date"`
}

// ParsedStockItem includes the properly parsed time
type ParsedStockItem struct {
	StockID   int
	StockDate time.Time
	Price     int // Added Price field
}

// StockGroup represents a group of stocks with similar dates
type StockGroup struct {
	GroupID     int               `json:"group_id"`
	UserID      int               `json:"user_id"`      // Added UserID field
	CustomerID  int               `json:"customer_id"`  // Added CustomerID field
	Items       []ParsedStockItem `json:"items"`
	OrderDate   string            `json:"order_date"`   // Added OrderDate field
	Status      string            `json:"status"`       // Added Status field
}

// GroupedStocks represents the collection of all stock groups
type GroupedStocks struct {
	Groups []StockGroup `json:"groups"`
}

// LoadBatchFile loads stock data from a specific batch file
func LoadBatchFile(n int) ([]ParsedStockItem, error) {
	// Construct the file path
	batchPath := filepath.Join("unsold_stocks_batches", fmt.Sprintf("batch_%d.json", n))
	
	// Open the file
	file, err := os.Open(batchPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open batch file %d: %w", n, err)
	}
	defer file.Close()
	
	// Read the file contents
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read batch file %d: %w", n, err)
	}
	
	// Parse JSON into slice of StockItem
	var stockItems []StockItem
	if err := json.Unmarshal(data, &stockItems); err != nil {
		return nil, fmt.Errorf("failed to parse batch file %d: %w", n, err)
	}
	
	// Convert to ParsedStockItem with properly parsed time
	parsedItems := make([]ParsedStockItem, 0, len(stockItems))
	for _, item := range stockItems {
		parsedTime, err := time.Parse("2006-01-02T15:04:05", item.StockDate)
		if err != nil {
			return nil, fmt.Errorf("failed to parse date '%s' in batch file %d: %w", item.StockDate, n, err)
		}
		
		// Generate random price between 10000 and 30000
		randomPrice := 10000 + rand.Intn(20001) // 10000 + random value up to 20000
		
		parsedItems = append(parsedItems, ParsedStockItem{
			StockID:   item.StockID,
			StockDate: parsedTime,
			Price:     randomPrice,
		})
	}
	
	return parsedItems, nil
}

// LoadUserIDs loads user IDs with role 4 from the JSON file
func LoadUserIDs() ([]int, error) {
	// Open the file
	file, err := os.Open("user_ids_role_4.json")
	if err != nil {
		return nil, fmt.Errorf("failed to open user_ids_role_4.json: %w", err)
	}
	defer file.Close()
	
	// Read the file contents
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read user_ids_role_4.json: %w", err)
	}
	
	// Parse JSON into slice of integers
	var userIDs []int
	if err := json.Unmarshal(data, &userIDs); err != nil {
		return nil, fmt.Errorf("failed to parse user_ids_role_4.json: %w", err)
	}
	
	return userIDs, nil
}

// LoadCustomerIDs loads customer IDs from the JSON file
func LoadCustomerIDs() ([]int, error) {
	// Open the file
	file, err := os.Open("customer_ids.json")
	if err != nil {
		return nil, fmt.Errorf("failed to open customer_ids.json: %w", err)
	}
	defer file.Close()
	
	// Read the file contents
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read customer_ids.json: %w", err)
	}
	
	// Parse JSON into slice of integers
	var customerIDs []int
	if err := json.Unmarshal(data, &customerIDs); err != nil {
		return nil, fmt.Errorf("failed to parse customer_ids.json: %w", err)
	}
	
	return customerIDs, nil
}

// SoldStockID represents a sold stock ID from the API
type SoldStockID struct {
	StockID int `json:"stock_id"`
}

// LoadSoldStockIDs loads stock IDs that are already sold
func LoadSoldStockIDs() (map[int]bool, error) {
	soldStockMap := make(map[int]bool)
	cacheFile := "sold_stock_ids_cache.json"
	
	// Try to load from cache first for better performance
	if data, err := ioutil.ReadFile(cacheFile); err == nil {
		var soldStockIDs []int
		if err := json.Unmarshal(data, &soldStockIDs); err == nil {
			fmt.Printf("Loaded %d sold stock IDs from cache\n", len(soldStockIDs))
			for _, id := range soldStockIDs {
				soldStockMap[id] = true
			}
			return soldStockMap, nil
		}
		// If cache loading fails, continue to fetch from API
		fmt.Printf("Warning: Failed to load sold stock IDs from cache: %v\n", err)
	}
	
	// If cache is not available, fetch from the API
	fmt.Println("Fetching sold stock IDs from API...")
	
	baseURL := "https://postgrest.kaminjitt.com/honeystock"
	offset := 0
	limit := 10000 // Fetch in batches to avoid timeout or memory issues
	
	for {
		url := fmt.Sprintf("%s?is_sold=eq.true&select=stock_id&order=stock_date.asc&limit=%d&offset=%d", 
			baseURL, limit, offset)
		
		resp, err := http.Get(url)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch sold stock IDs: %w", err)
		}
		defer resp.Body.Close()
		
		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("API returned error status: %s", resp.Status)
		}
		
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read API response: %w", err)
		}
		
		var soldStockIDs []SoldStockID
		if err := json.Unmarshal(body, &soldStockIDs); err != nil {
			return nil, fmt.Errorf("failed to parse API response: %w", err)
		}
		
		if len(soldStockIDs) == 0 {
			break // No more results
		}
		
		// Add the sold stock IDs to the map
		for _, item := range soldStockIDs {
			soldStockMap[item.StockID] = true
		}
		
		fmt.Printf("Fetched %d sold stock IDs (total so far: %d)\n", 
			len(soldStockIDs), len(soldStockMap))
		
		// Prepare for the next batch
		offset += limit
		
		// Be kind to the server
		time.Sleep(500 * time.Millisecond)
	}
	
	// Cache the results for future runs
	allIDs := make([]int, 0, len(soldStockMap))
	for id := range soldStockMap {
		allIDs = append(allIDs, id)
	}
	
	if data, err := json.Marshal(allIDs); err == nil {
		if err := ioutil.WriteFile(cacheFile, data, 0644); err == nil {
			fmt.Printf("Cached %d sold stock IDs for future use\n", len(soldStockMap))
		} else {
			fmt.Printf("Warning: Failed to write sold stock ID cache: %v\n", err)
		}
	}
	
	fmt.Printf("Total sold stock IDs: %d\n", len(soldStockMap))
	return soldStockMap, nil
}

// FilterOutSoldStocks removes stocks that are already sold
func FilterOutSoldStocks(stocks []ParsedStockItem, soldStockMap map[int]bool) []ParsedStockItem {
	if len(soldStockMap) == 0 {
		return stocks // Nothing to filter
	}
	
	result := make([]ParsedStockItem, 0, len(stocks))
	for _, stock := range stocks {
		if !soldStockMap[stock.StockID] {
			result = append(result, stock)
		}
	}
	
	fmt.Printf("Filtered out %d sold stocks, %d unsold stocks remaining\n", 
		len(stocks) - len(result), len(result))
	
	return result
}

// GroupStocksByDate groups the stocks so that each group has between 1000-10000 items
// and the date difference within each group is no more than 100 days
func GroupStocksByDate(stocks []ParsedStockItem) GroupedStocks {
	// Shuffle the stocks to randomize before grouping
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(stocks), func(i, j int) {
		stocks[i], stocks[j] = stocks[j], stocks[i]
	})
	
	var result GroupedStocks
	remainingStocks := make([]ParsedStockItem, len(stocks))
	copy(remainingStocks, stocks)
	
	groupID := 1
	
	for len(remainingStocks) > 0 {
		// Determine the size for this group
		// But don't exceed the number of remaining stocks
		maxSize := 2000
		minSize := 100
		
		if len(remainingStocks) < minSize {
			// If we have fewer than minSize stocks left, put them all in one group
			groupSize := len(remainingStocks)
			group := createGroup(remainingStocks[:groupSize], groupID)
			result.Groups = append(result.Groups, group)
			break
		}
		
		groupSize := minSize + rand.Intn(maxSize-minSize+1)
		if groupSize > len(remainingStocks) {
			groupSize = len(remainingStocks)
		}
		
		// Try to create a group that satisfies the date constraint
		var validGroup StockGroup
		var validGroupSize int
		
		// Sort the candidate items by date
		candidateItems := make([]ParsedStockItem, groupSize)
		copy(candidateItems, remainingStocks[:groupSize])
		
		// Try different subsets until we find one that satisfies the constraint
		for attempts := 0; attempts < 10; attempts++ {
			// Shuffle the candidate items to try a different combination
			rand.Shuffle(len(candidateItems), func(i, j int) {
				candidateItems[i], candidateItems[j] = candidateItems[j], candidateItems[i]
			})
			
			// Find earliest and latest dates
			var earliestDate, latestDate time.Time
			for i, item := range candidateItems {
				if i == 0 || item.StockDate.Before(earliestDate) {
					earliestDate = item.StockDate
				}
				if i == 0 || item.StockDate.After(latestDate) {
					latestDate = item.StockDate
				}
			}
			
			// Check if the date difference is within 100 days
			daysDiff := latestDate.Sub(earliestDate).Hours() / 24
			if daysDiff <= 100 {
				validGroup = createGroup(candidateItems, groupID)
				validGroupSize = len(candidateItems)
				break
			}
		}
		
		// If no valid group was found, create a smaller group
		if validGroup.GroupID == 0 {
			// Take a smaller subset that's more likely to fit within 100 days
			smallerSize := minSize
			smallerGroup := createGroup(remainingStocks[:smallerSize], groupID)
			result.Groups = append(result.Groups, smallerGroup)
			validGroupSize = smallerSize
		} else {
			result.Groups = append(result.Groups, validGroup)
		}
		
		// Remove the used items from remainingStocks
		remainingStocks = remainingStocks[validGroupSize:]
		groupID++
	}
	
	return result
}

// createGroup creates a stock group with the given items and ID
func createGroup(items []ParsedStockItem, groupID int) StockGroup {
	return StockGroup{
		GroupID: groupID,
		Items:   items,
		// UserID and CustomerID will be assigned later when userIDs and customerIDs are loaded
	}
}

// SaveGroupedStocks saves the grouped stocks to a JSON file
func SaveGroupedStocks(groupedStocks GroupedStocks, filename string) error {
	// Convert back to a format suitable for JSON
	type JsonStockItem struct {
		StockID   int    `json:"stock_id"`
		StockDate string `json:"stock_date"`
		Price     int    `json:"price"`
	}
	
	type JsonStockGroup struct {
		GroupID     int            `json:"group_id"`
		Items       []JsonStockItem `json:"items"`
		UserID      int            `json:"user_id"`
		CustomerID  int            `json:"customer_id"`
		OrderDate   string         `json:"order_date"`
		Status      string         `json:"status"`
	}
	
	type JsonGroupedStocks struct {
		Groups []JsonStockGroup `json:"groups"`
	}
	
	jsonGrouped := JsonGroupedStocks{
		Groups: make([]JsonStockGroup, len(groupedStocks.Groups)),
	}
	
	for i, group := range groupedStocks.Groups {
		jsonGroup := JsonStockGroup{
			GroupID:     group.GroupID,
			Items:       make([]JsonStockItem, len(group.Items)),
			UserID:      group.UserID,
			CustomerID:  group.CustomerID,
			OrderDate:   group.OrderDate,
			Status:      group.Status,
		}
		
		for j, item := range group.Items {
			jsonGroup.Items[j] = JsonStockItem{
				StockID:   item.StockID,
				StockDate: item.StockDate.Format("2006-01-02T15:04:05"),
				Price:     item.Price,  // Include the price in JSON output
			}
		}
		
		jsonGrouped.Groups[i] = jsonGroup
	}
	
	// Marshal to JSON
	data, err := json.MarshalIndent(jsonGrouped, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal grouped stocks to JSON: %w", err)
	}
	
	// Write to file
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write grouped stocks to file: %w", err)
	}
	
	return nil
}

// PostOrderGroup sends a single order group to the API
func PostOrderGroup(group StockGroup) (bool, error) {
	// Prepare the order data for the API
	type ApiStockItem struct {
		StockID   int    `json:"stock_id"`
		StockDate string `json:"stock_date"`
		Price     int    `json:"price"`
	}
	
	type ApiOrder struct {
		GroupID    int           `json:"group_id"`
		UserID     int           `json:"user_id"`
		CustomerID int           `json:"customer_id"`
		Items      []ApiStockItem `json:"items"`
		OrderDate  string        `json:"order_date"`
		Status     string        `json:"status"`
	}
	
	apiOrder := ApiOrder{
		GroupID:    group.GroupID,
		UserID:     group.UserID,
		CustomerID: group.CustomerID,
		Items:      make([]ApiStockItem, len(group.Items)),
		OrderDate:  group.OrderDate,
		Status:     group.Status,
	}
	
	for i, item := range group.Items {
		apiOrder.Items[i] = ApiStockItem{
			StockID:   item.StockID,
			StockDate: item.StockDate.Format("2006-01-02T15:04:05"),
			Price:     item.Price,
		}
	}
	
	// Convert to JSON
	jsonData, err := json.Marshal(apiOrder)
	if err != nil {
		return false, fmt.Errorf("failed to marshal order data: %w", err)
	}
	
	// Send the request
	url := "https://api.kaminjitt.com/api/order"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return false, fmt.Errorf("failed to post order to API: %w", err)
	}
	defer resp.Body.Close()
	
	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Errorf("failed to read API response: %w", err)
	}
	
	// Check for already sold pattern
	alreadySoldPattern := regexp.MustCompile(`is already sold`)
	if alreadySoldPattern.Match(body) {
		fmt.Printf("Order group %d contains already sold items, skipping\n", group.GroupID)
		return false, nil
	}
	
	// Check status code
	if resp.StatusCode != http.StatusCreated {
		return false, fmt.Errorf("API returned status %d: %s", resp.StatusCode, body)
	}
	
	return true, nil
}

func main() {
	// Make sure we seed the random number generator
	rand.Seed(time.Now().UnixNano())
	
	// Load sold stock IDs to filter them out
	soldStockMap, err := LoadSoldStockIDs()
	if err != nil {
		fmt.Printf("Error loading sold stock IDs: %v\n", err)
		fmt.Println("Continuing without filtering sold stocks...")
		soldStockMap = make(map[int]bool)
	}
	
	// Load user IDs and customer IDs first so we can assign them to groups
	userIDs, err := LoadUserIDs()
	if err != nil {
		fmt.Printf("Error loading user IDs: %v\n", err)
		return
	}
	
	customerIDs, err := LoadCustomerIDs()
	if err != nil {
		fmt.Printf("Error loading customer IDs: %v\n", err)
		return
	}
	
	// Print the loaded user IDs and customer IDs
	fmt.Printf("Loaded %d user IDs\n", len(userIDs))
	if len(userIDs) > 0 {
		fmt.Printf("First few user IDs: %v\n", userIDs[:min(5, len(userIDs))])
	}
	
	fmt.Printf("Loaded %d customer IDs\n", len(customerIDs))
	if len(customerIDs) > 0 {
		fmt.Printf("First few customer IDs: %v\n", customerIDs[:min(5, len(customerIDs))])
	}
	
	// Process all 115 batch files
	numBatches := 100 // Temporarily set to 2 for testing
	
	// Create output directory if it doesn't exist
	outputDir := "unsold_stocks_grouped_batches"
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Printf("Error creating output directory: %v\n", err)
		return
	}
	
	for i := 1; i <= numBatches; i++ {
		fmt.Printf("Processing batch %d of %d...\n", i, numBatches)
		
		// Load stocks from current batch
		stocks, err := LoadBatchFile(i)
		if err != nil {
			fmt.Printf("Error loading batch file %d: %v\n", i, err)
			continue
		}
		fmt.Printf("Loaded %d stocks from batch_%d.json\n", len(stocks), i)
		
		if len(stocks) == 0 {
			fmt.Printf("Batch %d has no stocks, skipping\n", i)
			continue
		}
		
			// Filter out stocks that are already sold
		stocks = FilterOutSoldStocks(stocks, soldStockMap)
		
		if len(stocks) == 0 {
			fmt.Printf("After filtering, batch %d has no unsold stocks left, skipping\n", i)
			continue
		}
		
		// Group stocks by date
		fmt.Printf("Grouping stocks for batch %d...\n", i)
		groupedStocks := GroupStocksByDate(stocks)
		
		// Assign random user_id, customer_id, order_date, and status to each group
		for j := range groupedStocks.Groups {
			// Assign random user ID and customer ID
			if len(userIDs) > 0 {
				groupedStocks.Groups[j].UserID = userIDs[rand.Intn(len(userIDs))]
			}
			
			if len(customerIDs) > 0 {
				groupedStocks.Groups[j].CustomerID = customerIDs[rand.Intn(len(customerIDs))]
				}
				
				// Find the newest item date in the group
				var newestDate time.Time
				for k, item := range groupedStocks.Groups[j].Items {
					if k == 0 || item.StockDate.After(newestDate) {
						newestDate = item.StockDate
					}
				}
				
				// Add random days (1-60) to the newest date
				randomDays := 1 + rand.Intn(60) // Random number between 1 and 60
				orderDate := newestDate.AddDate(0, 0, randomDays)
				
				// Format the order date as "YYYY-MM-DD HH:MM:SS"
				groupedStocks.Groups[j].OrderDate = orderDate.Format("2006-01-02 15:04:05")
				
				// Set the status to "Derived"
				groupedStocks.Groups[j].Status = "Derived"
			}
		
		fmt.Printf("Created %d groups for batch %d\n", len(groupedStocks.Groups), i)
		
		// Print some stats about the groups (limiting to first 3 groups only)
		for j, group := range groupedStocks.Groups {
			if j >= 3 && len(groupedStocks.Groups) > 3 {
				fmt.Printf("... and %d more groups\n", len(groupedStocks.Groups)-3)
				break
			}
			
			// Find min and max dates
			var minDate, maxDate time.Time
			for k, item := range group.Items {
				if k == 0 || item.StockDate.Before(minDate) {
					minDate = item.StockDate
				}
				if k == 0 || item.StockDate.After(maxDate) {
					maxDate = item.StockDate
				}
			}
			
			daysDiff := maxDate.Sub(minDate).Hours() / 24
			fmt.Printf("Group #%d: %d items, date range: %s to %s (%.1f days)\n",
				group.GroupID, len(group.Items), 
				minDate.Format("2006-01-02"), maxDate.Format("2006-01-02"), 
				daysDiff)
		}
		
		// Save the grouped stocks to a file
		outputFile := filepath.Join(outputDir, fmt.Sprintf("grouped_batch_%d.json", i))
		if err := SaveGroupedStocks(groupedStocks, outputFile); err != nil {
			fmt.Printf("Error saving grouped stocks for batch %d: %v\n", i, err)
			continue
		} else {
			fmt.Printf("Grouped stocks for batch %d saved to %s\n", i, outputFile)
		}
		
		// Post each group to the API
		fmt.Printf("Starting to post %d groups from batch %d to the API...\n", 
			len(groupedStocks.Groups), i)
		
		for j, group := range groupedStocks.Groups {
			fmt.Printf("Posting group %d/%d from batch %d...\n", 
				j+1, len(groupedStocks.Groups), i)
			
			// Try up to 3 times in case of temporary errors
			var success bool
			var postErr error
			
			for attempt := 1; attempt <= 3; attempt++ {
				if attempt > 1 {
					fmt.Printf("Retry attempt %d for group %d...\n", attempt, group.GroupID)
					// Wait before retrying
					time.Sleep(2 * time.Second)
				}
				
				success, postErr = PostOrderGroup(group)
				if success || postErr == nil {
					break
				}
				
				fmt.Printf("Error posting group %d (attempt %d): %v\n", 
					group.GroupID, attempt, postErr)
			}
			
			if success {
				fmt.Printf("Successfully posted group %d from batch %d\n", 
					group.GroupID, i)
			} else if postErr != nil {
				fmt.Printf("Failed to post group %d from batch %d after multiple attempts: %v\n", 
					group.GroupID, i, postErr)
			}
			
			// Wait between requests to avoid overloading the API
			time.Sleep(500 * time.Millisecond)
		}
		
		fmt.Printf("Finished posting all groups from batch %d\n", i)
		
		// Add a separator line for readability
		fmt.Println("----------------------------------------")
	}
	
	fmt.Printf("Finished processing all %d batches\n", numBatches)
}

// min returns the smaller of a and b
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
