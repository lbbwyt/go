// common
package logic

type SchoolInfo struct {
	SchoolId   int    `json:"school_id"`
	Province   string `json:"province"`
	City       string `json:"city"`
	SchoolType int    `school_type`
	SchoolName string `json:"school_name"`
}
