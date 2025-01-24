// package main

// import (
// 	"log"

// 	"swoop/locations"
// 	"swoop/pkg/config"
// 	database "swoop/pkg/db"
// )

// func main() {
// 	err := config.Init()
// 	if err != nil {
// 		log.Panicln(err)
// 	}
// 	log.Println("Seeding Database")
// 	db := database.Connect(config.DB())
// 	api := locations.NewAPI(db)
// 	// seed locations
// 	log.Println("Seeding *Location* data")
// 	heritageGolfLinks := api.CreateLocation("Heritage Golf Links", "Tucker, GA")
// 	log.Println("Seeding *Menus* data")
// 	// seed menus
// 	heritageGolfLinksMenu := api.CreateMenu("Heritage Grill", "Heritage Grill", heritageGolfLinks.ID)
// 	log.Println(heritageGolfLinksMenu)
// 	// seed categories
// 	log.Println("Seeding *Categories* data")
// 	heritageGolfLinksStarters := api.CreateCategory("Starters", "Starters", heritageGolfLinksMenu.ID)
// 	heritageGolfLinksGreens := api.CreateCategory("Greens", "Greens", heritageGolfLinksMenu.ID)
// 	heritageGolfLinksChefsSignatures := api.CreateCategory("Chef's Signatures", "Chef's Signatures", heritageGolfLinksMenu.ID)
// 	heritageGolfLinksClassics := api.CreateCategory("Classics", "Classics", heritageGolfLinksMenu.ID)
// 	heritageGolfLinksBurgersAndHotDogs := api.CreateCategory("Burgers & Hot Dogs", "Burgers & Hot Dogs", heritageGolfLinksMenu.ID)
// 	heritageGolfLinksPizza := api.CreateCategory("Pizza", "Pizza", heritageGolfLinksMenu.ID)
// 	heritageGolfLinksSides := api.CreateCategory("Sides", "Sides", heritageGolfLinksMenu.ID)
// 	heritageGolfLinksBeverages := api.CreateCategory("Beverages", "Beverages", heritageGolfLinksMenu.ID)
// 	heritageGolfLinksBeerAndSeltzers := api.CreateCategory("Beer & Seltzers", "Beer & Seltzers", heritageGolfLinksMenu.ID)
// 	heritageGolfLinksCocktails := api.CreateCategory("Cocktails", "Cocktails", heritageGolfLinksMenu.ID)
// 	log.Println(heritageGolfLinksStarters)
// 	log.Println(heritageGolfLinksGreens)
// 	log.Println(heritageGolfLinksChefsSignatures)
// 	log.Println(heritageGolfLinksClassics)
// 	log.Println(heritageGolfLinksBurgersAndHotDogs)
// 	log.Println(heritageGolfLinksPizza)
// 	log.Println(heritageGolfLinksSides)
// 	log.Println(heritageGolfLinksBeverages)
// 	log.Println(heritageGolfLinksBeerAndSeltzers)
// 	log.Println(heritageGolfLinksCocktails)
// 	// seed item
// 	log.Println("Seeding *Items* data")
// 	heritageGolfLinksQuesadilla := api.CreateItem("Quesadilla", "Crispy Tortilla, Sautéed Peppers, Onions with Side of Salsa & Sour Cream", 8, heritageGolfLinksStarters.ID)
// 	_ = api.CreateItem("Big Bang Shrimp", "8 crispy fried shrimp tossed in sweet chili sauce", 10, heritageGolfLinksStarters.ID)
// 	heritageGolfLinksBonelessWings := api.CreateItem("Boneless Wings", "10 boneless wings tossed in your desired flavor served with your choose of dipping sauce.", 11, heritageGolfLinksStarters.ID)
// 	heritageGolfLinksWings := api.CreateItem("Wings", "10 boneless wings tossed in your desired flavor served with your choose of dipping sauce.", 7, heritageGolfLinksStarters.ID)
// 	_ = api.CreateItem("Spring Rolls", "Served with a side of Thai sauce", 5, heritageGolfLinksStarters.ID)
// 	_ = api.CreateItem("Chips & Salsa ", "Crispy & Warm Tortilla Chips with Home Made Salsa", 5, heritageGolfLinksStarters.ID)
// 	_ = api.CreateItem("Chips & Queso", "Crispy & Warm Tortilla Chips with Creamy Queso Dip", 7, heritageGolfLinksStarters.ID)
// 	heritageGolfLinksMileHighNachos := api.CreateItem("Mile High Nachos", "Crispy & Warm Tortilla Chips Layered with Queso & Topped with Tomatoes, Lettuce, Jalapenos, Salsa & Sour Cream", 8, heritageGolfLinksStarters.ID)
// 	_ = api.CreateItem("Mozzarella Sticks", "Six (6) Crispy Fried Mozzarella Sticks Served with Marinara Sauce", 6, heritageGolfLinksStarters.ID)
// 	heritageGolfLinksHouseSalad := api.CreateItem("House Salad", "Salad Mix, Tomatoes, Cucumbers, Onions & Shredded Cheese", 7, heritageGolfLinksGreens.ID)
// 	heritageGolfLinksCobbSalad := api.CreateItem("Cobb Salad", "Salad Mix, Blue Cheese Crumbles, Bacon, Tomatoes, Onion, Eggs with Grilled or Fried Chicken", 11, heritageGolfLinksGreens.ID)
// 	heritageGolfLinksChefSalad := api.CreateItem("Chef Salad", "Salad Mix, Onions, Cucumbers, Shredded Cheese, Tomatoes, Ham & Turkey", 11, heritageGolfLinksGreens.ID)
// 	heritageGolfLinksChickenStirFry := api.CreateItem("Chicken Stir Fry", "Chicken with white rice, broccoli, carrots, sugar snap peas, red and green pepper mix with chefs signature spices and toasted sesame seeds.", 11, heritageGolfLinksChefsSignatures.ID)
// 	heritageGolfLinksFishandChips := api.CreateItem("Fish and Chips", "Breaded tilapia in signature spices served with choice of 1 side.", 10, heritageGolfLinksChefsSignatures.ID)
// 	heritageGolfLinksBeefandBroccoli := api.CreateItem("Beef and Broccoli", "Beef and broccoli served on bed of white rice with signature Asian sauce.", 10, heritageGolfLinksChefsSignatures.ID)
// 	heritageGolfLinksBuffaloChickenWrap := api.CreateItem("Buffalo Chicken Wrap", "Flour tortilla filled with grilled or fried chicken, romaine lettuce, tomato, onions, corn kernels, Gorgonzola dressing with diced bacon.", 10, heritageGolfLinksClassics.ID)
// 	heritageGolfLinksChickenSalad := api.CreateItem("Chicken Salad", "House made chicken salad, lettuce and tomatoes, on your choice of croissant or sliced bread.", 9, heritageGolfLinksClassics.ID)
// 	heritageGolfLinksTunaSalad := api.CreateItem("Tuna Salad", "House made tuna salad, lettuce and tomatoes, on your choice of croissant or sliced bread.", 9, heritageGolfLinksClassics.ID)
// 	heritageGolfLinksHeritageClub := api.CreateItem("Heritage Club", "Flour tortilla filled with grilled or fried chicken, romaine lettuce, tomato, onions, corn kernels, Gorgonzola dressing with diced bacon.", 10, heritageGolfLinksClassics.ID)
// 	heritageGolfLinksTendersandFries := api.CreateItem("Tenders and Fries", "Hand breaded tenders and fries and served with honey mustard.", 8, heritageGolfLinksClassics.ID)
// 	heritageGolfLinksKoreanBBQChickenWrap := api.CreateItem("Korean BBQ Chicken Wrap", "Flour tortilla filled with grilled or fried chicken, romaine lettuce, tomato, onions, corn kernels, Gorgonzola dressing with diced bacon.", 10, heritageGolfLinksClassics.ID)
// 	heritageGolfLinksChickenSandwich := api.CreateItem("Chicken Sandwich", "Swiss cheese, lettuce, tomato, Swiss cheese with honey mustard dressing.", 10, heritageGolfLinksClassics.ID)
// 	heritageGolfLinksClassicBLT := api.CreateItem("Classic BLT", "Four (4) Slices of Bacon, Lettuce, Tomato, On Sliced Bread", 8, heritageGolfLinksClassics.ID)
// 	heritageGolfLinksBYOBBurger := api.CreateItem("BYOB Burger", "Lettuce, Tomatoes, Pickles, Onions with American Cheese", 10, heritageGolfLinksBurgersAndHotDogs.ID)
// 	heritageGolfLinksVeggieBlackBeanBurger := api.CreateItem("Veggie Black Bean Burger", "Veggie black bean burger with lettuce, tomatoes and onion.", 10, heritageGolfLinksBurgersAndHotDogs.ID)
// 	heritageGolfLinksBlackandBlueBurger := api.CreateItem("Black and Blue Burger", "Blackened seasoned burger, Gorgonzola dressing and bacon", 10, heritageGolfLinksBurgersAndHotDogs.ID)
// 	heritageGolfLinksHeritageDog := api.CreateItem("Heritage Dog", "Quarter Pound All Beef Hot Dog", 4, heritageGolfLinksBurgersAndHotDogs.ID)
// 	heritageGolfLinksHeritageCheesyDog := api.CreateItem("Heritage Cheesy Dog", "Quarter Pound All Beef Hot Dog with American Cheese Melted", 5, heritageGolfLinksBurgersAndHotDogs.ID)
// 	heritageGolfLinksHeritageSouthwestDog := api.CreateItem("Heritage Southwest Dog", "Quarter Pound All Beef Hot Dog Covered in Queso Cheese, Jalapenos & Diced Tomatoes", 6, heritageGolfLinksBurgersAndHotDogs.ID)
// 	heritageGolfLinks12inchClassicMozzarellaAndMarinaraPizza := api.CreateItem("12-inch Classic Mozzarella & Marinara Pizza", "2 Toppings", 9.99, heritageGolfLinksPizza.ID)
// 	heritageGolfLinksBBQChickenPizza := api.CreateItem("BBQ Chicken Pizza", "Sweet Chili BBQ Sauce, Grilled Chicken, Spinach, Red Onion, Bell Pepper, Mushrooms", 11.99, heritageGolfLinksPizza.ID)
// 	heritageGolfLinksHawaiianPizza := api.CreateItem("Hawaiian Pizza", "Ham, Pineapple, Red Onion", 11.99, heritageGolfLinksPizza.ID)
// 	heritageGolfLinksTheLoadedSupreme := api.CreateItem("The Loaded Supreme", "Pepperoni, Sausage, Olives, Red Onion, Mushrooms, Bell Pepper", 11.99, heritageGolfLinksPizza.ID)
// 	_ = api.CreateItem("Fries", "", 3, heritageGolfLinksSides.ID)
// 	_ = api.CreateItem("Sweet Potato Fries", "", 3, heritageGolfLinksSides.ID)
// 	_ = api.CreateItem("Potato Salad", "", 3, heritageGolfLinksSides.ID)
// 	_ = api.CreateItem("Onion Rings", "", 3, heritageGolfLinksSides.ID)
// 	_ = api.CreateItem("Coleslaw", "", 3, heritageGolfLinksSides.ID)
// 	_ = api.CreateItem("Tots", "", 3, heritageGolfLinksSides.ID)
// 	_ = api.CreateItem("Bottled Water", "", 2, heritageGolfLinksBeverages.ID)
// 	heritageGolfLinksCannedSoda := api.CreateItem("Canned Soda", "Coke, Diet Coke, Sprite, Diet Sprite, Dr. Pepper, Diet Dr. Pepper", 1.5, heritageGolfLinksBeverages.ID)
// 	heritageGolfLinksFountainDrink := api.CreateItem("Fountain Drink", "Coke, Diet Coke, Sprite, Diet Sprite, Dr. Pepper, Diet Dr. Pepper", 2.5, heritageGolfLinksBeverages.ID)
// 	heritageGolfLinksBottledGatorade := api.CreateItem("Bottled Gatorade", "Lemon-Lime, Fruit Punch, Orange, Arctic Freeze, Cool Blue", 3.5, heritageGolfLinksBeverages.ID)
// 	heritageGolfLinksBlueMoon := api.CreateItem("Blue Moon", "", 5, heritageGolfLinksBeerAndSeltzers.ID)
// 	heritageGolfLinksBudLight := api.CreateItem("Bud Light", "", 5, heritageGolfLinksBeerAndSeltzers.ID)
// 	heritageGolfLinksBudLightLime := api.CreateItem("Bud Light Lime", "", 5, heritageGolfLinksBeerAndSeltzers.ID)
// 	heritageGolfLinksBudweiser := api.CreateItem("Budweiser", "", 5, heritageGolfLinksBeerAndSeltzers.ID)
// 	heritageGolfLinksCoorsLight := api.CreateItem("Coors Light", "", 5, heritageGolfLinksBeerAndSeltzers.ID)
// 	heritageGolfLinksCorona := api.CreateItem("Corona", "", 5, heritageGolfLinksBeerAndSeltzers.ID)
// 	heritageGolfLinksCoronaLight := api.CreateItem("Corona Light", "", 5, heritageGolfLinksBeerAndSeltzers.ID)
// 	heritageGolfLinksDosEquisLager := api.CreateItem("Dos Equis Lager", "", 5, heritageGolfLinksBeerAndSeltzers.ID)
// 	heritageGolfLinksMichelobUltra := api.CreateItem("Michelob Ultra", "", 5, heritageGolfLinksBeerAndSeltzers.ID)
// 	heritageGolfLinksMillerLite := api.CreateItem("Miller Lite", "", 5, heritageGolfLinksBeerAndSeltzers.ID)
// 	heritageGolfLinksStellaArtois := api.CreateItem("Stella Artois", "", 5, heritageGolfLinksBeerAndSeltzers.ID)
// 	heritageGolfLinksHeineken := api.CreateItem("Heineken", "", 5, heritageGolfLinksBeerAndSeltzers.ID)
// 	heritageGolfLinksTruly := api.CreateItem("Truly", "", 5, heritageGolfLinksBeerAndSeltzers.ID)
// 	heritageGolfLinksWhiteClaw := api.CreateItem("White Claw", "", 5, heritageGolfLinksBeerAndSeltzers.ID)
// 	heritageGolfLinksScrewDriver := api.CreateItem("Screw Driver", "Vodka and Fresh Squeezed Orange Juice", 10, heritageGolfLinksCocktails.ID)
// 	heritageGolfLinksBloodyMary := api.CreateItem("Bloody Mary", "Vodka, house bloody mary mix, fresh celery, fresh-grated horseradish, rosemary skewered olive, cherry tomato, spicy pepper", 13, heritageGolfLinksCocktails.ID)
// 	_ = api.CreateItem("Mimosa", "Sparkling Wine, Orange Juice", 8, heritageGolfLinksCocktails.ID)
// 	heritageGolfLinksJohnDaly := api.CreateItem("John Daly", "Fresh Brewed Iced Tea, Lemonade, and Vodka.", 12, heritageGolfLinksCocktails.ID)
// 	heritageGolfLinksPaloma := api.CreateItem("Paloma", "Tequila, Grapefruit Juice, Fresh Lime Juice, Agave Syrup, Club Soda", 11, heritageGolfLinksCocktails.ID)
// 	heritageGolfLinksMoscowMule := api.CreateItem("Moscow Mule", "Vodka, lime juice, mint, ginger beer", 11, heritageGolfLinksCocktails.ID)
// 	heritageGolfLinksRumAndCoke := api.CreateItem("Rum & Coke", "", 11, heritageGolfLinksCocktails.ID)
// 	heritageGolfLinksWhiskeyAndCoke := api.CreateItem("Whiskey & Coke", "", 11, heritageGolfLinksCocktails.ID)
// 	heritageGolfLinksAzelea := api.CreateItem("Azelea", "Lemon juice, pineapple juice, vodka and grenadine.", 11, heritageGolfLinksCocktails.ID)
// 	heritageGolfLinksTransfusion := api.CreateItem("Transfusion", "Vodka, ginger ale and grape juice", 11, heritageGolfLinksCocktails.ID)
// 	heritageGolfLinksRedBullAndVodka := api.CreateItem("Red Bull & Vodka", "", 11, heritageGolfLinksCocktails.ID)
// 	heritageGolfLinksCutwaterCannedCocktail := api.CreateItem("Cutwater Canned Cocktail", "Tequila Margarita, Vodka Mule, Lime Vodka Soda, Cucumber Vodka Soda, Tequila Paloma, Rum & Cola, Rum Mint Mojito, Spicy Bloody Mary, Gin & Tonic, White Russian", 6, heritageGolfLinksCocktails.ID)
// 	// seed item options
// 	log.Println("Seeding *Options* data")
// 	_ = api.CreateOption("Options", "Options", 1, 1, heritageGolfLinksQuesadilla.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Chicken", Description: "Chicken", Price: 0},
// 		locations.OptionItem{Name: "Steak", Description: "Steak", Price: 0}})
// 	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, heritageGolfLinksQuesadilla.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "No Onions", Description: "No Onions", Price: 0}})
// 	_ = api.CreateOption("Wings Sauce", "Wings Sauce", 1, 1, heritageGolfLinksBonelessWings.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Buffalo Sauce", Description: "Buffalo Sauce", Price: 0},
// 		locations.OptionItem{Name: "BBQ Sauce", Description: "BBQ Sauce", Price: 0},
// 		locations.OptionItem{Name: "Lemon Pepper Sauce", Description: "Lemon Pepper Sauce", Price: 0},
// 		locations.OptionItem{Name: "Teriyaki Sauce", Description: "Teriyaki Sauce", Price: 0}})
// 	_ = api.CreateOption("Wings Dipping Sauce", "Wings Dipping Sauce", 1, 1, heritageGolfLinksBonelessWings.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Ranch", Description: "Ranch", Price: 0},
// 		locations.OptionItem{Name: "Blue Cheese", Description: "Blue Cheese", Price: 0}})
// 	_ = api.CreateOption("Extras", "Extras", 1, 1, heritageGolfLinksWings.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "6 Count", Description: "6 Count", Price: 0},
// 		locations.OptionItem{Name: "12 Count", Description: "12 Count", Price: 5}})
// 	_ = api.CreateOption("Wings Sauce", "Wings Sauce", 1, 1, heritageGolfLinksWings.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Buffalo Sauce", Description: "Buffalo Sauce", Price: 0},
// 		locations.OptionItem{Name: "BBQ Sauce", Description: "BBQ Sauce", Price: 0},
// 		locations.OptionItem{Name: "Lemon Pepper Sauce", Description: "Lemon Pepper Sauce", Price: 0},
// 		locations.OptionItem{Name: "Teriyaki Sauce", Description: "Teriyaki Sauce", Price: 0}})
// 	_ = api.CreateOption("Wings Dipping Sauce", "Wings Dipping Sauce", 1, 1, heritageGolfLinksWings.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Ranch", Description: "Ranch", Price: 0},
// 		locations.OptionItem{Name: "Blue Cheese", Description: "Blue Cheese", Price: 0}})
// 	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, heritageGolfLinksMileHighNachos.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "No Tomatoes", Description: "No Tomatoes", Price: 0},
// 		locations.OptionItem{Name: "No Lettuce", Description: "No Lettuce", Price: 0},
// 		locations.OptionItem{Name: "No Jalapenos", Description: "No Jalapenos", Price: 0},
// 		locations.OptionItem{Name: "No Salsa", Description: "No Salsa", Price: 0},
// 		locations.OptionItem{Name: "No Sour Cream", Description: "No Sour Cream", Price: 0}})
// 	_ = api.CreateOption("Salad Dressing", "Salad Dressing", 1, 1, heritageGolfLinksHouseSalad.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Balsalmic", Description: "Balsalmic", Price: 0},
// 		locations.OptionItem{Name: "Caesar", Description: "Caesar", Price: 0},
// 		locations.OptionItem{Name: "Ranch", Description: "Ranch", Price: 0},
// 		locations.OptionItem{Name: "Bleu Cheese", Description: "Bleu Cheese", Price: 0},
// 		locations.OptionItem{Name: "Honey Mustard", Description: "Honey Mustard", Price: 0},
// 		locations.OptionItem{Name: "Italian", Description: "Italian", Price: 0}})
// 	_ = api.CreateOption("Extras", "Extras", 0, 0, heritageGolfLinksHouseSalad.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Fried Chicken", Description: "Fried Chicken", Price: 3},
// 		locations.OptionItem{Name: "Grilled Chicken", Description: "Grilled Chicken", Price: 3}})
// 	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, heritageGolfLinksHouseSalad.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "No Onions", Description: "No Onions", Price: 0},
// 		locations.OptionItem{Name: "No Tomatoes", Description: "No Tomatoes", Price: 0},
// 		locations.OptionItem{Name: "No Cucumbers", Description: "No Cucumbers", Price: 0},
// 		locations.OptionItem{Name: "No Cheese", Description: "No Cheese", Price: 0}})
// 	_ = api.CreateOption("Options", "Options", 1, 1, heritageGolfLinksCobbSalad.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Fried Chicken", Description: "Fried Chicken", Price: 0},
// 		locations.OptionItem{Name: "Grilled Chicken", Description: "Grilled Chicken", Price: 0}})
// 	_ = api.CreateOption("Salad Dressing", "Salad Dressing", 1, 1, heritageGolfLinksCobbSalad.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Balsalmic", Description: "Balsalmic", Price: 0},
// 		locations.OptionItem{Name: "Caesar", Description: "Caesar", Price: 0},
// 		locations.OptionItem{Name: "Ranch", Description: "Ranch", Price: 0},
// 		locations.OptionItem{Name: "Bleu Cheese", Description: "Bleu Cheese", Price: 0},
// 		locations.OptionItem{Name: "Honey Mustard", Description: "Honey Mustard", Price: 0},
// 		locations.OptionItem{Name: "Italian", Description: "Italian", Price: 0}})
// 	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, heritageGolfLinksCobbSalad.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "No Blue Cheese", Description: "No Blue Cheese", Price: 0},
// 		locations.OptionItem{Name: "No Bacon", Description: "No Bacon", Price: 0},
// 		locations.OptionItem{Name: "No Tomatoes", Description: "No Tomatoes", Price: 0},
// 		locations.OptionItem{Name: "No Onions", Description: "No Onions", Price: 0},
// 		locations.OptionItem{Name: "No Eggs", Description: "No Eggs", Price: 0}})
// 	_ = api.CreateOption("Salad Dressing", "Salad Dressing", 1, 1, heritageGolfLinksChefSalad.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Balsalmic", Description: "Balsalmic", Price: 0},
// 		locations.OptionItem{Name: "Caesar", Description: "Caesar", Price: 0},
// 		locations.OptionItem{Name: "Ranch", Description: "Ranch", Price: 0},
// 		locations.OptionItem{Name: "Bleu Cheese", Description: "Bleu Cheese", Price: 0},
// 		locations.OptionItem{Name: "Honey Mustard", Description: "Honey Mustard", Price: 0},
// 		locations.OptionItem{Name: "Italian", Description: "Italian", Price: 0}})
// 	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, heritageGolfLinksChefSalad.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "No Onions", Description: "No Onions", Price: 0},
// 		locations.OptionItem{Name: "No Cucumber", Description: "No Cucumber", Price: 0},
// 		locations.OptionItem{Name: "No Cheese", Description: "No Cheese", Price: 0},
// 		locations.OptionItem{Name: "No Tomatoes", Description: "No Tomatoes", Price: 0},
// 		locations.OptionItem{Name: "No Ham", Description: "No Ham", Price: 0},
// 		locations.OptionItem{Name: "No Turkey", Description: "No Turkey", Price: 0}})
// 	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, heritageGolfLinksChickenStirFry.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "No Bacon", Description: "No Bacon", Price: 0},
// 		locations.OptionItem{Name: "No Lettuce", Description: "No Lettuce", Price: 0},
// 		locations.OptionItem{Name: "No Tomatoes", Description: "No Tomatoes", Price: 0},
// 		locations.OptionItem{Name: "No Cheese", Description: "No Cheese", Price: 0}})
// 	_ = api.CreateOption("Sides", "Sides", 1, 1, heritageGolfLinksFishandChips.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Chips", Description: "Chips", Price: 0},
// 		locations.OptionItem{Name: "Fries", Description: "Fries", Price: 0},
// 		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0},
// 		locations.OptionItem{Name: "Sweet Potato Fries ", Description: "Sweet Potato Fries ", Price: 0},
// 		locations.OptionItem{Name: "Cole Slaw", Description: "Cole Slaw", Price: 0},
// 		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 0},
// 		locations.OptionItem{Name: "Tater Tots ", Description: "Tater Tots ", Price: 0}})
// 	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, heritageGolfLinksFishandChips.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "No Bacon", Description: "No Bacon", Price: 0},
// 		locations.OptionItem{Name: "No Lettuce", Description: "No Lettuce", Price: 0},
// 		locations.OptionItem{Name: "No Tomatoes", Description: "No Tomatoes", Price: 0},
// 		locations.OptionItem{Name: "No Cheese", Description: "No Cheese", Price: 0}})
// 	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, heritageGolfLinksBeefandBroccoli.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "No Bacon", Description: "No Bacon", Price: 0},
// 		locations.OptionItem{Name: "No Lettuce", Description: "No Lettuce", Price: 0},
// 		locations.OptionItem{Name: "No Tomatoes", Description: "No Tomatoes", Price: 0},
// 		locations.OptionItem{Name: "No Cheese", Description: "No Cheese", Price: 0}})
// 	_ = api.CreateOption("Options", "Options", 1, 1, heritageGolfLinksBuffaloChickenWrap.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Fried Chicken", Description: "Fried Chicken", Price: 0},
// 		locations.OptionItem{Name: "Grilled Chicken", Description: "Grilled Chicken", Price: 0}})
// 	_ = api.CreateOption("Sauce", "Sauce", 1, 1, heritageGolfLinksBuffaloChickenWrap.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Mild", Description: "Mild", Price: 0},
// 		locations.OptionItem{Name: "Medium", Description: "Medium", Price: 0},
// 		locations.OptionItem{Name: "Hot", Description: "Hot", Price: 0}})
// 	_ = api.CreateOption("Sides", "Sides", 1, 1, heritageGolfLinksBuffaloChickenWrap.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Fries", Description: "Fries", Price: 0},
// 		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0},
// 		locations.OptionItem{Name: "Sweet Potato Fries ", Description: "Sweet Potato Fries ", Price: 0},
// 		locations.OptionItem{Name: "Cole Slaw", Description: "Cole Slaw", Price: 0},
// 		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 0},
// 		locations.OptionItem{Name: "Tater Tots ", Description: "Tater Tots ", Price: 0}})
// 	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, heritageGolfLinksBuffaloChickenWrap.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "No Lettuce", Description: "No Lettuce", Price: 0},
// 		locations.OptionItem{Name: "No Tomatoes", Description: "No Tomatoes", Price: 0},
// 		locations.OptionItem{Name: "No Cheese", Description: "No Cheese", Price: 0}})
// 	_ = api.CreateOption("Choice of Bread", "Choice of Bread", 1, 1, heritageGolfLinksChickenSalad.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Croissant", Description: "Croissant", Price: 0},
// 		locations.OptionItem{Name: "Sliced Bread", Description: "Sliced Bread", Price: 0}})
// 	_ = api.CreateOption("Sides", "Sides", 1, 1, heritageGolfLinksChickenSalad.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Fries", Description: "Fries", Price: 0},
// 		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0},
// 		locations.OptionItem{Name: "Sweet Potato Fries ", Description: "Sweet Potato Fries ", Price: 0},
// 		locations.OptionItem{Name: "Cole Slaw", Description: "Cole Slaw", Price: 0},
// 		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 0},
// 		locations.OptionItem{Name: "Tater Tots ", Description: "Tater Tots ", Price: 0}})
// 	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, heritageGolfLinksChickenSalad.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "No Lettuce", Description: "No Lettuce", Price: 0},
// 		locations.OptionItem{Name: "No Tomatoes", Description: "No Tomatoes", Price: 0}})
// 	_ = api.CreateOption("Choice of Bread", "Choice of Bread", 1, 1, heritageGolfLinksTunaSalad.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Croissant", Description: "Croissant", Price: 0},
// 		locations.OptionItem{Name: "Sliced Bread", Description: "Sliced Bread", Price: 0}})
// 	_ = api.CreateOption("Sides", "Sides", 1, 1, heritageGolfLinksTunaSalad.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Fries", Description: "Fries", Price: 0},
// 		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0},
// 		locations.OptionItem{Name: "Sweet Potato Fries ", Description: "Sweet Potato Fries ", Price: 0},
// 		locations.OptionItem{Name: "Cole Slaw", Description: "Cole Slaw", Price: 0},
// 		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 0},
// 		locations.OptionItem{Name: "Tater Tots ", Description: "Tater Tots ", Price: 0}})
// 	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, heritageGolfLinksTunaSalad.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "No Lettuce", Description: "No Lettuce", Price: 0},
// 		locations.OptionItem{Name: "No Tomatoes", Description: "No Tomatoes", Price: 0}})
// 	_ = api.CreateOption("Sides", "Sides", 1, 1, heritageGolfLinksHeritageClub.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Fries", Description: "Fries", Price: 0},
// 		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0},
// 		locations.OptionItem{Name: "Sweet Potato Fries ", Description: "Sweet Potato Fries ", Price: 0},
// 		locations.OptionItem{Name: "Cole Slaw", Description: "Cole Slaw", Price: 0},
// 		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 0},
// 		locations.OptionItem{Name: "Tater Tots ", Description: "Tater Tots ", Price: 0}})
// 	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, heritageGolfLinksHeritageClub.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "No Bacon", Description: "No Bacon", Price: 0},
// 		locations.OptionItem{Name: "No Lettuce", Description: "No Lettuce", Price: 0},
// 		locations.OptionItem{Name: "No Tomatoes", Description: "No Tomatoes", Price: 0},
// 		locations.OptionItem{Name: "No Cheese", Description: "No Cheese", Price: 0}})
// 	_ = api.CreateOption("Wings Dipping Sauce", "Wings Dipping Sauce", 1, 1, heritageGolfLinksTendersandFries.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Honey Mustard", Description: "Honey Mustard", Price: 0},
// 		locations.OptionItem{Name: "Ranch", Description: "Ranch", Price: 0}})
// 	_ = api.CreateOption("Sauce", "Sauce", 1, 1, heritageGolfLinksTendersandFries.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "No Sauce", Description: "No Sauce", Price: 0},
// 		locations.OptionItem{Name: "Mild", Description: "Mild", Price: 0},
// 		locations.OptionItem{Name: "Medium", Description: "Medium", Price: 0},
// 		locations.OptionItem{Name: "Hot", Description: "Hot", Price: 0}})
// 	_ = api.CreateOption("Sides", "Sides", 1, 1, heritageGolfLinksTendersandFries.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Fries", Description: "Fries", Price: 0},
// 		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0},
// 		locations.OptionItem{Name: "Sweet Potato Fries ", Description: "Sweet Potato Fries ", Price: 0},
// 		locations.OptionItem{Name: "Cole Slaw", Description: "Cole Slaw", Price: 0},
// 		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 0},
// 		locations.OptionItem{Name: "Tater Tots ", Description: "Tater Tots ", Price: 0}})
// 	_ = api.CreateOption("Options", "Options", 1, 1, heritageGolfLinksKoreanBBQChickenWrap.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Fried Chicken", Description: "Fried Chicken", Price: 0},
// 		locations.OptionItem{Name: "Grilled Chicken", Description: "Grilled Chicken", Price: 0}})
// 	_ = api.CreateOption("Sauce", "Sauce", 1, 1, heritageGolfLinksKoreanBBQChickenWrap.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "No Sauce", Description: "No Sauce", Price: 0},
// 		locations.OptionItem{Name: "Mild", Description: "Mild", Price: 0},
// 		locations.OptionItem{Name: "Medium", Description: "Medium", Price: 0},
// 		locations.OptionItem{Name: "Hot", Description: "Hot", Price: 0}})
// 	_ = api.CreateOption("Sides", "Sides", 1, 1, heritageGolfLinksKoreanBBQChickenWrap.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Fries", Description: "Fries", Price: 0},
// 		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0},
// 		locations.OptionItem{Name: "Sweet Potato Fries ", Description: "Sweet Potato Fries ", Price: 0},
// 		locations.OptionItem{Name: "Cole Slaw", Description: "Cole Slaw", Price: 0},
// 		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 0},
// 		locations.OptionItem{Name: "Tater Tots ", Description: "Tater Tots ", Price: 0}})
// 	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, heritageGolfLinksKoreanBBQChickenWrap.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "No Tomatoes", Description: "No Tomatoes", Price: 0},
// 		locations.OptionItem{Name: "No Onions", Description: "No Onions", Price: 0},
// 		locations.OptionItem{Name: "No Corn", Description: "No Corn", Price: 0},
// 		locations.OptionItem{Name: "No Bacon", Description: "No Bacon", Price: 0}})
// 	_ = api.CreateOption("Options", "Options", 1, 1, heritageGolfLinksChickenSandwich.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Fried Chicken", Description: "Fried Chicken", Price: 0},
// 		locations.OptionItem{Name: "Grilled Chicken", Description: "Grilled Chicken", Price: 0}})
// 	_ = api.CreateOption("Sides", "Sides", 1, 1, heritageGolfLinksChickenSandwich.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Fries", Description: "Fries", Price: 0},
// 		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0},
// 		locations.OptionItem{Name: "Sweet Potato Fries ", Description: "Sweet Potato Fries ", Price: 0},
// 		locations.OptionItem{Name: "Cole Slaw", Description: "Cole Slaw", Price: 0},
// 		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 0},
// 		locations.OptionItem{Name: "Tater Tots ", Description: "Tater Tots ", Price: 0}})
// 	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, heritageGolfLinksChickenSandwich.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "No Lettuce", Description: "No Lettuce", Price: 0},
// 		locations.OptionItem{Name: "No Tomatoes", Description: "No Tomatoes", Price: 0}})
// 	_ = api.CreateOption("Sides", "Sides", 1, 1, heritageGolfLinksClassicBLT.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Fries", Description: "Fries", Price: 0},
// 		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0},
// 		locations.OptionItem{Name: "Sweet Potato Fries ", Description: "Sweet Potato Fries ", Price: 0},
// 		locations.OptionItem{Name: "Cole Slaw", Description: "Cole Slaw", Price: 0},
// 		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 0},
// 		locations.OptionItem{Name: "Tater Tots ", Description: "Tater Tots ", Price: 0}})
// 	_ = api.CreateOption("Meat Temperature", "Meat Temperature", 1, 1, heritageGolfLinksBYOBBurger.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Well Done", Description: "Well Done", Price: 0},
// 		locations.OptionItem{Name: "Medium-Well", Description: "Medium-Well", Price: 0},
// 		locations.OptionItem{Name: "Medium", Description: "Medium", Price: 0},
// 		locations.OptionItem{Name: "Medium-Rare", Description: "Medium-Rare", Price: 0},
// 		locations.OptionItem{Name: "Rare", Description: "Rare", Price: 0}})
// 	_ = api.CreateOption("Sides", "Sides", 1, 1, heritageGolfLinksBYOBBurger.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Fries", Description: "Fries", Price: 0},
// 		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0},
// 		locations.OptionItem{Name: "Sweet Potato Fries ", Description: "Sweet Potato Fries ", Price: 0},
// 		locations.OptionItem{Name: "Cole Slaw", Description: "Cole Slaw", Price: 0},
// 		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 0},
// 		locations.OptionItem{Name: "Tater Tots ", Description: "Tater Tots ", Price: 0}})
// 	_ = api.CreateOption("Options", "Options", 0, 0, heritageGolfLinksBYOBBurger.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Caramelized Onions", Description: "Caramelized Onions", Price: 0.5},
// 		locations.OptionItem{Name: "Jalepenos", Description: "Jalepenos", Price: 0.5},
// 		locations.OptionItem{Name: "Onion Ring", Description: "Onion Ring", Price: 0.5},
// 		locations.OptionItem{Name: "Fried Egg", Description: "Fried Egg", Price: 1},
// 		locations.OptionItem{Name: "Bacon", Description: "Bacon", Price: 1},
// 		locations.OptionItem{Name: "Mayonnaise", Description: "Mayonnaise", Price: 0},
// 		locations.OptionItem{Name: "Mustard", Description: "Mustard", Price: 0},
// 		locations.OptionItem{Name: "Ketchup", Description: "Ketchup", Price: 0}})
// 	_ = api.CreateOption("Cheese", "Cheese", 0, 0, heritageGolfLinksBYOBBurger.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Cheddar", Description: "Cheddar", Price: 0.5},
// 		locations.OptionItem{Name: "Swiss", Description: "Swiss", Price: 0.5},
// 		locations.OptionItem{Name: "American", Description: "American", Price: 0.5},
// 		locations.OptionItem{Name: "Provolone", Description: "Provolone", Price: 0.5},
// 		locations.OptionItem{Name: "Pepper Jack", Description: "Pepper Jack", Price: 0.5}})
// 	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, heritageGolfLinksBYOBBurger.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "No Onions", Description: "No Onions", Price: 0},
// 		locations.OptionItem{Name: "No Tomatoes", Description: "No Tomatoes", Price: 0},
// 		locations.OptionItem{Name: "No Cheese", Description: "No Cheese", Price: 0}})
// 	_ = api.CreateOption("Sides", "Sides", 1, 1, heritageGolfLinksVeggieBlackBeanBurger.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Fries", Description: "Fries", Price: 0},
// 		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0},
// 		locations.OptionItem{Name: "Sweet Potato Fries ", Description: "Sweet Potato Fries ", Price: 0},
// 		locations.OptionItem{Name: "Cole Slaw", Description: "Cole Slaw", Price: 0},
// 		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 0},
// 		locations.OptionItem{Name: "Tater Tots ", Description: "Tater Tots ", Price: 0}})
// 	_ = api.CreateOption("Options", "Options", 0, 0, heritageGolfLinksVeggieBlackBeanBurger.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Caramelized Onions", Description: "Caramelized Onions", Price: 0.5},
// 		locations.OptionItem{Name: "Jalepenos", Description: "Jalepenos", Price: 0.5},
// 		locations.OptionItem{Name: "Onion Ring", Description: "Onion Ring", Price: 0.5},
// 		locations.OptionItem{Name: "Fried Egg", Description: "Fried Egg", Price: 1},
// 		locations.OptionItem{Name: "Bacon", Description: "Bacon", Price: 1},
// 		locations.OptionItem{Name: "Mayonnaise", Description: "Mayonnaise", Price: 0},
// 		locations.OptionItem{Name: "Mustard", Description: "Mustard", Price: 0},
// 		locations.OptionItem{Name: "Ketchup", Description: "Ketchup", Price: 0}})
// 	_ = api.CreateOption("Cheese", "Cheese", 0, 0, heritageGolfLinksVeggieBlackBeanBurger.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Cheddar", Description: "Cheddar", Price: 0.5},
// 		locations.OptionItem{Name: "Swiss", Description: "Swiss", Price: 0.5},
// 		locations.OptionItem{Name: "American", Description: "American", Price: 0.5},
// 		locations.OptionItem{Name: "Provolone", Description: "Provolone", Price: 0.5},
// 		locations.OptionItem{Name: "Pepper Jack", Description: "Pepper Jack", Price: 0.5}})
// 	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, heritageGolfLinksVeggieBlackBeanBurger.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "No Onions", Description: "No Onions", Price: 0},
// 		locations.OptionItem{Name: "No Tomatoes", Description: "No Tomatoes", Price: 0},
// 		locations.OptionItem{Name: "No Cheese", Description: "No Cheese", Price: 0}})
// 	_ = api.CreateOption("Meat Temperature", "Meat Temperature", 1, 1, heritageGolfLinksBlackandBlueBurger.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Well Done", Description: "Well Done", Price: 0},
// 		locations.OptionItem{Name: "Medium-Well", Description: "Medium-Well", Price: 0},
// 		locations.OptionItem{Name: "Medium", Description: "Medium", Price: 0},
// 		locations.OptionItem{Name: "Medium-Rare", Description: "Medium-Rare", Price: 0},
// 		locations.OptionItem{Name: "Rare", Description: "Rare", Price: 0}})
// 	_ = api.CreateOption("Sides", "Sides", 1, 1, heritageGolfLinksBlackandBlueBurger.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Fries", Description: "Fries", Price: 0},
// 		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0},
// 		locations.OptionItem{Name: "Sweet Potato Fries ", Description: "Sweet Potato Fries ", Price: 0},
// 		locations.OptionItem{Name: "Cole Slaw", Description: "Cole Slaw", Price: 0},
// 		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 0},
// 		locations.OptionItem{Name: "Tater Tots ", Description: "Tater Tots ", Price: 0}})
// 	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, heritageGolfLinksBlackandBlueBurger.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "No Bacon", Description: "No Bacon", Price: 0}})
// 	_ = api.CreateOption("Sides", "Sides", 1, 1, heritageGolfLinksHeritageDog.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Fries", Description: "Fries", Price: 0},
// 		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0},
// 		locations.OptionItem{Name: "Sweet Potato Fries ", Description: "Sweet Potato Fries ", Price: 0},
// 		locations.OptionItem{Name: "Cole Slaw", Description: "Cole Slaw", Price: 0},
// 		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 0},
// 		locations.OptionItem{Name: "Tater Tots ", Description: "Tater Tots ", Price: 0}})
// 	_ = api.CreateOption("Sides", "Sides", 1, 1, heritageGolfLinksHeritageCheesyDog.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Fries", Description: "Fries", Price: 0},
// 		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0},
// 		locations.OptionItem{Name: "Sweet Potato Fries ", Description: "Sweet Potato Fries ", Price: 0},
// 		locations.OptionItem{Name: "Cole Slaw", Description: "Cole Slaw", Price: 0},
// 		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 0},
// 		locations.OptionItem{Name: "Tater Tots ", Description: "Tater Tots ", Price: 0}})
// 	_ = api.CreateOption("Sides", "Sides", 1, 1, heritageGolfLinksHeritageSouthwestDog.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Fries", Description: "Fries", Price: 0},
// 		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0},
// 		locations.OptionItem{Name: "Sweet Potato Fries ", Description: "Sweet Potato Fries ", Price: 0},
// 		locations.OptionItem{Name: "Cole Slaw", Description: "Cole Slaw", Price: 0},
// 		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 0},
// 		locations.OptionItem{Name: "Tater Tots ", Description: "Tater Tots ", Price: 0}})
// 	_ = api.CreateOption("Pizza Toppings", "Pizza Toppings", 0, 2, heritageGolfLinks12inchClassicMozzarellaAndMarinaraPizza.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Pepperoni", Description: "Pepperoni", Price: 0},
// 		locations.OptionItem{Name: "Ham", Description: "Ham", Price: 0},
// 		locations.OptionItem{Name: "Bacon", Description: "Bacon", Price: 0},
// 		locations.OptionItem{Name: "Sausage", Description: "Sausage", Price: 0},
// 		locations.OptionItem{Name: "Ground Beef", Description: "Ground Beef", Price: 0},
// 		locations.OptionItem{Name: "Pineapple", Description: "Pineapple", Price: 0},
// 		locations.OptionItem{Name: "Black Olives", Description: "Black Olives", Price: 0},
// 		locations.OptionItem{Name: "Mushrooms", Description: "Mushrooms", Price: 0},
// 		locations.OptionItem{Name: "Bell Peppers", Description: "Bell Peppers", Price: 0}})
// 	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, heritageGolfLinksBBQChickenPizza.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "No Spinach", Description: "No Spinach", Price: 0},
// 		locations.OptionItem{Name: "No Onions", Description: "No Onions", Price: 0},
// 		locations.OptionItem{Name: "No Bell Peppers", Description: "No Bell Peppers", Price: 0},
// 		locations.OptionItem{Name: "No Mushrooms", Description: "No Mushrooms", Price: 0}})
// 	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, heritageGolfLinksHawaiianPizza.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "No Onions", Description: "No Onions", Price: 0}})
// 	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, heritageGolfLinksTheLoadedSupreme.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "No Onions", Description: "No Onions", Price: 0},
// 		locations.OptionItem{Name: "No Mushrooms", Description: "No Mushrooms", Price: 0},
// 		locations.OptionItem{Name: "No Bell Peppers", Description: "No Bell Peppers", Price: 0}})
// 	_ = api.CreateOption("Flavor", "Flavor", 1, 1, heritageGolfLinksCannedSoda.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Coke", Description: "Coke", Price: 0},
// 		locations.OptionItem{Name: "Diet Coke", Description: "Diet Coke", Price: 0},
// 		locations.OptionItem{Name: "Sprite", Description: "Sprite", Price: 0},
// 		locations.OptionItem{Name: "Diet Sprite", Description: "Diet Sprite", Price: 0},
// 		locations.OptionItem{Name: "Dr. Pepper", Description: "Dr. Pepper", Price: 0},
// 		locations.OptionItem{Name: "Diet Dr. Pepper", Description: "Diet Dr. Pepper", Price: 0}})
// 	_ = api.CreateOption("Flavor", "Flavor", 1, 1, heritageGolfLinksFountainDrink.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Coke", Description: "Coke", Price: 0},
// 		locations.OptionItem{Name: "Diet Coke", Description: "Diet Coke", Price: 0},
// 		locations.OptionItem{Name: "Sprite", Description: "Sprite", Price: 0},
// 		locations.OptionItem{Name: "Diet Sprite", Description: "Diet Sprite", Price: 0},
// 		locations.OptionItem{Name: "Dr. Pepper", Description: "Dr. Pepper", Price: 0},
// 		locations.OptionItem{Name: "Diet Dr. Pepper", Description: "Diet Dr. Pepper", Price: 0}})
// 	_ = api.CreateOption("Flavor", "Flavor", 1, 1, heritageGolfLinksBottledGatorade.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Lemon-Lime", Description: "Lemon-Lime", Price: 0},
// 		locations.OptionItem{Name: "Fruit Punch", Description: "Fruit Punch", Price: 0},
// 		locations.OptionItem{Name: "Orange", Description: "Orange", Price: 0},
// 		locations.OptionItem{Name: "Glacier Freeze", Description: "Glacier Freeze", Price: 0},
// 		locations.OptionItem{Name: "Cool Blue", Description: "Cool Blue", Price: 0}})
// 	_ = api.CreateOption("Options", "Options", 0, 0, heritageGolfLinksBlueMoon.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Bucket of 4", Description: "Bucket of 4", Price: 10}})
// 	_ = api.CreateOption("Options", "Options", 0, 0, heritageGolfLinksBudLight.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Bucket of 4", Description: "Bucket of 4", Price: 10}})
// 	_ = api.CreateOption("Options", "Options", 0, 0, heritageGolfLinksBudLightLime.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Bucket of 4", Description: "Bucket of 4", Price: 10}})
// 	_ = api.CreateOption("Options", "Options", 0, 0, heritageGolfLinksBudweiser.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Bucket of 4", Description: "Bucket of 4", Price: 10}})
// 	_ = api.CreateOption("Options", "Options", 0, 0, heritageGolfLinksCoorsLight.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Bucket of 4", Description: "Bucket of 4", Price: 10}})
// 	_ = api.CreateOption("Options", "Options", 0, 0, heritageGolfLinksCorona.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Bucket of 4", Description: "Bucket of 4", Price: 10}})
// 	_ = api.CreateOption("Options", "Options", 0, 0, heritageGolfLinksCoronaLight.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Bucket of 4", Description: "Bucket of 4", Price: 10}})
// 	_ = api.CreateOption("Options", "Options", 0, 0, heritageGolfLinksDosEquisLager.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Bucket of 4", Description: "Bucket of 4", Price: 10}})
// 	_ = api.CreateOption("Options", "Options", 0, 0, heritageGolfLinksMichelobUltra.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Bucket of 4", Description: "Bucket of 4", Price: 10}})
// 	_ = api.CreateOption("Options", "Options", 0, 0, heritageGolfLinksMillerLite.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Bucket of 4", Description: "Bucket of 4", Price: 10}})
// 	_ = api.CreateOption("Options", "Options", 0, 0, heritageGolfLinksStellaArtois.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Bucket of 4", Description: "Bucket of 4", Price: 10}})
// 	_ = api.CreateOption("Options", "Options", 0, 0, heritageGolfLinksHeineken.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Bucket of 4", Description: "Bucket of 4", Price: 10}})
// 	_ = api.CreateOption("Flavor", "Flavor", 1, 1, heritageGolfLinksTruly.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Colima Lime", Description: "Colima Lime", Price: 0},
// 		locations.OptionItem{Name: "Grapefruit & Pomelo", Description: "Grapefruit & Pomelo", Price: 0},
// 		locations.OptionItem{Name: "Wild Berry", Description: "Wild Berry", Price: 0},
// 		locations.OptionItem{Name: "Lemon & Yuzu", Description: "Lemon & Yuzu", Price: 0},
// 		locations.OptionItem{Name: "Pomegranate", Description: "Pomegranate", Price: 0},
// 		locations.OptionItem{Name: "Blueberry & Acai", Description: "Blueberry & Acai", Price: 0},
// 		locations.OptionItem{Name: "Sicilian Blood Orange", Description: "Sicilian Blood Orange", Price: 0},
// 		locations.OptionItem{Name: "Raspberry & Lime", Description: "Raspberry & Lime", Price: 0},
// 		locations.OptionItem{Name: "Pineapple", Description: "Pineapple", Price: 0},
// 		locations.OptionItem{Name: "Passion Fruit", Description: "Passion Fruit", Price: 0},
// 		locations.OptionItem{Name: "Options", Description: "Options", Price: 0},
// 		locations.OptionItem{Name: "0", Description: "0", Price: 0},
// 		locations.OptionItem{Name: "Bucket of 4", Description: "Bucket of 4", Price: 10}})
// 	_ = api.CreateOption("Options", "Options", 0, 0, heritageGolfLinksTruly.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Bucket of 4", Description: "Bucket of 4", Price: 10}})
// 	_ = api.CreateOption("Flavor", "Flavor", 1, 1, heritageGolfLinksWhiteClaw.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Black Cherry", Description: "Black Cherry", Price: 0},
// 		locations.OptionItem{Name: "Ruby Grapefruit", Description: "Ruby Grapefruit", Price: 0},
// 		locations.OptionItem{Name: "Natural Lime", Description: "Natural Lime", Price: 0},
// 		locations.OptionItem{Name: "Raspberry", Description: "Raspberry", Price: 0},
// 		locations.OptionItem{Name: "Mango", Description: "Mango", Price: 0}})
// 	_ = api.CreateOption("Options", "Options", 0, 0, heritageGolfLinksWhiteClaw.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Bucket of 4", Description: "Bucket of 4", Price: 10}})
// 	_ = api.CreateOption("Vodka", "Vodka", 1, 1, heritageGolfLinksScrewDriver.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Grey Goose", Description: "Grey Goose", Price: 4},
// 		locations.OptionItem{Name: "Titos", Description: "Titos", Price: 0},
// 		locations.OptionItem{Name: "Effen", Description: "Effen", Price: 0},
// 		locations.OptionItem{Name: "Absolut", Description: "Absolut", Price: 0},
// 		locations.OptionItem{Name: "Ketel One", Description: "Ketel One", Price: 0},
// 		locations.OptionItem{Name: "Stoli", Description: "Stoli", Price: 0}})
// 	_ = api.CreateOption("Vodka", "Vodka", 1, 1, heritageGolfLinksBloodyMary.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Grey Goose", Description: "Grey Goose", Price: 4},
// 		locations.OptionItem{Name: "Titos", Description: "Titos", Price: 0},
// 		locations.OptionItem{Name: "Effen", Description: "Effen", Price: 0},
// 		locations.OptionItem{Name: "Absolut", Description: "Absolut", Price: 0},
// 		locations.OptionItem{Name: "Ketel One", Description: "Ketel One", Price: 0},
// 		locations.OptionItem{Name: "Stoli", Description: "Stoli", Price: 0}})
// 	_ = api.CreateOption("Vodka", "Vodka", 1, 1, heritageGolfLinksJohnDaly.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Grey Goose", Description: "Grey Goose", Price: 4},
// 		locations.OptionItem{Name: "Titos", Description: "Titos", Price: 0},
// 		locations.OptionItem{Name: "Effen", Description: "Effen", Price: 0},
// 		locations.OptionItem{Name: "Absolut", Description: "Absolut", Price: 0},
// 		locations.OptionItem{Name: "Ketel One", Description: "Ketel One", Price: 0},
// 		locations.OptionItem{Name: "Stoli", Description: "Stoli", Price: 0}})
// 	_ = api.CreateOption("Tequila", "Tequila", 1, 1, heritageGolfLinksPaloma.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Jose Cuervo", Description: "Jose Cuervo", Price: 0},
// 		locations.OptionItem{Name: "Patron", Description: "Patron", Price: 4},
// 		locations.OptionItem{Name: "Suaza", Description: "Suaza", Price: 0},
// 		locations.OptionItem{Name: "1800", Description: "1800", Price: 0},
// 		locations.OptionItem{Name: "Hornitos", Description: "Hornitos", Price: 2},
// 		locations.OptionItem{Name: "Don Julio", Description: "Don Julio", Price: 4}})
// 	_ = api.CreateOption("Vodka", "Vodka", 1, 1, heritageGolfLinksMoscowMule.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Grey Goose", Description: "Grey Goose", Price: 4},
// 		locations.OptionItem{Name: "Titos", Description: "Titos", Price: 0},
// 		locations.OptionItem{Name: "Effen", Description: "Effen", Price: 0},
// 		locations.OptionItem{Name: "Absolut", Description: "Absolut", Price: 0},
// 		locations.OptionItem{Name: "Ketel One", Description: "Ketel One", Price: 0},
// 		locations.OptionItem{Name: "Stoli", Description: "Stoli", Price: 0}})
// 	_ = api.CreateOption("Rum", "Rum", 1, 1, heritageGolfLinksRumAndCoke.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Bacardi", Description: "Bacardi", Price: 0},
// 		locations.OptionItem{Name: "Captain Morgan", Description: "Captain Morgan", Price: 0},
// 		locations.OptionItem{Name: "Tanduay", Description: "Tanduay", Price: 0}})
// 	_ = api.CreateOption("Options", "Options", 0, 0, heritageGolfLinksRumAndCoke.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Diet Coke", Description: "Diet Coke", Price: 0}})
// 	_ = api.CreateOption("Whiskey", "Whiskey", 1, 1, heritageGolfLinksWhiskeyAndCoke.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Jack Daniels", Description: "Jack Daniels", Price: 0},
// 		locations.OptionItem{Name: "Seagrams", Description: "Seagrams", Price: 0},
// 		locations.OptionItem{Name: "Makers Mark", Description: "Makers Mark", Price: 0},
// 		locations.OptionItem{Name: "Crown Royal", Description: "Crown Royal", Price: 0},
// 		locations.OptionItem{Name: "Jameson", Description: "Jameson", Price: 0},
// 		locations.OptionItem{Name: "Jim Beam", Description: "Jim Beam", Price: 0}})
// 	_ = api.CreateOption("Options", "Options", 0, 0, heritageGolfLinksWhiskeyAndCoke.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Diet Coke", Description: "Diet Coke", Price: 0}})
// 	_ = api.CreateOption("Vodka", "Vodka", 1, 1, heritageGolfLinksAzelea.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Grey Goose", Description: "Grey Goose", Price: 4},
// 		locations.OptionItem{Name: "Titos", Description: "Titos", Price: 0},
// 		locations.OptionItem{Name: "Effen", Description: "Effen", Price: 0},
// 		locations.OptionItem{Name: "Absolut", Description: "Absolut", Price: 0},
// 		locations.OptionItem{Name: "Ketel One", Description: "Ketel One", Price: 0},
// 		locations.OptionItem{Name: "Stoli", Description: "Stoli", Price: 0}})
// 	_ = api.CreateOption("Vodka", "Vodka", 1, 1, heritageGolfLinksTransfusion.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Grey Goose", Description: "Grey Goose", Price: 4},
// 		locations.OptionItem{Name: "Titos", Description: "Titos", Price: 0},
// 		locations.OptionItem{Name: "Effen", Description: "Effen", Price: 0},
// 		locations.OptionItem{Name: "Absolut", Description: "Absolut", Price: 0},
// 		locations.OptionItem{Name: "Ketel One", Description: "Ketel One", Price: 0},
// 		locations.OptionItem{Name: "Stoli", Description: "Stoli", Price: 0}})
// 	_ = api.CreateOption("Vodka", "Vodka", 1, 1, heritageGolfLinksRedBullAndVodka.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Grey Goose", Description: "Grey Goose", Price: 4},
// 		locations.OptionItem{Name: "Titos", Description: "Titos", Price: 0},
// 		locations.OptionItem{Name: "Effen", Description: "Effen", Price: 0},
// 		locations.OptionItem{Name: "Absolut", Description: "Absolut", Price: 0},
// 		locations.OptionItem{Name: "Ketel One", Description: "Ketel One", Price: 0},
// 		locations.OptionItem{Name: "Stoli", Description: "Stoli", Price: 0}})
// 	_ = api.CreateOption("Flavor", "Flavor", 1, 1, heritageGolfLinksCutwaterCannedCocktail.ID, []locations.OptionItem{
// 		locations.OptionItem{Name: "Tequila Margarita", Description: "Tequila Margarita", Price: 0},
// 		locations.OptionItem{Name: "Vodka Mule", Description: "Vodka Mule", Price: 0},
// 		locations.OptionItem{Name: "Lime Vodka Soda", Description: "Lime Vodka Soda", Price: 0},
// 		locations.OptionItem{Name: "Cucumber Vodka Soda", Description: "Cucumber Vodka Soda", Price: 0},
// 		locations.OptionItem{Name: "Tequila Paloma", Description: "Tequila Paloma", Price: 0},
// 		locations.OptionItem{Name: "Rum & Cola", Description: "Rum & Cola", Price: 0},
// 		locations.OptionItem{Name: "Rum Mint Mojito", Description: "Rum Mint Mojito", Price: 0},
// 		locations.OptionItem{Name: "Spicy Bloody Mary", Description: "Spicy Bloody Mary", Price: 0},
// 		locations.OptionItem{Name: "Gin & Tonic", Description: "Gin & Tonic", Price: 0},
// 		locations.OptionItem{Name: "White Russian", Description: "White Russian", Price: 0}})
// 	// seed orders
// 	log.Println("Seeding Completed")
// }
