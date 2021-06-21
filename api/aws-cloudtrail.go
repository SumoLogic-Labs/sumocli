package api

type AWSCloudTrailCollection struct {
	ApiVersion string        `json:"api.version"`
	Source     AWSCloudTrail `json:"source"`
}

type AWSCloudTrailCollectionResponse struct {
	Source AWSCloudTrailResponse `json:"source"`
}

type AWSCloudTrail struct {
	SourceType                 string              `json:"sourceType"`
	Name                       string              `json:"name"`
	Category                   string              `json:"category"`
	ContentType                string              `json:"contentType"`
	ThirdPartyRef              ThirdPartyReference `json:"thirdPartyRef"`
	ScanInterval               int                 `json:"scanInterval"`
	Paused                     bool                `json:"paused"`
	AutomaticDateParsing       bool                `json:"automaticDateParsing"`
	MultilineProcessingEnabled bool                `json:"multilineProcessingEnabled"`
	UseAutolineMatching        bool                `json:"useAutolineMatching"`
	ForceTimeZone              bool                `json:"forceTimeZone"`
	TimeZone                   string              `json:"timeZone"`
	Filters                    []SourceFilters     `json:"filters"`
	CutoffTimestamp            int                 `json:"cutoffTimestamp,omitempty"`
	CutoffRelativeTime         string              `json:"cutoffRelativeTime"`
	Encoding                   string              `json:"encoding"`
	Fields                     map[string]string   `json:"fields"`
}

type AWSCloudTrailResponse struct {
	Alive                      bool                `json:"alive"`
	Id                         int                 `json:"id"`
	Url                        string              `json:"url"`
	SourceType                 string              `json:"sourceType"`
	Name                       string              `json:"name"`
	Category                   string              `json:"category"`
	ContentType                string              `json:"contentType"`
	ThirdPartyRef              ThirdPartyReference `json:"thirdPartyRef"`
	ScanInterval               int                 `json:"scanInterval"`
	Paused                     bool                `json:"paused"`
	AutomaticDateParsing       bool                `json:"automaticDateParsing"`
	MultilineProcessingEnabled bool                `json:"multilineProcessingEnabled"`
	UseAutolineMatching        bool                `json:"useAutolineMatching"`
	ForceTimeZone              bool                `json:"forceTimeZone"`
	TimeZone                   string              `json:"timeZone"`
	Filters                    []SourceFilters     `json:"filters"`
	CutoffTimestamp            int                 `json:"cutoffTimestamp,omitempty"`
	CutoffRelativeTime         string              `json:"cutoffRelativeTime"`
	Encoding                   string              `json:"encoding"`
	Fields                     map[string]string   `json:"fields"`
}

type ThirdPartyReference struct {
	Resources []ThirdPartyReferenceResources `json:"resources"`
}

type ThirdPartyReferenceResources struct {
	ServiceType    string                                     `json:"serviceType"`
	Path           ThirdPartyReferenceResourcesPath           `json:"path"`
	Authentication ThirdPartyReferenceResourcesAuthentication `json:"authentication"`
}

type ThirdPartyReferenceResourcesPath struct {
	Type                      string                                   `json:"type"`
	BucketName                string                                   `json:"bucketName"`
	PathExpression            string                                   `json:"pathExpression"`
	SnsTopicOrSubscriptionArn ThirdPartyReferenceResourcesPathSnsTopic `json:"snsTopicOrSubscriptionArn,omitempty"`
}

type ThirdPartyReferenceResourcesPathSnsTopic struct {
	IsSuccess bool   `json:"isSuccess"`
	Arn       string `json:"arn"`
}

type ThirdPartyReferenceResourcesAuthentication struct {
	Type    string `json:"type"`
	AwsId   string `json:"awsId,omitempty"`
	AwsKey  string `json:"awsKey,omitempty"`
	RoleArn string `json:"roleARN,omitempty"`
}
