def ReverseWords str
	words = str.split(" ")

	words.each_with_index do |word, index|
		words[index] = word.reverse!
	end

	return words.join(" ")

end

puts ReverseWords("The quick brown fox jumps over the lazy dog.")
puts "<= Got, Expecting: ehT kciuq nworb xof spmuj revo eht yzal .god"
