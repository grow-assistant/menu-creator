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
	gatesFourGolfandCountryClub := api.CreateLocation("Gates Four Golf & Country Club", "Fayetteville, NC")
	log.Println("Seeding *Menus* data")
	// seed menus
	gatesFourGolfandCountryClubJPsBarandGrill := api.CreateMenu("JPs Bar and Grill", "JPs Bar and Grill", gatesFourGolfandCountryClub.ID)
	log.Println(gatesFourGolfandCountryClubJPsBarandGrill)
	// seed categories
	log.Println("Seeding *Categories* data")
	gatesFourGolfandCountryClubJPsBarandGrillDinnerEntrees := api.CreateCategory("Dinner Entrées", "Dinner Entrées", gatesFourGolfandCountryClubJPsBarandGrill.ID)
	gatesFourGolfandCountryClubJPsBarandGrillDinnerSpecialties := api.CreateCategory("Dinner Specialties", "Dinner Specialties", gatesFourGolfandCountryClubJPsBarandGrill.ID)
	gatesFourGolfandCountryClubJPsBarandGrillDinnerSides := api.CreateCategory("Dinner - Sides", "Dinner - Sides", gatesFourGolfandCountryClubJPsBarandGrill.ID)
	gatesFourGolfandCountryClubJPsBarandGrillDinnerDesserts := api.CreateCategory("Dinner - Desserts", "Dinner - Desserts", gatesFourGolfandCountryClubJPsBarandGrill.ID)
	gatesFourGolfandCountryClubJPsBarandGrillLunchAppetizers := api.CreateCategory("Lunch - Appetizers", "Lunch - Appetizers", gatesFourGolfandCountryClubJPsBarandGrill.ID)
	gatesFourGolfandCountryClubJPsBarandGrillLunchSoupsAndSalads := api.CreateCategory("Lunch - Soups & Salads", "Lunch - Soups & Salads", gatesFourGolfandCountryClubJPsBarandGrill.ID)
	gatesFourGolfandCountryClubJPsBarandGrillLunchSpecialties := api.CreateCategory("Lunch - Specialties", "Lunch - Specialties", gatesFourGolfandCountryClubJPsBarandGrill.ID)
	gatesFourGolfandCountryClubJPsBarandGrillLunchSandwiches := api.CreateCategory("Lunch - Sandwiches", "Lunch - Sandwiches", gatesFourGolfandCountryClubJPsBarandGrill.ID)
	gatesFourGolfandCountryClubJPsBarandGrillLunchBurgers := api.CreateCategory("Lunch - Burgers", "Lunch - Burgers", gatesFourGolfandCountryClubJPsBarandGrill.ID)
	gatesFourGolfandCountryClubJPsBarandGrillLunchSides := api.CreateCategory("Lunch - Sides", "Lunch - Sides", gatesFourGolfandCountryClubJPsBarandGrill.ID)
	log.Println(gatesFourGolfandCountryClubJPsBarandGrillDinnerEntrees)
	log.Println(gatesFourGolfandCountryClubJPsBarandGrillDinnerSpecialties)
	log.Println(gatesFourGolfandCountryClubJPsBarandGrillDinnerSides)
	log.Println(gatesFourGolfandCountryClubJPsBarandGrillDinnerDesserts)
	log.Println(gatesFourGolfandCountryClubJPsBarandGrillLunchAppetizers)
	log.Println(gatesFourGolfandCountryClubJPsBarandGrillLunchSoupsAndSalads)
	log.Println(gatesFourGolfandCountryClubJPsBarandGrillLunchSpecialties)
	log.Println(gatesFourGolfandCountryClubJPsBarandGrillLunchSandwiches)
	log.Println(gatesFourGolfandCountryClubJPsBarandGrillLunchBurgers)
	log.Println(gatesFourGolfandCountryClubJPsBarandGrillLunchSides)
	// seed item
	log.Println("Seeding *Items* data")
	gatesFourGolfandCountryClubChickenCaprese := api.CreateItem("Chicken Caprese", "A pan-seared chicken breast covered with fresh basil, tomatoes, and mozzarella cheese.", 18, gatesFourGolfandCountryClubJPsBarandGrillDinnerEntrees.ID)
	gatesFourGolfandCountryClubCountryFriedChicken := api.CreateItem("Country Fried Chicken", "A breaded chicken breast, fried golden and smothered with a white country-style gravy.", 18, gatesFourGolfandCountryClubJPsBarandGrillDinnerEntrees.ID)
	gatesFourGolfandCountryClubStuffedChicken := api.CreateItem("Stuffed Chicken", "A breaded chicken breast coated with breadcrumb, stuffed with an herb blended cream cheese and spinach.", 18, gatesFourGolfandCountryClubJPsBarandGrillDinnerEntrees.ID)
	gatesFourGolfandCountryClubSearedSalmon := api.CreateItem("Seared Salmon", "An 8oz portion of fresh Atlantic salmon, seared with blackened seasoning, or topped with a lemon caper beurre blanc.", 21, gatesFourGolfandCountryClubJPsBarandGrillDinnerEntrees.ID)
	gatesFourGolfandCountryClubHamburgerSteak := api.CreateItem("Hamburger Steak", "An 8oz angus beef steak seasoned, pan fried, and smothered in a homemade country brown gravy.", 16, gatesFourGolfandCountryClubJPsBarandGrillDinnerEntrees.ID)
	gatesFourGolfandCountryClubTheDriver := api.CreateItem("The Driver", "A fresh 14oz USDA Choice cut boneless Ribeye, grilled your way with our parmesan peppercorn butter.", 38, gatesFourGolfandCountryClubJPsBarandGrillDinnerEntrees.ID)
	gatesFourGolfandCountryClubTheWood := api.CreateItem("The Wood", "An 8oz filet, grilled your way with our parmesan peppercorn butter.", 28, gatesFourGolfandCountryClubJPsBarandGrillDinnerEntrees.ID)
	gatesFourGolfandCountryClubTheWedge := api.CreateItem("The Wedge", "A 10oz sirloin, grilled your way with our parmesan peppercorn butter.", 26, gatesFourGolfandCountryClubJPsBarandGrillDinnerEntrees.ID)
	gatesFourGolfandCountryClubCountryFriedPorkChop := api.CreateItem("Country Fried Pork Chop", "A 6oz bone-in pork chop, breaded, fried, and smothered in country-style white gravy.", 18, gatesFourGolfandCountryClubJPsBarandGrillDinnerEntrees.ID)
	gatesFourGolfandCountryClubBakedSpaghetti := api.CreateItem("Baked Spaghetti", "Pasta with homemade marinara, peppers and onions, and mozzarella cheese. Served with breadsticks.", 13, gatesFourGolfandCountryClubJPsBarandGrillDinnerSpecialties.ID)
	gatesFourGolfandCountryClubPastaAlfredo := api.CreateItem("Pasta Alfredo", "Your choice of pasta tossed with our homemade alfredo sauce. Served with breadsticks.", 11, gatesFourGolfandCountryClubJPsBarandGrillDinnerSpecialties.ID)
	gatesFourGolfandCountryClubPastaPrimavera := api.CreateItem("Pasta Primavera", "Your choice of pasta tossed with fresh steamed broccoli, breadsticks, zucchini, and bell peppers. Served with breadsticks.", 15, gatesFourGolfandCountryClubJPsBarandGrillDinnerSpecialties.ID)
	gatesFourGolfandCountryClubShrimpPlatter := api.CreateItem("Shrimp Platter", "Eight jumbo Atlantic caught shrimp fried, grilled, or blackened, served with French fries.", 18, gatesFourGolfandCountryClubJPsBarandGrillDinnerSpecialties.ID)
	_ = api.CreateItem("Shrimp and Grits", "A bowl of southern-style cheesy grits topped with eight blackened jumbo shrimp.", 18, gatesFourGolfandCountryClubJPsBarandGrillDinnerSpecialties.ID)
	_ = api.CreateItem("Rustic Mash Potato", "Rustic Mash Potato", 2, gatesFourGolfandCountryClubJPsBarandGrillDinnerSides.ID)
	_ = api.CreateItem("Jasmine Rice", "Jasmine Rice", 2, gatesFourGolfandCountryClubJPsBarandGrillDinnerSides.ID)
	_ = api.CreateItem("Baked Potato", "Baked Potato", 2, gatesFourGolfandCountryClubJPsBarandGrillDinnerSides.ID)
	_ = api.CreateItem("Sweet Potato", "Sweet Potato", 2, gatesFourGolfandCountryClubJPsBarandGrillDinnerSides.ID)
	_ = api.CreateItem("French Fries", "French Fries", 2, gatesFourGolfandCountryClubJPsBarandGrillDinnerSides.ID)
	_ = api.CreateItem("Sweet Potato Fries", "Sweet Potato Fries", 3, gatesFourGolfandCountryClubJPsBarandGrillDinnerSides.ID)
	_ = api.CreateItem("Vegetable of the Day", "Vegetable of the Day", 3, gatesFourGolfandCountryClubJPsBarandGrillDinnerSides.ID)
	_ = api.CreateItem("Asparagus", "Asparagus", 4, gatesFourGolfandCountryClubJPsBarandGrillDinnerSides.ID)
	_ = api.CreateItem("Add Bacon, Cheddar to Potato", "Add Bacon, Cheddar to Potato", 2, gatesFourGolfandCountryClubJPsBarandGrillDinnerSides.ID)
	_ = api.CreateItem("Ultimate Chocolate Cake", "Ultimate Chocolate Cake", 5, gatesFourGolfandCountryClubJPsBarandGrillDinnerDesserts.ID)
	_ = api.CreateItem("Key Lime Pie", "Key Lime Pie", 5, gatesFourGolfandCountryClubJPsBarandGrillDinnerDesserts.ID)
	_ = api.CreateItem("Carrot Cake", "Carrot Cake", 5, gatesFourGolfandCountryClubJPsBarandGrillDinnerDesserts.ID)
	_ = api.CreateItem("Brownie", "Brownie", 3, gatesFourGolfandCountryClubJPsBarandGrillDinnerDesserts.ID)
	_ = api.CreateItem("Vanilla Ice Cream Scoop", "Vanilla Ice Cream Scoop", 2, gatesFourGolfandCountryClubJPsBarandGrillDinnerDesserts.ID)
	gatesFourGolfandCountryClubMilkshakes := api.CreateItem("Milkshakes", "Milkshakes", 5, gatesFourGolfandCountryClubJPsBarandGrillDinnerDesserts.ID)
	_ = api.CreateItem("Calamari", "A mix of tentacles and rings, lightly coated and fried, served with a side of cocktail.", 12, gatesFourGolfandCountryClubJPsBarandGrillLunchAppetizers.ID)
	gatesFourGolfandCountryClubChickenWings := api.CreateItem("Chicken Wings", "Fried jumbo chicken wings, your choice of six or twelve, tossed in one of our signature sauces: Buffalo, BBQ, Blazing BBQ, Incinerator, Garlic Parm.", 7, gatesFourGolfandCountryClubJPsBarandGrillLunchAppetizers.ID)
	gatesFourGolfandCountryClubClubNachos := api.CreateItem("Club Nachos", "Your choice of chicken or beef, atop warm corn tortilla chips with jalapenos, onions, and cheddar cheese.", 12, gatesFourGolfandCountryClubJPsBarandGrillLunchAppetizers.ID)
	_ = api.CreateItem("Loaded Skins", "Six halves, loaded down with cheddar cheese, bacon, and green onion with a side of sour cream.", 10, gatesFourGolfandCountryClubJPsBarandGrillLunchAppetizers.ID)
	_ = api.CreateItem("Mozzarella Sticks", "Six of our breaded mozzarella sticks, fried golden and served with marinara.", 8, gatesFourGolfandCountryClubJPsBarandGrillLunchAppetizers.ID)
	_ = api.CreateItem("Par Three Platter", "A half portion of quesadilla, three mozzarella sticks, and six of our stuffed mushrooms, served with marinara, salsa, and Cajun aioli.", 13, gatesFourGolfandCountryClubJPsBarandGrillLunchAppetizers.ID)
	_ = api.CreateItem("BBQ Pork Quesadilla", "A full 12\" quesadilla stuffed with smoked pork, pepper jack, onions, and jalapenos.", 12, gatesFourGolfandCountryClubJPsBarandGrillLunchAppetizers.ID)
	_ = api.CreateItem("Loaded Quesadilla", "A full 12\" quesadilla loaded with peppers, onions, and chicken. Served with salsa and sour cream.", 11, gatesFourGolfandCountryClubJPsBarandGrillLunchAppetizers.ID)
	_ = api.CreateItem("Stuffed Mushrooms", "A plate of our baked button mushrooms stuffed with our blend of seasonings and cream cheese.", 8, gatesFourGolfandCountryClubJPsBarandGrillLunchAppetizers.ID)
	_ = api.CreateItem("Soup of the Day (Cup)", "Chef’s daily soup selection.", 4, gatesFourGolfandCountryClubJPsBarandGrillLunchSoupsAndSalads.ID)
	_ = api.CreateItem("Soup of the Day (Bowl)", "Chef’s daily soup selection.", 6, gatesFourGolfandCountryClubJPsBarandGrillLunchSoupsAndSalads.ID)
	_ = api.CreateItem("Side House or Caesar Salad", "Small portion of house or Caesar salad.", 5, gatesFourGolfandCountryClubJPsBarandGrillLunchSoupsAndSalads.ID)
	gatesFourGolfandCountryClubHouseSalad := api.CreateItem("House Salad", "Fresh mixed greens with tomato, cucumber, mushrooms, onions, and cheddar.", 10, gatesFourGolfandCountryClubJPsBarandGrillLunchSoupsAndSalads.ID)
	gatesFourGolfandCountryClubCaesarSalad := api.CreateItem("Caesar Salad", "Chopped romaine tossed with parmesan, croutons, and Caesar dressing.", 9, gatesFourGolfandCountryClubJPsBarandGrillLunchSoupsAndSalads.ID)
	gatesFourGolfandCountryClubClubSalad := api.CreateItem("Club Salad", "Fresh mixed greens with tomato, cucumber, mushrooms, onions, cheddar, smoked ham and turkey, and bacon.", 13, gatesFourGolfandCountryClubJPsBarandGrillLunchSoupsAndSalads.ID)
	gatesFourGolfandCountryClubGatesFourSpinachSalad := api.CreateItem("Gates Four Spinach Salad", "Baby leaf spinach tossed with candied nuts, craisins, bacon, onion, blue cheese crumbles, and balsamic dressing.", 14, gatesFourGolfandCountryClubJPsBarandGrillLunchSoupsAndSalads.ID)
	gatesFourGolfandCountryClubAhiTunaSteak := api.CreateItem("Ahi Tuna Steak", "A yellowfin tuna steak crusted with sesame and lightly seared and sliced. Served with your choice of side.", 14, gatesFourGolfandCountryClubJPsBarandGrillLunchSpecialties.ID)
	_ = api.CreateItem("Bang Bang Tempura", "Eight battered jumbo shrimp tossed in our sweet and sour sauce. Served with your choice of side.", 15, gatesFourGolfandCountryClubJPsBarandGrillLunchSpecialties.ID)
	_ = api.CreateItem("Fish and Chips", "Two battered cod filets fried golden and served with an order of French fries and tartar sauce.", 16, gatesFourGolfandCountryClubJPsBarandGrillLunchSpecialties.ID)
	gatesFourGolfandCountryClubChickenFajitas := api.CreateItem("Chicken Fajitas", "Two 6\" flour tortillas with grilled fajita chicken, grilled peppers and onions, and cotija cheese. Served with your choice of side.", 14, gatesFourGolfandCountryClubJPsBarandGrillLunchSpecialties.ID)
	_ = api.CreateItem("Bacon Mac & Cheese", "Penne pasta tossed in our homemade bacon cheddar sauce.", 9, gatesFourGolfandCountryClubJPsBarandGrillLunchSpecialties.ID)
	gatesFourGolfandCountryClubPortobelloBurger := api.CreateItem("Portobello Burger", "A portobello mushroom cap seasoned and grilled, served on a seeded gluten-free bun with lettuce and tomato. Served with your choice of side.", 12, gatesFourGolfandCountryClubJPsBarandGrillLunchSpecialties.ID)
	gatesFourGolfandCountryClubIrongateDeli := api.CreateItem("Irongate Deli", "Your choice of sliced deli meat, bread or wrap, and cheese with lettuce, tomato, and onion. Served with your choice of side.", 10, gatesFourGolfandCountryClubJPsBarandGrillLunchSandwiches.ID)
	gatesFourGolfandCountryClubPhillyCheesesteak := api.CreateItem("Philly Cheesesteak", "Sliced roasted steak with peppers, onions, and mushrooms, topped with provolone cheese in a footlong hoagie roll. Served with your choice of side.", 16, gatesFourGolfandCountryClubJPsBarandGrillLunchSandwiches.ID)
	gatesFourGolfandCountryClubTheItalian := api.CreateItem("The Italian", "Ham, capicola, and pepperoni with lettuce, oil and vinegar, and banana peppers in a footlong hoagie roll. Served with your choice of side.", 15, gatesFourGolfandCountryClubJPsBarandGrillLunchSandwiches.ID)
	gatesFourGolfandCountryClubBlackenedSalmonSandwich := api.CreateItem("Blackened Salmon Sandwich", "A filet of North Atlantic caught salmon seared and topped with lettuce and tomato on a potato bun. Served with your choice of side.", 18, gatesFourGolfandCountryClubJPsBarandGrillLunchSandwiches.ID)
	gatesFourGolfandCountryClubJPsClubSandwich := api.CreateItem("JP’s Club Sandwich", "Black forest ham, roasted turkey, and smoked bacon with Swiss and cheddar cheeses, crisp leaf lettuce, tomato, and mayo on three slices of your choice of bread. Served with your choice of side.", 14, gatesFourGolfandCountryClubJPsBarandGrillLunchSandwiches.ID)
	gatesFourGolfandCountryClubBuffaloChickenSandwich := api.CreateItem("Buffalo Chicken Sandwich", "A fried 8oz chicken breast, tossed in buffalo sauce, topped with pepper jack cheese, bacon, leaf lettuce, and tomato on a potato bun. Served with your choice of side.", 17, gatesFourGolfandCountryClubJPsBarandGrillLunchSandwiches.ID)
	gatesFourGolfandCountryClubJPsAce := api.CreateItem("JP’s Ace", "Angus hotdog with chili, coleslaw, mustard, diced onions. Side Optional", 4, gatesFourGolfandCountryClubJPsBarandGrillLunchSandwiches.ID)
	gatesFourGolfandCountryClubGatesFourBurger := api.CreateItem("Gates Four Burger", "A half-pound hamburger grilled with American cheese, lettuce, onion, and tomato. Served with your choice of side.", 14, gatesFourGolfandCountryClubJPsBarandGrillLunchBurgers.ID)
	gatesFourGolfandCountryClubBigSkyBurger := api.CreateItem("Big Sky Burger", "Our half-pound burger grilled with barbecue sauce, applewood smoked bacon, cheddar cheese, and crispy fried onions. Served with your choice of side.", 15, gatesFourGolfandCountryClubJPsBarandGrillLunchBurgers.ID)
	gatesFourGolfandCountryClubCarolinaBurger := api.CreateItem("Carolina Burger", "Our half-pound burger grilled with chili, American cheese, coleslaw, mustard, and onions. Served with your choice of sides.", 15, gatesFourGolfandCountryClubJPsBarandGrillLunchBurgers.ID)
	_ = api.CreateItem("French Fries", "Crispy, golden, and perfectly salted.", 2, gatesFourGolfandCountryClubJPsBarandGrillLunchSides.ID)
	_ = api.CreateItem("Sweet Potato Fries", "Sweet, crispy, and addicting.", 2, gatesFourGolfandCountryClubJPsBarandGrillLunchSides.ID)
	_ = api.CreateItem("Fruit", "Seasonal, fresh, and refreshing.", 2, gatesFourGolfandCountryClubJPsBarandGrillLunchSides.ID)
	_ = api.CreateItem("Potato Salad", "Creamy, classic, and loaded with flavor.", 2, gatesFourGolfandCountryClubJPsBarandGrillLunchSides.ID)
	_ = api.CreateItem("Pasta Salad", "Chilled, zesty, and packed with veggies.", 2, gatesFourGolfandCountryClubJPsBarandGrillLunchSides.ID)
	_ = api.CreateItem("Onion Rings", "Thick-cut, crispy, and golden brown.", 3, gatesFourGolfandCountryClubJPsBarandGrillLunchSides.ID)
	// seed item options
	log.Println("Seeding *Options* data")
	_ = api.CreateOption("Extras", "Extras", 0, 0, gatesFourGolfandCountryClubChickenCaprese.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Add Sauteed Mushrooms", Description: "Add Sauteed Mushrooms", Price: 2},
		locations.OptionItem{Name: "Add Three Fried Shrimp", Description: "Add Three Fried Shrimp", Price: 4},
		locations.OptionItem{Name: "Add Six Shrimp", Description: "Add Six Shrimp", Price: 6},
		locations.OptionItem{Name: "Oscar Style", Description: "Oscar Style", Price: 8},
		locations.OptionItem{Name: "Add Chicken", Description: "Add Chicken", Price: 4},
		locations.OptionItem{Name: "Add Salmon", Description: "Add Salmon", Price: 10}})
	_ = api.CreateOption("Choice of 2 Sides", "Choice of 2 Sides", 0, 0, gatesFourGolfandCountryClubChickenCaprese.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Rustic Mash Potato", Description: "Rustic Mash Potato", Price: 0},
		locations.OptionItem{Name: "Jasmine Rice", Description: "Jasmine Rice", Price: 0},
		locations.OptionItem{Name: "Baked Potato", Description: "Baked Potato", Price: 0},
		locations.OptionItem{Name: "Sweet Potato", Description: "Sweet Potato", Price: 0},
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Vegetable of the Day", Description: "Vegetable of the Day", Price: 0},
		locations.OptionItem{Name: "Asparagus", Description: "Asparagus", Price: 0},
		locations.OptionItem{Name: "Add Bacon, Cheddar to Potato", Description: "Add Bacon, Cheddar to Potato", Price: 0}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, gatesFourGolfandCountryClubChickenCaprese.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Remove Tomatoes", Description: "Remove Tomatoes", Price: 0}})
	_ = api.CreateOption("Extras", "Extras", 0, 0, gatesFourGolfandCountryClubCountryFriedChicken.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Add Sauteed Mushrooms", Description: "Add Sauteed Mushrooms", Price: 2},
		locations.OptionItem{Name: "Add Three Fried Shrimp", Description: "Add Three Fried Shrimp", Price: 4},
		locations.OptionItem{Name: "Add Six Shrimp", Description: "Add Six Shrimp", Price: 6},
		locations.OptionItem{Name: "Oscar Style", Description: "Oscar Style", Price: 8},
		locations.OptionItem{Name: "Add Chicken", Description: "Add Chicken", Price: 4},
		locations.OptionItem{Name: "Add Salmon", Description: "Add Salmon", Price: 10}})
	_ = api.CreateOption("Choice of 2 Sides", "Choice of 2 Sides", 0, 0, gatesFourGolfandCountryClubCountryFriedChicken.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Rustic Mash Potato", Description: "Rustic Mash Potato", Price: 0},
		locations.OptionItem{Name: "Jasmine Rice", Description: "Jasmine Rice", Price: 0},
		locations.OptionItem{Name: "Baked Potato", Description: "Baked Potato", Price: 0},
		locations.OptionItem{Name: "Sweet Potato", Description: "Sweet Potato", Price: 0},
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Vegetable of the Day", Description: "Vegetable of the Day", Price: 0},
		locations.OptionItem{Name: "Asparagus", Description: "Asparagus", Price: 0},
		locations.OptionItem{Name: "Add Bacon, Cheddar to Potato", Description: "Add Bacon, Cheddar to Potato", Price: 0}})
	_ = api.CreateOption("Extras", "Extras", 0, 0, gatesFourGolfandCountryClubStuffedChicken.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Add Sauteed Mushrooms", Description: "Add Sauteed Mushrooms", Price: 2},
		locations.OptionItem{Name: "Add Three Fried Shrimp", Description: "Add Three Fried Shrimp", Price: 4},
		locations.OptionItem{Name: "Add Six Shrimp", Description: "Add Six Shrimp", Price: 6},
		locations.OptionItem{Name: "Oscar Style", Description: "Oscar Style", Price: 8},
		locations.OptionItem{Name: "Add Chicken", Description: "Add Chicken", Price: 4},
		locations.OptionItem{Name: "Add Salmon", Description: "Add Salmon", Price: 10}})
	_ = api.CreateOption("Choice of 2 Sides", "Choice of 2 Sides", 0, 0, gatesFourGolfandCountryClubStuffedChicken.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Rustic Mash Potato", Description: "Rustic Mash Potato", Price: 0},
		locations.OptionItem{Name: "Jasmine Rice", Description: "Jasmine Rice", Price: 0},
		locations.OptionItem{Name: "Baked Potato", Description: "Baked Potato", Price: 0},
		locations.OptionItem{Name: "Sweet Potato", Description: "Sweet Potato", Price: 0},
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Vegetable of the Day", Description: "Vegetable of the Day", Price: 0},
		locations.OptionItem{Name: "Asparagus", Description: "Asparagus", Price: 0},
		locations.OptionItem{Name: "Add Bacon, Cheddar to Potato", Description: "Add Bacon, Cheddar to Potato", Price: 0}})
	_ = api.CreateOption("Extras", "Extras", 0, 0, gatesFourGolfandCountryClubSearedSalmon.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Add Sauteed Mushrooms", Description: "Add Sauteed Mushrooms", Price: 2},
		locations.OptionItem{Name: "Add Three Fried Shrimp", Description: "Add Three Fried Shrimp", Price: 4},
		locations.OptionItem{Name: "Add Six Shrimp", Description: "Add Six Shrimp", Price: 6},
		locations.OptionItem{Name: "Oscar Style", Description: "Oscar Style", Price: 8},
		locations.OptionItem{Name: "Add Chicken", Description: "Add Chicken", Price: 4},
		locations.OptionItem{Name: "Add Salmon", Description: "Add Salmon", Price: 10}})
	_ = api.CreateOption("Choice of 2 Sides", "Choice of 2 Sides", 0, 0, gatesFourGolfandCountryClubSearedSalmon.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Rustic Mash Potato", Description: "Rustic Mash Potato", Price: 0},
		locations.OptionItem{Name: "Jasmine Rice", Description: "Jasmine Rice", Price: 0},
		locations.OptionItem{Name: "Baked Potato", Description: "Baked Potato", Price: 0},
		locations.OptionItem{Name: "Sweet Potato", Description: "Sweet Potato", Price: 0},
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Vegetable of the Day", Description: "Vegetable of the Day", Price: 0},
		locations.OptionItem{Name: "Asparagus", Description: "Asparagus", Price: 0},
		locations.OptionItem{Name: "Add Bacon, Cheddar to Potato", Description: "Add Bacon, Cheddar to Potato", Price: 0}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, gatesFourGolfandCountryClubSearedSalmon.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Remove Lemon Caper Sauce", Description: "Remove Lemon Caper Sauce", Price: 0}})
	_ = api.CreateOption("Extras", "Extras", 0, 0, gatesFourGolfandCountryClubHamburgerSteak.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Add Sauteed Mushrooms", Description: "Add Sauteed Mushrooms", Price: 2},
		locations.OptionItem{Name: "Add Three Fried Shrimp", Description: "Add Three Fried Shrimp", Price: 4},
		locations.OptionItem{Name: "Add Six Shrimp", Description: "Add Six Shrimp", Price: 6},
		locations.OptionItem{Name: "Oscar Style", Description: "Oscar Style", Price: 8},
		locations.OptionItem{Name: "Add Chicken", Description: "Add Chicken", Price: 4},
		locations.OptionItem{Name: "Add Salmon", Description: "Add Salmon", Price: 10}})
	_ = api.CreateOption("Choice of 2 Sides", "Choice of 2 Sides", 0, 0, gatesFourGolfandCountryClubHamburgerSteak.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Rustic Mash Potato", Description: "Rustic Mash Potato", Price: 0},
		locations.OptionItem{Name: "Jasmine Rice", Description: "Jasmine Rice", Price: 0},
		locations.OptionItem{Name: "Baked Potato", Description: "Baked Potato", Price: 0},
		locations.OptionItem{Name: "Sweet Potato", Description: "Sweet Potato", Price: 0},
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Vegetable of the Day", Description: "Vegetable of the Day", Price: 0},
		locations.OptionItem{Name: "Asparagus", Description: "Asparagus", Price: 0},
		locations.OptionItem{Name: "Add Bacon, Cheddar to Potato", Description: "Add Bacon, Cheddar to Potato", Price: 0}})
	_ = api.CreateOption("Meat Temperature", "Meat Temperature", 1, 1, gatesFourGolfandCountryClubTheDriver.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Well Done", Description: "Well Done", Price: 0},
		locations.OptionItem{Name: "Medium-Well", Description: "Medium-Well", Price: 0},
		locations.OptionItem{Name: "Medium", Description: "Medium", Price: 0},
		locations.OptionItem{Name: "Medium-Rare", Description: "Medium-Rare", Price: 0},
		locations.OptionItem{Name: "Rare", Description: "Rare", Price: 0}})
	_ = api.CreateOption("Extras", "Extras", 0, 0, gatesFourGolfandCountryClubTheDriver.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Add Sauteed Mushrooms", Description: "Add Sauteed Mushrooms", Price: 2},
		locations.OptionItem{Name: "Add Three Fried Shrimp", Description: "Add Three Fried Shrimp", Price: 4},
		locations.OptionItem{Name: "Add Six Shrimp", Description: "Add Six Shrimp", Price: 6},
		locations.OptionItem{Name: "Oscar Style", Description: "Oscar Style", Price: 8},
		locations.OptionItem{Name: "Add Chicken", Description: "Add Chicken", Price: 4},
		locations.OptionItem{Name: "Add Salmon", Description: "Add Salmon", Price: 10}})
	_ = api.CreateOption("Choice of 2 Sides", "Choice of 2 Sides", 0, 0, gatesFourGolfandCountryClubTheDriver.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Rustic Mash Potato", Description: "Rustic Mash Potato", Price: 0},
		locations.OptionItem{Name: "Jasmine Rice", Description: "Jasmine Rice", Price: 0},
		locations.OptionItem{Name: "Baked Potato", Description: "Baked Potato", Price: 0},
		locations.OptionItem{Name: "Sweet Potato", Description: "Sweet Potato", Price: 0},
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Vegetable of the Day", Description: "Vegetable of the Day", Price: 0},
		locations.OptionItem{Name: "Asparagus", Description: "Asparagus", Price: 0},
		locations.OptionItem{Name: "Add Bacon, Cheddar to Potato", Description: "Add Bacon, Cheddar to Potato", Price: 0}})
	_ = api.CreateOption("Meat Temperature", "Meat Temperature", 1, 1, gatesFourGolfandCountryClubTheWood.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Well Done", Description: "Well Done", Price: 0},
		locations.OptionItem{Name: "Medium-Well", Description: "Medium-Well", Price: 0},
		locations.OptionItem{Name: "Medium", Description: "Medium", Price: 0},
		locations.OptionItem{Name: "Medium-Rare", Description: "Medium-Rare", Price: 0},
		locations.OptionItem{Name: "Rare", Description: "Rare", Price: 0}})
	_ = api.CreateOption("Extras", "Extras", 0, 0, gatesFourGolfandCountryClubTheWood.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Add Sauteed Mushrooms", Description: "Add Sauteed Mushrooms", Price: 2},
		locations.OptionItem{Name: "Add Three Fried Shrimp", Description: "Add Three Fried Shrimp", Price: 4},
		locations.OptionItem{Name: "Add Six Shrimp", Description: "Add Six Shrimp", Price: 6},
		locations.OptionItem{Name: "Oscar Style", Description: "Oscar Style", Price: 8},
		locations.OptionItem{Name: "Add Chicken", Description: "Add Chicken", Price: 4},
		locations.OptionItem{Name: "Add Salmon", Description: "Add Salmon", Price: 10}})
	_ = api.CreateOption("Choice of 2 Sides", "Choice of 2 Sides", 0, 0, gatesFourGolfandCountryClubTheWood.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Rustic Mash Potato", Description: "Rustic Mash Potato", Price: 0},
		locations.OptionItem{Name: "Jasmine Rice", Description: "Jasmine Rice", Price: 0},
		locations.OptionItem{Name: "Baked Potato", Description: "Baked Potato", Price: 0},
		locations.OptionItem{Name: "Sweet Potato", Description: "Sweet Potato", Price: 0},
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Vegetable of the Day", Description: "Vegetable of the Day", Price: 0},
		locations.OptionItem{Name: "Asparagus", Description: "Asparagus", Price: 0},
		locations.OptionItem{Name: "Add Bacon, Cheddar to Potato", Description: "Add Bacon, Cheddar to Potato", Price: 0}})
	_ = api.CreateOption("Meat Temperature", "Meat Temperature", 1, 1, gatesFourGolfandCountryClubTheWedge.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Well Done", Description: "Well Done", Price: 0},
		locations.OptionItem{Name: "Medium-Well", Description: "Medium-Well", Price: 0},
		locations.OptionItem{Name: "Medium", Description: "Medium", Price: 0},
		locations.OptionItem{Name: "Medium-Rare", Description: "Medium-Rare", Price: 0},
		locations.OptionItem{Name: "Rare", Description: "Rare", Price: 0}})
	_ = api.CreateOption("Extras", "Extras", 0, 0, gatesFourGolfandCountryClubTheWedge.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Add Sauteed Mushrooms", Description: "Add Sauteed Mushrooms", Price: 2},
		locations.OptionItem{Name: "Add Three Fried Shrimp", Description: "Add Three Fried Shrimp", Price: 4},
		locations.OptionItem{Name: "Add Six Shrimp", Description: "Add Six Shrimp", Price: 6},
		locations.OptionItem{Name: "Oscar Style", Description: "Oscar Style", Price: 8},
		locations.OptionItem{Name: "Add Chicken", Description: "Add Chicken", Price: 4},
		locations.OptionItem{Name: "Add Salmon", Description: "Add Salmon", Price: 10}})
	_ = api.CreateOption("Choice of 2 Sides", "Choice of 2 Sides", 0, 0, gatesFourGolfandCountryClubTheWedge.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Rustic Mash Potato", Description: "Rustic Mash Potato", Price: 0},
		locations.OptionItem{Name: "Jasmine Rice", Description: "Jasmine Rice", Price: 0},
		locations.OptionItem{Name: "Baked Potato", Description: "Baked Potato", Price: 0},
		locations.OptionItem{Name: "Sweet Potato", Description: "Sweet Potato", Price: 0},
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Vegetable of the Day", Description: "Vegetable of the Day", Price: 0},
		locations.OptionItem{Name: "Asparagus", Description: "Asparagus", Price: 0},
		locations.OptionItem{Name: "Add Bacon, Cheddar to Potato", Description: "Add Bacon, Cheddar to Potato", Price: 0}})
	_ = api.CreateOption("Extras", "Extras", 0, 0, gatesFourGolfandCountryClubCountryFriedPorkChop.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Add Sauteed Mushrooms", Description: "Add Sauteed Mushrooms", Price: 2},
		locations.OptionItem{Name: "Add Three Fried Shrimp", Description: "Add Three Fried Shrimp", Price: 4},
		locations.OptionItem{Name: "Add Six Shrimp", Description: "Add Six Shrimp", Price: 6},
		locations.OptionItem{Name: "Oscar Style", Description: "Oscar Style", Price: 8},
		locations.OptionItem{Name: "Add Chicken", Description: "Add Chicken", Price: 4},
		locations.OptionItem{Name: "Add Salmon", Description: "Add Salmon", Price: 10}})
	_ = api.CreateOption("Choice of 2 Sides", "Choice of 2 Sides", 0, 0, gatesFourGolfandCountryClubCountryFriedPorkChop.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Rustic Mash Potato", Description: "Rustic Mash Potato", Price: 0},
		locations.OptionItem{Name: "Jasmine Rice", Description: "Jasmine Rice", Price: 0},
		locations.OptionItem{Name: "Baked Potato", Description: "Baked Potato", Price: 0},
		locations.OptionItem{Name: "Sweet Potato", Description: "Sweet Potato", Price: 0},
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Vegetable of the Day", Description: "Vegetable of the Day", Price: 0},
		locations.OptionItem{Name: "Asparagus", Description: "Asparagus", Price: 0},
		locations.OptionItem{Name: "Add Bacon, Cheddar to Potato", Description: "Add Bacon, Cheddar to Potato", Price: 0}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, gatesFourGolfandCountryClubBakedSpaghetti.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Remove Onions", Description: "Remove Onions", Price: 0},
		locations.OptionItem{Name: "Remove Peppers", Description: "Remove Peppers", Price: 0}})
	_ = api.CreateOption("Options", "Options", 1, 1, gatesFourGolfandCountryClubPastaAlfredo.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Fettuccine", Description: "Fettuccine", Price: 0},
		locations.OptionItem{Name: "Linguine", Description: "Linguine", Price: 0},
		locations.OptionItem{Name: "Penne", Description: "Penne", Price: 0},
		locations.OptionItem{Name: "Spaghetti", Description: "Spaghetti", Price: 0}})
	_ = api.CreateOption("Options", "Options", 1, 1, gatesFourGolfandCountryClubPastaPrimavera.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Fettuccine", Description: "Fettuccine", Price: 0},
		locations.OptionItem{Name: "Linguine", Description: "Linguine", Price: 0},
		locations.OptionItem{Name: "Penne", Description: "Penne", Price: 0},
		locations.OptionItem{Name: "Spaghetti", Description: "Spaghetti", Price: 0}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, gatesFourGolfandCountryClubPastaPrimavera.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Remove Broccoli", Description: "Remove Broccoli", Price: 0},
		locations.OptionItem{Name: "Remove Zucchini", Description: "Remove Zucchini", Price: 0},
		locations.OptionItem{Name: "Remove Bell Peppers", Description: "Remove Bell Peppers", Price: 0}})
	_ = api.CreateOption("Options", "Options", 1, 1, gatesFourGolfandCountryClubShrimpPlatter.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Grilled", Description: "Grilled", Price: 0},
		locations.OptionItem{Name: "Blackened", Description: "Blackened", Price: 0},
		locations.OptionItem{Name: "Fried", Description: "Fried", Price: 0}})
	_ = api.CreateOption("Options", "Options", 1, 1, gatesFourGolfandCountryClubMilkshakes.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Vanilla", Description: "Vanilla", Price: 0},
		locations.OptionItem{Name: "Chocolate", Description: "Chocolate", Price: 0}})
	_ = api.CreateOption("Wings Sauce", "Wings Sauce", 1, 1, gatesFourGolfandCountryClubChickenWings.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Buffalo", Description: "Buffalo", Price: 0},
		locations.OptionItem{Name: "BBQ", Description: "BBQ", Price: 0},
		locations.OptionItem{Name: "Blazing BBQ", Description: "Blazing BBQ", Price: 0},
		locations.OptionItem{Name: "Incinerator", Description: "Incinerator", Price: 0},
		locations.OptionItem{Name: "Garlic Parm", Description: "Garlic Parm", Price: 0}})
	_ = api.CreateOption("Wings Dipping Sauce", "Wings Dipping Sauce", 1, 1, gatesFourGolfandCountryClubChickenWings.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Ranch", Description: "Ranch", Price: 0},
		locations.OptionItem{Name: "Blue Cheese", Description: "Blue Cheese", Price: 0}})
	_ = api.CreateOption("Options", "Options", 1, 1, gatesFourGolfandCountryClubChickenWings.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Six Wings", Description: "Six Wings", Price: 0},
		locations.OptionItem{Name: "Twelve Wings", Description: "Twelve Wings", Price: 5}})
	_ = api.CreateOption("Options", "Options", 1, 1, gatesFourGolfandCountryClubClubNachos.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Chicken", Description: "Chicken", Price: 0},
		locations.OptionItem{Name: "Beef", Description: "Beef", Price: 0},
		locations.OptionItem{Name: "No Protein", Description: "No Protein", Price: 0}})
	_ = api.CreateOption("Salad Dressing", "Salad Dressing", 1, 1, gatesFourGolfandCountryClubHouseSalad.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Balsalmic", Description: "Balsalmic", Price: 0},
		locations.OptionItem{Name: "Caesar", Description: "Caesar", Price: 0},
		locations.OptionItem{Name: "Ranch", Description: "Ranch", Price: 0},
		locations.OptionItem{Name: "Bleu Cheese", Description: "Bleu Cheese", Price: 0},
		locations.OptionItem{Name: "Honey Mustard", Description: "Honey Mustard", Price: 0},
		locations.OptionItem{Name: "Italian", Description: "Italian", Price: 0},
		locations.OptionItem{Name: "Thousand Island", Description: "Thousand Island", Price: 0},
		locations.OptionItem{Name: "Balsalmic Vinaigrette", Description: "Balsalmic Vinaigrette", Price: 0}})
	_ = api.CreateOption("Extras", "Extras", 0, 0, gatesFourGolfandCountryClubHouseSalad.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Fried Chicken", Description: "Fried Chicken", Price: 6},
		locations.OptionItem{Name: "Grilled Chicken", Description: "Grilled Chicken", Price: 6},
		locations.OptionItem{Name: "6oz Salmon", Description: "6oz Salmon", Price: 10},
		locations.OptionItem{Name: "3 Grilled Shrimp", Description: "3 Grilled Shrimp", Price: 4},
		locations.OptionItem{Name: "3 Fried Shrimp", Description: "3 Fried Shrimp", Price: 4}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, gatesFourGolfandCountryClubHouseSalad.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Tomatoes", Description: "No Tomatoes", Price: 0},
		locations.OptionItem{Name: "No Cucumbers", Description: "No Cucumbers", Price: 0},
		locations.OptionItem{Name: "No Onions", Description: "No Onions", Price: 0},
		locations.OptionItem{Name: "No Mushrooms", Description: "No Mushrooms", Price: 0}})
	_ = api.CreateOption("Extras", "Extras", 0, 0, gatesFourGolfandCountryClubCaesarSalad.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Fried Chicken", Description: "Fried Chicken", Price: 6},
		locations.OptionItem{Name: "Grilled Chicken", Description: "Grilled Chicken", Price: 6},
		locations.OptionItem{Name: "6oz Salmon", Description: "6oz Salmon", Price: 10},
		locations.OptionItem{Name: "3 Grilled Shrimp", Description: "3 Grilled Shrimp", Price: 4},
		locations.OptionItem{Name: "3 Fried Shrimp", Description: "3 Fried Shrimp", Price: 4}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, gatesFourGolfandCountryClubCaesarSalad.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Croutons", Description: "No Croutons", Price: 0}})
	_ = api.CreateOption("Extras", "Extras", 0, 0, gatesFourGolfandCountryClubClubSalad.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Fried Chicken", Description: "Fried Chicken", Price: 6},
		locations.OptionItem{Name: "Grilled Chicken", Description: "Grilled Chicken", Price: 6},
		locations.OptionItem{Name: "6oz Salmon", Description: "6oz Salmon", Price: 10},
		locations.OptionItem{Name: "3 Grilled Shrimp", Description: "3 Grilled Shrimp", Price: 4},
		locations.OptionItem{Name: "3 Fried Shrimp", Description: "3 Fried Shrimp", Price: 4}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, gatesFourGolfandCountryClubClubSalad.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Ham", Description: "No Ham", Price: 0},
		locations.OptionItem{Name: "No Turkey", Description: "No Turkey", Price: 0},
		locations.OptionItem{Name: "No Bacon", Description: "No Bacon", Price: 0}})
	_ = api.CreateOption("Extras", "Extras", 0, 0, gatesFourGolfandCountryClubGatesFourSpinachSalad.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Fried Chicken", Description: "Fried Chicken", Price: 6},
		locations.OptionItem{Name: "Grilled Chicken", Description: "Grilled Chicken", Price: 6},
		locations.OptionItem{Name: "6oz Salmon", Description: "6oz Salmon", Price: 10},
		locations.OptionItem{Name: "3 Grilled Shrimp", Description: "3 Grilled Shrimp", Price: 4},
		locations.OptionItem{Name: "3 Fried Shrimp", Description: "3 Fried Shrimp", Price: 4}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, gatesFourGolfandCountryClubGatesFourSpinachSalad.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Craisins", Description: "No Craisins", Price: 0},
		locations.OptionItem{Name: "No Tomatoes", Description: "No Tomatoes", Price: 0},
		locations.OptionItem{Name: "No Candied Nuts", Description: "No Candied Nuts", Price: 0},
		locations.OptionItem{Name: "No Onions", Description: "No Onions", Price: 0},
		locations.OptionItem{Name: "No Blue Cheese", Description: "No Blue Cheese", Price: 0},
		locations.OptionItem{Name: "No Bacon", Description: "No Bacon", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, gatesFourGolfandCountryClubAhiTunaSteak.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0},
		locations.OptionItem{Name: "Fruit", Description: "Fruit", Price: 0},
		locations.OptionItem{Name: "Homemade Chips", Description: "Homemade Chips", Price: 0},
		locations.OptionItem{Name: "Pasta Salad", Description: "Pasta Salad", Price: 0},
		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, gatesFourGolfandCountryClubChickenFajitas.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0},
		locations.OptionItem{Name: "Fruit", Description: "Fruit", Price: 0},
		locations.OptionItem{Name: "Homemade Chips", Description: "Homemade Chips", Price: 0},
		locations.OptionItem{Name: "Pasta Salad", Description: "Pasta Salad", Price: 0},
		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 0}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, gatesFourGolfandCountryClubChickenFajitas.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Peppers", Description: "No Peppers", Price: 0},
		locations.OptionItem{Name: "No Onions", Description: "No Onions", Price: 0},
		locations.OptionItem{Name: "No Cojita Cheese", Description: "No Cojita Cheese", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, gatesFourGolfandCountryClubPortobelloBurger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0},
		locations.OptionItem{Name: "Fruit", Description: "Fruit", Price: 0},
		locations.OptionItem{Name: "Homemade Chips", Description: "Homemade Chips", Price: 0},
		locations.OptionItem{Name: "Pasta Salad", Description: "Pasta Salad", Price: 0},
		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 0}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, gatesFourGolfandCountryClubPortobelloBurger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Lettuce", Description: "No Lettuce", Price: 0},
		locations.OptionItem{Name: "No Tomato", Description: "No Tomato", Price: 0}})
	_ = api.CreateOption("Choice of Meat", "Choice of Meat", 1, 1, gatesFourGolfandCountryClubIrongateDeli.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Ham", Description: "Ham", Price: 0},
		locations.OptionItem{Name: "Roast Beef", Description: "Roast Beef", Price: 0},
		locations.OptionItem{Name: "Capicola", Description: "Capicola", Price: 0},
		locations.OptionItem{Name: "Turkey", Description: "Turkey", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, gatesFourGolfandCountryClubIrongateDeli.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0},
		locations.OptionItem{Name: "Fruit", Description: "Fruit", Price: 0},
		locations.OptionItem{Name: "Homemade Chips", Description: "Homemade Chips", Price: 0},
		locations.OptionItem{Name: "Pasta Salad", Description: "Pasta Salad", Price: 0},
		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 0}})
	_ = api.CreateOption("Choice of Bread", "Choice of Bread", 1, 1, gatesFourGolfandCountryClubIrongateDeli.ID, []locations.OptionItem{
		locations.OptionItem{Name: "White", Description: "White", Price: 0},
		locations.OptionItem{Name: "Wheat", Description: "Wheat", Price: 0},
		locations.OptionItem{Name: "Flour Wrap", Description: "Flour Wrap", Price: 0},
		locations.OptionItem{Name: "Garlic Wrap", Description: "Garlic Wrap", Price: 0},
		locations.OptionItem{Name: "Gluten Free Wrap", Description: "Gluten Free Wrap", Price: 1}})
	_ = api.CreateOption("Choice of Cheese", "Choice of Cheese", 1, 1, gatesFourGolfandCountryClubIrongateDeli.ID, []locations.OptionItem{
		locations.OptionItem{Name: "American", Description: "American", Price: 0},
		locations.OptionItem{Name: "Cheddar", Description: "Cheddar", Price: 0},
		locations.OptionItem{Name: "Swiss", Description: "Swiss", Price: 0},
		locations.OptionItem{Name: "Provolone", Description: "Provolone", Price: 0},
		locations.OptionItem{Name: "Pepperjack", Description: "Pepperjack", Price: 0},
		locations.OptionItem{Name: "No Cheese", Description: "No Cheese", Price: 0}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, gatesFourGolfandCountryClubIrongateDeli.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Lettuce", Description: "No Lettuce", Price: 0},
		locations.OptionItem{Name: "No Tomato", Description: "No Tomato", Price: 0},
		locations.OptionItem{Name: "No Onions", Description: "No Onions", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, gatesFourGolfandCountryClubPhillyCheesesteak.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0},
		locations.OptionItem{Name: "Fruit", Description: "Fruit", Price: 0},
		locations.OptionItem{Name: "Homemade Chips", Description: "Homemade Chips", Price: 0},
		locations.OptionItem{Name: "Pasta Salad", Description: "Pasta Salad", Price: 0},
		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 0}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, gatesFourGolfandCountryClubPhillyCheesesteak.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Peppers", Description: "No Peppers", Price: 0},
		locations.OptionItem{Name: "No Onions", Description: "No Onions", Price: 0},
		locations.OptionItem{Name: "No Mushrooms", Description: "No Mushrooms", Price: 0},
		locations.OptionItem{Name: "No Provolone", Description: "No Provolone", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, gatesFourGolfandCountryClubTheItalian.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0},
		locations.OptionItem{Name: "Fruit", Description: "Fruit", Price: 0},
		locations.OptionItem{Name: "Homemade Chips", Description: "Homemade Chips", Price: 0},
		locations.OptionItem{Name: "Pasta Salad", Description: "Pasta Salad", Price: 0},
		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 0}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, gatesFourGolfandCountryClubTheItalian.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Lettuce", Description: "No Lettuce", Price: 0},
		locations.OptionItem{Name: "No Oil & Vinegar", Description: "No Oil & Vinegar", Price: 0},
		locations.OptionItem{Name: "No Banana Peppers", Description: "No Banana Peppers", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, gatesFourGolfandCountryClubBlackenedSalmonSandwich.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0},
		locations.OptionItem{Name: "Fruit", Description: "Fruit", Price: 0},
		locations.OptionItem{Name: "Homemade Chips", Description: "Homemade Chips", Price: 0},
		locations.OptionItem{Name: "Pasta Salad", Description: "Pasta Salad", Price: 0},
		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 0}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, gatesFourGolfandCountryClubBlackenedSalmonSandwich.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Lettuce", Description: "No Lettuce", Price: 0},
		locations.OptionItem{Name: "No Tomato", Description: "No Tomato", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, gatesFourGolfandCountryClubJPsClubSandwich.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0},
		locations.OptionItem{Name: "Fruit", Description: "Fruit", Price: 0},
		locations.OptionItem{Name: "Homemade Chips", Description: "Homemade Chips", Price: 0},
		locations.OptionItem{Name: "Pasta Salad", Description: "Pasta Salad", Price: 0},
		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 0}})
	_ = api.CreateOption("Choice of Bread", "Choice of Bread", 1, 1, gatesFourGolfandCountryClubJPsClubSandwich.ID, []locations.OptionItem{
		locations.OptionItem{Name: "White", Description: "White", Price: 0},
		locations.OptionItem{Name: "Wheat", Description: "Wheat", Price: 0},
		locations.OptionItem{Name: "Flour Wrap", Description: "Flour Wrap", Price: 0},
		locations.OptionItem{Name: "Garlic Wrap", Description: "Garlic Wrap", Price: 0},
		locations.OptionItem{Name: "Gluten Free Wrap", Description: "Gluten Free Wrap", Price: 1}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, gatesFourGolfandCountryClubJPsClubSandwich.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Lettuce", Description: "No Lettuce", Price: 0},
		locations.OptionItem{Name: "No Tomato", Description: "No Tomato", Price: 0},
		locations.OptionItem{Name: "No Onions", Description: "No Onions", Price: 0},
		locations.OptionItem{Name: "No Mayo", Description: "No Mayo", Price: 0},
		locations.OptionItem{Name: "No Bacon", Description: "No Bacon", Price: 0},
		locations.OptionItem{Name: "No Swiss", Description: "No Swiss", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, gatesFourGolfandCountryClubBuffaloChickenSandwich.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0},
		locations.OptionItem{Name: "Fruit", Description: "Fruit", Price: 0},
		locations.OptionItem{Name: "Homemade Chips", Description: "Homemade Chips", Price: 0},
		locations.OptionItem{Name: "Pasta Salad", Description: "Pasta Salad", Price: 0},
		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 0}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, gatesFourGolfandCountryClubBuffaloChickenSandwich.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Lettuce", Description: "No Lettuce", Price: 0},
		locations.OptionItem{Name: "No Tomato", Description: "No Tomato", Price: 0},
		locations.OptionItem{Name: "No Onions", Description: "No Onions", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 0, 0, gatesFourGolfandCountryClubJPsAce.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 2},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 2},
		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 2},
		locations.OptionItem{Name: "Fruit", Description: "Fruit", Price: 2},
		locations.OptionItem{Name: "Homemade Chips", Description: "Homemade Chips", Price: 2},
		locations.OptionItem{Name: "Pasta Salad", Description: "Pasta Salad", Price: 2},
		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 2}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, gatesFourGolfandCountryClubJPsAce.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Chili", Description: "No Chili", Price: 0},
		locations.OptionItem{Name: "No Coleslaw", Description: "No Coleslaw", Price: 0},
		locations.OptionItem{Name: "No Mustard", Description: "No Mustard", Price: 0},
		locations.OptionItem{Name: "No Diced Onions", Description: "No Diced Onions", Price: 0}})
	_ = api.CreateOption("Meat Temperature", "Meat Temperature", 1, 1, gatesFourGolfandCountryClubGatesFourBurger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Well Done", Description: "Well Done", Price: 0},
		locations.OptionItem{Name: "Medium-Well", Description: "Medium-Well", Price: 0},
		locations.OptionItem{Name: "Medium", Description: "Medium", Price: 0},
		locations.OptionItem{Name: "Medium-Rare", Description: "Medium-Rare", Price: 0},
		locations.OptionItem{Name: "Rare", Description: "Rare", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, gatesFourGolfandCountryClubGatesFourBurger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0},
		locations.OptionItem{Name: "Fruit", Description: "Fruit", Price: 0},
		locations.OptionItem{Name: "Homemade Chips", Description: "Homemade Chips", Price: 0},
		locations.OptionItem{Name: "Pasta Salad", Description: "Pasta Salad", Price: 0},
		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 0}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, gatesFourGolfandCountryClubGatesFourBurger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Lettuce", Description: "No Lettuce", Price: 0},
		locations.OptionItem{Name: "No Tomato", Description: "No Tomato", Price: 0},
		locations.OptionItem{Name: "No Onions", Description: "No Onions", Price: 0},
		locations.OptionItem{Name: "No Cheese", Description: "No Cheese", Price: 0}})
	_ = api.CreateOption("Meat Temperature", "Meat Temperature", 1, 1, gatesFourGolfandCountryClubBigSkyBurger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Well Done", Description: "Well Done", Price: 0},
		locations.OptionItem{Name: "Medium-Well", Description: "Medium-Well", Price: 0},
		locations.OptionItem{Name: "Medium", Description: "Medium", Price: 0},
		locations.OptionItem{Name: "Medium-Rare", Description: "Medium-Rare", Price: 0},
		locations.OptionItem{Name: "Rare", Description: "Rare", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, gatesFourGolfandCountryClubBigSkyBurger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0},
		locations.OptionItem{Name: "Fruit", Description: "Fruit", Price: 0},
		locations.OptionItem{Name: "Homemade Chips", Description: "Homemade Chips", Price: 0},
		locations.OptionItem{Name: "Pasta Salad", Description: "Pasta Salad", Price: 0},
		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 0}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, gatesFourGolfandCountryClubBigSkyBurger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Bacon", Description: "No Bacon", Price: 0},
		locations.OptionItem{Name: "No Cheddar", Description: "No Cheddar", Price: 0},
		locations.OptionItem{Name: "No Crispy Onions", Description: "No Crispy Onions", Price: 0}})
	_ = api.CreateOption("Meat Temperature", "Meat Temperature", 1, 1, gatesFourGolfandCountryClubCarolinaBurger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "Well Done", Description: "Well Done", Price: 0},
		locations.OptionItem{Name: "Medium-Well", Description: "Medium-Well", Price: 0},
		locations.OptionItem{Name: "Medium", Description: "Medium", Price: 0},
		locations.OptionItem{Name: "Medium-Rare", Description: "Medium-Rare", Price: 0},
		locations.OptionItem{Name: "Rare", Description: "Rare", Price: 0}})
	_ = api.CreateOption("Choice of Side", "Choice of Side", 1, 1, gatesFourGolfandCountryClubCarolinaBurger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "French Fries", Description: "French Fries", Price: 0},
		locations.OptionItem{Name: "Sweet Potato Fries", Description: "Sweet Potato Fries", Price: 0},
		locations.OptionItem{Name: "Onion Rings", Description: "Onion Rings", Price: 0},
		locations.OptionItem{Name: "Fruit", Description: "Fruit", Price: 0},
		locations.OptionItem{Name: "Homemade Chips", Description: "Homemade Chips", Price: 0},
		locations.OptionItem{Name: "Pasta Salad", Description: "Pasta Salad", Price: 0},
		locations.OptionItem{Name: "Potato Salad", Description: "Potato Salad", Price: 0}})
	_ = api.CreateOption("Remove Options", "Remove Options", 0, 0, gatesFourGolfandCountryClubCarolinaBurger.ID, []locations.OptionItem{
		locations.OptionItem{Name: "No Mustard", Description: "No Mustard", Price: 0},
		locations.OptionItem{Name: "No Onions", Description: "No Onions", Price: 0},
		locations.OptionItem{Name: "No Coleslaw", Description: "No Coleslaw", Price: 0}})
	// seed orders
	log.Println("Seeding Completed")
}