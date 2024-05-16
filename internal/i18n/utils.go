package i18n

var translations = map[string]map[string]map[string]string{}

func AddTranslations(page string, pageTranslations map[string]map[string]string) {
	if _, exists := translations[page]; !exists {
		translations[page] = pageTranslations
	} else {
		for key, langs := range pageTranslations {
			if _, exists := translations[page][key]; !exists {
				translations[page][key] = langs
			} else {
				for lang, text := range langs {
					translations[page][key][lang] = text
				}
			}
		}
	}
}

func GetTranslation(page, key, lang string) string {
	if pages, ok := translations[page]; ok {
		if keys, ok := pages[key]; ok {
			if text, ok := keys[lang]; ok {
				return text
			}
			return keys["en"]
		}
	}
	return ""
}
