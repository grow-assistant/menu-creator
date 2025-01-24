package main

import (
	"log"

	"swoop/locations"
	"swoop/pkg/config"
	database "swoop/pkg/db"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Panicln(err)
	}
	log.Println("Seeding Database")
	db := database.Connect(config.DB())
	api := locations.NewAPI(db)
	// seed locations
	log.Println("Seeding *Location* data")
	franklinHillsCountryClub := api.CreateLocation("Franklin Hills Country Club", "Franklin, MI")
	log.Println("Seeding *Menus* data")
	// seed menus
	franklinHillsCountryClubFranklinHillsCountryClub := api.CreateMenu("Franklin Hills Country Club", "Franklin Hills Country Club", franklinHillsCountryClub.ID)
	log.Println(franklinHillsCountryClubFranklinHillsCountryClub)
	// seed categories
	log.Println("Seeding *Categories* data")
	franklinHillsCountryClubFranklinHillsCountryClubBreakfast := api.CreateCategory("Breakfast", "Breakfast", franklinHillsCountryClubFranklinHillsCountryClub.ID)
	franklinHillsCountryClubFranklinHillsCountryClubHotSandwiches := api.CreateCategory("Hot Sandwiches", "Hot Sandwiches", franklinHillsCountryClubFranklinHillsCountryClub.ID)
	franklinHillsCountryClubFranklinHillsCountryClubTacos := api.CreateCategory("Tacos", "Tacos", franklinHillsCountryClubFranklinHillsCountryClub.ID)
	franklinHillsCountryClubFranklinHillsCountryClubColdSandwichesAndSalads := api.CreateCategory("Cold Sandwiches & Salads", "Cold Sandwiches & Salads", franklinHillsCountryClubFranklinHillsCountryClub.ID)
	franklinHillsCountryClubFranklinHillsCountryClubSides := api.CreateCategory("Sides", "Sides", franklinHillsCountryClubFranklinHillsCountryClub.ID)
	franklinHillsCountryClubFranklinHillsCountryClubDessertAndSnacks := api.CreateCategory("Dessert & Snacks", "Dessert & Snacks", franklinHillsCountryClubFranklinHillsCountryClub.ID)
	franklinHillsCountryClubFranklinHillsCountryClubBeverages := api.CreateCategory("Beverages", "Beverages", franklinHillsCountryClubFranklinHillsCountryClub.ID)
	franklinHillsCountryClubFranklinHillsCountryClubAlcoholicBeverages := api.CreateCategory("Alcoholic Beverages", "Alcoholic Beverages", franklinHillsCountryClubFranklinHillsCountryClub.ID)
	log.Println(franklinHillsCountryClubFranklinHillsCountryClubBreakfast)
	log.Println(franklinHillsCountryClubFranklinHillsCountryClubHotSandwiches)
	log.Println(franklinHillsCountryClubFranklinHillsCountryClubTacos)
	log.Println(franklinHillsCountryClubFranklinHillsCountryClubColdSandwichesAndSalads)
	log.Println(franklinHillsCountryClubFranklinHillsCountryClubSides)
	log.Println(franklinHillsCountryClubFranklinHillsCountryClubDessertAndSnacks)
	log.Println(franklinHillsCountryClubFranklinHillsCountryClubBeverages)
	log.Println(franklinHillsCountryClubFranklinHillsCountryClubAlcoholicBeverages)
	// seed item
	log.Println("Seeding *Items* data")
	_ = api.CreateItem("Breakfast Sandwich", "egg, canadian bacon, cheese, english muffin", 10, franklinHillsCountryClubFranklinHillsCountryClubBreakfast.ID)
	_ = api.CreateItem("Vegan Breakfast Sandwich", "plant based sausage patty, tomato, avocado, english muffin", 10, franklinHillsCountryClubFranklinHillsCountryClubBreakfast.ID)
	_ = api.CreateItem("Breakfast Panini", "egg, bacon, tomato, pesto, cheddar cheese, schiacciata bread", 11, franklinHillsCountryClubFranklinHillsCountryClubBreakfast.ID)
	_ = api.CreateItem("Breakfast Tacos", "corn tortillas, egg, tomato, cilantro, queso fresco, avocado, house-made salsa", 11, franklinHillsCountryClubFranklinHillsCountryClubBreakfast.ID)
	_ = api.CreateItem("Bagel with Cream Cheese", "", 4, franklinHillsCountryClubFranklinHillsCountryClubBreakfast.ID)
	_ = api.CreateItem("Hard Boiled Egg", "", 2, franklinHillsCountryClubFranklinHillsCountryClubBreakfast.ID)
	franklinHillsCountryClubBallParkHotDog := api.CreateItem("Ball Park Hot Dog", "", 7, franklinHillsCountryClubFranklinHillsCountryClubHotSandwiches.ID)
	franklinHillsCountryClubKosherHotDog := api.CreateItem("Kosher Hot Dog", "", 9, franklinHillsCountryClubFranklinHillsCountryClubHotSandwiches.ID)
	franklinHillsCountryClubHamburger := api.CreateItem("Hamburger", "lettuce, tomato, onion", 13, franklinHillsCountryClubFranklinHillsCountryClubHotSandwiches.ID)
	franklinHillsCountryClubCheeseburger := api.CreateItem("Cheeseburger", "lettuce, tomato, onion, cheese", 13, franklinHillsCountryClubFranklinHillsCountryClubHotSandwiches.ID)
	franklinHillsCountryClubImpossibleBurger := api.CreateItem("Impossible Burger", "", 13, franklinHillsCountryClubFranklinHillsCountryClubHotSandwiches.ID)
	franklinHillsCountryClubTurkeyBurger := api.CreateItem("Turkey Burger", "", 13, franklinHillsCountryClubFranklinHillsCountryClubHotSandwiches.ID)
	_ = api.CreateItem("Grilled Cheese", "", 5, franklinHillsCountryClubFranklinHillsCountryClubHotSandwiches.ID)
	_ = api.CreateItem("Grilled Tuna Melt", "with mozzarella cheese", 13, franklinHillsCountryClubFranklinHillsCountryClubHotSandwiches.ID)
	_ = api.CreateItem("Grilled Chicken Pita", "hummus, lettuce, tomato, sliced pickles, pita", 13, franklinHillsCountryClubFranklinHillsCountryClubHotSandwiches.ID)
	_ = api.CreateItem("Turkey Panini", "sliced turkey, apple, pesto, cheddar cheese, schiacciata bread", 13, franklinHillsCountryClubFranklinHillsCountryClubHotSandwiches.ID)
	_ = api.CreateItem("Grilled Vegetable Panini", "zucchini, red pepper, asparagus, portobello, hummus, schiacciata bread", 13, franklinHillsCountryClubFranklinHillsCountryClubHotSandwiches.ID)
	_ = api.CreateItem("Chicken Quesadilla Panini", "grilled chicken, tomato, onion, cilantro, cheese, house-made salsa", 13, franklinHillsCountryClubFranklinHillsCountryClubHotSandwiches.ID)
	_ = api.CreateItem("Chicken Tacos", "grilled chicken onions, cilantro, queso fresco, avocado, house-made salsa", 13, franklinHillsCountryClubFranklinHillsCountryClubTacos.ID)
	_ = api.CreateItem("Smoked Brisket Tacos", "smoked burnt ends, onions, cilantro, avocado, queso fresco, house made salsa", 14, franklinHillsCountryClubFranklinHillsCountryClubTacos.ID)
	_ = api.CreateItem("Quinoa Crumble Tacos", "taco seasoned quinoa crumbles, onion, cilantro, avocado, queso fresco, house-made salsa", 12, franklinHillsCountryClubFranklinHillsCountryClubTacos.ID)
	franklinHillsCountryClubTunaorChickenSalad := api.CreateItem("Tuna or Chicken Salad", "", 12, franklinHillsCountryClubFranklinHillsCountryClubColdSandwichesAndSalads.ID)
	_ = api.CreateItem("Egg Salad", "", 9, franklinHillsCountryClubFranklinHillsCountryClubColdSandwichesAndSalads.ID)
	_ = api.CreateItem("Sliced Turkey", "", 12, franklinHillsCountryClubFranklinHillsCountryClubColdSandwichesAndSalads.ID)
	_ = api.CreateItem("Peanut Butter & Jelly", "", 7, franklinHillsCountryClubFranklinHillsCountryClubColdSandwichesAndSalads.ID)
	_ = api.CreateItem("Carrot & Celery Sticks", "with hummus", 4, franklinHillsCountryClubFranklinHillsCountryClubSides.ID)
	_ = api.CreateItem("Chicken Noodle Soup", "", 6, franklinHillsCountryClubFranklinHillsCountryClubSides.ID)
	_ = api.CreateItem("Soup du Jour", "", 6, franklinHillsCountryClubFranklinHillsCountryClubSides.ID)
	_ = api.CreateItem("Frozen Grapes", "", 3, franklinHillsCountryClubFranklinHillsCountryClubSides.ID)
	_ = api.CreateItem("Jerky", "", 6, franklinHillsCountryClubFranklinHillsCountryClubSides.ID)
	_ = api.CreateItem("1st & 10th Tee Bars", "", 6, franklinHillsCountryClubFranklinHillsCountryClubSides.ID)
	_ = api.CreateItem("Fresh Fruit Cup", "", 5, franklinHillsCountryClubFranklinHillsCountryClubSides.ID)
	_ = api.CreateItem("Ice Cream Bars", "", 4, franklinHillsCountryClubFranklinHillsCountryClubDessertAndSnacks.ID)
	_ = api.CreateItem("Italian Ices", "", 4, franklinHillsCountryClubFranklinHillsCountryClubDessertAndSnacks.ID)
	_ = api.CreateItem("Cookies", "", 2, franklinHillsCountryClubFranklinHillsCountryClubDessertAndSnacks.ID)
	_ = api.CreateItem("Rice Krispy Treat", "", 2, franklinHillsCountryClubFranklinHillsCountryClubDessertAndSnacks.ID)
	_ = api.CreateItem("Chips or Pretzels", "", 2, franklinHillsCountryClubFranklinHillsCountryClubDessertAndSnacks.ID)
	_ = api.CreateItem("Candy Bars", "", 2, franklinHillsCountryClubFranklinHillsCountryClubDessertAndSnacks.ID)
	_ = api.CreateItem("Soft Drinks", "", 3, franklinHillsCountryClubFranklinHillsCountryClubBeverages.ID)
	franklinHillsCountryClubGatorade := api.CreateItem("Gatorade", "", 4, franklinHillsCountryClubFranklinHillsCountryClubBeverages.ID)
	_ = api.CreateItem("Juices", "", 3, franklinHillsCountryClubFranklinHillsCountryClubBeverages.ID)
	_ = api.CreateItem("Iced Tea", "", 3, franklinHillsCountryClubFranklinHillsCountryClubBeverages.ID)
	_ = api.CreateItem("Lemonade", "tea-lade", 3, franklinHillsCountryClubFranklinHillsCountryClubBeverages.ID)
	franklinHillsCountryClubCoffee := api.CreateItem("Coffee", "", 3, franklinHillsCountryClubFranklinHillsCountryClubBeverages.ID)
	_ = api.CreateItem("Milk", "", 3, franklinHillsCountryClubFranklinHillsCountryClubBeverages.ID)
	_ = api.CreateItem("Perrier", "", 4, franklinHillsCountryClubFranklinHillsCountryClubBeverages.ID)
	franklinHillsCountryClubBeer := api.CreateItem("Beer", "", 7, franklinHillsCountryClubFranklinHillsCountryClubAlcoholicBeverages.ID)
	franklinHillsCountryClubHighNoon := api.CreateItem("High Noon", "", 7, franklinHillsCountryClubFranklinHillsCountryClubAlcoholicBeverages.ID)
	franklinHillsCountryClubBloodyMary := api.CreateItem("Bloody Mary", "vodka, house bloody mary mix, fresh celery, fresh-grated horseradish, rosemary skewered olive, cherry tomato, spicy pepper", 15, franklinHillsCountryClubFranklinHillsCountryClubAlcoholicBeverages.ID)
	franklinHillsCountryClubTransfusion := api.CreateItem("Transfusion", "vodka, ginger ale and grape juice", 15, franklinHillsCountryClubFranklinHillsCountryClubAlcoholicBeverages.ID)
	_ = api.CreateItem("Tito’s Vodka", "", 14, franklinHillsCountryClubFranklinHillsCountryClubAlcoholicBeverages.ID)
	_ = api.CreateItem("Don Julio Tequila", "", 25, franklinHillsCountryClubFranklinHillsCountryClubAlcoholicBeverages.ID)
	_ = api.CreateItem("Jefferson Reserve Bourbon", "", 25, franklinHillsCountryClubFranklinHillsCountryClubAlcoholicBeverages.ID)
	_ = api.CreateItem("Bombay Sapphire Gin", "", 15, franklinHillsCountryClubFranklinHillsCountryClubAlcoholicBeverages.ID)
	_ = api.CreateItem("Glenlivet 12 Yr", "", 22, franklinHillsCountryClubFranklinHillsCountryClubAlcoholicBeverages.ID)
	_ = api.CreateItem("Chandon Brut", "", 11, franklinHillsCountryClubFranklinHillsCountryClubAlcoholicBeverages.ID)
	_ = api.CreateItem("Rose", "", 11, franklinHillsCountryClubFranklinHillsCountryClubAlcoholicBeverages.ID)
	_ = api.CreateItem("Sauvignon Blanc", "", 11, franklinHillsCountryClubFranklinHillsCountryClubAlcoholicBeverages.ID)
	_ = api.CreateItem("Sangria", "", 11, franklinHillsCountryClubFranklinHillsCountryClubAlcoholicBeverages.ID)
	// seed item options
	log.Println("Seeding *Options* data")
	_ = api.CreateOption("Options", "Options", 0, 0, franklinHillsCountryClubBallParkHotDog.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Ketchup", Description: "Ketchup", Price: 0},
		locations.OptionItem{Name: "Mustard", Description: "Mustard", Price: 0},
		locations.OptionItem{Name: "Relish", Description: "Relish", Price: 0},
		locations.OptionItem{Name: "Onions", Description: "Onions", Price: 0}})
	_ = api.CreateOption("Options", "Options", 0, 0, franklinHillsCountryClubKosherHotDog.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Ketchup", Description: "Ketchup", Price: 0},
		locations.OptionItem{Name: "Mustard", Description: "Mustard", Price: 0},
		locations.OptionItem{Name: "Relish", Description: "Relish", Price: 0},
		locations.OptionItem{Name: "Onions", Description: "Onions", Price: 0}})
	_ = api.CreateOption("Meat Temperature", "Meat Temperature", 1, 1, franklinHillsCountryClubHamburger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Well Done", Description: "Well Done", Price: 0},
		locations.OptionItem{Name: "Medium-Well", Description: "Medium-Well", Price: 0},
		locations.OptionItem{Name: "Medium", Description: "Medium", Price: 0},
		locations.OptionItem{Name: "Medium-Rare", Description: "Medium-Rare", Price: 0}})
	_ = api.CreateOption("Options", "Options", 0, 0, franklinHillsCountryClubHamburger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Mayonnaise", Description: "Mayonnaise", Price: 0},
		locations.OptionItem{Name: "Mustard", Description: "Mustard", Price: 0},
		locations.OptionItem{Name: "Ketchup", Description: "Ketchup", Price: 0}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, franklinHillsCountryClubHamburger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Lettuce", Description: "No Lettuce", Price: 0},
		locations.OptionItem{Name: "No Tomato", Description: "No Tomato", Price: 0},
		locations.OptionItem{Name: "No Onion", Description: "No Onion", Price: 0}})
	_ = api.CreateOption("Meat Temperature", "Meat Temperature", 1, 1, franklinHillsCountryClubCheeseburger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Well Done", Description: "Well Done", Price: 0},
		locations.OptionItem{Name: "Medium-Well", Description: "Medium-Well", Price: 0},
		locations.OptionItem{Name: "Medium", Description: "Medium", Price: 0},
		locations.OptionItem{Name: "Medium-Rare", Description: "Medium-Rare", Price: 0}})
	_ = api.CreateOption("Cheese", "Cheese", 0, 0, franklinHillsCountryClubCheeseburger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "American Cheese", Description: "American Cheese", Price: 0},
		locations.OptionItem{Name: "Cheddar Jack", Description: "Cheddar Jack", Price: 0},
		locations.OptionItem{Name: "Cheddar Cheese", Description: "Cheddar Cheese", Price: 0}})
	_ = api.CreateOption("Options", "Options", 0, 0, franklinHillsCountryClubCheeseburger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Mayonnaise", Description: "Mayonnaise", Price: 0},
		locations.OptionItem{Name: "Mustard", Description: "Mustard", Price: 0},
		locations.OptionItem{Name: "Ketchup", Description: "Ketchup", Price: 0}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, franklinHillsCountryClubCheeseburger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Lettuce", Description: "No Lettuce", Price: 0},
		locations.OptionItem{Name: "No Tomato", Description: "No Tomato", Price: 0},
		locations.OptionItem{Name: "No Onion", Description: "No Onion", Price: 0},
		locations.OptionItem{Name: "No Cheese", Description: "No Cheese", Price: 0}})
	_ = api.CreateOption("Options", "Options", 0, 0, franklinHillsCountryClubImpossibleBurger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Mayonnaise", Description: "Mayonnaise", Price: 0},
		locations.OptionItem{Name: "Mustard", Description: "Mustard", Price: 0},
		locations.OptionItem{Name: "Ketchup", Description: "Ketchup", Price: 0}})
	_ = api.CreateOption("Options", "Options", 0, 0, franklinHillsCountryClubTurkeyBurger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Mayonnaise", Description: "Mayonnaise", Price: 0},
		locations.OptionItem{Name: "Mustard", Description: "Mustard", Price: 0},
		locations.OptionItem{Name: "Ketchup", Description: "Ketchup", Price: 0}})
	_ = api.CreateOption("Bread", "Bread", 1, 1, franklinHillsCountryClubTunaorChickenSalad.ID, []locations.OptionItem{
		locations.OptionItem{Name: "White", Description: "White", Price: 0},
		locations.OptionItem{Name: "Wheat", Description: "Wheat", Price: 0},
		locations.OptionItem{Name: "Rye", Description: "Rye", Price: 0}})
	_ = api.CreateOption("Options", "Options", 1, 1, franklinHillsCountryClubTunaorChickenSalad.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Chicken Salad", Description: "Chicken Salad", Price: 0},
		locations.OptionItem{Name: "Tuna Salad", Description: "Tuna Salad", Price: 0}})
	_ = api.CreateOption("Flavor", "Flavor", 1, 1, franklinHillsCountryClubGatorade.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Lemon-Lime", Description: "Lemon-Lime", Price: 0},
		locations.OptionItem{Name: "Fruit Punch", Description: "Fruit Punch", Price: 0},
		locations.OptionItem{Name: "Orange", Description: "Orange", Price: 0},
		locations.OptionItem{Name: "Glacier Freeze", Description: "Glacier Freeze", Price: 0},
		locations.OptionItem{Name: "Cool Blue", Description: "Cool Blue", Price: 0}})
	_ = api.CreateOption("Extras", "Extras", 0, 0, franklinHillsCountryClubCoffee.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Cream", Description: "Cream", Price: 0},
		locations.OptionItem{Name: "Sugar", Description: "Sugar", Price: 0},
		locations.OptionItem{Name: "Splenda", Description: "Splenda", Price: 0}})
	_ = api.CreateOption("Options", "Options", 1, 1, franklinHillsCountryClubBeer.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Molson Gold", Description: "Molson Gold", Price: 0},
		locations.OptionItem{Name: "Ultra", Description: "Ultra", Price: 0},
		locations.OptionItem{Name: "Miller Light", Description: "Miller Light", Price: 0},
		locations.OptionItem{Name: "Heineken", Description: "Heineken", Price: 0},
		locations.OptionItem{Name: "Corona", Description: "Corona", Price: 0}})
	_ = api.CreateOption("Options", "Options", 1, 1, franklinHillsCountryClubHighNoon.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Black Cherry", Description: "Black Cherry", Price: 0},
		locations.OptionItem{Name: "Ruby Grapefruit", Description: "Ruby Grapefruit", Price: 0},
		locations.OptionItem{Name: "Natural Lime", Description: "Natural Lime", Price: 0},
		locations.OptionItem{Name: "Raspberry", Description: "Raspberry", Price: 0},
		locations.OptionItem{Name: "Mango", Description: "Mango", Price: 0}})
	_ = api.CreateOption("Vodka", "Vodka", 1, 1, franklinHillsCountryClubBloodyMary.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Grey Goose", Description: "Grey Goose", Price: 4},
		locations.OptionItem{Name: "Titos", Description: "Titos", Price: 0},
		locations.OptionItem{Name: "Effen", Description: "Effen", Price: 0},
		locations.OptionItem{Name: "Absolut", Description: "Absolut", Price: 0},
		locations.OptionItem{Name: "Ketel One", Description: "Ketel One", Price: 0},
		locations.OptionItem{Name: "Stoli", Description: "Stoli", Price: 0}})
	_ = api.CreateOption("Vodka", "Vodka", 1, 1, franklinHillsCountryClubTransfusion.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Grey Goose", Description: "Grey Goose", Price: 4},
		locations.OptionItem{Name: "Titos", Description: "Titos", Price: 0},
		locations.OptionItem{Name: "Effen", Description: "Effen", Price: 0},
		locations.OptionItem{Name: "Absolut", Description: "Absolut", Price: 0},
		locations.OptionItem{Name: "Ketel One", Description: "Ketel One", Price: 0},
		locations.OptionItem{Name: "Stoli", Description: "Stoli", Price: 0}})
	// seed orders
	log.Println("Seeding Completed")
}