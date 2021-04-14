package model

type (
	// Statistics - Statistics
	Statistics struct {
		Conditions []Conditions `json:"conditions"`
		Result     Result       `json:"result" validate:"required"`
		Others     Others       `json:"others,omitempty" validate:"omitempty"`
	}

	// Conditions - Conditions
	Conditions struct {
		FieldName string   `json:"field_name" validate:"required"`
		Operators Operator `json:"operators" validate:"required"`
	}

	// Result - Result
	Result struct {
		GroupBy []StatisticsGroupBy `json:"group_by" validate:"required"`
		Fields  []Field             `json:"fields" validate:"required"`
	}

	// Others - Others
	Others struct {
		Sort  Sort `json:"sort,omitempty" validate:"omitempty,required"`
		Limit int  `json:"limit,omitempty" validate:"omitempty,required"`
	}

	// Operator - Operator
	Operator struct {
		Equal              interface{} `json:"eq"`
		GreaterThan        interface{} `json:"gt"`
		LessThan           interface{} `json:"lt"`
		GreaterThanOrEqual interface{} `json:"gte"`
		LessThanOrEqual    interface{} `json:"lte"`
		NotEqual           interface{} `json:"ne"`
		Contains           interface{} `json:"contains"`
	}

	// Field - Field
	Field struct {
		FieldName string      `json:"field_name" validate:"required"`
		Column    interface{} `json:"column" validate:"required"`
		Operator  string      `json:"operator" validate:"required"`
	}

	// Sort - Sort
	Sort struct {
		Type  string `json:"type,omitempty" validate:"omitempty,required"`
		Field string `json:"field,omitempty" validate:"omitempty,required"`
	}

	// StatisticsGroupBy - StatisticsGroupBy
	StatisticsGroupBy struct {
		Key   string `json:"key" validate:"required"`
		Value string `json:"value" validate:"required"`
	}
)
