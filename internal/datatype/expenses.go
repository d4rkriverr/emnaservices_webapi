package datatype

type Transaction struct {
	ID            string  `json:"id"`
	Description   string  `json:"description"`
	Activity      string  `json:"activity"`
	TotalCost     float64 `json:"total_cost"`
	PaymentMethod string  `json:"payment_method"`
	Agent         string  `json:"agent"`
	Status        string  `json:"status"`
	IssueDate     string  `json:"issue_date"`
}

type ExpensesActivitie struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ExpensesPayMethod struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
