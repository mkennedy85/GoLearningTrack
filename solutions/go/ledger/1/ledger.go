package ledger

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

var (
	ErrInvalidCurrency = errors.New("invalid currency")
	ErrInvalidLocale   = errors.New("invalid locale")
	ErrInvalidDate     = errors.New("invalid date")
)

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	if err := validateCurrency(currency); err != nil {
		return "", err
	}

	if err := validateLocale(locale); err != nil {
		return "", err
	}

	entriesCopy := append([]Entry(nil), entries...)
	sortEntries(entriesCopy)

	var b strings.Builder
	b.WriteString(formatHeader(locale))

	for _, entry := range entriesCopy {
		row, err := formatRow(currency, locale, entry)
		if err != nil {
			return "", err
		}

		b.WriteString(row)
	}

	return b.String(), nil
}

func validateCurrency(currency string) error {
	switch currency {
	case "USD", "EUR":
		return nil
	default:
		return ErrInvalidCurrency
	}
}

func validateLocale(locale string) error {
	switch locale {
	case "en-US", "nl-NL":
		return nil
	default:
		return ErrInvalidLocale
	}
}

func sortEntries(entries []Entry) {
	sort.Slice(entries, func(i, j int) bool {
		a, b := entries[i], entries[j]

		if a.Date != b.Date {
			return a.Date < b.Date
		}

		if a.Description != b.Description {
			return a.Description < b.Description
		}

		return a.Change < b.Change
	})
}

func formatHeader(locale string) string {
	switch locale {
	case "nl-NL":
		return fmt.Sprintf("%-10s | %-25s | %-13s\n", "Datum", "Omschrijving", "Verandering")
	default:
		return fmt.Sprintf("%-10s | %-25s | %-13s\n", "Date", "Description", "Change")
	}
}

func formatRow(currency, locale string, entry Entry) (string, error) {
	date, err := formatDate(locale, entry.Date)
	if err != nil {
		return "", err
	}

	description := formatDescription(entry.Description)

	change, err := formatChange(currency, locale, entry.Change)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%-10s | %-25s | %13s\n", date, description, change), nil
}

func formatDate(locale, date string) (string, error) {
	if len(date) != 10 || date[4] != '-' || date[7] != '-' {
		return "", ErrInvalidDate
	}

	year := date[:4]
	month := date[5:7]
	day := date[8:10]

	switch locale {
	case "nl-NL":
		return day + "-" + month + "-" + year, nil
	default:
		return month + "/" + day + "/" + year, nil
	}
}

func formatDescription(description string) string {
	if len(description) > 25 {
		return description[:22] + "..."
	}

	return description
}

func formatChange(currency, locale string, cents int) (string, error) {
	symbol := currencySymbol(currency)

	negative := cents < 0
	if negative {
		cents = -cents
	}

	whole := cents / 100
	decimal := cents % 100

	switch locale {
	case "nl-NL":
		amount := fmt.Sprintf("%s %s,%02d", symbol, groupThousands(strconv.Itoa(whole), "."), decimal)
		if negative {
			amount = strings.Replace(amount, symbol+" ", symbol+" -", 1)
		}
		return amount + " ", nil

	default:
		amount := fmt.Sprintf("%s%s.%02d", symbol, groupThousands(strconv.Itoa(whole), ","), decimal)
		if negative {
			return "(" + amount + ")", nil
		}
		return amount + " ", nil
	}
}

func currencySymbol(currency string) string {
	switch currency {
	case "EUR":
		return "€"
	default:
		return "$"
	}
}

func groupThousands(s, sep string) string {
	if len(s) <= 3 {
		return s
	}

	var parts []string
	for len(s) > 3 {
		parts = append(parts, s[len(s)-3:])
		s = s[:len(s)-3]
	}

	parts = append(parts, s)

	for i, j := 0, len(parts)-1; i < j; i, j = i+1, j-1 {
		parts[i], parts[j] = parts[j], parts[i]
	}

	return strings.Join(parts, sep)
}