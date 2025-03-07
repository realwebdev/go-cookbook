// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"time"
// )

// // Define the struct with the updated field names and data types
// type FoodItem struct {
// 	ID                int64   `json:"id"`
// 	DeleteAt          int64   `json:"deleteAt"` // int64 for deleted_at (epoch)
// 	NameE             string  `json:"nameE"`
// 	NameA             string  `json:"nameA"`
// 	RecipeE           string  `json:"recipeE"`
// 	RecipeA           string  `json:"recipeA"`
// 	Calories          float64 `json:"calories"`
// 	Color             string  `json:"color"`
// 	CreatedByID       int64   `json:"createdById"`
// 	CreatedAt         int64   `json:"createdAt"` // converted to epoch
// 	ModifiedByID      int64   `json:"modifiedById"`
// 	UpdatedAt         int64   `json:"updatedAt"` // converted to epoch
// 	RestaurantID      int64   `json:"restaurantId"`
// 	DefaultMediaID    int64   `json:"defaultMediaId"`
// 	WeightInGram      float64 `json:"weightInGram"` // renamed from widget_in_gram
// 	ShortDescriptionE string  `json:"shortDescriptionE"`
// 	ShortDescriptionA string  `json:"shortDescriptionA"`
// 	Carbs             float64 `json:"carbs"`
// 	Protein           float64 `json:"protein"`
// 	Fat               float64 `json:"fat"`
// 	IsHomeMade        bool    `json:"isHomeMade"`
// 	DietaryFiber      float64 `json:"dietaryFiber"`
// 	TransFat          float64 `json:"transFat"`
// 	SaturatedFat      float64 `json:"saturatedFat"`
// 	Calcium           float64 `json:"calcium"`
// 	Sodium            float64 `json:"sodium"`
// 	Iron              float64 `json:"iron"`
// 	Sugar             float64 `json:"sugar"`
// 	Image             string  `json:"image"`
// 	IsSupermarket     bool    `json:"isSupermarket"`
// }

// // Helper function to convert time to epoch
// func timeToEpoch(timeStr string) int64 {
// 	layout := "2006-01-02 15:04:05"
// 	t, err := time.Parse(layout, timeStr)
// 	if err != nil {
// 		fmt.Printf("Error parsing time: %v\n", err)
// 		return 0
// 	}
// 	return t.Unix()
// }

// func main() {
// 	// Read the JSON data from the file
// 	fileData, err := ioutil.ReadFile("foodmenu.json")
// 	if err != nil {
// 		fmt.Printf("Error reading file: %v\n", err)
// 		return
// 	}

// 	// Unmarshal the data into a slice of map (to process dynamic structure)
// 	var foodItems []map[string]interface{}
// 	err = json.Unmarshal(fileData, &foodItems)
// 	if err != nil {
// 		fmt.Printf("Error unmarshalling JSON: %v\n", err)
// 		return
// 	}

// 	// Create a new slice to store the updated data
// 	var updatedItems []FoodItem

// 	// Iterate over each item in the original data
// 	for _, item := range foodItems {
// 		foodItem := FoodItem{}

// 		if id, ok := item["id"].(float64); ok {
// 			foodItem.ID = int64(id)
// 		} else {
// 			log.Printf("Error: id is not a float64")
// 			continue
// 		}

// 		foodItem.DeleteAt = 0 // default to 0 if deleted_at is nil

// 		if nameE, ok := item["name_e"].(string); ok {
// 			foodItem.NameE = nameE
// 		} else {
// 			log.Printf("Error: name_e is not a string")
// 			continue
// 		}

// 		if nameA, ok := item["name_a"].(string); ok {
// 			foodItem.NameA = nameA
// 		} else {
// 			log.Printf("Error: name_a is not a string")
// 			continue
// 		}

// 		if recipeE, ok := item["recipe_e"].(string); ok {
// 			foodItem.RecipeE = recipeE
// 		} else {
// 			log.Printf("Error: recipe_e is not a string")
// 			continue
// 		}

// 		if recipeA, ok := item["recipe_a"].(string); ok {
// 			foodItem.RecipeA = recipeA
// 		} else {
// 			log.Printf("Error: recipe_a is not a string")
// 			continue
// 		}

// 		if calories, ok := item["calories"].(float64); ok {
// 			foodItem.Calories = calories
// 		} else {
// 			log.Printf("Error: calories is not a float64")
// 			continue
// 		}

// 		if color, ok := item["color"].(string); ok {
// 			foodItem.Color = color
// 		} else {
// 			log.Printf("Error: color is not a string")
// 			continue
// 		}

// 		if createdByID, ok := item["created_by_id"].(float64); ok {
// 			foodItem.CreatedByID = int64(createdByID)
// 		} else {
// 			log.Printf("Error: created_by_id is not a float64")
// 			continue
// 		}

// 		if modifiedByID, ok := item["modified_by_id"].(float64); ok {
// 			foodItem.ModifiedByID = int64(modifiedByID)
// 		} else {
// 			log.Printf("Error: modified_by_id is not a float64")
// 			continue
// 		}

// 		if restaurantID, ok := item["restaurant_id"].(float64); ok {
// 			foodItem.RestaurantID = int64(restaurantID)
// 		} else {
// 			log.Printf("Error: restaurant_id is not a float64")
// 			continue
// 		}

// 		if defaultMediaID, ok := item["default_media_id"].(float64); ok {
// 			foodItem.DefaultMediaID = int64(defaultMediaID)
// 		} else {
// 			log.Printf("Error: default_media_id is not a float64")
// 			continue
// 		}

// 		if weightInGram, ok := item["widget_in_gram"].(float64); ok {
// 			foodItem.WeightInGram = weightInGram
// 		} else {
// 			log.Printf("Error: widget_in_gram is not a float64")
// 			continue
// 		}

// 		if shortDescriptionE, ok := item["short_description_e"].(string); ok {
// 			foodItem.ShortDescriptionE = shortDescriptionE
// 		} else {
// 			log.Printf("Error: short_description_e is not a string")
// 			continue
// 		}

// 		if shortDescriptionA, ok := item["short_description_a"].(string); ok {
// 			foodItem.ShortDescriptionA = shortDescriptionA
// 		} else {
// 			log.Printf("Error: short_description_a is not a string")
// 			continue
// 		}

// 		if carbs, ok := item["carbs"].(float64); ok {
// 			foodItem.Carbs = carbs
// 		} else {
// 			log.Printf("Error: carbs is not a float64")
// 			continue
// 		}

// 		if protein, ok := item["protein"].(float64); ok {
// 			foodItem.Protein = protein
// 		} else {
// 			log.Printf("Error: protein is not a float64")
// 			continue
// 		}

// 		if fat, ok := item["fat"].(float64); ok {
// 			foodItem.Fat = fat
// 		} else {
// 			log.Printf("Error: fat is not a float64")
// 			continue
// 		}

// 		if isHomeMade, ok := item["is_home_made"].(bool); ok {
// 			foodItem.IsHomeMade = isHomeMade
// 		} else {
// 			log.Printf("Error: is_home_made is not a bool")
// 			continue
// 		}

// 		if dietaryFiber, ok := item["dietary_fiber"].(float64); ok {
// 			foodItem.DietaryFiber = dietaryFiber
// 		} else {
// 			log.Printf("Error: dietary_fiber is not a float64")
// 			continue
// 		}

// 		if transFat, ok := item["trans_fat"].(float64); ok {
// 			foodItem.TransFat = transFat
// 		} else {
// 			log.Printf("Error: trans_fat is not a float64")
// 			continue
// 		}

// 		if saturatedFat, ok := item["saturated_fat"].(float64); ok {
// 			foodItem.SaturatedFat = saturatedFat
// 		} else {
// 			log.Printf("Error: saturated_fat is not a float64")
// 			continue
// 		}

// 		if calcium, ok := item["calcium"].(float64); ok {
// 			foodItem.Calcium = calcium
// 		} else {
// 			log.Printf("Error: calcium is not a float64")
// 			continue
// 		}

// 		if sodium, ok := item["sodium"].(float64); ok {
// 			foodItem.Sodium = sodium
// 		} else {
// 			log.Printf("Error: sodium is not a float64")
// 			continue
// 		}

// 		if iron, ok := item["iron"].(float64); ok {
// 			foodItem.Iron = iron
// 		} else {
// 			log.Printf("Error: iron is not a float64")
// 			continue
// 		}

// 		if sugar, ok := item["sugar"].(float64); ok {
// 			foodItem.Sugar = sugar
// 		} else {
// 			log.Printf("Error: sugar is not a float64")
// 			continue
// 		}

// 		if image, ok := item["image"].(string); ok {
// 			foodItem.Image = image
// 		} else {
// 			log.Printf("Error: image is not a string")
// 			continue
// 		}

// 		if isSupermarket, ok := item["is_supermarket"].(bool); ok {
// 			foodItem.IsSupermarket = isSupermarket
// 		} else {
// 			log.Printf("Error: is_supermarket is not a bool")
// 			continue
// 		}

// 		updatedItems = append(updatedItems, foodItem)

// 		// foodItem := FoodItem{
// 		// 	ID:                int64(item["id"].(float64)),
// 		// 	DeleteAt:          0, // default to 0 if deleted_at is nil
// 		// 	NameE:             item["name_e"].(string),
// 		// 	NameA:             item["name_a"].(string),
// 		// 	RecipeE:           item["recipe_e"].(string),
// 		// 	RecipeA:           item["recipe_a"].(string),
// 		// 	Calories:          item["calories"].(float64),
// 		// 	Color:             item["color"].(string),
// 		// 	CreatedByID:       int64(item["created_by_id"].(float64)),
// 		// 	ModifiedByID:      int64(item["modified_by_id"].(float64)),
// 		// 	RestaurantID:      int64(item["restaurant_id"].(float64)),
// 		// 	DefaultMediaID:    int64(item["default_media_id"].(float64)),
// 		// 	WeightInGram:      item["widget_in_gram"].(float64),
// 		// 	ShortDescriptionE: item["short_description_e"].(string),
// 		// 	ShortDescriptionA: item["short_description_a"].(string),
// 		// 	Carbs:             item["carbs"].(float64),
// 		// 	Protein:           item["protein"].(float64),
// 		// 	Fat:               item["fat"].(float64),
// 		// 	IsHomeMade:        int(item["is_home_made"].(float64)),
// 		// 	DietaryFiber:      item["dietary_fiber"].(float64),
// 		// 	TransFat:          item["trans_fat"].(float64),
// 		// 	SaturatedFat:      item["saturated_fat"].(float64),
// 		// 	Calcium:           item["calcium"].(float64),
// 		// 	Sodium:            item["sodium"].(float64),
// 		// 	Iron:              item["iron"].(float64),
// 		// 	Sugar:             item["suger"].(float64), // renamed suger to sugar
// 		// 	Image:             item["image"].(string),
// 		// 	IsSupermarket:     int(item["is_supermarket"].(float64)),
// 		// }

// 		// Convert creation_date and last_update_date to epoch
// 		if item["creation_date"] != nil {
// 			foodItem.CreatedAt = timeToEpoch(item["creation_date"].(string))
// 		}
// 		if item["last_update_date"] != nil {
// 			foodItem.UpdatedAt = timeToEpoch(item["last_update_date"].(string))
// 		}

// 		// Check if deleted_at is nil, otherwise convert to epoch
// 		if item["deleted_at"] != nil {
// 			foodItem.DeleteAt = timeToEpoch(item["deleted_at"].(string))
// 		}

// 		// Append the updated item to the new slice
// 		updatedItems = append(updatedItems, foodItem)
// 	}

// 	// Marshal the updated data into JSON
// 	updatedJSON, err := json.MarshalIndent(updatedItems, "", "  ")
// 	if err != nil {
// 		fmt.Printf("Error marshalling updated data: %v\n", err)
// 		return
// 	}

// 	// Write the updated data to a new file
// 	err = ioutil.WriteFile("foodmenuchanged.json", updatedJSON, 0644)
// 	if err != nil {
// 		fmt.Printf("Error writing to file: %v\n", err)
// 		return
// 	}

// 	fmt.Println("Data successfully updated and saved to foodmenuchanged.json")
// }

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/rs/xid"
)

// Struct to match the original JSON data

type FoodItem struct {
	Slug              string  `json:"slug"`
	Id                string  `json:"id"`        // Changed from int to int64
	DeletedAt         int64   `json:"deletedAt"` // Changed from datetime to epoch (int64)
	NameE             string  `json:"nameE"`
	NameA             string  `json:"nameA"`
	RecipeE           string  `json:"recipeE"`
	RecipeA           string  `json:"recipeA"`
	Calories          float64 `json:"calories"`
	Color             string  `json:"color"`
	CreatedById       string  `json:"createdById"`
	CreatedAt         int64   `json:"createdAt"` // Changed from datetime to epoch (int64)
	ModifiedById      string  `json:"modifiedById"`
	UpdatedAt         int64   `json:"updatedAt"` // Changed from datetime to epoch (int64)
	RestaurantId      uint64  `json:"restaurantId"`
	DefaultMediaId    uint64  `json:"defaultMediaId"`
	WeightInGram      float64 `json:"weightInGram"` // Renamed from widget_in_gram
	ShortDescriptionE string  `json:"shortDescriptionE"`
	ShortDescriptionA string  `json:"shortDescriptionA"`
	Carbs             float64 `json:"carbs"`
	Protein           float64 `json:"protein"`
	Fat               float64 `json:"fat"`
	IsHomeMade        bool    `json:"isHomeMade"` // Changed from tinyint(1) to bool
	DietaryFiber      float64 `json:"dietaryFiber"`
	TransFat          float64 `json:"transFat"`
	SaturatedFat      float64 `json:"saturatedFat"`
	Calcium           float64 `json:"calcium"`
	Sodium            float64 `json:"sodium"`
	Iron              float64 `json:"iron"`
	Sugar             float64 `json:"sugar"` // Fixed spelling from suger to sugar
	Image             string  `json:"image"`
	IsSupermarket     bool    `json:"isSupermarket"` // Changed from tinyint(1) to bool
	FoodFilter        string  `json:"foodFilter"`    // this is a new field to add to the data // category
	RestaurantName    string  `json:"restaurantName"`
	IsCafe            bool    `json:"isCafe"`
	CafeName          string  `json:"cafeName"`
	IsRamadan         bool    `json:"isRamadan"`
	IsRestraunt       bool    `json:"isRestraunt"`
}

// IsPotato          bool    `json:"isPotato"`
// IsSalad           bool    `json:"isSalad"`
// IsSoup            bool    `json:"isSoup"`

// IsSaucesAndDips   bool    `json:"isSaucesAndDips"`
// IsProtein         bool    `json:"isProtein"`
// IsNuts            bool    `json:"isNuts"`
// IsMilkAndCheese   bool    `json:"isMilkAndCheese"`
// IsJuice           bool    `json:"isJuice"`
// IsIcecream        bool    `json:"isIcecream"`
// IsHatabBakery     bool    `json:"isHatabBakery"`
// IsGrainAndFlour   bool    `json:"isGrainAndFlour"`
// IsFrozenProducts  bool    `json:"isFrozenProducts"`
// IsFruit           bool    `json:"isFruit"`
// IsFatayer         bool    `json:"isFatayer"`
// IsDates           bool    `json:"isDates"`
// IsCannedFood      bool    `json:"isCannedFood"`
// IsChocolate       bool    `json:"isChocolate"`
// IsChips           bool    `json:"isChips"`
// IsCornFlakes      bool    `json:"isCornFlakes"`
// IsCoffee          bool    `json:"isCoffee"`
// IsCakeAndSweet    bool    `json:"isCakeAndSweet"`
// IsBread           bool    `json:"isBread"`
// IsBiscuit         bool    `json:"isBiscuit"`
// IsAppetizers      bool    `json:"isAppetizers"`
// IsArabicFood      bool    `json:"isArabicFood"`
// IsKarazApproved   bool    `json:"isKarazApproved"`

/*
	restaurant
	cafe
	homeMade
	supermarket
	bread
	cakeAndSweet
	pastaAndNoodles
	Rice
	salad
	soup
	hatabBakery
	cornFlakes
	coffee
	milkAndCheese
	vegetables
	fruit
	frozenProducts
	chocolate
	biscuit
	nuts
	cannedFood
	juice
	fatayer
	chips
	grainAndFlour
	protein
	appetizers
	arabicFood
	karazApproved
	dates
	icecream
	saucesAndDips
	potatoes
	ramadan
*/

// Helper function to convert time to epoch
func timeToEpoch(datetimeStr string) int64 {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, datetimeStr)
	if err != nil {
		log.Printf("Error parsing time: %s\n", err)
		return 0
	}
	return t.Unix()
}

func main() {
	// Read the foodmenu.json file
	jsonFile, err := os.Open("foodmenu.json")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	// Unmarshal JSON data
	var foodItems []map[string]interface{}
	if err := json.Unmarshal(byteValue, &foodItems); err != nil {
		log.Fatalf("Error unmarshaling JSON: %s", err)
	}

	// Process each item
	for _, item := range foodItems {
		// Rename and convert fields
		if lastUpdate, ok := item["last_update_date"].(string); ok {
			item["updatedAt"] = timeToEpoch(lastUpdate)
		}
		delete(item, "last_update_date")

		if creationDate, ok := item["creation_date"].(string); ok {
			item["createdAt"] = timeToEpoch(creationDate)
		}
		delete(item, "creation_date")

		if weightInGram, ok := item["widget_in_gram"].(float64); ok {
			item["weightInGram"] = weightInGram
		}
		delete(item, "widget_in_gram")

		if deletedAt, ok := item["deleted_at"].(string); ok {
			item["deletedAt"] = timeToEpoch(deletedAt)
		} else if item["deleted_at"] == nil {
			item["deletedAt"] = 0 // Set to epoch forever
		}
		delete(item, "deleted_at")

		// Change id to slug and data type to string
		//guid := xid.New()
		//return guid.String()
		// Change id to slug and data type to string
		if _, ok := item["id"].(float64); ok {
			guid := xid.New()
			item["slug"] = guid.String()
		}
		// delete(item, "id")

		// convert id to string
		if id, ok := item["id"].(float64); ok {
			item["id"] = fmt.Sprintf("%d", int(id))
		}

		// Convert modified_by_id to string and rename to modifiedById
		if modifiedByID, ok := item["modified_by_id"].(float64); ok {
			item["modifiedById"] = fmt.Sprintf("%d", int(modifiedByID))
		}
		delete(item, "modified_by_id")

		// Convert created_by_id to string and rename to createdById
		if createdByID, ok := item["created_by_id"].(float64); ok {
			item["createdById"] = fmt.Sprintf("%d", int(createdByID))
		}
		delete(item, "created_by_id")

		// is supermarket to bool
		if isSupermarket, ok := item["is_supermarket"].(string); ok {
			if isSupermarket == "1" {
				item["isSupermarket"] = true
			} else {
				item["isSupermarket"] = false
			}
		}
		delete(item, "is_supermarket")

		// Add the foodFilter field based on conditions
		if restaurantID, ok := item["restaurant_id"].(float64); ok && int64(restaurantID) == 40 {
			item["foodFilter"] = "cakeAndSweet"
		}

		//

		// Add the foodFilter field based on conditions
		if restaurantID, ok := item["restaurant_id"].(float64); ok {
			// Example condition: set foodFilter to "bread" if restaurant_id == 15
			if uint64(restaurantID) == 15 {
				item["isRestraunt"] = true
				item["restaurantName"] = "KFC"
				item["isCafe"] = false
				item["cafeName"] = ""
			} else if uint64(restaurantID) == 18 {
				item["isRestraunt"] = true
				item["restaurantName"] = "Hardee's"
				item["isCafe"] = false
				item["cafeName"] = ""
			} else if uint64(restaurantID) == 19 {
				item["isRestraunt"] = true
				item["restaurantName"] = "Pizza Hut"
				item["isCafe"] = false
				item["cafeName"] = ""
			} else if uint64(restaurantID) == 20 {
				item["isRestraunt"] = true
				item["restaurantName"] = "FireGrill"
				item["isCafe"] = false
				item["cafeName"] = ""
			} else if uint64(restaurantID) == 21 {
				item["isRestraunt"] = true
				item["restaurantName"] = "McDonald's"
				item["isCafe"] = false
				item["cafeName"] = ""
			} else if uint64(restaurantID) == 22 {
				item["isRestraunt"] = true
				item["restaurantName"] = "Papa John's Pizza"
				item["isCafe"] = false
				item["cafeName"] = ""
			} else if uint64(restaurantID) == 24 {
				item["isRestraunt"] = true
				item["restaurantName"] = "Domino's Pizza"
				item["isCafe"] = false
				item["cafeName"] = ""
			} else if uint64(restaurantID) == 25 {
				item["isRestraunt"] = true
				item["restaurantName"] = "Burger King"
				item["isCafe"] = false
				item["cafeName"] = ""
			} else if uint64(restaurantID) == 26 {
				item["isRestraunt"] = true
				item["restaurantName"] = "Subway"
				item["isCafe"] = false
				item["cafeName"] = ""
			} else if uint64(restaurantID) == 28 {
				item["isRestraunt"] = true
				item["restaurantName"] = "Baskin - Robbins"
				item["isCafe"] = false
				item["cafeName"] = ""
			} else if uint64(restaurantID) == 31 {
				item["isRestraunt"] = true
				item["restaurantName"] = "Cinnabon"
				item["isCafe"] = true
				item["cafeName"] = "Cinnabon"
			} else if uint64(restaurantID) == 34 {
				item["isRestraunt"] = true
				item["restaurantName"] = "Wendy's"
				item["isCafe"] = false
				item["cafeName"] = ""
			} else if uint64(restaurantID) == 36 {
				item["isRestraunt"] = true
				item["restaurantName"] = "Albaik"
				item["isCafe"] = false
				item["cafeName"] = ""
			} else if uint64(restaurantID) == 39 {
				item["isRestraunt"] = true
				item["restaurantName"] = "Krispy Kreme"
				item["isCafe"] = false
				item["cafeName"] = ""
			} else if uint64(restaurantID) == 41 {
				item["isRestraunt"] = true
				item["restaurantName"] = "Pinkberry"
				item["isCafe"] = false
				item["cafeName"] = ""
			} else if uint64(restaurantID) == 45 {
				item["isRestraunt"] = true
				item["restaurantName"] = "Applebee's"
				item["isCafe"] = false
				item["cafeName"] = ""
			} else if uint64(restaurantID) == 47 {
				item["isRestraunt"] = true
				item["restaurantName"] = "Five Guys"
				item["isCafe"] = false
				item["cafeName"] = ""
			} else if uint64(restaurantID) == 48 {
				item["isRestraunt"] = true
				item["restaurantName"] = "Nando's"
				item["isCafe"] = false
				item["cafeName"] = ""
			} else if uint64(restaurantID) == 27 {
				item["restaurantName"] = "starbucks"
				item["isCafe"] = true
				item["cafeName"] = "starbucks"
			} else if uint64(restaurantID) == 37 {
				item["restaurantName"] = "Dunkin' Donuts"
				item["isCafe"] = true
				item["cafeName"] = "Dunkin' Donuts"
			} else if uint64(restaurantID) == 40 {
				item["restaurantName"] = "Barns Cafe"
				item["isCafe"] = true
				item["cafeName"] = "Barns Cafe"
			} else if uint64(restaurantID) == 34 {
				item["isRestraunt"] = true
				item["restaurantName"] = "Albaik"
				item["isCafe"] = false
				item["cafeName"] = ""
			} else {
				item["isRestraunt"] = false
				item["restaurantName"] = ""
				item["isCafe"] = false
				item["cafeName"] = ""
			}
		}

		// // based on isSupermarket
		// if isSupermarket, ok := item["is_supermarket"].(float64); ok {
		// 	if isSupermarket == 1 {
		// 		item["isSup"] = "supermarket"
		// 	}
		// }

		// based on homeMade
		if isHomeMade, ok := item["is_home_made"].(float64); ok {
			if isHomeMade == 1 {
				item["isRamadan"] = true
			} else {
				item["isRamadan"] = false
			}
		}

		// // based on name
		// if nameE, ok := item["name_e"].(string); ok {
		// 	if strings.Contains(nameE, "potato waffle") || strings.Contains(nameE, "potato smiles") || strings.Contains(nameE, "potato slices") || strings.Contains(nameE, "potato salad") || strings.Contains(nameE, "baby potato") || strings.Contains(nameE, "Indian Bombay potato") {
		// 		item["isPotatoes"] = true
		// 		// if contain mayyonaie, dip, ketchup then is isSaucesAndDips
		// 	} else if strings.Contains(nameE, "mayonnaise") || strings.Contains(nameE, "tomato ketchup") || strings.Contains(nameE, "sauce") || strings.Contains(nameE, "dip") || strings.Contains(nameE, "ranch") {
		// 		item["isSaucesAndDips"] = true
		// 	} else if strings.Contains(nameE, "icecream") || strings.Contains(nameE, "ice cream") {
		// 		item["isIcecream"] = true
		// 	} else if strings.Contains(nameE, "dates") {
		// 		item["isDates"] = true
		// 	} else if strings.Contains(nameE, "Zollipops") || strings.Contains(nameE, "Whisps -") || strings.Contains(nameE, "mini cookies") || strings.Contains(nameE, "GREEN MOCKTAILS") || strings.Contains(nameE, "GREEN FRUIT") || strings.Contains(nameE, "GREEN - LEMON") || strings.Contains(nameE, "wafers vanilla") || strings.Contains(nameE, "peanut butter with") || strings.Contains(nameE, "Keto Cookie") || strings.Contains(nameE, "soft baked mini") {
		// 		item["isKarazApproved"] = true
		// 	} else if strings.Contains(nameE, "soup") {
		// 		item["IsSoup"] = true
		// 	} else if strings.Contains(nameE, "salad") || strings.Contains(nameE, "coleslaw") || strings.Contains(nameE, "baba ghanoush with") {
		// 		item["IsSalad"] = true
		// 	} else if strings.Contains(nameE, "rice") || strings.Contains(nameE, "Biryani") || strings.Contains(nameE, "Chicken Kabsa") || strings.Contains(nameE, "Quinoa") || strings.Contains(nameE, "Makloubeh") {
		// 		item["IsRice"] = true
		// 	} else if strings.Contains(nameE, "pasta")  || strings.Contains(nameE, "Spaghetti") || strings.Contains(nameE, "Macaroni") || strings.Contains(nameE, "lasagna") || strings.Contains(nameE, "fettuccine") || strings.Contains(nameE, "indomie -") || strings.Contains(nameE, "noodles") || strings.Contains(nameE, "chow mein") {
		// 		item["IsPastaAndNoodles"] = true
		// 		// chocolate puff, strawberry puff, panna cotta, lemon meringue, chocolate cheesecake, custard, tart, halwa, zalabia, baba ghanoush with bread, kunafa, baklawa, cinnabon - classic roll, pineaple dried, apricots dried, jelly, creme brulee, moussee
		// 	} else if strings.Contains(nameE, "cake") || strings.Contains(nameE, "sweet") || strings.Contains(nameE, "puff") || strings.Contains(nameE, "panna cotta") || strings.Contains(nameE, "lemon meringue") || strings.Contains(nameE, "cheesecake") || strings.Contains(nameE, "custard") || strings.Contains(nameE, "tart") || strings.Contains(nameE, "halwa") || strings.Contains(nameE, "zalabia") || strings.Contains(nameE, "baba ghanoush with bread") || strings.Contains(nameE, "kunafa") || strings.Contains(nameE, "baklawa") || strings.Contains(nameE, "cinnabon - classic roll") || strings.Contains(nameE, "pineaple dried") || strings.Contains(nameE, "apricots dried") || strings.Contains(nameE, "jelly") || strings.Contains(nameE, "creme brulee") || strings.Contains(nameE, "moussee") {
		// 		item["IsCakeAndSweet"] = true
		// 		// corn flakes, cereals, fitness, nesquik, Kelloog's,cheerios frootloops
		// 	} else if strings.Contains(nameE, "corn flakes") || strings.Contains(nameE, "cereals") || strings.Contains(nameE, "fitness") || strings.Contains(nameE, "nesquik") || strings.Contains(nameE, "Kelloog's") || strings.Contains(nameE, "cheerios") || strings.Contains(nameE, "frootloops") {
		// 		item["IsCornFlakes"] = true
		// 	 // green olive tapenade, AlHatab, altannor, potato roll bread, whole wheat mimi burger, rusk
		// 	} else if strings.Contains(nameE, "green olive tapenade") || strings.Contains(nameE, "AlHatab") || strings.Contains(nameE, "altannor") || strings.Contains(nameE, "potato roll bread") || strings.Contains(nameE, "whole wheat mimi burger") || strings.Contains(nameE, "rusk") {
		// 		item["IsHatabBakery"] = true
		// 	// chicken burger, beef burger, hummus and bread, L'usine -,  croissant, poppy seed roll, bagel sesame, almarai - double layered, facaccia,
		// 	} else if strings.Contains(nameE, "chicken burger") || strings.Contains(nameE, "beef burger") || strings.Contains(nameE, "hummus and bread") || strings.Contains(nameE, "L'usine -") || strings.Contains(nameE, "croissant") || strings.Contains(nameE, "poppy seed roll") || strings.Contains(nameE, "bagel sesame") || strings.Contains(nameE, "almarai - double layered") || strings.Contains(nameE, "facaccia") {
		// 		item["IsBread"] = true
		// 		// almarai one, Hot Chocolate, Latte, Espresso, Cappuccino, Coffee Mate
		// 	} else if strings.Contains(nameE, "almarai one") || strings.Contains(nameE, "Hot Chocolate") || strings.Contains(nameE, "Latte") || strings.Contains(nameE, "Espresso") || strings.Contains(nameE, "Cappuccino") || strings.Contains(nameE, "Coffee Mate") {
		// 		item["IsCoffee"] = true
		// 		// ayran, Almarai -protein, milk, puck -, butter, yogurt
		// 	} else if strings.Contains(nameE, "ayran") || strings.Contains(nameE, "Almarai -protein") || strings.Contains(nameE, "milk") || strings.Contains(nameE, "puck -") || strings.Contains(nameE, "butter") || strings.Contains(nameE, "yogurt") {
		// 		item["IsMilkAndCheese"] = true
		// 		// Avocado, bamya, baby potatoes, Indian Bombay potato, Fried CauliFlower, onions, sweet corn, sweet corn on the cob, cabbage, zucchini, cauliflower, broccoli, radish, pumpkin, ginger, mushrom, eggplant, sweet potato
		// 	} else if strings.Contains(nameE, "Avocado") || strings.Contains(nameE, "bamya") || strings.Contains(nameE, "baby potatoes") || strings.Contains(nameE, "Indian Bombay potato") || strings.Contains(nameE, "Fried CauliFlower") || strings.Contains(nameE, "onions") || strings.Contains(nameE, "sweet corn") || strings.Contains(nameE, "sweet corn on the cob") || strings.Contains(nameE, "cabbage") || strings.Contains(nameE, "zucchini") || strings.Contains(nameE, "cauliflower") || strings.Contains(nameE, "broccoli") || strings.Contains(nameE, "radish") || strings.Contains(nameE, "pumpkin") || strings.Contains(nameE, "ginger") || strings.Contains(nameE, "mushrom") || strings.Contains(nameE, "eggplant") || strings.Contains(nameE, "sweet potato") {
		// 		item["IsVegetables"] = true
		// 		// apple, banana, orange, strawberry, pineapple, watermelon, melon, grapes, kiwi, mango, peach, pear, cherry, blueberry, raspberry, blackberry, cranberry, pomegranate, lemon, lime, grapefruit, avocado, coconut, date, fig, guava, kiwi, lychee, papaya, passion fruit, persimmon, plum, tangerine, apricot, nectarine, dragon fruit, star fruit, cantaloupe, honeydew, jackfruit, kiwano, kumquat, longan, rambutan, sapodilla, soursop, ugli fruit, yuzu, fruit yougurt, coconut yougurt, pomelo, nectarine, raisins, dates, kiwi, clementine
		// 	} else if strings.Contains(nameE, "apple") || strings.Contains(nameE, "banana") || strings.Contains(nameE, "orange") || strings.Contains(nameE, "strawberry") || strings.Contains(nameE, "pineapple") || strings.Contains(nameE, "watermelon") || strings.Contains(nameE, "melon") || strings.Contains(nameE, "grapes") || strings.Contains(nameE, "kiwi") || strings.Contains(nameE, "mango") || strings.Contains(nameE, "peach") || strings.Contains(nameE, "pear") || strings.Contains(nameE, "cherry") || strings.Contains(nameE, "blueberry") || strings.Contains(nameE, "raspberry") || strings.Contains(nameE, "blackberry") || strings.Contains(nameE, "cranberry") || strings.Contains(nameE, "pomegranate") || strings.Contains(nameE, "lemon") || strings.Contains(nameE, "lime") || strings.Contains(nameE, "grapefruit") || strings.Contains(nameE, "avocado") || strings.Contains(nameE, "coconut") || strings.Contains(nameE, "date") || strings.Contains(nameE, "fig") || strings.Contains(nameE, "guava") || strings.Contains(nameE, "kiwi") || strings.Contains(nameE, "lychee") || strings.Contains(nameE, "papaya") || strings.Contains(nameE, "passion fruit") || strings.Contains(nameE, "persimmon") || strings.Contains(nameE, "plum") || strings.Contains(nameE, "tangerine") || strings.Contains(nameE, "apricot") || strings.Contains(nameE, "nectarine") || strings.Contains(nameE, "dragon fruit") || strings.Contains(nameE, "star fruit") || strings.Contains(nameE, "cantaloupe") || strings.Contains(nameE, "honeydew") || strings.Contains(nameE, "jackfruit") || strings.Contains(nameE, "kiwano") || strings.Contains(nameE, "kumquat") || strings.Contains(nameE, "longan") || strings.Contains(nameE, "rambutan") || strings.Contains(nameE, "sapodilla") || strings.Contains(nameE, "soursop") || strings.Contains(nameE, "ugli fruit") || strings.Contains(nameE, "yuzu") || strings.Contains(nameE, "fruit yougurt") || strings.Contains(nameE, "coconut yougurt") || strings.Contains(nameE, "pomelo") || strings.Contains(nameE, "nectarine") || strings.Contains(nameE, "raisins") || strings.Contains(nameE, "dates") || strings.Contains(nameE, "kiwi") || strings.Contains(nameE, "clementine") {
		// 		item["IsFruit"] = true

		// 		// frozen producs chicken nuggets, onion rings, small pizza crust, vegetariana pizza, chicken pizza - sunbulah, sunbulah
		// 	} else if strings.Contains(nameE, "chicken nuggets") || strings.Contains(nameE, "onion rings") || strings.Contains(nameE, "small pizza crust") || strings.Contains(nameE, "vegetariana pizza") || strings.Contains(nameE, "chicken pizza - sunbulah") || strings.Contains(nameE, "sunbulah") {
		// 		item["IsFrozenProducts"] = true
		// 		// chocolate  chocolate tart, chocolate mousse, GALAXY, CADBURY,
		// 	} else if strings.Contains(nameE, "chocolate") || strings.Contains(nameE, "chocolate tart") || strings.Contains(nameE, "chocolate mousse") || strings.Contains(nameE, "GALAXY") || strings.Contains(nameE, "CADBURY") {
		// 		item["IsChocolate"] = true
		// 		// biscuit, biscuits, oreo, lotus, digestive, marie, petit beurre, prawn cracker, ice cream waffle biscuits, ritz crackers, almarai 7days-,  highkey -, wafer, Ginger biscuit, Fig roll, chocolate chip cookie, malted milk
		// 	} else if strings.Contains(nameE, "biscuit") || strings.Contains(nameE, "biscuits") || strings.Contains(nameE, "oreo") || strings.Contains(nameE, "lotus") || strings.Contains(nameE, "digestive") || strings.Contains(nameE, "marie") || strings.Contains(nameE, "petit beurre") || strings.Contains(nameE, "prawn cracker") || strings.Contains(nameE, "ice cream waffle biscuits") || strings.Contains(nameE, "ritz crackers") || strings.Contains(nameE, "almarai 7days-") || strings.Contains(nameE, "highkey -") || strings.Contains(nameE, "wafer") || strings.Contains(nameE, "Ginger biscuit") || strings.Contains(nameE, "Fig roll") || strings.Contains(nameE, "chocolate chip cookie") || strings.Contains(nameE, "malted milk") {
		// 		item["IsBiscuit"] = true
		// 		// nuts, nut, almond, cashew, pistachio, walnut, hazelnut, pecan, bombay mix, DELUXE MIXED, SWEET & SALTY, sunflower seeds salted,
		// 	} else if strings.Contains(nameE, "nuts") || strings.Contains(nameE, "nut") || strings.Contains(nameE, "almond") || strings.Contains(nameE, "cashew") || strings.Contains(nameE, "pistachio") || strings.Contains(nameE, "walnut") || strings.Contains(nameE, "hazelnut") || strings.Contains(nameE, "pecan") || strings.Contains(nameE, "bombay mix") || strings.Contains(nameE, "DELUXE MIXED") || strings.Contains(nameE, "SWEET & SALTY") || strings.Contains(nameE, "sunflower seeds salted") {
		// 		item["IsNuts"] = true
		// 		// canned, canned food, canned products, canned fruit, canned vegetables

		// 	}else if strings.Contains(nameE, "vegetables") || strings.Contains(nameE, "vegetable") || strings.Contains(nameE, "salad") || strings.Contains(nameE, "coleslaw") || strings.Contains(nameE, "baba ghanoush with") {
		// 		item["IsVegetables"] = true
		// 	} else if strings.Contains(nameE, "fruit") || strings.Contains(nameE, "fruits") || strings.Contains(nameE, "fruit salad") || strings.Contains(nameE, "fruit platter") || strings.Contains(nameE, "fruit cocktail") || strings.Contains(nameE, "fruit juice") || strings.Contains(nameE, "fruit smoothie") {
		// 		item["IsFruit"] = true
		// 	} else if strings.Contains(nameE, "frozen") || strings.Contains(nameE, "frozen products") || strings.Contains(nameE, "frozen food") || strings.Contains(nameE, "frozen yogurt") || strings.Contains(nameE, "frozen yogurt") || strings.Contains(nameE, "frozen dessert") {
		// 		item["IsFrozenProducts"] = true
		// 	} else if strings.Contains(nameE, "chocolate") {
		// 		item["IsChocolate"] = true
		// 	} else if strings.Contains(nameE, "biscuit") || strings.Contains(nameE, "biscuits") {
		// 		item["IsBiscuit"] = true
		// 	} else if strings.Contains(nameE, "nuts") || strings.Contains(nameE, "nut") {
		// 		item["IsNuts"] = true
		// 	} else if strings.Contains(nameE, "canned") || strings.Contains(nameE, "canned food") || strings.Contains(nameE, "canned products") || strings.Contains(nameE, "canned fruit") || strings.Contains(nameE, "canned vegetables") {
		// 		item["IsCannedFood"] = true
		// 	} else if strings.Contains(nameE, "juice") || strings.Contains(nameE, "juices") || strings.Contains(nameE, "juice bar") || strings.Contains(nameE, "juice shop") || strings.Contains(nameE, "juice corner") || strings.Contains(nameE, "juice center") || strings.Contains(nameE, "juice factory") {
		// 		item["IsJuice"] = true
		// 	} else if strings.Contains(nameE, "fatayer") || strings.Contains(nameE, "fatayer") {

		// 	} else if

		// 	} else {
		// 		item["isPotatoes"] = false
		// 		item["isSaucesAndDips"] = false
		// 		item["isIcecream"] = false
		// 		item["isDates"] = false
		// 		item["isKarazApproved"] = false
		// 		item["IsSoup"] = false
		// 		item["IsSalad"] = false
		// 	}

		// }

		/*
			FoodFilter        string  `json:"foodFilter"`    // this is a new field to add to the data // category
			RestaurantName    string  `json:"restaurantName"`
			IsCafe            bool    `json:"isCafe"`
			CafeName          string  `json:"cafeName"`
		*/

		// Convert field names to camelCase
		for key, value := range item {
			camelCaseKey := toCamelCase(key)
			delete(item, key)
			item[camelCaseKey] = value
		}
	}

	// Marshal the modified data back to JSON
	updatedJSON, err := json.MarshalIndent(foodItems, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling JSON: %s", err)
	}

	// Write the updated data to foodmenuchanged.json
	if err := os.WriteFile("foodmenuchanged.json", updatedJSON, 0644); err != nil {
		log.Fatalf("Error writing to file: %s", err)
	}

	fmt.Println("Data successfully updated and saved to foodmenuchanged.json")
}

// Helper function to convert snake_case to camelCase
func toCamelCase(input string) string {
	isToUpper := false
	result := ""

	for i, v := range input {
		if v == '_' {
			isToUpper = true
			continue
		}

		if isToUpper {
			result += strings.ToUpper(string(v))
			isToUpper = false
		} else {
			if i == 0 {
				result += strings.ToLower(string(v))
			} else {
				result += string(v)
			}
		}
	}

	return result
}

/*
restaurant
cafe
homeMade
supermarket
bread
cakeAndSweet
pastaAndNoodles
Rice
salad
soup
hatabBakery
cornFlakes
coffee
milkAndCheese
vegetables
fruit
frozenProducts
chocolate
biscuit
nuts
cannedFood
juice
fatayer
chips
grainAndFlour
protein
appetizers
arabicFood
karazApproved
dates
icecream
saucesAndDips
potatoes
ramadan
*/

// type FoodItem struct {
// 	ID                int     `json:"id"`
// 	DeletedAt         *string `json:"deletedAt"`
// 	NameE             string  `json:"nameE"`
// 	NameA             string  `json:"nameA"`
// 	RecipeE           string  `json:"recipeE"`
// 	RecipeA           string  `json:"recipeA"`
// 	Calories          float64 `json:"calories"`
// 	Color             string  `json:"color"`
// 	CreatedByID       int     `json:"createdById"`
// 	CreatedAt         int64   `json:"createdAt"`
// 	ModifiedByID      int     `json:"modifiedById"`
// 	UpdatedAt         int64   `json:"updatedAt"`
// 	RestaurantID      int     `json:"restaurantId"`
// 	DefaultMediaID    int     `json:"defaultMediaId"`
// 	WeightInGram      float64 `json:"weightInGram"`
// 	ShortDescriptionE string  `json:"shortDescriptionE"`
// 	ShortDescriptionA string  `json:"shortDescriptionA"`
// 	Carbs             float64 `json:"carbs"`
// 	Protein           float64 `json:"protein"`
// 	Fat               float64 `json:"fat"`
// 	IsHomeMade        int     `json:"isHomeMade"`
// 	DietaryFiber      float64 `json:"dietaryFiber"`
// 	TransFat          float64 `json:"transFat"`
// 	SaturatedFat      float64 `json:"saturatedFat"`
// 	Calcium           float64 `json:"calcium"`
// 	Sodium            float64 `json:"sodium"`
// 	Iron              float64 `json:"iron"`
// 	Sugar             float64 `json:"sugar"`
// 	Image             string  `json:"image"`
// 	IsSupermarket     string  `json:"isSupermarket"`
// }
