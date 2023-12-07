package entity

type Response struct {
	Body struct {
		Storage struct {
			Value          string `json:"value"`
			Representation string `json:"representation"`
			Expandable     struct {
				Content string `json:"content"`
			} `json:"_expandable"`
		} `json:"storage"`
		Expandable struct {
			Editor              string `json:"editor"`
			View                string `json:"view"`
			ExportView          string `json:"export_view"`
			StyledView          string `json:"styled_view"`
			AnonymousExportView string `json:"anonymous_export_view"`
		} `json:"_expandable"`
	} `json:"body"`
}
