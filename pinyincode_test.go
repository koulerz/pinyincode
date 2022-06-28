package pinyincode

import "testing"

func TestCodeOnlyPinyin(t *testing.T) {
	var hans = "我-是_中Z国 人r123."
	var piny = "wszgr"
	if res := Code(hans, FlagOnlyPinyin); res != piny {
		t.Errorf("[%s] pinyin excepted [%s] but [%s]", hans, piny, res)
	}
}

func TestCodeWithEnglish(t *testing.T) {
	var hans = "我-是_中Z国 人r123."
	var piny = "wszZgrr"
	if res := Code(hans, FlagWithEnglish); res != piny {
		t.Errorf("[%s] pinyin excepted [%s] but [%s]", hans, piny, res)
	}
}

func TestCodeWithNumber(t *testing.T) {
	var hans = "我-是_中Z国 人r123."
	var piny = "wszgr123"
	if res := Code(hans, FlagWithNumber); res != piny {
		t.Errorf("[%s] pinyin excepted [%s] but [%s]", hans, piny, res)
	}
}

func TestCodeWithHyphen(t *testing.T) {
	var hans = "我-是_中Z国 人r123."
	var piny = "w-szgr"
	if res := Code(hans, FlagWithHyphen); res != piny {
		t.Errorf("[%s] pinyin excepted [%s] but [%s]", hans, piny, res)
	}
}

func TestCodeWithUnderscore(t *testing.T) {
	var hans = "我-是_中Z国 人r123."
	var piny = "ws_zgr"
	if res := Code(hans, FlagWithUnderscore); res != piny {
		t.Errorf("[%s] pinyin excepted [%s] but [%s]", hans, piny, res)
	}
}

func TestCodeWithEnglishAndNumber(t *testing.T) {
	var hans = "我-是_中Z国 人r123."
	var piny = "wszZgrr123"
	if res := Code(hans, FlagWithEnglish|FlagWithNumber); res != piny {
		t.Errorf("[%s] pinyin excepted [%s] but [%s]", hans, piny, res)
	}
}

func TestCodeWithEnglishAndHyphen(t *testing.T) {
	var hans = "我-是_中Z国 人r123."
	var piny = "w-szZgrr"
	if res := Code(hans, FlagWithEnglish|FlagWithHyphen); res != piny {
		t.Errorf("[%s] pinyin excepted [%s] but [%s]", hans, piny, res)
	}
}

func TestCodeWithNumberAndUnderscore(t *testing.T) {
	var hans = "我-是_中Z国 人r123."
	var piny = "ws_zgr123"
	if res := Code(hans, FlagWithNumber|FlagWithUnderscore); res != piny {
		t.Errorf("[%s] pinyin excepted [%s] but [%s]", hans, piny, res)
	}
}

func TestCodeWithAll(t *testing.T) {
	var hans = "我-是_中Z国 人r123."
	var piny = "wszgr"
	if res := Code(hans, FlagOnlyPinyin|FlagWithEnglish|FlagWithNumber|FlagWithHyphen|FlagWithUnderscore); res != piny {
		t.Errorf("[%s] pinyin excepted [%s] but [%s]", hans, piny, res)
	}
}
