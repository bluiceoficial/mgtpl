# MGTPL

<!-- Badge opcional para deixar claro visualmente -->
![Status](https://img.shields.io/badge/status-arquivado--archived-red.svg)

> **Aviso importante:** Esta biblioteca foi **oficialmente arquivada** e não receberá mais atualizações, correções de bugs ou suporte para novas versões.

MGTPL é uma biblioteca leve e prática para gerar HTML dinamicamente em Go. Ela permite criar tags HTML, aninhar elementos, trabalhar com tags que não fecham, atributos sem valor e gerar páginas completas de forma programática.

---

## ✨ Recursos

* Criação de **qualquer tag HTML** com conteúdo e atributos
* Suporte a **tags que não fecham** (img, br, meta, entre outras)
* Suporte a **atributos sem valor** (checked, disabled,readonly, entre outros)
* Geração automática de **DOCTYPE HTML5**
* Extensível: permite adicionar novas tags e atributos customizados
* Criação de páginas completas e aninhamento de tags

---

## 📦 Instalação

Dentro do seu projeto Go:

```bash
go get github.com/mugomes/mgtpl
```

---

## 🚀 Uso Básico

### Criando uma instância do template

```go
tpl := mgtpl.New()
```

Cria uma instância do MGTPL, necessária para gerar tags HTML dinamicamente.

---

### Adicionando uma tag que não precisa fechar

```go
tpl.AddNotClose("customtag")
```

Permite criar uma tag personalizada que não precisa de fechamento automático. Por exemplo: `<customtag>`.

---

### Adicionando um atributo sem valor

```go
tpl.AddAttributeNoValue("myattribute")
```

Permite que certos atributos sejam adicionados sem valor. Por exemplo: `<customtag myattribute>`.

---

### Criando o DOCTYPE HTML5

```go
html := tpl.Doctype()
```

Gera automaticamente o DOCTYPE HTML5 (`<!DOCTYPE html>`), recomendado no início do HTML.

---

### Criando tags HTML com conteúdo e atributos

```go
htmlBody := tpl.Tag("body",
	tpl.Tag("h1", "Hello World!"),
	tpl.Tag("p", "Servidor de teste usando MGTPL"),
	tpl.Tag("a", map[string]string{"href": "#"}, "Link"),
	tpl.Tag("input", map[string]string{"type": "checkbox", "checked": ""}),
)
```

* `tpl.Tag(name, args...)` cria uma tag HTML.
* `args` podem ser strings (conteúdo) ou `map[string]string` (atributos).
* Suporta tags que não fecham e atributos sem valor.
* Permite **aninhamento de tags**.

---

### Criando o HTML completo

```go
html := fmt.Sprintf("%s%s", tpl.Doctype(),
	tpl.Tag("html", map[string]string{"lang": "pt-BR"},
		tpl.Tag("head",
			tpl.Tag("title", "Exemplo Completo MGTPL"),
			tpl.Tag("meta", map[string]string{"charset": "utf-8"}),
		),
		htmlBody,
	),
)
```

**Descrição:**
Combina o DOCTYPE e as tags HTML criadas para gerar uma página completa.

---

## 🌐 Exemplo com servidor HTTP

```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	tpl := mgtpl.New()

	// Adicionando tag e atributo customizados
	tpl.AddNotClose("customtag")
	tpl.AddAttributeNoValue("myattribute")

	// Criando HTML completo
	htmlBody := tpl.Tag("body",
		tpl.Tag("h1", "Bem-vindo ao MGTPL!"),
		tpl.Tag("p", "Página gerada dinamicamente."),
		tpl.Tag("a", map[string]string{"href": "#"}, "Link"),
		tpl.Tag("br"),
		tpl.Tag("input", map[string]string{"type": "checkbox", "checked": ""}),
		tpl.Tag("customtag", map[string]string{"myattribute": ""}),
		tpl.Tag("button", map[string]string{"onclick": "alert('Você clicou!')"}, "Clique aqui"),
	)

	html := fmt.Sprintf("%s%s", tpl.Doctype(),
		tpl.Tag("html", map[string]string{"lang": "pt-BR"},
			tpl.Tag("head",
				tpl.Tag("title", "Exemplo MGTPL com Servidor"),
				tpl.Tag("meta", map[string]string{"charset": "utf-8"}),
			),
			htmlBody,
		),
	)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(html))
})

log.Println("🚀 Servidor rodando em http://localhost:8000")
log.Fatal(http.ListenAndServe(":8000", nil))
```

* Cria um servidor HTTP que entrega o HTML gerado dinamicamente.
* Mostra uso de tags normais, tags que não fecham, atributos sem valor e tags personalizadas.
* Permite criar páginas completas com cabeçalho, corpo e elementos interativos.

---

## 👤 Autor

**Murilo Gomes Julio**

🔗 [https://mugomes.github.io](https://mugomes.github.io)

📺 [https://youtube.com/@mugomesoficial](https://youtube.com/@mugomesoficial)

---

## License

Copyright (c) 2026 Murilo Gomes Julio

Licensed under the [MIT](https://github.com/mugomes/mgtpl/blob/main/LICENSE) license.

All contributions to the MGTPL are subject to this license.
