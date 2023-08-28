# Program that takes a string input & returns the first character that
# is not repeated anywhere in the string, Case insensitive while checking
# Case sensitive while returning

def FirstNonRepeating str
	# return empty string if the string is empty
	if str.length <= 0
		return ""
	end

	# array to store unique characters
	unique = []
	characters = str.chars

	# create a hash & store the number of times a char occurs
	character_count = Hash.new(0)

	# increment count
	characters.each do |char|
		character_count[char.downcase] += 1
	end

	characters.each do |char|
		if character_count[char.downcase] == 1
			unique << char
		end
	end

	if unique.empty?
		return ""
	end

	return unique.first
end


puts "#{FirstNonRepeating("abba")} <= Got, Expected: "
puts "#{FirstNonRepeating("~><#~><")} <= Got, Expected: #"
puts "#{FirstNonRepeating("sTreSS")}, <= Got, Expexted: T"
puts "#{FirstNonRepeating("hello world, eh?")} <= Got, Expected: w"