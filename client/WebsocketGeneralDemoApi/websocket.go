package WebsocketGeneralDemoApi

type GeneralUpstreamTextEvent struct {
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	Object *struct {
		StrField  *string   `json:"strField,omitempty" xml:"strField,omitempty"`
		IntField  *int32    `json:"intField,omitempty" xml:"intField,omitempty"`
		InnerList []*string `json:"innerList,omitempty" xml:"innerList,omitempty"`
	} `json:"object,omitempty" xml:"object,omitempty"`
	List []*struct {
		Key2     *bool   `json:"boolField,omitempty" xml:"boolField,omitempty"`
		StrField *string `json:"strField,omitempty" xml:"strField,omitempty"`
	} `json:"list,omitempty" xml:"list,omitempty"`
	Map map[string]interface{} `json:"map,omitempty" xml:"map,omitempty"`
}

type GeneralUpstreamBinaryEvent interface{}

type UpstreamDefaultTextEvent interface{}

type DownstreamDefaultTextEvent interface{}

type GeneralDownstreamTextEvent struct {
	AudioId        *int32  `json:"audioId,omitempty" xml:"audioId,omitempty"`
	AudioType      *string `json:"audioType,omitempty" xml:"audioType,omitempty"`
	ProcessStatus  *string `json:"processStatus,omitempty" xml:"processStatus,omitempty"`
	AdditionalConf *struct {
		Timeout        *int64  `json:"timeout,omitempty" xml:"timeout,omitempty"`
		MaxAudioLength *string `json:"maxAudioLength,omitempty" xml:"maxAudioLength,omitempty"`
	} `json:"additionalConf,omitempty" xml:"additionalConf,omitempty"`
}

type GeneralDownstreamBinaryEvent interface{}
