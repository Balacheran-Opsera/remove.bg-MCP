package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// RemoveBgMultipart represents the RemoveBgMultipart schema from the OpenAPI specification
type RemoveBgMultipart struct {
	Size string `json:"size,omitempty"` // Maximum output image resolution: "preview" (default) = Resize image to 0.25 megapixels (e.g. 625×400 pixels) – 0.25 credits per image, "full" = Use original image resolution, up to 25 megapixels (e.g. 6250x4000) with formats ZIP or JPG, or up to 10 megapixels (e.g. 4000x2500) with PNG – 1 credit per image), "auto" = Use highest available resolution (based on image size and available credits). For backwards-compatibility this parameter also accepts the values "medium" (up to 1.5 megapixels) and "hd" (up to 4 megapixels) for 1 credit per image. The value "full" is also available under the name "4k" and the value "preview" is aliased as "small" and "regular".
	TypeField string `json:"type,omitempty"` // Foreground type: "auto" = Automatically detect kind of foreground, "person" = Use person(s) as foreground, "product" = Use product(s) as foreground. "car" = Use car as foreground,
	Image_file string `json:"image_file,omitempty"` // Source image file (binary). (If this parameter is present, the other image source parameters must be empty.)
	Image_file_b64 string `json:"image_file_b64,omitempty"` // Source image file (base64-encoded string). (If this parameter is present, the other image source parameters must be empty.)
	Position string `json:"position,omitempty"` // Positions the subject within the image canvas. Can be "original" (default unless "scale" is given), "center" (default when "scale" is given) or a value from "0%" to "100%" (both horizontal and vertical) or two values (horizontal, vertical).
	Semitransparency bool `json:"semitransparency,omitempty"` // Whether to have semi-transparent regions in the result (default: true). NOTE: Semitransparency is currently only supported for car windows (this might change in the future). Other objects are returned without semitransparency, even if set to true.
	Bg_image_url string `json:"bg_image_url,omitempty"` // Adds a background image from a URL. The image is centered and resized to fill the canvas while preserving the aspect ratio, unless it already has the exact same dimensions as the foreground image. (If this parameter is present, the other bg_ parameters must be empty.)
	Image_url string `json:"image_url,omitempty"` // Source image URL. (If this parameter is present, the other image source parameters must be empty.)
	Add_shadow bool `json:"add_shadow,omitempty"` // Whether to add an artificial shadow to the result (default: false). NOTE: Adding shadows is currently only supported for car photos. Other subjects are returned without shadow, even if set to true (this might change in the future).
	Bg_color string `json:"bg_color,omitempty"` // Adds a solid color background. Can be a hex color code (e.g. 81d4fa, fff) or a color name (e.g. green). For semi-transparency, 4-/8-digit hex codes are also supported (e.g. 81d4fa77). (If this parameter is present, the other bg_ parameters must be empty.)
	Crop bool `json:"crop,omitempty"` // Whether to crop off all empty regions (default: false). Note that cropping has no effect on the amount of charged credits.
	Format string `json:"format,omitempty"` // Result image format: "auto" = Use PNG format if transparent regions exist, otherwise use JPG format (default), "png" = PNG format with alpha transparency, "jpg" = JPG format, no transparency, "zip" = ZIP format, contains color image and alpha matte image, supports transparency (recommended).
	Scale string `json:"scale,omitempty"` // Scales the subject relative to the total image size. Can be any value from "10%" to "100%", or "original" (default). Scaling the subject implies "position=center" (unless specified otherwise).
	Roi string `json:"roi,omitempty"` // Region of interest: Only contents of this rectangular region can be detected as foreground. Everything outside is considered background and will be removed. The rectangle is defined as two x/y coordinates in the format "x1 y1 x2 y2". The coordinates can be in absolute pixels (suffix 'px') or relative to the width/height of the image (suffix '%'). By default, the whole image is the region of interest ("0% 0% 100% 100%").
	Type_level string `json:"type_level,omitempty"` // Classification level of the detected foreground type: "none" = No classification (X-Type Header won't bet set on the response) "1" = Use coarse classification classes: [person, product, animal, car, other] "2" = Use more specific classification classes: [person, product, animal, car, car_interior, car_part, transportation, graphics, other] "latest" = Always use the latest classification classes available
	Bg_image_file string `json:"bg_image_file,omitempty"` // Adds a background image from a file (binary). The image is centered and resized to fill the canvas while preserving the aspect ratio, unless it already has the exact same dimensions as the foreground image. (If this parameter is present, the other bg_ parameters must be empty.)
	Channels string `json:"channels,omitempty"` // Request either the finalized image ("rgba", default) or an alpha mask ("alpha"). Note: Since remove.bg also applies RGB color corrections on edges, using only the alpha mask often leads to a lower final image quality. Therefore "rgba" is recommended.
	Crop_margin string `json:"crop_margin,omitempty"` // Adds a margin around the cropped subject (default: 0). Can be an absolute value (e.g. "30px") or relative to the subject size (e.g. "10%"). Can be a single value (all sides), two values (top/bottom and left/right) or four values (top, right, bottom, left). This parameter only has an effect when "crop=true". The maximum margin that can be added on each side is 50% of the subject dimensions or 500 pixels.
}

// AuthFailed represents the AuthFailed schema from the OpenAPI specification
type AuthFailed struct {
	Errors []map[string]interface{} `json:"errors,omitempty"`
}

// ImprovementProgramJson represents the ImprovementProgramJson schema from the OpenAPI specification
type ImprovementProgramJson struct {
	Image_url string `json:"image_url,omitempty"` // Source image URL. (If this parameter is present, the other image source parameters must be empty.)
	Tag string `json:"tag,omitempty"` // Images with the same tag are grouped together.
	Image_file_b64 string `json:"image_file_b64,omitempty"` // Source image file (base64-encoded string). (If this parameter is present, the other image source parameters must be empty.)
	Image_filename string `json:"image_filename,omitempty"` // Filename of the image, if not provided it will be autodetected form the submitted data.
}

// ImprovementProgramJsonResponse represents the ImprovementProgramJsonResponse schema from the OpenAPI specification
type ImprovementProgramJsonResponse struct {
	Id string `json:"id,omitempty"` // ID of the submitted image
}

// ImprovementProgramMultipart represents the ImprovementProgramMultipart schema from the OpenAPI specification
type ImprovementProgramMultipart struct {
	Image_file_b64 string `json:"image_file_b64,omitempty"` // Source image file (base64-encoded string). (If this parameter is present, the other image source parameters must be empty.)
	Image_filename string `json:"image_filename,omitempty"` // Filename of the image, if not provided it will be autodetected form the submitted data.
	Image_url string `json:"image_url,omitempty"` // Source image URL. (If this parameter is present, the other image source parameters must be empty.)
	Tag string `json:"tag,omitempty"` // Images with the same tag are grouped together.
	Image_file string `json:"image_file,omitempty"` // Source image file (binary). (If this parameter is present, the other image source parameters must be empty.)
}

// RateLimit represents the RateLimit schema from the OpenAPI specification
type RateLimit struct {
	Errors []map[string]interface{} `json:"errors,omitempty"`
}

// RemoveBgJson represents the RemoveBgJson schema from the OpenAPI specification
type RemoveBgJson struct {
	Crop bool `json:"crop,omitempty"` // Whether to crop off all empty regions (default: false). Note that cropping has no effect on the amount of charged credits.
	Format string `json:"format,omitempty"` // Result image format: "auto" = Use PNG format if transparent regions exist, otherwise use JPG format (default), "png" = PNG format with alpha transparency, "jpg" = JPG format, no transparency, "zip" = ZIP format, contains color image and alpha matte image, supports transparency (recommended).
	Scale string `json:"scale,omitempty"` // Scales the subject relative to the total image size. Can be any value from "10%" to "100%", or "original" (default). Scaling the subject implies "position=center" (unless specified otherwise).
	Image_file_b64 string `json:"image_file_b64,omitempty"` // Source image file (base64-encoded string). (If this parameter is present, the other image source parameters must be empty.)
	Size string `json:"size,omitempty"` // Maximum output image resolution: "preview" (default) = Resize image to 0.25 megapixels (e.g. 625×400 pixels) – 0.25 credits per image, "full" = Use original image resolution, up to 25 megapixels (e.g. 6250x4000) with formats ZIP or JPG, or up to 10 megapixels (e.g. 4000x2500) with PNG – 1 credit per image), "auto" = Use highest available resolution (based on image size and available credits). For backwards-compatibility this parameter also accepts the values "medium" (up to 1.5 megapixels) and "hd" (up to 4 megapixels) for 1 credit per image. The value "full" is also available under the name "4k" and the value "preview" is aliased as "small" and "regular".
	Channels string `json:"channels,omitempty"` // Request either the finalized image ("rgba", default) or an alpha mask ("alpha"). Note: Since remove.bg also applies RGB color corrections on edges, using only the alpha mask often leads to a lower final image quality. Therefore "rgba" is recommended.
	Bg_image_url string `json:"bg_image_url,omitempty"` // Adds a background image from a URL. The image is centered and resized to fill the canvas while preserving the aspect ratio, unless it already has the exact same dimensions as the foreground image. (If this parameter is present, the other bg_ parameters must be empty.)
	Add_shadow bool `json:"add_shadow,omitempty"` // Whether to add an artificial shadow to the result (default: false). NOTE: Adding shadows is currently only supported for car photos. Other subjects are returned without shadow, even if set to true (this might change in the future).
	Image_url string `json:"image_url,omitempty"` // Source image URL. (If this parameter is present, the other image source parameters must be empty.)
	Semitransparency bool `json:"semitransparency,omitempty"` // Whether to have semi-transparent regions in the result (default: true). NOTE: Semitransparency is currently only supported for car windows (this might change in the future). Other objects are returned without semitransparency, even if set to true.
	Position string `json:"position,omitempty"` // Positions the subject within the image canvas. Can be "original" (default unless "scale" is given), "center" (default when "scale" is given) or a value from "0%" to "100%" (both horizontal and vertical) or two values (horizontal, vertical).
	TypeField string `json:"type,omitempty"` // Foreground type: "auto" = Automatically detect kind of foreground, "person" = Use person(s) as foreground, "product" = Use product(s) as foreground. "car" = Use car as foreground,
	Bg_color string `json:"bg_color,omitempty"` // Adds a solid color background. Can be a hex color code (e.g. 81d4fa, fff) or a color name (e.g. green). For semi-transparency, 4-/8-digit hex codes are also supported (e.g. 81d4fa77). (If this parameter is present, the other bg_ parameters must be empty.)
	Type_level string `json:"type_level,omitempty"` // Classification level of the detected foreground type: "none" = No classification (X-Type Header won't bet set on the response) "1" = Use coarse classification classes: [person, product, animal, car, other] "2" = Use more specific classification classes: [person, product, animal, car, car_interior, car_part, transportation, graphics, other] "latest" = Always use the latest classification classes available
	Crop_margin string `json:"crop_margin,omitempty"` // Adds a margin around the cropped subject (default: 0). Can be an absolute value (e.g. "30px") or relative to the subject size (e.g. "10%"). Can be a single value (all sides), two values (top/bottom and left/right) or four values (top, right, bottom, left). This parameter only has an effect when "crop=true". The maximum margin that can be added on each side is 50% of the subject dimensions or 500 pixels.
	Roi string `json:"roi,omitempty"` // Region of interest: Only contents of this rectangular region can be detected as foreground. Everything outside is considered background and will be removed. The rectangle is defined as two x/y coordinates in the format "x1 y1 x2 y2". The coordinates can be in absolute pixels (suffix 'px') or relative to the width/height of the image (suffix '%'). By default, the whole image is the region of interest ("0% 0% 100% 100%").
}

// RemoveBgJsonResponse represents the RemoveBgJsonResponse schema from the OpenAPI specification
type RemoveBgJsonResponse struct {
	Data map[string]interface{} `json:"data,omitempty"`
}
