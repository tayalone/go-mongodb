package dto

/*Create Todo Body Params */
type Create struct {
	Task string `json:"task" binding:"required"`
}

type Update struct {
	Task      *string `json:"task" binding:"required_without_all=Completed,omitempty"`
	Completed *bool   `json:"completed" binding:"required_without_all=Task ,omitempty"`
}

type GetId struct {
	ID string `uri:"id" binding:"required"`
}
