package model

type (
	// InnerJoin - InnerJoin
	InnerJoin struct {
		LocalBucketID  string           `json:"local_bucket_id,omitempty"`
		JoinWith       []JoinWith       `json:"join_with" validate:"required"`
		GroupBy        GroupBy          `json:"group_by,omitempty"`
		OtherAggregate []OtherAggregate `json:"other_aggregate,omitempty"`
	}

	// JoinWith - JoinWith
	JoinWith struct {
		JoinBucketID string                 `json:"join_bucket_id" validate:"required"`
		LocalField   string                 `json:"local_field" validate:"required"`
		ForeignField string                 `json:"foreign_field" validate:"required"`
		Filter       map[string]interface{} `json:"filter,omitempty"`
		Sort         []SortFilterStruct     `json:"sort,omitempty"`
		Select       []string               `json:"select,omitempty"`
	}

	//SortFilterStruct - SortFilterStruct
	SortFilterStruct struct {
		Field string `json:"field" validate:"required"`
		Order string `json:"order" validate:"required"`
	}

	// GroupBy - GroupBy
	GroupBy struct {
		GroupID   []GroupID   `json:"group_id" validate:"required"`
		StatField []StatField `json:"stat_field" validate:"required"`
	}

	// GroupID - GroupID
	GroupID struct {
		Key   string `json:"key" validate:"required"`
		Value string `json:"value" validate:"required"`
	}

	// StatField - StatField
	StatField struct {
		Field    string      `json:"field" validate:"required"`
		Value    interface{} `json:"value" validate:"required"`
		Operator string      `json:"operator" validate:"required"`
	}

	// OtherAggregate - OtherAggregate
	OtherAggregate struct {
		Operator   string      `json:"operator" validate:"required"`
		Defination interface{} `json:"defination" validate:"required"`
	}
)

type (
	// LogicClause - LogicClause
	LogicClause struct {
		Operator string                 `json:"operator" validate:"required"`
		Clause   map[string]interface{} `json:"clause"`
		SubQuery []LogicClause          `json:"sub_query"`
	}

	// LogicClauseWithSort - LogicClauseWithSort
	LogicClauseWithSort struct {
		Query LogicClause       `json:"query" validate:"required"`
		Sort  map[string]string `json:"sort"`
	}
)
