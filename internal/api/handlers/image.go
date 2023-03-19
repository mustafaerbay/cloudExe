package handlers

import "mime/multipart"

type ImageHandler struct {
	Image *multipart.FileHeader
}

