def SquareSum numbers
	sum = 0

	numbers.each do |num|
		sum += (num * num)
	end

	return sum
end

puts SquareSum [0,3,4,5]