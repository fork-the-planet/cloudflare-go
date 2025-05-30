// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/tidwall/gjson"
)

// DLPProfileCustomService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPProfileCustomService] method instead.
type DLPProfileCustomService struct {
	Options []option.RequestOption
}

// NewDLPProfileCustomService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDLPProfileCustomService(opts ...option.RequestOption) (r *DLPProfileCustomService) {
	r = &DLPProfileCustomService{}
	r.Options = opts
	return
}

// Creates a DLP custom profile.
func (r *DLPProfileCustomService) New(ctx context.Context, params DLPProfileCustomNewParams, opts ...option.RequestOption) (res *DLPProfileCustomNewResponseUnion, err error) {
	var env DLPProfileCustomNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a DLP custom profile.
func (r *DLPProfileCustomService) Update(ctx context.Context, profileID string, params DLPProfileCustomUpdateParams, opts ...option.RequestOption) (res *Profile, err error) {
	var env DLPProfileCustomUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom/%s", params.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a DLP custom profile.
func (r *DLPProfileCustomService) Delete(ctx context.Context, profileID string, body DLPProfileCustomDeleteParams, opts ...option.RequestOption) (res *DLPProfileCustomDeleteResponse, err error) {
	var env DLPProfileCustomDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom/%s", body.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a custom DLP profile by id.
func (r *DLPProfileCustomService) Get(ctx context.Context, profileID string, query DLPProfileCustomGetParams, opts ...option.RequestOption) (res *Profile, err error) {
	var env DLPProfileCustomGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom/%s", query.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Pattern struct {
	Regex string `json:"regex,required"`
	// Deprecated: deprecated
	Validation PatternValidation `json:"validation"`
	JSON       patternJSON       `json:"-"`
}

// patternJSON contains the JSON metadata for the struct [Pattern]
type patternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Pattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r patternJSON) RawJSON() string {
	return r.raw
}

type PatternValidation string

const (
	PatternValidationLuhn PatternValidation = "luhn"
)

func (r PatternValidation) IsKnown() bool {
	switch r {
	case PatternValidationLuhn:
		return true
	}
	return false
}

type PatternParam struct {
	Regex param.Field[string] `json:"regex,required"`
	// Deprecated: deprecated
	Validation param.Field[PatternValidation] `json:"validation"`
}

func (r PatternParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Union satisfied by [DLPProfileCustomNewResponseCustomProfile],
// [DLPProfileCustomNewResponsePredefinedProfile],
// [DLPProfileCustomNewResponseIntegrationProfile] or
// [DLPProfileCustomNewResponseArray].
type DLPProfileCustomNewResponseUnion interface {
	implementsDLPProfileCustomNewResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileCustomNewResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseCustomProfile{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponsePredefinedProfile{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseIntegrationProfile{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseArray{}),
		},
	)
}

type DLPProfileCustomNewResponseCustomProfile struct {
	// The id of the profile (uuid).
	ID string `json:"id,required" format:"uuid"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount int64 `json:"allowed_match_count,required"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness ContextAwareness `json:"context_awareness,required"`
	// When the profile was created.
	CreatedAt time.Time                                       `json:"created_at,required" format:"date-time"`
	Entries   []DLPProfileCustomNewResponseCustomProfileEntry `json:"entries,required"`
	// The name of the profile.
	Name       string                                       `json:"name,required"`
	OCREnabled bool                                         `json:"ocr_enabled,required"`
	Type       DLPProfileCustomNewResponseCustomProfileType `json:"type,required"`
	// When the profile was lasted updated.
	UpdatedAt           time.Time                                                   `json:"updated_at,required" format:"date-time"`
	AIContextEnabled    bool                                                        `json:"ai_context_enabled"`
	ConfidenceThreshold DLPProfileCustomNewResponseCustomProfileConfidenceThreshold `json:"confidence_threshold"`
	// The description of the profile.
	Description string                                       `json:"description,nullable"`
	JSON        dlpProfileCustomNewResponseCustomProfileJSON `json:"-"`
}

// dlpProfileCustomNewResponseCustomProfileJSON contains the JSON metadata for the
// struct [DLPProfileCustomNewResponseCustomProfile]
type dlpProfileCustomNewResponseCustomProfileJSON struct {
	ID                  apijson.Field
	AllowedMatchCount   apijson.Field
	ContextAwareness    apijson.Field
	CreatedAt           apijson.Field
	Entries             apijson.Field
	Name                apijson.Field
	OCREnabled          apijson.Field
	Type                apijson.Field
	UpdatedAt           apijson.Field
	AIContextEnabled    apijson.Field
	ConfidenceThreshold apijson.Field
	Description         apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseCustomProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseCustomProfileJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseCustomProfile) implementsDLPProfileCustomNewResponseUnion() {}

type DLPProfileCustomNewResponseCustomProfileEntry struct {
	ID      string                                              `json:"id,required" format:"uuid"`
	Enabled bool                                                `json:"enabled,required"`
	Name    string                                              `json:"name,required"`
	Type    DLPProfileCustomNewResponseCustomProfileEntriesType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPProfileCustomNewResponseCustomProfileEntriesPredefinedEntryConfidence].
	Confidence interface{} `json:"confidence"`
	CreatedAt  time.Time   `json:"created_at" format:"date-time"`
	Pattern    Pattern     `json:"pattern"`
	ProfileID  string      `json:"profile_id,nullable" format:"uuid"`
	Secret     bool        `json:"secret"`
	UpdatedAt  time.Time   `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                                       `json:"word_list"`
	JSON     dlpProfileCustomNewResponseCustomProfileEntryJSON `json:"-"`
	union    DLPProfileCustomNewResponseCustomProfileEntriesUnion
}

// dlpProfileCustomNewResponseCustomProfileEntryJSON contains the JSON metadata for
// the struct [DLPProfileCustomNewResponseCustomProfileEntry]
type dlpProfileCustomNewResponseCustomProfileEntryJSON struct {
	ID            apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	CaseSensitive apijson.Field
	Confidence    apijson.Field
	CreatedAt     apijson.Field
	Pattern       apijson.Field
	ProfileID     apijson.Field
	Secret        apijson.Field
	UpdatedAt     apijson.Field
	WordList      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r dlpProfileCustomNewResponseCustomProfileEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileCustomNewResponseCustomProfileEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileCustomNewResponseCustomProfileEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileCustomNewResponseCustomProfileEntriesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DLPProfileCustomNewResponseCustomProfileEntriesCustomEntry],
// [DLPProfileCustomNewResponseCustomProfileEntriesPredefinedEntry],
// [DLPProfileCustomNewResponseCustomProfileEntriesIntegrationEntry],
// [DLPProfileCustomNewResponseCustomProfileEntriesExactDataEntry],
// [DLPProfileCustomNewResponseCustomProfileEntriesWordListEntry].
func (r DLPProfileCustomNewResponseCustomProfileEntry) AsUnion() DLPProfileCustomNewResponseCustomProfileEntriesUnion {
	return r.union
}

// Union satisfied by [DLPProfileCustomNewResponseCustomProfileEntriesCustomEntry],
// [DLPProfileCustomNewResponseCustomProfileEntriesPredefinedEntry],
// [DLPProfileCustomNewResponseCustomProfileEntriesIntegrationEntry],
// [DLPProfileCustomNewResponseCustomProfileEntriesExactDataEntry] or
// [DLPProfileCustomNewResponseCustomProfileEntriesWordListEntry].
type DLPProfileCustomNewResponseCustomProfileEntriesUnion interface {
	implementsDLPProfileCustomNewResponseCustomProfileEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileCustomNewResponseCustomProfileEntriesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseCustomProfileEntriesCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseCustomProfileEntriesPredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseCustomProfileEntriesIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseCustomProfileEntriesExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseCustomProfileEntriesWordListEntry{}),
		},
	)
}

type DLPProfileCustomNewResponseCustomProfileEntriesCustomEntry struct {
	ID        string                                                         `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                      `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                           `json:"enabled,required"`
	Name      string                                                         `json:"name,required"`
	Pattern   Pattern                                                        `json:"pattern,required"`
	Type      DLPProfileCustomNewResponseCustomProfileEntriesCustomEntryType `json:"type,required"`
	UpdatedAt time.Time                                                      `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                         `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponseCustomProfileEntriesCustomEntryJSON `json:"-"`
}

// dlpProfileCustomNewResponseCustomProfileEntriesCustomEntryJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomNewResponseCustomProfileEntriesCustomEntry]
type dlpProfileCustomNewResponseCustomProfileEntriesCustomEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseCustomProfileEntriesCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseCustomProfileEntriesCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseCustomProfileEntriesCustomEntry) implementsDLPProfileCustomNewResponseCustomProfileEntry() {
}

type DLPProfileCustomNewResponseCustomProfileEntriesCustomEntryType string

const (
	DLPProfileCustomNewResponseCustomProfileEntriesCustomEntryTypeCustom DLPProfileCustomNewResponseCustomProfileEntriesCustomEntryType = "custom"
)

func (r DLPProfileCustomNewResponseCustomProfileEntriesCustomEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomProfileEntriesCustomEntryTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomProfileEntriesPredefinedEntry struct {
	ID         string                                                                   `json:"id,required" format:"uuid"`
	Confidence DLPProfileCustomNewResponseCustomProfileEntriesPredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                                                     `json:"enabled,required"`
	Name       string                                                                   `json:"name,required"`
	Type       DLPProfileCustomNewResponseCustomProfileEntriesPredefinedEntryType       `json:"type,required"`
	ProfileID  string                                                                   `json:"profile_id,nullable" format:"uuid"`
	JSON       dlpProfileCustomNewResponseCustomProfileEntriesPredefinedEntryJSON       `json:"-"`
}

// dlpProfileCustomNewResponseCustomProfileEntriesPredefinedEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomNewResponseCustomProfileEntriesPredefinedEntry]
type dlpProfileCustomNewResponseCustomProfileEntriesPredefinedEntryJSON struct {
	ID          apijson.Field
	Confidence  apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseCustomProfileEntriesPredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseCustomProfileEntriesPredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseCustomProfileEntriesPredefinedEntry) implementsDLPProfileCustomNewResponseCustomProfileEntry() {
}

type DLPProfileCustomNewResponseCustomProfileEntriesPredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                                         `json:"available,required"`
	JSON      dlpProfileCustomNewResponseCustomProfileEntriesPredefinedEntryConfidenceJSON `json:"-"`
}

// dlpProfileCustomNewResponseCustomProfileEntriesPredefinedEntryConfidenceJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomNewResponseCustomProfileEntriesPredefinedEntryConfidence]
type dlpProfileCustomNewResponseCustomProfileEntriesPredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseCustomProfileEntriesPredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseCustomProfileEntriesPredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomNewResponseCustomProfileEntriesPredefinedEntryType string

const (
	DLPProfileCustomNewResponseCustomProfileEntriesPredefinedEntryTypePredefined DLPProfileCustomNewResponseCustomProfileEntriesPredefinedEntryType = "predefined"
)

func (r DLPProfileCustomNewResponseCustomProfileEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomProfileEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomProfileEntriesIntegrationEntry struct {
	ID        string                                                              `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                           `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                                `json:"enabled,required"`
	Name      string                                                              `json:"name,required"`
	Type      DLPProfileCustomNewResponseCustomProfileEntriesIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                                                           `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                              `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponseCustomProfileEntriesIntegrationEntryJSON `json:"-"`
}

// dlpProfileCustomNewResponseCustomProfileEntriesIntegrationEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomNewResponseCustomProfileEntriesIntegrationEntry]
type dlpProfileCustomNewResponseCustomProfileEntriesIntegrationEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseCustomProfileEntriesIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseCustomProfileEntriesIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseCustomProfileEntriesIntegrationEntry) implementsDLPProfileCustomNewResponseCustomProfileEntry() {
}

type DLPProfileCustomNewResponseCustomProfileEntriesIntegrationEntryType string

const (
	DLPProfileCustomNewResponseCustomProfileEntriesIntegrationEntryTypeIntegration DLPProfileCustomNewResponseCustomProfileEntriesIntegrationEntryType = "integration"
)

func (r DLPProfileCustomNewResponseCustomProfileEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomProfileEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomProfileEntriesExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                                              `json:"case_sensitive,required"`
	CreatedAt     time.Time                                                         `json:"created_at,required" format:"date-time"`
	Enabled       bool                                                              `json:"enabled,required"`
	Name          string                                                            `json:"name,required"`
	Secret        bool                                                              `json:"secret,required"`
	Type          DLPProfileCustomNewResponseCustomProfileEntriesExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                                                         `json:"updated_at,required" format:"date-time"`
	JSON          dlpProfileCustomNewResponseCustomProfileEntriesExactDataEntryJSON `json:"-"`
}

// dlpProfileCustomNewResponseCustomProfileEntriesExactDataEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomNewResponseCustomProfileEntriesExactDataEntry]
type dlpProfileCustomNewResponseCustomProfileEntriesExactDataEntryJSON struct {
	ID            apijson.Field
	CaseSensitive apijson.Field
	CreatedAt     apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Secret        apijson.Field
	Type          apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseCustomProfileEntriesExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseCustomProfileEntriesExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseCustomProfileEntriesExactDataEntry) implementsDLPProfileCustomNewResponseCustomProfileEntry() {
}

type DLPProfileCustomNewResponseCustomProfileEntriesExactDataEntryType string

const (
	DLPProfileCustomNewResponseCustomProfileEntriesExactDataEntryTypeExactData DLPProfileCustomNewResponseCustomProfileEntriesExactDataEntryType = "exact_data"
)

func (r DLPProfileCustomNewResponseCustomProfileEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomProfileEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomProfileEntriesWordListEntry struct {
	ID        string                                                           `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                        `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                             `json:"enabled,required"`
	Name      string                                                           `json:"name,required"`
	Type      DLPProfileCustomNewResponseCustomProfileEntriesWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                                                        `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                                      `json:"word_list,required"`
	ProfileID string                                                           `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponseCustomProfileEntriesWordListEntryJSON `json:"-"`
}

// dlpProfileCustomNewResponseCustomProfileEntriesWordListEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomNewResponseCustomProfileEntriesWordListEntry]
type dlpProfileCustomNewResponseCustomProfileEntriesWordListEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	WordList    apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseCustomProfileEntriesWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseCustomProfileEntriesWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseCustomProfileEntriesWordListEntry) implementsDLPProfileCustomNewResponseCustomProfileEntry() {
}

type DLPProfileCustomNewResponseCustomProfileEntriesWordListEntryType string

const (
	DLPProfileCustomNewResponseCustomProfileEntriesWordListEntryTypeWordList DLPProfileCustomNewResponseCustomProfileEntriesWordListEntryType = "word_list"
)

func (r DLPProfileCustomNewResponseCustomProfileEntriesWordListEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomProfileEntriesWordListEntryTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomProfileEntriesType string

const (
	DLPProfileCustomNewResponseCustomProfileEntriesTypeCustom      DLPProfileCustomNewResponseCustomProfileEntriesType = "custom"
	DLPProfileCustomNewResponseCustomProfileEntriesTypePredefined  DLPProfileCustomNewResponseCustomProfileEntriesType = "predefined"
	DLPProfileCustomNewResponseCustomProfileEntriesTypeIntegration DLPProfileCustomNewResponseCustomProfileEntriesType = "integration"
	DLPProfileCustomNewResponseCustomProfileEntriesTypeExactData   DLPProfileCustomNewResponseCustomProfileEntriesType = "exact_data"
	DLPProfileCustomNewResponseCustomProfileEntriesTypeWordList    DLPProfileCustomNewResponseCustomProfileEntriesType = "word_list"
)

func (r DLPProfileCustomNewResponseCustomProfileEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomProfileEntriesTypeCustom, DLPProfileCustomNewResponseCustomProfileEntriesTypePredefined, DLPProfileCustomNewResponseCustomProfileEntriesTypeIntegration, DLPProfileCustomNewResponseCustomProfileEntriesTypeExactData, DLPProfileCustomNewResponseCustomProfileEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomProfileType string

const (
	DLPProfileCustomNewResponseCustomProfileTypeCustom DLPProfileCustomNewResponseCustomProfileType = "custom"
)

func (r DLPProfileCustomNewResponseCustomProfileType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomProfileTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomProfileConfidenceThreshold string

const (
	DLPProfileCustomNewResponseCustomProfileConfidenceThresholdLow      DLPProfileCustomNewResponseCustomProfileConfidenceThreshold = "low"
	DLPProfileCustomNewResponseCustomProfileConfidenceThresholdMedium   DLPProfileCustomNewResponseCustomProfileConfidenceThreshold = "medium"
	DLPProfileCustomNewResponseCustomProfileConfidenceThresholdHigh     DLPProfileCustomNewResponseCustomProfileConfidenceThreshold = "high"
	DLPProfileCustomNewResponseCustomProfileConfidenceThresholdVeryHigh DLPProfileCustomNewResponseCustomProfileConfidenceThreshold = "very_high"
)

func (r DLPProfileCustomNewResponseCustomProfileConfidenceThreshold) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomProfileConfidenceThresholdLow, DLPProfileCustomNewResponseCustomProfileConfidenceThresholdMedium, DLPProfileCustomNewResponseCustomProfileConfidenceThresholdHigh, DLPProfileCustomNewResponseCustomProfileConfidenceThresholdVeryHigh:
		return true
	}
	return false
}

type DLPProfileCustomNewResponsePredefinedProfile struct {
	// The id of the predefined profile (uuid).
	ID                string                                              `json:"id,required" format:"uuid"`
	AllowedMatchCount int64                                               `json:"allowed_match_count,required"`
	Entries           []DLPProfileCustomNewResponsePredefinedProfileEntry `json:"entries,required"`
	// The name of the predefined profile.
	Name                string                                                          `json:"name,required"`
	Type                DLPProfileCustomNewResponsePredefinedProfileType                `json:"type,required"`
	AIContextEnabled    bool                                                            `json:"ai_context_enabled"`
	ConfidenceThreshold DLPProfileCustomNewResponsePredefinedProfileConfidenceThreshold `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness ContextAwareness `json:"context_awareness"`
	OCREnabled       bool             `json:"ocr_enabled"`
	// Whether this profile can be accessed by anyone.
	OpenAccess bool                                             `json:"open_access"`
	JSON       dlpProfileCustomNewResponsePredefinedProfileJSON `json:"-"`
}

// dlpProfileCustomNewResponsePredefinedProfileJSON contains the JSON metadata for
// the struct [DLPProfileCustomNewResponsePredefinedProfile]
type dlpProfileCustomNewResponsePredefinedProfileJSON struct {
	ID                  apijson.Field
	AllowedMatchCount   apijson.Field
	Entries             apijson.Field
	Name                apijson.Field
	Type                apijson.Field
	AIContextEnabled    apijson.Field
	ConfidenceThreshold apijson.Field
	ContextAwareness    apijson.Field
	OCREnabled          apijson.Field
	OpenAccess          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponsePredefinedProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponsePredefinedProfileJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponsePredefinedProfile) implementsDLPProfileCustomNewResponseUnion() {}

type DLPProfileCustomNewResponsePredefinedProfileEntry struct {
	ID      string                                                  `json:"id,required" format:"uuid"`
	Enabled bool                                                    `json:"enabled,required"`
	Name    string                                                  `json:"name,required"`
	Type    DLPProfileCustomNewResponsePredefinedProfileEntriesType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPProfileCustomNewResponsePredefinedProfileEntriesPredefinedEntryConfidence].
	Confidence interface{} `json:"confidence"`
	CreatedAt  time.Time   `json:"created_at" format:"date-time"`
	Pattern    Pattern     `json:"pattern"`
	ProfileID  string      `json:"profile_id,nullable" format:"uuid"`
	Secret     bool        `json:"secret"`
	UpdatedAt  time.Time   `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                                           `json:"word_list"`
	JSON     dlpProfileCustomNewResponsePredefinedProfileEntryJSON `json:"-"`
	union    DLPProfileCustomNewResponsePredefinedProfileEntriesUnion
}

// dlpProfileCustomNewResponsePredefinedProfileEntryJSON contains the JSON metadata
// for the struct [DLPProfileCustomNewResponsePredefinedProfileEntry]
type dlpProfileCustomNewResponsePredefinedProfileEntryJSON struct {
	ID            apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	CaseSensitive apijson.Field
	Confidence    apijson.Field
	CreatedAt     apijson.Field
	Pattern       apijson.Field
	ProfileID     apijson.Field
	Secret        apijson.Field
	UpdatedAt     apijson.Field
	WordList      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r dlpProfileCustomNewResponsePredefinedProfileEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileCustomNewResponsePredefinedProfileEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileCustomNewResponsePredefinedProfileEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileCustomNewResponsePredefinedProfileEntriesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DLPProfileCustomNewResponsePredefinedProfileEntriesCustomEntry],
// [DLPProfileCustomNewResponsePredefinedProfileEntriesPredefinedEntry],
// [DLPProfileCustomNewResponsePredefinedProfileEntriesIntegrationEntry],
// [DLPProfileCustomNewResponsePredefinedProfileEntriesExactDataEntry],
// [DLPProfileCustomNewResponsePredefinedProfileEntriesWordListEntry].
func (r DLPProfileCustomNewResponsePredefinedProfileEntry) AsUnion() DLPProfileCustomNewResponsePredefinedProfileEntriesUnion {
	return r.union
}

// Union satisfied by
// [DLPProfileCustomNewResponsePredefinedProfileEntriesCustomEntry],
// [DLPProfileCustomNewResponsePredefinedProfileEntriesPredefinedEntry],
// [DLPProfileCustomNewResponsePredefinedProfileEntriesIntegrationEntry],
// [DLPProfileCustomNewResponsePredefinedProfileEntriesExactDataEntry] or
// [DLPProfileCustomNewResponsePredefinedProfileEntriesWordListEntry].
type DLPProfileCustomNewResponsePredefinedProfileEntriesUnion interface {
	implementsDLPProfileCustomNewResponsePredefinedProfileEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileCustomNewResponsePredefinedProfileEntriesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponsePredefinedProfileEntriesCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponsePredefinedProfileEntriesPredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponsePredefinedProfileEntriesIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponsePredefinedProfileEntriesExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponsePredefinedProfileEntriesWordListEntry{}),
		},
	)
}

type DLPProfileCustomNewResponsePredefinedProfileEntriesCustomEntry struct {
	ID        string                                                             `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                          `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                               `json:"enabled,required"`
	Name      string                                                             `json:"name,required"`
	Pattern   Pattern                                                            `json:"pattern,required"`
	Type      DLPProfileCustomNewResponsePredefinedProfileEntriesCustomEntryType `json:"type,required"`
	UpdatedAt time.Time                                                          `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                             `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponsePredefinedProfileEntriesCustomEntryJSON `json:"-"`
}

// dlpProfileCustomNewResponsePredefinedProfileEntriesCustomEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomNewResponsePredefinedProfileEntriesCustomEntry]
type dlpProfileCustomNewResponsePredefinedProfileEntriesCustomEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponsePredefinedProfileEntriesCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponsePredefinedProfileEntriesCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponsePredefinedProfileEntriesCustomEntry) implementsDLPProfileCustomNewResponsePredefinedProfileEntry() {
}

type DLPProfileCustomNewResponsePredefinedProfileEntriesCustomEntryType string

const (
	DLPProfileCustomNewResponsePredefinedProfileEntriesCustomEntryTypeCustom DLPProfileCustomNewResponsePredefinedProfileEntriesCustomEntryType = "custom"
)

func (r DLPProfileCustomNewResponsePredefinedProfileEntriesCustomEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponsePredefinedProfileEntriesCustomEntryTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomNewResponsePredefinedProfileEntriesPredefinedEntry struct {
	ID         string                                                                       `json:"id,required" format:"uuid"`
	Confidence DLPProfileCustomNewResponsePredefinedProfileEntriesPredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                                                         `json:"enabled,required"`
	Name       string                                                                       `json:"name,required"`
	Type       DLPProfileCustomNewResponsePredefinedProfileEntriesPredefinedEntryType       `json:"type,required"`
	ProfileID  string                                                                       `json:"profile_id,nullable" format:"uuid"`
	JSON       dlpProfileCustomNewResponsePredefinedProfileEntriesPredefinedEntryJSON       `json:"-"`
}

// dlpProfileCustomNewResponsePredefinedProfileEntriesPredefinedEntryJSON contains
// the JSON metadata for the struct
// [DLPProfileCustomNewResponsePredefinedProfileEntriesPredefinedEntry]
type dlpProfileCustomNewResponsePredefinedProfileEntriesPredefinedEntryJSON struct {
	ID          apijson.Field
	Confidence  apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponsePredefinedProfileEntriesPredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponsePredefinedProfileEntriesPredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponsePredefinedProfileEntriesPredefinedEntry) implementsDLPProfileCustomNewResponsePredefinedProfileEntry() {
}

type DLPProfileCustomNewResponsePredefinedProfileEntriesPredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                                             `json:"available,required"`
	JSON      dlpProfileCustomNewResponsePredefinedProfileEntriesPredefinedEntryConfidenceJSON `json:"-"`
}

// dlpProfileCustomNewResponsePredefinedProfileEntriesPredefinedEntryConfidenceJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomNewResponsePredefinedProfileEntriesPredefinedEntryConfidence]
type dlpProfileCustomNewResponsePredefinedProfileEntriesPredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponsePredefinedProfileEntriesPredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponsePredefinedProfileEntriesPredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomNewResponsePredefinedProfileEntriesPredefinedEntryType string

const (
	DLPProfileCustomNewResponsePredefinedProfileEntriesPredefinedEntryTypePredefined DLPProfileCustomNewResponsePredefinedProfileEntriesPredefinedEntryType = "predefined"
)

func (r DLPProfileCustomNewResponsePredefinedProfileEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponsePredefinedProfileEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomNewResponsePredefinedProfileEntriesIntegrationEntry struct {
	ID        string                                                                  `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                               `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                                    `json:"enabled,required"`
	Name      string                                                                  `json:"name,required"`
	Type      DLPProfileCustomNewResponsePredefinedProfileEntriesIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                                                               `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                                  `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponsePredefinedProfileEntriesIntegrationEntryJSON `json:"-"`
}

// dlpProfileCustomNewResponsePredefinedProfileEntriesIntegrationEntryJSON contains
// the JSON metadata for the struct
// [DLPProfileCustomNewResponsePredefinedProfileEntriesIntegrationEntry]
type dlpProfileCustomNewResponsePredefinedProfileEntriesIntegrationEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponsePredefinedProfileEntriesIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponsePredefinedProfileEntriesIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponsePredefinedProfileEntriesIntegrationEntry) implementsDLPProfileCustomNewResponsePredefinedProfileEntry() {
}

type DLPProfileCustomNewResponsePredefinedProfileEntriesIntegrationEntryType string

const (
	DLPProfileCustomNewResponsePredefinedProfileEntriesIntegrationEntryTypeIntegration DLPProfileCustomNewResponsePredefinedProfileEntriesIntegrationEntryType = "integration"
)

func (r DLPProfileCustomNewResponsePredefinedProfileEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponsePredefinedProfileEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomNewResponsePredefinedProfileEntriesExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                                                  `json:"case_sensitive,required"`
	CreatedAt     time.Time                                                             `json:"created_at,required" format:"date-time"`
	Enabled       bool                                                                  `json:"enabled,required"`
	Name          string                                                                `json:"name,required"`
	Secret        bool                                                                  `json:"secret,required"`
	Type          DLPProfileCustomNewResponsePredefinedProfileEntriesExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                                                             `json:"updated_at,required" format:"date-time"`
	JSON          dlpProfileCustomNewResponsePredefinedProfileEntriesExactDataEntryJSON `json:"-"`
}

// dlpProfileCustomNewResponsePredefinedProfileEntriesExactDataEntryJSON contains
// the JSON metadata for the struct
// [DLPProfileCustomNewResponsePredefinedProfileEntriesExactDataEntry]
type dlpProfileCustomNewResponsePredefinedProfileEntriesExactDataEntryJSON struct {
	ID            apijson.Field
	CaseSensitive apijson.Field
	CreatedAt     apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Secret        apijson.Field
	Type          apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponsePredefinedProfileEntriesExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponsePredefinedProfileEntriesExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponsePredefinedProfileEntriesExactDataEntry) implementsDLPProfileCustomNewResponsePredefinedProfileEntry() {
}

type DLPProfileCustomNewResponsePredefinedProfileEntriesExactDataEntryType string

const (
	DLPProfileCustomNewResponsePredefinedProfileEntriesExactDataEntryTypeExactData DLPProfileCustomNewResponsePredefinedProfileEntriesExactDataEntryType = "exact_data"
)

func (r DLPProfileCustomNewResponsePredefinedProfileEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponsePredefinedProfileEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomNewResponsePredefinedProfileEntriesWordListEntry struct {
	ID        string                                                               `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                            `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                                 `json:"enabled,required"`
	Name      string                                                               `json:"name,required"`
	Type      DLPProfileCustomNewResponsePredefinedProfileEntriesWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                                                            `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                                          `json:"word_list,required"`
	ProfileID string                                                               `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponsePredefinedProfileEntriesWordListEntryJSON `json:"-"`
}

// dlpProfileCustomNewResponsePredefinedProfileEntriesWordListEntryJSON contains
// the JSON metadata for the struct
// [DLPProfileCustomNewResponsePredefinedProfileEntriesWordListEntry]
type dlpProfileCustomNewResponsePredefinedProfileEntriesWordListEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	WordList    apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponsePredefinedProfileEntriesWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponsePredefinedProfileEntriesWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponsePredefinedProfileEntriesWordListEntry) implementsDLPProfileCustomNewResponsePredefinedProfileEntry() {
}

type DLPProfileCustomNewResponsePredefinedProfileEntriesWordListEntryType string

const (
	DLPProfileCustomNewResponsePredefinedProfileEntriesWordListEntryTypeWordList DLPProfileCustomNewResponsePredefinedProfileEntriesWordListEntryType = "word_list"
)

func (r DLPProfileCustomNewResponsePredefinedProfileEntriesWordListEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponsePredefinedProfileEntriesWordListEntryTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomNewResponsePredefinedProfileEntriesType string

const (
	DLPProfileCustomNewResponsePredefinedProfileEntriesTypeCustom      DLPProfileCustomNewResponsePredefinedProfileEntriesType = "custom"
	DLPProfileCustomNewResponsePredefinedProfileEntriesTypePredefined  DLPProfileCustomNewResponsePredefinedProfileEntriesType = "predefined"
	DLPProfileCustomNewResponsePredefinedProfileEntriesTypeIntegration DLPProfileCustomNewResponsePredefinedProfileEntriesType = "integration"
	DLPProfileCustomNewResponsePredefinedProfileEntriesTypeExactData   DLPProfileCustomNewResponsePredefinedProfileEntriesType = "exact_data"
	DLPProfileCustomNewResponsePredefinedProfileEntriesTypeWordList    DLPProfileCustomNewResponsePredefinedProfileEntriesType = "word_list"
)

func (r DLPProfileCustomNewResponsePredefinedProfileEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponsePredefinedProfileEntriesTypeCustom, DLPProfileCustomNewResponsePredefinedProfileEntriesTypePredefined, DLPProfileCustomNewResponsePredefinedProfileEntriesTypeIntegration, DLPProfileCustomNewResponsePredefinedProfileEntriesTypeExactData, DLPProfileCustomNewResponsePredefinedProfileEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomNewResponsePredefinedProfileType string

const (
	DLPProfileCustomNewResponsePredefinedProfileTypePredefined DLPProfileCustomNewResponsePredefinedProfileType = "predefined"
)

func (r DLPProfileCustomNewResponsePredefinedProfileType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponsePredefinedProfileTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomNewResponsePredefinedProfileConfidenceThreshold string

const (
	DLPProfileCustomNewResponsePredefinedProfileConfidenceThresholdLow      DLPProfileCustomNewResponsePredefinedProfileConfidenceThreshold = "low"
	DLPProfileCustomNewResponsePredefinedProfileConfidenceThresholdMedium   DLPProfileCustomNewResponsePredefinedProfileConfidenceThreshold = "medium"
	DLPProfileCustomNewResponsePredefinedProfileConfidenceThresholdHigh     DLPProfileCustomNewResponsePredefinedProfileConfidenceThreshold = "high"
	DLPProfileCustomNewResponsePredefinedProfileConfidenceThresholdVeryHigh DLPProfileCustomNewResponsePredefinedProfileConfidenceThreshold = "very_high"
)

func (r DLPProfileCustomNewResponsePredefinedProfileConfidenceThreshold) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponsePredefinedProfileConfidenceThresholdLow, DLPProfileCustomNewResponsePredefinedProfileConfidenceThresholdMedium, DLPProfileCustomNewResponsePredefinedProfileConfidenceThresholdHigh, DLPProfileCustomNewResponsePredefinedProfileConfidenceThresholdVeryHigh:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegrationProfile struct {
	ID        string                                               `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                            `json:"created_at,required" format:"date-time"`
	Entries   []DLPProfileCustomNewResponseIntegrationProfileEntry `json:"entries,required"`
	Name      string                                               `json:"name,required"`
	Type      DLPProfileCustomNewResponseIntegrationProfileType    `json:"type,required"`
	UpdatedAt time.Time                                            `json:"updated_at,required" format:"date-time"`
	// The description of the profile.
	Description string                                            `json:"description,nullable"`
	JSON        dlpProfileCustomNewResponseIntegrationProfileJSON `json:"-"`
}

// dlpProfileCustomNewResponseIntegrationProfileJSON contains the JSON metadata for
// the struct [DLPProfileCustomNewResponseIntegrationProfile]
type dlpProfileCustomNewResponseIntegrationProfileJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Entries     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseIntegrationProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseIntegrationProfileJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseIntegrationProfile) implementsDLPProfileCustomNewResponseUnion() {}

type DLPProfileCustomNewResponseIntegrationProfileEntry struct {
	ID      string                                                   `json:"id,required" format:"uuid"`
	Enabled bool                                                     `json:"enabled,required"`
	Name    string                                                   `json:"name,required"`
	Type    DLPProfileCustomNewResponseIntegrationProfileEntriesType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPProfileCustomNewResponseIntegrationProfileEntriesPredefinedEntryConfidence].
	Confidence interface{} `json:"confidence"`
	CreatedAt  time.Time   `json:"created_at" format:"date-time"`
	Pattern    Pattern     `json:"pattern"`
	ProfileID  string      `json:"profile_id,nullable" format:"uuid"`
	Secret     bool        `json:"secret"`
	UpdatedAt  time.Time   `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                                            `json:"word_list"`
	JSON     dlpProfileCustomNewResponseIntegrationProfileEntryJSON `json:"-"`
	union    DLPProfileCustomNewResponseIntegrationProfileEntriesUnion
}

// dlpProfileCustomNewResponseIntegrationProfileEntryJSON contains the JSON
// metadata for the struct [DLPProfileCustomNewResponseIntegrationProfileEntry]
type dlpProfileCustomNewResponseIntegrationProfileEntryJSON struct {
	ID            apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	CaseSensitive apijson.Field
	Confidence    apijson.Field
	CreatedAt     apijson.Field
	Pattern       apijson.Field
	ProfileID     apijson.Field
	Secret        apijson.Field
	UpdatedAt     apijson.Field
	WordList      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r dlpProfileCustomNewResponseIntegrationProfileEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileCustomNewResponseIntegrationProfileEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileCustomNewResponseIntegrationProfileEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileCustomNewResponseIntegrationProfileEntriesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DLPProfileCustomNewResponseIntegrationProfileEntriesCustomEntry],
// [DLPProfileCustomNewResponseIntegrationProfileEntriesPredefinedEntry],
// [DLPProfileCustomNewResponseIntegrationProfileEntriesIntegrationEntry],
// [DLPProfileCustomNewResponseIntegrationProfileEntriesExactDataEntry],
// [DLPProfileCustomNewResponseIntegrationProfileEntriesWordListEntry].
func (r DLPProfileCustomNewResponseIntegrationProfileEntry) AsUnion() DLPProfileCustomNewResponseIntegrationProfileEntriesUnion {
	return r.union
}

// Union satisfied by
// [DLPProfileCustomNewResponseIntegrationProfileEntriesCustomEntry],
// [DLPProfileCustomNewResponseIntegrationProfileEntriesPredefinedEntry],
// [DLPProfileCustomNewResponseIntegrationProfileEntriesIntegrationEntry],
// [DLPProfileCustomNewResponseIntegrationProfileEntriesExactDataEntry] or
// [DLPProfileCustomNewResponseIntegrationProfileEntriesWordListEntry].
type DLPProfileCustomNewResponseIntegrationProfileEntriesUnion interface {
	implementsDLPProfileCustomNewResponseIntegrationProfileEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileCustomNewResponseIntegrationProfileEntriesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseIntegrationProfileEntriesCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseIntegrationProfileEntriesPredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseIntegrationProfileEntriesIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseIntegrationProfileEntriesExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseIntegrationProfileEntriesWordListEntry{}),
		},
	)
}

type DLPProfileCustomNewResponseIntegrationProfileEntriesCustomEntry struct {
	ID        string                                                              `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                           `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                                `json:"enabled,required"`
	Name      string                                                              `json:"name,required"`
	Pattern   Pattern                                                             `json:"pattern,required"`
	Type      DLPProfileCustomNewResponseIntegrationProfileEntriesCustomEntryType `json:"type,required"`
	UpdatedAt time.Time                                                           `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                              `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponseIntegrationProfileEntriesCustomEntryJSON `json:"-"`
}

// dlpProfileCustomNewResponseIntegrationProfileEntriesCustomEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomNewResponseIntegrationProfileEntriesCustomEntry]
type dlpProfileCustomNewResponseIntegrationProfileEntriesCustomEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseIntegrationProfileEntriesCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseIntegrationProfileEntriesCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseIntegrationProfileEntriesCustomEntry) implementsDLPProfileCustomNewResponseIntegrationProfileEntry() {
}

type DLPProfileCustomNewResponseIntegrationProfileEntriesCustomEntryType string

const (
	DLPProfileCustomNewResponseIntegrationProfileEntriesCustomEntryTypeCustom DLPProfileCustomNewResponseIntegrationProfileEntriesCustomEntryType = "custom"
)

func (r DLPProfileCustomNewResponseIntegrationProfileEntriesCustomEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationProfileEntriesCustomEntryTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegrationProfileEntriesPredefinedEntry struct {
	ID         string                                                                        `json:"id,required" format:"uuid"`
	Confidence DLPProfileCustomNewResponseIntegrationProfileEntriesPredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                                                          `json:"enabled,required"`
	Name       string                                                                        `json:"name,required"`
	Type       DLPProfileCustomNewResponseIntegrationProfileEntriesPredefinedEntryType       `json:"type,required"`
	ProfileID  string                                                                        `json:"profile_id,nullable" format:"uuid"`
	JSON       dlpProfileCustomNewResponseIntegrationProfileEntriesPredefinedEntryJSON       `json:"-"`
}

// dlpProfileCustomNewResponseIntegrationProfileEntriesPredefinedEntryJSON contains
// the JSON metadata for the struct
// [DLPProfileCustomNewResponseIntegrationProfileEntriesPredefinedEntry]
type dlpProfileCustomNewResponseIntegrationProfileEntriesPredefinedEntryJSON struct {
	ID          apijson.Field
	Confidence  apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseIntegrationProfileEntriesPredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseIntegrationProfileEntriesPredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseIntegrationProfileEntriesPredefinedEntry) implementsDLPProfileCustomNewResponseIntegrationProfileEntry() {
}

type DLPProfileCustomNewResponseIntegrationProfileEntriesPredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                                              `json:"available,required"`
	JSON      dlpProfileCustomNewResponseIntegrationProfileEntriesPredefinedEntryConfidenceJSON `json:"-"`
}

// dlpProfileCustomNewResponseIntegrationProfileEntriesPredefinedEntryConfidenceJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomNewResponseIntegrationProfileEntriesPredefinedEntryConfidence]
type dlpProfileCustomNewResponseIntegrationProfileEntriesPredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseIntegrationProfileEntriesPredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseIntegrationProfileEntriesPredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomNewResponseIntegrationProfileEntriesPredefinedEntryType string

const (
	DLPProfileCustomNewResponseIntegrationProfileEntriesPredefinedEntryTypePredefined DLPProfileCustomNewResponseIntegrationProfileEntriesPredefinedEntryType = "predefined"
)

func (r DLPProfileCustomNewResponseIntegrationProfileEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationProfileEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegrationProfileEntriesIntegrationEntry struct {
	ID        string                                                                   `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                                `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                                     `json:"enabled,required"`
	Name      string                                                                   `json:"name,required"`
	Type      DLPProfileCustomNewResponseIntegrationProfileEntriesIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                                                                `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                                   `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponseIntegrationProfileEntriesIntegrationEntryJSON `json:"-"`
}

// dlpProfileCustomNewResponseIntegrationProfileEntriesIntegrationEntryJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomNewResponseIntegrationProfileEntriesIntegrationEntry]
type dlpProfileCustomNewResponseIntegrationProfileEntriesIntegrationEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseIntegrationProfileEntriesIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseIntegrationProfileEntriesIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseIntegrationProfileEntriesIntegrationEntry) implementsDLPProfileCustomNewResponseIntegrationProfileEntry() {
}

type DLPProfileCustomNewResponseIntegrationProfileEntriesIntegrationEntryType string

const (
	DLPProfileCustomNewResponseIntegrationProfileEntriesIntegrationEntryTypeIntegration DLPProfileCustomNewResponseIntegrationProfileEntriesIntegrationEntryType = "integration"
)

func (r DLPProfileCustomNewResponseIntegrationProfileEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationProfileEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegrationProfileEntriesExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                                                   `json:"case_sensitive,required"`
	CreatedAt     time.Time                                                              `json:"created_at,required" format:"date-time"`
	Enabled       bool                                                                   `json:"enabled,required"`
	Name          string                                                                 `json:"name,required"`
	Secret        bool                                                                   `json:"secret,required"`
	Type          DLPProfileCustomNewResponseIntegrationProfileEntriesExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                                                              `json:"updated_at,required" format:"date-time"`
	JSON          dlpProfileCustomNewResponseIntegrationProfileEntriesExactDataEntryJSON `json:"-"`
}

// dlpProfileCustomNewResponseIntegrationProfileEntriesExactDataEntryJSON contains
// the JSON metadata for the struct
// [DLPProfileCustomNewResponseIntegrationProfileEntriesExactDataEntry]
type dlpProfileCustomNewResponseIntegrationProfileEntriesExactDataEntryJSON struct {
	ID            apijson.Field
	CaseSensitive apijson.Field
	CreatedAt     apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Secret        apijson.Field
	Type          apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseIntegrationProfileEntriesExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseIntegrationProfileEntriesExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseIntegrationProfileEntriesExactDataEntry) implementsDLPProfileCustomNewResponseIntegrationProfileEntry() {
}

type DLPProfileCustomNewResponseIntegrationProfileEntriesExactDataEntryType string

const (
	DLPProfileCustomNewResponseIntegrationProfileEntriesExactDataEntryTypeExactData DLPProfileCustomNewResponseIntegrationProfileEntriesExactDataEntryType = "exact_data"
)

func (r DLPProfileCustomNewResponseIntegrationProfileEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationProfileEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegrationProfileEntriesWordListEntry struct {
	ID        string                                                                `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                             `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                                  `json:"enabled,required"`
	Name      string                                                                `json:"name,required"`
	Type      DLPProfileCustomNewResponseIntegrationProfileEntriesWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                                                             `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                                           `json:"word_list,required"`
	ProfileID string                                                                `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponseIntegrationProfileEntriesWordListEntryJSON `json:"-"`
}

// dlpProfileCustomNewResponseIntegrationProfileEntriesWordListEntryJSON contains
// the JSON metadata for the struct
// [DLPProfileCustomNewResponseIntegrationProfileEntriesWordListEntry]
type dlpProfileCustomNewResponseIntegrationProfileEntriesWordListEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	WordList    apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseIntegrationProfileEntriesWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseIntegrationProfileEntriesWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseIntegrationProfileEntriesWordListEntry) implementsDLPProfileCustomNewResponseIntegrationProfileEntry() {
}

type DLPProfileCustomNewResponseIntegrationProfileEntriesWordListEntryType string

const (
	DLPProfileCustomNewResponseIntegrationProfileEntriesWordListEntryTypeWordList DLPProfileCustomNewResponseIntegrationProfileEntriesWordListEntryType = "word_list"
)

func (r DLPProfileCustomNewResponseIntegrationProfileEntriesWordListEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationProfileEntriesWordListEntryTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegrationProfileEntriesType string

const (
	DLPProfileCustomNewResponseIntegrationProfileEntriesTypeCustom      DLPProfileCustomNewResponseIntegrationProfileEntriesType = "custom"
	DLPProfileCustomNewResponseIntegrationProfileEntriesTypePredefined  DLPProfileCustomNewResponseIntegrationProfileEntriesType = "predefined"
	DLPProfileCustomNewResponseIntegrationProfileEntriesTypeIntegration DLPProfileCustomNewResponseIntegrationProfileEntriesType = "integration"
	DLPProfileCustomNewResponseIntegrationProfileEntriesTypeExactData   DLPProfileCustomNewResponseIntegrationProfileEntriesType = "exact_data"
	DLPProfileCustomNewResponseIntegrationProfileEntriesTypeWordList    DLPProfileCustomNewResponseIntegrationProfileEntriesType = "word_list"
)

func (r DLPProfileCustomNewResponseIntegrationProfileEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationProfileEntriesTypeCustom, DLPProfileCustomNewResponseIntegrationProfileEntriesTypePredefined, DLPProfileCustomNewResponseIntegrationProfileEntriesTypeIntegration, DLPProfileCustomNewResponseIntegrationProfileEntriesTypeExactData, DLPProfileCustomNewResponseIntegrationProfileEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegrationProfileType string

const (
	DLPProfileCustomNewResponseIntegrationProfileTypeIntegration DLPProfileCustomNewResponseIntegrationProfileType = "integration"
)

func (r DLPProfileCustomNewResponseIntegrationProfileType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationProfileTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseArray []Profile

func (r DLPProfileCustomNewResponseArray) implementsDLPProfileCustomNewResponseUnion() {}

type DLPProfileCustomDeleteResponse = interface{}

type DLPProfileCustomNewParams struct {
	AccountID param.Field[string]                `path:"account_id,required"`
	Body      DLPProfileCustomNewParamsBodyUnion `json:"body,required"`
}

func (r DLPProfileCustomNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DLPProfileCustomNewParamsBody struct {
	AIContextEnabled param.Field[bool] `json:"ai_context_enabled"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount   param.Field[int64]  `json:"allowed_match_count"`
	ConfidenceThreshold param.Field[string] `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness param.Field[ContextAwarenessParam] `json:"context_awareness"`
	// The description of the profile.
	Description   param.Field[string]      `json:"description"`
	Entries       param.Field[interface{}] `json:"entries"`
	Name          param.Field[string]      `json:"name"`
	OCREnabled    param.Field[bool]        `json:"ocr_enabled"`
	Profiles      param.Field[interface{}] `json:"profiles"`
	SharedEntries param.Field[interface{}] `json:"shared_entries"`
}

func (r DLPProfileCustomNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsBody) implementsDLPProfileCustomNewParamsBodyUnion() {}

// Satisfied by [zero_trust.DLPProfileCustomNewParamsBodyProfiles],
// [zero_trust.DLPProfileCustomNewParamsBodyDLPNewCustomProfile],
// [DLPProfileCustomNewParamsBody].
type DLPProfileCustomNewParamsBodyUnion interface {
	implementsDLPProfileCustomNewParamsBodyUnion()
}

// Deprecated: deprecated
type DLPProfileCustomNewParamsBodyProfiles struct {
	Profiles param.Field[[]DLPProfileCustomNewParamsBodyProfilesProfile] `json:"profiles,required"`
}

func (r DLPProfileCustomNewParamsBodyProfiles) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsBodyProfiles) implementsDLPProfileCustomNewParamsBodyUnion() {}

type DLPProfileCustomNewParamsBodyProfilesProfile struct {
	Entries          param.Field[[]DLPProfileCustomNewParamsBodyProfilesProfilesEntryUnion] `json:"entries,required"`
	Name             param.Field[string]                                                    `json:"name,required"`
	AIContextEnabled param.Field[bool]                                                      `json:"ai_context_enabled"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount   param.Field[int64]  `json:"allowed_match_count"`
	ConfidenceThreshold param.Field[string] `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness param.Field[ContextAwarenessParam] `json:"context_awareness"`
	// The description of the profile.
	Description param.Field[string] `json:"description"`
	OCREnabled  param.Field[bool]   `json:"ocr_enabled"`
	// Entries from other profiles (e.g. pre-defined Cloudflare profiles, or your
	// Microsoft Information Protection profiles).
	SharedEntries param.Field[[]DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntryUnion] `json:"shared_entries"`
}

func (r DLPProfileCustomNewParamsBodyProfilesProfile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPProfileCustomNewParamsBodyProfilesProfilesEntry struct {
	Enabled param.Field[bool]         `json:"enabled,required"`
	Name    param.Field[string]       `json:"name,required"`
	Pattern param.Field[PatternParam] `json:"pattern"`
	Words   param.Field[interface{}]  `json:"words"`
}

func (r DLPProfileCustomNewParamsBodyProfilesProfilesEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsBodyProfilesProfilesEntry) implementsDLPProfileCustomNewParamsBodyProfilesProfilesEntryUnion() {
}

// Satisfied by
// [zero_trust.DLPProfileCustomNewParamsBodyProfilesProfilesEntriesDLPNewCustomEntry],
// [zero_trust.DLPProfileCustomNewParamsBodyProfilesProfilesEntriesDLPNewWordListEntry],
// [DLPProfileCustomNewParamsBodyProfilesProfilesEntry].
type DLPProfileCustomNewParamsBodyProfilesProfilesEntryUnion interface {
	implementsDLPProfileCustomNewParamsBodyProfilesProfilesEntryUnion()
}

type DLPProfileCustomNewParamsBodyProfilesProfilesEntriesDLPNewCustomEntry struct {
	Enabled param.Field[bool]         `json:"enabled,required"`
	Name    param.Field[string]       `json:"name,required"`
	Pattern param.Field[PatternParam] `json:"pattern,required"`
}

func (r DLPProfileCustomNewParamsBodyProfilesProfilesEntriesDLPNewCustomEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsBodyProfilesProfilesEntriesDLPNewCustomEntry) implementsDLPProfileCustomNewParamsBodyProfilesProfilesEntryUnion() {
}

type DLPProfileCustomNewParamsBodyProfilesProfilesEntriesDLPNewWordListEntry struct {
	Enabled param.Field[bool]     `json:"enabled,required"`
	Name    param.Field[string]   `json:"name,required"`
	Words   param.Field[[]string] `json:"words,required"`
}

func (r DLPProfileCustomNewParamsBodyProfilesProfilesEntriesDLPNewWordListEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsBodyProfilesProfilesEntriesDLPNewWordListEntry) implementsDLPProfileCustomNewParamsBodyProfilesProfilesEntryUnion() {
}

type DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntry struct {
	Enabled   param.Field[bool]                                                                `json:"enabled,required"`
	EntryID   param.Field[string]                                                              `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesEntryType] `json:"entry_type,required"`
}

func (r DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntry) implementsDLPProfileCustomNewParamsBodyProfilesProfilesSharedEntryUnion() {
}

// Satisfied by
// [zero_trust.DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesObject],
// [zero_trust.DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesObject],
// [zero_trust.DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesObject],
// [zero_trust.DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesObject],
// [DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntry].
type DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntryUnion interface {
	implementsDLPProfileCustomNewParamsBodyProfilesProfilesSharedEntryUnion()
}

type DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesObject struct {
	Enabled   param.Field[bool]                                                                      `json:"enabled,required"`
	EntryID   param.Field[string]                                                                    `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesObjectEntryType] `json:"entry_type,required"`
}

func (r DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesObject) implementsDLPProfileCustomNewParamsBodyProfilesProfilesSharedEntryUnion() {
}

type DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesObjectEntryType string

const (
	DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesObjectEntryTypeCustom DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesObjectEntryType = "custom"
)

func (r DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesObjectEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesObjectEntryTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesEntryType string

const (
	DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesEntryTypeCustom      DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesEntryType = "custom"
	DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesEntryTypePredefined  DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesEntryType = "predefined"
	DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesEntryTypeIntegration DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesEntryType = "integration"
	DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesEntryTypeExactData   DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesEntryType = "exact_data"
)

func (r DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesEntryTypeCustom, DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesEntryTypePredefined, DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesEntryTypeIntegration, DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesEntryTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomNewParamsBodyDLPNewCustomProfile struct {
	Entries          param.Field[[]DLPProfileCustomNewParamsBodyDLPNewCustomProfileEntryUnion] `json:"entries,required"`
	Name             param.Field[string]                                                       `json:"name,required"`
	AIContextEnabled param.Field[bool]                                                         `json:"ai_context_enabled"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount   param.Field[int64]  `json:"allowed_match_count"`
	ConfidenceThreshold param.Field[string] `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness param.Field[ContextAwarenessParam] `json:"context_awareness"`
	// The description of the profile.
	Description param.Field[string] `json:"description"`
	OCREnabled  param.Field[bool]   `json:"ocr_enabled"`
	// Entries from other profiles (e.g. pre-defined Cloudflare profiles, or your
	// Microsoft Information Protection profiles).
	SharedEntries param.Field[[]DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntryUnion] `json:"shared_entries"`
}

func (r DLPProfileCustomNewParamsBodyDLPNewCustomProfile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsBodyDLPNewCustomProfile) implementsDLPProfileCustomNewParamsBodyUnion() {
}

type DLPProfileCustomNewParamsBodyDLPNewCustomProfileEntry struct {
	Enabled param.Field[bool]         `json:"enabled,required"`
	Name    param.Field[string]       `json:"name,required"`
	Pattern param.Field[PatternParam] `json:"pattern"`
	Words   param.Field[interface{}]  `json:"words"`
}

func (r DLPProfileCustomNewParamsBodyDLPNewCustomProfileEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsBodyDLPNewCustomProfileEntry) implementsDLPProfileCustomNewParamsBodyDLPNewCustomProfileEntryUnion() {
}

// Satisfied by
// [zero_trust.DLPProfileCustomNewParamsBodyDLPNewCustomProfileEntriesDLPNewCustomEntry],
// [zero_trust.DLPProfileCustomNewParamsBodyDLPNewCustomProfileEntriesDLPNewWordListEntry],
// [DLPProfileCustomNewParamsBodyDLPNewCustomProfileEntry].
type DLPProfileCustomNewParamsBodyDLPNewCustomProfileEntryUnion interface {
	implementsDLPProfileCustomNewParamsBodyDLPNewCustomProfileEntryUnion()
}

type DLPProfileCustomNewParamsBodyDLPNewCustomProfileEntriesDLPNewCustomEntry struct {
	Enabled param.Field[bool]         `json:"enabled,required"`
	Name    param.Field[string]       `json:"name,required"`
	Pattern param.Field[PatternParam] `json:"pattern,required"`
}

func (r DLPProfileCustomNewParamsBodyDLPNewCustomProfileEntriesDLPNewCustomEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsBodyDLPNewCustomProfileEntriesDLPNewCustomEntry) implementsDLPProfileCustomNewParamsBodyDLPNewCustomProfileEntryUnion() {
}

type DLPProfileCustomNewParamsBodyDLPNewCustomProfileEntriesDLPNewWordListEntry struct {
	Enabled param.Field[bool]     `json:"enabled,required"`
	Name    param.Field[string]   `json:"name,required"`
	Words   param.Field[[]string] `json:"words,required"`
}

func (r DLPProfileCustomNewParamsBodyDLPNewCustomProfileEntriesDLPNewWordListEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsBodyDLPNewCustomProfileEntriesDLPNewWordListEntry) implementsDLPProfileCustomNewParamsBodyDLPNewCustomProfileEntryUnion() {
}

type DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntry struct {
	Enabled   param.Field[bool]                                                                   `json:"enabled,required"`
	EntryID   param.Field[string]                                                                 `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesEntryType] `json:"entry_type,required"`
}

func (r DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntry) implementsDLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntryUnion() {
}

// Satisfied by
// [zero_trust.DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesObject],
// [zero_trust.DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesObject],
// [zero_trust.DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesObject],
// [zero_trust.DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesObject],
// [DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntry].
type DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntryUnion interface {
	implementsDLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntryUnion()
}

type DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesObject struct {
	Enabled   param.Field[bool]                                                                         `json:"enabled,required"`
	EntryID   param.Field[string]                                                                       `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesObjectEntryType] `json:"entry_type,required"`
}

func (r DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesObject) implementsDLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntryUnion() {
}

type DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesObjectEntryType string

const (
	DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesObjectEntryTypeCustom DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesObjectEntryType = "custom"
)

func (r DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesObjectEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesObjectEntryTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesEntryType string

const (
	DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesEntryTypeCustom      DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesEntryType = "custom"
	DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesEntryTypePredefined  DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesEntryType = "predefined"
	DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesEntryTypeIntegration DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesEntryType = "integration"
	DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesEntryTypeExactData   DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesEntryType = "exact_data"
)

func (r DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesEntryTypeCustom, DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesEntryTypePredefined, DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesEntryTypeIntegration, DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesEntryTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseEnvelope struct {
	Errors   []DLPProfileCustomNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPProfileCustomNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPProfileCustomNewResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPProfileCustomNewResponseUnion           `json:"result"`
	JSON    dlpProfileCustomNewResponseEnvelopeJSON    `json:"-"`
}

// dlpProfileCustomNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPProfileCustomNewResponseEnvelope]
type dlpProfileCustomNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomNewResponseEnvelopeErrors struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           DLPProfileCustomNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpProfileCustomNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpProfileCustomNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPProfileCustomNewResponseEnvelopeErrors]
type dlpProfileCustomNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomNewResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    dlpProfileCustomNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpProfileCustomNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [DLPProfileCustomNewResponseEnvelopeErrorsSource]
type dlpProfileCustomNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomNewResponseEnvelopeMessages struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           DLPProfileCustomNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpProfileCustomNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpProfileCustomNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DLPProfileCustomNewResponseEnvelopeMessages]
type dlpProfileCustomNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    dlpProfileCustomNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpProfileCustomNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [DLPProfileCustomNewResponseEnvelopeMessagesSource]
type dlpProfileCustomNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPProfileCustomNewResponseEnvelopeSuccess bool

const (
	DLPProfileCustomNewResponseEnvelopeSuccessTrue DLPProfileCustomNewResponseEnvelopeSuccess = true
)

func (r DLPProfileCustomNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPProfileCustomUpdateParams struct {
	AccountID           param.Field[string] `path:"account_id,required"`
	Name                param.Field[string] `json:"name,required"`
	AIContextEnabled    param.Field[bool]   `json:"ai_context_enabled"`
	AllowedMatchCount   param.Field[int64]  `json:"allowed_match_count"`
	ConfidenceThreshold param.Field[string] `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness param.Field[ContextAwarenessParam] `json:"context_awareness"`
	// The description of the profile.
	Description param.Field[string] `json:"description"`
	// Custom entries from this profile. If this field is omitted, entries owned by
	// this profile will not be changed.
	Entries    param.Field[[]DLPProfileCustomUpdateParamsEntryUnion] `json:"entries"`
	OCREnabled param.Field[bool]                                     `json:"ocr_enabled"`
	// Other entries, e.g. predefined or integration.
	SharedEntries param.Field[[]DLPProfileCustomUpdateParamsSharedEntryUnion] `json:"shared_entries"`
}

func (r DLPProfileCustomUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPProfileCustomUpdateParamsEntry struct {
	Enabled param.Field[bool]         `json:"enabled,required"`
	Name    param.Field[string]       `json:"name,required"`
	Pattern param.Field[PatternParam] `json:"pattern,required"`
	EntryID param.Field[string]       `json:"entry_id" format:"uuid"`
}

func (r DLPProfileCustomUpdateParamsEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomUpdateParamsEntry) implementsDLPProfileCustomUpdateParamsEntryUnion() {}

// Satisfied by
// [zero_trust.DLPProfileCustomUpdateParamsEntriesDLPNewCustomEntryWithID],
// [zero_trust.DLPProfileCustomUpdateParamsEntriesDLPNewCustomEntry],
// [DLPProfileCustomUpdateParamsEntry].
type DLPProfileCustomUpdateParamsEntryUnion interface {
	implementsDLPProfileCustomUpdateParamsEntryUnion()
}

type DLPProfileCustomUpdateParamsEntriesDLPNewCustomEntryWithID struct {
	Enabled param.Field[bool]         `json:"enabled,required"`
	EntryID param.Field[string]       `json:"entry_id,required" format:"uuid"`
	Name    param.Field[string]       `json:"name,required"`
	Pattern param.Field[PatternParam] `json:"pattern,required"`
}

func (r DLPProfileCustomUpdateParamsEntriesDLPNewCustomEntryWithID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomUpdateParamsEntriesDLPNewCustomEntryWithID) implementsDLPProfileCustomUpdateParamsEntryUnion() {
}

type DLPProfileCustomUpdateParamsEntriesDLPNewCustomEntry struct {
	Enabled param.Field[bool]         `json:"enabled,required"`
	Name    param.Field[string]       `json:"name,required"`
	Pattern param.Field[PatternParam] `json:"pattern,required"`
}

func (r DLPProfileCustomUpdateParamsEntriesDLPNewCustomEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomUpdateParamsEntriesDLPNewCustomEntry) implementsDLPProfileCustomUpdateParamsEntryUnion() {
}

type DLPProfileCustomUpdateParamsSharedEntry struct {
	Enabled   param.Field[bool]                                               `json:"enabled,required"`
	EntryID   param.Field[string]                                             `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[DLPProfileCustomUpdateParamsSharedEntriesEntryType] `json:"entry_type,required"`
}

func (r DLPProfileCustomUpdateParamsSharedEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomUpdateParamsSharedEntry) implementsDLPProfileCustomUpdateParamsSharedEntryUnion() {
}

// Satisfied by [zero_trust.DLPProfileCustomUpdateParamsSharedEntriesObject],
// [zero_trust.DLPProfileCustomUpdateParamsSharedEntriesObject],
// [zero_trust.DLPProfileCustomUpdateParamsSharedEntriesObject],
// [DLPProfileCustomUpdateParamsSharedEntry].
type DLPProfileCustomUpdateParamsSharedEntryUnion interface {
	implementsDLPProfileCustomUpdateParamsSharedEntryUnion()
}

type DLPProfileCustomUpdateParamsSharedEntriesObject struct {
	Enabled   param.Field[bool]                                                     `json:"enabled,required"`
	EntryID   param.Field[string]                                                   `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[DLPProfileCustomUpdateParamsSharedEntriesObjectEntryType] `json:"entry_type,required"`
}

func (r DLPProfileCustomUpdateParamsSharedEntriesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomUpdateParamsSharedEntriesObject) implementsDLPProfileCustomUpdateParamsSharedEntryUnion() {
}

type DLPProfileCustomUpdateParamsSharedEntriesObjectEntryType string

const (
	DLPProfileCustomUpdateParamsSharedEntriesObjectEntryTypePredefined DLPProfileCustomUpdateParamsSharedEntriesObjectEntryType = "predefined"
)

func (r DLPProfileCustomUpdateParamsSharedEntriesObjectEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateParamsSharedEntriesObjectEntryTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomUpdateParamsSharedEntriesEntryType string

const (
	DLPProfileCustomUpdateParamsSharedEntriesEntryTypePredefined  DLPProfileCustomUpdateParamsSharedEntriesEntryType = "predefined"
	DLPProfileCustomUpdateParamsSharedEntriesEntryTypeIntegration DLPProfileCustomUpdateParamsSharedEntriesEntryType = "integration"
	DLPProfileCustomUpdateParamsSharedEntriesEntryTypeExactData   DLPProfileCustomUpdateParamsSharedEntriesEntryType = "exact_data"
)

func (r DLPProfileCustomUpdateParamsSharedEntriesEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateParamsSharedEntriesEntryTypePredefined, DLPProfileCustomUpdateParamsSharedEntriesEntryTypeIntegration, DLPProfileCustomUpdateParamsSharedEntriesEntryTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseEnvelope struct {
	Errors   []DLPProfileCustomUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPProfileCustomUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPProfileCustomUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  Profile                                       `json:"result"`
	JSON    dlpProfileCustomUpdateResponseEnvelopeJSON    `json:"-"`
}

// dlpProfileCustomUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPProfileCustomUpdateResponseEnvelope]
type dlpProfileCustomUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomUpdateResponseEnvelopeErrors struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           DLPProfileCustomUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpProfileCustomUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpProfileCustomUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DLPProfileCustomUpdateResponseEnvelopeErrors]
type dlpProfileCustomUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    dlpProfileCustomUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DLPProfileCustomUpdateResponseEnvelopeErrorsSource]
type dlpProfileCustomUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomUpdateResponseEnvelopeMessages struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           DLPProfileCustomUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpProfileCustomUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpProfileCustomUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DLPProfileCustomUpdateResponseEnvelopeMessages]
type dlpProfileCustomUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    dlpProfileCustomUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [DLPProfileCustomUpdateResponseEnvelopeMessagesSource]
type dlpProfileCustomUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPProfileCustomUpdateResponseEnvelopeSuccess bool

const (
	DLPProfileCustomUpdateResponseEnvelopeSuccessTrue DLPProfileCustomUpdateResponseEnvelopeSuccess = true
)

func (r DLPProfileCustomUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPProfileCustomDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPProfileCustomDeleteResponseEnvelope struct {
	Errors   []DLPProfileCustomDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPProfileCustomDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPProfileCustomDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPProfileCustomDeleteResponse                `json:"result,nullable"`
	JSON    dlpProfileCustomDeleteResponseEnvelopeJSON    `json:"-"`
}

// dlpProfileCustomDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPProfileCustomDeleteResponseEnvelope]
type dlpProfileCustomDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomDeleteResponseEnvelopeErrors struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           DLPProfileCustomDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpProfileCustomDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpProfileCustomDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DLPProfileCustomDeleteResponseEnvelopeErrors]
type dlpProfileCustomDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfileCustomDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    dlpProfileCustomDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpProfileCustomDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DLPProfileCustomDeleteResponseEnvelopeErrorsSource]
type dlpProfileCustomDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomDeleteResponseEnvelopeMessages struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           DLPProfileCustomDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpProfileCustomDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpProfileCustomDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DLPProfileCustomDeleteResponseEnvelopeMessages]
type dlpProfileCustomDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfileCustomDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    dlpProfileCustomDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpProfileCustomDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [DLPProfileCustomDeleteResponseEnvelopeMessagesSource]
type dlpProfileCustomDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPProfileCustomDeleteResponseEnvelopeSuccess bool

const (
	DLPProfileCustomDeleteResponseEnvelopeSuccessTrue DLPProfileCustomDeleteResponseEnvelopeSuccess = true
)

func (r DLPProfileCustomDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPProfileCustomDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPProfileCustomGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPProfileCustomGetResponseEnvelope struct {
	Errors   []DLPProfileCustomGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPProfileCustomGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPProfileCustomGetResponseEnvelopeSuccess `json:"success,required"`
	Result  Profile                                    `json:"result"`
	JSON    dlpProfileCustomGetResponseEnvelopeJSON    `json:"-"`
}

// dlpProfileCustomGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPProfileCustomGetResponseEnvelope]
type dlpProfileCustomGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomGetResponseEnvelopeErrors struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           DLPProfileCustomGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpProfileCustomGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpProfileCustomGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPProfileCustomGetResponseEnvelopeErrors]
type dlpProfileCustomGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomGetResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    dlpProfileCustomGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpProfileCustomGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [DLPProfileCustomGetResponseEnvelopeErrorsSource]
type dlpProfileCustomGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomGetResponseEnvelopeMessages struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           DLPProfileCustomGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpProfileCustomGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpProfileCustomGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DLPProfileCustomGetResponseEnvelopeMessages]
type dlpProfileCustomGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    dlpProfileCustomGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpProfileCustomGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [DLPProfileCustomGetResponseEnvelopeMessagesSource]
type dlpProfileCustomGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPProfileCustomGetResponseEnvelopeSuccess bool

const (
	DLPProfileCustomGetResponseEnvelopeSuccessTrue DLPProfileCustomGetResponseEnvelopeSuccess = true
)

func (r DLPProfileCustomGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
