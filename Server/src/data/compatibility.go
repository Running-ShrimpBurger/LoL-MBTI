package data

func GetGood(mbti string) string {
	Good := make(map[string]string)
	Good["entj"] = "isfp"
	Good["entp"] = "isfj"
	Good["intj"] = "esfp"
	Good["intp"] = "esfj"
	Good["estj"] = "infp"
	Good["esfj"] = "intp"
	Good["istj"] = "enfp"
	Good["isfj"] = "entp"
	Good["enfj"] = "istp"
	Good["enfp"] = "istj"
	Good["infj"] = "estp"
	Good["infp"] = "estj"
	Good["estp"] = "infj"
	Good["esfp"] = "intj"
	Good["istp"] = "enfj"
	Good["isfp"] = "entj"

	return Good[mbti]
}

func GetBad(mbti string) string {
	Bad := make(map[string]string)
	Bad["entj"] = "isfj"
	Bad["entp"] = "isfp"
	Bad["intj"] = "esfj"
	Bad["intp"] = "esfp"
	Bad["estj"] = "infj"
	Bad["esfj"] = "intj"
	Bad["istj"] = "enfj"
	Bad["isfj"] = "entj"
	Bad["enfj"] = "istj"
	Bad["enfp"] = "istp"
	Bad["infj"] = "estj"
	Bad["infp"] = "estp"
	Bad["estp"] = "infp"
	Bad["esfp"] = "intp"
	Bad["istp"] = "enfp"
	Bad["isfp"] = "entp"

	return Bad[mbti]
}
