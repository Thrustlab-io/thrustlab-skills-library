package clay

import (
	"encoding/json"
	"fmt"
)

func (c *Client) SearchCompaniesByIndustry(params SearchCompaniesParams) (*SearchCompaniesResult, error) {
	if params.WorkbookID == "" {
		return nil, fmt.Errorf("no workbook specified")
	}

	countries := ensureSlice(params.Countries)
	sizes := ensureSlice(params.CompanySizes)
	keywords := ensureSlice(params.Keywords)
	annualRevenues := ensureSlice(params.AnnualRevenues)

	var minMember, maxMember interface{}
	if params.MinimumMemberCount != nil {
		minMember = *params.MinimumMemberCount
	}
	if params.MaximumMemberCount != nil {
		maxMember = *params.MaximumMemberCount
	}

	payload := map[string]any{
		"workbookId":       params.WorkbookID,
		"wizardId":         "find-companies",
		"wizardStepId":     "companies-search",
		"sessionId":        GenerateSessionID(),
		"currentStepIndex": 0,
		"outputs":          []any{},
		"firstUseCase":     nil,
		"parentFolderId":   nil,
		"formInputs": map[string]any{
			"clientSettings": map[string]any{
				"tableType": "company",
			},
			"requiredDataPoint": nil,
			"basicFields": []map[string]any{
				{"name": "Name", "dataType": "text", "formulaText": "{{source}}.name"},
				{"name": "Description", "dataType": "text", "formulaText": "{{source}}.description"},
				{"name": "Primary Industry", "dataType": "text", "formulaText": "{{source}}.industry"},
				{"name": "Size", "dataType": "select", "formulaText": "{{source}}.size", "options": []map[string]any{
					{"id": "700e1eee-532f-4156-a562-4ee21c39048d", "text": "Self-employed", "color": "yellow"},
					{"id": "2d73a8b9-a0a1-461e-9eb2-5e2018f0bc49", "text": "2-10 employees", "color": "blue"},
					{"id": "06f139d8-6876-4b2e-9ea9-eb35bfa05a33", "text": "11-50 employees", "color": "green"},
					{"id": "d19ea21b-b725-484a-8613-99292ae65800", "text": "51-200 employees", "color": "red"},
					{"id": "7b157fc4-0534-4891-b0f2-65c5c62c2f9f", "text": "201-500 employees", "color": "violet"},
					{"id": "947e06c6-e8b3-496c-b773-a75962fda7ae", "text": "501-1,000 employees", "color": "grey"},
					{"id": "d5218b6c-d7bd-4c33-bc64-cbb4215d2293", "text": "1,001-5,000 employees", "color": "orange"},
					{"id": "590700a8-3f25-4b14-beff-fc85eb91dd10", "text": "5,001-10,000 employees", "color": "pink"},
					{"id": "a2992f14-887c-4333-93bf-90ea20836633", "text": "10,001+ employees", "color": "yellow"},
				}},
				{"name": "Type", "dataType": "text", "formulaText": "{{source}}.type"},
				{"name": "Location", "dataType": "text", "formulaText": "{{source}}.location"},
				{"name": "Country", "dataType": "text", "formulaText": "{{source}}.country"},
				{"name": "Domain", "dataType": "url", "formulaText": "{{source}}.domain"},
				{"name": "LinkedIn URL", "dataType": "url", "formulaText": "{{source}}.linkedin_url", "isDedupeField": true},
			},
			"type": "companies",
			"typeSettings": map[string]any{
				"name":               "Find companies",
				"iconType":           "Buildings",
				"actionKey":          MixrankActionKey,
				"actionPackageId":    MixrankActionPackageID,
				"previewTextPath":    "name",
				"defaultPreviewText": "Profile",
				"recordsPath":        "companies",
				"idPath":             "linkedin_company_id",
				"scheduleConfig": map[string]any{
					"runSettings": "once",
				},
				"inputs": map[string]any{
					"industries":                        params.Industries,
					"country_names":                     countries,
					"country_names_exclude":              []string{},
					"sizes":                             sizes,
					"description_keywords":              keywords,
					"description_keywords_exclude":      []string{},
					"limit":                             nil,
					"types":                             []string{},
					"locations":                         []string{},
					"locations_exclude":                 []string{},
					"funding_amounts":                   []string{},
					"annual_revenues":                   annualRevenues,
					"industries_exclude":                []string{},
					"minimum_follower_count":            nil,
					"minimum_member_count":              minMember,
					"maximum_member_count":              maxMember,
					"semantic_description":              "",
					"company_identifier":                []string{},
					"startFromCompanyType":              "company_identifier",
					"exclude_company_identifiers_mixed": []string{},
					"exclude_entities_configuration":    []string{},
					"exclude_entities_bitmap":           nil,
					"previous_entities_bitmap":          nil,
					"derived_industries":                []string{},
					"derived_subindustries":             []string{},
					"derived_subindustries_exclude":     []string{},
					"derived_revenue_streams":           []string{},
					"derived_business_types":            []string{},
					"tableId":                           nil,
					"domainFieldId":                     nil,
					"useRadialKnn":                      false,
					"radialKnnMinScore":                 nil,
					"has_resolved_domain":               nil,
					"resolved_domain_is_live":           nil,
					"resolved_domain_redirects":         nil,
					"name":                              "",
				},
				"hasEvaluatedInputs": true,
				"previewActionKey":   "find-lists-of-companies-with-mixrank-source-preview",
			},
		},
	}

	url := fmt.Sprintf("%s/workspaces/%s/wizard/evaluate-step", c.APIBase, c.WorkspaceID)
	resp, err := c.makeRequest("POST", url, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to create company search: %w\nResponse: %s", err, string(resp))
	}

	var result map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	output := result["output"].(map[string]any)
	table := output["table"].(map[string]any)

	return &SearchCompaniesResult{
		TableID:     fmt.Sprintf("%v", table["tableId"]),
		TableName:   fmt.Sprintf("%v", table["tableName"]),
		RecordCount: output["recordCount"].(float64),
	}, nil
}

func (c *Client) SearchBusinessesByGeography(params SearchBusinessesParams) (*SearchBusinessesResult, error) {
	if params.WorkbookID == "" {
		return nil, fmt.Errorf("no workbook specified")
	}

	if params.NumResults <= 0 {
		params.NumResults = 100
	}
	if params.TableName == "" {
		params.TableName = "\u26a1\ufe0f Find local businesses using Google Maps Table"
	}
	if params.TableEmoji == "" {
		params.TableEmoji = "\U0001f319"
	}

	isFreeText := params.Query != ""
	businessTypes := params.BusinessTypes
	if isFreeText && len(businessTypes) == 0 {
		businessTypes = []string{"store"}
	}
	if !isFreeText && len(businessTypes) == 0 {
		return nil, fmt.Errorf("either business_types or query must be provided")
	}

	businessTypesJSON, _ := json.Marshal(businessTypes)

	inputs := map[string]any{
		"usePreferredGoogleApi":       "true",
		"map":                         fmt.Sprintf(`{"latitude":%f,"longitude":%f,"proximity":%d}`, params.Latitude, params.Longitude, int(params.ProximityKm)),
		"numResults":                  fmt.Sprintf("%d", params.NumResults),
		"dynamicFields|businessTypes": string(businessTypesJSON),
	}

	if isFreeText {
		inputs["dynamicFields|searchType"] = `"freeText"`
		inputs["dynamicFields|searchType_displayName"] = `"Free text"`
		inputs["dynamicFields|query"] = fmt.Sprintf(`"%s"`, params.Query)
	} else {
		inputs["dynamicFields|searchType"] = `"businessTypes"`
		inputs["dynamicFields|searchType_displayName"] = `"Business types"`
	}

	payload := map[string]any{
		"icon": map[string]string{
			"emoji": params.TableEmoji,
		},
		"workspaceId": c.WorkspaceID,
		"type":        "company",
		"template":    "empty",
		"name":        params.TableName,
		"workbookId":  params.WorkbookID,
		"callerName":  "source creator modal",
		"sourceSettings": map[string]any{
			"addSource": map[string]any{
				"name": "Find local businesses using Google Maps",
				"source": map[string]any{
					"name":        "\u26a1\ufe0f Find local businesses using Google Maps",
					"workspaceId": c.WorkspaceID,
					"type":        "v3-action",
					"typeSettings": map[string]any{
						"name":        "Find local businesses using Google Maps",
						"description": "Pull local businesses from a specific location on Google Maps",
						"stages":      []string{"Inputs"},
						"categories":  []string{"FIND"},
						"customSignalSettings": map[string]any{
							"categories":  []string{"SOURCING"},
							"rank":        3,
							"title":       "Monitor local businesses using Google Maps",
							"description": "Monitor local businesses from a specific location on Google Maps",
						},
						"iconType":              "GoogleMapsSource",
						"actionKey":             GoogleMapsActionKey,
						"actionPackageId":       GoogleMapsActionPackageID,
						"defaultPreviewText":    "Business Found",
						"recordsPath":           "results",
						"idPath":                "id",
						"isAdmin":               false,
						"dedupeOnUniqueIds":     true,
						"isPaginationAvailable": true,
						"tableType":             "company",
						"costEstimate":          1,
						"sourceTableOutputs": map[string]any{
							"spreadsheet": map[string]any{
								"newFieldsNameToPaths": map[string]any{
									"Name":            map[string]any{"path": []string{"name"}, "type": "text"},
									"Google Maps URL": map[string]any{"path": []string{"googleMapsPlaceLink"}, "type": "url"},
									"Description":     map[string]any{"path": []string{"description"}, "type": "text"},
									"Website":         map[string]any{"path": []string{"website"}, "type": "url"},
									"Phone":           map[string]any{"path": []string{"phone"}, "type": "text"},
									"Address":         map[string]any{"path": []string{"address"}, "type": "text"},
									"Rating":          map[string]any{"path": []string{"rating"}, "type": "number"},
									"Reviews Count":   map[string]any{"path": []string{"reviews.count"}, "type": "number"},
								},
							},
						},
						"scheduleConfig": map[string]any{
							"runSettings": "once",
						},
						"inputs": inputs,
					},
				},
				"isPinned": true,
			},
			"addBasicFields": []map[string]any{
				{"name": "Name", "dataType": "text", "formulaText": "{{source}}.name"},
				{"name": "Google Maps URL", "dataType": "url", "formulaText": "{{source}}.googleMapsPlaceLink"},
				{"name": "Description", "dataType": "text", "formulaText": "{{source}}.description"},
				{"name": "Website", "dataType": "url", "formulaText": "{{source}}.website"},
				{"name": "Phone", "dataType": "text", "formulaText": "{{source}}.phone"},
				{"name": "Address", "dataType": "text", "formulaText": "{{source}}.address"},
				{"name": "Rating", "dataType": "number", "formulaText": "{{source}}.rating"},
				{"name": "Reviews Count", "dataType": "number", "formulaText": "{{source}}.reviews.count"},
			},
		},
	}

	url := fmt.Sprintf("%s/tables", c.APIBase)
	resp, err := c.makeRequest("POST", url, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to create geography search: %w\nResponse: %s", err, string(resp))
	}

	var result map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	tableID, tableName := extractTableInfo(result)

	return &SearchBusinessesResult{
		TableID:   tableID,
		TableName: tableName,
	}, nil
}

func ensureSlice(s []string) []string {
	if s == nil {
		return []string{}
	}
	return s
}

func extractTableInfo(result map[string]any) (string, string) {
	var tableID, tableName string
	if tableInfo, ok := result["table"].(map[string]any); ok {
		if id, ok := tableInfo["id"].(string); ok {
			tableID = id
		}
		if name, ok := tableInfo["name"].(string); ok {
			tableName = name
		}
	}
	if tableID == "" {
		if id, ok := result["id"].(string); ok {
			tableID = id
		}
	}
	if tableName == "" {
		if name, ok := result["name"].(string); ok {
			tableName = name
		}
	}
	return tableID, tableName
}
