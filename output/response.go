package output

func ShowSorryResponse(logger Logger, decoration *Decoration) {
	logger.DPrintf("%s%sDon't worry 👌%s, everyone makes mistakes\n", decoration.Bold, decoration.Yellow, decoration.Normal)
}

func ShowThanksResponse(logger Logger, decoration *Decoration) {
	logger.DPrintf("%s%sYour welcome 😊%s\n", decoration.Bold, decoration.Yellow, decoration.Normal)
	logger.DPrintf("Please consider donating ☕☕☕ to:\n")
	logger.DPrintf("%shttps://paypal.me/gofrendi%s\n", decoration.Yellow, decoration.Normal)
	logger.DPrintf("Also, follow Zaruba at 🐤 %shttps://twitter.com/zarubastalchmst%s\n", decoration.Yellow, decoration.Normal)
}
