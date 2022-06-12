package entity

/*
{
	blog: {
		title: {
			type: string,
			example: "title",
		}
		content: {
			type: string
			example: "Hello, World!"
		}
		created_at: {
			type: datetime,
			example: "2020-01-01T00:00:00+09:00"
		}
		etc...
	}
}
*/

type Propertie struct {
	Type    string `json:"type"`
	Example string `json:"example"`
}
