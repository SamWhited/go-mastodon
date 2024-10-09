package mastodon

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// AdminSuspend suspends an account.
func (c *Client) AdminSuspend(ctx context.Context, id ID) error {
	params := url.Values{}
	params.Add("type", "suspend")
	params.Add("text", "Suspended by anti-spam bot")

	return c.doAPI(ctx, http.MethodPost, fmt.Sprintf("/api/v1/admin/accounts/%s/action", url.PathEscape(string(id))), params, nil, nil)
}
