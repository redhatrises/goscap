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
		Text    string `xml:",chardata"`
		ID      string `xml:"id,attr"`
		Title   string `xml:"title"`
		Actions struct {
			Text          string `xml:",chardata"`
			Negate        string `xml:"negate,attr"`
			Operation     string `xml:"operation,attr"`
			TestActionRef []struct {
				Text     string `xml:",chardata"`
				Priority string `xml:"priority,attr"`
				Negate   string `xml:"negate,attr"`
			} `xml:"test_action_ref"`
		} `xml:"actions"`
	} `xml:"questionnaire"`

	BooleanQuestionTestAction []struct {
		Text        string `xml:",chardata"`
		QuestionRef string `xml:"question_ref,attr"`
		ID          string `xml:"id,attr"`
		WhenTrue    struct {
			Text          string `xml:",chardata"`
			Result        string `xml:"result"`
			TestActionRef struct {
				Text   string `xml:",chardata"`
				Negate string `xml:"negate,attr"`
			} `xml:"test_action_ref"`
		} `xml:"when_true"`
		WhenFalse struct {
			Text          string `xml:",chardata"`
			TestActionRef struct {
				Text   string `xml:",chardata"`
				Negate string `xml:"negate,attr"`
			} `xml:"test_action_ref"`
			Result string `xml:"result"`
		} `xml:"when_false"`
	} `xml:"boolean_question_test_action"`
	ChoiceQuestionTestAction []struct {
		Text        string `xml:",chardata"`
		QuestionRef string `xml:"question_ref,attr"`
		ID          string `xml:"id,attr"`
		WhenChoice  []struct {
			Text          string   `xml:",chardata"`
			TestActionRef string   `xml:"test_action_ref"`
			ChoiceRef     []string `xml:"choice_ref"`
			Result        string   `xml:"result"`
		} `xml:"when_choice"`
	} `xml:"choice_question_test_action"`
	NumericQuestionTestAction struct {
		Text        string `xml:",chardata"`
		QuestionRef string `xml:"question_ref,attr"`
		ID          string `xml:"id,attr"`
		WhenEquals  struct {
			Text   string `xml:",chardata"`
			Result string `xml:"result"`
			Value  string `xml:"value"`
		} `xml:"when_equals"`
		WhenRange []struct {
			Text   string `xml:",chardata"`
			Result string `xml:"result"`
			Range  struct {
				Text string `xml:",chardata"`
				Min  string `xml:"min"`
				Max  string `xml:"max"`
			} `xml:"range"`
		} `xml:"when_range"`
	} `xml:"numeric_question_test_action"`
	StringQuestionTestAction struct {
		Text        string `xml:",chardata"`
		QuestionRef string `xml:"question_ref,attr"`
		ID          string `xml:"id,attr"`
		WhenPattern []struct {
			Text    string `xml:",chardata"`
			Result  string `xml:"result"`
			Pattern string `xml:"pattern"`
		} `xml:"when_pattern"`
	} `xml:"string_question_test_action"`
	BooleanQuestion []struct {
		Text         string `xml:",chardata"`
		ID           string `xml:"id,attr"`
		Model        string `xml:"model,attr"`
		QuestionText string `xml:"question_text"`
	} `xml:"boolean_question"`
	ChoiceQuestion []struct {
		Text             string `xml:",chardata"`
		ID               string `xml:"id,attr"`
		DefaultAnswerRef string `xml:"default_answer_ref,attr"`
		QuestionText     string `xml:"question_text"`
		Choice           []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
		} `xml:"choice"`
		ChoiceGroupRef []string `xml:"choice_group_ref"`
	} `xml:"choice_question"`
	NumericQuestion struct {
		Text         string `xml:",chardata"`
		ID           string `xml:"id,attr"`
		QuestionText string `xml:"question_text"`
	} `xml:"numeric_question"`
	StringQuestion struct {
		Text         string `xml:",chardata"`
		ID           string `xml:"id,attr"`
		QuestionText string `xml:"question_text"`
	} `xml:"string_question"`
	ChoiceGroup []struct {
		Text   string `xml:",chardata"`
		ID     string `xml:"id,attr"`
		Choice []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
		} `xml:"choice"`
	} `xml:"choice_group"`
}
