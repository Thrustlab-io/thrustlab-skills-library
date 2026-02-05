# Clay.com Table Creation Operations - API Analysis

## Summary

Found **2 table creation operations** in the HAR file:

1. **Search businesses by Industry** (LinkedIn/Mixrank data)
2. **Search businesses by Google Maps Geography**

---

## 1. Search Businesses by Industry (Mixrank/LinkedIn)

### Endpoint
```
POST https://api.clay.com/v3/workspaces/757984/wizard/evaluate-step
```

### Key Parameters

**Wizard Config:**
- `wizardId`: `"find-companies"`
- `wizardStepId`: `"companies-search"`
- `workbookId`: `"wb_0t9zx9hiFdeR4VDUCHj"`

**Search Criteria:**
```json
{
  "actionKey": "find-lists-of-companies-with-mixrank-source",
  "actionPackageId": "e251a70e-46d7-4f3a-b3ef-a211ad3d8bd2",
  "inputs": {
    "industries": ["Accounting"],
    "country_names": ["Belgium"],
    "sizes": ["50"],
    "description_keywords": ["boekhouding"],
    "limit": 25000
  }
}
```

**Table Schema Created:**
- Name (text)
- Description (text)
- Primary Industry (text)
- Size (select with 9 options: Self-employed, 2-10, 11-50, etc.)
- Type (text)
- Location (text)
- Country (text)
- Domain (url)
- LinkedIn URL (url, dedupe field)

**Response:**
```json
{
  "table": {
    "tableId": "t_0ta00b4zvV4kUUdu9ec",
    "tableName": "Find companies Table",
    "isNewTable": true
  },
  "recordCount": 2
}
```

---

## 2. Search Businesses by Google Maps Geography

### Endpoint
```
POST https://api.clay.com/v3/tables
```

### Key Parameters

**Table Config:**
```json
{
  "workspaceId": "757984",
  "type": "company",
  "template": "empty",
  "icon": {"emoji": "🌙"},
  "name": "⚡️ Find local businesses using Google Maps Table",
  "workbookId": "wb_0t9zx9hiFdeR4VDUCHj"
}
```

**Source Settings:**
```json
{
  "actionKey": "google-review-source-v3",
  "actionPackageId": "3282a1c7-6bb0-497e-a34b-32268e104e55",
  "inputs": {
    "usePreferredGoogleApi": "true",
    "map": {
      "latitude": 51.04945049862644,
      "longitude": 3.725325757953186,
      "proximity": 13
    },
    "dynamicFields|searchType": "businessTypes",
    "dynamicFields|businessTypes": ["art_gallery"]
  }
}
```

**Table Schema Created:**
- Name (text)
- Google Maps URL (url)
- Description (text)
- Website (url)
- Phone (text)
- Address (text)
- Rating (number)
- Reviews Count (number)

**Response:**
```json
{
  "id": "t_0ta00h54TAPZKcPnWoZ",
  "name": "⚡️ Find local businesses using Google Maps Table",
  "type": "company"
}
```

---

## Key Differences

| Feature | Industry Search | Google Maps Search |
|---------|----------------|-------------------|
| **Endpoint** | `/wizard/evaluate-step` | `/tables` |
| **Data Source** | Mixrank/LinkedIn | Google Maps |
| **Geography Type** | Country-level | Lat/Lng + radius |
| **Search By** | Industry + keywords | Business types |
| **ID Field** | `linkedin_company_id` | Google Maps `id` |
| **Dedupe** | LinkedIn URL | Google Maps URL |

---

## Authentication

Both requests use:
- **Cookie**: `claysession=s%3A...` (session token)
- **Header**: `X-Clay-Frontend-Version: v20260205_151537Z_19e7945c5e`
- **Origin**: `https://app.clay.com`

---

## Files Created

1. [clay-post-requests-filtered.json](clay-post-requests-filtered.json) - All Clay POST requests
2. [google-maps-table-creation.json](google-maps-table-creation.json) - Google Maps operation
3. [clay-table-creation-requests.json](clay-table-creation-requests.json) - Both operations
4. This summary document

---

## Next Steps for MCP Server

To replicate these operations programmatically:

1. **Extract session cookie** from authenticated browser session
2. **Implement POST requests** with proper headers
3. **Build request builders** for both search types:
   - Industry search: Use wizard flow
   - Geography search: Use direct table creation
4. **Handle responses** to get table IDs and monitor progress
5. **Iterate on prompts** by modifying `inputs` parameters
