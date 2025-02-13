package seeding

import (
	"github.com/grow-assistant/menu-creator/internal/models"
)

func CreateAmeliaRiverClub() *models.Location {
	// seed locations
	location := models.CreateLocation("The Amelia River Club", "Fernandina Beach, FL")
	// Create menu
	menu := models.CreateMenu("River Bar & Grill", "River Bar & Grill", location.ID)
	location.Menus = append(location.Menus, *menu)
	// Create categories
	fairwayGreens := models.CreateCategory("The Teebox Fairway Greens", "The Teebox Fairway Greens", menu.ID)
	atTheTurn := models.CreateCategory("At The Turn", "At The Turn", menu.ID)
	handhelds := models.CreateCategory("Handhelds", "Handhelds", menu.ID)
	breakfast := models.CreateCategory("Breakfast", "Breakfast", menu.ID)
	beverages := models.CreateCategory("Beverages", "Beverages", menu.ID)
	cocktails := models.CreateCategory("Cocktails", "Cocktails", menu.ID)

	menu.Categories = append(menu.Categories,
		*fairwayGreens,
		*atTheTurn,
		*handhelds,
		*breakfast,
		*beverages,
		*cocktails,
	)
	// Create items
	chiliNachos := models.CreateItem(
		"Chili Cheese Nachos",
		"Smothered in chili, cheddar cheese, lettuce, tomato & topped with jalapenos. Served with a side of sour cream & salsa. Add chicken for $4",
		12,
		fairwayGreens.ID,
	)
	fairwayGreens.Items = append(fairwayGreens.Items, *chiliNachos)

	chipsAndSalsa := models.CreateItem("Chips & Salsa", "", 4, fairwayGreens.ID)
	fairwayGreens.Items = append(fairwayGreens.Items, *chipsAndSalsa)

	chickenWings := models.CreateItem(
		"Chicken Wings",
		"8 chicken wings tossed in our house sauce served with ranch or blue cheese",
		12,
		fairwayGreens.ID,
	)
	fairwayGreens.Items = append(fairwayGreens.Items, *chickenWings)
	onionRings := models.CreateItem(
		"Onion Rings",
		"Served with a spicy aioli",
		8,
		fairwayGreens.ID,
	)
	fairwayGreens.Items = append(fairwayGreens.Items, *onionRings)

	signatureSalad := models.CreateItem(
		"Signature Salad",
		"Mixed Greens tossed in an herb vinaigrette topped with feta cheese, artichoke hearts, roasted red peppers, cucumbers, & kalamata olives",
		12,
		fairwayGreens.ID,
	)
	fairwayGreens.Items = append(fairwayGreens.Items, *signatureSalad)

	caesarSalad := models.CreateItem(
		"Caesar Salad",
		"Traditional Caesar with herb butter croutons",
		11,
		fairwayGreens.ID,
	)
	fairwayGreens.Items = append(fairwayGreens.Items, *caesarSalad)

	houseSalad := models.CreateItem(
		"House Salad",
		"Mixed greens, cherry tomatoes, cucumbers, red onion, cheddar cheese. Served with a balsamic vinaigrette",
		10,
		fairwayGreens.ID,
	)
	fairwayGreens.Items = append(fairwayGreens.Items, *houseSalad)

	parThree := models.CreateItem(
		"The Par Three",
		"Chicken salad, tuna salad & egg salad with leaf lettuce & sliced tomatoes",
		14,
		fairwayGreens.ID,
	)
	fairwayGreens.Items = append(fairwayGreens.Items, *parThree)
	// Create At The Turn items
	hotDog := models.CreateItem(
		"Hot Dog",
		"Top with relish, onions or cole slaw. Chili cheese dog for $2. Served with kettle chips",
		9,
		atTheTurn.ID,
	)
	atTheTurn.Items = append(atTheTurn.Items, *hotDog)

	beerBrat := models.CreateItem(
		"IPA Beer Brat",
		"Smothered in grilled peppers & onions. Served with kettle chips",
		10,
		atTheTurn.ID,
	)
	atTheTurn.Items = append(atTheTurn.Items, *beerBrat)

	cheeseburger := models.CreateItem(
		"Cheeseburger",
		"Lettuce, tomato & onion with your choice of American, Cheddar, Provolone or Swiss cheese. Served with fries",
		14,
		atTheTurn.ID,
	)
	atTheTurn.Items = append(atTheTurn.Items, *cheeseburger)

	chickenSandwich := models.CreateItem(
		"Chicken Sandwich",
		"Grilled chicken breast with Provolone lettuce, tomato & onion. Served with fries. Add bacon $2",
		13,
		atTheTurn.ID,
	)
	atTheTurn.Items = append(atTheTurn.Items, *chickenSandwich)

	riverTacos := models.CreateItem(
		"River Tacos",
		"Two flour tortillas served with a chipotle lime, lettuce, tomato & cheddar cheese. Served with chips & salsa",
		13,
		atTheTurn.ID,
	)
	atTheTurn.Items = append(atTheTurn.Items, *riverTacos)
	// Create Handhelds items
	classicRuben := models.CreateItem(
		"Classic Ruben",
		"Corned beef with melted Swiss cheese, sourkraut & thousand island dressing on toasted rye bread",
		13,
		handhelds.ID,
	)
	handhelds.Items = append(handhelds.Items, *classicRuben)

	golfClubWrap := models.CreateItem(
		"Golf Club Wrap",
		"Turkey, bacon, lettuce, tomato, Swiss cheese & chipotle mayo in a spinach wrap",
		12,
		handhelds.ID,
	)
	handhelds.Items = append(handhelds.Items, *golfClubWrap)

	sandwedge := models.CreateItem(
		"Sandwedge",
		"Choice of BLT, turkey, chicken salad, or tuna salad on wheat bread with lettuce, tomato & onion",
		10,
		handhelds.ID,
	)
	handhelds.Items = append(handhelds.Items, *sandwedge)

	ahiTunaWrap := models.CreateItem(
		"Ahi Tuna Wrap",
		"Seared Ahi Tuna with lettuce, tomato, wasabi aoli & served cold in a spinach wrap",
		15,
		handhelds.ID,
	)
	handhelds.Items = append(handhelds.Items, *ahiTunaWrap)
	// Create Breakfast items
	breakfastPod := models.CreateItem(
		"Breakfast Pod",
		"One fried egg on a flour tortilla with American cheese & bacon",
		5,
		breakfast.ID,
	)
	breakfast.Items = append(breakfast.Items, *breakfastPod)

	breakfastSandwich := models.CreateItem(
		"Breakfast Sandwich",
		"Two fried eggs with American cheese, bacon, sausage or ham, on toast",
		10,
		breakfast.ID,
	)
	breakfast.Items = append(breakfast.Items, *breakfastSandwich)

	morningRiverWrap := models.CreateItem(
		"Morning River Wrap",
		"Scrambled eggs, sautéed onions and peppers. Choice of bacon, ham or sausage, with cheddar cheese & salsa in a garlic herb tortilla",
		12,
		breakfast.ID,
	)
	breakfast.Items = append(breakfast.Items, *morningRiverWrap)
	holeInOneBreakfast := models.CreateItem(
		"Hole In One Breakfast",
		"Two eggs cooked your way, with bacon, sausage, or ham, toast & tater tots",
		12,
		breakfast.ID,
	)
	breakfast.Items = append(breakfast.Items, *holeInOneBreakfast)

	cinnamonFrenchToast := models.CreateItem(
		"Cinnamon Brioche French Toast",
		"French toast with powdered sugar & maple syrup",
		12,
		breakfast.ID,
	)
	breakfast.Items = append(breakfast.Items, *cinnamonFrenchToast)

	hangoverBowl := models.CreateItem(
		"Hangover Bowl",
		"Two over easy eggs with choice of meat over grits, topped with cheddar cheese",
		12,
		breakfast.ID,
	)
	breakfast.Items = append(breakfast.Items, *hangoverBowl)

	cornedBeefHash := models.CreateItem(
		"Corned Beef Hash",
		"Corned beef hash with two eggs cooked your way & toast",
		12,
		breakfast.ID,
	)
	breakfast.Items = append(breakfast.Items, *cornedBeefHash)

	chefsQuiche := models.CreateItem(
		"Chef's Quiche",
		"Bacon, ham, Swiss cheese, sautéed peppers and green onions",
		8,
		breakfast.ID,
	)
	breakfast.Items = append(breakfast.Items, *chefsQuiche)

	breakfastSides := models.CreateItem(
		"Breakfast Sides",
		"Bacon, ham, Swiss cheese, sautéed peppers and green onions",
		4,
		breakfast.ID,
	)
	breakfast.Items = append(breakfast.Items, *breakfastSides)
	icedTea := models.CreateItem(
		"Iced Tea",
		"",
		3,
		beverages.ID,
	)
	beverages.Items = append(beverages.Items, *icedTea)

	bottledWater := models.CreateItem(
		"Bottled Water",
		"",
		3,
		beverages.ID,
	)
	beverages.Items = append(beverages.Items, *bottledWater)

	fountainDrink := models.CreateItem(
		"Fountain Drink",
		"Pepsi, Diet Pepsi, Sierra Mist, Dr. Pepper, Mug Root Beer",
		3,
		beverages.ID,
	)
	beverages.Items = append(beverages.Items, *fountainDrink)

	pinkLemonade := models.CreateItem(
		"Pink Lemonade",
		"",
		3,
		beverages.ID,
	)
	beverages.Items = append(beverages.Items, *pinkLemonade)
	screwDriver := models.CreateItem(
		"Screw Driver",
		"Vodka and Fresh Squeezed Orange Juice",
		10,
		cocktails.ID,
	)
	cocktails.Items = append(cocktails.Items, *screwDriver)

	bloodyMary := models.CreateItem(
		"Bloody Mary",
		"Vodka, house bloody mary mix, fresh celery, fresh-grated horseradish, rosemary skewered olive, cherry tomato, spicy pepper",
		13,
		cocktails.ID,
	)
	cocktails.Items = append(cocktails.Items, *bloodyMary)

	mimosa := models.CreateItem(
		"Mimosa",
		"Sparkling Wine, Orange Juice",
		8,
		cocktails.ID,
	)
	cocktails.Items = append(cocktails.Items, *mimosa)

	johnDaly := models.CreateItem(
		"John Daly",
		"Fresh Brewed Iced Tea, Lemonade, and Vodka.",
		12,
		cocktails.ID,
	)
	cocktails.Items = append(cocktails.Items, *johnDaly)

	paloma := models.CreateItem(
		"Paloma",
		"Tequila, Grapefruit Juice, Fresh Lime Juice, Agave Syrup, Club Soda",
		11,
		cocktails.ID,
	)
	cocktails.Items = append(cocktails.Items, *paloma)

	moscowMule := models.CreateItem(
		"Moscow Mule",
		"Vodka, lime juice, mint, ginger beer",
		11,
		cocktails.ID,
	)
	cocktails.Items = append(cocktails.Items, *moscowMule)

	rumAndCoke := models.CreateItem(
		"Rum & Coke",
		"",
		11,
		cocktails.ID,
	)
	cocktails.Items = append(cocktails.Items, *rumAndCoke)

	whiskeyAndCoke := models.CreateItem(
		"Whiskey & Coke",
		"",
		11,
		cocktails.ID,
	)
	cocktails.Items = append(cocktails.Items, *whiskeyAndCoke)

	azelea := models.CreateItem(
		"Azelea",
		"Lemon juice, pineapple juice, vodka and grenadine.",
		11,
		cocktails.ID,
	)
	cocktails.Items = append(cocktails.Items, *azelea)

	cranberryKamikaze := models.CreateItem(
		"Cranberry Kamikaze",
		"Lime juice, triple sec, cranberry juice and vodka.",
		11,
		cocktails.ID,
	)
	cocktails.Items = append(cocktails.Items, *cranberryKamikaze)

	transfusion := models.CreateItem(
		"Transfusion",
		"Vodka, ginger ale and grape juice",
		11,
		cocktails.ID,
	)
	cocktails.Items = append(cocktails.Items, *transfusion)

	redBullAndVodka := models.CreateItem(
		"Red Bull & Vodka",
		"",
		11,
		cocktails.ID,
	)
	cocktails.Items = append(cocktails.Items, *redBullAndVodka)
	// Create options
	extras := models.CreateOption(
		"Extras",
		"Extras",
		0,
		0,
		chiliNachos.ID,
		[]models.OptionItem{
			{Name: "Add Chicken", Description: "Add Chicken", Price: 4},
		},
	)
	chiliNachos.Options = append(chiliNachos.Options, *extras)
	ingredients := models.CreateOption(
		"Ingredients",
		"Ingredients",
		0,
		0,
		chiliNachos.ID,
		[]models.OptionItem{
			{Name: "No Lettuce", Description: "No Lettuce", Price: 0},
			{Name: "No Jalapenos", Description: "No Jalapenos", Price: 0},
			{Name: "No Tomatos", Description: "No Tomatos", Price: 0},
		},
	)
	chiliNachos.Options = append(chiliNachos.Options, *ingredients)

	return location
	dippingSauce := models.CreateOption(
		"Dipping Sauce",
		"Dipping Sauce",
		1,
		1,
		chickenWings.ID,
		[]models.OptionItem{
			{Name: "Ranch", Description: "Ranch", Price: 0},
			{Name: "Blue Cheese", Description: "Blue Cheese", Price: 0},
		},
	)
	chickenWings.Options = append(chickenWings.Options, *dippingSauce)
	ingredients := models.CreateOption(
		"Ingredients",
		"Ingredients",
		0,
		0,
		chickenWings.ID,
		[]models.OptionItem{
			{Name: "No House Sauce", Description: "No House Sauce", Price: 0},
		},
	)
	chickenWings.Options = append(chickenWings.Options, *ingredients)
	addProtein := models.CreateOption(
		"Add Protein",
		"Add Protein",
		0,
		0,
		onionRings.ID,
		[]models.OptionItem{
			{Name: "Salmon", Description: "Salmon", Price: 5},
			{Name: "Mahi Mahi", Description: "Mahi Mahi", Price: 5},
			{Name: "Chicken", Description: "Chicken", Price: 4},
			{Name: "Shrimp", Description: "Shrimp", Price: 5},
			{Name: "Ahi Tuna", Description: "Ahi Tuna", Price: 6},
		},
	)
	onionRings.Options = append(onionRings.Options, *addProtein)
	_ = api.CreateOption("Add Protein", "Add Protein", 0, 0, theAmeliaRiverClubRiverBarandGrillTheTeeboxFairwayGreensSignatureSalad.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Salmon", Description: "Salmon", Price: 5},
		locations.OptionItem{Name: "Mahi Mahi", Description: "Mahi Mahi", Price: 5},
		locations.OptionItem{Name: "Chicken", Description: "Chicken", Price: 4},
		locations.OptionItem{Name: "Shrimp", Description: "Shrimp", Price: 5},
		locations.OptionItem{Name: "Ahi Tuna", Description: "Ahi Tuna", Price: 6}})
	_ = api.CreateOption("Ingredients", "Ingredients", 0, 0, theAmeliaRiverClubRiverBarandGrillTheTeeboxFairwayGreensSignatureSalad.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Feta Cheese", Description: "No Feta Cheese", Price: 0},
		locations.OptionItem{Name: "No Artichoke Hearts", Description: "No Artichoke Hearts", Price: 0},
		locations.OptionItem{Name: "No Roasted Red Peppers", Description: "No Roasted Red Peppers", Price: 0},
		locations.OptionItem{Name: "No Cucumbers", Description: "No Cucumbers", Price: 0},
		locations.OptionItem{Name: "No Kalamata Olives", Description: "No Kalamata Olives", Price: 0}})
	_ = api.CreateOption("Add Protein", "Add Protein", 0, 0, theAmeliaRiverClubRiverBarandGrillTheTeeboxFairwayGreensCaesarSalad.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Salmon", Description: "Salmon", Price: 5},
		locations.OptionItem{Name: "Mahi Mahi", Description: "Mahi Mahi", Price: 5},
		locations.OptionItem{Name: "Chicken", Description: "Chicken", Price: 4},
		locations.OptionItem{Name: "Shrimp", Description: "Shrimp", Price: 5},
		locations.OptionItem{Name: "Ahi Tuna", Description: "Ahi Tuna", Price: 6}})
	_ = api.CreateOption("Ingredients", "Ingredients", 0, 0, theAmeliaRiverClubRiverBarandGrillTheTeeboxFairwayGreensCaesarSalad.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Croutons", Description: "No Croutons", Price: 0}})
	_ = api.CreateOption("Add Protein", "Add Protein", 0, 0, theAmeliaRiverClubRiverBarandGrillTheTeeboxFairwayGreensHouseSalad.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Salmon", Description: "Salmon", Price: 5},
		locations.OptionItem{Name: "Mahi Mahi", Description: "Mahi Mahi", Price: 5},
		locations.OptionItem{Name: "Chicken", Description: "Chicken", Price: 4},
		locations.OptionItem{Name: "Shrimp", Description: "Shrimp", Price: 5},
		locations.OptionItem{Name: "Ahi Tuna", Description: "Ahi Tuna", Price: 6}})
	_ = api.CreateOption("Ingredients", "Ingredients", 0, 0, theAmeliaRiverClubRiverBarandGrillTheTeeboxFairwayGreensHouseSalad.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Tomatoes", Description: "No Tomatoes", Price: 0},
		locations.OptionItem{Name: "No Cucumbers", Description: "No Cucumbers", Price: 0},
		locations.OptionItem{Name: "No Red Onion", Description: "No Red Onion", Price: 0},
		locations.OptionItem{Name: "No Cheddar Cheese", Description: "No Cheddar Cheese", Price: 0}})
	_ = api.CreateOption("Options", "Options", 1, 1, theAmeliaRiverClubRiverBarandGrillTheTeeboxFairwayGreensTheParThree.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Chicken Salad", Description: "Chicken Salad", Price: 0},
		locations.OptionItem{Name: "Tuna Salad", Description: "Tuna Salad", Price: 0},
		locations.OptionItem{Name: "Egg Salad", Description: "Egg Salad", Price: 0}})
	_ = api.CreateOption("Ingredients", "Ingredients", 0, 0, theAmeliaRiverClubRiverBarandGrillTheTeeboxFairwayGreensTheParThree.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Lettuce", Description: "No Lettuce", Price: 0},
		locations.OptionItem{Name: "No Tomatoes", Description: "No Tomatoes", Price: 0}})
	_ = api.CreateOption("Extras", "Extras", 0, 0, theAmeliaRiverClubRiverBarandGrillAtTheTurnHotDog.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Add Chili Cheese", Description: "Add Chili Cheese", Price: 2}})
	_ = api.CreateOption("Ingredients", "Ingredients", 0, 0, theAmeliaRiverClubRiverBarandGrillAtTheTurnHotDog.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Relish", Description: "No Relish", Price: 0},
		locations.OptionItem{Name: "No Onions", Description: "No Onions", Price: 0},
		locations.OptionItem{Name: "No Cole Slaw", Description: "No Cole Slaw", Price: 0}})
	_ = api.CreateOption("Ingredients", "Ingredients", 0, 0, theAmeliaRiverClubRiverBarandGrillAtTheTurnIPABeerBrat.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Peppers", Description: "No Peppers", Price: 0},
		locations.OptionItem{Name: "No Onions", Description: "No Onions", Price: 0}})
	_ = api.CreateOption("Cheese", "Cheese", 1, 1, theAmeliaRiverClubRiverBarandGrillAtTheTurnCheeseburger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "American", Description: "American", Price: 0},
		locations.OptionItem{Name: "Cheddar", Description: "Cheddar", Price: 0},
		locations.OptionItem{Name: "Provolone", Description: "Provolone", Price: 0},
		locations.OptionItem{Name: "Swiss Cheese", Description: "Swiss Cheese", Price: 0},
		locations.OptionItem{Name: "No Cheese", Description: "No Cheese", Price: 0}})
	_ = api.CreateOption("Ingredients", "Ingredients", 0, 0, theAmeliaRiverClubRiverBarandGrillAtTheTurnCheeseburger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Lettuce", Description: "No Lettuce", Price: 0},
		locations.OptionItem{Name: "No Tomato", Description: "No Tomato", Price: 0},
		locations.OptionItem{Name: "No Onion", Description: "No Onion", Price: 0}})
	_ = api.CreateOption("Extras", "Extras", 0, 0, theAmeliaRiverClubRiverBarandGrillAtTheTurnChickenSandwich.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Add Bacon", Description: "Add Bacon", Price: 2}})
	_ = api.CreateOption("Ingredients", "Ingredients", 0, 0, theAmeliaRiverClubRiverBarandGrillAtTheTurnChickenSandwich.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Lettuce", Description: "No Lettuce", Price: 0},
		locations.OptionItem{Name: "No Tomato", Description: "No Tomato", Price: 0},
		locations.OptionItem{Name: "No Onion", Description: "No Onion", Price: 0},
		locations.OptionItem{Name: "No Cheese", Description: "No Cheese", Price: 0}})
	_ = api.CreateOption("Options", "Options", 1, 1, theAmeliaRiverClubRiverBarandGrillAtTheTurnRiverTacos.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Chicken", Description: "Chicken", Price: 0},
		locations.OptionItem{Name: "Mahi Mahi", Description: "Mahi Mahi", Price: 2}})
	_ = api.CreateOption("Ingredients", "Ingredients", 0, 0, theAmeliaRiverClubRiverBarandGrillAtTheTurnRiverTacos.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Lettuce", Description: "No Lettuce", Price: 0},
		locations.OptionItem{Name: "No Tomato", Description: "No Tomato", Price: 0},
		locations.OptionItem{Name: "No Cheese", Description: "No Cheese", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, theAmeliaRiverClubRiverBarandGrillHandheldsClassicRuben.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Kettle Chips", Description: "Kettle Chips", Price: 0},
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Cole Slaw", Description: "Cole Slaw", Price: 0}})
	_ = api.CreateOption("Ingredients", "Ingredients", 0, 0, theAmeliaRiverClubRiverBarandGrillHandheldsClassicRuben.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Swiss Cheese", Description: "No Swiss Cheese", Price: 0},
		locations.OptionItem{Name: "No Sourkraut", Description: "No Sourkraut", Price: 0},
		locations.OptionItem{Name: "No Dressing", Description: "No Dressing", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, theAmeliaRiverClubRiverBarandGrillHandheldsGolfClubWrap.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Kettle Chips", Description: "Kettle Chips", Price: 0},
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Cole Slaw", Description: "Cole Slaw", Price: 0}})
	_ = api.CreateOption("Ingredients", "Ingredients", 0, 0, theAmeliaRiverClubRiverBarandGrillHandheldsGolfClubWrap.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Lettuce", Description: "No Lettuce", Price: 0},
		locations.OptionItem{Name: "No Tomato", Description: "No Tomato", Price: 0},
		locations.OptionItem{Name: "No Swiss Cheese", Description: "No Swiss Cheese", Price: 0},
		locations.OptionItem{Name: "No Bacon", Description: "No Bacon", Price: 0},
		locations.OptionItem{Name: "No Mayo", Description: "No Mayo", Price: 0}})
	_ = api.CreateOption("Options", "Options", 1, 1, theAmeliaRiverClubRiverBarandGrillHandheldsSandwedge.ID, []locations.OptionItem{
		locations.OptionItem{Name: "BLT", Description: "BLT", Price: 0},
		locations.OptionItem{Name: "Turkey", Description: "Turkey", Price: 0},
		locations.OptionItem{Name: "Chicken Salad", Description: "Chicken Salad", Price: 0},
		locations.OptionItem{Name: "Tuna Salad", Description: "Tuna Salad", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, theAmeliaRiverClubRiverBarandGrillHandheldsSandwedge.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Kettle Chips", Description: "Kettle Chips", Price: 0},
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Cole Slaw", Description: "Cole Slaw", Price: 0}})
	_ = api.CreateOption("Ingredients", "Ingredients", 0, 0, theAmeliaRiverClubRiverBarandGrillHandheldsSandwedge.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Lettuce", Description: "No Lettuce", Price: 0},
		locations.OptionItem{Name: "No Tomato", Description: "No Tomato", Price: 0},
		locations.OptionItem{Name: "No Onion", Description: "No Onion", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, theAmeliaRiverClubRiverBarandGrillHandheldsAhiTunaWrap.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Kettle Chips", Description: "Kettle Chips", Price: 0},
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Cole Slaw", Description: "Cole Slaw", Price: 0}})
	_ = api.CreateOption("Ingredients", "Ingredients", 0, 0, theAmeliaRiverClubRiverBarandGrillHandheldsAhiTunaWrap.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Lettuce", Description: "No Lettuce", Price: 0},
		locations.OptionItem{Name: "No Tomato", Description: "No Tomato", Price: 0},
		locations.OptionItem{Name: "No Wasabi Aioli", Description: "No Wasabi Aioli", Price: 0}})
	_ = api.CreateOption("Ingredients", "Ingredients", 0, 0, theAmeliaRiverClubRiverBarandGrillBreakfastBreakfastPod.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Cheese", Description: "No Cheese", Price: 0},
		locations.OptionItem{Name: "No Bacon", Description: "No Bacon", Price: 0}})
	_ = api.CreateOption("Choice of Meat", "Choice of Meat", 1, 1, theAmeliaRiverClubRiverBarandGrillBreakfastBreakfastSandwich.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Bacon", Description: "Bacon", Price: 0},
		locations.OptionItem{Name: "Sausage", Description: "Sausage", Price: 0},
		locations.OptionItem{Name: "Ham", Description: "Ham", Price: 0}})
	_ = api.CreateOption("Ingredients", "Ingredients", 0, 0, theAmeliaRiverClubRiverBarandGrillBreakfastBreakfastSandwich.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Cheese", Description: "No Cheese", Price: 0}})
	_ = api.CreateOption("Choice of Meat", "Choice of Meat", 1, 1, theAmeliaRiverClubRiverBarandGrillBreakfastMorningRiverWrap.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Bacon", Description: "Bacon", Price: 0},
		locations.OptionItem{Name: "Sausage", Description: "Sausage", Price: 0},
		locations.OptionItem{Name: "Ham", Description: "Ham", Price: 0}})
	_ = api.CreateOption("Ingredients", "Ingredients", 0, 0, theAmeliaRiverClubRiverBarandGrillBreakfastMorningRiverWrap.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Onions", Description: "No Onions", Price: 0},
		locations.OptionItem{Name: "No Peppers", Description: "No Peppers", Price: 0},
		locations.OptionItem{Name: "No Cheese", Description: "No Cheese", Price: 0},
		locations.OptionItem{Name: "No Salsa", Description: "No Salsa", Price: 0}})
	_ = api.CreateOption("Eggs", "Eggs", 1, 1, theAmeliaRiverClubRiverBarandGrillBreakfastHoleInOneBreakfast.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Over Easy", Description: "Over Easy", Price: 0},
		locations.OptionItem{Name: "Over Medium", Description: "Over Medium", Price: 0},
		locations.OptionItem{Name: "Over Hard", Description: "Over Hard", Price: 0},
		locations.OptionItem{Name: "Scrambled", Description: "Scrambled", Price: 0},
		locations.OptionItem{Name: "Sunny Side Up", Description: "Sunny Side Up", Price: 0}})
	_ = api.CreateOption("Choice of Meat", "Choice of Meat", 1, 1, theAmeliaRiverClubRiverBarandGrillBreakfastHoleInOneBreakfast.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Bacon", Description: "Bacon", Price: 0},
		locations.OptionItem{Name: "Sausage", Description: "Sausage", Price: 0},
		locations.OptionItem{Name: "Ham", Description: "Ham", Price: 0}})
	_ = api.CreateOption("Eggs", "Eggs", 1, 1, theAmeliaRiverClubRiverBarandGrillBreakfastHangoverBowl.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Over Easy", Description: "Over Easy", Price: 0},
		locations.OptionItem{Name: "Over Medium", Description: "Over Medium", Price: 0},
		locations.OptionItem{Name: "Over Hard", Description: "Over Hard", Price: 0},
		locations.OptionItem{Name: "Scrambled", Description: "Scrambled", Price: 0},
		locations.OptionItem{Name: "Sunny Side Up", Description: "Sunny Side Up", Price: 0}})
	_ = api.CreateOption("Ingredients", "Ingredients", 0, 0, theAmeliaRiverClubRiverBarandGrillBreakfastHangoverBowl.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Cheese", Description: "No Cheese", Price: 0},
		locations.OptionItem{Name: "No Grits", Description: "No Grits", Price: 0}})
	_ = api.CreateOption("Eggs", "Eggs", 1, 1, theAmeliaRiverClubRiverBarandGrillBreakfastCornedBeefHash.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Over Easy", Description: "Over Easy", Price: 0},
		locations.OptionItem{Name: "Over Medium", Description: "Over Medium", Price: 0},
		locations.OptionItem{Name: "Over Hard", Description: "Over Hard", Price: 0},
		locations.OptionItem{Name: "Scrambled", Description: "Scrambled", Price: 0},
		locations.OptionItem{Name: "Sunny Side Up", Description: "Sunny Side Up", Price: 0}})
	_ = api.CreateOption("Ingredients", "Ingredients", 0, 0, theAmeliaRiverClubRiverBarandGrillBreakfastCornedBeefHash.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Toast", Description: "No Toast", Price: 0}})
	_ = api.CreateOption("Options", "Options", 1, 5, theAmeliaRiverClubRiverBarandGrillBreakfastBreakfastSides.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Tator Tots", Description: "Tator Tots", Price: 0},
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Bacon", Description: "Bacon", Price: 0},
		locations.OptionItem{Name: "Ham", Description: "Ham", Price: 0},
		locations.OptionItem{Name: "Grits", Description: "Grits", Price: 0},
		locations.OptionItem{Name: "Sausage", Description: "Sausage", Price: 0},
		locations.OptionItem{Name: "Muffin", Description: "Muffin", Price: 0}})
	_ = api.CreateOption("Flavor", "Flavor", 1, 1, theAmeliaRiverClubFountainDrink.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Coke", Description: "Coke", Price: 0},
		locations.OptionItem{Name: "Diet Coke", Description: "Diet Coke", Price: 0},
		locations.OptionItem{Name: "Sprite", Description: "Sprite", Price: 0},
		locations.OptionItem{Name: "Diet Sprite", Description: "Diet Sprite", Price: 0},
		locations.OptionItem{Name: "Dr. Pepper", Description: "Dr. Pepper", Price: 0},
		locations.OptionItem{Name: "Diet Dr. Pepper", Description: "Diet Dr. Pepper", Price: 0}})
	_ = api.CreateOption("Extras", "Extras", 0, 0, theAmeliaRiverClubPinkLemonade.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Cream", Description: "Cream", Price: 0},
		locations.OptionItem{Name: "Sugar", Description: "Sugar", Price: 0},
		locations.OptionItem{Name: "Splenda", Description: "Splenda", Price: 0}})
	_ = api.CreateOption("Vodka", "Vodka", 1, 1, theAmeliaRiverClubScrewDriver.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Grey Goose", Description: "Grey Goose", Price: 4},
		locations.OptionItem{Name: "Titos", Description: "Titos", Price: 0},
		locations.OptionItem{Name: "Effen", Description: "Effen", Price: 0},
		locations.OptionItem{Name: "Absolut", Description: "Absolut", Price: 0},
		locations.OptionItem{Name: "Ketel One", Description: "Ketel One", Price: 0},
		locations.OptionItem{Name: "Stoli", Description: "Stoli", Price: 0}})
	_ = api.CreateOption("Vodka", "Vodka", 1, 1, theAmeliaRiverClubBloodyMary.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Grey Goose", Description: "Grey Goose", Price: 4},
		locations.OptionItem{Name: "Titos", Description: "Titos", Price: 0},
		locations.OptionItem{Name: "Effen", Description: "Effen", Price: 0},
		locations.OptionItem{Name: "Absolut", Description: "Absolut", Price: 0},
		locations.OptionItem{Name: "Ketel One", Description: "Ketel One", Price: 0},
		locations.OptionItem{Name: "Stoli", Description: "Stoli", Price: 0}})
	_ = api.CreateOption("Vodka", "Vodka", 1, 1, theAmeliaRiverClubJohnDaly.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Grey Goose", Description: "Grey Goose", Price: 4},
		locations.OptionItem{Name: "Titos", Description: "Titos", Price: 0},
		locations.OptionItem{Name: "Effen", Description: "Effen", Price: 0},
		locations.OptionItem{Name: "Absolut", Description: "Absolut", Price: 0},
		locations.OptionItem{Name: "Ketel One", Description: "Ketel One", Price: 0},
		locations.OptionItem{Name: "Stoli", Description: "Stoli", Price: 0}})
	_ = api.CreateOption("Tequila", "Tequila", 1, 1, theAmeliaRiverClubPaloma.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Jose Cuervo", Description: "Jose Cuervo", Price: 0},
		locations.OptionItem{Name: "Patron", Description: "Patron", Price: 4},
		locations.OptionItem{Name: "Suaza", Description: "Suaza", Price: 0},
		locations.OptionItem{Name: "1800", Description: "1800", Price: 0},
		locations.OptionItem{Name: "Hornitos", Description: "Hornitos", Price: 2},
		locations.OptionItem{Name: "Don Julio", Description: "Don Julio", Price: 4}})
	_ = api.CreateOption("Vodka", "Vodka", 1, 1, theAmeliaRiverClubMoscowMule.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Grey Goose", Description: "Grey Goose", Price: 4},
		locations.OptionItem{Name: "Titos", Description: "Titos", Price: 0},
		locations.OptionItem{Name: "Effen", Description: "Effen", Price: 0},
		locations.OptionItem{Name: "Absolut", Description: "Absolut", Price: 0},
		locations.OptionItem{Name: "Ketel One", Description: "Ketel One", Price: 0},
		locations.OptionItem{Name: "Stoli", Description: "Stoli", Price: 0}})
	_ = api.CreateOption("Rum", "Rum", 1, 1, theAmeliaRiverClubRumAndCoke.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Bacardi", Description: "Bacardi", Price: 0},
		locations.OptionItem{Name: "Captain Morgan", Description: "Captain Morgan", Price: 0},
		locations.OptionItem{Name: "Tanduay", Description: "Tanduay", Price: 0}})
	_ = api.CreateOption("Options", "Options", 0, 0, theAmeliaRiverClubRumAndCoke.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Diet Coke", Description: "Diet Coke", Price: 0}})
	_ = api.CreateOption("Whiskey", "Whiskey", 1, 1, theAmeliaRiverClubWhiskeyAndCoke.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Jack Daniels", Description: "Jack Daniels", Price: 0},
		locations.OptionItem{Name: "Seagrams", Description: "Seagrams", Price: 0},
		locations.OptionItem{Name: "Makers Mark", Description: "Makers Mark", Price: 0},
		locations.OptionItem{Name: "Crown Royal", Description: "Crown Royal", Price: 0},
		locations.OptionItem{Name: "Jameson", Description: "Jameson", Price: 0},
		locations.OptionItem{Name: "Jim Beam", Description: "Jim Beam", Price: 0}})
	_ = api.CreateOption("Options", "Options", 0, 0, theAmeliaRiverClubWhiskeyAndCoke.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Diet Coke", Description: "Diet Coke", Price: 0}})
	_ = api.CreateOption("Vodka", "Vodka", 1, 1, theAmeliaRiverClubAzelea.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Grey Goose", Description: "Grey Goose", Price: 4},
		locations.OptionItem{Name: "Titos", Description: "Titos", Price: 0},
		locations.OptionItem{Name: "Effen", Description: "Effen", Price: 0},
		locations.OptionItem{Name: "Absolut", Description: "Absolut", Price: 0},
		locations.OptionItem{Name: "Ketel One", Description: "Ketel One", Price: 0},
		locations.OptionItem{Name: "Stoli", Description: "Stoli", Price: 0}})
	_ = api.CreateOption("Vodka", "Vodka", 1, 1, theAmeliaRiverClubCranberryKamikaze.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Grey Goose", Description: "Grey Goose", Price: 4},
		locations.OptionItem{Name: "Titos", Description: "Titos", Price: 0},
		locations.OptionItem{Name: "Effen", Description: "Effen", Price: 0},
		locations.OptionItem{Name: "Absolut", Description: "Absolut", Price: 0},
		locations.OptionItem{Name: "Ketel One", Description: "Ketel One", Price: 0},
		locations.OptionItem{Name: "Stoli", Description: "Stoli", Price: 0}})
	_ = api.CreateOption("Vodka", "Vodka", 1, 1, theAmeliaRiverClubTransfusion.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Grey Goose", Description: "Grey Goose", Price: 4},
		locations.OptionItem{Name: "Titos", Description: "Titos", Price: 0},
		locations.OptionItem{Name: "Effen", Description: "Effen", Price: 0},
		locations.OptionItem{Name: "Absolut", Description: "Absolut", Price: 0},
		locations.OptionItem{Name: "Ketel One", Description: "Ketel One", Price: 0},
		locations.OptionItem{Name: "Stoli", Description: "Stoli", Price: 0}})
	_ = api.CreateOption("Vodka", "Vodka", 1, 1, theAmeliaRiverClubRedBullAndVodka.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Grey Goose", Description: "Grey Goose", Price: 4},
		locations.OptionItem{Name: "Titos", Description: "Titos", Price: 0},
		locations.OptionItem{Name: "Effen", Description: "Effen", Price: 0},
		locations.OptionItem{Name: "Absolut", Description: "Absolut", Price: 0},
		locations.OptionItem{Name: "Ketel One", Description: "Ketel One", Price: 0},
		locations.OptionItem{Name: "Stoli", Description: "Stoli", Price: 0}})
	// seed orders
	log.Println("Seeding Completed")
}
