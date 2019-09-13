type Ocil struct {
	XMLName        xml.Name `xml:"ocil"`
	Text           string   `xml:",chardata"`
	Xsi            string   `xml:"xsi,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	Xmlns          string   `xml:"xmlns,attr"`
	
	Generator      struct {
		Text          string `xml:",chardata"`
		SchemaVersion string `xml:"schema_version"`
		Timestamp     string `xml:"timestamp"`
		Author        struct {
			Text         string `xml:",chardata"`
			Organization string `xml:"organization,attr"`
		} `xml:"author"`
	} `xml:"generator"`

	Document struct {
		Text        string   `xml:",chardata"`
		Title       string   `xml:"title"`
		Description []string `xml:"description"`
		Notice      string   `xml:"notice"`
	} `xml:"document"`

	Questionnaire []struct {
		Text        string `xml:",chardata"`
		ID          string `xml:"id,attr"`
		Priority    string `xml:"priority,attr"`
		Title       string `xml:"title"`
		Description string `xml:"description"`
		Actions     struct {
			Text          string `xml:",chardata"`
			Priority      string `xml:"priority,attr"`
			Operation     string `xml:"operation,attr"`
			TestActionRef []struct {
				Text     string `xml:",chardata"`
				Priority string `xml:"priority,attr"`
			} `xml:"test_action_ref"`
		} `xml:"actions"`
	} `xml:"questionnaire"`

	BooleanQuestionTestAction []struct {
		Text        string `xml:",chardata"`
		ID          string `xml:"id,attr"`
		QuestionRef string `xml:"question_ref,attr"`
		WhenTrue    struct {
			Text   string `xml:",chardata"`
			Result string `xml:"result"`
		} `xml:"when_true"`
		WhenFalse struct {
			Text   string `xml:",chardata"`
			Result string `xml:"result"`
		} `xml:"when_false"`
	} `xml:"boolean_question_test_action"`
	BooleanQuestion []struct {
		Text         string `xml:",chardata"`
		ID           string `xml:"id,attr"`
		Model        string `xml:"model,attr"`
		QuestionText string `xml:"question_text"`
	} `xml:"boolean_question"`
}
