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
	pinetreePool := api.CreateLocation("Pinetree - Pool", "Kennesaw, GA")
	log.Println("Seeding *Menus* data")
	// seed menus
	pinetreePoolPoolMenu := api.CreateMenu("Pool Menu", "Pool Menu", pinetreePool.ID)
	log.Println(pinetreePoolPoolMenu)
	// seed categories
	log.Println("Seeding *Categories* data")
	pinetreePoolPoolMenuShareables := api.CreateCategory("Shareables", "Shareables", pinetreePoolPoolMenu.ID)
	pinetreePoolPoolMenuSalads := api.CreateCategory("Salads", "Salads", pinetreePoolPoolMenu.ID)
	pinetreePoolPoolMenuHandhelds := api.CreateCategory("Handhelds", "Handhelds", pinetreePoolPoolMenu.ID)
	pinetreePoolPoolMenuCondiments := api.CreateCategory("Condiments", "Condiments", pinetreePoolPoolMenu.ID)
	log.Println(pinetreePoolPoolMenuShareables)
	log.Println(pinetreePoolPoolMenuSalads)
	log.Println(pinetreePoolPoolMenuHandhelds)
	log.Println(pinetreePoolPoolMenuCondiments)
	// seed item
	log.Println("Seeding *Items* data")
	_ = api.CreateItem("Chips & Queso", "", 8, pinetreePoolPoolMenuShareables.ID)
	_ = api.CreateItem("Chips & Salsa", "", 6, pinetreePoolPoolMenuShareables.ID)
	_ = api.CreateItem("Chips & Guacamole", "", 8, pinetreePoolPoolMenuShareables.ID)
	pinetreePoolNachos := api.CreateItem("Nachos", "chicken or steak, cheddar-jack cheese, fresh pico de gallo, sour cream, and pickled jalapenos", 13, pinetreePoolPoolMenuShareables.ID)
	pinetreePoolBasketofFries := api.CreateItem("Basket of Fries", "", 7, pinetreePoolPoolMenuShareables.ID)
	_ = api.CreateItem("Fresh Fruit", "", 8, pinetreePoolPoolMenuShareables.ID)
	pinetreePoolChefsSalad := api.CreateItem("Chef's Salad", "ham, turkey, bacon, egg, tomato, cucumber, onion, mushrooms, croutons, cheddar jack", 16, pinetreePoolPoolMenuSalads.ID)
	pinetreePoolGardenSalad := api.CreateItem("Garden Salad ", "lettuce, tomato, cucumber, jack cheese, and garlic croutons", 9, pinetreePoolPoolMenuSalads.ID)
	pinetreePoolCheeseQuesadilla := api.CreateItem("Cheese Quesadilla", "with cheddar-jack cheese, garnished with fresh pico de gallo, sour cream, and pickled jalapenos", 7, pinetreePoolPoolMenuHandhelds.ID)
	pinetreePoolChickenQuesadilla := api.CreateItem("Chicken Quesadilla", "with peppers, onions, cheddar-jack cheese, garnished with fresh pico de gallo, sour cream, and pickled jalapenos", 13, pinetreePoolPoolMenuHandhelds.ID)
	pinetreePoolSteakQuesadilla := api.CreateItem("Steak Quesadilla", "with peppers, onions, cheddar-jack cheese, garnished with fresh pico de gallo, sour cream, and pickled jalapenos", 15, pinetreePoolPoolMenuHandhelds.ID)
	pinetreePoolAllBeefHotDog := api.CreateItem("All Beef Hot Dog", "classic beef hot dog topped with ketchup, mustard, relish, and/or onions", 8, pinetreePoolPoolMenuHandhelds.ID)
	pinetreePoolChickenCaesarWrap := api.CreateItem("Chicken Caesar Wrap", "grilled, blackened or fried chicken, romaine lettuce, parmesan and caesar dressing", 13, pinetreePoolPoolMenuHandhelds.ID)
	pinetreePoolHamburger := api.CreateItem("Hamburger", "lettuce, tomato, onion, cheese", 13, pinetreePoolPoolMenuHandhelds.ID)
	pinetreePoolChickenSandwich := api.CreateItem("Chicken Sandwich", "grilled or fried chicken breast with provolone, lettuce, tomato & onion. add bacon $2", 12, pinetreePoolPoolMenuHandhelds.ID)
	pinetreePoolTripleGrilledCheeseSandwich := api.CreateItem("Triple Grilled Cheese Sandwich", "", 11, pinetreePoolPoolMenuHandhelds.ID)
	pinetreePoolChickenorTunaSalad := api.CreateItem("Chicken or Tuna Salad", "shredded lettuce, tomato, served as a wrap or your choice of bread", 12, pinetreePoolPoolMenuHandhelds.ID)
	pinetreePoolBlackenedSalmonWrap := api.CreateItem("Blackened Salmon Wrap", "blackened salmon, romaine lettuce, parmesan and caesar dressing", 13, pinetreePoolPoolMenuHandhelds.ID)
	_ = api.CreateItem("Balsamic Vinaigrette", "", 2, pinetreePoolPoolMenuCondiments.ID)
	_ = api.CreateItem("Blue Cheese Dressing", "", 2, pinetreePoolPoolMenuCondiments.ID)
	_ = api.CreateItem("Honey Mustard", "", 2, pinetreePoolPoolMenuCondiments.ID)
	_ = api.CreateItem("Ranch", "", 2, pinetreePoolPoolMenuCondiments.ID)
	_ = api.CreateItem("Thousand Island", "", 2, pinetreePoolPoolMenuCondiments.ID)
	_ = api.CreateItem("Italian Dressing", "", 2, pinetreePoolPoolMenuCondiments.ID)
	_ = api.CreateItem("Ketchup", "", 2, pinetreePoolPoolMenuCondiments.ID)
	_ = api.CreateItem("BBQ Sauce", "", 2, pinetreePoolPoolMenuCondiments.ID)
	_ = api.CreateItem("Mayo", "", 2, pinetreePoolPoolMenuCondiments.ID)
	_ = api.CreateItem("Mustard", "", 2, pinetreePoolPoolMenuCondiments.ID)
	// seed item options
	log.Println("Seeding *Options* data")
	_ = api.CreateOption("Extras", "Extras", 1, 1, pinetreePoolNachos.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Chicken", Description: "Chicken", Price: 0},
		locations.OptionItem{Name: "Steak", Description: "Steak", Price: 2},
		locations.OptionItem{Name: "No Protein", Description: "No Protein", Price: 0}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, pinetreePoolNachos.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Jalapenos", Description: "No Jalapenos", Price: 0},
		locations.OptionItem{Name: "No Pico", Description: "No Pico", Price: 0},
		locations.OptionItem{Name: "No Sour Cream", Description: "No Sour Cream", Price: 0}})
	_ = api.CreateOption("Options", "Options", 0, 0, pinetreePoolBasketofFries.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Ketchup", Description: "Ketchup", Price: 0},
		locations.OptionItem{Name: "Ranch", Description: "Ranch", Price: 0}})
	_ = api.CreateOption("Salad Dressing", "Salad Dressing", 1, 1, pinetreePoolChefsSalad.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Balsamic Vinaigrette", Description: "Balsamic Vinaigrette", Price: 0},
		locations.OptionItem{Name: "Blue Cheese", Description: "Blue Cheese", Price: 0},
		locations.OptionItem{Name: "Honey Mustard", Description: "Honey Mustard", Price: 0},
		locations.OptionItem{Name: "Ranch", Description: "Ranch", Price: 0},
		locations.OptionItem{Name: "Thousand Island", Description: "Thousand Island", Price: 0},
		locations.OptionItem{Name: "Italian", Description: "Italian", Price: 0}})
	_ = api.CreateOption("Add Protein", "Add Protein", 0, 0, pinetreePoolChefsSalad.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Grilled Chicken", Description: "Grilled Chicken", Price: 6},
		locations.OptionItem{Name: "Fried Chicken", Description: "Fried Chicken", Price: 6},
		locations.OptionItem{Name: "Grilled Salmon", Description: "Grilled Salmon", Price: 8}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, pinetreePoolChefsSalad.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Ham", Description: "No Ham", Price: 0},
		locations.OptionItem{Name: "No Turkey", Description: "No Turkey", Price: 0},
		locations.OptionItem{Name: "No Bacon", Description: "No Bacon", Price: 0},
		locations.OptionItem{Name: "No Egg", Description: "No Egg", Price: 0},
		locations.OptionItem{Name: "No Tomatoes", Description: "No Tomatoes", Price: 0},
		locations.OptionItem{Name: "No Cucumbers", Description: "No Cucumbers", Price: 0},
		locations.OptionItem{Name: "No Croutons", Description: "No Croutons", Price: 0},
		locations.OptionItem{Name: "No Onions", Description: "No Onions", Price: 0}})
	_ = api.CreateOption("Salad Dressing", "Salad Dressing", 1, 1, pinetreePoolGardenSalad.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Balsamic Vinaigrette", Description: "Balsamic Vinaigrette", Price: 0},
		locations.OptionItem{Name: "Blue Cheese", Description: "Blue Cheese", Price: 0},
		locations.OptionItem{Name: "Honey Mustard", Description: "Honey Mustard", Price: 0},
		locations.OptionItem{Name: "Ranch", Description: "Ranch", Price: 0},
		locations.OptionItem{Name: "Thousand Island", Description: "Thousand Island", Price: 0},
		locations.OptionItem{Name: "Italian", Description: "Italian", Price: 0}})
	_ = api.CreateOption("Add Protein", "Add Protein", 0, 0, pinetreePoolGardenSalad.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Grilled Chicken", Description: "Grilled Chicken", Price: 6},
		locations.OptionItem{Name: "Fried Chicken", Description: "Fried Chicken", Price: 6},
		locations.OptionItem{Name: "Grilled Salmon", Description: "Grilled Salmon", Price: 8}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, pinetreePoolGardenSalad.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Tomatoes", Description: "No Tomatoes", Price: 0},
		locations.OptionItem{Name: "No Cucumbers", Description: "No Cucumbers", Price: 0},
		locations.OptionItem{Name: "No Croutons", Description: "No Croutons", Price: 0}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, pinetreePoolCheeseQuesadilla.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Sour Cream", Description: "No Sour Cream", Price: 0},
		locations.OptionItem{Name: "No Jalapenos", Description: "No Jalapenos", Price: 0},
		locations.OptionItem{Name: "No Pico de Gallo", Description: "No Pico de Gallo", Price: 0}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, pinetreePoolChickenQuesadilla.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Peppers & Onions", Description: "No Peppers & Onions", Price: 0},
		locations.OptionItem{Name: "No Sour Cream", Description: "No Sour Cream", Price: 0},
		locations.OptionItem{Name: "No Jalapenos", Description: "No Jalapenos", Price: 0},
		locations.OptionItem{Name: "No Pico de Gallo", Description: "No Pico de Gallo", Price: 0}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, pinetreePoolSteakQuesadilla.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Peppers & Onions", Description: "No Peppers & Onions", Price: 0},
		locations.OptionItem{Name: "No Sour Cream", Description: "No Sour Cream", Price: 0},
		locations.OptionItem{Name: "No Jalapenos", Description: "No Jalapenos", Price: 0},
		locations.OptionItem{Name: "No Pico de Gallo", Description: "No Pico de Gallo", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, pinetreePoolAllBeefHotDog.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Fresh Fruit", Description: "Fresh Fruit", Price: 0},
		locations.OptionItem{Name: "Tater Tots", Description: "Tater Tots", Price: 0},
		locations.OptionItem{Name: "Broccoli Raisin Salad", Description: "Broccoli Raisin Salad", Price: 0},
		locations.OptionItem{Name: "Pasta Salad", Description: "Pasta Salad", Price: 0}})
	_ = api.CreateOption("Options", "Options", 0, 0, pinetreePoolAllBeefHotDog.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Ketchup", Description: "Ketchup", Price: 0},
		locations.OptionItem{Name: "Mustard", Description: "Mustard", Price: 0},
		locations.OptionItem{Name: "Relish", Description: "Relish", Price: 0},
		locations.OptionItem{Name: "Onions", Description: "Onions", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, pinetreePoolChickenCaesarWrap.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Fresh Fruit", Description: "Fresh Fruit", Price: 0},
		locations.OptionItem{Name: "Tater Tots", Description: "Tater Tots", Price: 0},
		locations.OptionItem{Name: "Kettle Chips", Description: "Kettle Chips", Price: 0},
		locations.OptionItem{Name: "Broccoli Raisin Salad", Description: "Broccoli Raisin Salad", Price: 0},
		locations.OptionItem{Name: "Pasta Salad", Description: "Pasta Salad", Price: 0},
		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0}})
	_ = api.CreateOption("Options", "Options", 1, 1, pinetreePoolChickenCaesarWrap.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Grilled Chicken", Description: "Grilled Chicken", Price: 0},
		locations.OptionItem{Name: "Fried Chicken", Description: "Fried Chicken", Price: 0},
		locations.OptionItem{Name: "Blackened Chicken", Description: "Blackened Chicken", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, pinetreePoolChickenCaesarWrap.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Fresh Fruit", Description: "Fresh Fruit", Price: 0},
		locations.OptionItem{Name: "Tater Tots", Description: "Tater Tots", Price: 0},
		locations.OptionItem{Name: "Broccoli Raisin Salad", Description: "Broccoli Raisin Salad", Price: 0},
		locations.OptionItem{Name: "Pasta Salad", Description: "Pasta Salad", Price: 0}})
	_ = api.CreateOption("Meat Temperature", "Meat Temperature", 1, 1, pinetreePoolHamburger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Well Done", Description: "Well Done", Price: 0},
		locations.OptionItem{Name: "Medium-Well", Description: "Medium-Well", Price: 0},
		locations.OptionItem{Name: "Medium", Description: "Medium", Price: 0},
		locations.OptionItem{Name: "Medium-Rare", Description: "Medium-Rare", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, pinetreePoolHamburger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Fresh Fruit", Description: "Fresh Fruit", Price: 0},
		locations.OptionItem{Name: "Tater Tots", Description: "Tater Tots", Price: 0},
		locations.OptionItem{Name: "Broccoli Raisin Salad", Description: "Broccoli Raisin Salad", Price: 0},
		locations.OptionItem{Name: "Pasta Salad", Description: "Pasta Salad", Price: 0}})
	_ = api.CreateOption("Extras", "Extras", 0, 0, pinetreePoolHamburger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Bacon", Description: "Bacon", Price: 2}})
	_ = api.CreateOption("Cheese", "Cheese", 0, 0, pinetreePoolHamburger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "American Cheese", Description: "American Cheese", Price: 0},
		locations.OptionItem{Name: "Cheddar Jack", Description: "Cheddar Jack", Price: 0},
		locations.OptionItem{Name: "Cheddar Cheese", Description: "Cheddar Cheese", Price: 0}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, pinetreePoolHamburger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Onions", Description: "No Onions", Price: 0},
		locations.OptionItem{Name: "No Tomatoes", Description: "No Tomatoes", Price: 0},
		locations.OptionItem{Name: "No Onions", Description: "No Onions", Price: 0}})
	_ = api.CreateOption("Options", "Options", 1, 1, pinetreePoolChickenSandwich.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Grilled Chicken", Description: "Grilled Chicken", Price: 0},
		locations.OptionItem{Name: "Fried Chicken", Description: "Fried Chicken", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, pinetreePoolChickenSandwich.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Fresh Fruit", Description: "Fresh Fruit", Price: 0},
		locations.OptionItem{Name: "Tater Tots", Description: "Tater Tots", Price: 0},
		locations.OptionItem{Name: "Broccoli Raisin Salad", Description: "Broccoli Raisin Salad", Price: 0},
		locations.OptionItem{Name: "Pasta Salad", Description: "Pasta Salad", Price: 0}})
	_ = api.CreateOption("Extras", "Extras", 0, 0, pinetreePoolChickenSandwich.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Add Bacon", Description: "Add Bacon", Price: 2}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, pinetreePoolChickenSandwich.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Lettuce", Description: "No Lettuce", Price: 0},
		locations.OptionItem{Name: "No Tomato", Description: "No Tomato", Price: 0},
		locations.OptionItem{Name: "No Onion", Description: "No Onion", Price: 0},
		locations.OptionItem{Name: "No Cheese", Description: "No Cheese", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, pinetreePoolTripleGrilledCheeseSandwich.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Fresh Fruit", Description: "Fresh Fruit", Price: 0},
		locations.OptionItem{Name: "Tater Tots", Description: "Tater Tots", Price: 0},
		locations.OptionItem{Name: "Broccoli Raisin Salad", Description: "Broccoli Raisin Salad", Price: 0},
		locations.OptionItem{Name: "Pasta Salad", Description: "Pasta Salad", Price: 0}})
	_ = api.CreateOption("Bread", "Bread", 1, 1, pinetreePoolChickenorTunaSalad.ID, []locations.OptionItem{
		locations.OptionItem{Name: "White", Description: "White", Price: 0},
		locations.OptionItem{Name: "Wheat", Description: "Wheat", Price: 0},
		locations.OptionItem{Name: "Rye", Description: "Rye", Price: 0},
		locations.OptionItem{Name: "Wrap", Description: "Wrap", Price: 0},
		locations.OptionItem{Name: "In a Cup", Description: "In a Cup", Price: 0}})
	_ = api.CreateOption("Options", "Options", 1, 1, pinetreePoolChickenorTunaSalad.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Tuna", Description: "Tuna", Price: 0},
		locations.OptionItem{Name: "Chicken Salad", Description: "Chicken Salad", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, pinetreePoolChickenorTunaSalad.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Fresh Fruit", Description: "Fresh Fruit", Price: 0},
		locations.OptionItem{Name: "Tater Tots", Description: "Tater Tots", Price: 0},
		locations.OptionItem{Name: "Broccoli Raisin Salad", Description: "Broccoli Raisin Salad", Price: 0},
		locations.OptionItem{Name: "Pasta Salad", Description: "Pasta Salad", Price: 0}})
	_ = api.CreateOption("Cheese", "Cheese", 1, 1, pinetreePoolChickenorTunaSalad.ID, []locations.OptionItem{
		locations.OptionItem{Name: "American Cheese", Description: "American Cheese", Price: 0},
		locations.OptionItem{Name: "Cheddar Jack", Description: "Cheddar Jack", Price: 0},
		locations.OptionItem{Name: "Cheddar Cheese", Description: "Cheddar Cheese", Price: 0},
		locations.OptionItem{Name: "Provolone", Description: "Provolone", Price: 0},
		locations.OptionItem{Name: "No Cheese", Description: "No Cheese", Price: 0}})
	_ = api.CreateOption("Additional Options", "Additional Options", 0, 0, pinetreePoolChickenorTunaSalad.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Toasted Bread", Description: "Toasted Bread", Price: 0}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, pinetreePoolChickenorTunaSalad.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Tomato", Description: "No Tomato", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, pinetreePoolBlackenedSalmonWrap.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Fresh Fruit", Description: "Fresh Fruit", Price: 0},
		locations.OptionItem{Name: "Tater Tots", Description: "Tater Tots", Price: 0},
		locations.OptionItem{Name: "Kettle Chips", Description: "Kettle Chips", Price: 0},
		locations.OptionItem{Name: "Broccoli Raisin Salad", Description: "Broccoli Raisin Salad", Price: 0},
		locations.OptionItem{Name: "Pasta Salad", Description: "Pasta Salad", Price: 0},
		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0}})
	_ = api.CreateOption("Temperature", "Temperature", 1, 1, pinetreePoolBlackenedSalmonWrap.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Well Done", Description: "Well Done", Price: 0},
		locations.OptionItem{Name: "Medium-Well", Description: "Medium-Well", Price: 0},
		locations.OptionItem{Name: "Medium", Description: "Medium", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, pinetreePoolBlackenedSalmonWrap.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Fresh Fruit", Description: "Fresh Fruit", Price: 0},
		locations.OptionItem{Name: "Tater Tots", Description: "Tater Tots", Price: 0},
		locations.OptionItem{Name: "Broccoli Raisin Salad", Description: "Broccoli Raisin Salad", Price: 0},
		locations.OptionItem{Name: "Pasta Salad", Description: "Pasta Salad", Price: 0}})
	// seed orders
	log.Println("Seeding Completed")
}