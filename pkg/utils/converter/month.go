package converter

import constants "warabiz/api/pkg/constants/general"

func ConvertMonthtoRoman(month int) string {
	switch month {
	case constants.NumJan:
		return constants.RomanJan
	case constants.NumFeb:
		return constants.RomanFeb
	case constants.NumMar:
		return constants.RomanMar
	case constants.NumApr:
		return constants.RomanApr
	case constants.NumMay:
		return constants.RomanMay
	case constants.NumJune:
		return constants.RomanJune
	case constants.NumJuly:
		return constants.RomanJuly
	case constants.NumAug:
		return constants.RomanAug
	case constants.NumSep:
		return constants.RomanSep
	case constants.NumOct:
		return constants.RomanOct
	case constants.NumNov:
		return constants.RomanNov
	case constants.NumDec:
		return constants.RomanDec
	}
	return ""
}

func ConvertMonthtoStringEng(month int) string {
	switch month {
	case constants.NumJan:
		return constants.MonthJan
	case constants.NumFeb:
		return constants.MonthFeb
	case constants.NumMar:
		return constants.MonthMar
	case constants.NumApr:
		return constants.MonthApr
	case constants.NumMay:
		return constants.MonthMay
	case constants.NumJune:
		return constants.MonthJune
	case constants.NumJuly:
		return constants.MonthJuly
	case constants.NumAug:
		return constants.MonthAug
	case constants.NumSep:
		return constants.MonthSep
	case constants.NumOct:
		return constants.MonthOct
	case constants.NumNov:
		return constants.MonthNov
	case constants.NumDec:
		return constants.MonthDec
	}
	return ""
}

func ConvertMonthtoStringInd(month int) string {
	switch month {
	case constants.NumJan:
		return constants.BulanJan
	case constants.NumFeb:
		return constants.BulanFeb
	case constants.NumMar:
		return constants.BulanMar
	case constants.NumApr:
		return constants.BulanApr
	case constants.NumMay:
		return constants.BulanMay
	case constants.NumJune:
		return constants.BulanJune
	case constants.NumJuly:
		return constants.BulanJuly
	case constants.NumAug:
		return constants.BulanAug
	case constants.NumSep:
		return constants.BulanSep
	case constants.NumOct:
		return constants.BulanOct
	case constants.NumNov:
		return constants.BulanNov
	case constants.NumDec:
		return constants.BulanDec
	}
	return ""
}
