package models

type SoftVulDetails struct {
	Cpe                  string   `json:"CPE"`
	Cve_id               string   `json:"CVE_ID"`
	Cvss                 string   `json:"CVSS"`
	Cvss_severity        string   `json:"CVSS_SEVERITY"`
	Cvss_vector          string   `json:"CVSS_VECTOR"`
	Cwe                  []string `json:"CWE"`
	Description          string   `json:"DESCRIPTION"`
	Exploitability_score string   `json:"EXPLOITABILITY_SCORE"`
	Impact_score         string   `json:"IMPACT_SCORE"`
	Product              string   `json:"PRODUCT"`
	Vendor               string   `json:"VENDOR"`
	Version              string   `json:"VERSION"`
}

type ResponsesofVulDetails struct {
	Data_list  []SoftVulDetails
	Total_Rows int64 `json:"total_rows"`
}