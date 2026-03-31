// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package website exports the static content as an embed.FS.
package website

import (
	"embed"
	"io/fs"
)

// Content returns the go.dev website's static content,
// overlaying Vietnamese translations from _content_vi on top of _content.
func Content() fs.FS {
	vi := subdir(embeddedVI, "_content_vi")
	en := subdir(embedded, "_content")
	return NewOverlayFS(vi, en)
}

// TourOnly returns the content needed only for the standalone tour.
func TourOnly() fs.FS {
	return subdir(tourOnly, "_content")
}

// NewOverlayFS returns a filesystem that tries overlay first, falling back to base
// for files not present in overlay.
func NewOverlayFS(overlay, base fs.FS) fs.FS {
	return &overlayFS{overlay, base}
}

type overlayFS struct {
	overlay fs.FS
	base    fs.FS
}

func (o *overlayFS) Open(name string) (fs.File, error) {
	f, err := o.overlay.Open(name)
	if err == nil {
		return f, nil
	}
	return o.base.Open(name)
}

//go:embed _content
var embedded embed.FS

//go:embed _content_vi
var embeddedVI embed.FS

//go:embed _content/favicon.ico
//go:embed _content/images/go-logo-white.svg
//go:embed _content/images/icons
//go:embed _content/js/playground.js
//go:embed _content/tour
var tourOnly embed.FS

func subdir(fsys fs.FS, path string) fs.FS {
	s, err := fs.Sub(fsys, path)
	if err != nil {
		panic(err)
	}
	return s
}
