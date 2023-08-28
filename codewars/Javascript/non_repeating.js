/*
	Program that takes a string input & returns the first character that
	is not repeated anywhere in the string, Case insensitive while checking
	Case sensitive while returning
*/

const FirstNonRepeating = (string) => {
	if (string.length <= 0) {
		return ""
	}

	let characterCounts = {};

  for (let i = 0; i < string.length; i++) {
    const char = string[i].toLowerCase();
    if (characterCounts[char]) {
      characterCounts[char]++;
    } else {
      characterCounts[char] = 1;
    }
  }

  for (let i = 0; i < string.length; i++) {
    const char = string[i].toLowerCase();
    if (characterCounts[char] === 1) {
      return string[i];
    }
  }

  return "";
}

console.log(FirstNonRepeating("abba"), "<= Got, Expected: ")
console.log(FirstNonRepeating("~><#~><"), "<= Got, Expected: #")
console.log(FirstNonRepeating("sTreSS"), "<= Got, Expected: T")
console.log(FirstNonRepeating("hello world, eh?"), "<= Got, Expected: w")