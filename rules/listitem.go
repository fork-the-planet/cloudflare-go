// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rules

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// ListItemService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewListItemService] method instead.
type ListItemService struct {
	Options []option.RequestOption
}

// NewListItemService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewListItemService(opts ...option.RequestOption) (r *ListItemService) {
	r = &ListItemService{}
	r.Options = opts
	return
}

// Appends new items to the list.
//
// This operation is asynchronous. To get current the operation status, invoke the
// `Get bulk operation status` endpoint with the returned `operation_id`.
func (r *ListItemService) New(ctx context.Context, listID string, params ListItemNewParams, opts ...option.RequestOption) (res *ListItemNewResponse, err error) {
	var env ListItemNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if listID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rules/lists/%s/items", params.AccountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Removes all existing items from the list and adds the provided items to the
// list.
//
// This operation is asynchronous. To get current the operation status, invoke the
// `Get bulk operation status` endpoint with the returned `operation_id`.
func (r *ListItemService) Update(ctx context.Context, listID string, params ListItemUpdateParams, opts ...option.RequestOption) (res *ListItemUpdateResponse, err error) {
	var env ListItemUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if listID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rules/lists/%s/items", params.AccountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches all the items in the list.
func (r *ListItemService) List(ctx context.Context, listID string, params ListItemListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[ListItemListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if listID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rules/lists/%s/items", params.AccountID, listID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Fetches all the items in the list.
func (r *ListItemService) ListAutoPaging(ctx context.Context, listID string, params ListItemListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[ListItemListResponse] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, listID, params, opts...))
}

// Removes one or more items from a list.
//
// This operation is asynchronous. To get current the operation status, invoke the
// `Get bulk operation status` endpoint with the returned `operation_id`.
func (r *ListItemService) Delete(ctx context.Context, listID string, body ListItemDeleteParams, opts ...option.RequestOption) (res *ListItemDeleteResponse, err error) {
	var env ListItemDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if listID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rules/lists/%s/items", body.AccountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a list item in the list.
func (r *ListItemService) Get(ctx context.Context, listID string, itemID string, query ListItemGetParams, opts ...option.RequestOption) (res *ListItemGetResponse, err error) {
	var env ListItemGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if listID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
	if itemID == "" {
		err = errors.New("missing required item_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rules/lists/%s/items/%s", query.AccountID, listID, itemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ListCursor struct {
	After  string         `json:"after"`
	Before string         `json:"before"`
	JSON   listCursorJSON `json:"-"`
}

// listCursorJSON contains the JSON metadata for the struct [ListCursor]
type listCursorJSON struct {
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListCursor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listCursorJSON) RawJSON() string {
	return r.raw
}

type ListItemNewResponse struct {
	// The unique operation ID of the asynchronous action.
	OperationID string                  `json:"operation_id"`
	JSON        listItemNewResponseJSON `json:"-"`
}

// listItemNewResponseJSON contains the JSON metadata for the struct
// [ListItemNewResponse]
type listItemNewResponseJSON struct {
	OperationID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListItemNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemNewResponseJSON) RawJSON() string {
	return r.raw
}

type ListItemUpdateResponse struct {
	// The unique operation ID of the asynchronous action.
	OperationID string                     `json:"operation_id"`
	JSON        listItemUpdateResponseJSON `json:"-"`
}

// listItemUpdateResponseJSON contains the JSON metadata for the struct
// [ListItemUpdateResponse]
type listItemUpdateResponseJSON struct {
	OperationID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListItemUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ListItemListResponse struct {
	// The unique ID of the list.
	ID string `json:"id"`
	// Defines a non-negative 32 bit integer.
	ASN int64 `json:"asn"`
	// Defines an informative summary of the list item.
	Comment string `json:"comment"`
	// The RFC 3339 timestamp of when the item was created.
	CreatedOn string `json:"created_on"`
	// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
	// 0 to 9, wildcards (\*), and the hyphen (-).
	Hostname Hostname `json:"hostname"`
	// An IPv4 address, an IPv4 CIDR, an IPv6 address, or an IPv6 CIDR.
	IP string `json:"ip"`
	// The RFC 3339 timestamp of when the item was last modified.
	ModifiedOn string `json:"modified_on"`
	// The definition of the redirect.
	Redirect Redirect                 `json:"redirect"`
	JSON     listItemListResponseJSON `json:"-"`
}

// listItemListResponseJSON contains the JSON metadata for the struct
// [ListItemListResponse]
type listItemListResponseJSON struct {
	ID          apijson.Field
	ASN         apijson.Field
	Comment     apijson.Field
	CreatedOn   apijson.Field
	Hostname    apijson.Field
	IP          apijson.Field
	ModifiedOn  apijson.Field
	Redirect    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListItemListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemListResponseJSON) RawJSON() string {
	return r.raw
}

type ListItemDeleteResponse struct {
	// The unique operation ID of the asynchronous action.
	OperationID string                     `json:"operation_id"`
	JSON        listItemDeleteResponseJSON `json:"-"`
}

// listItemDeleteResponseJSON contains the JSON metadata for the struct
// [ListItemDeleteResponse]
type listItemDeleteResponseJSON struct {
	OperationID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListItemDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ListItemGetResponse struct {
	// The unique ID of the list.
	ID string `json:"id"`
	// Defines a non-negative 32 bit integer.
	ASN int64 `json:"asn"`
	// Defines an informative summary of the list item.
	Comment string `json:"comment"`
	// The RFC 3339 timestamp of when the item was created.
	CreatedOn string `json:"created_on"`
	// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
	// 0 to 9, wildcards (\*), and the hyphen (-).
	Hostname Hostname `json:"hostname"`
	// An IPv4 address, an IPv4 CIDR, an IPv6 address, or an IPv6 CIDR.
	IP string `json:"ip"`
	// The RFC 3339 timestamp of when the item was last modified.
	ModifiedOn string `json:"modified_on"`
	// The definition of the redirect.
	Redirect Redirect                `json:"redirect"`
	JSON     listItemGetResponseJSON `json:"-"`
}

// listItemGetResponseJSON contains the JSON metadata for the struct
// [ListItemGetResponse]
type listItemGetResponseJSON struct {
	ID          apijson.Field
	ASN         apijson.Field
	Comment     apijson.Field
	CreatedOn   apijson.Field
	Hostname    apijson.Field
	IP          apijson.Field
	ModifiedOn  apijson.Field
	Redirect    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListItemGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemGetResponseJSON) RawJSON() string {
	return r.raw
}

type ListItemNewParams struct {
	// Defines an identifier.
	AccountID param.Field[string]     `path:"account_id,required"`
	Body      []ListItemNewParamsBody `json:"body,required"`
}

func (r ListItemNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ListItemNewParamsBody struct {
	// Defines a non-negative 32 bit integer.
	ASN param.Field[int64] `json:"asn"`
	// Defines an informative summary of the list item.
	Comment param.Field[string] `json:"comment"`
	// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
	// 0 to 9, wildcards (\*), and the hyphen (-).
	Hostname param.Field[HostnameParam] `json:"hostname"`
	// An IPv4 address, an IPv4 CIDR, an IPv6 address, or an IPv6 CIDR.
	IP param.Field[string] `json:"ip"`
	// The definition of the redirect.
	Redirect param.Field[RedirectParam] `json:"redirect"`
}

func (r ListItemNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ListItemNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ListItemNewResponse   `json:"result,required"`
	// Defines whether the API call was successful.
	Success ListItemNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    listItemNewResponseEnvelopeJSON    `json:"-"`
}

// listItemNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ListItemNewResponseEnvelope]
type listItemNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListItemNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type ListItemNewResponseEnvelopeSuccess bool

const (
	ListItemNewResponseEnvelopeSuccessTrue ListItemNewResponseEnvelopeSuccess = true
)

func (r ListItemNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ListItemNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ListItemUpdateParams struct {
	// Defines an identifier.
	AccountID param.Field[string]        `path:"account_id,required"`
	Body      []ListItemUpdateParamsBody `json:"body,required"`
}

func (r ListItemUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ListItemUpdateParamsBody struct {
	// Defines a non-negative 32 bit integer.
	ASN param.Field[int64] `json:"asn"`
	// Defines an informative summary of the list item.
	Comment param.Field[string] `json:"comment"`
	// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
	// 0 to 9, wildcards (\*), and the hyphen (-).
	Hostname param.Field[HostnameParam] `json:"hostname"`
	// An IPv4 address, an IPv4 CIDR, an IPv6 address, or an IPv6 CIDR.
	IP param.Field[string] `json:"ip"`
	// The definition of the redirect.
	Redirect param.Field[RedirectParam] `json:"redirect"`
}

func (r ListItemUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ListItemUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo  `json:"errors,required"`
	Messages []shared.ResponseInfo  `json:"messages,required"`
	Result   ListItemUpdateResponse `json:"result,required"`
	// Defines whether the API call was successful.
	Success ListItemUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    listItemUpdateResponseEnvelopeJSON    `json:"-"`
}

// listItemUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [ListItemUpdateResponseEnvelope]
type listItemUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListItemUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type ListItemUpdateResponseEnvelopeSuccess bool

const (
	ListItemUpdateResponseEnvelopeSuccessTrue ListItemUpdateResponseEnvelopeSuccess = true
)

func (r ListItemUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ListItemUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ListItemListParams struct {
	// Defines an identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// The pagination cursor. An opaque string token indicating the position from which
	// to continue when requesting the next/previous set of records. Cursor values are
	// provided under `result_info.cursors` in the response. You should make no
	// assumptions about a cursor's content or length.
	Cursor param.Field[string] `query:"cursor"`
	// Amount of results to include in each paginated response. A non-negative 32 bit
	// integer.
	PerPage param.Field[int64] `query:"per_page"`
	// A search query to filter returned items. Its meaning depends on the list type:
	// IP addresses must start with the provided string, hostnames and bulk redirects
	// must contain the string, and ASNs must match the string exactly.
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [ListItemListParams]'s query parameters as `url.Values`.
func (r ListItemListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ListItemDeleteParams struct {
	// Defines an identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ListItemDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo  `json:"errors,required"`
	Messages []shared.ResponseInfo  `json:"messages,required"`
	Result   ListItemDeleteResponse `json:"result,required"`
	// Defines whether the API call was successful.
	Success ListItemDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    listItemDeleteResponseEnvelopeJSON    `json:"-"`
}

// listItemDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [ListItemDeleteResponseEnvelope]
type listItemDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListItemDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type ListItemDeleteResponseEnvelopeSuccess bool

const (
	ListItemDeleteResponseEnvelopeSuccessTrue ListItemDeleteResponseEnvelopeSuccess = true
)

func (r ListItemDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ListItemDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ListItemGetParams struct {
	// Defines an identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ListItemGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ListItemGetResponse   `json:"result,required"`
	// Defines whether the API call was successful.
	Success ListItemGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    listItemGetResponseEnvelopeJSON    `json:"-"`
}

// listItemGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ListItemGetResponseEnvelope]
type listItemGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListItemGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type ListItemGetResponseEnvelopeSuccess bool

const (
	ListItemGetResponseEnvelopeSuccessTrue ListItemGetResponseEnvelopeSuccess = true
)

func (r ListItemGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ListItemGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
