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

	industriesResource := mcp.NewResource(
		"clay://industries",
		"List of all LinkedIn industries for company search",
		mcp.WithResourceDescription("Complete list of valid industry names used by Clay's company search (sourced from LinkedIn)"),
		mcp.WithMIMEType("text/plain"),
	)
	s.AddResource(industriesResource, listIndustriesHandler)
}

func listCommandsHandler(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
	commands := `Clay MCP Server - Available Commands
==========================================

Workspace Management Commands (NEW):
-------------------------------------
1. add_workspace_profile
   Add or update a named workspace profile for easy switching
   Parameters:
   - name (required): Profile name (e.g., 'thrustlab', 'client-acme')
   - workspace_id (required): Clay workspace ID
   - session_cookie (required): Session cookie value
   - description (optional): Description for this workspace
   - frontend_version (optional): Frontend version header
   - set_as_default (optional): Set as default workspace (default: false)

2. list_workspace_profiles
   List all saved workspace profiles and show current active workspace

3. switch_workspace
   Switch to a different workspace at runtime (no restart required!)
   Parameters:
   - name (required): Profile name to switch to

4. get_current_workspace
   Get information about the currently active workspace

5. remove_workspace_profile
   Remove a saved workspace profile
   Parameters:
   - name (required): Profile name to remove

Legacy Configuration Commands:
-------------------------------
6. set_workspace_id
   Set your Clay workspace ID for API access (requires restart)
   Parameters:
   - workspace_id (required): Your Clay workspace ID (e.g., 757984)

7. set_session_cookie
   Set your Clay session cookie for authentication (requires restart)
   Parameters:
   - session_cookie (required): Session cookie value (starts with s%3A...)

Workbook Commands:
------------------
8. create_workbook
   Create a new Clay workbook and set it as default
   Parameters:
   - name (required): Name for the new workbook
   - folder_id (optional): Folder ID to create the workbook in

Search Commands:
----------------
9. search_companies_by_industry
   Search for companies by industry using Clay's Mixrank/LinkedIn data
   Parameters:
   - workbook_id (optional): Uses current workbook if not specified
   - industries (optional): Comma-separated LinkedIn industries (see clay://industries for valid values)
   - keywords (optional): Comma-separated description keywords (e.g., 'accounting,boekhouding,SaaS')
   - countries (optional): Comma-separated country names
   - company_sizes (optional): Comma-separated size codes (1,2,10,50,200,500,1000,5000,10000)
   - limit (optional): Maximum results (default: 25000)
   - annual_revenues (optional): Comma-separated revenue ranges (e.g., '1M-5M,10M-25M')
   - min_linkedin_members (optional): Minimum number of LinkedIn members (e.g., 100)
   - max_linkedin_members (optional): Maximum number of LinkedIn members (e.g., 200)

10. search_businesses_by_geography
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

Quick Start with Workspace Profiles:
------------------------------------
1. Add a workspace profile:
   add_workspace_profile(name='thrustlab', workspace_id='757984', session_cookie='s%3A...', set_as_default=true)

2. List all profiles:
   list_workspace_profiles()

3. Switch between workspaces instantly:
   switch_workspace(name='client-acme')

Resources:
----------
- clay://commands - This list of commands
- clay://business-types - List of all valid Google Maps business types
- clay://industries - List of all valid LinkedIn industries for company search
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

func listIndustriesHandler(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
	industries := `LinkedIn Industries (used by Clay company search)
===================================================

Accounting, Airlines/Aviation, Alternative Dispute Resolution, Alternative Medicine,
Animation, Apparel & Fashion, Architecture & Planning, Arts & Crafts,
Automotive, Aviation & Aerospace, Banking, Biotechnology, Broadcast Media,
Building Materials, Business Supplies & Equipment, Capital Markets,
Chemicals, Civic & Social Organization, Civil Engineering,
Commercial Real Estate, Computer & Network Security, Computer Games,
Computer Hardware, Computer Networking, Computer Software, Construction,
Consumer Electronics, Consumer Goods, Consumer Services, Cosmetics,
Dairy, Defense & Space, Design, E-learning, Education Management,
Electrical/Electronic Manufacturing, Entertainment, Environmental Services,
Events Services, Executive Office, Facilities Services,
Farming, Financial Services, Fine Art, Fishery, Food & Beverages,
Food Production, Fundraising, Furniture, Gambling & Casinos,
Glass, Ceramics & Concrete, Government Administration,
Government Relations, Graphic Design, Health, Wellness & Fitness,
Higher Education, Hospital & Health Care, Hospitality, Human Resources,
Import & Export, Individual & Family Services, Industrial Automation,
Information Services, Information Technology & Services, Insurance,
International Affairs, International Trade & Development,
Internet, Investment Banking, Investment Management,
Judiciary, Law Enforcement, Law Practice, Legal Services,
Legislative Office, Leisure, Travel & Tourism, Libraries,
Linguistics, Logistics & Supply Chain, Luxury Goods & Jewelry,
Machinery, Management Consulting, Maritime, Market Research,
Marketing & Advertising, Mechanical or Industrial Engineering,
Media Production, Medical Device, Medical Practice, Mental Health Care,
Military, Mining & Metals, Mobile Games, Motion Pictures & Film,
Museums & Institutions, Music, Nanotechnology, Newspapers,
Non-profit Organization Management, Oil & Energy,
Online Media, Outsourcing/Offshoring, Package/Freight Delivery,
Packaging & Containers, Paper & Forest Products,
Performing Arts, Pharmaceuticals, Philanthropy, Photography,
Plastics, Political Organization, Primary/Secondary Education,
Printing, Professional Training & Coaching, Program Development,
Public Policy, Public Relations & Communications, Public Safety,
Publishing, Railroad Manufacture, Ranching, Real Estate,
Recreational Facilities & Services, Religious Institutions,
Renewables & Environment, Research, Restaurants,
Retail, Security & Investigations, Semiconductors,
Shipbuilding, Sporting Goods, Sports, Staffing & Recruiting,
Supermarkets, Telecommunications, Textiles, Think Tanks,
Tobacco, Translation & Localization, Transportation/Trucking/Railroad,
Utilities, Venture Capital & Private Equity, Veterinary, Warehousing,
Wholesale, Wine & Spirits, Wireless, Writing & Editing

Note: These industry names are sourced from LinkedIn and are used by Clay's
company search (Mixrank/LinkedIn data source). Use exact names as listed above.
Multiple industries can be specified as comma-separated values.
`

	content := mcp.TextResourceContents{
		URI:      req.Params.URI,
		MIMEType: "text/plain",
		Text:     industries,
	}

	return []mcp.ResourceContents{&content}, nil
}
