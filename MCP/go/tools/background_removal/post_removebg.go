package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/background-removal-api/mcp-server/config"
	"github.com/background-removal-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Post_removebgHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.RemoveBgJson
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/removebg", cfg.BaseURL)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Fallback to single auth parameter
		if cfg.APIKey != "" {
			req.Header.Set("X-API-Key", cfg.APIKey)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.RemoveBgJsonResponse
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreatePost_removebgTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_removebg",
		mcp.WithDescription("Remove the background of an image"),
		mcp.WithString("image_url", mcp.Description("Input parameter: Source image URL. (If this parameter is present, the other image source parameters must be empty.)")),
		mcp.WithBoolean("semitransparency", mcp.Description("Input parameter: Whether to have semi-transparent regions in the result (default: true). NOTE: Semitransparency is currently only supported for car windows (this might change in the future). Other objects are returned without semitransparency, even if set to true.\n")),
		mcp.WithString("position", mcp.Description("Input parameter: Positions the subject within the image canvas. Can be \"original\" (default unless \"scale\" is given), \"center\" (default when \"scale\" is given) or a value from \"0%\" to \"100%\" (both horizontal and vertical) or two values (horizontal, vertical).\n")),
		mcp.WithString("type", mcp.Description("Input parameter: Foreground type:\n\"auto\" = Automatically detect kind of foreground,\n\"person\" = Use person(s) as foreground,\n\"product\" = Use product(s) as foreground.\n\"car\" = Use car as foreground,\n")),
		mcp.WithString("bg_color", mcp.Description("Input parameter: Adds a solid color background. Can be a hex color code (e.g. 81d4fa, fff) or a color name (e.g. green). For semi-transparency, 4-/8-digit hex codes are also supported (e.g. 81d4fa77). (If this parameter is present, the other bg_ parameters must be empty.)\n")),
		mcp.WithString("type_level", mcp.Description("Input parameter: Classification level of the detected foreground type:\n\"none\" = No classification (X-Type Header won't bet set on the response)\n\"1\" = Use coarse classification classes: [person, product, animal, car, other]\n\"2\" = Use more specific classification classes: [person, product, animal, car, car_interior, car_part, transportation, graphics, other]\n\"latest\" = Always use the latest classification classes available\n")),
		mcp.WithString("crop_margin", mcp.Description("Input parameter: Adds a margin around the cropped subject (default: 0). Can be an absolute value (e.g. \"30px\") or relative to the subject size (e.g. \"10%\"). Can be a single value (all sides), two values (top/bottom and left/right) or four values (top, right, bottom, left). This parameter only has an effect when \"crop=true\". The maximum margin that can be added on each side is 50% of the subject dimensions or 500 pixels.\n")),
		mcp.WithString("roi", mcp.Description("Input parameter: Region of interest: Only contents of this rectangular region can be detected as foreground. Everything outside is considered background and will be removed. The rectangle is defined as two x/y coordinates in the format \"x1 y1 x2 y2\". The coordinates can be in absolute pixels (suffix 'px') or relative to the width/height of the image (suffix '%'). By default, the whole image is the region of interest (\"0% 0% 100% 100%\").\n")),
		mcp.WithBoolean("crop", mcp.Description("Input parameter: Whether to crop off all empty regions (default: false). Note that cropping has no effect on the amount of charged credits.\n")),
		mcp.WithString("format", mcp.Description("Input parameter: Result image format:\n\"auto\" = Use PNG format if transparent regions exist, otherwise use JPG format (default),\n\"png\" = PNG format with alpha transparency,\n\"jpg\" = JPG format, no transparency,\n\"zip\" = ZIP format, contains color image and alpha matte image, supports transparency (recommended).\n")),
		mcp.WithString("scale", mcp.Description("Input parameter: Scales the subject relative to the total image size. Can be any value from \"10%\" to \"100%\", or \"original\" (default). Scaling the subject implies \"position=center\" (unless specified otherwise).\n")),
		mcp.WithString("image_file_b64", mcp.Description("Input parameter: Source image file (base64-encoded string). (If this parameter is present, the other image source parameters must be empty.)")),
		mcp.WithString("size", mcp.Description("Input parameter: Maximum output image resolution:\n\"preview\" (default) = Resize image to 0.25 megapixels (e.g. 625×400 pixels) – 0.25 credits per image,\n\"full\" = Use original image resolution, up to 25 megapixels (e.g. 6250x4000) with formats ZIP or JPG, or up to 10 megapixels (e.g. 4000x2500) with PNG – 1 credit per image),\n\"auto\" = Use highest available resolution (based on image size and available credits).\n\nFor backwards-compatibility this parameter also accepts the values \"medium\" (up to 1.5 megapixels) and \"hd\" (up to 4 megapixels) for 1 credit per image. The value \"full\" is also available under the name \"4k\" and the value \"preview\" is aliased as \"small\" and \"regular\".\n")),
		mcp.WithString("channels", mcp.Description("Input parameter: Request either the finalized image (\"rgba\", default) or an alpha mask (\"alpha\"). Note: Since remove.bg also applies RGB color corrections on edges, using only the alpha mask often leads to a lower final image quality. Therefore \"rgba\" is recommended.\n")),
		mcp.WithString("bg_image_url", mcp.Description("Input parameter: Adds a background image from a URL. The image is centered and resized to fill the canvas while preserving the aspect ratio, unless it already has the exact same dimensions as the foreground image. (If this parameter is present, the other bg_ parameters must be empty.)")),
		mcp.WithBoolean("add_shadow", mcp.Description("Input parameter: Whether to add an artificial shadow to the result (default: false). NOTE: Adding shadows is currently only supported for car photos. Other subjects are returned without shadow, even if set to true (this might change in the future).\n")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Post_removebgHandler(cfg),
	}
}
