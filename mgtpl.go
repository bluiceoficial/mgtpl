// Copyright (C) 2026 Murilo Gomes Julio
// SPDX-License-Identifier: MIT
//
// Site: https://mugomes.github.io

package mgtpl

import (
	"fmt"
	"slices"
	"strings"
)

type MGTPL struct {
	notClose         []string
	attributeNoValue []string
}

func New() *MGTPL {
	return &MGTPL{
		notClose: []string{
			"area", "base", "br", "col", "embed", "hr",
			"img", "input", "link", "meta", "param", "source", "track", "wbr",
		},
		attributeNoValue: []string{
			"async", "autofocus", "autoplay", "checked",
			"controls", "default", "defer", "disabled", "download", "hidden",
			"loop", "multiple", "muted", "novalidate", "open", "readonly",
			"required", "reversed", "selected",
		},
	}
}

// Adiciona novas tags que não fecham
func (t *MGTPL) AddNotClose(values ...string) {
	t.notClose = append(t.notClose, values...)
}

// Adiciona atributos sem valor
func (t *MGTPL) AddAttributeNoValue(values ...string) {
	t.attributeNoValue = append(t.attributeNoValue, values...)
}

// Tag cria qualquer elemento HTML
func (t *MGTPL) Tag(name string, args ...any) string {
	var content strings.Builder
	var attrs strings.Builder

	// Atributos
	for _, arg := range args {
		switch v := arg.(type) {

		case map[string]string:
			for key, val := range v {
				if t.procurarValor(key, t.attributeNoValue) {
					fmt.Fprintf(&attrs, " %s", key)
				} else {
					fmt.Fprintf(&attrs, " %s=\"%s\"", key, val)
				}
			}

		case string:
			content.WriteString(v)
		}
	}

	// Tags que não fecham
	if t.procurarValor(name, t.notClose) {
		if name == "img" {
			return "<" + name + attrs.String() + " />" + content.String()
		}
		return "<" + name + attrs.String() + ">" + content.String()
	}

	// Tags normais
	return "<" + name + attrs.String() + ">" +
		content.String() +
		"</" + name + ">"
}

// Encontrar valores nos slices
func (t *MGTPL) procurarValor(palavra string, itens []string) bool {
	return slices.Contains(itens, palavra)
}

// Doctype retorna o DOCTYPE HTML5
func (t *MGTPL) Doctype() string {
	return "<!DOCTYPE html>\n"
}
