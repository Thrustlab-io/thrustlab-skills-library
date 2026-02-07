package main

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerResources(s *server.MCPServer) {
	commandsResource := mcp.NewResource(
		"clay://commands",
		"List of all available Clay MCP commands",
		mcp.WithResourceDescription("Complete list of available commands and their descriptions"),
		mcp.WithMIMEType("text/plain"),
	)
	s.AddResource(commandsResource, listCommandsHandler)

	businessTypesResource := mcp.NewResource(
		"clay://business-types",
		"List of all Google Maps business types",
		mcp.WithResourceDescription("Complete list of valid business types for geography searches"),
		mcp.WithMIMEType("text/plain"),
	)
	s.AddResource(businessTypesResource, listBusinessTypesHandler)
}

func listCommandsHandler(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
	commands := `Clay MCP Server - Available Commands
==========================================

Configuration Commands:
-----------------------
1. set_workspace_id
   Set your Clay workspace ID for API access
   Parameters:
   - workspace_id (required): Your Clay workspace ID (e.g., 757984)

2. set_session_cookie
   Set your Clay session cookie for authentication
   Parameters:
   - session_cookie (required): Session cookie value (starts with s%3A...)

Workbook Commands:
------------------
3. create_workbook
   Create a new Clay workbook and set it as default
   Parameters:
   - name (required): Name for the new workbook
   - folder_id (optional): Folder ID to create the workbook in

Search Commands:
----------------
4. search_companies_by_industry
   Search for companies by industry using Clay's Mixrank/LinkedIn data
   Parameters:
   - workbook_id (optional): Uses current workbook if not specified
   - industries (required): Comma-separated list (e.g., 'Accounting,Consulting')
   - countries (optional): Comma-separated country names
   - company_sizes (optional): Comma-separated size codes (1,2,10,50,200,500,1000,5000,10000)
   - keywords (optional): Comma-separated description keywords
   - limit (optional): Maximum results (default: 25000)
   - annual_revenues (optional): Comma-separated revenue ranges (e.g., '1M-5M,10M-25M')
   - minimum_member_count (optional): Minimum number of employees (e.g., 100)
   - maximum_member_count (optional): Maximum number of employees (e.g., 200)

5. search_businesses_by_geography
   Search for local businesses by geography using Google Maps
   Two search modes: business types or free text query (provide one or the other)
   Parameters:
   - workbook_id (optional): Uses current workbook if not specified
   - latitude (required): Latitude coordinate (e.g., 51.049)
   - longitude (required): Longitude coordinate (e.g., 3.725)
   - proximity_km (required): Search radius in kilometers
   - business_types (optional): Comma-separated types (see clay://business-types)
   - query (optional): Free text search (e.g., 'slager', 'pizza', 'kapper')
   - num_results (optional): Maximum number of results (default: 100)
   - table_name (optional): Custom table name
   - table_emoji (optional): Table emoji icon

Resources:
----------
- clay://commands - This list of commands
- clay://business-types - List of all valid Google Maps business types
`

	content := mcp.TextResourceContents{
		URI:      req.Params.URI,
		MIMEType: "text/plain",
		Text:     commands,
	}

	return []mcp.ResourceContents{&content}, nil
}

func listBusinessTypesHandler(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
	businessTypes := `Google Maps Business Types
===========================

Common Business Types:
----------------------
accounting, airport, amusement_park, aquarium, art_gallery, atm, bakery, bank, bar,
beauty_salon, bicycle_store, book_store, bowling_alley, bus_station, cafe, campground,
car_dealer, car_rental, car_repair, car_wash, casino, cemetery, church, city_hall,
clothing_store, convenience_store, courthouse, dentist, department_store, doctor,
drugstore, electrician, electronics_store, embassy, fire_station, florist, funeral_home,
furniture_store, gas_station, gym, hair_care, hardware_store, hindu_temple, home_goods_store,
hospital, insurance_agency, jewelry_store, laundry, lawyer, library, light_rail_station,
liquor_store, local_government_office, locksmith, lodging, meal_delivery, meal_takeaway,
mosque, movie_rental, movie_theater, moving_company, museum, night_club, painter, park,
parking, pet_store, pharmacy, physiotherapist, plumber, police, post_office, primary_school,
real_estate_agency, restaurant, roofing_contractor, rv_park, school, secondary_school,
shoe_store, shopping_mall, spa, stadium, storage, store, subway_station, supermarket,
synagogue, taxi_stand, tourist_attraction, train_station, transit_station, travel_agency,
university, veterinary_care, zoo

Usage Examples:
---------------
Single type:
  book_store

Multiple types (comma-separated):
  book_store,library,university

Restaurant types:
  restaurant,cafe,bar,meal_delivery

Retail:
  clothing_store,shoe_store,jewelry_store,department_store

Services:
  hair_care,beauty_salon,spa,gym

Note: Use underscores (_) not spaces in business type names.
`

	content := mcp.TextResourceContents{
		URI:      req.Params.URI,
		MIMEType: "text/plain",
		Text:     businessTypes,
	}

	return []mcp.ResourceContents{&content}, nil
}
