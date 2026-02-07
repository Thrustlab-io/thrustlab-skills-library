---
name: gtm-strategy-generator
description: Generates a GTM strategy with ICP definition, market mapping, and a concrete Clay search plan for immediate campaign execution.
argument-hint: "client website"
---
# General
Put your output as a single md file in <current_working_directory>/strategies/CLIENT.md

### **Research Protocol**

1. **Always start with web_fetch on client website** for primary intelligence
2. **Follow with 3-5 targeted web_search queries** using:
   - `"[company]" + industry terms + geography`
   - `"[competitor names]" + regulatory/compliance terms`
   - `"[market segment]" + local market qualifiers`
3. **Never use generic searches** - always include industry-specific terminology
4. **Document everything** with [Source: URL] citations

### **Strategy Sections**

Generate exactly **6 sections** in this order:

#### **1. Company Snapshot** (1 page max, bullet format)
- Business model, founding, HQ, funding
- Current market position and key differentiators
- Revenue model and pricing
- Market context and trends

#### **2. ICP & Buying Personas**
- **Ideal Customer Profile:** firmographics (size, revenue, industry, geography), technographics, behavioral signals, pain points
- **Disqualifiers:** clear list of who is NOT a fit
- **Buying Personas:** 2-3 key personas with titles, responsibilities, KPIs, and what they care about
- **Budget & Decision Process:** typical spend, decision cycle length, budget holders

#### **3. TAM to SAM Filtering**
- Total Addressable Market estimate with reasoning
- Serviceable Addressable Market with specific filters
- Priority segments ranked by fit and urgency

#### **4. Competitive Landscape**
- 3-5 direct competitors with positioning
- Gaps and opportunities to exploit
- Key differentiators vs. competition

#### **5. Messaging Framework**
- Core value proposition (1-2 sentences)
- Pain-to-outcome mapping per persona (table format, no full email copy)
- 3-5 key proof points / social proof angles

#### **6. Clay Search Plan**
This is the most critical section. Provide a concrete, executable plan for building prospect lists in Clay.com. Structure it as follows:

**Search Type:** State whether to use `search_companies_by_industry` (LinkedIn/Mixrank data) or `search_businesses_by_geography` (Google Maps) or both.

**For each search, specify ALL parameters:**

For company searches:
- `industries`: LinkedIn industries (e.g., "Computer Software", "Financial Services") — see clay://industries for valid values
- `keywords`: exact comma-separated keywords to use
- `countries`: target countries
- `company_sizes`: size codes (1, 2, 10, 50, 200, 500, 1000, 5000, 10000)
- `annual_revenues`: revenue ranges (e.g., 1M-5M, 5M-10M)
- `min_linkedin_members` / `max_linkedin_members`: if relevant

For geography searches:
- `latitude`, `longitude`, `proximity_km`: exact coordinates and radius
- `business_types` or `query`: what to search for
- `num_results`: how many results to pull

**Multiple searches:** If the ICP spans multiple segments, define separate searches for each. Label them (e.g., "Search 1: Professional Services in Belgium", "Search 2: Tech Companies in Netherlands").

**Expected volume:** Estimate how many companies each search should return.

**Prioritization:** Rank the searches by expected quality/fit.

### **Output Constraints**

- **No fluff phrases:** "leverage," "synergy," "revolutionary"
- **No long paragraphs:** Max 3 sentences per bullet
- **No theory without practice:** Every concept needs execution steps
- **No assumptions without labels:** Mark as [Assumption - needs validation]

### **When Information is Missing**

1. First attempt to find via web research
2. If unavailable, provide 3-5 specific clarifying questions
3. Proceed with clearly marked [Assumptions] and risk notes
