package i18n

import "testing"

// ResetTranslations resets the translations map for testing purposes
func ResetTranslations() {
	translations = map[string]map[string]map[string]string{}
}

func TestAddTranslations(t *testing.T) {
	ResetTranslations() // Reset the translations map before the test

	AddTranslations("dashboard", map[string]map[string]string{
		"add_task": {
			"en": "Add Task",
			"ml": "ടാസ്ക് ചേർക്കുക",
		},
		"clear_all": {
			"en": "Clear All",
			"ml": "എല്ലാം മായ്ക്കുക",
		},
	})

	if len(translations["dashboard"]) != 2 {
		t.Errorf("Expected 2 translations, got: %d", len(translations["dashboard"]))
	}
}

func TestGetTranslation(t *testing.T) {
	ResetTranslations() // Reset the translations map before the test

	AddTranslations("dashboard", map[string]map[string]string{
		"add_task": {
			"en": "Add Task",
			"ml": "ടാസ്ക് ചേർക്കുക",
		},
	})

	page := "dashboard"
	key := "add_task"
	lang := "en"

	got := GetTranslation(page, key, lang)

	if got != "Add Task" {
		t.Errorf("Expected 'Add Task', got: %s", got)
	}

	lang = "ml"

	got = GetTranslation(page, key, lang)

	if got != "ടാസ്ക് ചേർക്കുക" {
		t.Errorf("Expected 'ടാസ്ക് ചേർക്കുക', got: %s", got)
	}
}
