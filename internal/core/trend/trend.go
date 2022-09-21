package trend

type FindingStrategy int

const (
	LikeCountFindingStrategy FindingStrategy = 1
)

const MinimumLikeCountToAcceptAsTrendForLikeCountFindingStrategy = 1

func IsTrendAsLikeCountFindingStrategy(likeCount int) bool {
	return likeCount >= MinimumLikeCountToAcceptAsTrendForLikeCountFindingStrategy
}

func GetAlgorithmOfFindingStrategy() {

}
