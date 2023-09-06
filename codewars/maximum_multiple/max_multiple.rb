	# Given a divisor & a Bound. Find the largest integer N
	# such that :

	# N is divisible by divisor
	# N <= bound
	# n > 0

def MaxMultiple divisor, bound
	numbers = []

	for i in 1..bound
		numbers << i if i % divisor == 0
	end

	return numbers.last
end

puts "#{MaxMultiple(2,7)} <= Got, Expected: 6"
puts "#{MaxMultiple(37, 200)} <= Got, Expected 185"
puts "#{MaxMultiple(7, 17)} <= Got, Expected 14"