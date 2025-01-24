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
	canterburyGolf := api.CreateLocation("Canterbury Golf", "Beachwood, OH")
	log.Println("Seeding *Menus* data")
	// seed menus
	canterburyGolfCanterburyGolf := api.CreateMenu("Canterbury Golf", "Canterbury Golf", canterburyGolf.ID)
	log.Println(canterburyGolfCanterburyGolf)
	// seed categories
	log.Println("Seeding *Categories* data")
	canterburyGolfCanterburyGolfStarters := api.CreateCategory("Starters", "Starters", canterburyGolfCanterburyGolf.ID)
	canterburyGolfCanterburyGolfSalads := api.CreateCategory("Salads", "Salads", canterburyGolfCanterburyGolf.ID)
	canterburyGolfCanterburyGolfSandwiches := api.CreateCategory("Sandwiches", "Sandwiches", canterburyGolfCanterburyGolf.ID)
	canterburyGolfCanterburyGolfMainCourses := api.CreateCategory("Main Courses", "Main Courses", canterburyGolfCanterburyGolf.ID)
	canterburyGolfCanterburyGolfDesserts := api.CreateCategory("Desserts", "Desserts", canterburyGolfCanterburyGolf.ID)
	canterburyGolfCanterburyGolfBeer := api.CreateCategory("Beer", "Beer", canterburyGolfCanterburyGolf.ID)
	canterburyGolfCanterburyGolfCocktails := api.CreateCategory("Cocktails", "Cocktails", canterburyGolfCanterburyGolf.ID)
	log.Println(canterburyGolfCanterburyGolfStarters)
	log.Println(canterburyGolfCanterburyGolfSalads)
	log.Println(canterburyGolfCanterburyGolfSandwiches)
	log.Println(canterburyGolfCanterburyGolfMainCourses)
	log.Println(canterburyGolfCanterburyGolfDesserts)
	log.Println(canterburyGolfCanterburyGolfBeer)
	log.Println(canterburyGolfCanterburyGolfCocktails)
	// seed item
	log.Println("Seeding *Items* data")
	canterburyGolfChickenWings := api.CreateItem("Chicken Wings", "roasted hot peppers sofrito, blue cheese", 14, canterburyGolfCanterburyGolfStarters.ID)
	_ = api.CreateItem("Veal Meatballs", "tomato italian sausage sauce, ohio polenta", 9, canterburyGolfCanterburyGolfStarters.ID)
	_ = api.CreateItem("Crispy Calamari", "sweet chili sauce, peanuts, cilantro, scallions, sesame seeds, lime", 14, canterburyGolfCanterburyGolfStarters.ID)
	_ = api.CreateItem("Parmesan Brussel Sprouts", "crumbled bacon, truffle honey", 9, canterburyGolfCanterburyGolfStarters.ID)
	_ = api.CreateItem("Charcuterie Board for Two", "", 24, canterburyGolfCanterburyGolfStarters.ID)
	canterburyGolfCaesar := api.CreateItem("Caesar", "black pepper parmesan crouton", 9, canterburyGolfCanterburyGolfSalads.ID)
	canterburyGolfItalianChop := api.CreateItem("Italian Chop", "mixed greens, cucumber, tomato, salami, sweet peppers, olives, peperoncini, chickpeas, provolone", 9, canterburyGolfCanterburyGolfSalads.ID)
	canterburyGolfStrawberryGoatCheese := api.CreateItem("Strawberry Goat Cheese", "spiced pecans, raisins, radish, mixed greens, white balsamic dressing", 14, canterburyGolfCanterburyGolfSalads.ID)
	canterburyGolfIceberg := api.CreateItem("Iceberg", "bacon, tomato, egg, blue cheese, pickled beets, dried cherries, blue cheese dressing", 14, canterburyGolfCanterburyGolfSalads.ID)
	canterburyGolfCanterburyBurger := api.CreateItem("Canterbury Burger", "lettuce, tomato, onion, pickles, american cheese, canterbury burger sauce", 16, canterburyGolfCanterburyGolfSandwiches.ID)
	canterburyGolfRoastedPrimeRibFrenchDip := api.CreateItem("Roasted Prime Rib French Dip", "mushroom, onions, swiss cheese, au jus", 22, canterburyGolfCanterburyGolfSandwiches.ID)
	canterburyGolfGrilledChickenBreast := api.CreateItem("Grilled Chicken Breast", "pesto aioli, provolone, multigrain bun, roasted sweet bell peppers", 15, canterburyGolfCanterburyGolfSandwiches.ID)
	canterburyGolfSmokedTurkey := api.CreateItem("Smoked Turkey", "provolone cheese, coleslaw, tomato, chipotle aiole", 14, canterburyGolfCanterburyGolfSandwiches.ID)
	canterburyGolfWalnutChickenSalad := api.CreateItem("Walnut Chicken Salad", "lettuce, tomato", 13, canterburyGolfCanterburyGolfSandwiches.ID)
	canterburyGolfShrimpLobsterClub := api.CreateItem("Shrimp Lobster Club", "avocado, bacon, lettuce, tomato", 21, canterburyGolfCanterburyGolfSandwiches.ID)
	_ = api.CreateItem("Smoked Pork Chop", "ohio creamy polenta, roasted brussel sprouts with truffle honey", 28, canterburyGolfCanterburyGolfMainCourses.ID)
	_ = api.CreateItem("Pasta Bolognese", "house made gemelli, parmesan", 24, canterburyGolfCanterburyGolfMainCourses.ID)
	_ = api.CreateItem("Roasted Citrus Salmon", "citrus butter sauce, basil oil, apple slaw, honey glazed carrots", 26, canterburyGolfCanterburyGolfMainCourses.ID)
	_ = api.CreateItem("Grilled Shrimp Scampi", "linguine tossed in pesto cream sauce, cherry tomato, spinach, gremolata", 24, canterburyGolfCanterburyGolfMainCourses.ID)
	_ = api.CreateItem("Roasted Rosemary Chicken Breast", "capers, lemon poulet sauce, roasted fingerling potato hash", 22, canterburyGolfCanterburyGolfMainCourses.ID)
	_ = api.CreateItem("Braised Beef Short Rib", "mashed potatoes, roasted carrots and green beans, mushroom veal sauce", 38, canterburyGolfCanterburyGolfMainCourses.ID)
	_ = api.CreateItem("Creme Brulee", "", 9, canterburyGolfCanterburyGolfDesserts.ID)
	_ = api.CreateItem("Divot Pie", "", 16, canterburyGolfCanterburyGolfDesserts.ID)
	_ = api.CreateItem("Lemon Mascarpone Cheesecake", "", 10, canterburyGolfCanterburyGolfDesserts.ID)
	_ = api.CreateItem("Forbidden Chocolate Cake", "", 10, canterburyGolfCanterburyGolfDesserts.ID)
	_ = api.CreateItem("Bud", "Classic American Lager", 6, canterburyGolfCanterburyGolfBeer.ID)
	_ = api.CreateItem("Bud Light", "Light American Lager", 6, canterburyGolfCanterburyGolfBeer.ID)
	_ = api.CreateItem("Miller Lite", "Light Lager", 6, canterburyGolfCanterburyGolfBeer.ID)
	_ = api.CreateItem("Michelob Ultra", "Low-Calorie Lager", 6, canterburyGolfCanterburyGolfBeer.ID)
	_ = api.CreateItem("Coors Light", "Light Lager", 6, canterburyGolfCanterburyGolfBeer.ID)
	_ = api.CreateItem("Miller High Life", "American Lager", 6, canterburyGolfCanterburyGolfBeer.ID)
	_ = api.CreateItem("Corona", "Mexican Lager", 7, canterburyGolfCanterburyGolfBeer.ID)
	_ = api.CreateItem("Corona Light", "Light Mexican Lager", 7, canterburyGolfCanterburyGolfBeer.ID)
	_ = api.CreateItem("Dortmunder Gold", "Lager", 7, canterburyGolfCanterburyGolfBeer.ID)
	_ = api.CreateItem("Summer Shandy", "Lemon Shandy", 7, canterburyGolfCanterburyGolfBeer.ID)
	_ = api.CreateItem("Dogfish 60 Min IPA", "IPA", 7, canterburyGolfCanterburyGolfBeer.ID)
	_ = api.CreateItem("Coronado Orange Ave Wit", "Wheat Beer with Orange Flavor", 8, canterburyGolfCanterburyGolfBeer.ID)
	_ = api.CreateItem("Clown Shoes Spacecake IPA", "Double IPA", 11, canterburyGolfCanterburyGolfBeer.ID)
	_ = api.CreateItem("Masthead Hazy Headlines IPA", "Hazy IPA", 8, canterburyGolfCanterburyGolfBeer.ID)
	_ = api.CreateItem("Rhineguist Truth IPA", "IPA", 8, canterburyGolfCanterburyGolfBeer.ID)
	_ = api.CreateItem("Chough IPA", "IPA", 10, canterburyGolfCanterburyGolfBeer.ID)
	_ = api.CreateItem("Hi Wire Bob's Lager", "Lager", 7, canterburyGolfCanterburyGolfBeer.ID)
	canterburyGolfHighNoonVodka := api.CreateItem("High Noon (Vodka)", "Vodka Seltzer", 8, canterburyGolfCanterburyGolfBeer.ID)
	canterburyGolfHighNoonTequila := api.CreateItem("High Noon (Tequila)", "Tequila Seltzer", 8, canterburyGolfCanterburyGolfBeer.ID)
	_ = api.CreateItem("Run Wild IPA NA", "Non-Alcoholic IPA", 7, canterburyGolfCanterburyGolfBeer.ID)
	_ = api.CreateItem("Run Wild Upside Dawn NA", "Non-Alcoholic Ale", 7, canterburyGolfCanterburyGolfBeer.ID)
	_ = api.CreateItem("Becks NA", "Non-Alcoholic Beer", 7, canterburyGolfCanterburyGolfBeer.ID)
	_ = api.CreateItem("The Woodland", "Bourbon cocktail with cherry syrup, orange bitters, soda", 16, canterburyGolfCanterburyGolfCocktails.ID)
	_ = api.CreateItem("Spicy Jalapeno Margarita", "Tequila cocktail with coconut, lime, tajin rim", 14, canterburyGolfCanterburyGolfCocktails.ID)
	_ = api.CreateItem("Espresso Martini", "Vodka martini with coffee liqueur and cold brew", 12, canterburyGolfCanterburyGolfCocktails.ID)
	_ = api.CreateItem("Rose Aperol Spritz", "Aperol with passionfruit, lime, rose prosecco", 14, canterburyGolfCanterburyGolfCocktails.ID)
	_ = api.CreateItem("Greens Keeper", "Peach whiskey with ginger ale, lime juice, basil", 12, canterburyGolfCanterburyGolfCocktails.ID)
	_ = api.CreateItem("Rain Delay", "Gin cocktail with rosemary, lemon, honey syrup, soda", 12, canterburyGolfCanterburyGolfCocktails.ID)
	_ = api.CreateItem("Birdie Juice", "Vodka cocktail with elderflower, limecello, strawberries", 12, canterburyGolfCanterburyGolfCocktails.ID)
	// seed item options
	log.Println("Seeding *Options* data")
	_ = api.CreateOption("Wings Dipping Sauce", "Wings Dipping Sauce", 1, 1, canterburyGolfChickenWings.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Ranch", Description: "Ranch", Price: 0},
		locations.OptionItem{Name: "Blue Cheese", Description: "Blue Cheese", Price: 0}})
	_ = api.CreateOption("Add Protein", "Add Protein", 0, 0, canterburyGolfCaesar.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Grilled Chicken", Description: "Grilled Chicken", Price: 7},
		locations.OptionItem{Name: "Chicken Salad", Description: "Chicken Salad", Price: 7},
		locations.OptionItem{Name: "Grilled Salmon", Description: "Grilled Salmon", Price: 12},
		locations.OptionItem{Name: "Shrimp", Description: "Shrimp", Price: 9},
		locations.OptionItem{Name: "6 oz Strip", Description: "6 oz Strip", Price: 22}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, canterburyGolfCaesar.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Croutons", Description: "No Croutons", Price: 0}})
	_ = api.CreateOption("Salad Dressing", "Salad Dressing", 1, 1, canterburyGolfItalianChop.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Ranch", Description: "Ranch", Price: 0},
		locations.OptionItem{Name: "White Balsamic Vinaigrette", Description: "White Balsamic Vinaigrette", Price: 0},
		locations.OptionItem{Name: "Blue Cheese", Description: "Blue Cheese", Price: 0},
		locations.OptionItem{Name: "Caesar", Description: "Caesar", Price: 0},
		locations.OptionItem{Name: "Italian", Description: "Italian", Price: 0}})
	_ = api.CreateOption("Add Protein", "Add Protein", 0, 0, canterburyGolfItalianChop.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Grilled Chicken", Description: "Grilled Chicken", Price: 7},
		locations.OptionItem{Name: "Chicken Salad", Description: "Chicken Salad", Price: 7},
		locations.OptionItem{Name: "Grilled Salmon", Description: "Grilled Salmon", Price: 12},
		locations.OptionItem{Name: "Shrimp", Description: "Shrimp", Price: 9},
		locations.OptionItem{Name: "6 oz Strip", Description: "6 oz Strip", Price: 22}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, canterburyGolfItalianChop.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Cucumber", Description: "No Cucumber", Price: 0},
		locations.OptionItem{Name: "No Tomatoes", Description: "No Tomatoes", Price: 0},
		locations.OptionItem{Name: "No Salami", Description: "No Salami", Price: 0},
		locations.OptionItem{Name: "No Sweet Peppers", Description: "No Sweet Peppers", Price: 0},
		locations.OptionItem{Name: "No Olives", Description: "No Olives", Price: 0},
		locations.OptionItem{Name: "No Peperoncinis", Description: "No Peperoncinis", Price: 0},
		locations.OptionItem{Name: "No Chickpeas", Description: "No Chickpeas", Price: 0}})
	_ = api.CreateOption("Salad Dressing", "Salad Dressing", 1, 1, canterburyGolfStrawberryGoatCheese.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Ranch", Description: "Ranch", Price: 0},
		locations.OptionItem{Name: "White Balsamic Vinaigrette", Description: "White Balsamic Vinaigrette", Price: 0},
		locations.OptionItem{Name: "Blue Cheese", Description: "Blue Cheese", Price: 0},
		locations.OptionItem{Name: "Caesar", Description: "Caesar", Price: 0},
		locations.OptionItem{Name: "Italian", Description: "Italian", Price: 0}})
	_ = api.CreateOption("Add Protein", "Add Protein", 0, 0, canterburyGolfStrawberryGoatCheese.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Grilled Chicken", Description: "Grilled Chicken", Price: 7},
		locations.OptionItem{Name: "Chicken Salad", Description: "Chicken Salad", Price: 7},
		locations.OptionItem{Name: "Grilled Salmon", Description: "Grilled Salmon", Price: 12},
		locations.OptionItem{Name: "Shrimp", Description: "Shrimp", Price: 9},
		locations.OptionItem{Name: "6 oz Strip", Description: "6 oz Strip", Price: 22}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, canterburyGolfStrawberryGoatCheese.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Spiced Pecans", Description: "No Spiced Pecans", Price: 0},
		locations.OptionItem{Name: "No Raisins", Description: "No Raisins", Price: 0},
		locations.OptionItem{Name: "No Radish", Description: "No Radish", Price: 0}})
	_ = api.CreateOption("Salad Dressing", "Salad Dressing", 1, 1, canterburyGolfIceberg.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Ranch", Description: "Ranch", Price: 0},
		locations.OptionItem{Name: "White Balsamic Vinaigrette", Description: "White Balsamic Vinaigrette", Price: 0},
		locations.OptionItem{Name: "Blue Cheese", Description: "Blue Cheese", Price: 0},
		locations.OptionItem{Name: "Caesar", Description: "Caesar", Price: 0},
		locations.OptionItem{Name: "Italian", Description: "Italian", Price: 0}})
	_ = api.CreateOption("Add Protein", "Add Protein", 0, 0, canterburyGolfIceberg.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Grilled Chicken", Description: "Grilled Chicken", Price: 7},
		locations.OptionItem{Name: "Chicken Salad", Description: "Chicken Salad", Price: 7},
		locations.OptionItem{Name: "Grilled Salmon", Description: "Grilled Salmon", Price: 12},
		locations.OptionItem{Name: "Shrimp", Description: "Shrimp", Price: 9},
		locations.OptionItem{Name: "6 oz Strip", Description: "6 oz Strip", Price: 22}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, canterburyGolfIceberg.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Bacon", Description: "No Bacon", Price: 0},
		locations.OptionItem{Name: "No Tomatoes", Description: "No Tomatoes", Price: 0},
		locations.OptionItem{Name: "No Egg", Description: "No Egg", Price: 0},
		locations.OptionItem{Name: "No Blue Cheese", Description: "No Blue Cheese", Price: 0},
		locations.OptionItem{Name: "No Beets", Description: "No Beets", Price: 0},
		locations.OptionItem{Name: "No Dried Cherries", Description: "No Dried Cherries", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, canterburyGolfCanterburyBurger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Tator Tots", Description: "Tator Tots", Price: 0},
		locations.OptionItem{Name: "Fresh Fruit", Description: "Fresh Fruit", Price: 0}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, canterburyGolfCanterburyBurger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Lettuce", Description: "No Lettuce", Price: 0},
		locations.OptionItem{Name: "No Tomato", Description: "No Tomato", Price: 0},
		locations.OptionItem{Name: "No Onion", Description: "No Onion", Price: 0},
		locations.OptionItem{Name: "No Cheese", Description: "No Cheese", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, canterburyGolfRoastedPrimeRibFrenchDip.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Tator Tots", Description: "Tator Tots", Price: 0},
		locations.OptionItem{Name: "Fresh Fruit", Description: "Fresh Fruit", Price: 0}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, canterburyGolfRoastedPrimeRibFrenchDip.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Onions", Description: "No Onions", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, canterburyGolfGrilledChickenBreast.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Tator Tots", Description: "Tator Tots", Price: 0},
		locations.OptionItem{Name: "Fresh Fruit", Description: "Fresh Fruit", Price: 0}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, canterburyGolfGrilledChickenBreast.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Provolone", Description: "No Provolone", Price: 0},
		locations.OptionItem{Name: "No Roasted Bell Peppers", Description: "No Roasted Bell Peppers", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, canterburyGolfSmokedTurkey.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Tator Tots", Description: "Tator Tots", Price: 0},
		locations.OptionItem{Name: "Fresh Fruit", Description: "Fresh Fruit", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, canterburyGolfWalnutChickenSalad.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Tator Tots", Description: "Tator Tots", Price: 0},
		locations.OptionItem{Name: "Fresh Fruit", Description: "Fresh Fruit", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, canterburyGolfShrimpLobsterClub.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Tator Tots", Description: "Tator Tots", Price: 0},
		locations.OptionItem{Name: "Fresh Fruit", Description: "Fresh Fruit", Price: 0}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, canterburyGolfShrimpLobsterClub.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Avacado", Description: "No Avacado", Price: 0},
		locations.OptionItem{Name: "No Bacon", Description: "No Bacon", Price: 0},
		locations.OptionItem{Name: "No Lettuce", Description: "No Lettuce", Price: 0},
		locations.OptionItem{Name: "No Tomatoes", Description: "No Tomatoes", Price: 0}})
	_ = api.CreateOption("Flavor", "Flavor", 1, 1, canterburyGolfHighNoonVodka.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Black Cherry", Description: "Black Cherry", Price: 0},
		locations.OptionItem{Name: "Ruby Grapefruit", Description: "Ruby Grapefruit", Price: 0},
		locations.OptionItem{Name: "Natural Lime", Description: "Natural Lime", Price: 0},
		locations.OptionItem{Name: "Raspberry", Description: "Raspberry", Price: 0},
		locations.OptionItem{Name: "Mango", Description: "Mango", Price: 0}})
	_ = api.CreateOption("Flavor", "Flavor", 1, 1, canterburyGolfHighNoonTequila.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Strawberry", Description: "Strawberry", Price: 0},
		locations.OptionItem{Name: "Lime", Description: "Lime", Price: 0},
		locations.OptionItem{Name: "Grapefruit", Description: "Grapefruit", Price: 0},
		locations.OptionItem{Name: "Passionfruit", Description: "Passionfruit", Price: 0}})
	// seed orders
	log.Println("Seeding Completed")
}