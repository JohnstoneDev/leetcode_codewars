function ReverseWord(str){
	let words = str.split(" ")
	let reversedWords = []

	for (let word of words) {
		let reversed = word.split("")
		word = reversed.reverse().join("")
		reversedWords.push(word)
	}

	console.log(reversedWords.join(" ") == "ehT kciuq nworb xof spmuj revo eht yzal .god")

	return reversedWords.join(" ")
}


console.log(ReverseWord("The quick brown fox jumps over the lazy dog."), "<= Got, Expecting: ehT kciuq nworb xof spmuj revo eht yzal .god")