package utils

import (
	"bytes"
	"unicode/utf8"
)

var (
	// Mang cac ky tu goc co dau var
	SOURCE_CHARACTERS, LL_LENGTH = stringToRune(`ÀÁÂÃÈÉÊÌÍÒÓÔÕÙÚÝàáâãèéêìíòóôõùúýĂăĐđĨĩŨũƠơƯưẠạẢảẤấẦầẨẩẪẫẬậẮắẰằẲẳẴẵẶặẸẹẺẻẼẽẾếỀềỂểỄễỆệỈỉỊịỌọỎỏỐốỒồỔổỖỗỘộỚớỜờỞởỠỡỢợỤụỦủỨứỪừỬửỮữỰự`)
	// Mang cac ky tu thay the khong dau var
	DESTINATION_CHARACTERS, _ = stringToRune(`AAAAEEEIIOOOOUUYaaaaeeeiioooouuyAaDdIiUuOoUuAaAaAaAaAaAaAaAaAaAaAaAaEeEeEeEeEeEeEeEeIiIiOoOoOoOoOoOoOoOoOoOoOoOoUuUuUuUuUuUuUu`)
)

func stringToRune(s string) ([]string, int) {

	ll := utf8.RuneCountInString(s)

	var texts = make([]string, ll+1)

	var index = 0

	for _, runeValue := range s {

		texts[index] = string(runeValue)

		index++

	}

	return texts, ll

}

func binarySearch(sortedArray []string, key string, low int, high int) int {

	var middle int = (low + high) / 2

	if high < low {
		return -1
	}

	if key == sortedArray[middle] {

		return middle

	} else if key < sortedArray[middle] {

		return binarySearch(sortedArray, key, low, middle-1)

	} else {

		return binarySearch(sortedArray, key, middle+1, high)

	}

}

func removeAccentChar(ch string) string {

	var index int = binarySearch(SOURCE_CHARACTERS, ch, 0, LL_LENGTH)

	if index >= 0 {

		ch = DESTINATION_CHARACTERS[index]
	}
	return ch

}

func RemoveAccent(s string) string {
	var buffer bytes.Buffer

	for _, runeValue := range s {
		buffer.WriteString(removeAccentChar(string(runeValue)))
	}

	return buffer.String()
}
