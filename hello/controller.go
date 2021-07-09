// Copyright (c) 2021 Thomas Junk
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package hello

import "net/http"

func (h *HelloWorld) Welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}
