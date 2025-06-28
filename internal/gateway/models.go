package gateway


// Ошибка
type ErrorResponse struct {
	Error	string `json:"error" example:"some error"`
}

// Универсальное тело ответа задачи
type TaskResponse struct {
	ID			string `json:"id" example:"a81bc81b-dead-4e5d-abff-90865d1e13b1"`
	Title		string `json:"title" example:"Купить хлеб"`
	Description	string `json:"description" example:"В ближайшем магазине"`
	Status		string `json:"status" example:"TODO"`
	Deadline	string `json:"deadline" example:"2025-07-01T12:00:00Z"`
	CreatedAt	string `json:"createdAt" example:"2025-07-01T12:00:00Z"`
	UpdatedAt	string `json:"UpdatedAt" example:"2025-07-01T12:00:00Z"`
}

// Тело запроса на создание задачи
type CreateTaskRequest struct {
	Title       string `json:"title" example:"Купить хлеб"`
	Description string `json:"description" example:"В ближайшем магазине"`
	Deadline    string `json:"deadline" example:"2025-07-01T12:00:00Z"`
}

// Тело запроса на получение задачи
type GetTaskRequest struct {
	Id			string `form:"id" example:"a81bc81b-dead-4e5d-abff-90865d1e13b1"`
}

// Тело запроса на получение списка задач
type ListTasksRequest struct {}

// Тело запроса на обновление задачи
type UpdateTaskRequestBody struct {
	Status		string `json:"status" example:"TODO"`
}

// Квери параметры запроса на обновление задачи
type UpdateTaskRequestQuery struct {
	Id			string `form:"id" example:"a81bc81b-dead-4e5d-abff-90865d1e13b1"`
}

// Тело запроса на удаление задачи
type DeleteTaskRequest struct {
	Id			string `form:"id" example:"a81bc81b-dead-4e5d-abff-90865d1e13b1"`
}
