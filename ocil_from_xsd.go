package postal

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

// Must match the pattern ocil:[A-Za-z0-9_\-\.]+:artifact:[1-9][0-9]*
type ArtifactIDPattern string

// The ArtifactRefType type defines a single artifact reference
// that may be collected as part of a questionnaire assessment.
type ArtifactRefType struct {
	Idref    ArtifactIDPattern `xml:"idref,attr,omitempty"`
	Required bool              `xml:"required,attr,omitempty"`
}

func (t *ArtifactRefType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ArtifactRefType
	var overlay struct {
		*T
		Required *bool `xml:"required,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Required = (*bool)(&overlay.T.Required)
	return d.DecodeElement(&overlay, &start)
}

// The ArtifactRefsType type defines a collection of artifact
// references that may be collected as part of a questionnaire
// assessment.
type ArtifactRefsType struct {
	Artifact_ref []ArtifactRefType `xml:"http://scap.nist.gov/schema/ocil/2.0 artifact_ref"`
}

// The ArtifactResultType type defines structures containing
// information about the submitted artifact, its value, who provided and
// submitted it, and when it was submitted.
type ArtifactResultType struct {
	Artifact_value ArtifactValueType    `xml:"http://scap.nist.gov/schema/ocil/2.0 artifact_value"`
	Provider       ProviderValuePattern `xml:"http://scap.nist.gov/schema/ocil/2.0 provider"`
	Submitter      UserType             `xml:"http://scap.nist.gov/schema/ocil/2.0 submitter"`
	Artifact_ref   ArtifactIDPattern    `xml:"artifact_ref,attr"`
	Timestamp      time.Time            `xml:"timestamp,attr"`
}

func (t *ArtifactResultType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ArtifactResultType
	var layout struct {
		*T
		Timestamp *xsdDateTime `xml:"timestamp,attr"`
	}
	layout.T = (*T)(t)
	layout.Timestamp = (*xsdDateTime)(&layout.T.Timestamp)
	return e.EncodeElement(layout, start)
}
func (t *ArtifactResultType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ArtifactResultType
	var overlay struct {
		*T
		Timestamp *xsdDateTime `xml:"timestamp,attr"`
	}
	overlay.T = (*T)(t)
	overlay.Timestamp = (*xsdDateTime)(&overlay.T.Timestamp)
	return d.DecodeElement(&overlay, &start)
}

// The ArtifactResultsType type defines structures
// containing a set of artifact_result elements.
type ArtifactResultsType struct {
	Artifact_result []ArtifactResultType `xml:"http://scap.nist.gov/schema/ocil/2.0 artifact_result"`
}

// The ArtifactType type defines structures containing
// information about an artifact such as title, description, persistence,
// and if it's required to complete an answer to a question.
type ArtifactType struct {
	Title       TextType          `xml:"http://scap.nist.gov/schema/ocil/2.0 title"`
	Description TextType          `xml:"http://scap.nist.gov/schema/ocil/2.0 description"`
	Notes       []string          `xml:"http://scap.nist.gov/schema/ocil/2.0 notes,omitempty"`
	Id          ArtifactIDPattern `xml:"id,attr"`
	Persistent  bool              `xml:"persistent,attr,omitempty"`
	Revision    int               `xml:"revision,attr,omitempty"`
}

func (t *ArtifactType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ArtifactType
	var overlay struct {
		*T
		Persistent *bool `xml:"persistent,attr,omitempty"`
		Revision   *int  `xml:"revision,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Persistent = (*bool)(&overlay.T.Persistent)
	overlay.Revision = (*int)(&overlay.T.Revision)
	return d.DecodeElement(&overlay, &start)
}

// The ArtifactValueType type defines structures containing
// either the artifact data itself or a pointer to it.
type ArtifactValueType struct {
}

// The ArtifactsType type defines structures containing a
// set of artifact elements.
type ArtifactsType struct {
	Artifact []ArtifactType `xml:"http://scap.nist.gov/schema/ocil/2.0 artifact"`
}

// The data model that holds binary data-based artifacts.
type BinaryArtifactValueType struct {
	Data      []byte `xml:"http://scap.nist.gov/schema/ocil/2.0 data"`
	Mime_type string `xml:"mime_type,attr"`
}

func (t *BinaryArtifactValueType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BinaryArtifactValueType
	var layout struct {
		*T
		Data *xsdBase64Binary `xml:"http://scap.nist.gov/schema/ocil/2.0 data"`
	}
	layout.T = (*T)(t)
	layout.Data = (*xsdBase64Binary)(&layout.T.Data)
	return e.EncodeElement(layout, start)
}
func (t *BinaryArtifactValueType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BinaryArtifactValueType
	var overlay struct {
		*T
		Data *xsdBase64Binary `xml:"http://scap.nist.gov/schema/ocil/2.0 data"`
	}
	overlay.T = (*T)(t)
	overlay.Data = (*xsdBase64Binary)(&overlay.T.Data)
	return d.DecodeElement(&overlay, &start)
}

// May be one of MODEL_YES_NO, MODEL_TRUE_FALSE
type BooleanQuestionModelType string

// The BooleanQuestionResultType type defines structures
// containing a reference to a boolean_question, the response, and
// whether the question was successfully posed.
type BooleanQuestionResultType struct {
	Answer       bool              `xml:"http://scap.nist.gov/schema/ocil/2.0 answer,omitempty"`
	Question_ref QuestionIDPattern `xml:"question_ref,attr"`
	Response     UserResponseType  `xml:"response,attr,omitempty"`
}

func (t *BooleanQuestionResultType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BooleanQuestionResultType
	var overlay struct {
		*T
		Response *UserResponseType `xml:"response,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Response = (*UserResponseType)(&overlay.T.Response)
	return d.DecodeElement(&overlay, &start)
}

// The BooleanQuestionTestActionType type defines a
// structure that references a boolean_question and includes handlers for
// TRUE (YES) or FALSE (NO) responses.
type BooleanQuestionTestActionType struct {
	When_true           TestActionConditionType     `xml:"http://scap.nist.gov/schema/ocil/2.0 when_true"`
	When_false          TestActionConditionType     `xml:"http://scap.nist.gov/schema/ocil/2.0 when_false"`
	Title               TextType                    `xml:"http://scap.nist.gov/schema/ocil/2.0 title,omitempty"`
	When_unknown        TestActionConditionType     `xml:"http://scap.nist.gov/schema/ocil/2.0 when_unknown,omitempty"`
	When_not_tested     TestActionConditionType     `xml:"http://scap.nist.gov/schema/ocil/2.0 when_not_tested,omitempty"`
	When_not_applicable TestActionConditionType     `xml:"http://scap.nist.gov/schema/ocil/2.0 when_not_applicable,omitempty"`
	When_error          TestActionConditionType     `xml:"http://scap.nist.gov/schema/ocil/2.0 when_error,omitempty"`
	Notes               []string                    `xml:"http://scap.nist.gov/schema/ocil/2.0 notes,omitempty"`
	Question_ref        QuestionIDPattern           `xml:"question_ref,attr"`
	Id                  QuestionTestActionIDPattern `xml:"id,attr"`
	Revision            int                         `xml:"revision,attr,omitempty"`
}

func (t *BooleanQuestionTestActionType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BooleanQuestionTestActionType
	var overlay struct {
		*T
		Revision *int `xml:"revision,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Revision = (*int)(&overlay.T.Revision)
	return d.DecodeElement(&overlay, &start)
}

// The BooleanQuestionType type defines a question with
// valid responses of either {TRUE, FALSE} or {YES, NO}.
type BooleanQuestionType struct {
	Question_text  []QuestionTextType       `xml:"http://scap.nist.gov/schema/ocil/2.0 question_text"`
	Instructions   InstructionsType         `xml:"http://scap.nist.gov/schema/ocil/2.0 instructions,omitempty"`
	Notes          []string                 `xml:"http://scap.nist.gov/schema/ocil/2.0 notes,omitempty"`
	Default_answer bool                     `xml:"default_answer,attr,omitempty"`
	Model          BooleanQuestionModelType `xml:"model,attr,omitempty"`
	Id             QuestionIDPattern        `xml:"id,attr"`
	Revision       int                      `xml:"revision,attr,omitempty"`
}

func (t *BooleanQuestionType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BooleanQuestionType
	var overlay struct {
		*T
		Model    *BooleanQuestionModelType `xml:"model,attr,omitempty"`
		Revision *int                      `xml:"revision,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Model = (*BooleanQuestionModelType)(&overlay.T.Model)
	overlay.Revision = (*int)(&overlay.T.Revision)
	return d.DecodeElement(&overlay, &start)
}

// The ChoiceAnswerType type defines structures containing a
// choice_ref attribute that identifies the selected choice.
type ChoiceAnswerType struct {
	Choice_ref ChoiceIDPattern `xml:"choice_ref,attr,omitempty"`
}

// Must match the pattern ocil:[A-Za-z0-9_\-\.]+:choicegroup:[1-9][0-9]*
type ChoiceGroupIDPattern string

// The ChoiceGroupType type defines a group of choices that
// may then be reused in multiple choice_question elements. For example, a
// document may include multiple choice_questions with the options of
// "Good", "Fair", or "Poor". By defining these choices in a single
// choice_group, the author would not need to list them out explicitly in
// every choice_question.
type ChoiceGroupType struct {
	Choice []ChoiceType         `xml:"http://scap.nist.gov/schema/ocil/2.0 choice"`
	Id     ChoiceGroupIDPattern `xml:"id,attr"`
}

// Must match the pattern ocil:[A-Za-z0-9_\-\.]+:choice:[1-9][0-9]*
type ChoiceIDPattern string

// The ChoiceQuestionResultType type defines structures
// containing a reference to a choice_question, the response, and
// whether the question was successfully posed.
type ChoiceQuestionResultType struct {
	Answer       ChoiceAnswerType  `xml:"http://scap.nist.gov/schema/ocil/2.0 answer,omitempty"`
	Question_ref QuestionIDPattern `xml:"question_ref,attr"`
	Response     UserResponseType  `xml:"response,attr,omitempty"`
}

func (t *ChoiceQuestionResultType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ChoiceQuestionResultType
	var overlay struct {
		*T
		Response *UserResponseType `xml:"response,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Response = (*UserResponseType)(&overlay.T.Response)
	return d.DecodeElement(&overlay, &start)
}

// The ChoiceQuestionTestActionType type defines a structure
// that references a choice_question and includes handlers for the various
// choices set out in the choice_question.
type ChoiceQuestionTestActionType struct {
	When_choice         []ChoiceTestActionConditionType `xml:"http://scap.nist.gov/schema/ocil/2.0 when_choice"`
	Title               TextType                        `xml:"http://scap.nist.gov/schema/ocil/2.0 title,omitempty"`
	When_unknown        TestActionConditionType         `xml:"http://scap.nist.gov/schema/ocil/2.0 when_unknown,omitempty"`
	When_not_tested     TestActionConditionType         `xml:"http://scap.nist.gov/schema/ocil/2.0 when_not_tested,omitempty"`
	When_not_applicable TestActionConditionType         `xml:"http://scap.nist.gov/schema/ocil/2.0 when_not_applicable,omitempty"`
	When_error          TestActionConditionType         `xml:"http://scap.nist.gov/schema/ocil/2.0 when_error,omitempty"`
	Notes               []string                        `xml:"http://scap.nist.gov/schema/ocil/2.0 notes,omitempty"`
	Question_ref        QuestionIDPattern               `xml:"question_ref,attr"`
	Id                  QuestionTestActionIDPattern     `xml:"id,attr"`
	Revision            int                             `xml:"revision,attr,omitempty"`
}

func (t *ChoiceQuestionTestActionType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ChoiceQuestionTestActionType
	var overlay struct {
		*T
		Revision *int `xml:"revision,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Revision = (*int)(&overlay.T.Revision)
	return d.DecodeElement(&overlay, &start)
}

// The ChoiceQuestionType type defines a question with one
// or more acceptable answers specified by the author. The response will
// be one of these specified answers. Acceptable answers are specified
// either explicitly using the choice element or implicitly using the
// choice_group_ref element to reference a choice_group element. Choices
// are presented in the order in which they are provided. All the choices in
// a choice_group are inserted in the order in which they appear within the
// choice_group.
type ChoiceQuestionType struct {
	Choice             ChoiceType           `xml:"http://scap.nist.gov/schema/ocil/2.0 choice"`
	Choice_group_ref   ChoiceGroupIDPattern `xml:"http://scap.nist.gov/schema/ocil/2.0 choice_group_ref"`
	Question_text      []QuestionTextType   `xml:"http://scap.nist.gov/schema/ocil/2.0 question_text"`
	Instructions       InstructionsType     `xml:"http://scap.nist.gov/schema/ocil/2.0 instructions,omitempty"`
	Notes              []string             `xml:"http://scap.nist.gov/schema/ocil/2.0 notes,omitempty"`
	Default_answer_ref ChoiceIDPattern      `xml:"default_answer_ref,attr,omitempty"`
	Id                 QuestionIDPattern    `xml:"id,attr"`
	Revision           int                  `xml:"revision,attr,omitempty"`
}

func (t *ChoiceQuestionType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ChoiceQuestionType
	var overlay struct {
		*T
		Revision *int `xml:"revision,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Revision = (*int)(&overlay.T.Revision)
	return d.DecodeElement(&overlay, &start)
}

// The ChoiceTestActionConditionType type defines a structure that
// specifies the action to take in a choice_test_action when a particular
// choice is selected in response to a choice_question.
type ChoiceTestActionConditionType struct {
	Choice_ref      []ChoiceIDPattern `xml:"http://scap.nist.gov/schema/ocil/2.0 choice_ref"`
	Result          ResultType        `xml:"http://scap.nist.gov/schema/ocil/2.0 result"`
	Test_action_ref TestActionRefType `xml:"http://scap.nist.gov/schema/ocil/2.0 test_action_ref"`
	Artifact_refs   ArtifactRefsType  `xml:"http://scap.nist.gov/schema/ocil/2.0 artifact_refs,omitempty"`
}

// The ChoiceType type defines structures that hold
// information about one acceptable answer to a choice_question.
type ChoiceType struct {
	Value   string            `xml:",chardata"`
	Id      ChoiceIDPattern   `xml:"id,attr"`
	Var_ref VariableIDPattern `xml:"var_ref,attr,omitempty"`
}

// The CompoundTestActionType type describes the structures
// used to combine multiple test_action elements into a single
// result.
type CompoundTestActionType struct {
	Title       TextType       `xml:"http://scap.nist.gov/schema/ocil/2.0 title,omitempty"`
	Description TextType       `xml:"http://scap.nist.gov/schema/ocil/2.0 description,omitempty"`
	References  ReferencesType `xml:"http://scap.nist.gov/schema/ocil/2.0 references,omitempty"`
	Actions     OperationType  `xml:"http://scap.nist.gov/schema/ocil/2.0 actions"`
	Notes       []string       `xml:"http://scap.nist.gov/schema/ocil/2.0 notes,omitempty"`
	Revision    int            `xml:"revision,attr,omitempty"`
}

func (t *CompoundTestActionType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CompoundTestActionType
	var overlay struct {
		*T
		Revision *int `xml:"revision,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Revision = (*int)(&overlay.T.Revision)
	return d.DecodeElement(&overlay, &start)
}

// The ConstantVariableType type defines structures
// containing a value defined by the author of the
// document.
type ConstantVariableType struct {
	Value       string            `xml:"http://scap.nist.gov/schema/ocil/2.0 value"`
	Description TextType          `xml:"http://scap.nist.gov/schema/ocil/2.0 description,omitempty"`
	Notes       []string          `xml:"http://scap.nist.gov/schema/ocil/2.0 notes,omitempty"`
	Id          VariableIDPattern `xml:"id,attr"`
	Datatype    VariableDataType  `xml:"datatype,attr"`
	Revision    int               `xml:"revision,attr,omitempty"`
}

func (t *ConstantVariableType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ConstantVariableType
	var overlay struct {
		*T
		Revision *int `xml:"revision,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Revision = (*int)(&overlay.T.Revision)
	return d.DecodeElement(&overlay, &start)
}

// The DocumentType type describes structures used to provide
// document-level information, including title, descriptions, and notices.
type DocumentType struct {
	Title       string   `xml:"http://scap.nist.gov/schema/ocil/2.0 title"`
	Description []string `xml:"http://scap.nist.gov/schema/ocil/2.0 description,omitempty"`
	Notice      []string `xml:"http://scap.nist.gov/schema/ocil/2.0 notice,omitempty"`
}

// The base data structure that holds artifact values that
// are embedded into the results model.
type EmbeddedArtifactValueType struct {
	Mime_type string `xml:"mime_type,attr"`
}

// The EqualsTestActionConditionType defines a structure that specifies the
// action to take in a numeric_test_action when a particular value is given
// in response to a numeric_question.
type EqualsTestActionConditionType struct {
	Value           []float64         `xml:"http://scap.nist.gov/schema/ocil/2.0 value"`
	Result          ResultType        `xml:"http://scap.nist.gov/schema/ocil/2.0 result"`
	Test_action_ref TestActionRefType `xml:"http://scap.nist.gov/schema/ocil/2.0 test_action_ref"`
	Artifact_refs   ArtifactRefsType  `xml:"http://scap.nist.gov/schema/ocil/2.0 artifact_refs,omitempty"`
	Var_ref         VariableIDPattern `xml:"var_ref,attr,omitempty"`
}

// May be one of UNKNOWN, ERROR, NOT_TESTED, NOT_APPLICABLE
type ExceptionalResultType string

type ExtensionContainerType struct {
	Items []string `xml:",any"`
}

// The ExternalVariableType type defines structures
// containing a value defined elsewhere or some external
// source.
type ExternalVariableType struct {
	Description TextType          `xml:"http://scap.nist.gov/schema/ocil/2.0 description,omitempty"`
	Notes       []string          `xml:"http://scap.nist.gov/schema/ocil/2.0 notes,omitempty"`
	Id          VariableIDPattern `xml:"id,attr"`
	Datatype    VariableDataType  `xml:"datatype,attr"`
	Revision    int               `xml:"revision,attr,omitempty"`
}

func (t *ExternalVariableType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ExternalVariableType
	var overlay struct {
		*T
		Revision *int `xml:"revision,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Revision = (*int)(&overlay.T.Revision)
	return d.DecodeElement(&overlay, &start)
}

// The GeneratorType type defines an element that is used
// to hold information about when a particular OCIL document was generated,
// what version of the schema was used, what tool was used to generate the
// document, and what version of the tool was used.
//
// Additional generator information is also allowed although
// it is not part of the official OCIL language. Individual organizations
// can place generator information that they feel is important.
type GeneratorType struct {
	Product_name    string                 `xml:"http://scap.nist.gov/schema/ocil/2.0 product_name,omitempty"`
	Product_version string                 `xml:"http://scap.nist.gov/schema/ocil/2.0 product_version,omitempty"`
	Author          []UserType             `xml:"http://scap.nist.gov/schema/ocil/2.0 author,omitempty"`
	Schema_version  float64                `xml:"http://scap.nist.gov/schema/ocil/2.0 schema_version"`
	Timestamp       time.Time              `xml:"http://scap.nist.gov/schema/ocil/2.0 timestamp"`
	Additional_data ExtensionContainerType `xml:"http://scap.nist.gov/schema/ocil/2.0 additional_data,omitempty"`
}

func (t *GeneratorType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T GeneratorType
	var layout struct {
		*T
		Timestamp *xsdDateTime `xml:"http://scap.nist.gov/schema/ocil/2.0 timestamp"`
	}
	layout.T = (*T)(t)
	layout.Timestamp = (*xsdDateTime)(&layout.T.Timestamp)
	return e.EncodeElement(layout, start)
}
func (t *GeneratorType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T GeneratorType
	var overlay struct {
		*T
		Timestamp *xsdDateTime `xml:"http://scap.nist.gov/schema/ocil/2.0 timestamp"`
	}
	overlay.T = (*T)(t)
	overlay.Timestamp = (*xsdDateTime)(&overlay.T.Timestamp)
	return d.DecodeElement(&overlay, &start)
}

// The InstructionsType type defines a series of steps
// intended to guide the user in answering a question.
type InstructionsType struct {
	Title TextType   `xml:"http://scap.nist.gov/schema/ocil/2.0 title"`
	Step  []StepType `xml:"http://scap.nist.gov/schema/ocil/2.0 step"`
}

// The ItemBaseType complex type defines structures allowing
// a set of notes to be included. This type is inherited by many of the
// elements in the OCIL language.
type ItemBaseType struct {
	Notes    []string `xml:"http://scap.nist.gov/schema/ocil/2.0 notes,omitempty"`
	Revision int      `xml:"revision,attr,omitempty"`
}

func (t *ItemBaseType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ItemBaseType
	var overlay struct {
		*T
		Revision *int `xml:"revision,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Revision = (*int)(&overlay.T.Revision)
	return d.DecodeElement(&overlay, &start)
}

type Lang string

// The LocalVariableType type defines structures containing
// a value determined during evaluation. The value is determined based on
// the answer to the linked question. If one or more set elements are present, the
// value is computed based on those set elements.  The value stored in the first
// set element that produces a match pattern is used. If none of the set elements
// have a pattern that matches the response, then an error result is generated.
//
// If no set element is provided, the value used will be the same as the
// answer, with a few exceptions. The mappings are listed below.
// 1) If the question is a boolean question and the variable data type is
// NUMERIC, then the value based on the answer must be 1 for true and 0 for
// false.
// 2) If the question is a boolean question and the variable data type is TEXT,
// then the value is determined by the question's model as follows:
// a) MODEL_YES_NO: the value must be yes if true or no if false.
// b) MODEL_TRUE_FALSE: the value must be true if true or false if false.
// 3) If the question is a choice question, the variable data type must be TEXT
// and the value must be set to the text value of the choice.
// 4) If the question is a numeric question, the variable data type must be NUMERIC
// and the value must be set to the value of the answer.
// 5) If the question is a string question, the variable data type must be TEXT
// and the value must be set to the value of the answer.
//
// If a local variable is referenced and the value cannot be
// determined, then the referencing question or test action should cause an ERROR
// result to be generated by all referencing test actions.
type LocalVariableType struct {
	Set          string            `xml:"http://scap.nist.gov/schema/ocil/2.0 set,omitempty"`
	Description  TextType          `xml:"http://scap.nist.gov/schema/ocil/2.0 description,omitempty"`
	Notes        []string          `xml:"http://scap.nist.gov/schema/ocil/2.0 notes,omitempty"`
	Question_ref QuestionIDPattern `xml:"question_ref,attr"`
	Id           VariableIDPattern `xml:"id,attr"`
	Datatype     VariableDataType  `xml:"datatype,attr"`
	Revision     int               `xml:"revision,attr,omitempty"`
}

func (t *LocalVariableType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T LocalVariableType
	var overlay struct {
		*T
		Revision *int `xml:"revision,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Revision = (*int)(&overlay.T.Revision)
	return d.DecodeElement(&overlay, &start)
}

// The NamedItemBaseType complex type defines structures
// allowing a set of notes and the name of a target (system or user) to be
// included.
type NamedItemBaseType struct {
	Name     string   `xml:"http://scap.nist.gov/schema/ocil/2.0 name"`
	Notes    []string `xml:"http://scap.nist.gov/schema/ocil/2.0 notes,omitempty"`
	Revision int      `xml:"revision,attr,omitempty"`
}

func (t *NamedItemBaseType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T NamedItemBaseType
	var overlay struct {
		*T
		Revision *int `xml:"revision,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Revision = (*int)(&overlay.T.Revision)
	return d.DecodeElement(&overlay, &start)
}

// The NumericQuestionResultType type defines structures
// containing a reference to a numeric_question, the provided response,
// and whether the question was successfully posed.
type NumericQuestionResultType struct {
	Answer       float64           `xml:"http://scap.nist.gov/schema/ocil/2.0 answer,omitempty"`
	Question_ref QuestionIDPattern `xml:"question_ref,attr"`
	Response     UserResponseType  `xml:"response,attr,omitempty"`
}

func (t *NumericQuestionResultType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T NumericQuestionResultType
	var overlay struct {
		*T
		Response *UserResponseType `xml:"response,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Response = (*UserResponseType)(&overlay.T.Response)
	return d.DecodeElement(&overlay, &start)
}

// The NumericQuestionTestActionType type defines a
// structure that references a numeric_question and includes handlers that
// indicate actions to perform based on whether the response matches
// a particular value or falls within a particular range.
type NumericQuestionTestActionType struct {
	When_equals         []EqualsTestActionConditionType `xml:"http://scap.nist.gov/schema/ocil/2.0 when_equals"`
	When_range          []RangeTestActionConditionType  `xml:"http://scap.nist.gov/schema/ocil/2.0 when_range,omitempty"`
	Title               TextType                        `xml:"http://scap.nist.gov/schema/ocil/2.0 title,omitempty"`
	When_unknown        TestActionConditionType         `xml:"http://scap.nist.gov/schema/ocil/2.0 when_unknown,omitempty"`
	When_not_tested     TestActionConditionType         `xml:"http://scap.nist.gov/schema/ocil/2.0 when_not_tested,omitempty"`
	When_not_applicable TestActionConditionType         `xml:"http://scap.nist.gov/schema/ocil/2.0 when_not_applicable,omitempty"`
	When_error          TestActionConditionType         `xml:"http://scap.nist.gov/schema/ocil/2.0 when_error,omitempty"`
	Notes               []string                        `xml:"http://scap.nist.gov/schema/ocil/2.0 notes,omitempty"`
	Question_ref        QuestionIDPattern               `xml:"question_ref,attr"`
	Id                  QuestionTestActionIDPattern     `xml:"id,attr"`
	Revision            int                             `xml:"revision,attr,omitempty"`
}

func (t *NumericQuestionTestActionType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T NumericQuestionTestActionType
	var overlay struct {
		*T
		Revision *int `xml:"revision,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Revision = (*int)(&overlay.T.Revision)
	return d.DecodeElement(&overlay, &start)
}

// The NumericQuestionType type defines a question that
// requires a numeric answer. Acceptable values may be positive or negative
// and may include decimals.
type NumericQuestionType struct {
	Question_text  []QuestionTextType `xml:"http://scap.nist.gov/schema/ocil/2.0 question_text"`
	Instructions   InstructionsType   `xml:"http://scap.nist.gov/schema/ocil/2.0 instructions,omitempty"`
	Notes          []string           `xml:"http://scap.nist.gov/schema/ocil/2.0 notes,omitempty"`
	Default_answer float64            `xml:"default_answer,attr,omitempty"`
	Id             QuestionIDPattern  `xml:"id,attr"`
	Revision       int                `xml:"revision,attr,omitempty"`
}

func (t *NumericQuestionType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T NumericQuestionType
	var overlay struct {
		*T
		Revision *int `xml:"revision,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Revision = (*int)(&overlay.T.Revision)
	return d.DecodeElement(&overlay, &start)
}

// The OCILType represents the primary content model for the OCIL.
type OCILType struct {
	Generator      GeneratorType      `xml:"http://scap.nist.gov/schema/ocil/2.0 generator"`
	Document       DocumentType       `xml:"http://scap.nist.gov/schema/ocil/2.0 document,omitempty"`
	Questionnaires QuestionnairesType `xml:"http://scap.nist.gov/schema/ocil/2.0 questionnaires"`
	Test_actions   TestActionsType    `xml:"http://scap.nist.gov/schema/ocil/2.0 test_actions"`
	Questions      QuestionsType      `xml:"http://scap.nist.gov/schema/ocil/2.0 questions"`
	Artifacts      ArtifactsType      `xml:"http://scap.nist.gov/schema/ocil/2.0 artifacts,omitempty"`
	Variables      VariablesType      `xml:"http://scap.nist.gov/schema/ocil/2.0 variables,omitempty"`
	Results        ResultsType        `xml:"http://scap.nist.gov/schema/ocil/2.0 results,omitempty"`
}

// The OperationType type defines structures that hold a
// set of test_actions and provide instructions as to how to aggregate their
// individual results into a single result.
type OperationType struct {
	Test_action_ref []TestActionRefType `xml:"http://scap.nist.gov/schema/ocil/2.0 test_action_ref"`
	Operation       OperatorType        `xml:"operation,attr,omitempty"`
	Negate          bool                `xml:"negate,attr,omitempty"`
}

func (t *OperationType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T OperationType
	var overlay struct {
		*T
		Operation *OperatorType `xml:"operation,attr,omitempty"`
		Negate    *bool         `xml:"negate,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Operation = (*OperatorType)(&overlay.T.Operation)
	overlay.Negate = (*bool)(&overlay.T.Negate)
	return d.DecodeElement(&overlay, &start)
}

// May be one of AND, OR
type OperatorType string

// The PatternTestActionConditionType type defines a structure that
// specifies the action to take in a string_test_action when a string given
// in response to a string_question matches the given regular
// expression.
type PatternTestActionConditionType struct {
	Pattern         []PatternType     `xml:"http://scap.nist.gov/schema/ocil/2.0 pattern"`
	Result          ResultType        `xml:"http://scap.nist.gov/schema/ocil/2.0 result"`
	Test_action_ref TestActionRefType `xml:"http://scap.nist.gov/schema/ocil/2.0 test_action_ref"`
	Artifact_refs   ArtifactRefsType  `xml:"http://scap.nist.gov/schema/ocil/2.0 artifact_refs,omitempty"`
}

// The PatternType type defines a structure that specifies a
// regular expression against which a string will be compared.
type PatternType struct {
	Value   string            `xml:",chardata"`
	Var_ref VariableIDPattern `xml:"var_ref,attr,omitempty"`
}

// Must match the pattern ocil:[A-Za-z0-9_\-\.]+:user:[1-9][0-9]*|ocil:[A-Za-z0-9_\-\.]+:system:[1-9][0-9]*
type ProviderValuePattern string

// Must match the pattern ocil:[A-Za-z0-9_\-\.]+:question:[1-9][0-9]*
type QuestionIDPattern string

// The QuestionResultType complex type defines structures
// that hold information about a question and the response to
// it.
type QuestionResultType struct {
	Question_ref QuestionIDPattern `xml:"question_ref,attr"`
	Response     UserResponseType  `xml:"response,attr,omitempty"`
}

func (t *QuestionResultType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T QuestionResultType
	var overlay struct {
		*T
		Response *UserResponseType `xml:"response,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Response = (*UserResponseType)(&overlay.T.Response)
	return d.DecodeElement(&overlay, &start)
}

// The QuestionResultsType type defines structures
// containing computed results of all evaluated question
// types.
type QuestionResultsType struct {
	Question_result []QuestionResultType `xml:"http://scap.nist.gov/schema/ocil/2.0 question_result"`
}

// Must match the pattern ocil:[A-Za-z0-9_\-\.]+:testaction:[1-9][0-9]*
type QuestionTestActionIDPattern string

// The QuestionTestActionType type defines structures that
// are used to hold handlers for non-standard results (UNKNOWN, NOT_TESTED,
// NOT_APPLICABLE, and ERROR) received from a referenced question. All
// children of question_test_action extend this type.
type QuestionTestActionType struct {
	Title               TextType                    `xml:"http://scap.nist.gov/schema/ocil/2.0 title,omitempty"`
	When_unknown        TestActionConditionType     `xml:"http://scap.nist.gov/schema/ocil/2.0 when_unknown,omitempty"`
	When_not_tested     TestActionConditionType     `xml:"http://scap.nist.gov/schema/ocil/2.0 when_not_tested,omitempty"`
	When_not_applicable TestActionConditionType     `xml:"http://scap.nist.gov/schema/ocil/2.0 when_not_applicable,omitempty"`
	When_error          TestActionConditionType     `xml:"http://scap.nist.gov/schema/ocil/2.0 when_error,omitempty"`
	Notes               []string                    `xml:"http://scap.nist.gov/schema/ocil/2.0 notes,omitempty"`
	Question_ref        QuestionIDPattern           `xml:"question_ref,attr"`
	Id                  QuestionTestActionIDPattern `xml:"id,attr"`
	Revision            int                         `xml:"revision,attr,omitempty"`
}

func (t *QuestionTestActionType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T QuestionTestActionType
	var overlay struct {
		*T
		Revision *int `xml:"revision,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Revision = (*int)(&overlay.T.Revision)
	return d.DecodeElement(&overlay, &start)
}

// The QuestionTextType complex type defines a structure
// to hold the text and variables that comprise a question's text.
type QuestionTextType struct {
	Sub []SubstitutionTextType `xml:"http://scap.nist.gov/schema/ocil/2.0 sub,omitempty"`
}

// The QuestionType complex type defines a structure to
// describe a question and any instructions to help in determining an
// answer.
type QuestionType struct {
	Question_text []QuestionTextType `xml:"http://scap.nist.gov/schema/ocil/2.0 question_text"`
	Instructions  InstructionsType   `xml:"http://scap.nist.gov/schema/ocil/2.0 instructions,omitempty"`
	Notes         []string           `xml:"http://scap.nist.gov/schema/ocil/2.0 notes,omitempty"`
	Id            QuestionIDPattern  `xml:"id,attr"`
	Revision      int                `xml:"revision,attr,omitempty"`
}

func (t *QuestionType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T QuestionType
	var overlay struct {
		*T
		Revision *int `xml:"revision,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Revision = (*int)(&overlay.T.Revision)
	return d.DecodeElement(&overlay, &start)
}

// Must match the pattern ocil:[A-Za-z0-9_\-\.]+:questionnaire:[1-9][0-9]*
type QuestionnaireIDPattern string

// The QuestionnaireResultType type defines structures
// containing the computed result, associated artifacts and targets of a
// particular questionnaire.
type QuestionnaireResultType struct {
	Artifact_results  ArtifactResultsType    `xml:"http://scap.nist.gov/schema/ocil/2.0 artifact_results,omitempty"`
	Questionnaire_ref QuestionnaireIDPattern `xml:"questionnaire_ref,attr"`
	Result            ResultType             `xml:"result,attr"`
}

// The QuestionnaireResultsType type defines structures
// containing computed results of all the evaluated
// questionnaires.
type QuestionnaireResultsType struct {
	Questionnaire_result []QuestionnaireResultType `xml:"http://scap.nist.gov/schema/ocil/2.0 questionnaire_result"`
}

// The QuestionnaireType type defines a structure that
// represents a specific question or set of questions that evaluate to a
// single result. A questionnaire may contain multiple test_actions.
// test_actions may be nested and aggregated through an acceptable
// operation to produce the result of a check.
type QuestionnaireType struct {
	Title       TextType               `xml:"http://scap.nist.gov/schema/ocil/2.0 title,omitempty"`
	Description TextType               `xml:"http://scap.nist.gov/schema/ocil/2.0 description,omitempty"`
	References  ReferencesType         `xml:"http://scap.nist.gov/schema/ocil/2.0 references,omitempty"`
	Actions     OperationType          `xml:"http://scap.nist.gov/schema/ocil/2.0 actions"`
	Notes       []string               `xml:"http://scap.nist.gov/schema/ocil/2.0 notes,omitempty"`
	Id          QuestionnaireIDPattern `xml:"id,attr"`
	Child_only  bool                   `xml:"child_only,attr,omitempty"`
	Revision    int                    `xml:"revision,attr,omitempty"`
}

func (t *QuestionnaireType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T QuestionnaireType
	var overlay struct {
		*T
		Child_only *bool `xml:"child_only,attr,omitempty"`
		Revision   *int  `xml:"revision,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Child_only = (*bool)(&overlay.T.Child_only)
	overlay.Revision = (*int)(&overlay.T.Revision)
	return d.DecodeElement(&overlay, &start)
}

// The QuestionnairesType type defines a container for a set
// of questionnaire elements.
type QuestionnairesType struct {
	Questionnaire []QuestionnaireType `xml:"http://scap.nist.gov/schema/ocil/2.0 questionnaire"`
}

// The QuestionsType type defines structures containing a
// set of QuestionType and ChoiceGroupType elements.
type QuestionsType struct {
	Question     []QuestionType    `xml:"http://scap.nist.gov/schema/ocil/2.0 question"`
	Choice_group []ChoiceGroupType `xml:"http://scap.nist.gov/schema/ocil/2.0 choice_group,omitempty"`
}

// The RangeTestActionConditionType type defines a structure that specifies
// the action to take in a numeric_test_action when a value given
// in response to a numeric_question falls within the indicated range.
type RangeTestActionConditionType struct {
	Range           []RangeType       `xml:"http://scap.nist.gov/schema/ocil/2.0 range"`
	Result          ResultType        `xml:"http://scap.nist.gov/schema/ocil/2.0 result"`
	Test_action_ref TestActionRefType `xml:"http://scap.nist.gov/schema/ocil/2.0 test_action_ref"`
	Artifact_refs   ArtifactRefsType  `xml:"http://scap.nist.gov/schema/ocil/2.0 artifact_refs,omitempty"`
}

// The RangeType type defines a structure that specifies a
// range against which a numeric response is to be compared.
type RangeType struct {
	Min RangeValueType `xml:"http://scap.nist.gov/schema/ocil/2.0 min,omitempty"`
	Max RangeValueType `xml:"http://scap.nist.gov/schema/ocil/2.0 max,omitempty"`
}

// Defines a specific bound in a range.
type RangeValueType struct {
	Value     float64           `xml:",chardata"`
	Inclusive bool              `xml:"inclusive,attr,omitempty"`
	Var_ref   VariableIDPattern `xml:"var_ref,attr,omitempty"`
}

func (t *RangeValueType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T RangeValueType
	var overlay struct {
		*T
		Inclusive *bool `xml:"inclusive,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Inclusive = (*bool)(&overlay.T.Inclusive)
	return d.DecodeElement(&overlay, &start)
}

type Reference struct {
	Href string `xml:"href,attr,omitempty"`
}

// The data model that references external artifacts.
type ReferenceArtifactValueType struct {
	Reference Reference `xml:"http://scap.nist.gov/schema/ocil/2.0 reference"`
}

// The ReferenceType complex type defines structures used to
// hold information about an external reference given its URI and
// description.
//
// This structure may be used to reference other standards
// such as CVE, CCE, or CPE. To do so, the href attribute would give the
// relevant namespace. For example, the namespace of the current version of
// CPE is http://cpe.mitre.org/dictionary/2.0 and the body of this element
// would hold a specific CPE identifier. References to other information
// (documents, web pages, etc.) are also permitted.
type ReferenceType struct {
	Value string `xml:",chardata"`
	Href  string `xml:"href,attr,omitempty"`
	Lang  Lang   `xml:"lang,attr,omitempty"`
}

// The ReferencesType complex type contains a set of
// references.
type ReferencesType struct {
	Reference []ReferenceType `xml:"http://scap.nist.gov/schema/ocil/2.0 reference"`
}

// The ResultType simple type defines acceptable result
// values for questionnaires and test_actions.
//
// || P   | F | E | U | NT | NA ||
// ---------------||-----------------------------||------------------||------------------------------------------
// || 1+ | 0   | 0   | 0   | 0   | 0+ || Pass
// || 0+ | 1+ | 0+ | 0+ | 0+ | 0+ || Fail
// AND 	|| 0+ | 0   | 1+ | 0+ | 0+ | 0+ || Error
// || 0+ | 0   | 0   | 1+ | 0+ | 0+ || Unknown
// || 0+ | 0   | 0   | 0   | 1+ | 0+ || Not Tested
// || 0   | 0   | 0   | 0   | 0   | 1+ || Not Applicable
// || 0   | 0   | 0   | 0   | 0   | 0   || Not Tested
// ---------------||-----------------------------||------------------||------------------------------------------
// || 1+ | 0+ | 0+ | 0+ | 0+ | 0+ || Pass
// || 0   | 1+ | 0   | 0   | 0   | 0+ || Fail
// OR 	|| 0   | 0+ | 1+ | 0+ | 0+ | 0+ || Error
// || 0   | 0+ | 0   | 1+ | 0+ | 0+ || Unknown
// || 0   | 0+ | 0   | 0   | 1+ | 0+ || Not Tested
// || 0   | 0   | 0   | 0   | 0   | 1+ || Not Applicable
// || 0   | 0   | 0   | 0   | 0   | 0   || Not Tested
type ResultType string

// The ResultsType type defines structures containing
// results from questionnaires, test actions, questions, artifacts, and
// metadata about the start/end time of evaluation, any targets, and a short
// caption or title.
type ResultsType struct {
	Title                 TextType                 `xml:"http://scap.nist.gov/schema/ocil/2.0 title,omitempty"`
	Questionnaire_results QuestionnaireResultsType `xml:"http://scap.nist.gov/schema/ocil/2.0 questionnaire_results,omitempty"`
	Test_action_results   TestActionResultsType    `xml:"http://scap.nist.gov/schema/ocil/2.0 test_action_results,omitempty"`
	Question_results      QuestionResultsType      `xml:"http://scap.nist.gov/schema/ocil/2.0 question_results,omitempty"`
	Artifact_results      ArtifactResultsType      `xml:"http://scap.nist.gov/schema/ocil/2.0 artifact_results,omitempty"`
	Targets               TargetsType              `xml:"http://scap.nist.gov/schema/ocil/2.0 targets,omitempty"`
	Start_time            time.Time                `xml:"start_time,attr,omitempty"`
	End_time              time.Time                `xml:"end_time,attr,omitempty"`
}

func (t *ResultsType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ResultsType
	var layout struct {
		*T
		Start_time *xsdDateTime `xml:"start_time,attr,omitempty"`
		End_time   *xsdDateTime `xml:"end_time,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Start_time = (*xsdDateTime)(&layout.T.Start_time)
	layout.End_time = (*xsdDateTime)(&layout.T.End_time)
	return e.EncodeElement(layout, start)
}
func (t *ResultsType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ResultsType
	var overlay struct {
		*T
		Start_time *xsdDateTime `xml:"start_time,attr,omitempty"`
		End_time   *xsdDateTime `xml:"end_time,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Start_time = (*xsdDateTime)(&overlay.T.Start_time)
	overlay.End_time = (*xsdDateTime)(&overlay.T.End_time)
	return d.DecodeElement(&overlay, &start)
}

// The SetExpressionBaseType type is the base type of all set
// expressions.  It defines the value to use if the expression evaluates to
// TRUE.
type SetExpressionBaseType struct {
	Value string `xml:"http://scap.nist.gov/schema/ocil/2.0 value"`
}

// The SetExpressionBooleanType type defines criteria for
// evaluating the result of a question answer based on the boolean value entered.  If the
// answer matches the expression must evaluate to TRUE.
type SetExpressionBooleanType struct {
	Value     string `xml:"http://scap.nist.gov/schema/ocil/2.0 value"`
	ValueAttr bool   `xml:"value,attr"`
}

// The SetExpressionChoiceType type defines criteria for
// evaluating the result of a question answer based on the choice made.  If the
// referenced choice is selected the expression must evaluate to TRUE.
type SetExpressionChoiceType struct {
	Value      string          `xml:"http://scap.nist.gov/schema/ocil/2.0 value"`
	Choice_ref ChoiceIDPattern `xml:"choice_ref,attr"`
}

// The SetExpressionPatternType type defines criteria for
// evaluating the result of a question answer as a string.  If the pattern
// matches, the expression must evaluate to TRUE.
type SetExpressionPatternType struct {
	Value   string `xml:"http://scap.nist.gov/schema/ocil/2.0 value"`
	Pattern string `xml:"pattern,attr"`
}

// The SetExpressionRangeType type defines criteria for
// evaluating the result of a question answer based on the decimal value entered.  If the
// answer is within the range (inclusive) the expression must evaluate to TRUE.
type SetExpressionRangeType struct {
	Value string  `xml:"http://scap.nist.gov/schema/ocil/2.0 value"`
	Min   float64 `xml:"min,attr"`
	Max   float64 `xml:"max,attr"`
}

// The StepType complex type defines structures that
// describe one step (out of possibly multiple steps) that a user should
// take to respond to a question. The steps would appear as part of
// the question's instructions element.
type StepType struct {
	Description TextType        `xml:"http://scap.nist.gov/schema/ocil/2.0 description,omitempty"`
	Reference   []ReferenceType `xml:"http://scap.nist.gov/schema/ocil/2.0 reference,omitempty"`
	Step        []StepType      `xml:"http://scap.nist.gov/schema/ocil/2.0 step,omitempty"`
	Is_done     bool            `xml:"is_done,attr,omitempty"`
	Is_required bool            `xml:"is_required,attr,omitempty"`
}

func (t *StepType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T StepType
	var overlay struct {
		*T
		Is_done     *bool `xml:"is_done,attr,omitempty"`
		Is_required *bool `xml:"is_required,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Is_done = (*bool)(&overlay.T.Is_done)
	overlay.Is_required = (*bool)(&overlay.T.Is_required)
	return d.DecodeElement(&overlay, &start)
}

// The StringQuestionResultType type defines structures
// containing a reference to a string_question, the string provided in
// response, and whether the question was successfully
// posed.
type StringQuestionResultType struct {
	Answer       string            `xml:"http://scap.nist.gov/schema/ocil/2.0 answer,omitempty"`
	Question_ref QuestionIDPattern `xml:"question_ref,attr"`
	Response     UserResponseType  `xml:"response,attr,omitempty"`
}

func (t *StringQuestionResultType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T StringQuestionResultType
	var overlay struct {
		*T
		Response *UserResponseType `xml:"response,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Response = (*UserResponseType)(&overlay.T.Response)
	return d.DecodeElement(&overlay, &start)
}

// The StringQuestionTestActionType type defines a structure
// that references a string_question and includes handlers that indicate
// actions to perform based on whether the response matches a given
// regular expression.
type StringQuestionTestActionType struct {
	When_pattern        []PatternTestActionConditionType `xml:"http://scap.nist.gov/schema/ocil/2.0 when_pattern"`
	Title               TextType                         `xml:"http://scap.nist.gov/schema/ocil/2.0 title,omitempty"`
	When_unknown        TestActionConditionType          `xml:"http://scap.nist.gov/schema/ocil/2.0 when_unknown,omitempty"`
	When_not_tested     TestActionConditionType          `xml:"http://scap.nist.gov/schema/ocil/2.0 when_not_tested,omitempty"`
	When_not_applicable TestActionConditionType          `xml:"http://scap.nist.gov/schema/ocil/2.0 when_not_applicable,omitempty"`
	When_error          TestActionConditionType          `xml:"http://scap.nist.gov/schema/ocil/2.0 when_error,omitempty"`
	Notes               []string                         `xml:"http://scap.nist.gov/schema/ocil/2.0 notes,omitempty"`
	Question_ref        QuestionIDPattern                `xml:"question_ref,attr"`
	Id                  QuestionTestActionIDPattern      `xml:"id,attr"`
	Revision            int                              `xml:"revision,attr,omitempty"`
}

func (t *StringQuestionTestActionType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T StringQuestionTestActionType
	var overlay struct {
		*T
		Revision *int `xml:"revision,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Revision = (*int)(&overlay.T.Revision)
	return d.DecodeElement(&overlay, &start)
}

// The StringQuestionType type defines a question that
// requires a string answer.
type StringQuestionType struct {
	Question_text  []QuestionTextType `xml:"http://scap.nist.gov/schema/ocil/2.0 question_text"`
	Instructions   InstructionsType   `xml:"http://scap.nist.gov/schema/ocil/2.0 instructions,omitempty"`
	Notes          []string           `xml:"http://scap.nist.gov/schema/ocil/2.0 notes,omitempty"`
	Default_answer string             `xml:"default_answer,attr,omitempty"`
	Id             QuestionIDPattern  `xml:"id,attr"`
	Revision       int                `xml:"revision,attr,omitempty"`
}

func (t *StringQuestionType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T StringQuestionType
	var overlay struct {
		*T
		Revision *int `xml:"revision,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Revision = (*int)(&overlay.T.Revision)
	return d.DecodeElement(&overlay, &start)
}

// A type that is used to represent text from a variable that may be inserted into a text string within this model.
type SubstitutionTextType struct {
	Var_ref VariableIDPattern `xml:"var_ref,attr"`
}

// The SystemTargetType type defines structures containing
// information about the organization it belongs to, a set of ip addresses
// of computers/networks included in the system, descrioption about it, and
// the roles it performs.
type SystemTargetType struct {
	Organization string   `xml:"http://scap.nist.gov/schema/ocil/2.0 organization,omitempty"`
	Ipaddress    []string `xml:"http://scap.nist.gov/schema/ocil/2.0 ipaddress,omitempty"`
	Description  TextType `xml:"http://scap.nist.gov/schema/ocil/2.0 description,omitempty"`
	Name         string   `xml:"http://scap.nist.gov/schema/ocil/2.0 name"`
	Notes        []string `xml:"http://scap.nist.gov/schema/ocil/2.0 notes,omitempty"`
	Revision     int      `xml:"revision,attr,omitempty"`
}

func (t *SystemTargetType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T SystemTargetType
	var overlay struct {
		*T
		Revision *int `xml:"revision,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Revision = (*int)(&overlay.T.Revision)
	return d.DecodeElement(&overlay, &start)
}

// The TargetsType type defines structures containing a set
// of target elements.
type TargetsType struct {
	Target []NamedItemBaseType `xml:"http://scap.nist.gov/schema/ocil/2.0 target"`
}

// The TestActionConditionType complex type specifies processing
// instructions - either produce a result or move on to another test. The
// TestActionConditionType is extended by all handlers ("when_...") in
// test_actions.
type TestActionConditionType struct {
	Result          ResultType        `xml:"http://scap.nist.gov/schema/ocil/2.0 result"`
	Test_action_ref TestActionRefType `xml:"http://scap.nist.gov/schema/ocil/2.0 test_action_ref"`
	Artifact_refs   ArtifactRefsType  `xml:"http://scap.nist.gov/schema/ocil/2.0 artifact_refs,omitempty"`
}

// The TestActionRefType type defines a structure that holds
// a reference (id) to a test_action or questionnaire.
type TestActionRefType struct {
	TestActionRefValuePattern TestActionRefValuePattern `xml:",chardata"`
	Negate                    bool                      `xml:"negate,attr,omitempty"`
}

func (t *TestActionRefType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T TestActionRefType
	var overlay struct {
		*T
		Negate *bool `xml:"negate,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Negate = (*bool)(&overlay.T.Negate)
	return d.DecodeElement(&overlay, &start)
}

// Must match the pattern ocil:[A-Za-z0-9_\-\.]+:testaction:[1-9][0-9]*|ocil:[A-Za-z0-9_\-\.]+:questionnaire:[1-9][0-9]*
type TestActionRefValuePattern string

// The TestActionResultType type defines structures
// containing all computed results of a TestActionType. One of these
// elements will appear for each test_action evaluated.
type TestActionResultType struct {
	Artifact_results ArtifactResultsType       `xml:"http://scap.nist.gov/schema/ocil/2.0 artifact_results,omitempty"`
	Test_action_ref  TestActionRefValuePattern `xml:"test_action_ref,attr"`
	Result           ResultType                `xml:"result,attr"`
}

// The TestActionResultsType type defines structures
// containing computed results of all the evaluated test action
// types.
type TestActionResultsType struct {
	Test_action_result []TestActionResultType `xml:"http://scap.nist.gov/schema/ocil/2.0 test_action_result"`
}

// The TestActionsType type defines a container for a set of
// test action elements.
type TestActionsType struct {
	Test_action []ItemBaseType `xml:"http://scap.nist.gov/schema/ocil/2.0 test_action"`
}

// The data model that holds text-based artifacts.
type TextArtifactValueType struct {
	Data      string `xml:"http://scap.nist.gov/schema/ocil/2.0 data"`
	Mime_type string `xml:"mime_type,attr"`
}

// The TextType complex type defines an element that holds
// basic string information.
type TextType struct {
	Value string `xml:",chardata"`
	Lang  Lang   `xml:"lang,attr,omitempty"`
}

// The UserResponseType type defines structures containing
// the type of response. The question could have been answered or an
// exceptional condition may have occurred.
type UserResponseType string

// The UserType type defines structures containing
// information about a user such as name, organization, position, email, and
// role.
type UserType struct {
	Organization []string `xml:"http://scap.nist.gov/schema/ocil/2.0 organization,omitempty"`
	Position     []string `xml:"http://scap.nist.gov/schema/ocil/2.0 position,omitempty"`
	Email        []string `xml:"http://scap.nist.gov/schema/ocil/2.0 email,omitempty"`
	Name         string   `xml:"http://scap.nist.gov/schema/ocil/2.0 name"`
	Notes        []string `xml:"http://scap.nist.gov/schema/ocil/2.0 notes,omitempty"`
	Revision     int      `xml:"revision,attr,omitempty"`
}

func (t *UserType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T UserType
	var overlay struct {
		*T
		Revision *int `xml:"revision,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Revision = (*int)(&overlay.T.Revision)
	return d.DecodeElement(&overlay, &start)
}

// May be one of TEXT, NUMERIC
type VariableDataType string

// Must match the pattern ocil:[A-Za-z0-9_\-\.]+:variable:[1-9][0-9]*
type VariableIDPattern string

// The VariableSetType type defines structures containing
// information describing how to compute a variable value. It holds the
// patterns, choice_refs, range, or boolean values to be matched; and the
// appropriate value to be stored on the variable based on the
// match.
type VariableSetType struct {
	Expression []SetExpressionBaseType `xml:"http://scap.nist.gov/schema/ocil/2.0 expression"`
}

// The VariableType type defines structures used to hold a
// single value.
type VariableType struct {
	Description TextType          `xml:"http://scap.nist.gov/schema/ocil/2.0 description,omitempty"`
	Notes       []string          `xml:"http://scap.nist.gov/schema/ocil/2.0 notes,omitempty"`
	Id          VariableIDPattern `xml:"id,attr"`
	Datatype    VariableDataType  `xml:"datatype,attr"`
	Revision    int               `xml:"revision,attr,omitempty"`
}

func (t *VariableType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T VariableType
	var overlay struct {
		*T
		Revision *int `xml:"revision,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Revision = (*int)(&overlay.T.Revision)
	return d.DecodeElement(&overlay, &start)
}

// The VariablesType type defines structures containing a
// set of variables.
type VariablesType struct {
	Variable []VariableType `xml:"http://scap.nist.gov/schema/ocil/2.0 variable"`
}

// May be one of PASS, FAIL
type _anon1 string

// May be one of ANSWERED
type _anon2 string

type xsdBase64Binary []byte

func (b *xsdBase64Binary) UnmarshalText(text []byte) (err error) {
	*b, err = base64.StdEncoding.DecodeString(string(text))
	return
}
func (b xsdBase64Binary) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	enc := base64.NewEncoder(base64.StdEncoding, &buf)
	enc.Write([]byte(b))
	enc.Close()
	return buf.Bytes(), nil
}

type xsdDateTime time.Time

func (t *xsdDateTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalText() ([]byte, error) {
	return []byte((time.Time)(t).Format("2006-01-02T15:04:05.999999999")), nil
}
func (t xsdDateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDateTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
func _unmarshalTime(text []byte, t *time.Time, format string) (err error) {
	s := string(bytes.TrimSpace(text))
	*t, err = time.Parse(format, s)
	if _, ok := err.(*time.ParseError); ok {
		*t, err = time.Parse(format+"Z07:00", s)
	}
	return err
}

