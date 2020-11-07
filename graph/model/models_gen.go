package model

type AddMessage struct {
	Author  string `json:"author"`
	Ticket  string `json:"ticket"`
	Content string `json:"content"`
}

type AddTicket struct {
	Title       string `json:"title"`
	Owner       string `json:"owner"`
	Equipment   string `json:"equipment"`
	Description string `json:"description"`
}

type Login struct {
	ID     string `json:"_id" bson:"_id"`
	Status bool   `json:"status"`
	User   *User  `json:"user"`
}

type Message struct {
	ID        string `json:"_id" bson:"_id"`
	Author    string `json:"author"`
	Ticket    string `json:"ticket"`
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
}

type NewMessage struct {
	Author    string `json:"author"`
	Ticket    string `json:"ticket"`
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
}

type NewTicket struct {
	Title       string `json:"title"`
	Owner       string `json:"owner"`
	Status      string `json:"status"`
	Technical   string `json:"technical"`
	Equipment   string `json:"equipment"`
	Description string `json:"description"`
	Timestamp   string `json:"timestamp"`
}

type NewTv struct {
	Type   string `json:"type"`
	Model  string `json:"model"`
	Brand  string `json:"brand"`
	Serial string `json:"serial"`
	Owner  string `json:"owner"`
}

type NewUser struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type Ticket struct {
	ID          string `json:"_id" bson:"_id"`
	Title       string `json:"title"`
	Status      string `json:"status"`
	Equipment   string `json:"equipment"`
	Technical   string `json:"technical"`
	Timestamp   string `json:"timestamp"`
	Description string `json:"description"`
	Tv          *Tv    `json:"tv"`
}

type Tickets struct {
	ID          string `json:"_id" bson:"_id"`
	Title       string `json:"title"`
	Owner       string `json:"owner"`
	Status      string `json:"status"`
	Technical   string `json:"technical"`
	Equipment   string `json:"equipment"`
	Timestamp   string `json:"timestamp"`
	Description string `json:"description"`
}

type Tv struct {
	ID     string `json:"_id" bson:"_id"`
	Type   string `json:"type"`
	Model  string `json:"model"`
	Brand  string `json:"brand"`
	Serial string `json:"serial"`
	Owner  string `json:"owner"`
	User   *User  `json:"user"`
}

type Tvs struct {
	ID     string `json:"_id" bson:"_id"`
	Type   string `json:"type"`
	Model  string `json:"model"`
	Brand  string `json:"brand"`
	Serial string `json:"serial"`
	Owner  string `json:"owner"`
}

type User struct {
	ID      string `json:"_id" bson:"_id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Role    string `json:"role"`
}

type UserAuth struct {
	ID       string `json:"_id" bson:"_id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Password string `json:"password"`
}
