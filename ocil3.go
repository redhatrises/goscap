package main

import (
	"encoding/xml"
	"io/ioutil"
)


type Generator struct {
                XMLName xml.Name `xml:"generator"`
                // Text           string   `xml:",chardata"`
                ProductName    string `xml:"product_name,omitempty"`
                ProductVersion string `xml:"product_version,omitempty"`
                Author         []struct {
                        Text         string `xml:",chardata"`
                        Organization string `xml:"organization,attr"`
                } `xml:"author,omitempty"`
                SchemaVersion  string `xml:"schema_version"`
                Timestamp      string `xml:"timestamp"`
                AdditionalData string `xml:"additional_data,omitempty"`
        }


type Ocil struct {
	XMLName xml.Name `xml:"ocil"`
	// Text           string   `xml:",chardata"`
	XML            string `xml:"xml,attr"`
	Xmlns          string `xml:"xmlns,attr"`
	Sch            string `xml:"sch,attr"`
	SchemaLocation string `xml:"schemaLocation,attr"`
	Xsi            string `xml:"xsi,attr"`

	Generator struct {
                XMLName xml.Name `xml:"generator"`
                // Text           string   `xml:",chardata"`
                ProductName    string `xml:"product_name,omitempty"`
                ProductVersion string `xml:"product_version,omitempty"`
                Author         []struct {
                        Text         string `xml:",chardata"`
                        Organization string `xml:"organization,attr"`
                } `xml:"author,omitempty"`
                SchemaVersion  string `xml:"schema_version"`
                Timestamp      string `xml:"timestamp"`
                AdditionalData string `xml:"additional_data,omitempty"`
        }
	Document struct {
		Text        string   `xml:",chardata"`
		Title       string   `xml:"title"`
		Description []string `xml:"description,omitempty"`
		Notice      []string `xml:"notice,omitempty"`
	} `xml:"document,omitempty"`

	Questionnaires struct {
		// Text          string `xml:",chardata"`
		Questionnaire struct {
			Text        string `xml:",chardata"`
			ID          string `xml:"id,attr"`
			ChildOnly   string `xml:"child_only,attr,omitempty"`
			Title       string `xml:"title,omitempty"`
			Description string `xml:"description,omitempty"`
			References  string `xml:"references,omitempty"`
			Actions     struct {
				//Text          string `xml:",chardata"`
				TestActionRef []string `xml:"test_action_ref"`
				Operation     string   `xml:"operation,attr,omitempty"`
				Negate        string   `xml:"negate,attr,omitempty"`
			} `xml:"actions"`
			Notes []string `xml:"notes,omitempty"`
		} `xml:"questionnaire"`
	} `xml:"questionnaires"`

	TestActions struct {
		//Text                     string `xml:",chardata"`
		ChoiceQuestionTestAction struct {
			Text        string `xml:",chardata"`
			QuestionRef string `xml:"question_ref,attr"`
			ID          string `xml:"id,attr"`
			WhenChoice  struct {
				Text      string `xml:",chardata"`
				Result    string `xml:"result"`
				ChoiceRef string `xml:"choice_ref"`
			} `xml:"when_choice"`
		} `xml:"choice_question_test_action"`
	} `xml:"test_actions"`
	Questions struct {
		//Text           string `xml:",chardata"`
		ChoiceQuestion []struct {
			Text         string `xml:",chardata"`
			ID           string `xml:"id,attr"`
			QuestionText string `xml:"question_text"`
			Choice       struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id,attr"`
			} `xml:"choice"`
			ChoiceGroupRef []string `xml:"choice_group_ref"`
		} `xml:"choice_question"`
	} `xml:"questions"`
	Artifacts struct {
		Text     string `xml:",chardata"`
		Artifact []struct {
			Text        string `xml:",chardata"`
			ID          string `xml:"id,attr"`
			Title       string `xml:"title"`
			Description string `xml:"description"`
		} `xml:"artifact,omitempty"`
	} `xml:"artifacts"`
	Variables struct {
		Text     string `xml:",chardata"`
		Variable []struct {
			Text     string `xml:",chardata"`
			Type     string `xml:"type,attr"`
			ID       string `xml:"id,attr"`
			Datatype string `xml:"datatype,attr"`
			Value    string `xml:"value"`
		} `xml:"variable"`
		ConstantVariable struct {
			Text     string `xml:",chardata"`
			ID       string `xml:"id,attr"`
			Datatype string `xml:"datatype,attr"`
			Value    string `xml:"value"`
		} `xml:"constant_variable"`
	} `xml:"variables,omitempty"`
	Results struct {
		Text              string `xml:",chardata"`
		StartTime         string `xml:"start_time,attr"`
		EndTime           string `xml:"end_time,attr"`
		Title             string `xml:"title"`
		TestActionResults struct {
			Text             string `xml:",chardata"`
			TestActionResult struct {
				Text          string `xml:",chardata"`
				TestActionRef string `xml:"test_action_ref,attr"`
				Result        string `xml:"result,attr"`
			} `xml:"test_action_result"`
		} `xml:"test_action_results"`
		QuestionResults struct {
			Text                  string `xml:",chardata"`
			BooleanQuestionResult struct {
				Text        string `xml:",chardata"`
				QuestionRef string `xml:"question_ref,attr"`
				Answer      string `xml:"answer"`
			} `xml:"boolean_question_result"`
		} `xml:"question_results"`
		Targets struct {
			Text   string `xml:",chardata"`
			Target struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
				Name string `xml:"name"`
			} `xml:"target"`
		} `xml:"targets"`
	} `xml:"results,omitempty"`
}

func main() {
	note := &Ocil{}
	file, _ := xml.MarshalIndent(note, "", " ")

	_ = ioutil.WriteFile("notes1.xml", file, 0644)

}
