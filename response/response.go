package response

import "github.com/state-alchemists/zaruba/monitor"

func ShowSorryResponse(logger monitor.Logger, decoration *monitor.Decoration) {
	logger.DPrintf("%s%sDon't worry 👌%s, everyone makes mistakes\n", decoration.Bold, decoration.Yellow, decoration.Normal)
}

func ShowThanksResponse(logger monitor.Logger, decoration *monitor.Decoration) {
	logger.DPrintf("%s%sYour welcome 😊%s\n", decoration.Bold, decoration.Yellow, decoration.Normal)
	logger.DPrintf("Please consider donating ☕☕☕ to:\n")
	logger.DPrintf("%shttps://paypal.me/gofrendi%s\n", decoration.Yellow, decoration.Normal)
	logger.DPrintf("Also, follow Zaruba at 🐤 %shttps://twitter.com/zarubastalchmst%s\n", decoration.Yellow, decoration.Normal)
}

func ShowPleaseResponse(logger monitor.Logger, decoration *monitor.Decoration) {
	logger.DPrintf("%sPlease what?%s\n", decoration.Bold, decoration.Normal)
	logger.DPrintf("Here are several things you can try:\n")
	logger.DPrintf("* %szaruba please explain task %s%s[task-keyword]%s\n", decoration.Yellow, decoration.Normal, decoration.Blue, decoration.Normal)
	logger.DPrintf("* %szaruba please explain input %s%s[input-keyword]%s\n", decoration.Yellow, decoration.Normal, decoration.Blue, decoration.Normal)
	logger.DPrintf("* %szaruba please explain %s%s[task-or-input-keyword]%s\n", decoration.Yellow, decoration.Normal, decoration.Blue, decoration.Normal)
}
