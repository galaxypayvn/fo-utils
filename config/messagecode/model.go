package messagecode

type messageCode struct {
	HTTPCode int    `json:"http_code"`
	Message  string `json:"messasge"`
}

type strapiMessageCodeResp struct {
	Data  []strapiMessageCode `json:"data"`
	Error struct {
		Status  int                    `json:"status"`
		Name    string                 `json:"name"`
		Message string                 `json:"message"`
		Details ValidationErrorDetails `json:"details"`
	} `json:"error"`
	Meta strapiMeta `json:"meta"`
}

type strapiMessageCode struct {
	ID       int    `json:"id"`
	Code     int    `json:"code"`
	Locale   string `json:"locale"`
	Message  string `json:"message"`
	HTTPCode int    `json:"http_code"`
}

type strapiMeta struct {
	Pagination strapiPagination `json:"pagination"`
}

type strapiPagination struct {
	Page      int `json:"page"`
	PageSize  int `json:"pageSize"`
	PageCount int `json:"pageCount"`
	Total     int `json:"total"`
}

type CreateMessageCodeReq struct {
	Code        int    `json:"code"`
	HttpCode    int    `json:"http_code"`
	Message     string `json:"message"`
	Description string `json:"description"`
}

type SuccessResponse struct {
	Data struct {
		ID          int    `json:"id"`
		Code        int    `json:"code"`
		Message     string `json:"message"`
		HttpCode    int    `json:"http_code"`
		CreatedAt   string `json:"createdAt"`
		UpdatedAt   string `json:"updatedAt"`
		PublishedAt string `json:"publishedAt"`
		Locale      string `json:"locale"`
	} `json:"data"`
	Meta struct{} `json:"meta"`
}

type ValidationErrorDetail struct {
	Path    []string `json:"path"`
	Message string   `json:"message"`
	Name    string   `json:"name"`
}

type ValidationErrorDetails struct {
	Errors []ValidationErrorDetail `json:"errors"`
}

type ErrorResponse struct {
	Data  interface{} `json:"data"`
	Error struct {
		Status  int                    `json:"status"`
		Name    string                 `json:"name"`
		Message string                 `json:"message"`
		Details ValidationErrorDetails `json:"details"`
	} `json:"error"`
}
